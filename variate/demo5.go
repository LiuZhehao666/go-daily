package main

import "fmt"

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2

func main() {
	e := "hello" //只能再函数体中使用
	fmt.Println(x, y, a, b, c, d, e)
}
