---
title: 通道合成器
order: 5
---

# Channel Combiner - 通道合成器

## 简述

把指定的通道换成别的通道。根据图像指定色相、亮度、RGB 等通道，将该通道 α 化或变换到其他通道。也可以从其他图像中提取频道。

您可以使用此效果将任何通道视为灰度图像，具体方法是从“自”菜单中选择通道，然后从“到”菜单中选择“仅亮度”。

此效果适用于 8-bpc 颜色。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00414.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=115&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Channel_Combiner.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Channel_Combiner_cn.png)

| Channel Combiner | 通道合成器     | チャンネルコンバイナー |                           |                |            |
| ---------------- | -------------- | ---------------------- | ------------------------- | -------------- | ---------- |
| Use 2nd Layer    | 使用第二个图层 | 2 つめのレイヤーを使用 |                           |                |            |
| Source Layer     | 源图层         | ソースレイヤー         |                           |                |            |
| From             | 自             | 変更するカラー         | RGB to HLS                | RGB 到 HLS     |            |
|                  |                |                        | HLS to RGB                | HLS 到 RGB     |            |
|                  |                |                        | RGB to yuv                | RGB 到 YUV     |            |
|                  |                |                        | YUV to rgB                | YUV 到 RGB     |            |
|                  |                |                        | Straight to Premultiplied | 直接到预乘     |            |
|                  |                |                        | Red                       | 红色           | red        |
|                  |                |                        | Green                     | 绿色           | green      |
|                  |                |                        | Blue                      | 蓝色           | blue       |
|                  |                |                        | Alpha                     | Alpha          | アルファ   |
|                  |                |                        | Hue                       | 色相           | 色相       |
|                  |                |                        | Lightness                 | 亮度           | 輝度       |
|                  |                |                        | Luminance                 | 明亮度         | ルミナンス |
|                  |                |                        | Saturation                | 饱和度         | 彩度       |
|                  |                |                        | Saturation Multiplied     | 饱和度（合成） |            |
|                  |                |                        | Min RGB                   | RGB 最小值     |            |
|                  |                |                        | Max RGB                   | RGB 最大值     |            |
| To               | 至             | 変更後のカラー         | Red                       | 红色           | red        |
|                  |                |                        | Green                     | 绿色           | green      |
|                  |                |                        | Blue                      | 蓝色           | blue       |
|                  |                |                        | Alpha                     | Alpha          | アルファ   |
|                  |                |                        | Hue                       | 色相           | 色相       |
|                  |                |                        | Lightness                 | 亮度           | 輝度       |
|                  |                |                        | Saturation                | 饱和度         | 彩度       |
|                  |                |                        | Red only                  | 仅红色         |            |
|                  |                |                        | Green only                | 仅绿色         |            |
|                  |                |                        | Blue only                 | 仅蓝色         |            |
|                  |                |                        | Alpha only                | 仅 Apha        |            |
|                  |                |                        | Hue only                  | 仅色相         |            |
|                  |                |                        | Lightness only            | 仅亮度         |            |
|                  |                |                        | Saturation only           | 仅饱和度       |            |
| Invert           | 反转           | 反転                   |                           |                |            |
| Solid Alpha      | 纯色 Alpha     | 不透明アルファ         |                           |                |            |

## 参数详解

### 使用第二个图层

从源图层检索值，源图层可以是合成中的任何图层。

### **自**

要用作输入的值；菜单中的前几项是多通道组合的输入和输出选项，因此不需要设置“到”值。  
饱和度（合成）使用饱和度值乘亮度值，其中亮度是到黑色或白色的最短距离。例如，与亮蓝色或纯蓝色像素相比，深蓝色或浅蓝色像素的值较低。此选项表示最常见的像素饱和度值视图。

- RGB 最小值：使用红色、绿色和蓝色通道值中的最低值。
- RGB 最大值：使用红色、绿色和蓝色通道值中的最高值。

### **至**

向其应用值的通道。

- 仅红色、仅绿色和仅蓝色：仅将值应用到一个通道，并将其他颜色通道设置为零。
- 仅 Alpha”：将值应用到 Alpha 通道，并将颜色通道设置为 1.0。
- 仅色相：应用的色相值会与 50% 亮度和 100% 饱和度组合。
- 仅亮度：应用的亮度值会与 0% 饱和度组合，随后不会影响色相。
- 仅饱和度：应用的饱和度值会与 0% 色相和 50% 亮度组合。

### 反转

反转（从 1.0 中减去）输出通道值。

### 纯色 Alpha

使整个图层的 Alpha 通道值为 1.0（完全不透明）。
