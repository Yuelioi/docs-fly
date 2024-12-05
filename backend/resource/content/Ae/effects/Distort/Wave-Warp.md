---
title: 波形变形
order: 37
---

# Wave Warp - 波形变形

## 简述

可在图像上产生波形移动的效果，可以创建定期重复的波浪形扭曲动画。有各种不同的波形形状，包括正方形、圆形和正弦波形。

波形变形效果可在时间范围内以定速（无关键帧或表达式）自动设置动画。要改变速度，请设置关键帧或表达式。

此效果适用于 8-bpc 和 16-bpc 颜色。

### 特点

类型多、可以超出图层边界、

可以做波浪、旗子摇摆、烟雾、头发

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=65&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Wave_Warp.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Wave_Warp_cn.png)

|Wave Warp | 波形变形 | 波形ワープ | | | ||
![](https://cdn.yuelili.com/20211225123610.png)  
|---|---|---|---|---|---|---|
|Wave Type | 波浪类型 | 波形の種類 | Sine | 正弦 | サイン ||
![](https://cdn.yuelili.com/20211225123633.png)  
|| | | Square | 正方形 | Square ||
![](https://cdn.yuelili.com/20211225124115.png)  
|| | | Triangle | 三角形 | 三角形数 ||
![](https://cdn.yuelili.com/20211225124142.png)  
|| | | Sawtooth | 锯齿 | ジグザグ | ![](https://cdn.yuelili.com/20211225124207.png)|
|| | | Circle | 圆形 | 円 | ![](https://cdn.yuelili.com/20211225124221.png)|
|| | | Semicircle | 半圆形 | | ![](https://cdn.yuelili.com/20211225124239.png)|
|| | | Uncircle | 非圆形 | | ![](https://cdn.yuelili.com/20211225124251.png)|
|| | | Noise | 杂色 | ノイズ | ![](https://cdn.yuelili.com/20211225124302.png)|
|| | | Smooth Noise | 平滑杂色 | ||
![](https://cdn.yuelili.com/20211225124318.png)  
|Wave Height | 波形高度 | 波形の高さ | | | ||
|Wave Width | 波形宽度 | 波紋の幅 | | | ||
|Direction | 方向 | 方向 | | | ||
|Wave Speed | 波形速度 | 波紋の速度 | | | ||
|Pinning | 固定 | 固定 | None | 无 | なし ||
|| | | All Edges | 全部边缘 | ||
|| | | Center | 居中 | 中央 ||
|| | | Left Edge | 左侧边缘 | ||
|| | | Top Edge | 顶部边缘 | ||
|| | | Right Edge | 右侧边缘 | ||
|| | | Bottom Edge | 底部边缘 | ||
|| | | Horizontal Edges | 水平边缘 | ||
|| | | Vertical Edges | 垂直边缘 | ||
|Phase | 相位 | フェーズ | | | ||
|Antialiasing(Best Quality) | 消除锯齿（最佳品质） | アンチエイリアス(最高画質) | Low | 低 | 低 ||
|| | | Medium | 中 | 中 ||
|| | | High | 高 | 高 ||

## 参数详解

### 波形类型

您可以从 9 种类型中选择波浪的形状

![](https://cdn.yuelili.com/20211225121619.png)

### 波形高度

波峰之间的距离，以像素为单位。

### 波形宽度

调整波浪的宽度，以像素为单位。数值越大波浪越大。数值越小，浪数越多。

### 方向

波形在图像中移动的方向。例如，值为 225° 时，波形从右上角移动到左下角。

### 波形速度

指定波浪移动的速度，也可以算循环的周期。数字越大，动画动画越快。数值为 0 将波浪将完全静止，负数波浪的方向将反转。

值为 1：则 1、2、3、4 等整数秒状态一样

值为 2，则每 1/2，也就是 0.5 秒状态一样

### 固定

要固定的边缘，这些边缘的像素不生效。比如旗子一般固定左边缘

### 相位

开始波形循环的点。例如，值为 0°，则在波下坡的中点开始波形；值为 90°，则在波谷的最低点开始。

### 消除锯齿

设置在图像上完成的消除锯齿的数量，或边缘平滑量。在许多情况下，较低的设置可产生令人满意的结果；较高的设置会显著增加渲染时间。仅当图层品质设置为“最佳”时，才执行“消除锯齿”。

## 案例：火箭拖尾

示例展示了火箭在冒烟的同时飞行的情况。

![](https://cdn.yuelili.com/20211225121737.png)
