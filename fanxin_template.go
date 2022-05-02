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
