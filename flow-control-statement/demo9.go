package main

import "fmt"

func main() {
BREAKDEMO:
	for i := 0; i < 10; i++ {
		if i == 3 {
			break BREAKDEMO
		}
		fmt.Println(i)
	}
	fmt.Println("hhh")
}
