---
order: 8
title: 最小 最大
---

# Minimax - 最小 最大

## 简述

最小值/最大值效果可为像素的每个通道分配指定半径内该通道的最小值或最大值。

此效果可用于扩大或减小遮罩。例如，使用最小值效果和半径 1 时，黑色环绕的白色纯色区域各边会收缩一个像素。

图层的品质设置不会影响最小值/最大值效果。

此效果适用于 8-bpc 和 16-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20211212154802.gif)

### 视频教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=3&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 参数中英日对照

| Minimax           | 最小/最大    | 最大/最小          |                      |                  |          |
| ----------------- | ------------ | ------------------ | -------------------- | ---------------- | -------- |
| Operation         | 操作         | 操作               | Minimum              | 最小值           | 最小     |
|                   |              |                    | Maximum              | 最大值           | 最大     |
|                   |              |                    | Minimum then maximum | 先最小值再最大值 |          |
|                   |              |                    | Maximum Then Minimum | 先最大值再最小值 |          |
| Radius            | 半径         | 半径               |                      |                  |          |
| Channel           | 通道         | チャンネル         | Color                | 颜色             | カラー   |
|                   |              |                    | Alpha and color      | Alpha 和颜色     |          |
|                   |              |                    | Red                  | 红色             | red      |
|                   |              |                    | Green                | 绿色             | green    |
|                   |              |                    | Blue                 | 蓝色             | blue     |
|                   |              |                    | Apha                 | Alpha            | アルファ |
| Direction         | 方向         | 方向               | Horizontal vertical  | 水平和垂重       |          |
|                   |              |                    | Just horizontal      | 仅水平           |          |
|                   |              |                    | Just vertical        | 仅垂直           |          |
| Dont Shrink Edges | 不要收缩边缘 | エッジを縮小しない |                      |                  |          |

## 参数图片参考

![](http://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Minimax.png)
![](http://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Minimax_cn.png)

### 操作

最小值效果：可为像素的每个通道分配指定半径内指定通道的最小值。

最大值效果：可为像素的每个通道分配最大值。

先最小值再最大值效果：先执行最小值操作，然后执行最大值操作。

先最大值再最小值效果：先执行最大值操作，然后执行最小值操作。

### 方向

扫描的轴。“水平和垂直”选项用于扫描各个方向。

### 不缩小边缘

对于没有 α 通道的图像，将[操作]设置为[最小]，将[通道]设置为[颜色]的情况下，缩小的部分使图像周围变黑，选择此选项，不会出现该现象
