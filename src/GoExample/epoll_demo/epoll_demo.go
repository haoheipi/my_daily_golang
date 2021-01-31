package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("服务端进程id:", os.Getpid())
	listen, err := net.Listen("tcp", "0.0.0.0:9009")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}

	for {
		conn, err := listen.Accept() //等待建立连接
		if err != nil {
			fmt.Println("建立连接失败：", err)
			continue
		}
		//开启协程处理
		go func() {
			defer conn.Close()
			for {
				buf := make([]byte, 128)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("读出错：", err)
					return
				}
				fmt.Println("读取到的数据：", string(buf[:n]))
			}
		}()
	}

}
