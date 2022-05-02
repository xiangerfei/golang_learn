package main

import (
	"fmt"
	"time"
)

func makeTimeout(ch chan bool, t int) {
	time.Sleep(time.Second * time.Duration(t))
	ch <- true
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	timeout := make(chan bool, 1)

	go makeTimeout(timeout, 2)

	select {
	case msg1 := <-c1:
		fmt.Println("c1 received: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2 received: ", msg2)
	case <-timeout:
		fmt.Println("Timeout, exit.")
	}
}