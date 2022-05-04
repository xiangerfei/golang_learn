package main

import (
	"fmt"
	"time"
)

func main(){
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Microsecond * 1)
	}
	time.Sleep(time.Second * 3)
}