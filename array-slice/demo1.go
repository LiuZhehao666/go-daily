package main

import "fmt"

func main() {
	//数组初始化
	var strAry = [10]string{"aa", "bb", "cc", "dd", "ee"}
	//切片初始化
	//[]string代表切片元素的类型
	//0 表示这个切片初始化的长度
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:4]
	//字典初始化
	var dic = map[string]int{"apples": 1, "oranges": 2, "watermelon": 3}

	fmt.Println(strAry)
	fmt.Println(sliceAry)
	fmt.Println(dic)
}
