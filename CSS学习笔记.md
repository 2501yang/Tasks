# `CSS`学习笔记



### 语法

`CSS `规则由两个主要的部分构成：选择器，以及一条或多条声明:

![img](https://www.runoob.com/wp-content/uploads/2013/07/632877C9-2462-41D6-BD0E-F7317E4C42AC.jpg)

选择器是需要改变样式的 HTML 元素。

每条声明由一个属性和一个值组成。

总是以分号 **;** 结束，声明总以大括号 **{}** 括起来

注释以 **/\*** 开始, 以 ***/** 结束

### id 和 class

HTML元素以id属性来设置id选择器,CSS 中 id 选择器以 "#" 来定义 

#para1
{
    text-align:center;
    color:red;
}

class可以在多个元素中使用

class 选择器在 HTML 中以 class 属性表示, 在 CSS 中，类选择器以一个点 **.** 号显示

p.center
{
	text-align:center;
}

<p class="center">段落居中对齐</p>

<p id="para1">Hello World!</p>

## 外部样式

每个页面使用 <link> 标签链接到样式表。 <link> 标签在（文档的）头部

<head>
<link rel="stylesheet" type="text/css" href="mystyle.css">
</head>

浏览器会从文件 mystyle.css 中读到样式声明，并根据它来格式文档

编辑样式表应该以 .css 扩展名进行保存

hr {color:sienna;}
p {margin-left:20px;}
body {background-image:url("/images/back40.gif");}

## 内部样式

使用 <style> 标签在文档头部定义

<head>
<style>
hr {color:sienna;}
p {margin-left:20px;}
body {background-image:url("images/back40.gif");}
</style>
</head>



### CSS 背景

CSS 属性定义背景效果:

- background-color 背景颜色

<style>
h1
{
	background-color:#6495ed;
}
p
{
	background-color:#e0ffff;
}
div
{
	background-color:#b0c4de;
}
</style>

- background-image 背景图像**平铺**
- background-repeat**不平铺**

body
{
background-image:url('gradient2.png');
background-repeat:repeat-x;
}

background-position 属性改变图像在背景中的位置

body
{
background-image:url('img_tree.png');
background-repeat:no-repeat;
background-position:right top;
}

## 文本

#### 对齐方式

文本可居中或对齐到左或右,当text-align设置为"justify"，每一行被展开为宽度相等，左，右外边距是对齐

<style>
h1 {text-align:center;}
p.date {text-align:right;}
p.main {text-align:justify;}
</style>

##### 文本缩进属性是用来指定文本的第一行的缩进

 <style>
p {text-indent:50px;}
</style> 

##### 文本阴影


<style>
h1 {text-shadow:2px 2px #FF0000;}
</style>

#####  字体

p{font-family:"Times New Roman", Times, serif;}

- 正常 - 正常显示文本
- 斜体 - 以斜体字显示的文字
- 倾斜的文字 - 文字向一边倾斜

<style>
p.normal {font-style:normal;}
p.italic {font-style:italic;}
p.oblique {font-style:oblique;}
</style>
对应


<p class="normal">这是一个段落,正常。</p>
<p class="italic">这是一个段落,斜体。</p>
<p class="oblique">这是一个段落,斜体。</p>

#### 倾角

<div style="font-style: oblique 5deg;">This is a sentence.</div>
<div style="font-style: oblique 10deg;">This is a sentence.</div>
<div style="font-style: oblique 20deg;">This is a sentence.</div>
<div style="font-style: oblique 30deg;">This is a sentence.</div>
<div style="font-style: oblique 40deg;">This is a sentence.</div>

##### font-size 属性设置文本的大小

h1 {font-size:40px;}
h2 {font-size:30px;}
p {font-size:14px;}

**不同的字体粗细**

p.normal {font-weight:normal;}
p.thick {font-weight:bold;}
p.thicker {font-weight:900;}

### 链接

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<style>
a:link,a:visited
{
	display:block;
	font-weight:bold;
	color:#FFFFFF;
	background-color:#98bf21;
	width:120px;
	text-align:center;
	padding:4px;
	text-decoration:none;
}
a:hover,a:active
{
	background-color:#7A991A;
}
</style>
</head>
<body>
<a href="/css/" target="_blank">这是一个链接</a>
</body>
</html>    

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<style>
a.one:link {color:#ff0000;}
a.one:visited {color:#0000ff;}
a.one:hover {color:#ffcc00;}
a.two:link {color:#ff0000;}
a.two:visited {color:#0000ff;}
a.two:hover {font-size:150%;}
a.three:link {color:#ff0000;}
a.three:visited {color:#0000ff;}
a.three:hover {background:#66ff66;}
a.four:link {color:#ff0000;}
a.four:visited {color:#0000ff;}
a.four:hover {font-family:Georgia, serif;}
a.five:link {color:#ff0000;text-decoration:none;}
a.five:visited {color:#0000ff;text-decoration:none;}
a.five:hover {text-decoration:underline;}
</style>
</head>
<body>
<p>将鼠标移至链接上改变样式.</p>
<p><b><a class="one" href="/css/" target="_blank">这个链接改变颜色</a></b></p>
<p><b><a class="two" href="/css/" target="_blank">这个链接改变字体大小</a></b></p>
<p><b><a class="three" href="/css/" target="_blank">这个链接改变背景颜色</a></b></p>
<p><b><a class="four" href="/css/" target="_blank">这个链接改变字体类型</a></b></p>
<p><b><a class="five" href="/css/" target="_blank">这个链接改变文字修饰</a></b></p>
</body>
</html>

#### 边框

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<style>
p.none {border-style:none;}
p.dotted {border-style:dotted;}
p.dashed {border-style:dashed;}
p.solid {border-style:solid;}
p.double {border-style:double;}
p.groove {border-style:groove;}
p.ridge {border-style:ridge;}
p.inset {border-style:inset;}
p.outset {border-style:outset;}
p.hidden {border-style:hidden;}
p.mix {border-style: dotted dashed solid double;}
</style>
</head>
<body>
<p class="none">无边框。</p>
<p class="dotted">虚线边框。</p>
<p class="dashed">虚线边框。</p>
<p class="solid">实线边框。</p>
<p class="double">双边框。</p>
<p class="groove"> 凹槽边框。</p>
<p class="ridge">垄状边框。</p>
<p class="inset">嵌入边框。</p>
<p class="outset">外凸边框。</p>
<p class="hidden">隐藏边框。</p>
<p class="mix">混合边框</p>
</body>
</html>

#### 单独设置各边

**dotted solid double dashed**

p
{
	border-top-style:dotted;
	border-right-style:solid;
	border-bottom-style:double;
	border-left-style:dashed;
}

#### 轮廓颜色

p 
{
	border:1px solid red;
	outline-style:dotted;
	outline-color:#00ff00;
}



#### 对于相同样式的元素

h1 {
    color:green;
}
h2 {
    color:green;
}
p {
    color:green;
}

可以使用分组选择器

h1,h2,p
{
    color:green;
}

#### 重叠的元素

z-index属性指定了一个元素的堆叠顺序（哪个元素应该放在前面，或后面）

img
{
    position:absolute;
    left:0px;
    top:0px;
    z-index:-1;
}

#####  overflow 可以控制内容溢出元素框时在对应的元素区间内添加滚动条

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>菜鸟教程(runoob.com)</title>
<style>
#overflowTest {
    background: #4CAF50;
    color: white;
    padding: 15px;
    width: 80%;
    height: 100px;
    overflow: scroll;
    border: 1px solid #ccc;
}
</style>
</head>
<body>
<div id="overflowTest">
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
<p>这里的文本内容是可以滚动的，滚动条方向是垂直方向。</p>
</div>
</body>
</html>



#####  Float，会使元素向左或向右移动，其周围的元素也会重新排列

如图像是右浮动，下面的文本流将环绕在它左边

img
{
    float:right;
}

### 列表导航栏

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<style>
ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    width: 200px;
    background-color: #f1f1f1;
}
li a {
    display: block;
    color: #000;
    padding: 8px 16px;
    text-decoration: none;
}
/* 鼠标移动到选项上修改背景颜色 */
li a:hover {
    background-color: #555;
    color: white;
}
</style>
</head>
<body>
<ul>
<li><a href="#home">主页</a></li>
<li><a href="#news">新闻</a></li>
<li><a href="#contact">联系</a></li>
<li><a href="#about">关于</a></li>
</ul>
<p>背景颜色添加到链接中显示链接的区域</p>
<p>注意,整个区域是可点击的链接,而不仅仅是文本。</p>
</body>
</html>

#### 下拉菜单

<style>
.dropdown {
  position: relative;
  display: inline-block;
}
.dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
  padding: 12px 16px;
}
.dropdown:hover .dropdown-content {
  display: block;
}
</style>
<div class="dropdown">
  <span>鼠标移动这</span>
  <div class="dropdown-content">
    <p>xxx</p>
    <p>xxx</p>
  </div>
</div>

#### 添加底部箭头

.tooltip .tooltiptext::after {
    content: " ";
    position: absolute;
    top: 100%; /* 提示工具底部 */
    left: 50%;
    margin-left: -5px;
    border-width: 5px;
    border-style: solid;
    border-color: black transparent transparent transparent;
}



### 输入框 样式

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"> 
<title>菜鸟教程(runoob.com)</title> 
<style> 
input {
  width: 100%;
}
</style>
</head>
<body>
<p>全宽输入框:</p>
<form>
  <label for="fname">First Name</label>
  <input type="text" id="fname" name="fname">
</form>
</body>
</html>



##### 颜色

input[type=text] {
  background-color: #3CBC8D;
  color: white;
}

<!DOCTYPE html>
<html>
<head>
<style>
* {
  box-sizing: border-box;
}
input[type=text], select, textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 4px;
  resize: vertical;
}
label {
  padding: 12px 12px 12px 0;
  display: inline-block;
}
input[type=submit] {
  background-color: #4CAF50;
  color: white;
  padding: 12px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  float: right;
}
input[type=submit]:hover {
  background-color: #45a049;
}
.container {
  border-radius: 5px;
  background-color: #f2f2f2;
  padding: 20px;
}
.col-25 {
  float: left;
  width: 25%;
  margin-top: 6px;
}
.col-75 {
  float: left;
  width: 75%;
  margin-top: 6px;
}
/* 清除浮动 */
.row:after {
  content: "";
  display: table;
  clear: both;
}
/* 响应式布局 layout - 在屏幕宽度小于 600px 时， 设置为上下堆叠元素 */
@media screen and (max-width: 600px) {
  .col-25, .col-75, input[type=submit] {
    width: 100%;
    margin-top: 0;
  }
}
</style>
</head>
<body>
<h3>响应式表单</h3>
<p>响应式表带可以根据浏览器窗口的大小重新布局各个元素，我们可以通过重置浏览器窗口大小来查看效果：</p>


<div class="container">
  <form action="/action_page.php">
  <div class="row">
    <div class="col-25">
      <label for="fname">First Name</label>
    </div>
    <div class="col-75">
      <input type="text" id="fname" name="firstname" placeholder="Your name..">
    </div>
  </div>
  <div class="row">
    <div class="col-25">
      <label for="lname">Last Name</label>
    </div>
    <div class="col-75">
      <input type="text" id="lname" name="lastname" placeholder="Your last name..">
    </div>
  </div>
  <div class="row">
    <div class="col-25">
      <label for="country">Country</label>
    </div>
    <div class="col-75">
      <select id="country" name="country">
        <option value="australia">Australia</option>
        <option value="canada">Canada</option>
        <option value="usa">USA</option>
      </select>
    </div>
  </div>
  <div class="row">
    <div class="col-25">
      <label for="subject">Subject</label>
    </div>
    <div class="col-75">
      <textarea id="subject" name="subject" placeholder="Write something.." style="height:200px"></textarea>
    </div>
  </div>
  <div class="row">
    <input type="submit" value="Submit">
  </div>
  </form>
</div>


## 网页布局

##### 头部区域

.header {
  background-color: #F1F1F1;
  text-align: center;
  padding: 20px;
}

菜单导航

/* 导航条 */
.topnav {
  overflow: hidden;
  background-color: #333;
}

/* 导航链接 */
.topnav a {
  float: left;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

/* 链接 - 修改颜色 */
.topnav a:hover {
  background-color: #ddd;
  color: black;
}