package main

import (
	"encoding/json"
	"fmt"
	"strings"

	machineCode "github.com/super-l/machine-code"
)

func main() {
	if machineCode.MachineErr != nil {
		fmt.Println("获取机器码信息错误:" + machineCode.MachineErr.Error())
		return
	}

	machineJson, _ := json.Marshal(machineCode.Machine)
	fmt.Println("机器码信息汇总:" + string(machineJson))

	trafficIp, _ := machineCode.GetIpAddr()
	fmt.Println("当前活跃IP：" + trafficIp)

	allIp, _ := machineCode.GetIpAddrAll()
	fmt.Println("所有IP：" + strings.Join(allIp, " "))
}
