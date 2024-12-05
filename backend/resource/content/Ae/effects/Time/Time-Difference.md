---
title: 时差
order: 6
---

# Time Difference - 时差

## 简述

时差效果计算两个图层之间的色差，有助于进行颜色校正；在将干净的背景图版与前景素材进行匹配时，可使用它来提取色差。

它还适合用于创建遮罩，然后利用遮罩来应用烟、火或残影效果。在应用时差效果以查找素材中的色差之后，您可以使用颜色校正效果来应用颜色校正。

此效果适用于 8-bpc 颜色。

将时差与粒子场结合使用可以仅从移动部分散发粒子。

## 效果展示

![](https://cdn.yuelili.com/20220102112242.gif)

## 教程

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Time_Difference.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Timewarp_cn.png)

| Time Difference     | 时差             | 時間差             |                       |               |          |
| ------------------- | ---------------- | ------------------ | --------------------- | ------------- | -------- | --- |
| Target              | 目标             | 保存先             |                       |               |          |
| Time Offset (sec)   | 时间偏移量（秒） | 時間オフセット(秒) |                       |               |          |
| Contrast            | 对比度           | コントラスト       |                       |               |          |
| Absolute Difference | 绝对差值         | 差の絶対値         |                       |               |          |
| Alpha Channel       | Alpha 通道       |                    | Oriana                | 原始          | RAW      |
|                     |                  |                    | Target                | 目标          | 保存先   |
|                     |                  |                    | Blend                 | 混合          | ブレンド |
|                     |                  |                    | Max                   | 最大值        | 最大     |
|                     |                  |                    | Full On               | 完全打开      |          |
|                     |                  |                    | Lightness of Result   | 结果亮度      |          |
|                     |                  |                    | Max of Result         | 结果最大值    |          |
|                     |                  |                    | Alpha Difference      | A             | pha 差值 |     |
|                     |                  |                    | Alpha Difference Only | 仅 Alpha 差值 |          |

## 参数详解

### 目标

指定要与效果图层进行比较的图层。比较图层不需要处于打开状态

### 时间偏移

比较图层中的相对时间，将在此时间对图层进行比较，以秒为单位。如果此控件设置为 0.00，将在当前时间进行比较。例如，要在比较图层中的 3
秒处对效果图层进行比较，请将时间偏移值更改为 3。在选择带下划线的位移值时，您可以采用帧/帧速率格式输入特定帧偏移值。After Effects
会自动计算该值。例如，输入 3/30 可在 30-fps 合成中向前偏移三个帧。计算出的值是 0.1，即总时间的 10%。

### 对比度

调整比较结果。此控件可能特别有助于微调颜色校正。

### 绝对差值

将比较结果显示为绝对值。比较图层中任何与效果图层相同的区域由黑色表示，任何差异值由比黑色更明亮的颜色表示。如果不选择此选项，没有差异的比较区域将表示为灰色。

### Alpha 通道

指定计算 Alpha 通道的方式。

- 原始：使用效果图层的 Alpha 通道。
- 目标：使用目标图层的 Alpha 通道。
- 混合：混合目标图层和效果图层的 Alpha 通道。
- 最大值：使用更加不透明的 Alpha 通道。
- 完全打开：将 Alpha 通道设置为完全不透明。
- 结果亮度：使用 RGB 差值的亮度作为 Alpha。
- 结果最大值：使用 RGB 差异的最高值作为 Alpha。
- Alpha 差值：采用与计算 RGB 差值相同的方式计算效果图层和目标图层的 Alpha 通道的差值。
- 仅 Alpha 差值：仅计算 Alpha 通道中的差异。RGB 被设置为白色。

## 案例:调整图层的进阶应用

通过在调整图层应用此效果，[CrayEditz](https://www.youtube.com/watch?v=XAN43CBonXk&ab_channel=CrayEditz)
创造了一种基于运动的人物光影。因为调整图层=空，所以任何运动都会出现“色差”

![](https://cdn.yuelili.com/20220102113010.png)
