---
title: 复合模糊
order: 9
---

# Compound Blur - 复合模糊

## 简述

可根据控件图层（也称为模糊图层或模糊图）的明亮度值使效果图层中的像素变模糊。

淡化使相邻的颜色的边缘发生变化，使之整体模糊。模糊的方向可以设定为水平、垂直、四方。也可以制作简易的噪声去除和聚焦/退出的效果。

### 原理

复合模糊首先要选择一个模糊图层作为模糊参考(不选=选自身)

黑色区域完全不模糊

白色区域完全模糊

中间则是模糊过渡

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/image00582.jpg)

## 参数中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Compound_Blur.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Blur-Sharpen-Compound_Blur_cn.png)

| Compound Blur      | 复合模糊         | ブラー(合成)           |     |     |     |
| ------------------ | ---------------- | ---------------------- | --- | --- | --- |
| Blur Layer         | 模糊图层         | ブラーレイヤー         |     |     |     |
| Maximum Blur       | 最大模糊         | 最大ブラー             |     |     |     |
| Stretch Map to Fit | 伸缩对应图以适合 | マップをフィットさせる |     |     |     |
| Invert Blur        | 反转模糊         | ブラーを反転           |     |     |     |

## 参数详解

### Blur Layer 模糊图层

选择参考的模糊控制图层(一般可以用个渐变图层)

### Maximum Blur 最大模糊

最大模糊程度，以像素为单位。

### Stretch Map to Fit 伸缩对应图以适合

相当于给模糊控制图层加个效果(图层-变换-适应合成)；不然，模糊控制图层的影响效果居中作用于原图层上

### Invert Blur 反转模糊

黑变白 白变黑 模糊变不模糊 不模糊变模糊(字面意思 不赘述)

## 案例：文字经过烟雾变模糊

![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Compound_Blur0.png)
![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Compound_Blur1.png)
![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Compound_Blur2.png)
![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Compound_Blur3.png)
![](https://mir.yuelili.com/user/AE/effects/list/Blur-Sharpen-Compound_Blur4.png)
