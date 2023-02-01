package main

import (
	"encoding/json"
	"fmt"
)

// 结构体 struct 是带类型的字段的集合
type person struct {
	name string
	age  int
}

type student struct {
	person
	id int
}

func main() {
	p1 := person{"Bob", 20}
	fmt.Println(p1)

	// 初始化指定字段名字，省略的字段会初始化为零值
	p2 := person{name: "Alice", age: 20}
	fmt.Println(p2)

	// 使用 . 操作符访问字段
	fmt.Println(p1.name)

	// 对结构体指针使用 . 操作访问字段，自动解引用
	sp := &p1
	fmt.Println(sp.name)

	// 结构体是可变的，指针赋值修改原值
	sp.age = 99
	fmt.Println(sp.age)
	fmt.Println(p1.age)

	// 结构体序列化为 Json
	s2 := student{}
	s2.name = "Bob"
	s2.age = 20
	s2.id = 1
	s2json, e := json.Marshal(s2)
	fmt.Println(e)
	fmt.Println(string(s2json))

}
