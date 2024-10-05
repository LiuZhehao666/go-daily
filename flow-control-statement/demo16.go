package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto flag
		}
		fmt.Println(i)
	}
	return

flag:
	fmt.Println("结束循环")
}
