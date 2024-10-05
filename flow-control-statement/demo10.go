package main

import "fmt"

func main() {
CONTINUEDEMO:
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue CONTINUEDEMO
		}
		fmt.Println(i)
	}
	fmt.Println("hhh")
}
