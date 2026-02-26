package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run http_client.go <消息内容>")
		return
	}
	
	message := strings.Join(os.Args[1:], " ")
	
	// 构造请求
	reqBody := Request{Message: message}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("JSON 编码失败:", err)
		return
	}
	
	// 发送 HTTP 请求
	resp, err := http.Post("http://127.0.0.1:8080/api/message", 
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()
	
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	
	// 解析响应
	var respData Response
	err = json.Unmarshal(body, &respData)
	if err != nil {
		fmt.Println("JSON 解码失败:", err)
		return
	}
	
	fmt.Printf("服务器响应: %s\n", respData.Message)
	fmt.Printf("时间戳: %s\n", respData.Timestamp)
}