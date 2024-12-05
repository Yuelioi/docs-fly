---
title: Time 时间对象
order: 3
---

---
title: Time 时间对象
order: 3
category:

- PR脚本

---

## 时间对象

myTime = new Time();

**描述**：表示时间的对象。每秒有 254016000000 个 ticks 。可以用不同的表示形式访问，包括作为时间码字符串。

## 属性

### seconds 秒

myTime.seconds

**描述**：时间值，以秒表示。

**类型**：数字。读/写

```javascript
myTime = new Time();

myTime.seconds = 2; // 设置秒数

alert(myTime.seconds.toString()); // 2（秒）
```

### ticks 刻度

myTime.ticks

**描述**：时间值，以刻度表示。

**类型**：字符串。读写

```javascript
myTime = new Time();

myTime.ticks = "254016000000";

alert(myTime.seconds.toString()); // 1（秒）
```

## 方法

### getFormatted() 时间格式化

myTime.getFormatted(frameRate, displayFormat)

**描述**：将 Time 传递的值作为字符串返回，以指定的显示格式进行格式化。

**参数**

| 参数          | 类型   | 描述                                 |
| ------------- | ------ | ------------------------------------ |
| frameRate     | 字符串 | 要使用的帧速率，基于字符串的时间值。 |
| displayFormat | 整数   | 显示格式：                           |

TIMEDISPLAY_24 时间码 = 100;
TIMEDISPLAY_25 时间码 = 101;

TIMEDISPLAY_2997 丢帧时间码 = 102;
TIMEDISPLAY_2997 不丢帧时间码 = 103;
TIMEDISPLAY_30 时间码 = 104;
TIMEDISPLAY_50 时间码 = 105;
TIMEDISPLAY_5994 丢帧时间码 = 106;
TIMEDISPLAY_5994 不丢帧时间码= 107;
TIMEDISPLAY_60 时间码 = 108;
TIMEDISPLAY_Frames = 109;
TIMEDISPLAY_23976 时间码= 110;
TIMEDISPLAY_16mmFeetFrames = 111;
TIMEDISPLAY_35mmFeetFrames = 112;
TIMEDISPLAY_48 时间码 = 113;
TIMEDISPLAY_AudioSamplesTimecode = 200;
TIMEDISPLAY_AudioMsTimecode = 201;
|

**返回**：字符串。

示例：报错，不知道为啥！

```javascript
myTime = new Time();

myTime.ticks = "254016000000";

myTime.getFormatted("24", 108);
```

### setSecondsAsFraction() 时间除法运算

myTime.setSecondsAsFraction(numerator, denominator)

**描述**：将 Time 对象设置为分子除以分母的结果。

**参数**：分子和分母都是 ints。

**返回**：布尔值；true 如果成功。

```javascript
myTime = new Time(); // 定义新 Time 对象

myTime.setSecondsAsFraction(1, 4);

alert(myTime.seconds.toString()); // 0.25（秒）
```
