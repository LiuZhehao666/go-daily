package main

import "fmt"

func main() {
	age := 50
	switch { //switch后面不跟变量
	case age > 60:
		fmt.Println("老年人")
	case age < 60:
		fmt.Println("劳动者")
	default:
		fmt.Println("不是人")
	}
}
