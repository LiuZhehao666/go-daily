package main

import "fmt"

func main() {
	a, b, c, d := 20, 10, 15, 5
	var e int

	e = (a + b) * c / d
	fmt.Println(e)
	e = a + (b*c)/d
	fmt.Println(e)
	e = (a + b) * (c / d)
	fmt.Println(e)
	e = ((a + b) * c) / d
	fmt.Println(e)
}
