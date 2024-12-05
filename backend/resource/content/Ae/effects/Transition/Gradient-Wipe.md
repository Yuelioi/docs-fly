---
title: 渐变擦除
order: 13
---

# Gradient Wipe - 渐变擦除

## 简述

基于另一个图层的亮度值进行过渡。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=79&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Gradient_Wipe.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Gradient_Wipe_cn.png)

| Gradient Wipe         | 渐变擦除   | グラデーションワイプ   |                         |                |
| --------------------- | ---------- | ---------------------- | ----------------------- | -------------- |
| Transition Completion | 过渡完成   | 変換終了               |                         |                |
| Transition Softness   | 过渡柔和度 | 変換の柔らかさ         |                         |                |
| Gradient Layer        | 渐变图层   | グラデーションレイヤー |                         |                |
| Gradient Placement    | 渐变位置   | グラデーション置き換え | Tile Gradient           | 拼贴渐变       |
|                       |            |                        | Center Gradient         | 中心渐变       |
|                       |            |                        | Stretch Gradient to Fit | 伸缩渐变以适合 |
| Invert Gradient       | 反转渐变   | グラデーションを反転   |                         |                |

## 参数详解

### 过渡柔和度

每个像素渐变的程度。如果此值为 0%，则应用了该效果的图层中的像素将是完全不透明或完全透明。如果此值大于 0%，则在过渡的中间阶段像素是半透明的。

### 渐变位置

渐变图层的像素如何映射到应用了该效果的图层中的像素：

### 拼贴渐变

使用渐变图层的多个平铺副本。

### 中心渐变

在图层中心使用渐变图层的单个实例。

### 伸缩渐变以适合

在水平和垂直方向调整渐变图层的大小以适合图层的全部区域。

### 反转渐变

反转渐变图层的影响；渐变图层中的较浅像素以低于较深像素的“过渡完成”值创建透明度。

## 案例

<iframe src="https://player.bilibili.com/player.html?bvid=BV18E411C7qv&page=1&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>
