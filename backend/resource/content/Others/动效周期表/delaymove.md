---
title: "【动效周期表】DelayMove 延迟移动"
date: "2020-06-07"
categories: 
  - "archive"
---

<table style="border-collapse: collapse; width: 100%;"><tbody class="table1"><tr><td style="width: 25.4125%;"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/DelayMove.gif"></td><td style="width: 93.8898%;"><h2 style="font-size: 36pt;">DelayMove</h2><div></div><span style="font-size: 24pt;">延迟移动</span><div></div>将对象的移动延迟指定秒数的效果。 您可以复制几图层并逐渐延迟，使它们看起来像是跟随。 很少使用，但请记住，您也可以做到。</td></tr></tbody></table>

## 制作方法

在Aftereffects中使用时，

1.在关键帧的属性表达式中输入：

```javascript
valueAtTime(time-0.2) // 延迟0.2秒
```

## 应用

1. WiggleMove图层副本，每0.2秒延迟一次。

位置属性表达式

类似于thisComp.layer("shape layer 1").transform.position.valueAtTime(time-0.2)

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/repeattrim/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/DelayMove.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/WiggleMove/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleMove.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/DelayMove-Ex001.gif"></td></tr></tbody></table>

## 例

![](https://mir.yuelili.com/user/AE/mg/foxcodex/DelayMove-Ex001.gif)

## 相似/相关

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/motionblur/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/MotionBlur.gif"></a></td><td><a href="https://yuelili.com/archive/WiggleMove/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleMove.gif"></a></td></tr></tbody></table>
