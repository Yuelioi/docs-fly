---
title: 形状
order: 1
---

# 形状

![](https://mir.yuelili.com/2021/07/cd4156f412bf886fe8f6bf1e10678914.png)

| Group     | 组（空） |                  | 混合模式不赘述 |                 |                |          |        |
| --------- | -------- | ---------------- | -------------- | --------------- | -------------- | -------- | ------ |
|           |          | Transform：Group | 变换：组       |                 |                |          |        |
|           |          |                  |                | Anchor Point    | 锚点           |          |        |
|           |          |                  |                | Position        | 位置           |          |        |
|           |          |                  |                | Scale           | 比例           |          |        |
|           |          |                  |                | Skew            | 倾斜           |          |        |
|           |          |                  |                | Skew Axis       | 倾斜轴         |          |        |
|           |          |                  |                | Rotation        | 旋转           |          |        |
|           |          |                  |                | Opacity         | 不透明度       |          |        |
| Rectangle | 矩形     |                  |                |                 |                |          |        |
|           |          | Rectangle Path   | 矩形路径       |                 | 混合模式不赘述 |          |        |
|           |          |                  |                | Size            | 大小           |          |        |
|           |          |                  |                | Position        | 位置           |          |        |
|           |          |                  |                | Roundness       | 圆度           |          |        |
| Ellipse   | 椭圆     |                  |                |                 | 混合模式不赘述 |          |        |
|           |          | Ellipse Path     | 椭圆路径       |                 |                |          |        |
|           |          |                  |                | Size            | 大小           |          |        |
|           |          |                  |                | Position        | 位置           |          |        |
| Polystar  | 多边星形 |                  |                |                 |                |          |        |
|           |          | Polystar Path    | 多边星形路径   |                 | 混合模式不赘述 |          |        |
|           |          |                  |                | Type            | 类型           |          |        |
|           |          |                  |                |                 |                | Polystar | 多边形 |
|           |          |                  |                |                 |                | Star     | 星形   |
|           |          |                  |                | Points          | 点             |          |        |
|           |          |                  |                | Position        | 位置           |          |        |
|           |          |                  |                | Rotation        | 旋转           |          |        |
|           |          |                  |                | Inner radius    | 内径           |          |        |
|           |          |                  |                | Outer radius    | 外径           |          |        |
|           |          |                  |                | Inner roundness | 内圆度         |          |        |
|           |          |                  |                | Outer roundness | 外圆度         |          |        |
| Path      | 路径     |                  |                |                 |                |          |        |

## 形状类型

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-type.png)

形状共有五种类型。矩形、圆角矩形、椭圆、多边形、星形，但是有两种可以通过其他进行衍生（比如圆角矩形 =
矩形+圆度，多边形和星形，统称为多边星形），故有 3 大类型：矩形、椭圆、多边星形

## 矩形&圆角矩形

![](https://mir.yuelili.com/user/AE/mg/shape-layer/rectangle-property1.png)

| ![](https://mir.yuelili.com/user/AE/mg/shape-layer/rectangle-property.png) | ![](https://mir.yuelili.com/user/AE/mg/shape-layer/rectangle-property2.png) |
| --------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |

size 大小：矩形的长和宽，如果固定长宽比，再改成一样大小，此时矩形就会变成正方形。

position 位置：此为矩形路径的位置，不同于变换-位置属性。因为一个形状图层可以塞一堆形状路径，用于单独控制一个路径的位置，也是必不可少的。

roundness 圆度：基于四个角向内进行画圆，圆的半径值就是圆度。增加圆度，甚至可以把正方形变成圆形，矩形 → 圆角矩形 → 圆。

## 圆形

size 大小：椭圆的长和宽，长宽相等，则是圆形。

position 位置：圆形的位置，不同于变换-位置属性。因为一个形状图层可以塞一堆形状路径，用于单独控制一个路径的位置，也是必不可少的。

## 多边形

Points 点：控制多边形的边数  
Position 位置：不同于变换-位置属性。因为一个形状图层可以塞一堆形状路径，用于单独控制一个路径的位置，也是必不可少的。  
Rotation 旋转：旋转  
Outer radius 外径：控制多边形的大小，参考矩形大小  
Outer roundness 外圆度：基于边的角点，进行圆化。见下图

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Outer-roundness.png)

## 星形

![](https://mir.yuelili.com/user/AE/mg/shape-layer/star-shape.png)

Points 点：控制多边形的边数  
Position 位置：不同于变换-位置属性。因为一个形状图层可以塞一堆形状路径，用于单独控制一个路径的位置，也是必不可少的。  
Rotation 旋转：旋转  
Inner radius 内径：控制星形内侧半径的大小，见上图  
Outer radius 外径：控制星形外侧半径的大小，见上图  
Inner roundness 内圆度：基于边的角点，进行圆化。见上图  
Outer roundness 外圆度：基于边的角点，进行圆化。见上图

示例 1：

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Inner-roundness.png)

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Outer-roundness2.png)

内径/外径有一个为 0，则两者图像一致

可以看出：

1.外径内径只不过是一个代称，当内径大于外径，那么内径就变成外径了。

2.其中一个为 0 时，正负也没区别

示例 2：

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Outer-roundness7.png)

---

![](https://mir.yuelili.com/user/AE/mg/shape-layer/Outer-roundness6.png)  
![](https://mir.yuelili.com/user/AE/mg/shape-layer/Outer-roundness5.png)

## 技巧

tip1：双击形状工具进行创建，可以创建与合成大小一致的形状

tip2：按住 shift 创建，可以创建正方形/圆形/正多边形

tips3：按住 ctrl shift 进行创建，可以基于拖拽点，创建正方形/圆形/正多边形

tip4：ctrl alt home 可以使锚点居中

tips5：创建一个圆，结果是椭圆，请查看合成的像素长宽比是不是 1。

tips6：形状图层如何创建蒙版（创建时，单击以下按钮，就不会追加形状，而是新建蒙版了）

![](https://mir.yuelili.com/user/AE/mg/shape-layer/shape-mask.png)

top7：ctrl shift H 可以关闭边线显示
