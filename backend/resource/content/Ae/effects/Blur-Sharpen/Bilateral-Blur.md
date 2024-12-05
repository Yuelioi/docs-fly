---
title: 双向模糊
order: 1
---

# Bilateral Blur - 双向模糊

## 简述

通过 **对比度** 改变模糊的程度。可选择性地使图像变模糊，从而保留边缘和其他细节。

与低对比度区域相比，高对比度区域变模糊的程度低一些，这样就能在尽量保持图像轮廓和细节的状态下模糊。高对比度区域是指，像素值差别很大。

双向模糊效果和智能模糊效果之间的主要差异是：双向模糊效果仍然会使边缘和细节略微变模糊。在使用同样设置的情况下，与智能模糊效果实现的效果相比，双向模糊效果实现的效果更柔软、更梦幻。双向模糊效果实现的效果与
Adobe Photoshop 表面模糊滤镜的效果非常相似。

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/image00581.jpg)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Bilateral_Blur.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=118&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Bilateral_Blur.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Bilateral_Blur_cn.png)

| Bilateral Blur | 双向模糊 | ブラー(バイラテラル) |     |     |     |
| -------------- | -------- | -------------------- | --- | --- | --- |
| Radius         | 半径     | 半径                 |     |     |     |
| Threshold      | 阈值     | しきい値             |     |     |     |
| Colorize       | 彩色化   | 色付き               |     |     |     |

### Threshold 阈值

阈值越低，保留的细节越多。阈值越高，简化的效果越多，保留的细节越少。在边缘或其他突出细节区域，模糊半径会自动减少。阈值可确定哪些区域需要保留，以及哪些区域需要模糊化。

### Radius 半径

半径越大，意味确定每个像素值需要更多像素，因此会增加模糊程度。

### Colorize 彩色化

感觉叫“着色”更好点

默认勾选。对 R、G 和 B 颜色通道值的分别加权平均值。结果混合一起生成彩色图像。  
不勾选“彩色化”时，则对亮度进行模糊。结果生成灰白图像。
