---
title: CameraLayer
order: 2
---

# CameraLayer 摄像机对象

## CameraLayer object

app.project.item(index).layer(index)

描述：CameraLayer 对象表示合成中的摄像机图层。使用 LayerCollection.addCamera()创建摄像机。可以通过索引或名称在项目的图层集中对其进行访问。

父级关系：CameraLayer 是 Layer 对象的子类。Layer 的所有方法和属性均可用。

AE 属性

CameraLayer 没有定义其他属性，但具有与其他图层类型不同的 AE 属性。它具有以下属性和属性组：

| Marker        | 标记       |                 |                         |
| ------------- | ---------- | --------------- | ----------------------- |
| Transform     | 变换       |                 |                         |
|               |            | PointofInterest | 目标点                  |
|               |            | Position        | 位置                    |
|               |            | Scale           | 缩放                    |
|               |            | Orientation     | 方向                    |
|               |            | XRotation       | X 旋转                  |
|               |            | YRotation       | Y 旋转                  |
|               |            | Rotation        | 旋转                    |
|               |            | Opacity         | 不透明度                |
| CameraOptions | 摄像机选项 |                 |                         |
|               |            | 缩放            | Zoom                    |
|               |            | 景深            | Depth of Field          |
|               |            | 焦距            | Focus Distance          |
|               |            | 光圈            | Aperture                |
|               |            | 模概层次        | Blur level              |
|               |            | 光圈形状        | Iris Shape              |
|               |            | 光圈旋转        | Iris Rotation           |
|               |            | 光圈圆度        | Iris Roundness          |
|               |            | 光圈长宽比      | Iris Aspect Ratio       |
|               |            | 光圈衍射条纹    | Iris Diffraction Fringe |
|               |            | 高光增益        | Highlight Gain          |
|               |            | 高光阈值        | Highlight Threshold     |
|               |            | 高光饱和度      | Highlight Saturation    |
