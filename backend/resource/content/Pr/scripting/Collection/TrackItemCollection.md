---
title: TrackItemCollection 对象
order: 7
---

---
title: TrackItemCollection 对象
order: 9
category:

- PR脚本

---

## TrackItemCollection 对象

app.project.sequences[index].audioTracks[index].clips
app.project.sequences[index].videoTracks[index].clips

该 TrackItemCollection 对象代表的集合[TrackItem 对象](https://ppro-scripting.docsforadobe.dev/item/trackitem.html#trackitem)的轨道上。

TrackItemCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 TrackItemCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numItems

app.project.sequences[index].audioTracks[index].clips.numItems
app.project.sequences[index].videoTracks[index].clips.numItems

**描述**：轨道上的剪辑总数。

**类型**：整数，只读。
