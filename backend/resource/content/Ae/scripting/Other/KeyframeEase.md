---
title: KeyframeEase
order: 3
---

# KeyframeEase 关键帧缓动

## 关键帧缓动

myKey = new KeyframeEase(speed, influence);

描述：关键帧缓动对象（KeyframeEase）封装了 AE 属性的关键帧缓动设置。缓动程度取决于属性的 setTemporalEaseAtKey 方法设置的速度和影响值。本构造函数会创建一个 KeyframeEase 对象。

- speed：浮点值。设置速度（speed）属性。
- influence：浮点值，0.01~100.0。设置影响（influence）属性。

### 示例 1：设置第二个关键帧的缓入缓出

假定位置具有两个以上的关键帧。右键第二个关键帧 - 关键帧速度

![](https://cdn.yuelili.com/20211012171649.png)

![](https://cdn.yuelili.com/20211012171544.png)

```javascript
var easeIn = new KeyframeEase(0.5, 50);
var easeOut = new KeyframeEase(0.75, 85);
var myPositionProperty = app.project.activeItem.layer(1).property("Position");
myPositionProperty.setTemporalEaseAtKey(2, [easeIn], [easeOut]);
```

### 示例 2：缩放具有两或三个维度的时间属性

对于 2D 和 3D 属性，必须为每个尺寸设置一个 easeIn 和 easeOut 值，右键第二个关键帧 - 关键帧速度

![](https://cdn.yuelili.com/20211012171940.png)

![](https://cdn.yuelili.com/20211012171920.png)

```javascript
var easeIn = new KeyframeEase(0.5, 50);
var easeOut = new KeyframeEase(0.75, 85);
var myScaleProperty = app.project.activeItem.layer(1).property("Scale");
myScaleProperty.setTemporalEaseAtKey(2, [easeIn, easeIn, easeIn], [easeOut, easeOut, easeOut]);
```

## 属性

### KeyframeEase.influence

myKey.influence

![](https://mir.yuelili.com/2021/07/03ca9dfa94fbfc4e928639bdec1b1430.png)

![](https://mir.yuelili.com/2021/07/85dfc84604539ae9bc2943350a886df2.png)![](https://mir.yuelili.com/2021/07/25ca4531a399e908d29a6da36a0bd050.png)

描述：关键帧的影响值。右键关键帧 - 关键帧速度 - 影响。或者图表编辑器里，更改滑竿的长度。

类型：浮点值，0.1~100.0；读/写。

### KeyframeEase.speed

myKey.speed

描述：关键帧的速度值。单位取决于关键帧的类型，右键关键帧 - 关键帧速度

类型：浮点值；读/写。
