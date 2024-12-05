---
title: 中间值
order: 5
---

# Median - 中间值

## 简述

中间值效果将每个像素替换为另一像素，此像素具有指定半径的邻近像素的中间颜色值。当“半径”值较低时，此效果可用于减少某些类型的杂色。在“半径”值较高时，此效果为图像提供绘画风格的外观。

此效果适用于 8-bpc 、 16-bpc 颜色和 32-bpc。

与旧版的区别：支持 32bpc，其他没区别。（也许性能有提升？我猜的）

## 效果展示

| 正常图片                                        | 带 alpha 的，如文字（可以把周围圆滑，或者使像素渗入） |
| ----------------------------------------------- | ----------------------------------------------------- |
| ![](https://cdn.yuelili.com/20220103223140.png) |

![](https://cdn.yuelili.com/20220103223741.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=95&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

| Radius                   | 半径                | 半径 |
| ------------------------ | ------------------- | ---- |
| Operate On Alpha Channel | 在 Alpha 通道上运算 |      |

## 参数详解

### 半径

进行融合的半径范围

### 在 Alpha 通道上运算

如果是 alpha 的，记得勾选下

## 案例
