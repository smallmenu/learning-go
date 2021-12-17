package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	close(queue)

	// range 可以遍历通道取值
	// 所以，一个非空的通道是可以关闭的，并且，通道中的值仍然可以被接收到
	for q := range queue {
		fmt.Println(q)
	}
}
