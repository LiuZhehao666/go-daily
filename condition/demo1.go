package main

//if 语句基本写法
import "fmt"

func main() {
	var score = 80
	if score > 90 {
		fmt.Println("优秀")
	} else if score > 75 {
		fmt.Println("良好")
	} else if score > 60 {
		fmt.Println("合格")
	} else {
		fmt.Println("继续加油！")
	}
	fmt.Println(score)
}
