package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	jsonStr := `
		{
		  "name": {"first": "Tom", "last": "Anderson"},
		  "age":37,
		  "children": ["Sara","Alex","Jack"],
		  "fav.movie": "Deer Hunter",
		  "friends": [
			{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
			{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
			{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		  ]
		}
	`

	// gjson 提供了快速从 JSON 文档获取值的方法
	fmt.Println(gjson.Get(jsonStr, "name.first"))

	children := gjson.Get(jsonStr, "children.#")
	fmt.Println(children)
	fmt.Println(children.Type)
}
