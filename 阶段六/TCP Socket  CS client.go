package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("已连接到服务器")
	
	// 读取服务器欢迎消息
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Print(string(buffer[:n]))
	
	// 读取用户输入并发送
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入消息 (输入 quit 退出): ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		
		// 发送消息
		conn.Write([]byte(message + "\n"))
		
		// 读取服务器响应
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("连接断开")
			break
		}
		fmt.Print(string(buffer[:n]))
		
		if message == "quit" {
			break
		}
	}
}