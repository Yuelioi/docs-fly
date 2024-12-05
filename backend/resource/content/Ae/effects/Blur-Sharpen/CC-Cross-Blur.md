---
order: 4
title: CC交叉模糊
---

# CC Cross Blur - CC交叉模糊

## 简述

在绘图模式下合成垂直和水平模糊以[Blend]或[Screen]等绘制模式合成水平和垂直的模糊。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00562.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=14&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-CC_Cross_Blur.png)

| Radius X           | 半径 X       | 半径 X                   |     |     |     |
| ------------------ | ------------ | ------------------------ | --- | --- | --- |
| Radius Y           | 半径 Y       | 半径 Y                   |     |     |     |
| Transfer Mode      | 混合模式     | 描画モード               |     |     |     |
| Repeat Edge Pixels | 重复边缘像素 | エッジピクセルを繰り返す |     |     |     |
|                    |              |                          |     |     |     |

## 参数详解

### Radius X/ Y 半径 X/Y

控制模糊 X/Y 方向的半径

### Transfer Mode 混合模式

合成水平和垂直方向模糊的模式。

- Blend
- Add
- Screen
- Multiply
- LightenDarken

### Repeat Edge Pixels 重复边缘像素

因为算法是从周围像素取平均值，但是图像边缘之外没有像素了，一平均下来，边缘图像会半透明或者亮度变暗。

勾选此选项可使边缘锐化，以防边缘变暗、变透明。取消选择此选项则运行模糊算法，超出边缘的像素还是按零算。

![](https://mir.yuelili.com/user/AE/effects/ext/image00563.jpg)

### 示例 2

| 原图                                                                                                 | 只加一点半径 X(柔软的模糊效果 类似肥皂剧的感觉)                                                      | 同时加 X 和 Y(Cross 有交叉,十字型标记的意思) |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | -------------------------------------------- |
| ![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Cross_Blur1.png) | ![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Cross_Blur2.png) |

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Cross_Blur3.png)

###
