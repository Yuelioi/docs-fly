---
title: Global
order: 2
---

# Global 全局函数

## 全局函数

这些特定于 AfterEffects 的全局可用函数。任何 JavaScript 对象或函数都可以调用这些函数，这些函数使您可以在“信息”面板的小区域(只有三行)显示文本，可以将数字时间值与字符串值进行相互转换，也可以生成随机数。  
ExtendScript 定义了标准用户 I / O 的其他全局函数(alert，confirm 和 prompt)以及文件 I /
O 的静态函数；有关详细的参考信息，请参见《[JavaScript 工具指南](https://extendscript.docsforadobe.dev/)》。

### clearOutput() 信息面板清空

全名：clearOutput()

描述：清除“信息”面板中的文本。下面的 write()则是写入文本

![](https://cdn.yuelili.com/20210824212139.png)

参数：无。

返回：无。

### currentFormatToTime() 帧转秒

全名：currentFormatToTime(formattedTime, fps[, isDuration])

描述：给定帧速率，将帧转为秒数。

示例：时间值为 0:00:12，并且帧速率为 24 fps，则返回 0.5 秒(12/24)。帧速率为 30 fps，则返回 0.4 秒(12/30)。

参数：

- formattedTime：帧数，字符串。字符串类型取决于当前项目的时间显示格式。
- fps：每秒帧数，浮点值。
- isDuration：选填。true，则(从 0 帧开始计算)。false(默认值)，则从项目的起始帧开始计算(请参阅 Project displayStartFrame 属性)。

返回：浮点值，以秒为单位。

```javascript
var a = currentFormatToTime("1:25", 30);
alert(a);
```

1 秒 25 帧（30 帧每秒） = 1.8 秒（具体显示格式根据项目设置，比如丢帧时，返回 1.8333）

![](https://mir.yuelili.com/2021/07/3b3549b35a370256a0b8d7dd97a54eb9.png)

### generateRandomNumber() 生成随机数

全名：generateRandomNumber()

描述：生成随机数。

注意：建议使用此函数，而不是 Math.random(比如在使用 setValue 时)。此方法可避免 Math.random 由于 CPU 多线程的并发问题，无法在 AE
CC 2015(13.5.x)中返回随机值的问题。

返回：浮点伪随机数，范围为[0，1]。

备注：在 After Effects 13.6(CC 2015)中添加了此函数

```javascript
// 随机改变所有图层的X位置（包括锁定图层） var myComp = app.project.activeItem;
var x = 0;
for (var i = 1; i <= myComp.numLayers; i++) {
  // 如果用Math.random(),  不会正常工作;
  // 比如 x = 400 * (Math.random()) – 200;
  // 用generateRandomNumber()替换
  x = 400 * generateRandomNumber() - 200;
  var currentPos = myComp.layer(i).property("Position").value;
  myComp
    .layer(i)
    .property("Position")
    .setValue([currentPos[0] + x, currentPos[1]]);
}
```

有 2 个"我"文字图层，并且 x 位置发生了变化

![](https://mir.yuelili.com/2021/07/e953884b1f9e64d7d293498c74976a94.png)

### isValid() 判断对象是否存在

全名：isValid(obj)

描述：判断 AE 对象是否仍然存在(例如，合成，图层，蒙版等)。某些操作(例如 PropertyBase.moveTo())可能会使对象分配的现有变量无效。使用此函数，可以在访问这些变量之前，测试是否仍然有效。

参数：

- obj ，要检查的 After Effects 对象。

返回：布尔值。

```javascript
//在第1个图层上画3个蒙版，运行以下脚本 var layer =
app.project.activeItem.layer(1); // 记得要画3个蒙版，不然崩了勿怪
alert(isValid(layer)); //此时弹出"True＂
var mask1 = layer.mask(1);
var mask2 = layer.mask(2);
var mask3 = layer.mask(3);
mask3.moveTo(1); // 把第3个蒙版置顶
alert(isValid(mask1)); //弹出"false"; 其他蒙版同理
```

蒙版 3 跑到了顶部，蒙版 1 就不不是第 1 个蒙版了。故返回 false

![](https://mir.yuelili.com/2021/07/3212ec842d0ca58f7999123b705ce12a.png)

### timeToCurrentFormat() 秒转帧

全名：timeToCurrentFormat(time, fps[, isDuration])

描述：根据指定的帧速率，将时间秒转换为帧；

示例：时间为 0.5 秒，帧速率是 24 fps，则返回帧 0:00:12(当项目设置为显示为时间码时)。如果帧速率为 30
fps，则返回帧为 0:00:15。时间码字符串的格式由项目设置确定。(请参阅 Project displayStartFrame 属性)。

参数：

- formattedTime：帧数，字符串。字符串类型取决于当前项目的时间显示格式。
- fps：每秒帧数，浮点值。
- isDuration：选填。true，则(从 0 帧开始计算)。false(默认值)，则从项目的起始帧开始计算(请参阅 Project displayStartFrame 属性)。

返回：项目当前时间显示格式的字符串。

```javascript
t = timeToCurrentFormat(1.8,20) alert(t)
```

1.8 秒在 20 帧里，等于 1 秒 16 帧（数值是 1:16，具体显示格式根据项目设置，比如丢帧与不丢帧）  
![](https://mir.yuelili.com/2021/07/87b59e8d0100a695c76545e44fe15b70.png)

### write() 写入信息面板

全名：write(text)

描述：将文字写入“信息”面板，结尾不加换行符。

参数：

- text：要写入的字符串。如果太长（基于信息面板长度），则后面不显示。

返回：无。

```javascript
write("要写进去了");
write("要写更多进去了");
```

![](https://mir.yuelili.com/2021/07/949f6c253285fd6df81137ed88cf140c.png)

### writeLn() 写入信息面板

全名：writeLn(text)

描述：将信息写入“信息”面板，并在末尾添加换行符。

参数：

- text：要写入的字符串。

返回：无。

```javascript
writeLn("要写进去了");
writeLn("要写更多进去了");
```

![](https://mir.yuelili.com/2021/07/e7f98a38cf409bc3506e75d406b7bb2c.png)
