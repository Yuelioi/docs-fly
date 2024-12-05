---
title: TrackCollection 对象
order: 6
---

---
title: TrackCollection 对象
order: 8
category:

- PR脚本

---

## TrackCollection 对象

app.project.sequences[index].audioTracks
app.project.sequences[index].videoTracks

TrackCollection 对象表示序列中的[Track 对象](https://ppro-scripting.docsforadobe.dev/sequence/track.html#track)的集合。

父级关系：TrackCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。Collection 的所有方法和属性都可用。

## 属性

### numTracks 轨道总数

app.project.sequences[index].audioTracks.numTracks
app.project.sequences[index].videoTracks.numTracks

**描述**：序列中的轨道总数。

**类型**：整数，只读。

![](https://cdn.yuelili.com/20211028180656.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks.numTracks.toString());
```
