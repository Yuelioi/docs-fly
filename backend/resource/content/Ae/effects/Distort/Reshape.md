---
order: 27
title: 改变形状
---

# Reshape - 改变形状

## 简述

在边界蒙版范围内，创建由“源蒙版”到“目标蒙版”的变形过渡。除了改变图像的整个形状外，还可以轻松地对其进行变换，例如仅放大图像的一部分。

### 关键词

效果：过渡、流体位移

类型：扭曲

控制基于：

控制范围：点控制、区域控制、蒙版控制

生效范围：不透明度、位置、扭曲

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=84&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Reshap.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Reshap_cn.png)

| Reshape | 改变形状 | リシェープ |                      |          |                    |                   |            |              |
| ------- | -------- | ---------- | -------------------- | -------- | ------------------ | ----------------- | ---------- | ------------ |
|         |          |            | Source Mask          | 源蒙版   | リシェープ元マスク |                   |            |              |
|         |          |            | Destination Mask     | 目标蒙版 | リシェープ先マスク |                   |            |              |
|         |          |            | Boundary Mask        | 边界蒙版 | 境界マスク         |                   |            |              |
|         |          |            | Percent              | 百分比   | パーセント         |                   |            |              |
|         |          |            | Elasticity           | 弹性     | 弾力性             | Stiff             | 生硬       |              |
|         |          |            |                      |          |                    | Less stiff        | 不那么生硬 |              |
|         |          |            |                      |          |                    | Below normal      | 标准以下   |              |
|         |          |            |                      |          |                    | Normal            | 正常       | 標準         |
|         |          |            |                      |          |                    | Absolutely Normal | 绝对正常   |              |
|         |          |            |                      |          |                    | Above Average     | 平均以上   |              |
|         |          |            |                      |          |                    | Loose             | 松散       |              |
|         |          |            |                      |          |                    | Liquid            | 液体       |              |
|         |          |            |                      |          |                    | Super Fluid       | 超级流体   |              |
|         |          |            | Correspondence Point | 对应点   | 対応ポイント       |                   |            |              |
|         |          |            | Interpolation Method | 计算密度 | 補間法             | Discrete          | 分离       | 不連続       |
|         |          |            |                      |          |                    | Linear            | 线性       | リニア       |
|         |          |            |                      |          |                    | Smooth            | 平滑       | スムージング |

## 参数详解

### 百分比

调整变形强度，值越大，源蒙版指定范围内会变形越大，使其更适合目标蒙版。

![](https://cdn.yuelili.com/20211225013424.png)

### 弹性

变形时的精度可以设置：

生硬：最粗糙，尤其是曲线部分会变成锯齿状。

超级液体：变形会更加平滑。（需要更长渲染时间）

### 计算密度

动画时，可以从三个选项中指定关键帧之间转换的平滑度

- Discontinuity：逐帧计算。结果最准确，但渲染时间更长。
- 线性（默认）：在两个或多个关键帧之间有规律地线性变换。会忽视关键帧缓动，为独立线性计算
- 平滑：在三个或更多关键帧之间的曲线中平滑变化。会忽视关键帧缓动，为独立平滑计算

## 案例：变换图像

![](https://cdn.yuelili.com/20211225013815.gif)
