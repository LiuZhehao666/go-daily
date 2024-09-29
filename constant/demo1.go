package main

import (
	"fmt"
	"unsafe"
)

const (
	a = ""
	b = len(a)           //长度
	c = unsafe.Sizeof(a) //字节数
)

func main() {
	fmt.Println(a, b, c)
}
