---
title: Time-Conversion
order: 7
---

# Time Conversion

::: danger

如果您需要对素材中的时间码外观进行更多控制，请使用 timeToCurrentFormat 方法或其他 timeTo 方法生成时间码，而非使用时间码或编号效果。创建文本图层，向源文本属性添加表达式，并在表达式字段中输入 timeToCurrentFormat()。  
使用此方法，您可以对时间码文本进行格式设置以及动画制作。此外，时间码使用当前的项目设置定义的同一显示样式。
:::

## isDuration 参数

如果为 true，则时间为绝对时间（从第 0 帧开始计算）。如果为 false（默认值），则从项目开始帧开始计算时间（项目开始帧可以在项目设置-Frame Count 处进行设置）。于此同时，时间的绝对值会向小于 0 的方向取值

示例 1：时间为正，起始帧为 0 帧

```javascript
timeToFrames(1, 29.97, (isDuration = true)); //返回30，因为timeToFrames取整数，29.97算30
timeToFrames(1, 29.97, (isDuration = false)); //返回29，因为timeToFrames取整数，29.97 > 0，需要向接近0的方向取整，所以为29
```

示例 2：时间为负，起始帧为 1 帧

```javascript
timeToFrames(-1, 29.97, (isDuration = true)); //返回-30，因为timeToFrames取整数，-29.97算-30
timeToFrames(-1, 29.97, (isDuration = false)); //返回-29，因为timeToFrames取整数，29.97 < 0，需要向接近0的方向取整，所以为-29
```

![](https://mir.yuelili.com/user/AE/expression/a-z/isDuration1.png)

![](https://mir.yuelili.com/user/AE/expression/a-z/isDuration2.png)

## timeToFrames()

用法：timeToFrames(t= time + thisComp.displayStartTime, fps = 1.0 / thisComp.frameDuration,isDuration = false)

说明：将默认为当前合成时间的 t 转换为**整数**帧数。得出的值为括号里面的两个数(t 和 fps)相乘【时间\*帧速率=帧数】

参数：

- t ，数值，默认值为 time + thisComp.displayStartTime，表示当前时间+合成起始时间
- fps：每秒的帧数，也就是合成的帧速率(也就是 1 秒除以当前合成单帧持续时间)
- frameDuration: 单帧持续时间
- isDuration ，false 则采取丢帧方案，true 则使用正常帧。

类型：数值

示例：下图表示在 3s 时间点，返回本合成（29.97fps）位于 3s 的帧数：90 帧

![](https://mir.yuelili.com/user/AE/expression/exp-global3.png?imageView2/1/w/2332/h/526#)

## framesToTime()

用法：framesToTime(frames, fps = 1.0 /thisComp.frameDuration)

说明：把帧数转为时间

参数：

- frames 为要转换的帧数。
- fps 为每秒的帧数，默认为当前合成的帧速率【1 秒/单帧持续时间=帧速率】

类型：与 frames 参数对应的时间值。不一定是整数。【帧数/帧速率=时间】

示例：本合成（30fps）在 90 帧（frames=90）返回的秒数，即 3s

![](https://mir.yuelili.com/user/AE/expression/exp-global4.png?imageView2/1/w/1926/h/284#)

```javascript
framesToTime(90, 30);
```

## timeToTimecode()

用法：timeToTimecode(t= time + thisComp.displayStartTime, timecodeBase = 30, isDuration =false)

说明：将 t 的值转换为表示时间码的字符串。

参数：

- t 用于转换的值（时间）

- displayStartTime: 合成起始时间

- timecodeBase： 用于设置毫秒的显示方式，默认为 30，代表 1 秒显示 30 帧。最大为 999。建议还是根据合成的帧率进行设置，即 1.0 /thisComp.frameDuration

- isDuration ，false 则采取丢帧方案，true 则使用正常帧。[戳我查看详细原理](https://mir.yuelili.com/docs/expression_ae/expression-global#isduration%e5%8f%82%e6%95%b0)

类型：字符串

示例 1：timecodeBase 参数说明

![](https://mir.yuelili.com/user/AE/expression/a-z/timecodeBase.png)

```javascript
timeToTimecode(t= time + thisComp.displayStartTime, timecodeBase = 100, isDuration =false
```

20 帧的 3:01 在此刻显示为 3:05，不难理解，timecodeBase 为 100 时，原先一秒显示 20 帧，现在显示 100 帧。原先 01 帧，现在就显示 05 帧咯。

再举个例子：1 分钟=60 秒，相当于 60 的帧率，如果安到百进制上（100 帧率），那原来的 30 秒，相当于百进制的 50 秒咯。

示例 2：下图表示合成为 30fps 的 3:18 分的时间码

![](https://mir.yuelili.com/user/AE/expression/exp-global5.png?imageView2/1/w/2026/h/434#)

```javascript
timeToTimecode((t = time + thisComp.displayStartTime), (timecodeBase = 30), (isDuration = false)); //见上
timeToTimecode(3, 30, false); //简略写法：返回00:00:03:00
```

## timeToNTSCTimecode()

用法：timeToNTSCTimecode(t= time + thisComp.displayStartTime, ntscDropFrame = false, isDuration =false)

说明：将 t 转换为表示 NTSC 时间码的字符串

参数：

- 1： t ，数值类型，默认值 time + thisComp.displayStartTime 为当前时间

- 2：ntscDropFrame 为 False（默认值），用于丢帧与否。

- 3：isDuration ，false 则采取丢帧方案，true 则使用正常帧。[戳我查看详细原理](https://mir.yuelili.com/docs/expression_ae/expression-global#isduration%e5%8f%82%e6%95%b0)

- 4：有关 NTSC 制请自行百度

类型：字符串

示例：下图表示 30fps 的 3:18 分的时间码，由于选择未丢帧，所以返回 00:00:03:17

![](https://mir.yuelili.com/user/AE/expression/exp-global6.png?imageView2/1/w/2050/h/452#)

```javascript
timeToNTSCTimecode((t = time + thisComp.displayStartTime), (ntscDropFrame = false), (isDuration = false)); //见上
timeToNTSCTimecode(3.6, (ntscDropFrame = false), (isDuration = false)); //可以直接用3.6表示3:18 (30帧/s*0.6秒=18帧)
```

## timeToFeetAndFrames()

用法：timeToFeetAndFrames(t= time + thisComp.displayStartTime, fps = 1.0 / thisComp.frameDuration,framesPerFoot = 16, isDuration = false)

说明：将 t 的值转换为表示胶片和帧英尺数的字符串。

参数：

- t ，数值类型，默认值 time + thisComp.displayStartTime 为当前时间

- displayStartTime: 合成起始时间

- fps：设定每秒的帧数，默认为当前合成的帧速率(也就是 1.0/thisComp.frameDuration)

- frameDuration: 单帧持续时间

- framesPerFoot 参数指定一英尺胶片中的帧数。它默认为 16，是 35 毫米素材的最常见速率。

- isDuration ，false 则采取丢帧方案，true 则使用正常帧。[戳我查看详细原理](https://mir.yuelili.com/docs/expression_ae/expression-global#isduration%e5%8f%82%e6%95%b0)。

类型：字符串

示例：下图表示 30fps 的 3:18 分的胶片和帧英尺数

![](https://mir.yuelili.com/user/AE/expression/exp-global7.png?imageView2/1/w/1964/h/432#)

```javascript
timeToFeetAndFrames((t = time + thisComp.displayStartTime), (fps = 1.0 / thisComp.frameDuration), (framesPerFoot = 16), (isDuration = false)); //全都用表达式表示
timeToFeetAndFrames(3.6, 30, 16, 0); //用具体数值表示
```

## timeToCurrentFormat()

用法：timeToCurrentFormat(t= time + thisComp.displayStartTime, fps = 1.0 / thisComp.frameDuration,isDuration = false, ntscDropFrame =thisComp.ntscDropFrame)

说明：将 t 的值转换为表示采用当前项目设置显示格式的时间的字符串。项目设置显示格式：就是左上角那个时间码，可以设置为帧数显示,或者时间码显示（按住 ctrl 鼠标单击切换）

参数

- t ，数值类型，默认值 time + thisComp.displayStartTime 为当前时间

- displayStartTime: 合成起始时间

- fps：设定每秒的帧数，默认为当前合成的帧速率(也就是 1.0/thisComp.frameDuration)

- frameDuration: 单帧持续时间

- isDuration ，false 则采取丢帧方案，true 则使用正常帧。[戳我查看详细原理](https://mir.yuelili.com/docs/expression_ae/expression-global#isduration%e5%8f%82%e6%95%b0)
- 默认值：ntscDropFrame = thisComp.ntscDropFrame(当前合成有无丢帧,比如 29.97)

类型：字符串

示例 1：下图表示当前时间格式（时间码显示）

![](https://mir.yuelili.com/user/AE/expression/exp-global8.png?imageView2/1/w/2024/h/380#)

示例 2：下图表示当前时间格式（帧显示）

![](https://mir.yuelili.com/user/AE/expression/exp-global12.bmp?imageView2/1/w/2364/h/438#)
