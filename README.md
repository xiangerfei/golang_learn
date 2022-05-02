# golang_learn
开始系统学习golang

以下为网课笔记

#### go程序入口
1. package main 包里面的main函数

#### go命令介绍
```
1. go help      帮
2. go build     编译
3. go install   安装
4. go test      测试
5. go run       编译且运行
6. go mod 
    1. go mod init  
    2. go mod tidy
7. go tool 
    1. go tool pprof
    2. go tool trace
8. go vet       检查静态代码错误
9. go fmt       格式化代码
10. go doc      得到包的帮助
11. go get      获取包
```


#### 语法介绍

1. 关键字
```
break    default     func   interface   select
case     defer       go     map         struct
chan     else        goto   package     switch
const    if          range  type        continue
for      import      return fallthrough var

iota 
```
2. 基本类型
```
1. 整形
2. 浮点型
3. 布尔
rune 4个字节，可以表示一个中文
byte = uint8    
```
3. 变量申明
```
var a float64 = 43 等价于 var a = 64.0 等价于 a := 64 （自动类型推导, 语法糖） 函数内部这样使用
```

4. 变量作用域
```
// 大小写决定可见度
var (
    A = 3
    b = 2
)
A 可通过package.A 被其他模块访问
b 不可行。只能自己的package访问
```

5. 复合数据类型
```
array   值类型
struct  值类型
string
slice   引用类型
map     引用类型
chnanel 引用类型
interface 接口
function   函数
```



#### 系统库
Strings
```
var s string = "I am cool man"
	var arr_string []string= strings.Split(s, " ")
	for _, ele := range arr_string {
		fmt.Printf("%s\t", ele)
	}
	println()

	println(strings.Repeat("*", 30), "strings的用法", strings.Repeat("*", 30))
	fmt.Printf("查找\n")
	fmt.Printf(`是否以'man'结尾:`)
	println(strings.HasSuffix(s, "man"))
	fmt.Printf(`是否以'I'开头:`)
	println(strings.HasPrefix(s, "I"))
	fmt.Printf("第一个出现的am, 下标为%d\n", strings.Index(s, "am"))

	s1 := "hi"
	s2 := "Hello"
	s3 := "world"
	s4 := s1 + s2 + s3 // 第一种拼接
	s5 := fmt.Sprintf("%s %s %s", s1, s2, s3) //第二种
	s6 := strings.Join([]string{s1, s2, s3}, " ") // 第三种,效率高
	s7 := strings.Builder{}					// 第四种，最快
	s7.WriteString(s1)
	s7.WriteString(" ")
	s7.WriteString(s2)
	s7.WriteString(" ")
	s7.WriteString(s3)
	fmt.Printf("%s\n, %s\n, %s\n, %s", s4, s5, s6, s7.String())
```

#### 数组
1. 必须指定长度
2. 数据的操作
```
arrar_1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d\n" ,arrar_1)
	for index, ele := range arrar_1{
		fmt.Printf("index: %d, value:%d\n",index, ele)
	}
	fmt.Printf("数组的capacity 和 length")
	fmt.Printf("cap:%d \t len:%d", cap(arrar_1), len(arrar_1))

```


#### 结构体
```
type User struct {
		name string
		age uint8
	}

	println()

	var user User =  User{
		"yixiang",
		18}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)

	user_arr := []User{
		{"zhao", 18},
		{"zhao2", 19},
		{"zhao4", 20}}

	for index, u := range user_arr{
		fmt.Printf("%d, name: %s, age: %d", index, u.name, u.age)
		println()
}

```

#### 切片

1. 切片的结构如下，可以看出和数组的首地址不是一个概念
```
type slice struct{
    array unsafe.Pointer
    len int
    cap int
}
```
2. 初始化
```
var s []int // 申明 len = cap = 0
s = []int{} // 初始化 len = cap = 0
s = make([]int, 3) // 初始化, len = cap = 3
s = make([]int, 3, 5) // 初始化, len = 3, cap = 5
s = []int{1, 2, 3, 4, 5} // 初始化，len = cap = 5
s2d := [][]int{
    {1}, {2, 3}, // 二维数组各行的列数相等，二维切片各行的len可以不等。
}

```

3. 操作
```
append
copy
```

4. 切片会影响底层的引用地址。


#### map

1. map的特性
- map就是一个hashmap
- 如果有冲突拉一个链
- 如果冲突过多，需要扩容，避免O(1)复杂度变成O(n)。
2. map操作
```
var m map[string] int

m = make(map[string]int)

m = make(map[string]int, 5) // 容量为5，建议设置一个容量初始值，减少扩容概率。

m = map[string]int{"语文":0, "数学": 100}  // 初始化直接赋


m["英语"] = 98 // 增加
m["英语"] = 88 // 修改
delete(m["英语"]) // 删除
len(m) //长度
cap(m) // 不支持，不需要


// 判断是否存在
if value, exists := m["英语"]; exists {
    fmt.Println(value)
}else {
    fmt.Println("not exists")
}


// 遍历
for key, value := range m{
    fmt.Printf("key: %s, value: %d", key, value);
}
```

#### channel

1. chan 是一个引用类型
2. chan 可用于多goroutine 之间的通信
3. chan 操作

```
func main() {

	// 构建一个通道
	ch := make(chan int)

	// 申明一个同步组
	var wg sync.WaitGroup

	// 组里面增加一个
	wg.Add(1)
	// 开启一个并发匿名函数
	go func() {

		fmt.Println("start goroutine")

		// 通过通道通知main的goroutine
		ch <- 0

		//  wg.Down()作用是等待子进程结束
		fmt.Println("exit goroutine")
		wg.Add(-1) // 等价 wg.Down()

	}()

	time.Sleep(10000000)
	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch
	wg.Wait()
	fmt.Println("all done")

}
```


#### 结构体
```
type User struct {
	id uint64
	age uint8
	name, addr, desc string
}
// 也可以申明匿名结构体，C++也有。

func (self User) Hello(str string){
	fmt.Println(str, self.name)
}

func main(){

	var user User = User{
		1, 18, "yixiang", "anhui", "cool man"}
	user.Hello("hello")
}
```


#### 指针
```
ptr1 := *[]int
ptr2 := []*int
```


#### 控制语句
```
if 语句
switch 语句 (case default fallthrough)
for 循环 break continue
goto 语句与label
```


##### 不定长参数
```
func f(args ...int){

}

```

#### 异常处理
1. go没有类型try catch 类的异常处理。
2. 使用返回error 的方式来处理异常。



#### defer
1. 延迟调用
2. 函数返回前调用
3. 先注册的后处理
4. 函数体里面的变量后调用时赋值，语句的正常赋值。
5. 一般用于关闭文件，数据库关闭等操作。 

#### panic 和 recover
1. panic 产生一个错误。会打印函数堆栈后，exit(2)
2. 使用recover 捕捉panic，让函数正常执行, recover 必须放在defer里面
```

func panic_recover(){


    defer func(){

        if err := recover(); err !=nil {

            fmt.Println("Some thing error")
        }

    }()

    panic(errors.New("Occor an error"))

}

func main(){
    panic_recover()
    fmt.Println("error happend\n")
}
```


#### 接口
1. 一组方法的集合
2. 结构的值由两部分组成
    1. 一个指向具体的实现
    2. 指向具体实例的字段数据

3. 接口嵌入，类似结构体嵌入
```
代码见 interface.go
```



#### 类型断言
1. i.(type) 只能用在switch 后面
2. i.(int) 整体

#### 泛型

1. 类似模版函数
2. 1.17测试阶段，1.18可以试试

代码大概如下面
```
package main

import "fmt"

type Addable interface{
   type int, float64, int32, uint8, float32, int64}

func Add[T Addable](a, b T) T{
    return a + b
}

func main(){

    fmt.Println(Add(3.2, 4.2))
    fmt.Println(Add(3, 4))
}
```


#### 反射

```
1. 反射是在运行时期（不是编译期间）探知对象的类型信息、内存结构、更新变量、调用他们的方法 
2. api的用法
3. json的虚拟化和反虚拟化


```





#### 工程化

1. go mod init xiangerfer.com/db_flagship
2. go mod tidy
3. GOROOT  GOPATH 当前目录
4. package package_name 
    1. package_name 不必需要和目录一样
    2. 同一个目录下面的所有包名需要一样。
    3. 可以直接使用同包下面的变量，函数，结构体
    4. 不同包下需要import对应的包
    5. go文件的第一行声明包名。
5. 在包的申明砂上面写包的注释，也可以单独写在doc.go里面
6. 包里面的init()函数，在引入时被执行，在maim之前执行。
7. internal 包

```
import asd "go-course/package" // 别名

asd.Add()
```


#### 标准库
- time
    1. time.NewTicker(3 * time.Second) 周期性任务
    2. time.NewTimer(3 * time.Second)  定时任务
    3. time.After(3 * time.Second) //
- 数学库
- io 库
    1. 标准输入
    ```
    var s string
    fmt.Scan(&s)

    var i int
    fmt.Scanf("%d", &)
    ```
    2. bufio
    3. io
    4. io/ioutils
    5. log
    6. os
    7. path/filepath
    8. exec
#### 编码


```

```
