package main

import "fmt"

func main() {
	score := 80
	if score > 90 {
		fmt.Println("优秀")
	} else if score > 75 {
		fmt.Println("良好")
	} else if score > 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("继续加油！")
	}
}
