---
title: "【动效周期表】TipShape 尖端形状"
date: "2020-06-03"
categories: 
  - "archive"
---

 

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape.gif"></td><td><h2 style="font-size: 36pt;">TipShape</h2><div></div><span style="font-size: 24pt;">尖端形状</span><div></div>一种运动，用附在形状尖端的形状（例如箭头）伸展和收缩。 对于图形和解释而言，将箭头和数字关联起来非常方便。 是一个方便的动作，且很容易实现。 为了使延长线的时间与路径修整和图形顶端的位置同步， 我们必须编写一个稍微复杂的表达式。</td></tr></tbody></table>

## 制作方法

1.首先在形状图层上绘制原始图形的路径。自己手画一个！

![](https://cdn.yuelili.com/2021/20210809013233.png)

2.应用修剪路径并点击关键帧。

3.应用“效果” - “表达式控制” - “滑块控制”的效果。

4.拾取表达式鞭子链接修剪路径和滑块控制。（我用的开始）

![](https://cdn.yuelili.com/2021/20210809013438.png) 如果在移动滑块控件，这条线会来回伸缩

5.使用新的形状图层，创建要附加到尖端的形状。（箭头，圆圈等）

![](https://cdn.yuelili.com/2021/20210809013926.png)

6.复制在1.中创建形状的路径。粘贴到尖端图形的位置属性中。 （必须粘贴到尖端图层的变换的位置属性。）如果你用的多边形工具（没有单独的路径），需要右键多边形路径，转为贝塞尔路径，再复制

![](https://cdn.yuelili.com/2021/20210809020835.png)

![](https://cdn.yuelili.com/2021/20210809021033.png)

另外，必须在0帧粘贴。 \*请注意，不能很好地使用路径属性。 \*如果粘贴时间不在0帧，后面你得自己改表达式。

7.在圆形的“位置”会创建2秒钟的关键帧。

![](https://cdn.yuelili.com/2021/20210809014008.png)

8.给滑块K个动画 0~2s

我是K的 0到100

9.将以下表达式添加到路径形状的“修建路径”。（我用的开始）

effect("滑块控制")("滑块").valueAtTime(2 \*(effect("滑块控制")("滑块"))/ 100);

尖端图形会与路径同步。

![](https://cdn.yuelili.com/2021/20210809014845.gif)

## 应用

1.一种运动，通过在创建过程来复制圆点。 此外，尖端图形的高度和水平线的高度通过表达式同步。

<table><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/tipshape/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/move/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Move.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape-Ex001.gif"></td></tr></tbody></table>

2.使用尖端形状作为框架而不是图形的示例。

<table><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/tipshape/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/plus.png"></td><td><a href="https://yuelili.com/archive/slide/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Slide.gif"></a></td><td><img class="plus" src="https://mir.yuelili.com/user/AE/mg/foxcodex/tri.png"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape-Ex002.gif"></td></tr></tbody></table>

## 例

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape-Ex001.gif"></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/TipShape-Ex002.gif"></td></tr></tbody></table>

## 相似/相关

<table style="border-collapse: collapse;"><tbody class="table1"><tr><td><a href="https://yuelili.com/archive/shapetransform/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/ShapeTransform.gif"></a></td><td><a href="https://yuelili.com/archive/easing/"><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Easing.gif"></a></td><td><img src="https://mir.yuelili.com/user/AE/mg/foxcodex/Scale.gif"></td></tr></tbody></table>
