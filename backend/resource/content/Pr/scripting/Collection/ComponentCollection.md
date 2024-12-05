---
title: ComponentCollection 对象
order: 1
---

---
title: ComponentCollection 对象
order: 3
category:

- PR脚本

---

## ComponentCollection 对象

app.project.rootItem.children[index].videoComponents()
app.project.sequences[index].audioTracks[index].clips[index].components
app.project.sequences[index].videoTracks[index].clips[index].components

_增加一个说明_

ComponentCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 ComponentCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numItems 数量

app.project.rootItem.children[index].videoComponents().numItems
app.project.sequences[index].audioTracks[index].clips[index].components.numItems
app.project.sequences[index].videoTracks[index].clips[index].components.numItems

**描述**：_属性的数量_

**类型**：整数，只读。

![](https://cdn.yuelili.com/20211028181806.png)

示例：当前序列 第 1 个视频轨 第 1 个剪辑片段的属性个数

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].clips[0].components.numItems.toString()); // 属性的数量
```
