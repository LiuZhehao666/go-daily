package main

import "fmt"

func main() {
	if score := 80; score > 90 { //初始化写在条件语句内
		fmt.Println("优秀")
	} else if score > 75 {
		fmt.Println("良好")
	} else if score > 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("继续加油！")
	}
	//fmt.Println(score) //会报错
}
