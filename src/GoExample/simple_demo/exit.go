package main

import "fmt"
import "os"

func main() {

	// 当使用 `os.Exit` 时 `defer` 将_不会_ 执行，所以这里的 `fmt.Println`
	// 将永远不会被调用。
	defer fmt.Println("!")

	// 退出并且退出状态为 3。
	os.Exit(3)
}
