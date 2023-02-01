package main

import (
	"fmt"
)

func main() {

	// 基本的 switch，默认分支是 default
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 类型开关比较类型而非值，可以用来发现一个接口值的类型。
	// 变量 t 在每个分支中会有相应的类型。
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("don't know", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("abc")
}
