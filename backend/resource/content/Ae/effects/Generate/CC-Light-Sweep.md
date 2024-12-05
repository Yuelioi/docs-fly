---
title: CC扫光
order: 9
---

# CC Light Sweep - CC扫光

## 简述

使用效果灯光，给图层增加一个扫光（会考虑 Alpha）

## 效果展示

![](https://cdn.yuelili.com/20211230110108.png) |
![](https://cdn.yuelili.com/20211230110537.png) |
![](https://cdn.yuelili.com/20211230112216.png)  
|---|---|---|

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=77&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-CC_Light_Sweep.png)

## 参数详解

### 中心与方向

扫光的中心位置和方向，这样可以确定一个扫光路线

### 形状

线性：光会线性过度（极亮到周围不亮）

柔和：光纤比较均匀、柔和

尖锐：默认，中间光柱比较突出

| 线性                                            | 柔和 | 尖锐 |
| ----------------------------------------------- | ---- | ---- |
| ![](https://cdn.yuelili.com/20211230110926.png) |
| ![](https://cdn.yuelili.com/20211230110949.png) |

![](https://cdn.yuelili.com/20211230111007.png)

### 宽度

灯光生效范围

| 10                                              | 100 | 500 |
| ----------------------------------------------- | --- | --- |
| ![](https://cdn.yuelili.com/20211230111253.png) |
| ![](https://cdn.yuelili.com/20211230111334.png) |

![](https://cdn.yuelili.com/20211230111351.png)

### 扫光强度

最大 500，0 则没光

| 50                                              | 100 | 500 |
| ----------------------------------------------- | --- | --- |
| ![](https://cdn.yuelili.com/20211230111334.png) |
| ![](https://cdn.yuelili.com/20211230111608.png) |

![](https://cdn.yuelili.com/20211230111620.png)

### 边缘强度与厚度

这个一般要求图层带 Alpha，比如 logo、文字。

![](https://cdn.yuelili.com/20211230111828.png)

### 灯光颜色

字面意思

### 灯光接受

其实就是灯光效果的混合模式

Add：相加，默认的

Composite：不混合

CutOut：只显示灯光

## 案例
