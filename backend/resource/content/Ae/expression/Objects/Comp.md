---
title: Comp
order: 2
---

# Comp 合成

## layer()

用法：layer(index) / layer(name)

说明：根据图层索引/名称引用图层。一般要与其他参数配合使用

参数：iname 为字符串

类型：图层

示例：
![](https://mir.yuelili.com/user/AE/expression/exp-5-1.bmp)
![](https://mir.yuelili.com/user/AE/expression/exp-5-1.bmp)

```javascript
thisComp.layer(3).transform.scale; //返回[100,100]。第3个图层的缩放为100,100
thisComp.layer("Medium Solid2").transform.scale; //返回[100,100]。第3个图层的缩放为100,100
```

## layer(otherLayer,relIndex)

用法：layer(otherLayer,relIndex)

说明：根据其他图层引用新图层。一般要与其他参数配合使用

参数：otherLayer 是其他图层，relIndex 是相对顺序

类型：图层

示例：

![](https://mir.yuelili.com/user/AE/expression/exp-5-1.bmp)

```javascript
thisComp.layer(thisComp.layer(3), 1).name; //返回"<empty layer>2"。返回第4个图层的名称,第3个图层向下数一个即可。
```

## layerByComment()

用法：XXcomp.layerByComment(comment)

说明：根据图层的评论引用该图层。评论在时间轴右键-列数-评论开启。

参数：comment 为字符串

类型：图层

示例：

```javascript
thisComp.layerByComment("我好帅").name; //返回带有评论＂我好帅＂的图层的名字。
```

## marker

:::info Note
你不能通过标记来访问一个合成标记。如果你有老版本创建的项目，在表达式中使用合成标记的数字，你必须改变这些调用，使用 marker.key(name)来代替。因为合成标记的默认名称是一个数字，将引用转换为使用名称，通常只是在数字周围加上引号的问题。
:::

参数：无

类型:标记属性

示例：

```javascript
thisComp.layer("Solid 1").marker.key(1).comment; //返回图层第1个标记的评论
thisComp.marker.key(1).duration; //返回本合成第1个标记的持续时间
```

## marker.key

用法: key(index) / key(name)

说明: 通过标记索引值或者名称访问标记

类型:标记属性

示例:

```javascript
thisComp.marker.key(1).time;
thisComp.marker.key("0").time;
```

## marker.nearestKey

用法: marker.nearestKey(t)

说明: 当前时间最近的标记

类型:标记属性

示例:

```javascript
thisComp.marker.nearestKey(1).time; // 离1秒最近的标记
thisComp.marker.nearestKey(time).time; // 离当前时间最近的标记
```

## marker.numKeys

用法: 合成对象.numKeys

说明: 合成一共多少个标记

类型: 数字

示例:

```javascript
thisComp.marker.numKeys; // 5 (5个标记)
```

## numLayers

用法：XX 合成.numLayers

说明：用于返回合成的图层个数

参数：无

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/layer4-1.bmp)
![](https://mir.yuelili.com/user/AE/expression/layer4-1.bmp)

```javascript
thisComp.numLayers; //返回当前合成的图层个数，当前合成有4个图层，故返回4
```

## activeCamera

说明： 返回在当前帧渲染合成的摄像机对象。此摄像机未必是您通过其查看"合成"面板的摄像机。一般配合摄像机的其他属性使用，见示例

类型：摄像机

示例:

![](https://mir.yuelili.com/user/AE/expression/camera-1.bmp)
![](https://mir.yuelili.com/user/AE/expression/camera-1.bmp)

```javascript
thisComp.activeCamera; //2019返回[OBJECT CAMERA],2020返回"摄像机1"
thisComp.activeCamera.depthOfField; // 返回0,因为我没开摄像机的景深
```

合成设置图

![](https://mir.yuelili.com/user/AE/expression/exp-6-1.bmp)
![](https://mir.yuelili.com/user/AE/expression/exp-7-1.bmp)
![](https://mir.yuelili.com/user/AE/expression/exp-6-1.bmp)

## width

用法：XX.width

说明：宽度属性

参数：无

类型：数值

示例：

```javascript
thisComp.width; //返回1920，当前合成的宽度为1920像素
```

## height

用法：XX.height

说明：高度

参数：无

类型：高度值

示例：

```javascript
thisComp.height; //返回1080，当前合成的高度为1080像素
```

## duration

用法：XX.duration

说明：持续时间

参数：无

类型：数值

示例：

```javascript
thisComp.duration; //返回当前合成的持续时间。合成设置-持续时间里，当前为30s
```

## ntscDropFrame

用法：XXComp.ntscDropFrame

说明：判断合成是否为丢帧格式。关于丢帧请自行百度(29.97 就是丢帧)

参数：无

类型：布尔值

示例：

```javascript
thisComp.ntscDropFrame; //丢帧，返回True
```

## displayStartTime

用法：XX 合成.displayStartTime

说明：合成的开始时间，可以在 合成 - 合成设置(ctrl k) - 开始时间码处进行设定

参数：无

类型：数值

示例：

```javascript
thisComp.displayStartTime; //在合成-合成设置-起始时间(Start Timecode)进行设置。当前为0s
```

## frameDuration

用法：合成.frameDuration

说明：帧持续时间。合成为 20 帧/s 时，帧持续时间=1/20=0.05

参数：无

类型：数值

示例：

```javascript
thisComp.frameDuration; //返回0.03333。当前合成为30帧/s
```

## shutterAngle

用法：XXComp.shutterAngle

说明：快门角度值（以度为单位）。在合成-合成设置-快门角度(shutter angle)进行设置

参数：无

类型：数值

示例：

```javascript
thisComp.shutterAngle; //返回180。默认为180°
```

## shutterPhase

用法：XXComp.shutterPhase

说明：快门相位。在合成-合成设置-快门相位(shutter phase)进行设置

参数：无

类型：数值

示例：

```javascript
thisComp.shutterPhase; //返回90。
```

## bGColor

用法：XXComp.bgColor

说明：合成的背景色

参数：无

类型：四维数组，RGBA

示例：

```javascript
thisComp.bgColor; //合成-合成设置-背景颜色可以设置（见合成设置图）。当前为 [0,0,0,1] 也就是黑色。
```

## pixelAspect

说明：像素高宽比

类型：数值。

示例：  
![](https://mir.yuelili.com/user/AE/expression/exp-8-1.bmp)

```javascript
thisComp.pixelAspect; //0.909
```

## name

用法：XX.name

说明：某个对象或属性的名称

参数：无

类型：一般为字符串

示例：

```javascript
thisComp.layer("Black Solid 1").name; // 类型：Black Solid 1
thisComp.name; //类型：Comp1
```
