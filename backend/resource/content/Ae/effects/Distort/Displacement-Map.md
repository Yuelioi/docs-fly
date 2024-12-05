---
title: 置换图
order: 19
---

# Displacement Map - 置换图

## 简述

具体方法是根据“置换图层”像素的颜色值，水平和垂直置换像素。

比如设置为颜色。那么生效程度由置换图 **每个像素** 的颜色值决定。颜色值的范围是 0 - 255。每个值均转换为 -1 - 1 的缩放范围。

颜色值为 0，可生成最大的负置换（-1 \* 最大置换数量）。颜色值为 255，可生成最大的正置换。颜色值为 128，不生成任何置换。

比如置换数量是 100，A 像素颜色是 180，那么 A 像素会被移动到 （194 - 128）/128 \* 100 = 50（正向移动 50 px）

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

### 关键词

效果：过渡、流体、涂抹、位移

类型：噪波、扭曲

控制基于：颜色、不透明度

控制范围：像素级控制、区域控制

生效范围：不透明度、位置

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=82&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Displacement_Map.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Distort-Displacement_Map_cn.png)

## 参数详解

### 置换图层

作为置换源的图层

[ Source ] 只用参考图层  
[ Mask ] 用参考图层蒙版  
[ Effects and Masks ] 用参考图层蒙版与效果

![](https://cdn.yuelili.com/20211224155535.png)

### 用于水平/垂直替换

确定参考层中的哪些信息用于扭曲

### 最大水平/垂直更换

确定扭曲的程度

![](https://cdn.yuelili.com/20211224155718.png)

### 置换贴图特性

确定当原始图层的大小和引用图层的大小不同时如何插入图层

### 像素回绕 **Wrap Pixels**

检查以绘制到原始图层的大小。默认不检测

**扩展输出**

效果范围是否超出原始图层。默认扩展

![](https://cdn.yuelili.com/20211224155903.png)

### 友情效果

梯度渐变：用于制作规整的置换图层

分型杂色：疯狂的置换图层

马赛克：制作块状置换图

色调：控制位移方向（比如把白色映射到 128 灰，这样就不会向上位移了） 14:15~

### 技巧：任意角度置换

用两个变换夹住本效果：视频 16:30~
