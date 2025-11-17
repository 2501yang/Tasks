

# HTML学习笔记

###### VS code 编译器 字体放大按  Ctrl + = （等号键）

######                                     缩小：按  Ctrl + - （减号键）

### 基本框架

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>xxx</title>
</head>
<body>
    <h1>标题</h1>
    <p>段落</p>
</body>
</html>

- **<!DOCTYPE html>** 声明为 HTML5 文档
- **<html>** 元素是 HTML 页面的根元素
- **<head>** 元素包含了文档的元（meta）数据， **<meta charset="utf-8">** 定义网页编码格式为 **utf-8**，中文就不是乱码。
- **<title>** 元素描述了文档的标题
- **<body>** 元素包含了可见的页面内容
- **<h1>** 元素定义一个大标题
- **<p>** 元素定义一个段落



标题是通过 <h1> - <h6> 标签进行定义的。

<h1> 定义最大的标题 
<h6> 定义最小的标题

<hr> 创建水平线

可用于分隔内容

## 注释

<!-- xxx -->

**换行**

使用 **<br>**



## HTML 文本格式化标签

<b>粗体文本</b>

<em>着重文字</em>

<i>斜体字</i>

<small>小号字

<strong>加重语气</strong>

<sub>下标</sub>

<sup>上标</sup>

<ins>下划线</ins>

<del>删除线</del>



## HTML 超链接

<p>
<a href="/index.html">本文本</a> 是一个指向本网站中的一个页面的链接。</p>
<p><a href="https://www.xxx.com/">本文本</a> 是一个指向万维网上的页面的链接。</p>



### target：定义链接的打开方式。

- `_blank`: 在新窗口或新标签页中打开链接
- `_self`: 在当前窗口或标签页中打开链接（默认）
- `_parent`: 在父框架中打开链接
- `_top`: 在整个窗口中打开链接，取消任何框架



### rel：定义链接与目标页面的关系

nofollower: 表示搜索引擎不应跟踪该链接，常用于外部链接

noopener: 防止新的浏览上下文（页面）访问`window.opener`属性和`open`方法。

noreferrer: 不发送referer header（即不告诉目标网站你从哪里来的）

<a href="https://www.xxx.com" target="_blank" rel="noopener noreferrer">实例</a>



### id用于链接锚点，通常在同一页面中跳转到某个特定位置

<a id="tips">提示部分</a> 

<a href="#tips">跳到提示部分</a>


<a href="#section1">跳转到第1部分</a>

<div id="section1">这是第1部分</div>



## <meta> 标签- 使用实例

为搜索引擎定义关键词:

```html
<meta name="keywords" content="xxx">
```

为网页定义描述内容:

```html
<meta name="description" content="xxx">
```

定义网页作者:

```html
<meta name="author" content="xxx">
```

每30秒钟刷新当前页面:

```html
<meta http-equiv="refresh" content="30">
```



HTML <style> 元素
<style> 标签定义了HTML文档的样式文件引用地址.
在<style> 元素中你也可以直接添加样式来渲染 HTML 文档

<head>
<style type="text/css">
body {
    background-color:yellow;
}
p {
    color:blue
}
</style>
</head>

### title：定义链接的额外信息，当鼠标悬停在链接上时显示的工具提示

<a href="https://www.xxx.com" title="访问xxx 网站">访问 xxx</a>



## 使用CSS

背景色

<body style="background-color:white;">

<h2 style="background-color:red;">这是一个标题</h2>
<p style="background-color:yellow;">这是一个段落。</p>

</body>

font-family（字体），color（颜色），和font-size（字体大小）属性来定义字体

例：

<h1 style="font-family:verdana;">一个标题</h1>
<p style="font-family:arial;color:red;font-size:20px;">一个段落。</p>

####  text-align（文字对齐）属性指定文本的水平与垂直对齐方式

<h1 style="text-align:center;">居中对齐的标题</h1>

#### **图像**

<img src="url" alt="xxx">

URL 指存储图像的位置

alt 用来为图像定义一串预备的可替换的文本（在浏览器无法载入图像时，替换文本属性告诉读者她们失去的信息）

height（高度） 与 width（宽度）属性用于设置图像的高度与宽度   如 width="304" height="228"

矩形：(左上角顶点坐标为(x1,y1)，右下角顶点坐标为(x2,y2))

```
<area shape="rect" coords="x1,y1,x2,y2" href=url>
```

圆形：(圆心坐标为(X1,y1)，半径为r)

```html
<area shape="circle" coords="x1,y1,r" href=url>
```

### 表格

- **tr**：tr 是 table row 的缩写，表示表格的一行
- **td**：td 是 table data 的缩写，表示表格的数据单元格
- **th**：th 是 table header的缩写，表示表格的表头单元格（列）

<table>
  <thead>
    <tr>
      <th>列标题1</th>
      <th>列标题2</th>
      <th>列标题3</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>行1，列1</td>
      <td>行1，列2</td>
      <td>行1，列3</td>
    </tr>
    <tr>
      <td>行2，列1</td>
      <td>行2，列2</td>
      <td>行2，列3</td>
    </tr>
  </tbody>
</table>

##### 加边框

<table border="5">
    <tr>
        <td>Row 1, cell 1</td>
        <td>Row 1, cell 2</td>
    </tr>
</table>

#### 有序列表

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>

<h4>编号列表：</h4>
<ol>
 <li>Apples</li>
 <li>Bananas</li>
 <li>Lemons</li>
 <li>Oranges</li>
</ol>  

<h4>大写字母列表：</h4>
<ol type="A">
 <li>Apples</li>
 <li>Bananas</li>
 <li>Lemons</li>
 <li>Oranges</li>
</ol>  

<h4>小写字母列表：</h4>
<ol type="a">
 <li>Apples</li>
 <li>Bananas</li>
 <li>Lemons</li>
 <li>Oranges</li>
</ol>  

<h4>罗马数字列表：</h4>
<ol type="I">
 <li>Apples</li>
 <li>Bananas</li>
 <li>Lemons</li>
 <li>Oranges</li>
</ol>  

<h4>小写罗马数字列表：</h4>
<ol type="i">
 <li>Apples</li>
 <li>Bananas</li>
 <li>Lemons</li>
 <li>Oranges</li>
</ol>  
#### 嵌套列表

<h4>嵌套列表</h4>
<ul>
  <li>Coffee</li>
  <li>Tea
    <ul>
      <li>Black tea</li>
      <li>Green tea
        <ul>
          <li>China</li>
          <li>Africa</li>
        </ul>
      </li>
    </ul>
  </li>
  <li>Milk</li>
</ul>

#### 布局

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>
<table width="500" border="0">
<tr>
<td colspan="2" style="background-color:#FFA500;">
<h1>网页标题</h1>
</td>
</tr>
<tr>
<td style="background-color:#FFD700;width:10px;vertical-align:top;">
<b>菜单</b><br>
HTML<br>
CSS<br>
JavaScript
</td>
<td style="background-color:#eeeeee;height:200px;width:400px;vertical-align:top;">
内容在这里</td>
</tr>
<tr>
<td colspan="2" style="background-color:#FFA500;text-align:center;">
贾旭洋的制表</td>
</tr>
</table>




### 表单输入

- `<form>` 元素用于创建表单，`action` 属性定义了表单数据提交的目标 URL，`method` 属性定义了提交数据的 HTTP 方法（这里使用的是 "post"）。
- `<label>` 元素用于为表单元素添加标签，提高可访问性。
- `<input>` 元素是最常用的表单元素之一，它可以创建文本输入框、密码框、单选按钮、复选框等。`type` 属性定义了输入框的类型，`id` 属性用于关联 `<label>` 元素，`name` 属性用于标识表单字段。
- `<select>` 元素用于创建下拉列表，而 `<option>` 元素用于定义下拉列表中的选项。

<form action="/" method="post">
    <!-- 文本输入框 -->
    <label for="name">用户名:</label>
    <input type="text" id="name" name="name" required>
    <br>
    <!-- 密码输入框 -->
    <label for="password">密码:</label>
    <input type="password" id="password" name="password" required>
    <br>
    <!-- 单选按钮 -->
    <label>性别:</label>
    <input type="radio" id="male" name="gender" value="male" checked>
    <label for="male">男</label><input type="radio" id="female" name="gender" value="female"><label for="female">女</label>
    <br>
    <!-- 复选框 -->
    <input type="checkbox" id="subscribe" name="subscribe" checked><label for="subscribe">订阅推送信息</label>
    <br>
    <!-- 下拉列表 -->
    <label for="country">国家:</label>
    <select id="country" name="country">
        <option value="cn">CN</option>
        <option value="usa">USA</option>
        <option value="uk">UK</option>
    </select>
    <br>
    <!-- 提交按钮 -->
    <input type="submit" value="提交">

 文本域通过 **<input type="text">** 标签来设定

<form>
First name: <input type="text" name="firstname"><br>
Last name: <input type="text" name="lastname">
</form>

密码字段通过标签 **<input type="password">** 来定义

<form>
Password: <input type="password" name="pwd">
</form>

**<input type="radio">** 标签定义了表单的单选框选项

<form action="">
<input type="radio" name="sex" value="male">男<br>
<input type="radio" name="sex" value="female">女
</form>

**<input type="checkbox">** 定义了复选框

<form>
<input type="checkbox" name="vehicle[]" value="Bike">我喜欢x<br>
<input type="checkbox" name="vehicle[]" value="Car">我喜欢y
</form>

**<input type="submit">** 定义了提交按钮

<form name="input" action="html_form_action.php" method="get">
Username: <input type="text" name="user">
<input type="submit" value="Submit">
</form>

假如在上面的文本框内键入几个字母，然后点击确认按钮，那么输入数据会传送到 **html_form_action.php** 文件，该页面将显示出输入的结果

#### 下拉列表

<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
</head>
<body>
</head>
<body>
<form action="">
<select name="cars">
<option value="volvo">Volvo</option>
<option value="saab">Saab</option>
<option value="fiat">Fiat</option>
<option value="audi">Audi</option>
</select>
</form>





<!-- 以下表单使用 GET 请求发送数据到当前的 URL，method 默认为 GET。 -->

<form>
  <label>Name:
    <input name="submitted-name" autocomplete="name">
  </label>
  <button>Save</button>
</form>

<!-- 以下表单使用 POST 请求发送数据到当前的 URL。 -->
<form method="post">
  <label>Name:
    <input name="submitted-name" autocomplete="name">
  </label>
  <button>Save</button>
</form>

<!-- 表单使用 fieldset, legend, 和 label 标签 -->
<form method="post">
  <fieldset>
    <legend>Title</legend>
    <label><input type="radio" name="radio"> Select me</label>
  </fieldset>
</form>
### 颜色

| 黑   | #000000 | rgb(0,0,0)       |
| ---- | ------- | ---------------- |
| 红   | #FF0000 | rgb(255,0,0)     |
| 绿   | #00FF00 | rgb(0,255,0)     |
| 深蓝 | #0000FF | rgb(0,0,255)     |
| 黄   | #FFFF00 | rgb(255,255,0)   |
| 浅蓝 | #00FFFF | rgb(0,255,255)   |
| 梅红 | #FF00FF | rgb(255,0,255)   |
| 灰   | #C0C0C0 | rgb(192,192,192) |
| 白   | #FFFFFF | rgb(255,255,255) |

相对于使用 rgb(255,255,0)，使用 rgba(255,255,0,0.5) 可以实现设置颜色透明度的功能

<p style="background-color:rgb(255,255,0)">
通过 rbg 值设置背景颜色
</p>
<p style="background-color:rgba(255,255,0,0.25)">
通过 rbg 值设置背景颜色
</p>
<p style="background-color:rgba(255,255,0,0.5)">
通过 rbg 值设置背景颜色
</p>
<p style="background-color:rgba(255,255,0,0.75)">
通过 rbg 值设置背景颜色
</p>

这里的 0.5 表示透明度，范围 0~1，0 表示全透明



#### 脚本 

<script>
document.write("Hello World!")
</script>
<noscript>抱歉，你的浏览器不支持 JavaScript!</noscript>

只有在浏览器不支持脚本或者禁用脚本时，才会显示 <noscript> 元素中的内容



## HTML 实体

浏览器总是会截短 HTML 页面中的空格，如需在页面中增加空格的数量，需要使用 &nbsp



## URL

**scheme`://`host.domain`:`port`/`path`/`filename**

说明:

- scheme - 定义因特网服务的类型。最常见的类型是 http
- host - 定义域主机（http 的默认主机是 www）
- domain - 定义因特网域名，比如 runoob.com
- :port - 定义主机上的端口号（http 的默认端口号是 80）
- path - 定义服务器上的路径（如果省略，则文档必须位于网站的根目录中）。
- filename - 定义文档/资源的名称

## 
