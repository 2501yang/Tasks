# Gorm 学习笔记

### ORM 与 Gorm 核心概念

#### 一、ORM

ORM（Object-Relational Mapping，对象关系映射）是一种编程技术，核心作用是把编程语言中的「对象」和数据库中的「表」建立映射关系。
通俗理解：你不用再写复杂的 SQL 语句，而是通过操作 Go 语言的结构体（对象）来间接操作数据库表，ORM 框架会自动帮你把对象操作转换成对应的 SQL 语句。
核心作用：
1、简化数据库操作，降低学习成本（不用死记复杂 SQL）；
2、减少重复的 SQL 代码，提高开发效率；
3、屏蔽不同数据库的 SQL 差异（比如 MySQL 和 PostgreSQL 的语法差异）。

#### 二、Gorm 

Gorm 是 Go 语言生态中最流行的 ORM 框架，全称是 Go ORM。
核心优势：
1、语法简洁，符合 Go 语言的使用习惯；
2、功能强大，支持完整的 CRUD、关联查询、事务、迁移等；
3、兼容性好，支持 MySQL、PostgreSQL、SQLite、SQL Server 等主流数据库；
4、文档完善，中文文档齐全；
5、自动迁移（AutoMigrate）：可以根据结构体自动创建 / 更新数据库表，不用手动写建表 SQL。

### Gorm 基础操作

#### 模型定义

 Go 的结构体，对应数据库中的一张表，Gorm 通过结构体的字段、标签来映射表的列和属性

##### 用户表模型

// User 对应数据库中的 users 表（Gorm 会自动将结构体名转小写并加 s）
type User struct {
  // gorm.Model 是 Gorm 内置的基础模型，包含 ID、CreatedAt、UpdatedAt、DeletedAt 四个字段
  gorm.Model
  // 用户名，非空，唯一
  Username string `gorm:"type:varchar(50);not null;unique;comment:用户名"`
  // 年龄，默认值 0
  Age      int    `gorm:"type:int;default:0;comment:年龄"`
  // 邮箱，允许为空
  Email    string `gorm:"type:varchar(100);comment:邮箱"`
  // 忽略该字段（不映射到数据库表）
  Remark   string `gorm:"-"`
}

##### gorm.Model：内置结构体

type Model struct {
  ID                     uint           `gorm:"primarykey"` // 主键
  CreatedAt       time.Time      // 创建时间（自动填充）
  UpdatedAt      time.Time      // 更新时间（自动填充）
  DeletedAt        gorm.DeletedAt `gorm:"index"`      // 软删除标记（逻辑删除）
}

##### 结构体标签

type:varchar(50)：指定列的数据类型；
not null：列非空；
unique：列值唯一；
default:0：列默认值；
comment：列注释；
-：忽略该字段，不生成列

##### 自定义表名

// 方法1：在结构体中定义 TableName 方法
type User struct {
	ID   uint
	Name string
}

func (User) TableName() string {
	return "my_users"  // 表名变成 my_users
}

// 方法2：使用 gorm 标签
type Product struct {
	ID   uint
	Name string
} `gorm:"table:products"`

### Gorm 基础 CRUD 操作

1. ##### 新增（Create）

// 1. 新增单条数据
user := User{Username: "zhangsan", Age: 20, Email: "zhangsan@test.com"}
result := db.Create(&user) // 传入指针，新增后 user.ID 会被自动赋值
if result.Error != nil {
  fmt.Printf("新增失败：%v\n", result.Error)
} else {
  fmt.Printf("新增成功，ID：%d，影响行数：%d\n", user.ID, result.RowsAffected)
}

// 2. 批量新增
users := []User{
  {Username: "lisi", Age: 21},
  {Username: "wangwu", Age: 22},
}
db.Create(&users) // 批量新增

2. ##### 查询（Read）

// 1. 根据主键查询单条数据
var user User
// 查询 ID=1 的用户，First 会返回第一条匹配数据（主键升序）
db.Find(&user, 1) 
fmt.Printf("ID=1 的用户：%+v\n", user)

// 2. 查询所有数据
var users []User
db.Find(&users)
fmt.Printf("所有用户：%+v\n", users)

**条件查询**

var user User
var users []User

// 等于条件
db.Where("age = ?", 20).Find(&users) // 年龄=20 的所有用户

// 不等于条件
db.Where("age != ?", 20).Find(&users) // 年龄≠20 的所有用户

// 模糊查询（like）
db.Where("username LIKE ?", "%zhang%").Find(&users) // 用户名包含 zhang 的用户

// 范围查询
db.Where("age BETWEEN ? AND ?", 18, 25).Find(&users) // 年龄在 18-25 之间的用户

// 多条件（AND）
db.Where("age > ? AND email IS NOT NULL", 20).Find(&users) // 年龄>20 且邮箱非空的用户

// 多条件（OR）
db.Where("age = ?", 20).Or("username = ?", "lisi").Find(&users) // 年龄=20 或用户名=lisi 的用户

##### 分页

分页是查询大量数据时的必备操作，核心是 Offset（偏移量）和 Limit（每页条数）：

page := 1     // 当前页
pageSize := 2 // 每页条数
offset := (page - 1) * pageSize // 偏移量（跳过前 N 条）

// 分页查询：跳过 offset 条，取 pageSize 条
db.Offset(offset).Limit(pageSize).Find(&users)
fmt.Printf("第 %d 页数据：%+v\n", page, users)

// 统计总条数
var total int64
db.Model(&User{}).Count(&total) // 统计用户总数
fmt.Printf("用户总数：%d\n", total)

##### **排序**

// 按年龄降序排序（desc），升序用 asc（默认）
db.Order("age desc").Find(&users)
// 多字段排序：先按年龄降序，再按创建时间升序
db.Order("age desc, created_at asc").Find(&users)

##### 更新（Update）

// 1. 更新单个字段
db.Model(&User{ID: 1}).Update("age", 25) // 将 ID=1 的用户年龄改为 25

// 2. 更新多个字段
db.Model(&User{ID: 1}).Updates(User{Age: 26, Email: "new_zhangsan@test.com"})

// 3. 条件更新
// 将年龄>20 的用户邮箱改为 empty@test.com
db.Model(&User{}).Where("age > ?", 20).Update("email", "empty@test.com")

##### 删除（Delete）

Gorm 默认是软删除（逻辑删除），即不会真正删除数据，只是将 deleted_at 字段设置为当前时间

// 1. 软删除（默认）
db.Delete(&User{}, 1) // 删除 ID=1 的用户（仅标记 deleted_at）

// 2. 条件软删除
db.Where("age < ?", 18).Delete(&User{}) // 删除年龄<18 的用户

// 3. 硬删除（物理删除）
db.Unscoped().Delete(&User{}, 1) // 真正删除 ID=1 的用户

**恢复**

使用 Update 将 deleted_at 字段设置为 nil
result := db.Unscoped().Where("id = ?", 1).Update("deleted_at", nil)

// 检查是否恢复成功
if result.Error != nil {
    fmt.Println("恢复失败:", result.Error)
} else if result.RowsAffected == 0 {
    fmt.Println("未找到该数据")
} else {
    fmt.Println("恢复成功！")
}
