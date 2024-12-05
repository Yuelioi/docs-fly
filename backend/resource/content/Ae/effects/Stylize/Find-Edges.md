---
title: 查找边缘
order: 16
---

# Find Edges - 查找边缘

## 简述

查找边缘效果可确定具有大过渡（对比强烈）的图像区域，并可强调边缘，看似原始图像的草图。

边缘可在白色背景上显示为深色线条，也可在黑色背景上显示为彩色线条。

此效果适用于 8-bpc 颜色。

可用 GPU 加速渲染。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=89&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Find_Edges.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Find_Edges_cn.png)

| Find Edges | 查找边缘 | 輪郭検出 |                     |                |                    |
| ---------- | -------- | -------- | ------------------- | -------------- | ------------------ |
|            |          |          | Invert              | 反转           | 反転               |
|            |          |          | Blend With Original | 与原始图像混合 | 元の画像とブレンド |

## 参数详解

### 反转

在找到边缘之后反转图像。如果不选择“反转”，则边缘在白色背景上显示为暗线条。如果选择此控件，则边缘在黑色背景上显示为亮线条。

## 教程

过滤出来的线框增加颜色（tint）发光，再与原图像叠加

![](https://cdn.yuelili.com/20220102002016.png)

### 友情链接

tint 色调：用于去色

curves 色阶：用于控制范围
