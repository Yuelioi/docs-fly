---
title: 智能模糊
order: 15
---

# Smart Blur - 智能模糊

## 简述

智能模糊效果可模糊低对比度区域，同时保留 **边缘** 和 **高对比度** 区域

两者都可以通过将相似的颜色融合在一起而不仅仅是模糊图像来工作

高对比度边缘仅将那些高对比度边缘变为白色，而将其余图像变为黑色

一句话解释原理:首先根据阈值寻找 **高低对比度线** ,然后对线内区域进行 **模糊**

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/image00588.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=118&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 参数中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Smart_Blur.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Smart_Blur_cn.png)

| Smart Blur | 智能模糊 | ブラー(詳細) |              |          |                      |
| ---------- | -------- | ------------ | ------------ | -------- | -------------------- |
| Radius     | 半径     | 半径         |              |          |                      |
| Threshold  | 阈值     | しきい値     |              |          |                      |
| Mode       | 模式     | モード       | Normal       | 正常     | 標準                 |
|            |          |              | Edge Only    | 仅限边缘 | エッジのみ           |
|            |          |              | Overlay Only | 叠加边缘 | エッジのオーバーレイ |

## 参数详解

### Radius 半径

进行模糊的半径

### Threshold 阈值

判断高对比度区域的阈值

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Smart_Blur3.png)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Smart_Blur2.png)

### Mode 模式

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Smart_Blur4.png)

**-Normal 正常**

对整个图像的低对比度区域进行模糊，同时保留边缘和高对比度区域。阈值变大，会导致高对比度区域减少，因此模糊量增大(见下图)

**-Edge Only 仅限边缘**

高对比区域线会变白. 其余图像变为黑色

**-Overlay Only 叠加边缘**

会执行"仅限边缘"的操作，但不是将其余图像变成黑色，而是使高对比度线周围进行突出显示
