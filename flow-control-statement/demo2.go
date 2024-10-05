package main

import "fmt"

// if 特殊写法
func main() {
	if score := 80; score > 90 {
		fmt.Println("优秀")
	} else if score > 75 {
		fmt.Println("良好")
	} else if score > 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("继续加油！")
	}
}
