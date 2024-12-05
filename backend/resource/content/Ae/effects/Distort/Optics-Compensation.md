---
title: 光学补偿
order: 25
---

# Optics Compensation - 光学补偿

## 简述

使用光学补偿效果可添加或移除摄像机镜头扭曲效果。比如：把图像贴在老式往里凹的屏幕上，不对，那我为啥不用贝塞尔变形？

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=54&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Optics_Compensation.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Optics_Compensation_cn.png)

| Optics Compensation | 光学补偿    | レンズ補正 |                                     |                      |                                |
| ------------------- | ----------- | ---------- | ----------------------------------- | -------------------- | ------------------------------ |
| Field Of View(FOV)  | 视场（FOV） |            |                                     |                      |                                |
|                     |             |            | Reverse Lens Distortion             | 反转镜头扭曲         | レンズディストーションを反転   |
| FOV Orientation     | FOV 方向    |            | Horizontal                          | 水平                 | 水平                           |
|                     |             |            | Vertical                            | 垂直                 | 垂直                           |
|                     |             |            | Diagonal                            | 对角                 | 対角                           |
| View Center         | 视图中心    | 視界の中心 |                                     |                      |                                |
|                     |             |            | Optimal Pixels(Invalidates Reversal | 最佳像素（反转无效） | 最適ピクセル(反転を無効にする) |
| Resize              | 调整大小    | サイズ変更 |                                     |                      |                                |
|                     |             |            | off                                 | 关闭                 | 閉じる                         |
|                     |             |            | Max 2X                              | 最大 2X              |                                |
|                     |             |            | Max 4X                              | 最大 4X              |                                |
|                     |             |            | Unlimited                           | 无限                 |                                |

## 参数详解

### 视场

扭曲素材的 _视场_ ( _FOV_ )。初始值为 0，没有变化，数值越大，扭曲程度越高。

### 反转镜头扭曲

如果选中，将反转扭曲方向。例如，要移除广角镜扭曲，请将“视场”设置为 40.0，并选择“反转镜头扭曲”。选择“反转镜头扭曲”将启用“调整大小”控件。

![](https://cdn.yuelili.com/20211225015838.png)

### 视线方向（FOV 方向）

按照该方向歪曲[视线]

- 垂直
- 水平
- 对角线

### 视图中心

指定视图的备用中心点。默认当然是图层中心

### 最佳像素

通过扭曲保持尽可能多的像素信息。选择此选项后，FOV 值不能再反转。

### 调整大小

扩展该效果的生效范围。要使用此控件，请先选择“反转镜头扭曲”，然后选择选项。“关闭”不能调整图层的大小。

- 最大 2X：最多可将图层的大小调整到原始宽度和高度的两倍。
- 最大 4X：四倍。
- 无限：无限制范围。此选项可能需要大量内存。

## 案例

场景变化就像被扭曲一样，因此图像会产生影响。

![](https://cdn.yuelili.com/20211225015928.gif)
