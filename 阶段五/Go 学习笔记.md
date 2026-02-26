# Go 学习笔记

### Go 语言初识

##### 一、Go 前世今生

Go语言由来自Google公司的[Robert Griesemer](http://research.google.com/pubs/author96.html)，[Rob Pike](http://genius.cat-v.org/rob-pike/)和[Ken Thompson](http://genius.cat-v.org/ken-thompson/)三位大牛于2007年9月开始设计和实现，然后于2009年的11月对外正式发布。	

##### 二、Go 主要应用场景

Go语言中和并发编程相关的特性是全新的也是有效的，对数据抽象和面向对象编程的支持也很灵活。 同时还集成了自动垃圾收集技术用于更好地管理内存。Go语言尤其适合编写网络服务相关基础设施，同时也适合开发一些工具软件和系统软件。

##### 三、Go 的执行原理（Go命令相关内容）

Go语言的代码通过包（package）组织，包类似于其它语言里的库。每个源文件都以一条package声明语句开始，表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。缺少了必要的包或者导入了不需要的包，程序都无法编译通过。

Go语言提供的工具都通过一个单独的命令`go`调用，`go`命令有一系列子命令。最简单的一个子命令就是run。这个命令编译一个或多个**以.go结尾的源文件**，链接库文件，并运行最终生成的可执行文件。

##### 四、 Go 的基础语法

例：

```go
package main

import "fmt" //fmt 包实现了格式化输入/输出的函数

func main() {
    fmt.Println("Hello, 世界")//fmt.Println(...) 可以将字符串输出，并在最后自动增加换行字符 \n
}
```

注意 { 不能单独放在一行

读取用户输入

var name string
var age int
fmt.Scanln(&name, &age) // 读取两个值，分别赋给 name 和 age

##### 变量声明

1、指定变量类型，如果没有初始化，则变量默认为零值：           var 变量名字 类型 = 表达式

2、根据值自行判定变量类型：        “名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导

主要类型：int  float64  string bool

fmt.Printf是格式化输出，**%v**万能符号自适应输出

简短声明:=只能在函数内用，全局变量必须用var

字符串拼接用+号，比如"姓名：" + name

字符串可以用**反引号**整体输出

##### 基本数据类型转换 

T（v）:  将值v转换为类型T

var a int64=100

var b float64=float64（a）

#####  条件判断：if/else

if后的条件不用括号

 if里可以先声明变量（如num := score / 10），变量仅在if块内有效；

#####  循环：for（Go只有一种循环）

for替代while循环打印1-5
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}

#####  函数

函数定义：

func 函数名（形参列表）（返回值列表），参数格式是“参数名 类型”

支持多返回值

##### 数组

数组初始化： var numbers [5]int  初始值为 0

创建一个名为 numbers 的整数数组，并将其大小设置为 5，并初始化元素的值：

var numbers = [5]int{1, 2, 3, 4, 5}  或  numbers := [5]int{1, 2, 3, 4, 5}  也可以用[...]自动填充

若[ ] 中的数字不设置数组大小，会根据元素的个数来设置数组的大小

**注意：在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int 和 [10]int 是不同的类型。**

##### 切片（slice）

切片 是数组的引用 容量是可以自动增长的

初始化：例： var numbers []int= []int{1,2,3 }

内置的len和cap函数返回slice的长度和容量

例：

```Go
summer := months[6:9]  //引用month数组起始下标为6，只到但不包括9
```

make函数可以创建一个指定元素类型、长度和容量的slice。容量部分可以省略，此时容量等于长度。

```Go
make([]数据类型, len)
make([]数据类型, len, cap) 
```

for range遍历 例：

` s := []int{10, 20, 30}
  for i, v := range s { //不要索引可以 for _, v := range s {
  fmt.Printf("索引:%d 值:%d\n", i, v)
}`

**append函数**扩容

例：

```Go
package main
import "fmt"
func main() {
// 初始化一个切片
s := []int{1, 2, 3}
fmt.Println("原始切片:", s) // 输出: [1 2 3]
// 追加一个元素
s = append(s, 4)
fmt.Println("追加一个元素:", s) // 输出: [1 2 3 4]
// 追加多个元素
s = append(s, 5, 6, 7)
fmt.Println("追加多个元素:", s) // 输出: [1 2 3 4 5 6 7]
// 追加另一个切片（需要用 ... 展开）
anotherSlice := []int{8, 9, 10}
s = append(s, anotherSlice...)
fmt.Println("追加另一个切片:", s) // 输出: [1 2 3 4 5 6 7 8 9 10]
```
##### map

声明    var   map的变量名   map[类型]值类型

增删改查示例：

```Go
cities = make(map[string]string)// Go 没有内置的清空方法，可通过重新创建实现
// 创建一个空 map，键为 string 类型，值为 string 类型
cities := make(map[string]string)

// 创建并初始化 map
cities := map[string]string{
    "no1": "北京",
    "no2": "上海",
}

// 添加或更新键值对
cities["no3"] = "广州" // 新增
cities["no1"] = "深圳" // 修改

// 安全查找（推荐）
val, ok := cities["no1"]
if ok {
    fmt.Printf("找到 no1，值为 %v\n", val)
} else {
    fmt.Println("未找到 no1")
}

// 直接查找（无法区分键不存在和值为零值的情况）
val := cities["no1"]

// 遍历所有键值对
for key, val := range cities {
    fmt.Printf("key: %v, value: %v\n", key, val)
}
// 只遍历值
for _, val := range cities {
    fmt.Printf("value: %v\n", val)
}

```

**让无序的map按键的大小有序输出**

```go
package main
import (
	"fmt"   // 用于打印输出
	"sort"  // 用于排序
)
func main() {
	// 1. 定义一个map，键是int类型，值是string类型，模拟业务数据
	map1 := map[int]string{
		3: "张三",
		1: "李四",
		2: "王五",
		5: "赵六",
		4: "钱七",
	}
var keys []int // 步骤1：声明一个空的整型切片，用来存map的键
for k, _ := range map1 { // 步骤2：遍历map，把所有键捞出来
	keys = append(keys, k) // 把捞出来的键追加到切片里
}

sort.Ints(keys) // 步骤3：对存键的切片做升序排序
//fmt.Println(keys)  打印排序后的键，看看顺序对不对
// 步骤4：按排序后的键，去map里取对应的值并打印
for _, k := range keys {
	fmt.Printf("map1[%v]=%v \n", k, map1[k])
}
```
##### 结构体

定义

```Go
type 结构体名称 struct {
   属性         int
   属性       string
          
}

var  结构体变量名  结构体名称
```

名称大写表示可以被其他包引用，小写不行

##### 方法

在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个方法。

定义：

func (接收者变量 接收者类型) 方法名(参数列表) (返回值类型 为空可以省略) {
    方法体
}

常用于结构体

例：

func (s Student) ShowInfo() {
 	fmt.Printf("姓名：%s，年龄：%d\n", s.Name, s.Age)

}

##### 接口

定义接口：只写方法名，无方法体

type 接口名 interface {
    方法名1(参数列表) (返回值列表)
    方法名2(参数列表) (返回值列表)
    // ... 更多方法
}

实现接口

无需显式声明，只需让自定义类型实现接口的所有方法即可。

例：

```go
package main
import "fmt"
// 1. 定义核心接口：Usb
type Usb interface {
	Start() // 设备启动方法
	Stop()  // 设备停止方法
}
// 2. 定义Phone结构体
type Phone struct {
}
// 让Phone实现Usb接口的所有方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}
// 3. 定义Camera结构体
type Camera struct {
}
// 让Camera实现Usb接口的所有方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}
// 4. 定义Computer结构体，用于统一调用Usb设备
type Computer struct {
}
// 编写一个方法，接收任意实现了Usb接口的设备
func (com Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func main() {
	// 实例化设备
	phone := Phone{}
	camera := Camera{}
	computer := Computer{}
// 统一调用，体现多态
computer.Working(phone)
computer.Working(camera)
```
##### 错误处理

error：Go 内置的错误类型，专门存错误

err：通用习惯名，存错误的变量

fmt.Errorf()：用来创建一条自定义错误

用法：

// 1. 定义错误变量
var err error

// 2. 出错时创建错误
err = fmt.Errorf("写错误原因")

// 3. 判断是否出错：err != nil 表示有错
if err != nil {
    fmt.Println("出错了：", err)
    return
}

法2： 使用defer + recover来捕获和处理异常
defer func() {
   err := recover() // recover()是内置函数，用于捕获当前goroutine的panic
   if err != nil { // 如果err不为nil，说明捕获到了异常
        fmt.Println("err=", err)
        // 这里可以添加更多错误处理逻辑，比如日志记录、资源清理等
         }
     }()