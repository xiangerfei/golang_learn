package main

import (
	"fmt"
	"sync"
	"time"
)

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
