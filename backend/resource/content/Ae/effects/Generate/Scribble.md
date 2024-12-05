---
title: 涂写
order: 23
---

# Scribble - 涂写

## 简述

基于蒙版，生成类似涂写的效果

## 效果展示

![](https://cdn.yuelili.com/20211230230551.png) |
![](https://cdn.yuelili.com/20211230230605.png)  
---|---

## 教程

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-Scribble.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Generate-Scribble_cn.png)

| Scribble       | 涂写     | 落書き         |                        |                  |                                |                 |          |     |
| -------------- | -------- | -------------- | ---------------------- | ---------------- | ------------------------------ | --------------- | -------- | --- |
| Scribble       | 涂抹     | 落書き         | None                   | 无               | なし                           |                 |          |     |
|                |          |                | Sinale Mask            | 单个蒙版         | 単一マスク                     |                 |          |     |
|                |          |                | All Masks              | 所有蒙版         | すべてのマスク                 |                 |          |     |
|                |          |                | All Masks Using Modes  | 所有蒙版使用模式 | モードを使用するすべてのマスク |                 |          |
| Mask           | 蒙版     | マスク         |                        |                  |                                |                 |          |     |
| Fill Type      | 填充类型 | 塗りの種類     | Inside                 | 内部             | 内側                           |                 |          |     |
|                |          |                | Centered Edge          | 中心边缘         | エッジ中心                     |                 |          |     |
|                |          |                | Inside Dae             | 在边缘内         | エッジ内側                     |                 |          |     |
|                |          |                | Outside Edge           | 外面边缘         | エッジ外側                     |                 |          |     |
|                |          |                | Left Edge              | 左边             | 左エッジ                       |                 |          |     |
|                |          |                | Right Edge             | 右边             | 右エッジ                       |                 |          |     |
| Edge Option    | 边缘选项 |                |                        |                  |                                |                 |          |     |
|                |          |                | Edge Wdth              | 边缘宽度         | エッジの幅                     |                 |          |     |
|                |          |                | End Cap                | 末端端点         | 線端                           | Round           | 圆角     |     |
|                |          |                |                        |                  |                                | utt             | 平头     |     |
|                |          |                |                        |                  |                                | Projecting      | 投影     |     |
|                |          |                | Join                   | 连续             | 連続                           | Round           | 圆角     |     |
|                |          |                |                        |                  |                                | utt             | 平头     |     |
|                |          |                |                        |                  |                                | Projecting      | 投影     |     |
|                |          |                | Milar Limit            | 尖角限制         | 角の比率                       |                 |          |     |
|                |          |                | Start/End Apply To     | 开始/末端应用于  | 開始/終了の適用先              | Mask Path       | 蒙版路径 |     |
|                |          |                |                        |                  |                                | Scribble Result | 涂写结果 |     |
| Color          | 颜色     | カラー         |                        |                  |                                |                 |          |     |
| Opacity        | 不透明度 | 不透明度       |                        |                  |                                |                 |          |     |
| Angle          | 角度     | 角度           |                        |                  |                                |                 |          |     |
| Stroke Width   | 描边宽度 | 線幅           |                        |                  |                                |                 |          |     |
| Stroke Options | 描边选项 | ストローク     |                        |                  |                                |                 |          |     |
|                |          |                | Curviness              | 曲度             |                                |                 |          |     |
|                |          |                | Curviness Variation    | 曲度变化         |                                |                 |          |     |
|                |          |                | Spacing                | 间距             | 間隔                           |                 |          |     |
|                |          |                | Spacing Variation      | 间距变化         |                                |                 |          |     |
|                |          |                | Path Overlap           | 路径重叠         |                                |                 |          |     |
|                |          |                | Path Overlap Variation | 路径重叠变化     |                                |                 |          |     |
| Start          | 起始     | 開始           |                        |                  |                                |                 |          |     |
| End            | 结束     | 終了           |                        |                  |                                |                 |          |     |
|                |          |                | Fill Paths Sequentally | 顺序填充路径     |                                |                 |          |     |
| Wiggle Type    | 摆动类型 | ウィグルの種類 | Static                 | 静态             | 静かな                         |                 |          |     |
|                |          |                | Jumpy                  | 跳跃性           |                                |                 |          |     |
|                |          |                | Smooth                 | 平滑             | スムージング                   |                 |          |     |
| Wiggles/Second | 摇摆/秒  | ウィグル/秒    |                        |                  |                                |                 |          |     |
| Random Seed    | 随机植入 | ランダムシード |                        |                  |                                |                 |          |     |
| Composite      | 合成     | コンポジション | On Original Image      | 在原始图像上     | 元のイメージ                   |                 |          |     |
|                |          |                | On Transparent         | 在透明背景上     | 透明                           |                 |          |     |
|                |          |                | Reveal Original Image  | 显示原始图像     | 元のイメージを表示             |                 |          |     |

## 参数详解

很多参数与形状图层一致，不再赘述。

### 蒙版

根据蒙版选定涂写范围，当然还可以计算蒙版的叠加模式

### 填充类型

内部填充还是外部边缘，一共 6 个选项，任君挑选

### 边缘选项

跟形状图层的描边一样（见下）

![](https://cdn.yuelili.com/20211230231302.png)

### 颜色、不透明度、角度

这还用说明吗

### 描边宽度

### 描边选项

控制描边的各种参数。这里可以玩一玩。

- 曲度、曲度变化
- 间距、间距变化
- 路径重叠、路径重叠变化

![](https://cdn.yuelili.com/20211230233619.png)

### 起始与结束

与形状图层修剪路径一样

### 摆动类型

看图说话

| 静态                                            | 跳跃性 | 平滑 |
| ----------------------------------------------- | ------ | ---- |
| ![](https://cdn.yuelili.com/20211230233920.png) |
| ![](https://cdn.yuelili.com/20211230233948.png) |

![](https://cdn.yuelili.com/20211230234000.png)

### 摇摆/秒

抖动速度，参考 wiggle。0 则静止

### 随机植入

随机种子值，为状态的随机初始化值

### 合成

- 在原始图像上
- 在透明背景上
- 显示原始图像

## 案例
