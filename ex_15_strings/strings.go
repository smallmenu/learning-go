package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello，中国"
	strs := []string{"a", "b", "c"}
	strInt := "123456"
	ints := 123456

	// 字符串比较
	if str == "Hello，中国" {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	isContains := strings.Contains("test", "es")
	splits := strings.Split("a,b, c", ",")
	join := strings.Join(strs, "-")
	trimSpace := strings.TrimSpace("    		abc ")

	// 字符长度
	len := len(str)
	lenByte := utf8.RuneCountInString(str)

	// 字符串转数值
	// Atoi() 等同于 ParseInt(str, 10, 0)，更简洁
	atoi, _ := strconv.Atoi(strInt)
	parseInt, _ := strconv.ParseInt(strInt, 10, 64)
	// 数值转字符串
	formatInt := strconv.FormatInt(int64(ints), 10)

	fmt.Println("Type:		", reflect.TypeOf(strs))
	fmt.Println("Contains:	", isContains)
	fmt.Println("Split:	", splits)
	fmt.Println("Join:		", join)
	fmt.Println("Len:		", len)
	fmt.Println("LenByte:	", lenByte)
	fmt.Println("TrimSpace:", trimSpace)
	fmt.Println("Atoi:	", reflect.TypeOf(parseInt))
	fmt.Println("Atoi:	", atoi)
	fmt.Println("ParseInt:	", reflect.TypeOf(parseInt))
	fmt.Println("ParseInt:	", parseInt)
	fmt.Println("FormatInt:	", reflect.TypeOf(formatInt))
	fmt.Println("FormatInt:	", formatInt)

	// 使用 ` 定义字符串，避免转义双引号
	strBytes := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Println(reflect.TypeOf(strBytes))
	fmt.Println(strBytes)
}
