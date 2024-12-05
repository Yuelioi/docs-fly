---
title: CC胶枪
order: 6
---

# CC Glue Gun - CC胶枪

## 简述

本效果根据图层模拟胶枪射胶水

## 效果展示

![](https://cdn.yuelili.com/20220102235922.png)

| 彩带一枚                                        | 加个胶枪 | 原始图 VS 加了效果 |
| ----------------------------------------------- | -------- | ------------------ |
| ![](https://cdn.yuelili.com/20211228001819.png) | ![](https://cdn.yuelili.com/20211228002239.png) |![](https://cdn.yuelili.com/20211228002228.png)|

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=129&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-CC_Glue_Gun.png)

| CC Glue Gun    | CC 胶枪      | CC Glue Gun |                 |          |                  |               |          |                |
| -------------- | ------------ | ----------- | --------------- | -------- | ---------------- | ------------- | -------- | -------------- |
| Brush Position | 笔刷位置     |             |                 |          |                  |               |          |                |
| Stroke width   | 描边宽度     | 線幅        |                 |          |                  |               |          |                |
| Density        | 密度         | 濃度        |                 |          |                  |               |          |                |
| TimeSpan (sec) | 时间跨度（秒 |             |                 |          |                  |               |          |                |
| Reflection     | 反射         | 反射        |                 |          |                  |               |          |                |
| Strength       | 强度         | 密度        |                 |          |                  |               |          |                |
| Style          | 风格         |             |                 |          |                  |               |          |                |
|                |              |             | Paint style     |          |                  | Plain         | 直接     |                |
|                |              |             |                 |          |                  | Wobbly        | 摇摆     |                |
|                |              |             | Wobble width    | 摇摆宽度 |                  |               |          |                |
|                |              |             | Wobble Height   | 摇摆高度 |                  |               |          |                |
|                |              |             | Wobble Speed    | 摇摆速度 |                  |               |          |                |
| Light          | 灯光         | 照明        |                 |          |                  |               |          |                |
|                |              |             | Using           | 使用     |                  | Effect Light  | 效果灯光 |                |
|                |              |             |                 |          |                  | AE Lights     | AE 灯光  |                |
|                |              |             | Light Intensity | 灯光强度 | ライトの強さ     |               |          |                |
|                |              |             | Light Color     | 灯光颜色 | ライトのカラー   |               |          |                |
|                |              |             | Light Type      | 灯光类型 | ライトの種類     | Distant Light | 平行光   |                |
|                |              |             |                 |          |                  | Point Light   | 点光源   | ポイントソース |
|                |              |             | Light Height    | 灯光高度 | ライトの高さ     |               |          |                |
|                |              |             | Lisht posrtion  | 灯元位置 |                  |               |          |                |
|                |              |             | LIght Direction | 灯光方向 | ライトの方向     |               |          |                |
| Fading         | 阴影         | シャドウ    |                 |          |                  |               |          |                |
|                |              |             | Ambient         | 环境     | アンビエント     |               |          |                |
|                |              |             | Diffuse         | 扩散     | ディフュージョン |               |          |                |
|                |              |             | Specular        | 反射     | 反射             |               |          |                |
|                |              |             | roughness       | 粗糙度   |                  |               |          |                |
|                |              |             | Metel           | 金属     |                  |               |          |                |

## 参数详解

### 笔刷位置

胶枪移动轨迹，如果不 K 帧，那就是一个点

### 描边宽度

胶枪口的宽度，或者说喷出来胶水的宽度

### 密度

胶水浓度

| 密度 1                                          | 密度 5 | 密度 50 |
| ----------------------------------------------- | ------ | ------- |
| ![](https://cdn.yuelili.com/20211228003405.png) | ![](https://cdn.yuelili.com/20211228003433.png) |![](https://cdn.yuelili.com/20211228003455.png)|

### 时间跨度

胶水停留的时间。时间一到，会按喷射轨迹消失

### 反射

| 0 反射                                          | 200 反射 |
| ----------------------------------------------- | -------- |
| ![](https://cdn.yuelili.com/20211228010305.png) |![](https://cdn.yuelili.com/20211228010342.png)|

### 强度

胶水每个间隔点的大小。当然，宽度过大，也会覆盖掉一部分“强度”

![](https://cdn.yuelili.com/20211228004238.png)

### 风格

| 直接                                            | 摇摆（手抖手抖） |
| ----------------------------------------------- | ---------------- |
| ![](https://cdn.yuelili.com/20211228004548.png) |![](https://cdn.yuelili.com/20211228004559.png)|

**摇摆高度与宽度**

![](https://cdn.yuelili.com/20211228005557.png)

### 灯光与阴影

需要把阴影的环境拉到 100 以下

**灯光颜色**

以下示例在环境值 0 的情况下测试

| 白光所有都可以照得到                            | 红光只能照带红色区域的（绿色就没掺红，所有变黑了） | 乱七八糟颜色光，啥都照一点 |
| ----------------------------------------------- | -------------------------------------------------- | -------------------------- |
| ![](https://cdn.yuelili.com/20211228011502.png) | ![](https://cdn.yuelili.com/20211228011528.png) |![](https://cdn.yuelili.com/20211228011549.png)|

**阴影**

环境

决定灯光阴影生效还是原始图像生效。

| 0：纯灯光阴影生效                               | 50：各占一半 | 100：纯原始图像 |
| ----------------------------------------------- | ------------ | --------------- |
| ![](https://cdn.yuelili.com/20211228011827.png) | ![](https://cdn.yuelili.com/20211228011853.png) |![](https://cdn.yuelili.com/20211228011923.png)|

## 案例
