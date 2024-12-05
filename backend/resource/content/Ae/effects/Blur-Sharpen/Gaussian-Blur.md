---
title: 高斯模糊
order: 12
---

# Gaussian Blur - 高斯模糊

## 简述

使相邻的颜色的边缘发生变化，使之整体模糊。模糊的方向可以设定为水平、垂直、全部。也可以制作简易的噪声去除和聚焦/退出的效果。

高斯模糊效果可使图像变模糊，柔化图像并消除杂色。图层的品质设置不会影响高斯模糊效果。

此效果使用“重复边缘像素”选项和 8-bpc、16-bpc 与 32-bpc 颜色。此效果使用 GPU 加速来实现更快的渲染。

高斯模糊效果可替换高斯模糊（旧版）效果，并且会产生与后者稍微不同的渲染效果。

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=26&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

| Gaussian Blur   | 高斯模糊 | ブラー (ガウス) |                         |            |                |
| --------------- | -------- | --------------- | ----------------------- | ---------- | -------------- |
| Blurriness      | 模糊度   | ブラー          |                         |            |                |
| Blur Dimensions | 模糊方向 | ブラーの方向    | Horizontal and Vertical | 水平和垂直 | 水平および垂直 |
|                 |          |                 | Horizontal              | 水平       | 水平           |
|                 |          |                 | Vertical                | 垂直       | 垂直           |

## 参数图片参考

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Gaussian_Blur.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Gaussian_Blur_cn.png)

## 效果展示

下图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00577.jpg)

## 参数详解

### Blurriness 模糊度

模糊的程度 -- 字面意思

### Blur Dimensions 模糊方向

- Horizontal and Vertical 水平和垂直
- Horizontal 水平
- Vertical 垂直

### Repeat Edge Pixels 重复边缘像素

因为算法是从周围像素取平均值，但是图像边缘之外就没有像素了，一平均下来，边缘图像会半透明或者亮度变暗。勾选此选项可使边缘锐化，以防边缘变暗、变透明。取消选择此选项则运行模糊算法，超出边缘的像素还是按零算。

由下图看来 高斯模糊与快速方框模糊 迭代设置为 3 效果比较类似

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Fast_Gaussian_Blur1.png)
