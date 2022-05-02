package main

import "fmt"

func defer_t() (i int) {
	i = 9
	defer func() {
		fmt.Printf("func novar i: %d\n", i)
		i = 4 // 在return之前改变了i 
	}()

	defer func(t int) { //9
        
		fmt.Printf("func var i: %d\n", i) //5 最后计算, 赋值成了5
		fmt.Printf("func var t: %d\n", t)
		i = 4
	}(i)    // 9. 参数拷贝，赋值

	defer fmt.Printf("fmt.Println i: %d\n", i) // 9 拷贝赋值
	i = 5
	return // return的 i = 4
}

func defer_t1(){

	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println(2)
	fmt.Println("-------------")

}

func main() {

	i := defer_t()
	fmt.Printf("main return i: %d\n", i)
}
/*
fmt.Println i: 9
func var i: 5
func var t: 9
func novar i: 4
main return i: 4
*/
