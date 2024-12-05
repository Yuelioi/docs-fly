---
title: 保留细节放大
order: 18
---

# Detail preserving Upscale - 保留细节放大

## 简述

可以在放大图像的同时保留图像中的细节。保留细节放大效果的处理速度慢于其他缩放方式，比如在变换属性组中使用图层的原生两次线性或两次立方缩放。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=72&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Detail_preserving_Upscale.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Detail_preserving_Upscale_cn.png)

| Detail preserving Upscale | 保留细节放大 | アップスケール（ディテールを保持） |                    |              |                                |
| ------------------------- | ------------ | ---------------------------------- | ------------------ | ------------ | ------------------------------ |
|                           |              |                                    | Fit to Comp Width  | 适合复合宽度 | コンポジションの幅に合わせる   |
|                           |              |                                    | Fit to Comp Height | 适合复合高度 | コンポジションの高さに合わせる |
|                           |              |                                    | Scale              | 缩放         | スケール                       |
|                           |              |                                    | Reduce Noise       | 减少杂色     | ノイズを軽減                   |
|                           |              |                                    | Detail             | 详细信息     | 詳細                           |
|                           |              |                                    | Alpha              | Alpha        | アルファ                       |

## 参数详解

### 适配合成宽度/高度

设置缩放百分比，使图层的宽度/高度与合成的宽度/高度相匹配。

![](https://cdn.yuelili.com/20211225191736.png)

_自动适合宽度_

如果原始图像尺寸大于合成尺寸，则无法使用 [适合合成宽度] 和 [适合合成高度] 功能。

### 缩放

放大放大，当然最小值为 100%。

### 减少杂色

在缩放计算之前去除杂色。增大该值，以免将杂色误认为应保留的细节。

![](https://cdn.yuelili.com/20211225192033.png)

### 详细信息

增加该值将产生锐化边缘和增加对比度的效果。

![](https://cdn.yuelili.com/20211225192318.png)

### Alpha

处理 Alpah 通道。[保留详细信息]比[双三次]更流畅，但渲染更差劲。前提你的图片有 alpha 信息，比如扣好的 logo

## 案例
