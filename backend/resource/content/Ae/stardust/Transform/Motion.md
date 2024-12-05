---
title: 运动 Motion
order: 6
---

# 运动 Motion

## 介绍

向粒子添加额外的运动，使用此节点，可以按照预设的运动（例如，圆周运动/灯光路径）来移动粒子，使粒子跟随并定向到该路径。

## Motion Type 运动类型

- 灯光路径：新建灯光，k 个路径，粒子会在灯光路径上移动
- 圆环：在圆环上运动
- 正弦：在正弦曲线上运动
- 反弹：粒子之间有弹力，且给个初始的弹力
- 朝向：所有粒子会朝向灯光，可能会跟其他粒子属性有冲突。

![](http://cdn.yuelili.com/202020160044-3.png)![](http://cdn.yuelili.com/202020140242-P.png)![](http://cdn.yuelili.com/202020140246-H.png)

![](http://cdn.yuelili.com/202020160045-4.png)![](http://cdn.yuelili.com/202020140242-K.png)![](http://cdn.yuelili.com/202020140247-Z.png)

![](http://cdn.yuelili.com/202020160046-H.png)![](http://cdn.yuelili.com/202020140243-F.png)![](http://cdn.yuelili.com/202020140245-X.png)

![](http://cdn.yuelili.com/202020160046-Q.png)![](http://cdn.yuelili.com/202020140243-A.png)

[![](http://cdn.yuelili.com/202020160047-w.png)![](http://cdn.yuelili.com/202020140254-A.png)](http://cdn.yuelili.com/202020140254-A.png)![](http://cdn.yuelili.com/202020140253-V.png)

## 运动类型

### Light Path 灯光路径

Starting With 选择灯光

Speed 速度

Speed Random 速度随机

Max Delay（Seconds）最大延迟（秒）

Delay Random 延迟随机

Orient To Path 方向

### Circle 圆环

Origin Type 起点类型

Point 点

Particle Origin 粒子起点

Emitter Origin 发射器起点

Origin XY XY 轴起点

Origin Z Z 轴起点

Angle X X 轴角度

Angle Y Y 轴角度

Angle Z Z 轴角度

Offset X X 轴偏移

Offset Y Y 轴偏移

Offset Z Z 轴偏移

Speed 速度

Speed Random 速度随机

### Sine 正弦

Origin Type 起点类型

Point 点

Particle Origin 粒子起点

Emitter Origin 发射器起点

Motion Type 运动类型

Angle 角度

Point 点

Particle 粒子

Origin XY XY 轴起点

Origin Z Z 轴起点

Angle X X 轴角度

Angle Y Y 轴角度

Angle Z Z 轴角度

Speed 速度

Speed Random 速度随机

Amplitude 振幅

Amplitude Random 振幅随机

Frequency 频率

Decay 衰减

### Bounce 弹跳

Bounce Origin Type 弹跳起点类型

Point 点

Offset 偏移

Origin XY XY 轴起点

Origin Z Z 轴起点

Angle X X 轴角度

Angle Y Y 轴角度

Angle Z Z 轴角度

Offset X X 轴偏移

Offset Y Y 轴偏移

Offset Z Z 轴偏移

Amplitude 振幅

Frequency 频率

Decay 衰减

Bounce in（Seconds）弹跳（秒）

Max Delay（percent）最大延迟（百分比）

Delay Random 延迟随机

Bounce Random 弹跳随机

Position（percent）位置（百分比）

Scale（percent） 缩放（百分比）

### Look At 朝向

Starting With 灯光选择

Motion Over Life 运动跟随生命
