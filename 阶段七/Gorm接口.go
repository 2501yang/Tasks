package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 定义 GORM 模型（对应 MySQL 的 users 表）
type User struct {
	gorm.Model        // 内置字段：ID、CreatedAt、UpdatedAt、DeletedAt
	Username string `json:"username" gorm:"type:varchar(50);not null;unique"` // 用户名
	Age      int    `json:"age" gorm:"type:int;default:0"`                   // 年龄
	Email    string `json:"email" gorm:"type:varchar(100)"`                  // 邮箱
}

// 全局 DB 实例（新手简化版，实际项目可优化）
var db *gorm.DB

// 2. 初始化数据库连接
func initDB() error {
	// 替换为你的 MySQL 配置！！！
	dsn := "yang:123456@tcp(192.168.206.134:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect db failed: %v", err)
	}

	// 自动迁移：根据模型创建/更新表结构
	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("auto migrate failed: %v", err)
	}
	return nil
}

// 3. 接口处理函数 - 新增用户
func createUser(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 JSON
	w.Header().Set("Content-Type", "application/json")

	// 解析请求体中的 JSON 数据到 User 结构体
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body: " + err.Error()})
		return
	}

	// GORM 新增数据
	result := db.Create(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "create user failed: " + result.Error.Error()})
		return
	}

	// 返回成功结果
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// 4. 接口处理函数 - 查询单个用户（按 ID）
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 从 URL 获取 ID 参数（比如 /user/1）
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid user id"})
		return
	}

	// GORM 按主键查询
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found: " + result.Error.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// 5. 接口处理函数 - 查询用户列表
func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GORM 查询所有用户（排除软删除）
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "list users failed: " + result.Error.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"total":   len(users),
		"data":    users,
	})
}

// 6. 接口处理函数 - 更新用户
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 获取 ID 参数
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid user id"})
		return
	}

	// 解析更新数据
	var updateData User
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body: " + err.Error()})
		return
	}

	// GORM 更新（只更新传入的非零值字段）
	var user User
	result := db.Model(&user).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "update user failed: " + result.Error.Error()})
		return
	}

	// 查询更新后的用户
	db.First(&user, id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// 7. 接口处理函数 - 删除用户（软删除）
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 获取 ID 参数
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid user id"})
		return
	}

	// GORM 软删除
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "delete user failed: " + result.Error.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// 8. 主函数：初始化 + 注册路由 + 启动服务
func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		panic("init db failed: " + err.Error())
	}
	fmt.Println("db init success")

	// 注册路由
	mux := http.NewServeMux()
	// 新增用户：POST /users
	mux.HandleFunc("POST /users", createUser)
	// 查询单个用户：GET /users/{id}
	mux.HandleFunc("GET /users/{id}", getUser)
	// 查询用户列表：GET /users
	mux.HandleFunc("GET /users", listUsers)
	// 更新用户：PUT /users/{id}
	mux.HandleFunc("PUT /users/{id}", updateUser)
	// 删除用户：DELETE /users/{id}
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	// 启动 HTTP 服务
	fmt.Println("server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic("start server failed: " + err.Error())
	}
}