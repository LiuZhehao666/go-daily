package main

import "fmt"

var a1 = 1

// b2出了这个函数就死了，它的生命周期就是test函数的存活时间
func test() {
	b2 := 2
	fmt.Println(b2)
}
func main() {
	var c3 = 3
	fmt.Println(a1, c3) //这里输出b2就会出错
	test()
}
