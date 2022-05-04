/*
	go:通过通信来实现共享内存，而不是通过共享内存来通信
*/
package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

type Request struct {
	id   string
	name string
}

func TWaitGroup() {

	//requests := []*Request{}

	wg := &sync.WaitGroup{}
	wg.Add(1)

}

func TOnce() {

	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}

}
func main() {
	//go Balance()
	//m := sync.Mutex{}
	//m.Lock()
	//m.Unlock()
	//atomic.CompareAndSwapInt32()
	//
	//const (
	//	mutexLocked = 1 << iota // mutex is locked
	//	mutexWoken
	//	mutexStarving
	//	mutexWaiterShift = iota
	//)
	//fmt.Println(mutexLocked)
	//fmt.Println(mutexWoken)
	//fmt.Println(mutexStarving)
	//fmt.Println(mutexWaiterShift)

	//TWaitGroup()

	//TOnce()


}

