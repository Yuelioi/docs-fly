---
title: 范围扩散
order: 5
---

# Grow Bound - 范围扩散

## 简述

AE 的效果通常只作用于图层上，本效果可扩大效果在该图层的作用范围。

此效果对已启用“折叠变换”/“连续栅格化”的图层最为有用，因为这些图层使用与合成大小相同的缓冲区进行渲染。

例如，如果您将“投影”应用于部分超出合成帧的文本图层，则阴影会被剪切，因为只有位于合成帧中的文本部分才会投影。在“投影”效果防止剪切阴影之前应用“范围扩散”效果。您指定的像素数可增加图层缓冲区的高度和宽度。

## 效果展示

比如 CC 径向快速模糊

![](https://cdn.yuelili.com/20220102215615.png)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=29&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Utility-Grow_Bound.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Utility-Grow_Bound_cn.png)

## 参数详解

像素：扩散的效果作用域范围

## 案例
