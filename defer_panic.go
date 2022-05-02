package main


import "fmt"

// 测试defer 中发生了panic ,对其他的defer影响
func defer_panic_t(){

    fmt.Println("hello world 1")
    defer fmt.Println("hello world defer 2")

    n := 0
    defer func(){
//        panic(2/0)
        fmt.Println(2/n)
        fmt.Println("hello defer panic 3")
    }()
    
    defer fmt.Println("hello defer_panic_t 4")
}


func main(){
    
    defer_panic_t()

}
