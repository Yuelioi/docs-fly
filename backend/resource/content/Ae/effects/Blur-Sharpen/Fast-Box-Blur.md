---
title: 快速方框模糊
order: 11
---

# Fast Box Blur - 快速方框模糊

## 简述

利用 GPU 快速处理模糊，与老式模糊一样，不过可以用 GPU 加速

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00590.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=26&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

### ![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Fast_Box_Blur.png)

### ![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Fast_Box_Blur_cn.png)

| Fast Box Blur   | 快速方框模糊 | 高速ボックスブラー |                         |            |                |
| --------------- | ------------ | ------------------ | ----------------------- | ---------- | -------------- |
| Blur Radius     | 模糊半径     | ブラーの半径       |                         |            |                |
| Iterations      | 迭代         | 繰り返し           |                         |            |                |
| Blur Dimensions | 模糊方向     | ブラーの方向       | Horizontal and Vertical | 水平和垂直 | 水平および垂直 |
|                 |              |                    | Horizontal              | 水平       | 水平           |
|                 |              |                    | Vertical                | 垂直       | 垂直           |

## 参数详解

### Blur Radius 模糊半径

模糊的半径。你也可以理解为模糊的程度

### Iterations 迭代

将模糊效果连续应用到图像的次数，模糊的细致程度。迭代数越大，颜色之间的过渡越平滑，模糊程度越高，但渲染时间也越长。默认值会产生适中的效果。

### Blur Dimensions 模糊方向 (字面意思 不赘述)

### Horizontal and Vertical 水平和垂直

- Horizontal 水平
- Vertical 垂直

在迭代 1 次时，Box
Blur 的简单模糊处理非常方便。例如，在“水平”或“垂直”模式下使用时，“方框模糊”比“方向模糊”更好地近似运动模糊。摄影运动模糊应具有"框型"的外观，而不应具有高斯风格的平滑度。

在迭代 4 次时，Box Blur 在质量上与 Fast Blur 相同。达到 4 或 5 更好，产生更柔和，更圆的模糊。当模糊小的，明亮的浮点物体时，这种额外的质量会很好。

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Fast_Box_Blur1.png)
