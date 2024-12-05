---
title: 通道模糊
order: 8
---

# Channel Blur - 通道模糊

## 简述

分通道，为每个 RGBA 通道设置模糊量。可以得到在图像的轮廓部分产生色差一样的效果

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00579.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=30&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Channel_Blur.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Channel_Blur_cn.png)

| Channel Blur     | 通道模糊     | ブラー(チャンネル) |                         |            |                |
| ---------------- | ------------ | ------------------ | ----------------------- | ---------- | -------------- |
| Red Blurriness   | 红色模糊度   | 赤ブラー           |                         |            |                |
| Green Blurriness | 绿色模糊度   | 緑ブラー           |                         |            |                |
| Blue Blurriness  | 蓝色模糊度   | 青ブラー           |                         |            |                |
| Alpha Blurriness | Alpha 模糊度 | アルファブラ一     |                         |            |                |
| Edge Behavior    | 边缘特性     | エッジ動作         |                         |            |                |
| Blur Dimensions  | 模糊方向     | ブラーの方向       | Horizontal and Vertical | 水平和垂直 | 水平および垂直 |
|                  |              |                    | Horizontal              | 水平       | 水平           |
|                  |              |                    | Vertical                | 垂直       | 垂直           |

## 参数详解

首先科普一下下,生活中看到的各种颜色,是由红绿蓝三原色组成的

而通道模糊是只针对某颜色通道进行模糊,而其他通道不做变化

在这里的 Alpha 是指可见区域(有点像 mask) 你可以在颜色显示那里 选择 alpha 通道 查看应用范围

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur1.png)

### Red/Green/Blue/Alpha Blurriness 红/绿蓝/Alpha 模糊度

如下图所示,非常简单,就不再深入举例了 下图分别是(red green blue alpha 模糊度均取 50 的情况)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur_red.png)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur_green.png)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur_blue.png)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur_alpha.png)

### Edge Behavior 边缘特性 -- 重复边缘像素

因为算法是从周围像素取平均值，但是图像边缘之外就没有像素了，一平均下来， **图像的边缘**
会半透明或者亮度变暗。勾选此选项可使边缘锐化，以防边缘变暗、变透明。

取消选择此选项则运行模糊算法，超出边缘的像素还是按零算。

### Blur Dimensions 模糊方向

- Horizontal and Vertical 水平和垂直
- Horizontal 水平
- Vertical 垂直

### 拓展：降噪

本效果可以用来降噪

比如本图片的绿色通道有明显噪点,可以稍微增加点绿色通道的通道模糊来搞定(损失的部分细节可以用锐化解决)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Channel_Blur2.png)

(本图片摘自琳达 AE 全特效解析-2013)
