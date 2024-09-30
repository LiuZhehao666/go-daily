package main

import "fmt"

func main() {
	a, b := 20, 10
	var c int

	c = a + b
	fmt.Printf("a + b = %d\n", c)
	c = a - b
	fmt.Printf("a - b = %d\n", c)
	c = a * b
	fmt.Printf("a * b = %d\n", c)
	c = a / b
	fmt.Printf("a / b = %d\n", c)
	c = a % b
	fmt.Printf("取余 = %d\n", c)
	a++
	fmt.Printf("自增 = %d\n", a)
	a--
	fmt.Printf("自减 = %d\n", a)
}
