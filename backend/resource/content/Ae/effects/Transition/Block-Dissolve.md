---
title: 块溶解
order: 1
---

# Block Dissolve - 块溶解

## 简述

块溶解效果使图层消失在随机块中。可以像素为单位单独设置块的宽度和高度。使用“草图”品质时，块采用像素精度放置并具有清晰的边缘；使用“最佳”品质时，块使用子像素精度放置并具有模糊的边缘。

此效果适用于 8-bpc 和 16-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20220103010643.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=76&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Block_Dissolve.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Block_Dissolve_cn.png)

| Block Dissolve           | 块溶解               | ブロックディゾルブ |     |     |     |
| ------------------------ | -------------------- | ------------------ | --- | --- | --- |
| Transition Completion    | 过渡完成             | 変換終了           |     |     |     |
| Block Width              | 块宽度               | ブロック幅         |     |     |     |
| Block Height             | 块高度               | ブロック高さ       |     |     |     |
| Feather                  | 羽化                 | 境界のぼかし       |     |     |     |
| Soft Edges(Best Quality) | 柔化边缘（最高品质） |                    |     |     |     |

## 参数详解

### 过渡完成

完成程度。

### 块宽度/高度

每个“块”的长宽。

| 1X1                                             | 10X10 | 50X50 |
| ----------------------------------------------- | ----- | ----- |
| ![](https://cdn.yuelili.com/20220103011117.png) |
| ![](https://cdn.yuelili.com/20220103011051.png) |

![](https://cdn.yuelili.com/20220103011135.png)

### 羽化

羽化块边缘。如果羽化值极大，那么跟直接更改不透明 **差不多** 了。

### 柔化边缘（最高品质）

| 关闭（呈现块状）                                | 开启（默认） |
| ----------------------------------------------- | ------------ |
| ![](https://cdn.yuelili.com/20220103010643.png) |

![](https://cdn.yuelili.com/20220103010728.png)

## 案例
