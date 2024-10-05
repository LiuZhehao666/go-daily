package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch num {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("啥也不是")
	}
}
