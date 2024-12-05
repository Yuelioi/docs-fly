---
title: Source-Text
order: 1
---

# Source-Text

## getStyleAt()

用法：getStyleAt(charIndex, t = time)

说明：获取某秒的某个字符的样式

参数：charIndex 为字符的索引，从 0 开始排序。t 为时间，因为文字也是可以根据时间变化内容的，比如第 1 秒是 1，第 20 秒是 20。所以参数 t 也是有点用的

示例：

源文字

使用表达式后

![](https://mir.yuelili.com/user/AE/expression/a-z/source-text.png)

![](https://mir.yuelili.com/user/AE/expression/a-z/text-sample1.png)

```javascript
getStyleAt(1, (t = time)); //获取第1个字符的样式
```

## createStyle()

用法：createStyle() 或 createStyle().你的文字样式

说明：可以使用它创建一个空的样式，或者直接创建好所有样式

类型：新样式

示例 1：先创建空样式，后面再赋值

```javascript
styleA = createStyle(); //先创建一个空样式，赋值给styleA（样式A）
styleB = styleA.setFontSize(100); //给样式A增加一个样式内容，字体大小为100号，然后赋值给样式B
styleB; //调用样式B，此时文字变成100字号了
```

示例 2：直接创建样式（非空样式）

```javascript
createStyle().setFontSize(100); //字体变为100字号了
```

其他说明：有点像新建空变量，然后赋值，以及直接新建变量的同时进行赋值

参考如下：如果你学了表达式，那么以下可以加深理解。如果没有，可以跳过

```javascript
var a;
a = 10;
以及;
var a = 10;
```
