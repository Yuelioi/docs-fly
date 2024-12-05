---
title: Property
order: 12
---

# Property 属性

## value

用法：XX.value

说明：返回某属性在当前时间的值

参数：无

类型：根据属性

示例：

```javascript
thisComp.layer("Solid 2").transform.position.value; //返回位置值[960,540]
```

## valueAtTime(t)

用法：valueAtTime(t=time)

参数：t 是数值。

说明：属性在指定时间（以秒为单位）的值。

类型：数值或数组。

例如，要从一组（四个）值中随机选择每个帧的属性值，请在 0、1、2 和 3 秒将您的四个值设为关键帧，然后将以下表达式应用于该属性：

示例：

```javascript
thisComp.layer("Solid 2").transform.position.valueAtTime(2); //返回位置第2秒的位置值
thisComp.layer("Solid 2").transform.position.valueAtTime(time - 1); //位置的延迟动画
```

## velocity

说明：某属性的速度值

类型：数值或数组。返回当前时间的临时速度值。对于空间属性（例如位置），它返回正切矢量值。结果与属性的维度相同。

示例：方块从 0 秒到 2 秒，位置由[1000,1000]移动到[800,700]

```javascript
thisComp.layer(1).transform.position.velocity; //返回[-100,-150]
```

## velocityAtTime(t)

说明：某属性的速度值

类型：数值或数组。返回指定时间的临时速度值。

示例：方块从 0 秒到 2 秒，位置由[1000,1000]移动到[800,700]

```javascript
thisComp.layer(1).transform.position.velocityAtTime(1); //返回[-100,-150]
```

## speed

说明：返回 1 维的的正速度值。如果对一维属性使用，那速度可为负

类型：1 维数值。

示例：方块从 0 秒到 2 秒，位置由[1000,1000]移动到[700,600]

```javascript
thisComp.layer(1).transform.position.speed; //返回250。勾股定理会不会？x移动300，y移动400，空间移动500,2秒 >>> 250/s
```

## speedAtTime(t)

说明：返回指定时间的 1 维的正速度值。如果对一维属性使用，那速度可为负

类型：1 维数值。

示例：方块从 0 秒到 2 秒，位置由[1000,1000]移动到[700,600]

```javascript
thisComp.layer(1).transform.position.speedAtTime(1); //返回250。勾股定理啦？x移动300，y移动400，空间移动500,2秒 >>> 250/s
```

## wiggle()

用法：wiggle(freq,amp,octaves=1,amp_mult=.5,t=time)

说明：随机摇动（摆动）属性值。

类型：数值或数组。Wiggle 放入某个属性中，将返回与属性具有相同类型的值。

- 放入旋转属性，将返回数值；
- 放入 2D 属性，如 2D 位置，将返回二维数组值；
- 放入 3D 属性，如 3D 位置，返回三维数组值，
- 放入 4D 属性，如颜色，返回四维数组值，
- 以此类推。

参数类型：freq、amp、octaves、amp_mult 和 t 是数值。

- Freq：频率。每秒随机变化的次数。数值越高，变化的次数就越多。
- amp: 振幅。 随机波动的大小。例如，旋转属性的参数值为 100，振幅输入 50，它将在 50 到 150 之间随机。
- octaves : 八度音阶，默认为 1。像其他随机方法一样，wiggle()基于噪波进行计算，就像分形噪波一样。如果该参数变大，则噪波更细致，也就是频率之间会更细致。
- amp_mult：默认为 0.5。每个八度的振幅乘以 0 到 1 的次数。像八度音阶一样，越接近 0，细节越少；越接近 1，细节越多。
- t：默认情况下这是另一个参数，定义此方法根据哪个时间点进行随机。  
  如果使用 time，将使用当前时间进行随机；如果输入 5，则使用属性的第 5 秒来返回随机值，此时相当于一个随机的"固定"值。参考随机数种子。

![【AE表达式】wiggle构成四步图解_月离的万事屋](https://cdn.yuelili.com/202121172246-O.png)

示例 1：position 属性上使用

```javascript
wiggle(5, 20, 3, 0.5); //每秒产生 约5 次摆动，大小约 20 像素。除了主要摆动之外，其他两个级别的详细摆动发生的频率为每秒 10 次和 20 次摆动，各自的大小为 10 和 5 像素。
```

示例 2：对于二维属性（例如缩放），按相同值摆动两个维度：

```javascript
v = wiggle(5, 10);
[v[0], v[0]];
```

示例 3：二维属性，只能在 Y 轴上摇摆：

```javascript
freq = 3;
amp = 50;
w = wiggle(freq, amp);
[value[0], w[1]];
```

## temporalWiggle()

用法：temporalWiggle(freq,amp,octaves=1,amp_mult=.5,t=time)

说明：在摆动的时间对属性进行采样。此函数采样的是时间，而 wiggle 摆动的是值。其他跟 wiggle 一样

参数：freq、amp、octaves、amp_mult 和 t 是数值。详见 wiggle()

类型：数值或数组。

示例：给位置加个关键帧，0 秒[0,0] 4 秒[200,200]

```javascript
temporalWiggle(1, 2); //1秒波动一次，一次波动范围在2秒之内，在2秒时，有可能波动2秒，位置直接摇摆到[200,200]
```

## smooth()

用法：smooth(width=.2,samples=5,t=time)

说明：随着时间的推移平滑属性值，将值的大而短的偏差转换为更小、分布更均匀的偏差。

参数：width、samples 和 t 是数值。此平滑通过在指定时间向属性值应用框滤镜来完成。

width 。指定以秒为单位的当前帧在平滑计算中平均的时间量。这个不太理解

samples 。均匀分布的关键帧数目；想要更大的平滑度应使用更大的值（性能会下降）。选择一个奇数会把当前帧也包括在计算中

time 为采样时间，默认为当前时间

类型：数值或数组。

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/smooth.png)

```javascript
smooth(0.5, 5); //采样平均时间维0.5s。红点是当前关键帧，蓝色为sample个数，如果是奇数，则算上红色。
```

## loopIn()和 loopOut()

用法：loopIn(type="cycle",numKeyframes=0)

返回类型：数值或数组。

循环类型

cycle（默认）重复指定段。

pingpong 重复指定段，向前和向后交替。

offset 重复指定段，但会按段开始和结束时属性值的差异乘以段已循环的次数偏移每个周期。

continue 不重复指定段，但继续基于第一个或最后一个关键帧的速度对属性进行动画制作。此类型不接受 keyframes 或 duration 参数。

cycle

循环

![](https://cdn.yuelili.com/20220118130053.gif)

重复 K1~K3 的动画

pingpong

乒乓

![](https://cdn.yuelili.com/20220118130059.gif)

K1 到 K3 后，再从 K3 回到 K1，如此往复

offset

偏移

![](https://cdn.yuelili.com/20220118130105.gif)

K1 到 K3 后，以 K3 作为新的“K1”点（起始点），重复这段动画

continue

持续

![](https://cdn.yuelili.com/20220118130112.gif)

K1 到 K3 后，根据 K3 的方向与速度。直接按直线跑了

更多信息，请参阅  [loop 表达式详解](https://www.yuelili.com/?p=10578)。

## loopInDuration()和 loopOutDuration()

用法：loopInDuration(type="cycle",duration=0) 和 loopOutDuration(type="cycle",duration=0)

说明：跟 loopIn 一样，只不过不是按关键帧计算，而是以秒来计算循环

更多信息，请参阅  [loop 表达式详解](https://www.yuelili.com/?p=10578)。

## key()

用法：key(index)或 key(markerName)

说明：前者用于调用关键帧或标记，后者只能用于标记。一般要与其他属性一起使用

参数：index 为数值，markerName 为字符串。注意，标记名可以在标记设置-评论，进行改变。不过 AE2020 最新版本好像不可以用标记名调用了

类型：无

示例：

```javascript
thisLayer.marker.key(1).time; //返回图层第1个标记的时间
transform.key(2).value; //如返回[960,540]，位置属性第2个的关键帧值(如果没有关键帧会报错)
marker.key("标记1").time; //如返回15，返回图层标记1的时间
```

## nearestKey(t)

类型：Key 或 MarkerKey。一般要配合其他属性使用

说明：返回最接近指定时间的 Key 或 MarkerKey 对象。

示例：如果 3 秒 5 秒都有一个关键帧

```javascript
nearestKey(4).time; //返回最接近的关键帧的时间，也就是3秒
nearestKey(4.1).time; //返回最接近的关键帧的时间，也就是5秒
```

## numKeys

用法：XX 关键帧.numKeys

说明：关键帧数目，或者标记数目。

类型：数值。

示例：

```javascript
transform.position.numKeys; //返回本图层位置属性的关键帧个数
thisLayer.marker.numKeys; //返回本图层标记个数
```

## active

说明：是否激活属性

类型：布尔值

示例：

```javascript
effect("Threshold").active; //如果效果前的FX打开，且当前效果作用于时间轴上，则返回true
```

## enabled

说明：是否开启属性

类型：布尔值

示例：

```javascript
effect("Threshold").enabled; //如果效果前的FX打开，则返回true
```

## propertyGroup()

用法：propertyGroup(countUp=1)

说明：向上级进行索引，返回包含该属性的上一级

参数：countUp 为数值，默认为 1

类型：组。

示例：

![](https://mir.yuelili.com/user/AE/expression/a-z/propertyGroup.png)

```javascript
r = thisComp.layer("形状图层 1").content("椭圆 1").transform.rotation; //定义变量r，为椭圆的一个旋转属性

r.propertyGroup(1).name; //返回 变换。旋转的上1级是变换
r.propertyGroup(2).name; //返回 椭圆1。旋转的上2级是椭圆1
```

## propertyIndex

用法：XX 属性.propertyIndex

说明：属性在属性组的索引，属性组一般在蒙版、效果、文本动画、选择器、形状、跟踪器以及跟踪点。

类型：数值。

示例：

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Radial_Wipe.png)

```javascript
effect("Radial Wipe")("Start Angle").propertyIndex; //starAngle在Radial Wipe效果的第2个，故返回2
```

## name

说明：属性或属性组的名称。

类型：字符串。

示例：

```javascript
effect("Threshold")("Level").name; //返回Level
```
