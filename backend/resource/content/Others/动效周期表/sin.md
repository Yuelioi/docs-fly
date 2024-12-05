---
title: "【动效周期表】Sin 三角函数"
date: "2020-06-07"
categories: 
  - "archive"
---

<table style="border-collapse: collapse; width: 100%;"><tbody class="table1"><tr><td style="width: 25.4125%;"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Sin.gif"></td><td style="width: 93.8898%;"><h2 style="font-size: 36pt;">Sin</h2><div></div><span style="font-size: 24pt;">三角函数</span><div></div>Sin画波浪的运动。 当您要表达平稳连续运动时使用。 它不仅可以应用于位置，还可以应用于旋转和缩放等元素。 通过添加绝对值等数学元素，可实现广泛的应用范围。</td></tr></tbody></table>

## 制作方法

在AE中使用时，

1.在属性表达式中书写：Math.sin(time \* Math.PI)等。

## 应用

1.创建一个圆，移动锚点，然后将其旋转到“轨道”状态。

```javascript
Math.sin(time * 2 * Math.PI / 2)* 80 + 180;// 在rotation属性的表达式
```

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/repeattrim/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Sin.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/Orbit/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Orbit.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Sin-Ex001.gif"></td></tr></tbody></table>

## 例

![](https://mir.yuelili.com/user/AE/mg/foxcodex/Sin-Ex001.gif)

## 相似/相关

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/symrotate/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/SymRotate.gif"></a></td><td><a href="https://yuelili.com/archive/SymScale/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/SymScale.gif"></a></td><td><a href="https://yuelili.com/archive/Orbit/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Orbit.gif"></a></td></tr></tbody></table>
