---
title: ID遮罩
order: 7
---

# Id Matte - ID遮罩

## 简述

3D 程序使用唯一的对象 ID
标记场景中的元素。根据 RLA、RPF 文件的 ID 制作遮罩可以单独显示使用 3DCG 软件渲染的 RLA、RPF 文件中包含的对象 ID 的通道，并用于蒙版等。

ID 遮罩效果使用此信息创建遮罩，以此排除所需元素之外的内容。

对图像内的各个对象进行色彩调整和辉光等时方便的效果

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Id_Matte.png)

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Id_Matte_cn.png)

| Id Matte     | ID 遮罩  | ID マット      |             |         |                 |
| ------------ | -------- | -------------- | ----------- | ------- | --------------- |
| Aux Channel  | 辅助通道 | 補助チャンネル |             |         |                 |
| ID Selection | ID 选择  | ID 選択        | Object ID   | 对象 ID | オブジェクト ID |
|              |          |                | Meterial ID | 材质 ID | マテリアル ID   |
| Feather      | 羽化     | 境界のぼかし   |             |         |                 |
| Invert       | 反转     | 反転           |             |         |                 |
| Use Coverage | 使用范围 | 使用範囲       |             |         |                 |

## 效果展示 1

原始图像（左上）、选择左边的异形 ID 并应用 ID 遮罩效果（左下）、合成图（右下）

![ID 遮罩效果：原始图像、使用较近的异形作为 ID 选择并应用 ID
遮罩效果的图像、在新背景上合成的图像](https://mir.yuelili.com/user/source/ef_05.png)

## 效果展示 2

选择球 ID ，并应用效果

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-id_matte1.jpeg)

## 参数详解

### 辅助通道

根据对象 ID 或材质 ID 隔离元素。

### ID 选择

对象的 ID 值。指定 3D CG 软件设定的各个 ID，只显示该 ID 的对象或素材。

**注意**

确定对象的 ID，可以在选择效果后，从“效果控件”面板的“辅助通道”菜单中选择 ID
类型，然后使用选择工具在“合成”面板或“图层”面板中单击此对象。如果选择“对象 ID”作为“辅助通道”，则“ID 选择”会自动更新为单击对象的 ID。

### 羽化

沿遮罩边缘的羽化量。模糊显示对象的轮廓（蒙版轮廓）。

### 反转

反转选择。选择此选项则变成对选择的对象进行遮罩处理。

### 使用范围

平滑创建的蒙版边缘，沿遮罩边缘像素移除对象后存储的颜色，以此创建更纯净的遮罩。仅当 3D 图像包含存储对象后颜色的相关信息的覆盖范围通道时，此选项才起作用。

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-id_matte2.jpg)
