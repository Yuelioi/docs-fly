---
title: Path-Property
order: 10
---

# Path Property 路径属性

::: info Note
2018 之后可用
:::

## Path.points()

用法：path.points(t = time)

说明：获取路径上所有顶点的 x、y 坐标。

其他说明 1：图层蒙版路径顶点的坐标是相对于其左上角的的原点。此时移动图层，路径点会发生改变。下图白色为蒙版

![](https://cdn.yuelili.com/20210913231336.png)

其他说明 2：形状路径顶点的坐标，是相对于形状组的锚点原点，**不是图层锚点**。此时移动图层，路径点不变，因为跟图层本身没啥关系，是**形状组的锚点**！

![](https://cdn.yuelili.com/20210913232212.png)

![](https://cdn.yuelili.com/20210913232057.png)

类型：嵌套数组，四舍五入到小数点后四位。

参数：t （可选）数字。路径的时间（以秒为单位）。默认值为 time（当前时间）。

示例：在一个纯色图层上创建一个蒙版

```javascript
thisComp
  .layer("Solid 1")
  .mask("Mask 1")
  .maskPath.points((t = time));
//返回[[200,100],[100,200],[200,300],[300,200]]
```

## Path.inTangents()

用法：Path.inTangents(t = time)

说明：获取路径上所有入点手柄的二维坐标。坐标值是相对于父点。 [0,0] 代表在入点处不产生弯曲。

类型：嵌套数组，四舍五入到小数点后四位。

参数：t （可选）数字。采样路径的合成时间（以秒为单位）。默认值为时间（当前时间）。

示例：纯色图层绘制一个蒙版

![](https://cdn.yuelili.com/20210913234241.png)

```javascript
thisComp.layer("纯色 1").mask("蒙版 1").maskPath.inTangents();
//返回[[-200,0],[200,-100]]
```

## Path.outTangents()

用法：Path.outTangents(t = time)

说明：获取路径上所有出点手柄的二维坐标。坐标值是相对于父点。 [0,0] 代表在入点处不产生弯曲。

类型：嵌套数组，四舍五入到小数点后四位。

参数：t （可选）数字。采样路径的合成时间（以秒为单位）。默认值为时间（当前时间）。

示例：纯色图层绘制一个蒙版（跟入点一样，我就不单独截图了）

![](https://cdn.yuelili.com/20210913234241.png)

```javascript
thisComp.layer("纯色 1").mask("蒙版 1").maskPath.outTangents();
//返回[[-200,0],[200,-100]]
```

## Path.isClosed()

用法：path.isClosed()

说明：路径是开放的还是闭合的。如果路径是闭合的，则返回 true；如果路径是开放的，则返回 false。

类型：布尔值

参数：无

```javascript
thisComp.layer("Solid 1").mask("Mask 1").maskPath.isClosed();
//闭合，返回true
```

## Path.pointOnPath()

用法：Path.pointOnPath(percentage = 0.5, t = time)

说明：获取路径上任意点的 x、y 坐标。该点表示为路径弧线长度的百分比。第一个点为 0%，最后一个点为 100%。

如果路径是闭合的，那么 0% 和 100% 将返回相同的坐标。

参数：

- 百分比 （可选）0 到 1 之间的数字。默认值为 0.5。
- t （可选）数字。采样路径的合成时间（以秒为单位）。默认值为时间（当前时间）。

类型：数组。

注意：蒙版与形状，返回的点依据不一样。详见 Path.points()的其他说明

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/path-point2.png)

```javascript
thisComp.layer("Turquoise Solid 1").mask("Mask 1").maskPath.pointOnPath(0.5);
//第50%个点，也就是第三个点。返回[200,300]
```

## Path.tangentOnPath(percentage = 0.5, t = time)

用法：Path .tangentOnPath(percentage = 0.5, t = time)

说明：获取路径上任意点的出点手柄的标准化矢量 x、y 坐标。

标准化矢量：x² +y² = 1

参数：

- 百分比 （可选）0 到 1 之间的数字。采样沿路径的弧长的百分比。值小于 0 和大于 1 的部分会被剪掉。默认值为 0.5。
- t （可选）数字。采样路径的合成时间（以秒为单位）。默认值为时间（当前时间）。

类型：数组。

示例：获取某一出点的标注化矢量坐标

![](https://cdn.yuelili.com/20210914000913.png)

```javascript
thisComp.layer("Solid 1").mask("Mask 1").maskPath.tangentOnPath(0.5);
//返回[1,0]
```

## Path .normalOnPath()

用法：Path.normalOnPath(percentage = 0.5, t = time)

类型：数偶数组。

说明：获取路径上经过计算的任意点的**法线**的标准化矢量 x、y 坐标。

标准化矢量：x² +y² = 1

参数：

- 百分比 （可选）0 到 1 之间的数字。采样沿路径的弧长的百分比。值小于 0 和大于 1 的部分会被剪掉。默认值为 0.5。
- t （可选）数字。采样路径的合成时间（以秒为单位）。默认值为时间（当前时间）。

示例：返回 0,1 ，懒得截图了，你懂就好

![](https://cdn.yuelili.com/20210914000913.png)

```javascript
thisComp.layer("Solid 1").mask("Mask 1").maskPath.tangentOnPath(0.5);
//返回[0,1]
```

## Path.createPath()

用法：Path.createPath(points = [[0,0], [100,0], [100,100], [0,100]], inTangents = [], outTangents = [], is_closed = true)

类型：路径对象。

说明：根据坐标点和切线点创建路径对象。

参数：

- points：路径点的 [x,y] 坐标的复合数组,如 [ [x1,y1] , [x2,y2] ]。必需，除非不传递任何参数，例如，createPath()。
- is_closed （可选）布尔值。决定蒙版是否为闭合。如果为 true，最后一个点会连接到第一个点。默认为 true。
- inTangents，复合数组，表示路径入点手柄的 [x，y] 位移坐标。必需，
- outTangents，复合数组，表示路径出点手柄的 [x，y] 位移坐标。必需，

示例 1：以下表达式将通过不传递 inTangents 或 outTangents 参数，删除"蒙版 1"的曲线：

```javascript
myMask = mask("Mask 1").path;
myMask.createPath(myMask.points());
```

示例 2：下面的示例通过传递"蒙版 1"的 points 和 tangents，并通过将 is_closed 设置为 false 来将其转换为开放路径：

```javascript
myMask = mask("Mask 1").path;
myMask.createPath(myMask.points(), myMask.inTangents(), myMask.outTangents(), false);
```
