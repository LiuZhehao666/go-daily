package main

import "fmt"

func main() {
	var a string = "Hello" //声明变量类型
	fmt.Println(a)

	var b int //没有初始化则为零值
	fmt.Println(b)

	var c bool //bool的默认值是false
	fmt.Println(c)
}
