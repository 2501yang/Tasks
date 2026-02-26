package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 监听 TCP 端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("启动服务器失败:", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("TCP 服务器已启动，监听端口 8080...")
	
	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err)
			continue
		}
		
		// 启动 goroutine 处理客户端
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("客户端连接:", clientAddr)
	
	// 发送欢迎消息
	conn.Write([]byte("欢迎连接到 TCP 服务器!\n"))
	
	// 读取客户端消息
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("客户端断开:", clientAddr)
			break
		}
		
		message = strings.TrimSpace(message)
		fmt.Printf("收到 [%s]: %s\n", clientAddr, message)
		
		// 回显消息
		response := fmt.Sprintf("服务器收到: %s\n", message)
		conn.Write([]byte(response))
		
		// 检查退出命令
		if message == "quit" {
			conn.Write([]byte("再见!\n"))
			break
		}
	}
}
