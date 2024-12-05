---
order: 4
title: Space-Transforms
---

# 空间变换

使用图层空间变换方法将值从一个空间变换到其他空间

"from"方法可将值从合成空间或世界空间变换到图层空间。"to"方法可将值从图层空间变换到合成空间或世界空间。

每个变换方法各采用一个可选参数来确定计算变换的时间；通常使用当前时间（默认）。

当变换方向矢量（例如两个位置值之间的差值）时使用"Vec"变换方法。

当变换点（例如位置）时，使用简单的（非"Vec"）变换方法。

2D 图层，合成和世界空间是一样的。然而，对于 3D 图层，合成空间与活动摄像机有关，而世界空间独立于摄像机。  
![](https://mir.yuelili.com/user/AE/expression/a-z/tocomp.png)

合成空间：可以在合成设置里更改空间的大小，一般固定。上图点 X 在合成里就是[540,0]

图层空间：图层自己本身的空间。上图点 X（位于图层左上角）在图层里的位置始终固定，为[0,0]

世界空间：增加摄像机以后，各种点都有了空间上的深度感，也就是 Z 轴，因此有世界空间这一概念

注：点默认在合成空间里

## toComp()

用法：XXlayer.toComp(point, t = time)

说明：把点在图层上的坐标，转换为点在合成上的坐标

参数：point，点在图层的坐标值。time 是图层的动画时间，比如位置在 2 秒内从[0,0]到[100,100]，此时 t=2，这时位置的 point=[100,100]。

类型：点在合成中的坐标值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tocomp-sample1.png)

```javascript
//如图，在2000*2000合成新建500*500的纯色层。它的锚点为[250,250]位于图层正中心
//此时给锚点添加表达式
thisLayer.toComp(transform.anchorPoint); //返回[1000,1000]。锚点在合成的[1000,1000]位置处

//如果打上三维开关
thisLayer.toComp(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250,0]转换为[1000,1000,2776.8]
```

## fromComp()

用法：XXlayer.fromComp(point, t = time)

说明：把点在合成上的坐标，转换为点在图层上的坐标

参数：point，点在合成的坐标值。time 是图层的动画时间，比如位置在 2 秒内从[0,0]到[100,100]，此时 t=2，这时位置的 point=[100,100]。

类型：点在图层中的坐标值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tocomp-sample2.png)

```javascript
//如图，在2000*2000合成新建500*500的纯色层。它的锚点为[250,250]位于图层正中心
//此时给锚点添加表达式
thisLayer.fromComp(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250]转换为[-500,-500]

//如果打上三维开关
thisLayer.fromWorld(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250,0]转换为[-249.7,-249.7,-2776.8]
```

## toWorld()

用法：XXlayer.toWorld(point, t = time)

说明：将点从图层空间变换到与视角无关的世界空间。

参数：point，点在图层的坐标值。time 是图层的动画时间，比如位置在 2 秒内从[0,0]到[100,100]，此时 t=2，这时位置的 point=[100,100]。

类型：点在世界坐标系的坐标值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tocomp-sample1.png)

```javascript
//如图，在2000*2000合成新建500*500的纯色层。它的锚点为[250,250]位于图层正中心
//此时给锚点添加表达式
thisLayer.toWorld(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250]转换为[1000,1000]

//如果打上三维开关
thisLayer.toWorld(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250,0]转换为[1000,1000,0]
```

## fromWorld()

用法：XXlayer.fromWorld(point, t = time)

说明：将点从世界空间变换到与图层空间。

参数：point，点在世界的坐标值。time 是图层的动画时间，比如位置在 2 秒内从[0,0]到[100,100]，此时 t=2，这时位置的 point=[100,100]。

类型：点在图层的坐标值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/tocomp-sample2.png)

```javascript
//如图，在2000*2000合成新建500*500的纯色层。它的锚点为[250,250]位于图层正中心
//此时给锚点添加表达式
thisLayer.fromWorld(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250]转换为[-500,-500]

//如果打上三维开关
thisLayer.fromWorld(transform.anchorPoint); //图上为toComp表达式，本表达式结果会把[250,250,0]转换为[-500,-500,0]
```

## toCompVec(vec,t=time)

说明：将矢量从图层空间变换到合成空间。

返回类型：数组 [2 或 3]。

参数类型：vec 是数组 [2 或 3]，t 是数值。

详情见 toComp，只不过转的不是位置，而是矢量，矢量是有方向的。

示例：toCompVec([1,0])

## fromCompVec(vec,t=time)

返回类型：数组 [2 或 3]。

参数类型：vec 是数组 [2 或 3]，t 是数值。

将矢量从合成空间变换到图层空间。

详情见 fromComp，只不过转的不是位置，而是矢量，矢量是有方向的。

示例（2D 图层）：

dir=sub(position, thisComp.layer(2).position);

fromCompVec(dir)

## toWorldVec(vec,t=time)

返回类型：数组 [2 或 3]。

参数类型：vec 是数组 [2 或 3]，t 是数值。

将矢量从图层空间变换到世界空间。

详情见 toWorld，只不过转的不是位置，而是矢量，矢量是有方向的。

示例：

p1 = effect("Eye Bulge 1")("Bulge Center");

p2 = effect("Eye Bulge 2")("Bulge Center");

toWorld(sub(p1, p2))

## fromWorldVec(vec,t=time)

返回类型：数组 [2 或 3]。

参数类型：vec 是数组 [2 或 3]，t 是数值。

将矢量从世界空间变换到图层空间。

详情见 fromWorld，只不过转的不是位置，而是矢量，矢量是有方向的。

示例：fromWorld(thisComp.layer(2).position)

## fromCompToSurface(point,t=time)

返回类型：数组 [2]。

参数类型：point 是数组 [2 或 3]，t 是数值。

在从活动摄像机中进行查看时出现的位置将位于合成空间中的点投影到图层表面上的点（零 z 值）。此方法有助于设置效果控制点。仅用于 3D 图层。
