package main

import "fmt"

func main() {
	const (
		a = iota    //开始计数，从 0 开始
		b           //1
		c           //2
		d = "hello" //3，hello
		e           //4，相当于 e = "hello"
		f           //5，相当于 f = "hello"
		g = 1       //6，1
		h           //7，相当于 h = 1
		i           //8，相当于 i = 1
		j = iota    // 恢复计数，9
		k           //10
		l           //11
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
}
