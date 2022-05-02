package main

import "fmt"

func sub() func(){

    i := 10
    fmt.Printf("地址:%p\n", &i)
    b := func(){
        fmt.Printf("i addr %p\n", &i)
        i--
        fmt.Println(i)
    }

    return b
}

func main(){
    
    var f func() = sub()
    f()
    f()
    fmt.Println("over")

}
