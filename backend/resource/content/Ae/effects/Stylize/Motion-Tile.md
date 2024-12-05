---
title: 动态拼贴
order: 19
---

# Motion Tile - 动态拼贴

## 简述

动态拼贴效果可跨输出图像复制源图像。如果已启用运动模糊，则在更改拼贴的位置时，此效果会使用运动模糊来使移动更明显。

有点像图层“中继器”

此效果适用于 8-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20211228144928.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=12&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

### ![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Motion_Tile.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Motion_Tile_cn.png)

参数详解 #

### 拼贴中心

主要拼贴的中心。

![](https://cdn.yuelili.com/20211228145338.png)

### 拼贴宽度、拼贴高度

拼贴的尺寸，显示为输入图层尺寸的百分比。

50 = 原始图像缩小 50%

### 输出宽度、输出高度

输出图像的尺寸，显示为输入图层尺寸的百分比。

300 = 复制 3 个

### 镜像边缘

翻转邻近拼贴，以形成镜像图像。如果“相位”设置为 0，则选择此选项将使用周围的拼贴对图层边缘使用镜像功能。

### 相位

拼贴的水平或垂直位移。

| 水平（180°）                                    | 垂直（180°） |
| ----------------------------------------------- | ------------ |
| ![](https://cdn.yuelili.com/20211228145526.png) |

![](https://cdn.yuelili.com/20211228145500.png)

### 水平位移

使拼贴水平（而非垂直）位移。需要移动上一个参数（相位）

## 案例

被 AK 用来修复边缘像素。005.摄像机抖动
