package main

import (
	"fmt"
	"strings"
)

func print_map(m map[string]int){
	for k, v := range m{
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	print(strings.Repeat("*", 30), "\n")
}

func main()  {

	m := map[string]int{"Chinese": 100, "English": 88}

	// 判断是否存在
	if value, exists := m["Chinese"]; exists{
		fmt.Printf("%d\n", value)
	}else {
		print("not exists\n")
	}
	print_map(m)

	// 增加
	m["math"] = 98
	print_map(m)

	// 修改
	m["Chinese"] = 90
	print_map(m)

	// 删除
	delete(m, "English")
	print_map(m)

}
