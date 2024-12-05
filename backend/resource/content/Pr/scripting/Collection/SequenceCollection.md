---
title: SequenceCollection 序列集合对象
order: 5
---

---
title: SequenceCollection 序列集合对象
order: 7
category:

- PR脚本

---

## 序列集合对象

app.project.sequences

SequenceCollection 对象表示所有的集合 [序列对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)的活动项目。

SequenceCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 SequenceCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numSequences 序列总数

app.project.sequences.numSequences

**描述**：活动项目中的序列总数。

**类型**：整数，只读。

![](https://cdn.yuelili.com/20211028180245.png)

```javascript
alert(app.project.sequences.numSequences.toString()); // 当前项目有2序列
```
