---
title: CC图像擦除
order: 4
---

# CC Image Wipe - CC图像擦除

## 简述

基于置换图的 XX 信息，进行过渡。默认是由暗部区 → 高亮区。也可以使用 RGBA 等进行判断

## 效果展示

![](https://cdn.yuelili.com/20220103200858.gif)

![](https://cdn.yuelili.com/20220103193637.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=61&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-CC_Image_Wipe.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-CC_Image_Wipe_cn.png)

| CC Image Wipe   | CC 图像擦除 | CC Image Wipe  |                  |          |                      |            |        |          |
| --------------- | ----------- | -------------- | ---------------- | -------- | -------------------- | ---------- | ------ | -------- |
| Completion      | 完成        | 完了           |                  |          |                      |            |        |          |
| Border Softness | 边界柔和度  |                |                  |          |                      |            |        |          |
| Auto Softness   | 自动柔和度  |                |                  |          |                      |            |        |          |
| Gradient        | 渐变        | グラデーション |                  |          |                      |            |        |          |
|                 |             |                | Layer            | 图层     | レイヤー             |            |        |          |
|                 |             |                | Property         | 属性     | プロパティ           | Red        | 红     | 赤       |
|                 |             |                |                  |          |                      | Green      | 绿     | 緑       |
|                 |             |                |                  |          |                      | Blue       | 蓝     | 青       |
|                 |             |                |                  |          |                      | Alpha      | Alpha  | アルファ |
|                 |             |                |                  |          |                      | Luminance  | 亮度   | 輝度     |
|                 |             |                |                  |          |                      | Lightness  | 明亮   | 明るい   |
|                 |             |                |                  |          |                      | Hue        | 色相   | 色相     |
|                 |             |                |                  |          |                      | Saturation | 饱和度 | 彩度     |
|                 |             |                | Blur             | 模糊     | ブラー               |            |        |          |
|                 |             |                | Inverse Gradient | 反转渐变 | グラデーションを反転 |            |        |

## 参数详解

### 边界柔和度

越大越柔和，但是过大会。。。

| 0                                               | 20  | 100 |
| ----------------------------------------------- | --- | --- |
| ![](https://cdn.yuelili.com/20220103201020.png) |
| ![](https://cdn.yuelili.com/20220103201051.png) |

![](https://cdn.yuelili.com/20220103201105.png)

### 自动柔和度

试了，不知道为啥没效果

### 渐变

选择要进行过渡的图层以及判断标准

## 案例
