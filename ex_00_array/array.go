package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明 5 个元素的数组，默认会初始化 0 值
	var a [5]int
	fmt.Println(a)

	a[4] = 5
	fmt.Println(a[4])

	// 内置函数 len() 获取数组长度
	fmt.Println(len(a))

	// 声明并初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println(reflect.TypeOf(a))
}
