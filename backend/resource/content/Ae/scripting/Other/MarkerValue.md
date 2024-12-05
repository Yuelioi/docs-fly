---
title: MarkerValue
order: 4
---

# MarkerValue 标记值

## 标记值

new MarkerValue(comment, chapter, url, frameTarget, cuePointName, params)

描述：标记对象（MarkerValue）表示一个图层或合成的标记。设置标记值、获取值的其他函数，请参考 [Property
属性](https://www.yuelili.com/docs/ae-script/property-property/#62-toc-title)
（因为标记跟关键帧很多都是一样的）

设置 MarkerValue 对象：layerObject.property("Marker").setValueAtTime(time,
markerValueObject);

### 示例 1.新建一个标记

![](https://cdn.yuelili.com/20211012172953.png)

```javascript
// 新建一个标记值
var myMarker = new MarkerValue("我是标记");
// 把值加给一个图层或合成的"Marker"属性组里
app.project.activeItem.layer(1).property("Marker").setValueAtTime(2, myMarker);
```

#### 示例 2：获取/修改标记值

```javascript
var layer = app.project.activeItem.layer(1); // 定义一个图层
var keyComent4 = layer.property("Marker").valueAtTime(4.0, true).comment; // "获取"第4秒标记的评论
var compMarker = new MarkerValue("我是图层标记!"); // 定义评论值
compMarker.label = 9; // 修改标记的属性
var keyComent = layer.property("Marker").setValueAtKey(1, compMarker); // "修改"图层第1个标记的评论
var mkGroup = layer.property("Marker"); // 定义标记属性组
var key4 = mkGroup.keyValue(mkGroup.nearestKeyIndex(4.2)); // 最接近第4秒的标记
```

AE2018 标记设置,AE2019 及之后的标记设置（很多功能没有了）
| ![](https://mir.yuelili.com/2021/07/0b35a735cb997a6be046c7fdd0c7afda.png) |

![](https://cdn.yuelili.com/20211012172447.png)

## 属性

### chapter 标记章节

app.project.item(index).layer(index).property("Marker").keyValue(index).chapter

描述：标记的文本章节链接。章节链接可启动到 QuickTime 电影或其他支持章节标记的格式中的章节的跳转。

::: info 提示
AE2019 以后无了
:::

类型：字符串; 读/写。

### comment 标记注释

app.project.item(index).layer(index).property("Marker").keyValue(index).comment

![](https://mir.yuelili.com/2021/07/e0cf48ac32bf73ec5c21fa299c628a00.png)

描述：标记的文本注释。此注释显示在“时间轴”面板中图层标记旁边。

类型：字符串; 读/写。

### cuePointName 标记提示点

app.project.item(index).layer(index).property("Marker").keyValue(index).cuePointName

描述：Flash Video 提示点名称，如“标记”对话框中所示。

::: info 提示
AE2019 以后无了
:::

类型：字符串; 读/写。

### duration 标记持续时间

app.project.item(index).layer(index).property("Marker").keyValue(index).duration

描述：标记的持续时间(秒)。在“时间轴”面板中显示为延伸的短条。

![](https://mir.yuelili.com/2021/07/ecdbe12134b54e7cd42f303303d6a5da.png)![](https://mir.yuelili.com/2021/07/704398044a329bcb6c81eb44e3cc63f9.png)

类型：浮点; 读/写。

### eventCuePoint 标记事件提示点

app.project.item(index).layer(index).property("Marker").keyValue(index).eventCuePoint

描述：设置为 true 时，FlashVideo 提示点用于事件。否则，用于导航。

::: info 提示
AE2019 以后无了
:::

类型：布尔值 读/写。

### frameTarget 编辑文本帧目标

app.project.item(index).layer(index).property("Marker").keyValue(index).frameTarget

描述：此标记的文本帧目标。与 URL 值一起，它以网页中的特定帧为目标。

::: info 提示
AE2019 以后无了
:::

类型：字符串; 读/写。

### url 标记链接

app.project.item(index).layer(index).property("Marker").keyValue(index).url

描述：该标记的 URL。该 URL 是指向网页的自动链接。

::: info 提示
AE2019 以后无了
:::

类型：字符串; 读/写。

### label 标记标签

app.project.item(index).layer(index).property("Marker").keyValue(index).label

描述：合成或图层标记的标签颜色。颜色由其数字表示(“ 0”代表“无”，或者“
1 至 16”代表“标签”首选项中的一种预设颜色)。自定义标签颜色无法通过编程设置。

在 After Effects 16.0 或更高版本中可用。

类型：整数(0 到 16)；读/写。

### protectedRegion 标记受保护区

app.project.item(index).markerProperty.keyValue(index).protectedRegion

描述：“合成标记”对话框中的“受保护区域”状态。

说明：如果为 true，则合成标记为受保护区域。嵌套合成图层上的受保护区域标记，也将返回 true，否则为 false。图层标记不可用。

在 After Effects 16.0 或更高版本中可用。

类型：布尔值 读/写。

## 方法

### getParameters()

app.project.item(index).layer(index).property("Marker").keyValue(index).getParameters()

描述：返回与该标记值关联的提示点的 Flash Video 提示点参数的键值对。

参数：无。

返回：具有与每个参数名称匹配的属性的对象，其中包含该参数的值。

### setParameters()

app.project.item(index).layer(index).property("Marker").keyValue(index).setParameters(keyValuePairs)

描述：将与 Flash Video 提示点参数关联的一组键值对与此标记值关联的提示点相关联。看 MarkerValue 的标记版本更迭图

说明：一个提示点可以具有任意数量的参数，但是您只能通过用户界面添加三个参数。使用此方法可以添加三个以上的参数。

参数：keyValuePairs 对象，包含键值对应属性和值。toString()将属性和值转为字符串。

返回：无。

示例

```javascript
var mv = new MarkerValue("MyMarker");
var parms = {};
parms.timeToBlink = 1;
parms.assignMe = "A string";
mv.setParameters(parms);
myLayer.property("Marker").setValueAtTime(2, mv);
```
