---
title: "【动效周期表】RepeatTransform 连续变换"
date: "2020-06-03"
categories: 
  - "archive"
---

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RepeatTransform.gif"></td><td><h2 style="font-size: 36pt;">RepeatTransform</h2><div></div><span style="font-size: 24pt;">连续变换</span><div></div>连续执行“形状变换”的运动，该变换定期使图形变形。 变换是连续的，因此看起来很复杂，但是做起来很简单。 只需复制图层并移动时间即可复制。 与不透明度兼容，因为它通常由堆叠图层使用。 适用于诸如抽象图形，UI和加载屏幕之类的动作。</td></tr></tbody></table>

## 作成法

重复的形状图层定期进行交错变形。

## 应用

1.重复变换，连续变换图形，再乘以MaskWipe（蒙版擦除）。 确切地说，将相同的变形图放置在末端并用alpha遮罩切出。 看起来非常复杂，但仔细观察，其实具有相同的动作。

<table><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/repeattransform/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RepeatTransform.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/maskwipe/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/MaskWipe.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/MaskWipe-Ex001.gif"></td></tr></tbody></table>

## 例

![](https://mir.yuelili.com/user/AE/mg/foxcodex/MaskWipe-Ex001.gif)

## 相似/相关

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/shapetransform/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/ShapeTransform.gif"></a></td><td><a href="https://yuelili.com/archive/Rhythm/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Rhythm.gif"></a></td><td><a href="https://yuelili.com/archive/RepeatTrim/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/RepeatTrim.gif"></a></td></tr></tbody></table>
