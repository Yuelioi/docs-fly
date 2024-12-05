---
title: Vector
order: 8
---

# Vector 矢量数学

## 介绍

矢量数学函数是对数组进行运算的全局方法。与内置的 JavaScript 方法（例如 Math.sin）不同，这些方法不用 Math 前缀。

除非另有说明，矢量数学方法对维度的要求很宽松，返回输入数组对象最大维度的值，用零填充缺失的维度元素。

例如，表达式 add([10, 20], [1, 2, 3]) 返回 [11, 22, 3]。

书写规格：[111,222,333] 中括号里几个维度就写几个数字

说明: 案例均在 AE 文字图层的源文本中用表达式书写,案例中返回的意思是在屏幕中显示内容

## add()

用法：add(vec1, vec2 )

说明：两个矢量相加。vec 是数组的意思, 如果 2 维数组 add3 维数组，则 2 维数组会补一位 0，见示例 2

参数：两个矢量

类型：数组

示例:

```javascript
add([1, 2, 3], [10, 20, 30]); //返回[11,22,33]
add([1, 2], [10, 20, 30]); //返回[11,22,30] 等同于add([1,2,0],[10,20,30])
```

## sub()

用法：sub(vec1,vec2)

说明：两个矢量相减。

参数：vec1 和 vec2 是数组。

类型：数组

示例：

```javascript
sub([111, 222, 333], [100, 200, 300]); // 返回[11,22,33] 单纯各个元素相减
sub([111, 222, 333], [100, 200]); // 返回[11,22,333] 缺少的元素按0计算
```

## mul()

用法：mul(vec,amount)

说明：将矢量的每个元素与数量相乘

参数：vec 是数组，amount 是数值

类型：数组

示例：

```javascript
mul([111, 222, 333], 2); // 返回[222,444,666] 单纯地各个元素乘以后面数字
```

## div()

用法：div(vec,amount)

说明：用矢量的每个元素除以数量

参数：vec 是数组，amount 是数值

类型：数组

示例：

```javascript
div([100, 200, 300], 10); //返回[10,20,30] 单纯地用数组元素除以数量
```

## clamp()

用法：clamp(value,limit1,limit2)

说明：把 value 每个元素的值都限定在 limit1 和 limit2 之间

参数：value、limit1 和 limit2 是数值或数组

类型：数值或数组

![](https://mir.yuelili.com/user/docs/exp-a-z/exp-clamp.png)

示例：value 为数值

![](https://mir.yuelili.com/user/AE/expression/clamp.bmp)

```javascript
clamp(time, 1, 3); //在3秒时返回3,10秒时返回3。
//time本身为线性函数。受函数影响，大于1且低于3的值返回time本身，超过3的值返回3，被限制。
```

示例：value 为数组时

![](https://mir.yuelili.com/user/AE/expression/exp-global13.bmp)

```javascript
clamp([time, time * 2], [10, 30], [20, 40]); // 在12秒时返回[12,30]; 在18秒时返回[18,36]
//数组同理 相当于独立计算每个元素的clamp 然后把结果合一起
//注：clamp处理time时,time以秒为单位
```

## dot()

用法：dot(vec1,vec2)

说明：返回矢量参数的点（内）积。(同维相乘再相加）

参数：vec1 和 vec2 是数组

类型：数值

示例：

```javascript
dot([7, 5], [3, 6]); // 返回51 = 7*3 + 5*6
```

## cross()

用法：cross(vec1, vec2)

说明：返回 vec1 和 vec2 的向量叉积。同时与 vec1 和 vec2 垂直的一个矢量 C

参数：vec 为一个矢量（二维或三维）。

类型：矢量。

示例：

A = [4 -2 1];
B = [1 -1 3];
C = cross(A,B) //[-5,-11,-2]

//使用点积验证 C 是否与 A 和 B 垂直。 dot(C,A)==0 & dot(C,B)==0

![](https://mir.yuelili.com/user/AE/expression/exp-global14.bmp)  
具体叉积请百度

## normalize()

用法：normalize(vec)

说明：标准化矢量使其长度为 1，勾股定理会算吧。等价于 div(vec, length(vec))

参数：vec 是数组

类型：数组

示例：

```javascript
normalize([3, 4]); //返回[0.6,0.8] 长度计算 1²=0.6²+0.8²
```

![](https://mir.yuelili.com/user/AE/expression/exp-global15.bmp?imageView2/1/w/778/h/754#)

## length()

全名 1：length(vec)

说明：返回矢量 vec 的长度

参数：vec 是数组

类型：数值

示例：

```javascript
length([3, 4]); //返回5 勾股定理
```

全名 2：length(point1,point2)

说明：返回两点之间的距离。等同于 length(sub(point1, point2)

参数：point1 和 point2 是数组。point2 参数可选

类型：数值

示例：

```javascript
length(position, pointOfInterest); // 将此表达式添加到摄像机的焦距属性中，可以使焦平面锁定到摄像机的目标点，以便目标点对准焦点
```

## lookAt()

用法：lookAt(fromPoint,atPoint)

说明：如果对摄像机使用此表达式，请关闭自动方向。

参数 fromPoint 为数组 [3]。要定向的图层的世界空间中的位置。

参数 atPoint 为数组 [3]。要将图层指向的世界空间中的点。

类型：数组 [3]。返回值可用作"方向"属性的表达式，使图层点的 z 轴指向 atPoint。此方法对摄像机和光照特别有用。

示例：

```javascript
lookAt(position, thisComp.layer(1).position); // 聚光灯的方向属性的以下表达式会将光点指向同一合成中的 1 号图层的锚点
lookAt(position, thisComp.layer("Light 1").transform.position); //给3D纯色图层的orientation属性添加表达式后，图层和灯光的z轴朝向一致。移动灯光，图层也一直面朝向灯光
```

关于世界空间、图层空间详见 tocomp
