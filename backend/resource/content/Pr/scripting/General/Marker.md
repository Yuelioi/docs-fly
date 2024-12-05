---
title: Marker 标记对象
order: 3
---

---
title: Marker 标记对象
order: 4
category:

- PR脚本

---

## 标记对象

app.project.activeSequence.markers.getFirstMarker()
app.project.rootItem.children[index].getMarkers().getFirstMarker()

**描述**：[项目](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)和[序列](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)相关的**标志**。

## 属性

### comments 注释

app.project.activeSequence.markers.getFirstMarker().comments
app.project.rootItem.children[index].getMarkers().getFirstMarker().comments

**描述**：标记内的注释。

**类型**：字符串; 读/写。

示例：查看序列第 1 个标记注释

![](https://cdn.yuelili.com/20211028112507.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.markers.getFirstMarker().comments);
```

### end 标记结束

app.project.activeSequence.markers.getFirstMarker().end
app.project.rootItem.children[index].getMarkers().getFirstMarker().end

**描述**：[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，含有所述标记的结束的值。

**类型**：[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)；读/写。

示例：查看序列第 1 个标记的结束时间点

![](https://cdn.yuelili.com/20211028112507.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

var end = sequence.markers.getFirstMarker().end; // 获取结束点时间对象

alert(end.seconds.toString()); // 80（秒） 时间轴的1分20
```

### guid 标记标识符

app.project.activeSequence.markers.getFirstMarker().guid
app.project.rootItem.children[index].getMarkers().getFirstMarker().guid

**描述**：标记的唯一标识符，在实例化时创建。

**类型**：字符串; 只读。

示例：查看序列第 1 个标记的 ID

![](https://cdn.yuelili.com/20211028112507.png)

![](https://cdn.yuelili.com/20211028114337.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

var marker1 = sequence.markers.getFirstMarker(); // 第一个标记

alert(marker1.guid);
```

### name 标记名

app.project.activeSequence.markers.getFirstMarker().name
app.project.rootItem.children[index].getMarkers().getFirstMarker().name

**描述**：标记的名称。

**类型**：字符串; 读/写。

start 标记起始

app.project.activeSequence.markers.getFirstMarker().start
app.project.rootItem.children[index].getMarkers().getFirstMarker().start

**描述**：[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)含有所述标记的开始的值。

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)；读/写。

示例：查看序列第 1 个标记的起始时间点

![](https://cdn.yuelili.com/20211028112507.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

var end = sequence.markers.getFirstMarker().start; // 获取结束点时间对象

alert(end.seconds.toString()); // 60（秒） 时间轴的1分
```

### 标记类型

app.project.activeSequence.markers.getFirstMarker().type

app.project.rootItem.children[index].getMarkers().getFirstMarker().type

**描述**：标记的类型；“评论”、“章节”、“分段”或“WebLink”。注意：Premiere Pro 可以导入一些无法从 Premiere Pro 中创建的标记类型。

**类型**：字符串; 只读。

## 方法

### getColorByIndex() 获取标记颜色索引

app.project.activeSequence.markers.getFirstMarker().getColorByIndex(index)
app.project.rootItem.children[index].getMarkers().getFirstMarker().getColorByIndex(index)

Note!

此功能是在 Adob​​e Premire Pro 13.x 中添加的。

**描述**：获取标记颜色索引。

**参数**

| 参数  | 类型    | 描述                 |
| ----- | ------- | -------------------- |
| index | Integer | 要读取的标记的索引。 |

**返回**：将颜色索引作为 Integer.

### getWebLinkFrameTarget()

app.project.activeSequence.markers.getFirstMarker().getWebLinkFrameTarget()
app.project.rootItem.children[index].getMarkers().getFirstMarker().getWebLinkFrameTarget()

**描述**：从标记的 FrameTarget 字段中检索帧目标。

**参数**：无。

**返回**：返回 String 包含帧目标的 ，如果不成功则返回**0**。

### getWebLinkURL()

app.project.activeSequence.markers.getFirstMarker().getWebLinkURL()
app.project.rootItem.children[index].getMarkers().getFirstMarker().getWebLinkURL()

**描述**：从标记的 URL 字段中检索 URL。

**参数**：无。

**返回**：返回 String 包含 URL 的 ，如果不成功则返回**0**。

### setColorByIndex() 按索引设置标记颜色

app.project.activeSequence.markers.getFirstMarker().setColorByIndex(colorIndex, markerIndex)
app.project.rootItem.children[index].getMarkers().getFirstMarker().setColorByIndex(colorIndex, markerIndex)

Note!

此功能是在 Adob​​e Premire Pro 13.x 中添加的。

**描述**：按索引设置标记颜色。下面列出了颜色索引。

- 0 = 绿色
- 1 = 红色
- 2 = 紫色
- 3 = 橙色
- 4 = 黄色
- 5 = 白色
- 6 = 蓝色
- 7 = 青色

**参数**

| 参数        | 类型    | 描述                   |
| ----------- | ------- | ---------------------- |
| colorIndex  | Integer | 应用于标记的颜色索引。 |
| markerIndex | Integer | 要设置的标记的索引。   |

**返回**：返回 undefined。

### setTypeAsChapter() 设置标记为章节

app.project.activeSequence.markers.getFirstMarker().setTypeAsChapter()
app.project.rootItem.children[index].getMarkers().getFirstMarker().setTypeAsChapter()

**描述**：将标记类型设置为“章节”。

**参数**：无。

**返回**：如果成功则返回**0**。

### setTypeAsComment() 设置标记为注释

app.project.activeSequence.markers.getFirstMarker().setTypeAsComment()
app.project.rootItem.children[index].getMarkers().getFirstMarker().setTypeAsComment()

**描述**：将标记的类型设置为“注释”。

**参数**：无。

**返回**：如果成功则返回**0**。

### setTypeAsSegmentation() 设置标记为分段

app.project.activeSequence.markers.getFirstMarker().setTypeAsSegmentation()
app.project.rootItem.children[index].getMarkers().getFirstMarker().setTypeAsSegmentation()

**描述**：将标记的类型设置为“分段”。

**参数**：无。

**返回**：如果成功则返回**0**。

### setTypeAsWebLink() 设置标记为 Link

app.project.activeSequence.markers.getFirstMarker().setTypeAsWebLink()
app.project.rootItem.children[index].getMarkers().getFirstMarker().setTypeAsWebLink()

**描述**：将标记的类型设置为“WebLink”。

**参数**：无。

**返回**：如果成功则返回**0**。
