package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 模拟工作协程，使用 j, more := <-jobs 接收数据，根据第二个值 more 如果 jobs 已经关闭，该值将是 false
	// done 通道用于通知主线程
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}

	// 关闭一个通道，意味着不能再向这个通道发送值，该特性可以向通道的接收方传达已经完成的信息
	close(jobs)
	fmt.Println("send all jobs")

	// done 通道会阻塞到直到工作线程执行完毕，如果没有这个，程序可能就直接结束了，
	<-done
}
