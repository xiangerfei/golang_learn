package main



import (
    "fmt"
    "errors"
    )



func panic_recover(){

    
    defer func(){

        if err := recover(); err !=nil {
            
            fmt.Println("Some thing error")
        }

    }()

    panic(errors.New("Occor an error"))

}

func main(){
    panic_recover()
    fmt.Println("error happend\n")
}
