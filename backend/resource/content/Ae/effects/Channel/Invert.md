---
order: 7
title: 反转
---

# Invert - 反转

## 简述

反转效果可反转图像的颜色信息。

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=31&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00432.jpg)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Invert.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Invert_cn.png)

| Invert  | 反转 | 反転       |                        |            |                    |
| ------- | ---- | ---------- | ---------------------- | ---------- | ------------------ |
| Channel | 通道 | チャンネル | RGB                    | RGB        | RGB                |
|         |      |            | Red                    | 红色       | red                |
|         |      |            | Green                  | 绿色       | green              |
|         |      |            | Blue                   | 蓝色       | blue               |
|         |      |            | HLS                    | HLS        | HLS                |
|         |      |            | Hue                    | 色相       | 色相               |
|         |      |            | Lightness              | 亮度       | 輝度               |
|         |      |            | Saturation             | 饱和度     | 彩度               |
|         |      |            | YIQ                    | YIQ        | YIQ                |
|         |      |            | Luminance              | 明亮度     | ルミナンス         |
|         |      |            | In phase chrominance   | 相内彩色度 | I 信号クロミナンス |
|         |      |            | Quadrature Chrominance | 求积彩色度 | Q 信号クロミナンス |
|         |      |            | Alpha                  | Alpha      | アルファ           |

## 参数详解

![](https://mir.yuelili.com/user/AE/effects/ext/image00433.jpg)

### 通道

要反转的一个或多个通道。每个项目组均在特定颜色空间中运行，因此可以反转该颜色空间中的整个图像，也可以仅反转单个通道。

**-RGB/红色/绿色/蓝色**

RGB 用于反转所有三个附加的颜色通道。“红色”、“绿色”和“蓝色”各自用于反转单个颜色通道。

**-HLS/色相/亮度/饱和度**

HLS 用于反转所有三个计算的颜色通道。“色相”、“亮度”和“饱和度”各自用于反转单个颜色通道。

**-YIQ/明亮度/相内彩色度/求积彩色度**

YIQ 用于反转所有三个 NTSC 明亮度和彩色度通道。Y（“明亮度”）、I（“相内彩色度”）和 Q（“求积彩色度”）各自用于反转单个通道。

**-Alpha**

反转图像的 Alpha 通道。Alpha 通道不是颜色通道；它用于指定透明度。

### 与原始图像混合

效果图像的透明度。效果图像与原始图像混合的结果，并在上面合成效果图像结果。此值设置得越高，此效果对图层的影响越小。例如，如果将此值设置为
100%，则此效果在图层上不会产生明显结果；如果将此值设置为 0%，则不会隐约显出原始图像。

反转效果使用 GPU 加速来实现更快的渲染。
