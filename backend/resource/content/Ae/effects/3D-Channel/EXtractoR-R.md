---
title: 提取器R
order: 5
---

# EXtractoR - 提取器R

## 简述

从 Openexr 文件提取通道制作蒙版  
可以针对每个 RGBA 通道显示, 作为 HDI 文件的一种 - Openexr 文件中, 包含的 Z 深度通道来创建蒙版。

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-extractor1.jpeg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=124&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-EXtractoR_cn.png)

| Black Point | 黑场 | ブラックポイント |     |     |
| ----------- | ---- | ---------------- | --- | --- |
| White Point | 白场 | ホワイトポイント |     |     |
| UnMult      | 去黑 | UnMult           |     |     |

## 参数详解

### Channel Info 通道信息

将在 3D CG 软件等输出 Openexr 文件时设定的 3D 通道的要素设定为每个 RGBA 频道显示哪个要素。单击属性名称后，将显示用于设置频道显示的面板。

PROCESS

### Black Point 黑场

设定 Channel Info 中显示的通道的最黑显示值

### White Point 白场

设定 Channel Info 中显示的通道的最白显示值

### Unmult 去黑

重建 α 通道的黑色阶

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-extractor2.jpeg)

### 其他

[PDF 文档](http://www.fnordware.com/ProEXR/ProEXR_Manual.pdf) 英文
