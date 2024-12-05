---
order: 2
title: Genaral
---

# 常规

## width

说明：图层宽度值

类型：数值。

示例

```javascript
thisLayer.width; //返回1080。
```

## height

说明：图层高度值

类型：数值。

示例

```javascript
thisLayer.height; //返回1920
```

## index

说明：图层的索引号。从 1 开始计算

返回类型：数值。

示例

thisLayer.index //返回 1。第 1 个图层

## parent

说明：返回图层的父图层对象（如果有的话），父对象就是小鞭子勾的对象。

类型：图层、光照或摄像机对象。

示例：

thisLayer.parent.width

## hasParent

用法：XX.hasParent

说明：XX 是否有父对象

参数：无

类型：布尔值

示例：

```javascript
thisLayer.hasParent; //有父对象，返回True
```

## inPoint

用法：XX.inPoint

说明：图层入点时间。如果反转图层时间，则 inPoint 的值大于 outPoint 的值。

参数：无

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/sourceTime1.png)

thisLayer.inPoint.toFixed(1) //返回 1.6（秒）

## outPoint

用法：XX.inPoint

说明：图层出点时间。如果反转图层时间，则 inPoint 的值大于 outPoint 的值。

参数：无

类型：数值

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/sourceTime1.png)

```javascript
thisLayer.outPoint.toFixed(1); //返回7.6（秒），被挡住了，参考入点
```

## starTime

说明：

类型：数值。

返回图层的起始时间（以秒为单位）。

![](https://mir.yuelili.com/user/AE/expression/a-z/sourceTime1.png)

## hasVideo

用法：XX.hasVideo

说明：是否还有视频轨

参数：无

类型：布尔值

示例：

```javascript
thisComp.layer("Solid 2").hasAudio; //纯色层是否有视频轨：True
```

## hasAudio

用法：XX.hasAudio

说明：是否还有音轨

参数：无

类型：布尔值

示例：

```javascript
thisComp.layer("Solid 2").hasAudio; //纯色层是否有音频轨：False
```

## active

用法：XX.active

说明：返回图层/属性是否激活的布尔值。如果小眼睛打开,且当前时间点该图层存在于时间轴上，返回 true。对于图层来说，是加强版的 enabled

参数：无参数

类型：布尔值

示例:

```javascript
thisComp.layer("测试").active; //对图层使用。满足条件，返回True
effect("Threshold")("Level").active; //对效果使用。如果有该效果，返回true
```

## enabled

用法：XX.enabled

说明：图层是否打开小眼睛，就是是否显示的意思

参数：无

类型：布尔值

示例：

```javascript
thisComp.layer("Solid 2").enabled; //关上小眼睛：False
```

## audioActive

用法：XX 图层.audioActive

说明：返回某图层是否打开小喇叭，且当前时间点该图层是否存在于时间轴上。都满足则返回 True

参数：无

类型：布尔值

示例：

```javascript
thisComp.layer("示例.mp4").audioActive; // True
```

## sampleImage()

用法：XX 图层.sampleImage(point, radius = [.5, .5], postEffect = true, t = time)

说明：根据图层上某点进行采样，返回 RGBA 颜色值

参数：

point：二维位置数组，表示要进行采样的中心点

radius：二维数组，表示采样区域。[6,6]代表采样 12\*12 的正方形，结果是加权平均

postEffect：要采样的图层是否要计算效果和蒙版。Tue 就是算，加个调色效果就会变。False 就是不算，怎么加都算原始图层的。

t：时间，比如我想采样 3 秒后的图层，t=3

类型：RGBA 四维数组

示例：

```javascript
# 1.变色龙效果
给个颜色复杂的背景图：bg
新建一个圆形：圆1
给圆1的填充颜色增加表达式
sourceLayer = thisComp.layer("bg"); //指定变量sourceLayer为背景层
p = transform.position; //指定变量p为圆的位置值
col = sourceLayer.sampleImage(p)  //其他参数默认就行
//此时移动圆形，会根据背景的颜色变化而变化

# 2.云彩中根据不透明度闪烁的星星
很简单，同上，只不过给不透明度打个表达式
XX.sampleImage(p)[3] //因为表达式返回四维数组，是有不透明度的，所以用[3]调用不透明度即可
```
