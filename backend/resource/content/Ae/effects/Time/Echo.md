---
title: 残影
order: 3
---

# Echo - 残影

## 简述

残影效果组合图层中不同时间的帧。可以制作简单的 **视觉残影** 、 **拖尾** 或者更复杂的 **漩涡条纹** 效果。

仅当图层随时间发生变化时，才会生效。

此效果不会创建通过在合成内移动图层本身而创建的运动残影。非要创建这种残影，需先预合成图层，将所有属性移动到新合成中，然后在预合成图层上应用此效果。

适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20220102102312.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=40&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

<iframe src="https://player.bilibili.com/player.html?bvid=BV1bt411b7Xx&page=1&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Echo.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Echo_cn.png)

| Echo               | 残影           | エコー           |                    |              |            |
| ------------------ | -------------- | ---------------- | ------------------ | ------------ | ---------- |
| Echo Time(seconds) | 残影时间（秒） | エコー時間（秒） |                    |              |            |
| Number Of Echoes   | 残影数量       | エコーの数       |                    |              |            |
| Starting Intensity | 起始强度       | 開始強度         |                    |              |            |
| Decay              | 衰减           | 減衰             |                    |              |            |
| Echo Operator      | 残影运算符     | エコー演算子     | Add                | 相加         | 加算       |
|                    |                |                  | Maximum            | 最大值       | 最大       |
|                    |                |                  | Minimum            | 最小值       | 最小       |
|                    |                |                  | Screen             | 滤色         | スクリーン |
|                    |                |                  | Composite In Back  | 从后至前组合 |            |
|                    |                |                  | Composite In Front | 从前至后组合 |            |
|                    |                |                  | Blend              | 混合         | ブレンド   |

## 参数详解

### 残影时间（秒）

残影间隔，以秒为单位。负值为正常拖尾；正值为提前量。

### 残影数量

残影的数目。值为 2，加 **本体** 共有 3 个状态：

### 起始强度

残影序列中第一个图像的不透明度。

### 衰减

在残影序列，与前面的残影不透明度的比率。

例如衰减是 0.5，第 1 个残影是起始强度的一半；第二个残影是起始强度的四分之一。

### 残影运算符

用于合并残影的混合运算。

- 相加：通过将残影的像素值相加来 **合并** 残影。如果起始强度太高，此模式可能会快速过载并生成白色的线条。
- 最大值：通过获取所有残影的 **最大像素值** 来合并残影。
- 最小值：通过获取所有残影的 **最小像素值** 来合并残影。
- 滤色：通过在视觉上将残影夹在中间来模仿合并残影的效果。此模式与“相加”类似，但是它不会快速过载。（就是普通的滤色混合模式）
- 从后至前组合：使当前时间的图像置后，从而使每个残影在合成中依次置前。
- 从前至后组合：使当前时间的图像置前，从而使每个残影在合成中依次置后。
- 混合：使残影达到平衡。（与相加类似，但是有点不一样）

## 案例

漩涡条纹效果，需要残影间隔极地，比如 1/30

![](https://cdn.yuelili.com/20220102104502.png)
