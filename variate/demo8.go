package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Println(a, b)
	_, b = 1, 2
	fmt.Println(a, b)
}
