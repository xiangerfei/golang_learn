package main

import "fmt"

const SIZE = 4

type User struct {
	id uint64
	age uint8
	name, addr, desc string
}

func (self User) Hello(str string){
	fmt.Println(str, self.name)
}

func add(a, b interface{}){
	switch value := a.(type) {
	case int32:
		println("int32")
	case int64:
		println("int64")
	case int:
		print("int")
		fmt.Printf("%d %d", a, b)
	default:
		print("不是")
		fmt.Printf("%T", value)
	}
}


func main(){

	var user User = User{
		1, 18, "yixiang", "anhui", "cool man"}
	user.Hello("hello")

	fmt.Printf("ces%d",SIZE)

	add(12, 32)
}