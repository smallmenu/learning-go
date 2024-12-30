package main

import (
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	currentPID := os.Getpid()
	fmt.Printf("当前进程的 PID: %d\n", currentPID)

	// 获取当前进程的名称
	currentProcess, err := process.NewProcess(int32(currentPID))
	if err != nil {
		fmt.Println("获取当前进程失败:", err)
		return
	}
	currentName, err := currentProcess.Name()
	if err != nil {
		fmt.Println("获取进程名称失败:", err)
		return
	}
	fmt.Printf("当前进程的名称: %s\n", currentName)

	// 列出所有进程并检查当前进程
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("获取进程失败:", err)
		return
	}

	instanceCount := 0
	for _, p := range processes {
		pid := p.Pid
		_, err := p.Name()
		if err != nil {
			continue
		}
		if pid == int32(currentPID) {
			instanceCount++
		}
		//fmt.Printf("PID: %d, 名称: %s\n", pid, name)
	}

	// 判断当前进程实例是否唯一
	if instanceCount > 1 {
		fmt.Println("当前进程实例不是唯一的，其他实例正在运行。")
	} else {
		fmt.Println("当前进程实例是唯一的。")
	}

	time.Sleep(time.Duration(10) * time.Second)
}
