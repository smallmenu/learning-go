package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 测试字符串是否符合正则表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 通过 Compile 得到一个优化过的结构体
	compile, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(compile.MatchString("peach"))

	// 查找匹配到的字符串
	fmt.Println(compile.FindString("aaa   peach bbb"))

	// 返回完全匹配和局部匹配的字符串：peach 和 ea
	fmt.Println(compile.FindStringSubmatch("aaa   peach pinch"))

	// 返回全部的匹配项
	fmt.Println(compile.FindAllString("peach punch pinch", -1))
}
