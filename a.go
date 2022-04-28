package main

import (
	"fmt"
	"github.com/grd/statistics"
	"gofirst/mypath" // 这里gofirst用的是go mod init gofirst里面mod
	"gonum.org/v1/gonum/stat"
	"strconv"
	"strings"
)

func main() {

	b := 7
	c := 7
	d := mypath.Add(b, c)
	fmt.Println("hello world")
	fmt.Println(d)
	arr := []float64{1, 3, 4, 6, 7, 8}
	v := stat.Variance(arr, nil)
	fmt.Println("方差:%f", v)

	data := statistics.Int64{1, 2, 3, 4, 5}
	//方差计算
	variance := statistics.Variance(&data)
	//输出结果(2.5)
	fmt.Println(variance)

	var s string = "I am cool man"
	var arr_string []string = strings.Split(s, " ")
	for _, ele := range arr_string {
		fmt.Printf("%s\t", ele)
	}
	println()

	println(strings.Repeat("*", 30), "strings的用法", strings.Repeat("*", 30))
	fmt.Printf("查找\n")
	fmt.Printf(`是否以'man'结尾:`)
	println(strings.HasSuffix(s, "man"))
	fmt.Printf(`是否以'I'开头:`)
	println(strings.HasPrefix(s, "I"))
	fmt.Printf("第一个出现的am, 下标为%d\n", strings.Index(s, "am"))

	s1 := "hi"
	s2 := "Hello"
	s3 := "world"
	s4 := s1 + s2 + s3                            // 第一种拼接
	s5 := fmt.Sprintf("%s %s %s", s1, s2, s3)     //第二种
	s6 := strings.Join([]string{s1, s2, s3}, " ") // 第三种,效率高
	s7 := strings.Builder{}                       // 第四种，最快
	s7.WriteString(s1)
	s7.WriteString(" ")
	s7.WriteString(s2)
	s7.WriteString(" ")
	s7.WriteString(s3)
	fmt.Printf("%s\n, %s\n, %s\n, %s", s4, s5, s6, s7.String())


	fmt.Println("\n", strings.Repeat("*", 30), "类型转换", strings.Repeat("*", 30), "\n" )
	var s_zh string = "1234123"
	i_zh, error := strconv.Atoi(s_zh)
	if error == nil {
		println(i_zh)
	}else{
		println("转换有问题")
	}

	println()
	println()
	fmt.Println("\n", strings.Repeat("*", 30), "数组", strings.Repeat("*", 30), "\n" )

	arrar_1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d\n" ,arrar_1)
	for index, ele := range arrar_1{
		fmt.Printf("index: %d, value:%d\n",index, ele)
	}
	fmt.Printf("数组的capacity 和 length")
	fmt.Printf("cap:%d \t len:%d", cap(arrar_1), len(arrar_1))


	type User struct {
		name string
		age uint8
	}

	println()

	var user User =  User{
		"yixiang",
		18}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)

	user_arr := []User{
		{"zhao", 18},
		{"zhao2", 19},
		{"zhao4", 20}}

	for index, u := range user_arr{
		fmt.Printf("%d, name: %s, age: %d", index, u.name, u.age)
		println()
	}
	// 修改

	
}
