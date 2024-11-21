package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个 sync.Map 实例
	var safeMap sync.Map
	var wg sync.WaitGroup

	// 启动多个协程写入数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 使用 Store 方法将键值对存入 map
			safeMap.Store(fmt.Sprintf("key%d", i), i)
			fmt.Printf("Stored key%d: %d\n", i, i)
		}(i)
	}

	// 启动多个协程读取数据
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 使用 Load 方法读取数据
			if value, ok := safeMap.Load(fmt.Sprintf("key%d", i)); ok {
				fmt.Printf("Loaded key%d: %d\n", i, value)
			} else {
				fmt.Printf("key%d does not exist\n", i)
			}
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 演示 Delete 方法
	fmt.Println("Deleting key5")
	safeMap.Delete("key5")

	// 再次尝试读取 key5
	if value, ok := safeMap.Load("key5"); ok {
		fmt.Printf("Loaded key5: %d\n", value)
	} else {
		fmt.Println("key5 has been deleted")
	}
}
