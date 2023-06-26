package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WithCancel()

	WithDeadline()
}

func WithCancel() {
	// WithCancel 创建一个可以取消的 context，可以在 goroutine 中使用该 context 来控制 goroutine 的生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine canceled")
				return
			default:
				fmt.Println("goroutine running")
			}

			time.Sleep(time.Second * 1)
		}
	}()

	// 模拟操作，等待两秒
	time.Sleep(time.Second * 5)

	// 通知子协程退出，调用 cancel 函数时，goroutine会收到取消通知并结束执行
	cancel()

	// 延迟退出主协程，为了看到通知的信息
	time.Sleep(time.Second * 1)
}

func WithDeadline() {
	// 这个例子与上面的例子类似，只是使用 WithDeadline 创建了一个带有截止时间的 context
	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine deadline canceled")
				return
			default:
				fmt.Println("goroutine running")
			}

			time.Sleep(time.Second * 1)
		}
	}()

	// 模拟一些操作，比截止时间更长
	time.Sleep(5 * time.Second)
}
