---
title: 色调分离时间
order: 5
---

# Posterize Time - 色调分离时间

## 简述

色调分离时间效果会将图层锁定到某特定帧速率。它可以作为特殊效果独立使用，但它还有其他更微妙的用途。

例如，每秒 60 场的视频素材可以锁定到每秒 24 帧（然后场将以每秒 60
场的速度渲染）以呈现电影效果。此外，嵌套的合成也可以锁定到给定的帧速率。此效果在硬件设备中有时称为闪光。

为帧速率滑块的值设置动画可能会产生不可预知的结果。鉴于此，可用于帧速率的唯一插值法是“定格”。

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

**注意**

能够实现与向图层应用色调分离时间效果差不多相同结果的另一种方法是：预合成图层，在“合成设置”对话框中更改预合成的帧速率，然后在“高级”选项卡上为预合成设置“在嵌套时或在渲染队列中，保留帧速率”。此方法导致预合成保留自己的帧速率，而不是从包含它的合成那里继承那些设置。

## 效果展示

![](https://cdn.yuelili.com/20220102110827.gif)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=52&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Posterize_Time.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Time-Posterize_Time_cn.png)

## 参数详解

### 帧速率

控制当前图层的帧速率。

## 案例
