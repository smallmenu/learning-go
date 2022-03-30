package main

import "fmt"

type rect struct {
	width, height int
}

// area() 是一个拥有 *rect 类型的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 还可以通过值类型定义方法
func (r rect) perim() int {
	return (r.width + r.height) * 2
}

func main() {
	r := rect{width: 10, height: 20}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go 会自动处理值和指针之间的转换，想要避免调用方法时产生一个拷贝。都可以使用指针调用方法
	pr := &r
	fmt.Println("area:", pr.area())
	fmt.Println("perim:", pr.perim())

}
