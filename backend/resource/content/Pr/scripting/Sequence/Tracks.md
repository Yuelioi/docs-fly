---
title: Tracks 轨道对象
order: 2
---

---
title: Tracks 轨道对象
order: 5
category:

- PR脚本

---

## 轨道对象

app.project.sequences[index].audioTracks[index]
app.project.sequences[index].videoTracks[index]

**描述**：**轨道**对象（Tracks）表示[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)内视频或音频轨道

## 属性

### clips 剪辑

app.project.sequences[index].audioTracks[index].clips
app.project.sequences[index].videoTracks[index].clips

**描述**：按时间顺序包含在轨道内的[轨道项目](https://ppro-scripting.docsforadobe.dev/item/trackitem.html#trackitem)对象数组。

**类型**：[TrackItemCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/trackitemcollection.html#trackitemcollection)，只读；

![](https://cdn.yuelili.com/20211027180444.png)

### id 轨道序号

app.project.sequences[index].audioTracks[index].id
app.project.sequences[index].videoTracks[index].id

**描述**：这是在创建时分配给轨道的序号。

**类型**：整数，只读。

### mediaType 媒体类型

app.project.sequences[index].audioTracks[index].mediaType
app.project.sequences[index].videoTracks[index].mediaType

**描述**：此轨道中包含的媒体类型。

**类型**：字符串，只读；有效值为 Audio 和 Video。

### name 名称

app.project.sequences[index].audioTracks[index].name
app.project.sequences[index].videoTracks[index].name

**描述**：名称。

**类型**：字符串; 只读。

### transitions 过渡

app.project.sequences[index].audioTracks[index].transitions
app.project.sequences[index].videoTracks[index].transitions

**描述**：按时间顺序包含在轨道中的一系列过渡对象。

**类型**：[TrackItemCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/trackitemcollection.html#trackitemcollection)，只读；

## 方法

### insertClip() 插入剪辑

app.project.sequences[index].audioTracks[index].insertClip(projectItem, time)
app.project.sequences[index].videoTracks[index].insertClip(projectItem, time)

**描述**：在指定时间将“剪辑”（来自[ProjectItem 对象的](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)媒体片段）添加到轨道。届时将插入媒体。

**参数**

| 参数        | 类型                                                                                  | 描述                                   |
| ----------- | ------------------------------------------------------------------------------------- | -------------------------------------- |
| projectItem | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 从中获取媒体的项目项。                 |
| time        | String                                                                                | 添加项目项的时间，以**Ticks 为单位**。 |

**返回**：无。

### isMuted() 检测静音状态

app.project.sequences[index].audioTracks[index].isMuted()
app.project.sequences[index].videoTracks[index].isMuted()

**描述**：检索轨道的当前静音状态。

**参数**：无。

**返回**：如果音轨当前静音，则返回**true**；如果不是，则为**假**。

### overwriteClip() 覆盖剪辑

app.project.sequences[index].audioTracks[index].overwriteClip(projectItem, time)
app.project.sequences[index].videoTracks[index].overwriteClip(projectItem, time)

**描述**：在指定时间将“剪辑”（来自[ProjectItem 对象的](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)媒体片段）添加到轨道。这将覆盖当时的任何现有媒体。

**参数**

| 参数        | 类型                                                                                  | 描述                                   |
| ----------- | ------------------------------------------------------------------------------------- | -------------------------------------- |
| projectItem | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 从中获取媒体的项目项。                 |
| time        | String                                                                                | 添加项目项的时间，以**Ticks 为单位**。 |

**返回**：返回 true。

### setMute() 设置静音

app.project.sequences[index].audioTracks[index].setMute(isMuted)
app.project.sequences[index].videoTracks[index].setMute(isMuted)

**描述**：设置轨道的静音状态。

**参数**

| 参数    | 类型    | 描述                                           |
| ------- | ------- | ---------------------------------------------- |
| isMuted | Integer | 如果 1，则将轨道静音。如果 0，曲目将取消静音。 |

**返回**：如果成功则返回 0。
