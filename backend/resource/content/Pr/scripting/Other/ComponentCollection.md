---
title: ComponentCollection 属性对象集合
order: 1
---

---
title: ComponentCollection 属性对象集合
order: 4
category:

- PR脚本

---

## ComponentCollection 对象

app.project.rootItem.children[index].videoComponents()
app.project.sequences[index].audioTracks[index].clips[index].components
app.project.sequences[index].videoTracks[index].clips[index].components

描述：属性对象集合？

父级关系：ComponentCollection 是 Collection 对象的子类。Collection 的所有方法和属性都可用。

## 属性

### numItems

app.project.rootItem.children[index].videoComponents().numItems
app.project.sequences[index].audioTracks[index].clips[index].components.numItems
app.project.sequences[index].videoTracks[index].clips[index].components.numItems

**描述**：项目个数

**类型**：整数，只读。

示例：查看第一个 clip 的属性个数。（时间重映射是灰的，所以不算。不透明度是第 1 个，可能是英文排序）

![](https://cdn.yuelili.com/20211027143212.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].clips[0].components.numItems.toString()); //2
```
