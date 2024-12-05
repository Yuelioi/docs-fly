---
title: 发光
order: 17
---

# Glow - 发光

## 简述

通过找到图像的较亮区域，使其周围的像素变亮，以创建漫射的发光光环。

发光效果也可以模拟明亮的光照对象的过度曝光。您可以使发光基于图像的原始颜色，或基于其 Alpha 通道。

基于 Alpha 通道的发光仅在不透明和透明区域之间的图像边缘产生漫射亮度。您也可以使用发光效果创建两种颜色（A 和 B
颜色）之间的渐变发光，以及创建循环的多色效果。

在“最佳”品质渲染发光效果可更改图层的外观。如果使用 Adobe Photoshop
任意映射为发光着色，则此更改特别真实。请务必在渲染效果之前以“最佳”品质预览。

\*GPU 加速

## 效果展示

![](https://cdn.yuelili.com/20220102004445.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=109&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Glow.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Glow_cn.png)

| Glow              | 发光         | グロー           |                         |              |                |     |     |     |
| ----------------- | ------------ | ---------------- | ----------------------- | ------------ | -------------- | --- | --- | --- |
| Glow Based On     | 发光基于     | グロー基準       | Alpha Channel           | Alpha 通道   |                |     |     |     |
|                   |              |                  | Color Channels          | 颜色通道     |                |     |     |     |
| Glow Threshold    | 发光阔值     |                  |                         |              |                |     |     |     |
| Glow Radius       | 发光半径     | グローの半径     |                         |              |                |     |     |     |
| Glow Intensity    | 发光强度     | グロー強度       |                         |              |                |     |     |     |
| Composite Origina | 合成原始项目 | 元を合成         | On Top                  | 顶端         |                |     |     |     |
|                   |              |                  | Behind                  | 后面         |                |     |     |     |
|                   |              |                  | None                    | 无           | なし           |     |     |     |
| Glow Operation    | 发光操作     | グロー操作       |                         |              |                |     |     |     |
| Glow Colors       | 发光颜色     | グローのカラー   | Original Colors         | 原始颜色     |                |     |     |     |
|                   |              |                  | A & B Colors            | A 和 B 颜色  |                |     |     |     |
|                   |              |                  | Arbitrary Map           | 任意映射     |                |     |     |     |
| Color Looping     | 颜色循环     | カラールーピング | Sawtooth A>B            | 锯齿 A>B     |                |     |     |     |
|                   |              |                  | Sawtooth B>A            | 锯齿 B>A     |                |     |     |     |
|                   |              |                  | TriAngle A>B>A          | 三角形 A>B>A |                |     |     |     |
|                   |              |                  | Triangle B>A>B          | 三角形 B>A>B |                |     |     |     |
| Color Loops       | 颜色循环     | カラールーピング |                         |              |                |     |     |     |
| Color Phase       | 色彩相位     | カラーフェーズ   |                         |              |                |     |     |     |
| A& B Midpoint     | A 和 B 中点  |                  |                         |              |                |     |     |     |
| Color A           | 颜色 A       |                  |                         |              |                |     |     |     |
| Color B           | 颜色 B       |                  |                         |              |                |     |     |     |
| Glow Dimensions   | 发光维度     | グローの方向     | Horizontal and Vertical | 水平和垂直   | 水平および垂直 |     |
|                   |
|                   |              |                  | Horizontal              | 水平         | 水平           |     |     |     |
|                   |              |                  | Vertical                | 垂直         | 垂直           |     |     |     |

## 参数详解

### 发光基于

基于颜色发光还是透明度（更均匀）。

### 发光阈值

高于该阈值才会发光。值越低，发光区域越多

### 发光半径

从明亮区域开始延伸的距离，以像素为单位。

较大的值会产生漫射发光；较小的值会产生锐化边缘的发光。

### 发光强度

发光的亮度。

### 合成原始项目

指定如何合成效果结果和图层。

- 顶端：将发光效果放在图像顶端，以便使用为“发光操作”选择的混合方法。
- 后面：将发光效果放在图像后面，从而创建逆光结果。
- 无：只有发光效果，去除原图像。

### 发光操作

其实就是发光效果与原始图层的混合模式

### 发光颜色

发光的颜色。

“A 和 B 颜色”用于使用“颜色 A”和“颜色 B”控件指定的颜色，创建渐变发光。

### 颜色循环

选择“A 和 B 颜色”作为“发光颜色”的值时，使用的渐变曲线的形状。

### 颜色循环

可在选择两个或更多循环时，创建发光的多色环。值为 1，则循环显示为“发光颜色”指定的渐变（或任意图）。

![](https://cdn.yuelili.com/20220102003443.png)

### 色彩相位

在颜色循环中，开始循环的位置。默认情况下为第 1 个循环。

### A 和 B 中点

用于指定渐变中使用的两种颜色所占百分比。

值越低，使用 A 颜色越少。反之， B 颜色越少。

### 颜色 A、颜色 B

在选择“A 和 B 颜色”作为“发光颜色”的值时，发光的颜色。

A：发光边缘（未循环前）

B：发光核心（未循环前）

### 发光维度

发光是水平的、垂直的，还是两者兼有。

### 其他

加载发光形状预设请看教程最后 1 分钟，不常用倒是
