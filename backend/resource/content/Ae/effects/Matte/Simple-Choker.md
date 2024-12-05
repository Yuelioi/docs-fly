---
title: 简单阻塞工具
order: 4
---

# Simple Choker - 简单阻塞工具

## 简述

简单阻塞工具效果可以小缩小或扩展遮罩边缘，以便创建更整洁的遮罩。当您想用颜色键提取一些颜色并将它们与其他图像组合时，这很有效。

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=23&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://cdn.yuelili.com/20211225215635.png)

![](https://cdn.yuelili.com/20211225215640.png)

| Simple Choker | 简单阻塞工具 | チョーク       |              |          |        |
| ------------- | ------------ | -------------- | ------------ | -------- | ------ |
| View          | 视图         | 表示           | Final Outout | 最终输出 |        |
|               |              |                | Matte        | 遮罩     | マット |
| Choke Matte   | 阻塞遮罩     | チョークマット |              |          |        |

## 参数详解

### 视图

最终输出：显示应用此效果的图像，

遮罩：提供黑白视图以显示生效范围，黑色区域（表示透明度）和白色区域（表示不透明度）。

### 阻塞遮罩

调整 alpha 遮罩（透明部分）的范围。数值越大范围越小，数值越负范围越大。

![](https://cdn.yuelili.com/20211225215853.png)

## 案例：调整绿幕的裁剪范围

当然，你可以直接用 keylight 来扣，只是演示下这个效果的作用。

颜色键：大范围扣绿，但是人体会留有空洞

简单阻塞：增大阻塞，以弥补人体空洞，但是人体边缘又会变绿

![](https://cdn.yuelili.com/20211225215903.png)

颜色键：大范围扣绿，但是人体会留有空洞

简单阻塞：增大阻塞，以弥补人体空洞，但是人体边缘又会变绿

颜色键 2：再次扣人体边缘绿色

简单阻塞：细化边缘

![](https://cdn.yuelili.com/20211225220425.png)
