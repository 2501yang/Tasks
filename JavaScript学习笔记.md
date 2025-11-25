# JavaScript学习笔记

*JavaScript 与 Java 是两种完全不同的语言



### JavaScript语法

##### JavaScript 语句是用分号分隔

双斜杠 **//** 注释

#### 字面量

固定值称为字面量，可以是**数字**，**字符串**（使用单引号或双引号），**表达式**（如5 * 10）

**数组**（[40, 100, 1, 5, 25, 10]，**函数**（function myFunction(a, b) { return a * b;}）

**对象**（{firstName:"John", lastName:"Doe", age:50, eyeColor:"blue"}）

对象由花括号分隔。在括号内部，对象的属性以名称和值对的形式 (name : value) 来定义。属性由逗号分隔：

**特殊字符**：字符串写在单引号或双引号中可能无法解析使用反斜杠 (\) 来转义

例如  `  "We are the so-called \"Vikings\" from the north."`

### 变量

常见的是驼峰法的命名规则，如 lastName (而不是lastname)

变量可以重复声明（覆盖原变量）

##### 使用关键字 **var** 来定义变量， 使用等号来为变量赋值

例如

var x, y;

x = 5;
y = 6;

#### 数组输出例 


<script>
var i;
var cars = new Array();
cars[0] = "Saab";
cars[1] = "Volvo";
cars[2] = "BMW";
for (i=0;i<cars.length;i++)
{
document.write(cars[i] + "<br>");
}
</script>

##### JavaScript 会忽略多余的空格

var person="A";
var person = "A";

等效



HTML 中的 Javascript 脚本代码必须位于 **<script>** 与 **</script>** 标签之间

Javascript 脚本代码可被放置在 HTML 页面的 **<body>** 和 **<head>** 部分中

<head>中

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>JavaScript 函数</title> 
<script>
function myFunction(){
	document.getElementById("demo").innerHTML="JavaScript 函数";
}
</script>
</head>
<body>	
<h5>我的 Web 页面</h5>
<p id="demo">段落</p>
<button type="button" onclick="myFunction()">点击这里</button>
</body>
</html>

<body>中

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>JavaScript 函数</title> 
</head>
<body>
<h5>我的Web页面</h5>
<p id="demo">段落</p>
<button type="button" onclick="myFunction()">点击这里</button>
<script>
function myFunction(){
	document.getElementById("demo").innerHTML="JavaScript 函数";
}
</script>
</body>
</html>

###### 外部的 JavaScript

把脚本保存到外部文件中，外部文件通常包含被多个网页使用的代码。

外部 JavaScript 文件的文件扩展名是 .js

如需使用外部文件，请在 <script> 标签的 "src" 属性中设置该 .js 文件：

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>外部javascript</title> 
</head>
<body>
<h5>我的 Web 页面</h5>
<p id="demo">段落</p>
<button type="button" onclick="myFunction()">点击这里</button>
<p><b>注释：</b>myFunction 保存在名为 "myScript.js" 的外部文件中</p>
<script src="myScript.js"></script>
</body>
</html>

myScript.js 文件代码为

function myFunction()
{
    document.getElementById("demo").innerHTML="JavaScript 函数";
}





#### 创建JS文件

打开VS code在搜索框输入 js，创建 **test.js** 文件：

![img](https://www.runoob.com/wp-content/uploads/2024/04/c54f2a8f6a9b9eaf0f706c477d078204.png)

保存 test.js 文件代码，右击文件名，在集成终端执行命令

node test.js





#### 从 JavaScript 访问某个 HTML 元素

使用 document.getElementById(*id*) 

 "id" 属性来标识 HTML 元素， innerHTML 来获取或插入元素内容

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>
<h5>我的 Web 页面</h5>
<p id="demo">段落</p>
<script>
document.getElementById("demo").innerHTML="段落";
</script>
</body>
</html>

innerHTML="段落"  是用于修改元素的 HTML 





#### 将JavaScript直接写在HTML 文档中：

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title></title> 
</head>
<body>
<h4> Web 页面</h4>
<p>段落</p>
<button onclick="myFunction()">点击授时</button>
<script>
function myFunction() 
{
    document.write(Date());
}
</script>
</body>
</html>

<script>
document.write(xxx);
</script>

使用 document.write() 可以向文档写入内容

如果在文档已完成加载后执行 document.write，整个 HTML 页面将被覆盖



#### 调用带参数的函数

<!DOCTYPE html>
<html>	
<head> 
<meta charset="utf-8"> 
<title>调用带参数的函数</title> 
</head>
<body>
<p>点击这个按钮，来调用带参数的函数</p>
<button onclick="myFunction('Harry Potter','Wizard')">点击这里</button>
<script>
function myFunction(name,job){
	alert("Welcome " + name + ", the " + job);
}
</script>
</body>
</html>



#### 条件语句

`<script type="text/javascript">
var d = new Date();
var time = d.getHours();
if (time<10)
{
	document.write("<b>早上好</b>");
}
else if (time>=10 && time<20)
{
	document.write("<b>今天好</b>");
}
else
{
	document.write("<b>晚上好!</b>");
}
</script>`



#### [更多javascript实例学习](https://www.runoob.com/js/js-examples.html)