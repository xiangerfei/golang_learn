package main

import (
	"fmt"
	"strings"
)

const SIZE2 = 10

// slice 是值拷贝传递，其中
func func1(s *[]int){
	(*s)[1] = 100
	 fmt.Printf(strings.Repeat("*", 30))
	var cap1 int = cap(*s)
	fmt.Printf("s cap:", cap(*s))
	for i:= 0; i < 100; i++{
		(*s) = append((*s), i)
		if cap1 != cap(*s){
			fmt.Printf("new capacity: %d\n", cap(*s))
		}
	}

}

// 一旦cap超过了值，就会发生分离。
func func2(s []int){
	s[0] = 1000
	for i:=3; i < SIZE2*2; i++{
		s = append(s, i)
	}
	fmt.Printf("-----cap: %d \n", cap(s))
	fmt.Printf(strings.Repeat("*", 30))
	for i, value := range s{
		fmt.Printf("key: %d, value: %d \n", i, value)
	}

}

func main(){
	/* 验证引用传递
	s1 := []int{1, 2, 3, 4}

	func1(&s1)
	println()

	for i,value := range s1{
		println("key: %d, value: %d", i, value)
	}
	*/


	s2 := make([]int, SIZE2, SIZE2*2)
	fmt.Printf("cap: %d \n", cap(s2))
	fmt.Printf("len: %d \n", len(s2))
	func2(s2)


	// 这个时候发生了什么？
	for i, value := range s2{
		fmt.Printf("key: %d, value: %d \n", i, value)
	}



}
