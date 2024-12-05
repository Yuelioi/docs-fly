---
title: 闪光灯
order: 23
---

# Strobe Light - 闪光灯

## 简述

闪光灯效果可以定期或随机间隔在图层上执行算术运算，或使图层变透明。

例如，每 5 秒，图层会变为完全透明状态 1 秒，或者图层颜色按随机间隔反转。

此效果适用于 8-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20220102022404.gif)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=121&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Strobe_Light.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Stylize-Strobe_Light_cn.png)

| Strobe Light             | 闪光灯             | ストロボ                   |                         |                     |        |
| ------------------------ | ------------------ | -------------------------- | ----------------------- | ------------------- | ------ |
| Strobe Color             | 闪光颜色           | ストロボカラー             |                         |                     |        |
| Blend With Original      | 与原始图像混合     | 元の画像とブレンド         |                         |                     |        |
| Strobe Duration(secs)    | 闪光持续时间（秒） | ストロボデュレーション(秒) |                         |                     |
| Strobe Period (secs)     | 闪光间隔时间（秒） | ストロボ間隔(秒)           |                         |                     |        |
| Random Strobe Probablity | 随机闪光概率       | ランダムストロボの確率     |                         |                     |        |
| Strobe                   | 闪光               | 稲妻                       | Operates On Color Only  | 仅对颜色操作        |        |
|                          |                    |                            | Makes Layer Transparent | 使图层透明          |        |
| Strobe Operator          | 闪光运算符         | ストロボ演算子             | Copy                    | 复制                | コピー |
|                          |                    |                            | Add                     | 相加                | 加算   |
|                          |                    |                            | Sub...                  | ...等 12 种混合模式 |        |
| Random Seed              | 随机植入           | ランダムシード             |                         |                     |        |

## 参数详解

### 闪光颜色

闪光灯的颜色。

### 与原始图像混合

效果的影响程度，是效果图像与原始图像混合的结果。值越高，效果影响越小。 100%，则几乎没影响。

### 闪光持续时间（秒）

每个闪光持续的时间，以秒为单位。

### 闪光间隔时间（秒）

闪光循环间隔时间，以秒为单位。必须比持续时间大！

### 随机闪光概率

对所有给定帧应用闪光运算的概率。

### 闪光

- 使图层透明：可使图层变透明
- 仅对颜色操作：可使用“闪光运算符”指定的运算。

### 闪光运算

用于每个闪光的运算，然后混合模式。

## 案例

看教程吧，感觉用处不大
