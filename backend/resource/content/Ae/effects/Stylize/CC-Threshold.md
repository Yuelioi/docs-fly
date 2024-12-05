---
title: CC阈值
order: 10
---

# CC Threshold - CC阈值

## 简述

阈值的进阶版。

基于一个限制参数，裁剪图像信息。有更多参数控制。

## 效果展示

![](https://cdn.yuelili.com/20211228161445.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=17&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-CC_Threshold.png)

## 参数详解

### 阈值

分割点，高于该值的变纯白，低于该值的变纯黑。（RGB 则是变成纯 R、G、B）

### 通道 Channel

| 原图                                            | 基于亮度 | 基于 RGB | 基于饱和度 | Alpha |
| ----------------------------------------------- | -------- | -------- | ---------- | ----- |
| ![](https://cdn.yuelili.com/20211228161413.png) |
| ![](https://cdn.yuelili.com/20211228161445.png) |
| ![](https://cdn.yuelili.com/20211228161525.png) |
| ![](https://cdn.yuelili.com/20211228161547.png) |

![](https://cdn.yuelili.com/20211228162302.png)

## 案例

基于 Alpha 加 模糊效果（可以改变 Alpha） 就可以做图层边缘融化
