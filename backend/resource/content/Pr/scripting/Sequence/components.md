---
order: 3
title: components 组件对象
---

---
title: components 组件对象
order: 2
category:

- PR脚本

---

## 组件对象

app.project.sequences[index].audioTracks[index].clips[index].components[index]
app.project.sequences[index].videoTracks[index].clips[index].components[index]

**描述**：**组件**对象表示已有的或者后面添加到 trackItem 的东西。

## 属性

### displayName 显示名称

app.project.sequences[index].audioTracks[index].clips[index].components[index].displayName
app.project.sequences[index].videoTracks[index].clips[index].components[index].displayName

**描述**：组件的名称，本地化显示给用户。

**类型：**字符串; 只读。

示例：查看某一个 clip 的第一个属性名称（不透明度）

![](https://cdn.yuelili.com/20211027144056.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].clips[0].components[0].displayName); // 不透明度
```

### matchName 匹配名称

app.project.sequences[index].audioTracks[index].clips[index].components[index].matchName
app.project.sequences[index].videoTracks[index].clips[index].components[index].matchName

**描述**：组件的名称，因为它是从磁盘加载的；用于唯一标识效果插件。

**类型：**字符串; 只读。

示例：查看某一个 clip 的第一个属性匹配名称（不透明度）

![](https://cdn.yuelili.com/20211027144056.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].clips[0].components[0].matchName); // AE.ADBE Opactiy
```

### properties 组件属性

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties

**描述**：相关组件的属性；通常，这些是效果参数。

**类型：**组件数组，只读；（ComponentParamCollection 对象）。

示例：查看某一个 clip 的属性组里第一个属性名称（不透明度）

![](https://cdn.yuelili.com/20211027144056.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].clips[0].components[0].properties[0].displayName); // 不透明度
```
