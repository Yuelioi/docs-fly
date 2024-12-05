---
title: CC矢量模糊
order: 7
---

# CC Vector Blur - CC矢量模糊

## 简述

以分层为基础进行淡化。使用配置在时间轴上的其他层，使图像变形或模糊。

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur0.png)

![](https://mir.yuelili.com/user/AE/effects/ext/image00568.jpg)

## 参数中英日对照

### ![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-CC_Vector_Blur.png)

| Type             |            |            | Natural          | 自然     |      |
| ---------------- | ---------- | ---------- | ---------------- | -------- | ---- |
|                  |            |            | Constant         | 恒定     |      |
|                  |            |            | Perpendicular    | 垂直     | 垂直 |
|                  |            |            | Direction Center | 方向中心 |      |
|                  |            |            | Direction Fading | 方向衰减 |      |
| Amount           | 数量       | 量         |                  |          |      |
| Angle Offset     | 角度偏移   |            |                  |          |      |
| Ridge Smoothness | 岭平整度   |            |                  |          |      |
| Vector Map       | 矢量贴图   |            |                  |          |      |
| Property         | 属性       | プロパティ |                  |          |      |
| Map Softness     | 贴图柔软度 |            |                  |          |      |

## 参数详解

### Type

**区域组**

- Natural 自然(默认) 模糊通过点 a 与点 b 之间的区域生成
- Constant 恒定 模糊通过点 a 与点 b 之间的区域 并循环
- Perpendicular 垂直 与 Natural 的区域垂直

**方向组**

- Direction Center 方向中心 模糊通过点 A 到点 B 以及点 B 到点 A 方向上 双向模糊
- Direction Fading 方向衰减 同上,只不过仅在点 A 到点 B 方向上模糊

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur10.png)

### Amount 数量

模糊的程度

### Angle Offset 角度偏移

模糊的方向进行旋转 旋转 360° 就等于没旋转哦

### Ridge Smoothness/Revolution

岭平整度(区域组)/演化(方向组)

平整度:模糊的整体增加柔滑,羽化效果

演化: 看图就懂(N 个演化就多了 N 个褶 当然用的是 Direction Fading 用 Direction
Center 的话就 N\*2 个褶,因为这个效果是两个方向同时模糊的)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur12.png)

### Vector Map 矢量贴图

- None 无 也就是以自身为贴图
- Layer 选择一个图层作为贴图
- Source 图层源
- Mask 蒙版
- Mask & Effects 蒙版&效果

公式: 渐变原图选择 source 模式 = 纯色图层+渐变效果 选择 Mask&Effects

### Property 属性

![](https://mir.yuelili.com/user/AE/effects/list0/Blur-Sharpen-CC_Vector_Blur02.png)

基于某属性作为模糊依据

### Map Softness 贴图柔度

增加该值时，获得更平滑的模糊贴图,羽化 AB 点的边缘

### 示例

| 原图 1:一个猫猫头                                                                                     | 原图 2:一个渐变贴图(此时 Vector Map 可以选 source)                                                    | 原图 3:一个带渐变的纯色层,并且加了矩形蒙版 |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| ![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur1.png) | ![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur2.png) |

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur3.png)

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur_mask2.png)
![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur_mask3.png)

### 小技巧

(老马 2020/3/22 21 点 15 分左右直播) :

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-CC_Vector_Blur15.png)
