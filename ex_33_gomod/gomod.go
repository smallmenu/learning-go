package main

// `go mod tidy` 会自动拉取缺少的模块，移除不用的模块
import (
	"fmt"

	// 来自远程仓库
	"rsc.io/quote"

	// 使用 go mod edit -replace 修改为本地系统模块
	"github.com/learning-go/greetings"

	"log"
)

func main() {
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")

	// 调用本地系统模块的函数
	message, err := greetings.Hello("test")

	// Fatal() 会导致程序退出
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
