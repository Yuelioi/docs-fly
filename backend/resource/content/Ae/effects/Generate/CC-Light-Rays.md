---
title: CC光线放射
order: 8
---

# CC Light Rays - CC光线放射

## 简述

给定点光源，基于调整光的强度以及影响范围（半径）内的 **亮像素** ，向周围投射光。并以一种光学补偿的方式扭曲光线。黑色区域不会投射。

## 效果展示

![](https://cdn.yuelili.com/20211229212805.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=116&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-CC_Light_Rays.png)

| CC Light Rays     | CC 光线放射 |          |         |      |            |
| ----------------- | ----------- | -------- | ------- | ---- | ---------- |
| Intensity         | 强度        | 密度     |         |      |            |
| Center            | 中心        | 中心     |         |      |            |
| Radius            | 半径        | 半径     |         |      |            |
| Warp Softness     | 径向柔和度  |          |         |      |            |
| Shape             | 形状        | シェイプ | Round   | 圆形 | 円         |
|                   |             |          | Square  | 方形 |            |
| Direction         | 方向        | 方向     |         |      |            |
| Color from source | 来源的颜色  |          |         |      |            |
| Allow Brightening | 允许变亮    |          |         |      |            |
| Color             | 颜色        | カラー   |         |      |            |
| Trensfer mode     | 混合模式    |          | None    | 无   | なし       |
|                   |             |          | Add     | 相加 | 加算       |
|                   |             |          | Lighten | 变亮 | 明るく     |
|                   |             |          | Screen  | 屏幕 | スクリーン |

## 参数详解

原图

![](https://cdn.yuelili.com/20211229213806.png)

### 强度

光线强度。为 0 则效果不生效

| 100                                             | 100 | 1000 |
| ----------------------------------------------- | --- | ---- |
| ![](https://cdn.yuelili.com/20211229213908.png) |
| ![](https://cdn.yuelili.com/20211229213933.png) |

![](https://cdn.yuelili.com/20211229213849.png)

### 中心

灯光的中心位置

### 半径

灯光的半径

| 50                                              | 500 | 5000（超越次元了都） |
| ----------------------------------------------- | --- | -------------------- |
| ![](https://cdn.yuelili.com/20211229214453.png) |
| ![](https://cdn.yuelili.com/20211229214339.png) |

![](https://cdn.yuelili.com/20211229214225.png)

### 柔和度

光线的柔和度，同时还会降低灯源点亮度

| 0（完全不柔和，光线直来直去）光条纹             | 50  | 400（太柔了，直接可以投个虚影） |
| ----------------------------------------------- | --- | ------------------------------- |
| ![](https://cdn.yuelili.com/20211229215046.png) |
| ![](https://cdn.yuelili.com/20211229215321.png) |

![](https://cdn.yuelili.com/20211229215918.png)

### 灯光颜色

不用图层的颜色作为光了，自己设置

![](https://cdn.yuelili.com/20211229220159.png)

### 允许变亮

是否让“灯源”显示

| 允许                                            | 不允许 |
| ----------------------------------------------- | ------ |
| ![](https://cdn.yuelili.com/20211229220513.png) |

![](https://cdn.yuelili.com/20211229220530.png)

### 混合模式

这个效果的混合模式

## 案例
