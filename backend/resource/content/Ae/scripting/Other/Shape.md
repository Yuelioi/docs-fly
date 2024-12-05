---
order: 7
title: Shape
---

# Shape 形状

## 形状对象 Shape object

app.project.item(index).layer(index).property(index).property("maskShape").value

描述：Shape 对象将描述形状的信息，封装在形状层或蒙版形状中。它是形状图层的“路径” 和蒙版中的"蒙版路径"。使用构造函数 new
Shape()来创建一个新的空 Shape 对象，然后再分别设置属性。

![](https://mir.yuelili.com/2021/07/a341b3650266fa71efcd5dd7e8a3c8d3.png)

- 顶点：

形状具有一组锚点（或叫顶点），每个顶点有一对方向手柄（或叫切线向量）。切线向量(在非 RotoBezier 蒙版中)确定绘制锚点绘制的线的方向。每个顶点都有一个入点切线向量和一个出点切线向量。

- 切线值

是相对于关联顶点指定的一对 x，y 坐标，就是手柄的位置。例如顶点为[0,0]，切线值为[-1，-1]，位于顶点的上方和左侧，并且具有 45 度的斜率。

- 闭合形状

如果未闭合形状，则忽略第一个顶点的 inTangent 和最后一个顶点的 outTangent。如果形状是封闭的，则这两个向量将指定最终连接线段的方向控点，使其从最终顶点移出，并返回到第一个顶点。

RotoBezier 蒙版会自动计算其切线。(请参见 MaskPropertyGroup.rotoBezier)如果在 RotoBezier 蒙版中使用了形状，则切线值将被忽略。这意味着，对于 RotoBezier 蒙版，你可以构建一个形状的 vertices 属性，inTangents 和 outTangents 都为空即可。访问新形状时，其切线值将自动填充计算。

对于封闭的蒙版形状，宽度可变的蒙版羽化点可以存在于蒙版路径上的任何位置。羽化点是“蒙版路径”属性的一部分。通过出现的蒙版路径段的编号(相邻顶点之间的路径部分)的编号来参考特定的羽化点。

::: info 提示
蒙版上的羽化点按创建顺序在阵列中列出。
:::

示例 1：创建一个正方形蒙版。正方形是具有 4 个顶点的封闭形状。连接的直线段的 inTangents 和 outTangents 为默认值 0，不需要显式设置。

```javascript
var myShape = new Shape();
myShape.vertices = [
  [0, 0],
  [0, 100],
  [100, 100],
  [100, 0],
];
myShape.closed = true;
```

示例 2：创建一个“ U”形的蒙版。“ U”是具有与正方形相同的 4 个顶点的开放形状。

```javascript
var myShape = new Shape();
myShape.vertices = [
  [0, 0],
  [0, 100],
  [100, 100],
  [100, 0],
];
myShape.closed = false;
```

示例 3：创建一个椭圆形。椭圆是具有 4 个顶点且具有 inTangent 和 outTangent 值的封闭形状。

```javascript
var myShape = new Shape();
myShape.vertices = [
  [300, 50],
  [200, 150],
  [300, 250],
  [400, 150],
];
myShape.inTangents = [
  [55.23, 0],
  [0, -55.23],
  [-55.23, 0],
  [0, 55.23],
];
myShape.outTangents = [
  [-55.23, 0],
  [0, 55.23],
  [55.23, 0],
  [0, -55.23],
];
myShape.closed = true;
```

示例 4：创建一个带有两个羽化点的方形蒙版。一个大的方形蒙版，具有两个羽化点，一个更靠近第二个蒙版片段(位于底部边缘)的左端，半径为 30 个像素，另一个则位于第三个蒙版片段(位于右侧边缘)的中间，半径较大 100 像素。

```javascript
var myShape = new Shape();
myShape.vertices = [
  [100, 100],
  [100, 400],
  [400, 400],
  [400, 100],
]; // segments drawn counter clockwise
myShape.closed = true;
myShape.featherSegLocs = [1, 2]; // segments are numbered starting at 0, so second segment is 1
myShape.featherRelSegLocs = [0.15, 0.5]; // 0.15 is closer to the lower-left corner of the square
myShape.featherRadii = [30, 100]; // second feather point (onright- sidesegment) has a larger radius
```

## 属性

### closed 形状闭合

shapeObject.value.closed

描述：如果为 true，则将第一个和最后一个顶点连接起来以形成闭合曲线。如果为 false，则不绘制结束线段。

类型：布尔值 读/写。

### featherInterps 羽化插值

shapeObject.value.featherInterps

描述：包含每个羽化点的半径插值类型的数组(0 表示非保持羽化点，1 表示保持羽化点)。

::: info 提示
值以创建羽化点的顺序存储在数组中。
:::

类型：整数数组(0 或 1)；读/写。

### featherRadii 羽化半径

shapeObject.value.featherRadii

描述：包含每个羽化点的半径(羽化量)的数组；内部羽化点具有负值。

::: info 提示
值以创建羽化点的顺序存储在数组中。
:::

类型：浮点值数组；读/写。

### featherRelCornerAngles

shapeObject.value.featherRelCornerAngles

描述：一个数组，其中包含在蒙版路径上某个角处的弯曲外部羽化边界两侧的每个法线之间的每个羽化点的相对角度百分比。对于不在角上的羽化点，角度值为 0％。

::: info 提示
值以创建羽化点的顺序存储在数组中。
:::

类型：浮点百分比值数组(0 到 100)；读/写。

### featherRelSegLocs 羽化路径 X？

shapeObject.value.featherRelSegLocs

描述：一个数组，其中包含每个羽化点在其蒙版路径线段(顶点之间的蒙版路径的截面，从 0 开始)的从 0 到 1 的相对位置。

::: info 提示
值以创建羽化点的顺序存储在数组中。要将羽化点移动到其他蒙版路径段，请首先更改 featherSegLocs 属性值，然后更改此属性。
:::

类型：浮点值数组(0 到 1)；读/写。

### featherSegLocs 羽化路径段号

shapeObject.value.featherSegLocs

描述：包含每个羽化点的蒙版路径段号(顶点之间的蒙版路径的部分，从 0 开始编号)的数组。

::: info 提示
值以创建羽化点的顺序存储在数组中。通过更改其片段号(此属性)和(可选)其 featherRelSegLocs 属性值，将一个羽化点移至其他片段。
:::

类型：整数数组；读/写。

示例

```javascript
// Assuming a rectangle closed mask (segments numbered 0-3) has 3 mask feather points,
// move all 3 feather points to the first mask segment.

// Get the Shape object for the mask, assumed here to be the first mask on the layer.
var my_maskShape = layer.mask(1).property("ADBE Mask Shape").value;

// Check where mask feather points are located.
// Note: They are stored in the order that they are added.
var where_are_myMaskFeatherPoints = my_maskShape.featherSegLocs;

// Move all 3 feather points to the first mask segment (numbered 0).
my_maskShape.featherSegLocs = [0, 0, 0];

// Update the mask path.
layer.mask(1).property("ADBE Mask Shape").setValue(my_maskShape);
```

### featherTensions 羽化张量

shapeObject.value.featherTensions

描述：一个数组，其中包含每个羽化点的张力量，从 0(0％张力)到 1(100％张力)。

::: info 提示
值以创建羽化点的顺序存储在数组中。
:::

类型：浮点值数组(0 到 1)；读/写。

### featherTypes 羽化方向

shapeObject.value.featherTypes

描述：一个数组，其中包含每个羽化点的方向，可以是 0(外部羽化点)或 1(内部羽化点)。

::: info 提示
创建羽化点后，便无法更改其方向。
值以创建羽化点的顺序存储在数组中。
:::

类型：整数数组(0 或 1)；读/写。

### inTangents 形状入点手柄

shapeObject.value.inTangents

描述：与形状的顶点相关联的输入切线向量或方向控制柄。将每个向量指定为两个浮点值的数组，然后将向量收集到与该数组相同长度的 vertices 数组中。

每个切线值默认为[0,0]。当蒙版形状不是 RotoBezier 时，这将导致直线段。

如果形状在 RotoBezier 蒙版中，则所有切线值将被忽略，并且将自动计算切线。

类型：浮点对数组 读/写。

### outTangents 形状出点手柄

shapeObject.value.outTangents

描述：与形状的顶点关联的传出切线向量或方向手柄。将每个向量指定为两个浮点值的数组，然后将向量收集到与数组长度相同的 vertices 数组中。

每个切线值默认为[0,0]。当蒙版形状不是 RotoBezier 时，这将导致直线段。

如果形状在 RotoBezier 蒙版中，则所有切线值将被忽略，并且将自动计算切线。

类型：浮点对数组 读/写。

### vertices 形状顶点

shapeObject.value.vertices

描述：形状的锚点。将每个点指定为两个浮点值的数组，然后将点对收集到完整点集的数组中。

示例

```javascript
myShape.vertices = [
  [0, 0],
  [0, 1],
  [1, 1],
  [1, 0],
];
```

类型：浮点对数组 读/写。
