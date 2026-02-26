package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 请求/响应结构
type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 只处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST 方法", http.StatusMethodNotAllowed)
		return
	}
	
	// 解析 JSON 请求
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "无效的 JSON 格式", http.StatusBadRequest)
		return
	}
	
	fmt.Printf("收到消息: %s (来自 %s)\n", req.Message, r.RemoteAddr)
	
	// 构造响应
	resp := Response{
		Message:   "服务器收到: " + req.Message,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	
	// 返回 JSON 响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// 注册路由
	http.HandleFunc("/api/message", helloHandler)
	
	// 启动服务器
	fmt.Println("HTTP 服务器已启动，监听端口 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}