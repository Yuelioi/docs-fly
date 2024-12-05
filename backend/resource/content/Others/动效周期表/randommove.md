---
title: "【动效周期表】RandomMove 随机移动"
date: "2020-06-07"
categories: 
  - "archive"
---

<table style="border-collapse: collapse; width: 100%;"><tbody class="table1"><tr><td style="width: 25.4125%;"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove.gif"></td><td style="width: 93.8898%;"><h2 style="font-size: 36pt;">RandomMove</h2><div></div><span style="font-size: 24pt;">随机移动</span><div></div>该函数每个帧返回一个随机值。 应用于图层的position属性后，每一帧移动到不同的位置， 从而导致不连续的剧烈运动，如左图所示。 除了position属性，还有多种用途，例如用于文本图层。</td></tr></tbody></table>

## 制作方法

在AE中使用时：

1.如果要随机移动对象，请在图层的position属性设置表达式：random（100）;。

2.如果要使用随机数，可以编写一个表达式，但是可以使用“效果” /“文本” /“编号”轻松创建它。

## 应用

1.在圆的扩展/收缩属性中编写一个随机表达式，以便大小随机变化。 复制这些图层，降低了图层的不透明度，像是扬声器振动。

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/randommove/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/Scale/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Scale.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove-Ex002.gif"></td></tr></tbody></table>

2.以下表达式放在文本图层的源文本。 生成一个随机整数。

```javascript
y = random(100000);
x = Math.round(y);
if (x <0) str = "000000";
else if (x < 10) str = "00000";
else if (x < 100) str = "0000";
else if (x < 1000) str = "000";
else if (x < 10000) str = "00";
else if (x < 100000) str = "0";
str = x+str;
```

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/randommove/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove-Ex001.gif"></td></tr></tbody></table>

## 例

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove-Ex001.gif"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RandomMove-Ex002.gif"></td></tr></tbody></table>

## 相似/相关

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/wigglemove/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleMove.gif"></a></td><td><a href="https://yuelili.com/archive/WiggleRotate/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleRotate.gif"></a></td><td><a href="https://yuelili.com/archive/WiggleScale/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleScale.gif"></a></td></tr></tbody></table>
