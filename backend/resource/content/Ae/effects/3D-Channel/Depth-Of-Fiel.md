---
title: 场深度
order: 4
---

# Depth Of Field - 场深度

## 简述

可读取 3D 图像中的深度信息，并沿 z 轴,在任意位置对图像切片。例如，可以隐藏 3D 场景中的背景，也可以将对象置入 3D 场景中。

加上焦点模糊。从包含 Z 深度通道的 RLA、RPF 文件模拟照相机的景深，进行焦点的调整。可以通过参数设定聚焦范围、聚焦位置等。

注意：选择效果后，使用选择工具单击“合成”面板或“图层”面板中的相应对象，可以在“信息”面板中显示对象的深度。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-depth_of_field1.jpg)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Depth_Of_Field.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Depth_Of_Field_cn.png)

| Depth Of Field        | 场深度     | 被写界深度               |     |     |
| --------------------- | ---------- | ------------------------ | --- | --- |
| Focal Plane           | 焦平面     | フォーカルプレーン       |     |     |
| Maximun Radius        | 最大半径   |                          |     |     |
| Focal Plane Thickness | 焦平面厚度 | フォーカルプレーンの厚さ |     |     |
| Focal Bias            | 焦点偏移   | フォーカルバイアス       |     |     |

## 参数详解

### Focal Plane(焦平面)

焦点平面沿 z 轴到摄像机的距离。

### Maximum Radius(最大半径)

焦平面以外的对象会有模糊的程度。

### Focal Plane Thickness(焦平面厚度)

确定焦点对准焦平面任一侧的深度。

### Focal Bias(焦点偏移)

值越高，元素离开焦点的速度越快，且距焦平面的距离越远。  
例：调整参数，可以使机器人逐渐失焦

<![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Depth_Matte1.png)

### 示例 2

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-depth_of_field3.jpg)
