---
title: 深度遮罩
order: 3
---

# Depth Matte - 深度遮罩

## 简述

读取 3D 图像中的深度信息，并沿 z 轴在任意位置对图像切片。例如可以移除 3D 场景中的背景，也可以将对象插入 3D 场景中。

如何确定对象的深度：在选择效果后，使用选择工具单击“合成”面板或“图层”面板中的相应对象。

常用于根据原始图像的深度调整不透明度的材料，例如雾和烟。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-depth_matte1.jpg)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Depth_Matte.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Depth_Matte_cn.png)

| Depth Matte | 深度遮罩 | デプスマット |     |     |
| ----------- | -------- | ------------ | --- | --- |
| Depth       | 深度     | 深度         |     |     |
| Feather     | 羽化     | 境界のぼかし |     |     |
| Invert      | 反转     | 反転         |     |     |

## 参数详解

示例 2：原始就是一个黑疙瘩机器人，不贴图了

### 深度 Depth

对图像切片所在位置的 z 轴值。小于“深度”值的所有内容都将进行遮罩处理。

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Depth_Matte-1.png)

### 羽化 Feather

沿遮罩边缘的羽化量。

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Depth_Matte-2.png)

### 反转 Invert

选择此选项，改为对深度大于“深度”值的内容进行遮罩处理。（看外围的枪被羽化了）

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Depth_Matte-3.png)

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-depth_matte2.jpg)
