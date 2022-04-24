package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	f, err := os.Create(p)

	// 当函数返回我们不知道如何处理的错误时，使用 panic 中止操作
	// 以非 0 状态退出程序
	// 注意： Go 中通常会尽可能的使用返回值来标示错误
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")

	// defer 用于确保程序在实行完成后，会调用某个函数，和 Java 的 finally 类似
	defer closeFile(f)

	writeFile(f)
}
