#### sync

提供了同步原语


- sync.Mutex
互斥锁
```
type Mutex struct {
    state int32 // 状态
    sema  uint32 // 信号量
}
```
1. state 有32位
```
00000000 00000000 00000000 00001111
最后一位 mutexLocked  是否被锁定 
倒数第二位 mutexWoken  正常模式下被唤醒为 1
倒数第三位 mutexStarving 表示进入了饥饿状态 1
剩下的  表示等待有多少个goroutine等待互斥锁
```
2. 方法
```
m : &sync.Mutex{}

m.Lock()
m.Unlock()
```

- sync.RWMutex
读写锁
```
type RWMutex struct {
    w           Mutex
    writerSem   uint32 // 写等待
    readerSem   uint32 // 读等待
    readerCount int32  // 当前读的数量
    readerWait  int32  // 写操作被堵塞时读操作数
}
// 当读写锁被释放 Unlock 时首先会通知所有的读操作，然后才会释放持有的互斥锁，
// 保证读操作不会被连续的写操作『饿死』
```

方法
```
rw := &RWMutex{}

rw.RLock()
rw.RUnlock()

rw.Lock()
rw.Unlock()

```


- sync.WaitGroup

用于等待一系列的 Goroutine 的返回
```
type WaitGroup struct {
    noCopy noCopy
    state1 [3]uint32 // 存储当前结构体持有的状态和信号量
}
```
方法
```
Add // 当调用 Add 方法导致计数器归零并且还有等待的 Goroutine 时，
    // 就会通过 runtime_Semrelease 唤醒处于等待状态的所有 Goroutine。
Wait // Wait 会在当前计数器中保存的数据大于 0 时修改等待 
     // Goroutine 的个数 waiter 并调用 runtime_Semacquire 陷入睡眠状态。


Done(实际上就是Add(-1))
```

```
requests := []*Request{...}

wg := &sync.WaitGroup{}
wg.Add(len(requests))

for _, request := range requests {
    go func(r *Request) {
        defer wg.Done() // 这里会 -1
        
        // res, err := service.call(r)
    }(request)
}

wg.Wait()

```

3. 注意：
    1. Add 不能在和 Wait 方法在 Goroutine 中并发调用，一旦出现就会造成程序崩溃；
    2. WaitGroup 必须在 Wait 方法返回之后才能被重新使用；
    3. Done 只是对 Add 方法的简单封装，我们可以向 Add 方法传入任意负数（需要保证计数器非负）快速将计数器归零以唤醒其他等待的 Goroutine；
    4. 可以同时有多个 Goroutine 等待当前 WaitGroup 计数器的归零，这些 Goroutine 也会被『同时』唤醒；


- Once

```
type Once struct {
    done uint32 // 是否被执行过
    m    Mutex  // 锁
}
```

Once 结构体对外唯一暴露的方法就是 Do
```
func main() {
    o := &sync.Once{}
    for i := 0; i < 10; i++ {
        o.Do(func() {
            fmt.Println("only once")
        })
    }
}

$ go run main.go
only once
```


- Cond

通过 Cond 我们可以让一系列的 Goroutine 都在触发某个事件或者条件时才被唤醒，每一个 Cond 结构体都包含一个互斥锁 L

```
func main() {
    c := sync.NewCond(&sync.Mutex{})

    for i := 0; i < 10; i++ {
        go listen(c)
    }

    go broadcast(c)

    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt)
    <-ch
}

func broadcast(c *sync.Cond) {
    c.L.Lock()
    c.Broadcast()
    c.L.Unlock()
}

func listen(c *sync.Cond) {
    c.L.Lock()
    c.Wait()
    fmt.Println("listen")
    c.L.Unlock()
}

```
