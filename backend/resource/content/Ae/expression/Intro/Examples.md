---
title: Examples
order: 2
---

# 示例

:::info Note

Dan Ebberts 提供的基础示例
:::

---

## 获取工程名称 (AE 15.1+ only)

先获取完整路径,接着再处理路径字符串:

::: danger 月离备注
试了没用. $.os 应该是脚本语法,这里应该发错了
:::

```javascript
var aepName = thisProject.fullPath.split($.os.indexOf("Windows") > -1 ? "\\" : "/").pop();
```

也可以保留未保存工程的名称

```javascript
var aepName = thisProject.fullPath.split($.os.indexOf("Windows") > -1 ? "\\" : "/").pop();
aepName = aepName === "" ? "Unsaved" : aepName;
```

---

## 让图层圆形旋转

1. 选择图层,给位置打表达式

2.输入以下表达式:

```javascript
[thisComp.width / 2, thisComp.height / 2] + [Math.sin(time) * 50, -Math.cos(time) * 50];
```

---

## 时针旋转

使用螺旋挑鞭(表达式选择器)来连接各层之间的旋转值，使时钟上的指针旋转的同时，分针随其选择。

如果两个指针分别设置关键帧，这种类型的动画 K 起来费时费力，但使用挑鞭，分分钟搞定。

1. 绘制一个时针和分针(随便 2 个矩形就行)

2. 把它们的锚点拉倒图层尾部

3. 移动图层,使其锚点重合在合成中心

4. 为时针设置旋转关键帧。

5. 选择分针的旋转属性 : `动画 > 添加表达式.`

6. 拖动分针的螺旋挑鞭(表达式选择器)至时针,会以下表达式:

```javascript
thisComp.layer("时针").rotation;
```

7. 设置分针旋转速度为时针的 12 倍,后面加`* 12`即可:

```javascript
thisComp.layer("hour hand").rotation * 12;
```

---

## 图层在另外 2 个图层中间

This example expression positions and maintains one layer at a balanceddistance between two other layers.

1. Start with three layers.

2. Animate the positions of the first two layers in the Timeline panel.

3. Select the third layer, press P to reveal the Position property, and Alt-click (Windows) or Option-click (Mac OS) the stopwatch button to the left of the property name.

4. Enter the following in the expression field:

```javascript
(thisComp.layer(1).position + thisComp.layer(2).position) / 2;
```

---

## 创建图片拖尾

This example expression instructs a layer to be at the same position as thenext higher layer in the Timeline panel, but delayed by a specified amount oftime (in this case, 0.5 seconds). You can set similar expressions for theother geometric properties.

1. Start with two solid-color layers that are scaled to approximately 30% of the composition size. (See Solid-color layers and solid-color footage items.)

2. Animate the position of the first layer.

3. Select the second layer, press P to reveal the Position property, and Alt-click (Windows) or Option-click (Mac OS) the stopwatch button to the left of the property name.

4. Enter the following in the expression field:

```javascript
thisComp.layer(thisLayer, -1).position.valueAtTime(time - 0.5);
```

5. Duplicate the last layer five times by selecting it and pressing Ctrl+D (Windows) or Command+D (Mac OS) five times.

All layers follow the same path, and each is delayed 0.5 seconds from the
previous.

:::info Note

Dan Ebberts provides more examples and techniques for creating trails ofimages on his [MotionScript](http://www.motionscript.com/mastering-expressions/follow-the-leader.html).
:::

---

## Create a bulge between two layers

This example expression synchronizes the Bulge Center argument of the Bulgeeffect in one layer with the position of another layer. For example, you cancreate an effect that looks like a magnifying glass moving over a layer, withthe contents under the magnifying glass bulging as the lens (that is, theoverlying layer) moves. This expression uses the fromWorld method, which makesthe expression work correctly regardless of whether you move the magnifyingglass layer or the underlying layer. You can rotate or scale the underlyinglayer, and the expression stays intact.

You can also use other effects, such as Ripple, with this expression.

1. Start with two layers. Make one layer a magnifying glass or similar object with a hole in the middle and name it Magnifier. (See Creating layers.)

2. Animate the position of the magnifying glass layer. (See Motion paths.)

3. Apply the Bulge effect to the other layer. (See Apply an effect or animation preset.)

4. Select the Bulge Center property of the Bulge effect in the Timeline panel and choose Animation > Add Expression, or Alt-click (Windows) or Option-click (Mac OS) the stopwatch button for the property.

5. Select the default expression text and type the following:

```javascript
fromWorld(thisComp.layer("Magnifier").position);
```

---

## Fade opacity of a 3D layer based on distance from camera

Apply the following expression to the Opacity property of a 3D layer:

```javascript
startFade = 500; // Start fade 500 pixels from camera.
endFade = 1500; // End fade 1500 pixels from camera.
try {
  // Check whether there's a camera
  C = thisComp.activeCamera.toWorld([0, 0, 0]);
} catch (err) {
  // Nocamera, so assume 50mm
  w = thisComp.width * thisComp.pixelAspect;
  z = w / 2 / Math.tan(degreesToRadians(19.799));
  C = [0, 0, -z];
}
P = toWorld(anchorPoint);
d = length(C, P);
linear(d, startFade, endFade, 100, 0);
```

The fade starts at a distance of `500` pixels from the camera and is completeat `1500` pixels from the camera. The linear interpolation method is used tomap distance values to opacity values.

---

## Make a 3D layer invisible if facing away from camera

Apply the following expression to the Opacity property of a 3D layer:

```javascript
if (toCompVec([0, 0, 1])[2] > 0 ) value else 0
```

:::info Note

Dan Ebberts explains this expression on his[site](http://www.adobe.com/go/learn_ae_motionscriptinvisiblelayer).
:::

---

## Flip layer horizontally if facing away from camera

Apply the following expression to the Scale property of a 3D layer:

```javascript
if (toCompVec([0, 0, 1])[2] > 0 ) value 
else [-value[0], value[1], value[2]]
```

---

## Animate scale at each layer marker

Apply the following expression to a Scale property to make a layer wobble at
each marker:

```javascript
n = 0; t = 0;
if (marker.numKeys > 0){
  n =marker.nearestKey(time).index;
if (marker.key(n).time > time) n\--;
}
if (n >0) t = time - marker.key(n).time;
amp = 15;
freq = 5;
decay = 3.0;
angle =freq * 2 * Math.PI * t;
scaleFact = (100 + amp * Math.sin(angle) /Math.exp(decay * t)) / 100;
[value[0] * scaleFact, value[1] / scaleFact];

```

---

## Start or stop wiggle at specific time

You can use any expression in place of the wiggle expression used here, tobegin and end the influence of any expression at a specific time.

Apply the following expression to a property to wiggle it beginning at time 2
seconds:

```javascript
timeToStart = 2;
if (time > timeToStart) {
  wiggle(3, 25);
} else {
  value;
}
```

Apply the following expression to a property to stop wiggling it at time 4
seconds:

```javascript
timeToStop = 4;
if (time > timeToStop) {
  value;
} else {
  wiggle(3, 25);
}
```

Apply the following expression to a property to start wiggling it at time 2seconds and stop wiggling it at time 4 seconds:

```javascript
timeToStart = 2;
timeToStop = 4;
if (time > timeToStart && time < timeToStop) {
  wiggle(3, 25);
} else {
  value;
}
```

---

## Match camera focal plane to another layer

Apply the following expression to the Focus Distance property of a cameralayer to have its focus distance match the distance to the anchor point of a
layer named “target”:

```javascript
target = thisComp.layer("target");
V1 = target.toWorld(target.anchorPoint) - toWorld([0, 0, 0]);
V2 = toWorldVec([0, 0, 1]);
dot(V1, V2);
```

:::info Note

Dan Ebberts explains this expression example in detail on his[website](http://motionscript.com/design-guide/auto-focus.html).
:::
