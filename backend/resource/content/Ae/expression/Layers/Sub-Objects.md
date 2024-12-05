---
title: Sub-Objects
order: 5
---

# 子对象

## source

用法：source.XX

说明：当合成或者素材，作为图层使用时，返回合成或素材本身。一般要配合其他属性使用。

![](https://cdn.yuelili.com/20211024140901.png)

参数：无

类型：合成或者素材本身

示例 1：Comp1 里面有 Comp2

```javascript
source.duration; //返回Comp2的持续时间：40（s）
```

## sourceTime()

用法：sourceTime（t=time）

说明：返回图层源当前的时间值。

参数：t=time，为数值。你可以写具体的数值，或者写个 time（time 代表当前时间轴时间，相当于逐渐变换的数值）

类型：数值

示例：  
![](https://mir.yuelili.com/user/AE/expression/a-z/sourceTime1.png)

```javascript
//如果在当前时间输入：
thisLyaer.sourceTime(time); //返回0.5 。等于当前时间（2.5）减去图层源起点（2） = 2.5-2 =0.5
```

## sourceRectAtTime()

用法：sourceRectAtTime(t = time, includeExtents = false)

说明：

范围仅适用于形状图层和段落文本图层。

形状图层范围根据需要增加图层边界的大小。

段落文本图层返回段落框的边界。

参数：t 为时间数值，Extents 为布尔值

类型：

示例：[详细教程请戳我](https://www.yuelili.com/?p=10443)

## effect()

用法：effect(index)或 effect(name)

说明：根据索引或者名称引用效果。如果多个效果具有相同名称，则会使用最接近"效果控件"面板顶层的效果。

参数：index 为数值，name 为字符串。索引从 1 开始计算

类型：效果本身

示例：

```javascript
effect("阈值")("级别"); //返回阈值效果的级别数值
thisComp.layer("Solid 1").effect(1)("Level"); //返回纯色图层1下第1个效果(当然还是阈值啦)的level(级别)数值
```

## mask()

用法：mask（index）或 mask（name）

说明：根据蒙版序号或者蒙版名称引用蒙版。索引从 1 开始计算

参数：index 为数字，name 为字符串

类型：蒙版

示例：

```javascript
thisComp.layer("Solid 1").mask("Mask 1").maskFeather[0]; //返回蒙版1的羽化值
thisComp.layer("纯色层 1").mask(2).maskFeather[0]; //返回第2个蒙版的羽化值
```
