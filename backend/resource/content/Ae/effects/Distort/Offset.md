---
title: 偏移
order: 24
---

# Offset - 偏移

## 简述

偏移效果可在图层内平移图像。可以像电影胶片一样上下移动

用途之一是在图层中创建循环背景。在“最佳”品质时，使用子像素精度执行偏移操作。

此效果适用于 8-bpc 和 16-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20211212150544.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=2&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Offset.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Offset_cn.png)

| Offset | 偏移 | オフセット |                     |                |                    |
| ------ | ---- | ---------- | ------------------- | -------------- | ------------------ |
|        |      |            | Shift Center To     | 将中心转换为   | 中央をシフト       |
|        |      |            | Blend With Original | 与原始图像混合 | 元の画像とブレンド |

## 参数详解

将中心转换为：原始图像中心点的新位置。

![](https://cdn.yuelili.com/20211223020720.png)

与原始图像混合：效果图像的透明度。此值设置得越高，此效果对图层的影响越小。例如，此值设置为 100%，则没有效果；

## 案例

Lloyd Alvarez 在 [AE Enhancers
论坛](https://www.adobe.com/go/learn_ae_lloydoffsetprojector_cn)上提供了一个简单表达式，您可将其应用到“将中心转换为”属性，以模拟短时间的影片投影仪。

```javascript
fps = 24;
t = time * fps;
move = t * thisComp.height;
reset = Math.floor(move / thisComp.height);
moveReset = move - reset * thisComp.height;
[thisProperty[0], thisProperty[1] + moveReset];
```
