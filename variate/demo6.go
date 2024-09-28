package main

import "fmt"

var str2 = "world" //可以不使用

func main() {
	str1 := "abc"
	fmt.Println("hello", str1) //必须使用str1，否则会报错
}
