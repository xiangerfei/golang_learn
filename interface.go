package main

import "fmt"

type Animal interface{
    Say(string)

    Move()
}

type Monkey struct{
    Name string
    Age uint8
}

type Elephonete struct{
    Name string
    Age uint8
    Weight uint32
}

func (m *Monkey) Say(str string){
    fmt.Printf("monkey say:%s %s\n", m.Name, str)
}

func (m Monkey) Move(){
    fmt.Println("monkey move")
}

func (e Elephonete) Say(str string){
    fmt.Printf("el say:%s %s\n", e.Name, str)
}

func (e Elephonete) Move(){
    fmt.Println("el move")
}

func (e Elephonete) Love(){
    fmt.Println("el Love")
}

func interface_daemon(animal Animal){
    animal.Say("hello world")
    animal.Move()
}


func main(){
    m := &Monkey{
        "houzi", 18}

    interface_daemon(m)

    e := &Elephonete{
        "daxiang", 25, 2314}

    e2 := Elephonete{
        "daxiang", 2, 2314}

    interface_daemon(e)
    interface_daemon(e2)
        
}
