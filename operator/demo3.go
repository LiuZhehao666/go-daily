package main

import "fmt"

func main() {
	var a int = 1
	var p *int

	p = &a          //p中存的是a的地址
	fmt.Println(p)  //输出的是a的地址
	fmt.Println(*p) //*p是通过a的地址找到a的内容
}
