---
title: 变换
order: 32
---

# Transform - 变换

## 简述

为图层增加额外的变换属性。因为它是效果，所以只会影响图层内容，而不会像图层的变换一样，更改图层本身。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=11&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Transform.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Transform_cn.png)

## 参数详解

### 统一缩放

可以把 X，Y 分开

### 倾斜

字面意思

![](https://cdn.yuelili.com/20211224154336.png)

### 快门角度

如果选中，使用合成的快门角度。不选择则使用自己设置的。数值越高，运动模糊越强。（注意，自定义运动模糊只对本效果的参数生效）

### 采样

在绘制图层时如何使用插值方法。  
[ **双三次** **方法** ]产生的描绘比[ **双线性** ]更平滑，但渲染更久。

## 案例

### 其他

可以独立于合成的快门角度，为 **此图层** 的运动模糊指定快门角度。（必须为图层和合成启用运动模糊功能，此方法才有意义。）
