---
title: TrackItem 轨道对象
order: 2
---

---
title: TrackItem 轨道对象
order: 3
category:

- PR脚本

---

## TrackItem 对象

app.project.sequences[index].audioTracks[index].clips[index]
app.project.sequences[index].videoTracks[index].clips[index]

**描述**：**trackItem**对象表示[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)中视频或音频轨道上的的项。

## 属性

### components 组件

app.project.sequences[index].audioTracks[index].clips[index].components
app.project.sequences[index].videoTracks[index].clips[index].components

**描述**：与此 trackItem 关联的组件。可以包括过渡以及视频和音频效果。

**类型**：[ComponentCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/componentcollection.html#componentcollection)，只读；

### duration 持续时间

app.project.sequences[index].audioTracks[index].clips[index].duration
app.project.sequences[index].videoTracks[index].clips[index].duration

**描述**：trackItem 的持续时间。

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，只读。

示例：第 1 个视频轨第 1 个片段的持续时间

![](https://cdn.yuelili.com/20211028015150.png)

```javascript
var clip1 = app.project.sequences[0].videoTracks[0].clips[0];

alert(clip1.duration.seconds.toString()); // 5(秒）
```

### end 结束点

app.project.sequences[index].audioTracks[index].clips[index].end
app.project.sequences[index].videoTracks[index].clips[index].end

**描述**：trackItem 的结束时间。注意：这可能与 trackItem 的出点不同。被你裁剪就会变的裁剪出点

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，读/写。

```javascript
var clip1 = app.project.sequences[0].videoTracks[0].clips[0];

alert(clip1.end.seconds.toString()); // 5(秒）
```

### inPoint 入点

app.project.sequences[index].audioTracks[index].clips[index].inPoint
app.project.sequences[index].videoTracks[index].clips[index].inPoint

**描述**：媒体的入点。被你裁剪也不会变的真实入点

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，读/写。

```javascript
var clip1 = app.project.sequences[0].videoTracks[0].clips[0];

alert(clip1.inPoint.seconds.toString()); // 0（秒）
```

### matchName 匹配名称

app.project.sequences[index].audioTracks[index].clips[index].matchName
app.project.sequences[index].videoTracks[index].clips[index].matchName

**描述**：匹配名称

**类型**：字符串; 只读。

```javascript
var clip1 = app.project.sequences[0].videoTracks[0].clips[0];

alert(clip1.matchName); // 空白 不知道啥类型的clip会有匹配名称
```

### mediaType 媒体类型

app.project.sequences[index].audioTracks[index].clips[index].mediaType
app.project.sequences[index].videoTracks[index].clips[index].mediaType

**描述**：此 trackItem 提供的媒体的媒体类型。

**类型**：字符串，**Audio**或**Video**。

```javascript
var clip1 = app.project.sequences[0].audioTracks[0].clips[0];

alert(clip1.mediaType); // Audio（即使是视频里的音频）
```

### name 名称

app.project.sequences[index].audioTracks[index].clips[index].name
app.project.sequences[index].videoTracks[index].clips[index].name

**描述**：轨道项目的名称。

**类型**：字符串; 读/写。

```javascript
var clip1 = app.project.sequences[0].audioTracks[0].clips[0];

alert(clip1.name); // 123.mp4
```

### nodeId 节点 ID

app.project.sequences[index].audioTracks[index].clips[index].nodeId
app.project.sequences[index].videoTracks[index].clips[index].nodeId

**描述**：特有 ID

**类型**：字符串。

```javascript
var clip1 = app.project.sequences[0].audioTracks[0].clips[0];

alert(clip1.nodeId); //  000f4253
```

### outPoint 出点

app.project.sequences[index].audioTracks[index].clips[index].outPoint
app.project.sequences[index].videoTracks[index].clips[index].outPoint

**描述**：媒体的出点，在此 trackItem 中。被你裁剪也不会变的真实出点

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，读/写。

### projectItem 工程项目

app.project.sequences[index].audioTracks[index].clips[index].projectItem
app.project.sequences[index].videoTracks[index].clips[index].projectItem

**描述**：从中绘制媒体的[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)。

**类型**：[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)。

### start 开始时间

app.project.sequences[index].audioTracks[index].clips[index].start
app.project.sequences[index].videoTracks[index].clips[index].start

**描述**：trackItem 的开始时间。注意：可能与 trackItem 的入点不同。被你裁剪就会变的裁剪入点

**类型**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，读/写。

### type 类型

app.project.sequences[index].audioTracks[index].clips[index].type
app.project.sequences[index].videoTracks[index].clips[index].type

**描述**：此 trackItem 提供的媒体类型。

**类型**：数字，**1**表示视频，**2**表示音频。

## 方法

### getSpeed() 获取速度倍率

app.project.sequences[index].audioTracks[index].clips[index].getSpeed()
app.project.sequences[index].videoTracks[index].clips[index].getSpeed()

**描述**：应用于 trackItem 的速度倍率。

**参数**：无。

**返回**：float。默认速度为 1。

```javascript
var clip1 = app.project.sequences[0].videoTracks[0].clips[0];

alert(clip1.getSpeed().toString()); //  2 （200%）
```

### isAdjustmentLayer()

app.project.sequences[index].audioTracks[index].clips[index].isAdjustmentLayer()
app.project.sequences[index].videoTracks[index].clips[index].isAdjustmentLayer()

**描述**：判断 trackItem 是否为调整图层。

**参数**：无。

**返回**：是则返回 true；不是则为 false。

### isSpeedReversed() 速度反转

app.project.sequences[index].audioTracks[index].clips[index].isSpeedReversed()()
app.project.sequences[index].videoTracks[index].clips[index].isSpeedReversed()()

**描述**：是否反转。

**参数**：无。

**返回**：是则返回 1；不是则为 0。

### isSelected() 是否选择

app.project.sequences[index].audioTracks[index].clips[index].isSelected()
app.project.sequences[index].videoTracks[index].clips[index].isSelected()

**描述**：检索 trackItem 的当前选择状态。

**参数**：无。

**返回**：是则返回 true；不是则为 false。

### setSelected() 设置选择

app.project.sequences[index].audioTracks[index].clips[index].setSelected(state, updateUI)
app.project.sequences[index].videoTracks[index].clips[index].setSelected(state, updateUI)

**描述**：设置 trackItem 的选择状态。

**参数**

| 参数     | 类型    | 描述                                               |
| -------- | ------- | -------------------------------------------------- |
| state    | Integer | 如果 1，将选择轨道项目；如果 0，它将被取消选择。   |
| updateUI | Integer | 如果 1，Premiere Pro UI 将在进行此函数调用后更新。 |

**返回**：如果成功则返回**0**。

### getMatchName() 获取匹配名称

app.project.sequences[index].audioTracks[index].clips[index].getMatchName()
app.project.sequences[index].videoTracks[index].clips[index].getMatchName()

**描述**：检索 trackItem 的匹配名称。

**参数**：无。

**返回**：如果成功，**则以字符串形式**返回匹配名称。

### remove() 移除项目

app.project.sequences[index].audioTracks[index].clips[index].remove(inRipple, inAlignToVideo)
app.project.sequences[index].videoTracks[index].clips[index].remove(inRipple, inAlignToVideo)

**描述**：移除项目

**参数**

| 参数           | 类型    | 描述                                                                       |
| -------------- | ------- | -------------------------------------------------------------------------- |
| inRipple       | Boolean | 如果 1，后面的跟踪项目将向前移动，以填补空白；如果 0，后面项目将保持原位。 |
| inAlignToVideo | Boolean | 如果 1，Premiere Pro 会将移动的轨道项目与最近的视频帧的开头对齐。          |

**返回**：如果成功则返回**0**。

### disabled 禁用项目

app.project.sequences[index].audioTracks[index].clips[index].disabled
app.project.sequences[index].videoTracks[index].clips[index].disabled

**描述**：设置 trackItem 的禁用状态。读/写。

**参数**

| 参数            | 类型    | 描述                                                                 |
| --------------- | ------- | -------------------------------------------------------------------- |
| newDisableState | Boolean | 如果 true，则此 trackItem 将被禁用；如果 false，trackItem 将被启用。 |

**返回**：如果成功则返回**0**。

### move() 移动项目

app.project.sequences[index].audioTracks[index].clips[index].move(newInPoint)
app.project.sequences[index].videoTracks[index].clips[index].move(newInPoint)

**描述**：将轨道项目的 inPoint 移动到新时间。

**参数**

| 参数       | 类型 | 描述     |
| ---------- | ---- | -------- |
| newInPoint | Time | 新时间。 |

**返回**：如果成功则返回**0**。
