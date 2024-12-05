---
order: 26
title: 极坐标
---

# Polar Coordinates - 极坐标

## 简述

极坐标效果可扭曲图层，具体方法是将图层 (x,y)
坐标系中的每个像素调换到极坐标中的相应位置，反之亦然。此效果会产生反常的和令人惊讶的扭曲，扭曲结果根据选择的图像和控件的不同而有很大不同。标准坐标系通过从原点开始测量水平距离（X
轴）和垂直距离（Y 轴）来指定点。每个点均以 (x,y) 形式进行指定。极坐标系通过从原点开始测量半径长度 (r) 并从 X 轴开始测量其角度 ()
来指定点。每个点均以 (r,) 形式指定。

## 效果展示

![](https://www.yuelili.com/2021/07/cc7dbbb1d5fee61565d5bf773a95d994.jpeg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=42&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

### 原理

<iframe src="https://player.bilibili.com/player.html?bvid=BV1oq4y177gW&page=1&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Polar_Coordinates.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Polar_Coordinates_cn.png)

| Polar Coordinates  | 极坐标   | 極座標     |               |            |     |
| ------------------ | -------- | ---------- | ------------- | ---------- | --- |
| Interpolation      | 插值     | 補間       |               |            |     |
| Type of Conversion | 转换类型 | 変換の種類 | Rect to Polar | 矩形到极线 |     |
|                    |          |            | Polar to Rect | 极线到矩形 |     |

## 参数详解

![](https://www.yuelili.com/2021/07/f248377e62e84ef0e2dad353beec2fdf.jpeg)

### 插值

转换的程度

### 转换类型

极坐标的转换方式

**矩形到极线**

从垂直坐标系转极坐标系（如示例）

![](https://cdn.yuelili.com/20211223013805.png)

**极线到矩形**

从极坐标系转垂直坐标系

![](https://cdn.yuelili.com/20211223013819.png)

## 案例

### 高级百叶窗效果

极坐标

百叶窗

极坐标 2

变换

![](https://cdn.yuelili.com/20211223014604.png)
