---
title: Footage
order: 4
---

# Footage 素材

要将"项目"面板中的素材项目用作表达式中的对象，请使用全局 footage 方法，正如 footage("file_name") 中一样。您还可以使用图层（其源为素材项目）上的 source 属性访问素材对象。  
素材以下表达式方法与合成方法一毛一样，只是对象不一样而已，请参考上一讲"合成属性和方法"  
![](https://mir.yuelili.com/user/AE/expression/exp-13-1.bmp)

## width

用法：XX.width

说明：宽度属性

参数：无

类型：数值

示例：

```javascript
footage("1.mp4").width; //返回1920，当前素材的宽度为1920像素
```

## height

用法：XX.height

说明：高度

参数：无

类型：高度值

示例：

```javascript
footage("1.mp4").height; //返回1080，当前素材的高度为1080像素
```

## duration

用法：XX.duration

说明：持续时间

参数：无

类型：无

示例：

```javascript
footage("1.mp4").duration; //返回当前素材的持续时间。上图素材为10分02秒09帧
```

## frameDuration

用法：合成.frameDuration

说明：单帧持续时间。20 帧/s 时，单帧持续时间=1/20=0.05(秒)

参数：无

类型：数值

示例：

```javascript
footage("1.mp4").frameDuration; //返回素材单帧持续时间，0.03333。当前合成为30帧/s
```

## ntscDropFrame

用法：XXComp.ntscDropFrame

说明：判断合成是否为丢帧格式。关于丢帧请自行百度(29.97 就是丢帧)

参数：无

类型：布尔值

示例：

```javascript
footage("1.mp4").ntscDropFrame; //丢帧，返回True
```

## pixelAspect

说明：像素高宽比

类型：数值。

示例：  
![](https://mir.yuelili.com/user/AE/expression/exp-8-1.bmp)

```javascript
footage("1.mp4").pixelAspect; //0.909
```

## name

用法：XX.name

说明：某个对象或属性的名称

参数：无

类型：一般为字符串

示例：

```javascript
thisComp.layer("Black Solid 1").name; // 类型：Black Solid 1
footage("1.mp4").name; //类型：1.mp4
```

[请结合示例（戳我）](https://www.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

## sourceText

用法：footage("xx").sourceText

说明：返回 JSON 文件本身的内容

类型：字符串

示例：[1.json](https://mir.yuelili.com/user/AE/expression/a-z/1.json)

```javascript
footage("1.json").sourceText

//返回如下
{
"name": "yueli",
"age": 18
}
```

## sourceData

用法：footage("xx").sourceData

说明：返回 JSON 文件本身，作为数组

类型：对象

示例：[1.json(单击查看）](https://mir.yuelili.com/user/AE/expression/a-z/1.json)

案例：[戳我](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

示例：

```javascript
footage("1.json").sourceData; //返回对象本身
footage("1.json").sourceData.name; //返回yueli
footage("1.json").sourceData["age"]; //返回18。
```

## dataValue()

用法：footage("xx").dataValue(dataPath)

说明：返回 JSON 文件的值数据 (json 为键-值对形式。上面文件的 name，age 为键组，而冒号后面的"yueli"," 18"为值组）

参数：dataPath 为数组

类型：字符串

案例：[戳我](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

示例：[1.json](https://mir.yuelili.com/user/AE/expression/a-z/1.json)[(单击查看）](https://mir.yuelili.com/user/AE/expression/a-z/1.json)

```javascript
footage("1.json").dataValue([0]); //返回值组第1个，也就是yueli
footage("1.json").dataValue([1]); //返回值组第2个，也就是18

//如果值组的值还有多重嵌套
footage("XX.json").dataValue([0, 1]); //返回第1个值组下的第一个值
```

## dataKeyCount()

用法：footage("xx").dataKeyCount(dataPath)

说明：返回 JSON 文件样本数量

参数：dataPath 为数组

类型：数值

案例：[戳我](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

示例：[sample.mgjson 见参考样本](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)，以下返回样本个数

```javascript
footage("sample.mgjson").dataKeyCount([0]); //返回第1组的样本计数，52
footage("sample.mgjson").dataKeyCount([1]); //返回第2组的样本计数，52
```

## dataKeyTimes()

用法：footage("xx").dataKeyTimes(dataPath, t0 = startTime, t1=endTime)

说明：返回 .mgJSON 文件中指定动态数据流的采样时间（以秒为单位）。也可指定返回样本的时间跨度。

参数：

dataPath 数组，必选。层次结构中到动态数据流的路径。  
t0 （可选）数字。返回样本的时间跨度的开始时间 (以秒为单位)。默认设置为 startTime。  
t1（可选）数字。返回样本的时间跨度的结束时间 (以秒为单位)。默认设置为 endTime。

类型：表示采样时间的数组。

示例：[sample.mgjson 见参考样本](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

```javascript
footage("sample.mgjson").dataKeyTimes([0], 0.1, 0.2); //返回第1个子项在 0.1 秒和 0.2秒之间的样本时间。25帧的，所以间隔是0.04。[0.08,0.12,0.16,0.2]
```

## dataKeyValues()

用法：footage("xx").dataKeyValues(dataPath, t0 = startTime, t1=endTime)

说明：返回 .mgJSON 文件中指定动态数据流中的样本值。也可指定返回样本的时间跨度。默认情况下，返回动态数据流中 startTime 和 endTime 之间所有样本的时间，如 .mgJSON 文件中数据流的 samplesTemporalExtent 属性所定义。接受单数组值以将层次结构中的路径定义到所需的动态数据流。

参数：  
dataPath 数组，必选。层次结构中到动态数据流的路径。  
t0 （可选）数字。返回样本的时间跨度的开始时间 (以秒为单位)。默认设置为 startTime。  
t1 （可选）数字。返回样本的时间跨度的结束时间 (以秒为单位)。默认设置为 endTime。

类型：表示样本值的数组。

示例：[sample.mgjson 见参考样本](https://mir.yuelili.com/ae/expression/ae-expression/exp-data-force.html)

```javascript
footage("sample.mgjson").dataKeyValues([0], 0.1, 0.2); //返回第1个子项在 0.1秒和0.2秒之间的采样值。[2,3,5,8]
```

Updated on 2021 年 10 月 24 日

[8.Comp 合成](https://www.yuelili.com/docs/ae-expression/exp-comp/)
