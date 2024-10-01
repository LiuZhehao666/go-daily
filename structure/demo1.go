package main

import "fmt"

type Student struct {
	Name string
	ID   int
	Age  int
}

func main() {
	/*st := Student{
		//属性: 值
		Name: "jack", //需要尾随逗号
		ID:   123456,
		Age:  18,
	}*/
	st := Student{
		//值
		"jack",
		123456,
		18,
	}
	fmt.Println(st)
	fmt.Println(st.Name)
}
