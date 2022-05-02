package main

import "fmt"

// 不定长参数的使用


func add2(args ...int)int{
	print(len(args))
	println()
	for i, v := range args{
		fmt.Printf("i: %d v:%d\n", i, v)
	}
	return 1
}

func main(){

	add2(133, 2, 3, 4)

}
