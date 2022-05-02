package main

import (
	"errors"
	"fmt"
)

// 测试函数的参数不同
func fun1(a, b int) (c int, err error){
	c = a + b
	if true{
		return c, nil
	}
	return c, errors.New("error")
}

func fun2(a, b uint32) (c int){
	c = int(a) * int(b)
	if true{
		return c
	}
	return c
}

func main(){

	c, _:= fun1(3 ,3)
	fmt.Println(c)


}
