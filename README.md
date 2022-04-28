# golang_learn
开始系统学习golang

以下为网课笔记

#### go程序入口
1. package main 包里面的main函数

#### go命令介绍
```
1. go help      帮
2. go build     编译
3. go install   安装
4. go test      测试
5. go run       编译且运行
6. go mod 
    1. go mod init  
    2. go mod tidy
7. go tool 
    1. go tool pprof
    2. go tool trace
8. go vet       检查静态代码错误
9. go fmt       格式化代码
10. go doc      得到包的帮助
11. go get      获取包
```


#### 语法介绍

1. 关键字
```
break    default     func   interface   select
case     defer       go     map         struct
chan     else        goto   package     switch
const    if          range  type        continue
for      import      return fallthrough var

iota 
```
2. 基本类型
```
1. 整形
2. 浮点型
3. 布尔
rune 4个字节，可以表示一个中文
byte = uint8    
```
3. 变量申明
```
var a float64 = 43 等价于 var a = 64.0 等价于 a := 64 （自动类型推导, 语法糖） 函数内部这样使用
```

4. 变量作用域
```
// 大小写决定可见度
var (
    A = 3
    b = 2
)
A 可通过package.A 被其他模块访问
b 不可行。只能自己的package访问
```

5. 复合数据类型
```
array   值类型
struct  值类型
string
slice   引用类型
map     引用类型
chnanel 引用类型
interface 接口
function   函数
```



#### 系统库
Strings
```
var s string = "I am cool man"
	var arr_string []string= strings.Split(s, " ")
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
	s4 := s1 + s2 + s3 // 第一种拼接
	s5 := fmt.Sprintf("%s %s %s", s1, s2, s3) //第二种
	s6 := strings.Join([]string{s1, s2, s3}, " ") // 第三种,效率高
	s7 := strings.Builder{}					// 第四种，最快
	s7.WriteString(s1)
	s7.WriteString(" ")
	s7.WriteString(s2)
	s7.WriteString(" ")
	s7.WriteString(s3)
	fmt.Printf("%s\n, %s\n, %s\n, %s", s4, s5, s6, s7.String())
```

#### 数组
1. 必须指定长度
2. 数据的操作
```
arrar_1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d\n" ,arrar_1)
	for index, ele := range arrar_1{
		fmt.Printf("index: %d, value:%d\n",index, ele)
	}
	fmt.Printf("数组的capacity 和 length")
	fmt.Printf("cap:%d \t len:%d", cap(arrar_1), len(arrar_1))

```


#### 结构体
```
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

```

#### 切片

1. 切片的结构如下，可以看出和数组的首地址不是一个概念
```
type slice struct{
    array unsafe.Pointer
    len int
    cap int
}
```
2. 初始化
```
var s []int // 申明 len = cap = 0
s = []int{} // 初始化 len = cap = 0
s = make([]int, 3) // 初始化, len = cap = 3
s = make([]int, 3, 5) // 初始化, len = 3, cap = 5
s = []int{1, 2, 3, 4, 5} // 初始化，len = cap = 5
s2d := [][]int{
    {1}, {2, 3}, // 二维数组各行的列数相等，二维切片各行的len可以不等。
}

```

3. 操作
```
append
copy
```

4. 切片会影响底层的引用地址。
