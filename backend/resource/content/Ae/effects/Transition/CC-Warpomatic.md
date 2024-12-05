---
order: 11
title: CC自动翘曲过渡
---

# CC Warpomatic - CC自动翘曲过渡

## 简述

通过计算两张图片的差异，进行扭曲。

## 效果展示

![](https://cdn.yuelili.com/20220103195249.gif)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=61&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-CC_Warpomatic.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-CC_Warpomatic_cn.png)

| CC Warpomatic   | CC 自动翘曲过渡 | CC Warpomatic |                        |            |                  |
| --------------- | --------------- | ------------- | ---------------------- | ---------- | ---------------- |
| Completion      | 完成            | 完了          |                        |            |                  |
| Layer to Reveal | 显示图层        |               |                        |            |                  |
| Reactor         | 反应器          |               | Brightness             | 亮度       | 輝度             |
|                 |                 |               | Contrast Differences   | 对比度差异 |                  |
|                 |                 |               | Brightness Differences | 亮度差异   |                  |
|                 |                 |               | Local Differences      | 局部差异   |                  |
| Smoothness      | 平滑            | スムージング  |                        |            |                  |
| Warp Amount     | 弯曲量          |               |                        |            |                  |
| Warp Direction  | 变形方向        |               | Joint                  | 联合       |                  |
|                 |                 |               | Opposing               | 相反       |                  |
|                 |                 |               | Twisting               | 扭曲       | ディストーション |
| Blend Span      | 混合跨度        |               |                        |            |                  |

## 参数详解

### 反应器

基于 XX 差异进行扭曲

### 平滑弯曲量

扭曲的平滑度。

| 5                                               | 20  | 250（最大） |
| ----------------------------------------------- | --- | ----------- |
| ![](https://cdn.yuelili.com/20220103195732.png) |
| ![](https://cdn.yuelili.com/20220103195801.png) |

![](https://cdn.yuelili.com/20220103195814.png)

### 变形方向

变形的方向。

| Joint（联合）                                   | Opposing（相反） | Twisting（扭曲） |
| ----------------------------------------------- | ---------------- | ---------------- |
| ![](https://cdn.yuelili.com/20220103200054.png) |
| ![](https://cdn.yuelili.com/20220103200105.png) |

![](https://cdn.yuelili.com/20220103200114.png)

### 混合跨度

两图片过渡交融的持续时间，越短过渡越突兀，最小值为 1 。越大中间过渡时间越长

中间过渡是指：两图片都有点影响

![](https://cdn.yuelili.com/20220103200528.png)

## 案例
