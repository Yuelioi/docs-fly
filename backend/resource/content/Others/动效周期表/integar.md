---
title: "【动效周期表】StepTime 步进时刻"
date: "2020-06-07"
categories: 
  - "archive"
---

<table style="border-collapse: collapse; width: 100%;"><tbody class="table1"><tr><td style="width: 25.4125%;"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Integar.gif"></td><td style="width: 93.8898%;"><h2 style="font-size: 36pt;">StepTime</h2><div></div><span style="font-size: 24pt;">步进时刻</span><div></div>一种不连续移动对象的效果。 可以像延时一样表现运动，特别适合钟表和齿轮的运动。 使用移动和旋转方法应该很有用。</td></tr></tbody></table>

## 制作方法

在“AE”中使用时：

1.对要移动的图层的属性应用表达式。

2.用Math.floor（time）四舍五入时间的小数点。

3.如果觉得Math.floor（time）很慢，把“ time”更改为“ time \* 10”或乘以整数改变速率。 在旋转属性中写入Math.floor（time \* 5）或\* 10来移动以上示例中的指针。

## 应用

1.使用StepTime复刻齿轮运动。 通过组合使用StepTime的齿轮和不使用StepTime的齿轮来模拟齿轮感觉。

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/integar/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Integar.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/wigglerotate/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/WiggleRotate.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Integar-Ex001.gif"></td></tr></tbody></table>

## 例 ![](https://mir.yuelili.com/user/AE/mg/foxcodex/Integar-Ex001.gif)

## 相似/相关 ![](https://mir.yuelili.com/user/AE/mg/foxcodex/Sin.gif)
