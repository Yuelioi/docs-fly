---
order: 2
title: Marker 标记集合对象
---

---
title: Marker 标记集合对象
order: 4
category:

- PR脚本

---

## 标记集合对象

app.project.sequences[index].markers
app.project.rootItem.children[index].getMarkers()

MarkerCollection 对象表示在[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)和[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)的[标记集合对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)。

MarkerCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 MarkerCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numMarkers 标记数量

app.project.sequences[index].markers.numMarkers
app.project.rootItem.children[index].getMarkers().numMarkers

**描述**：项目项或序列中标记对象的计数。

**类型**：整数，只读。

## 方法

### createMarker() 创建标记

app.project.sequences[index].markers.createMarker(time)
app.project.rootItem.children[index].getMarkers().createMarker(time)

**描述**：在项目项或序列上创建新的[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)。

**参数**

| 参数 | 类型  | 描述                             |
| ---- | ----- | -------------------------------- |
| time | Float | 应创建标记的时间（以秒为单位）。 |

**返回**：如果成功则[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)。

### deleteMarker() 删除标记

app.project.sequences[index].markers.deleteMarker(marker)
app.project.rootItem.children[index].getMarkers().deleteMarker(marker)

**描述**：从集合中删除给定的标记对象。

**参数**

| 参数   | 类型                                                                           | 描述                       |
| ------ | ------------------------------------------------------------------------------ | -------------------------- |
| marker | [标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker) | 要从集合中删除的标记对象。 |

**返回**：Boolean.

**示例**：从活动序列中删除所有标记

```javascript
var markers = app.project.activeSequence.markers;

var marker = markers.getFirstMarker();

var count = markers.numMarkers;

while (marker) {
  markers.deleteMarker(marker);

  marker = markers.getFirstMarker();
}

alert("Removed " + count.toString() + " markers");
```

### getFirstMarker() 获取第 1 个标记

app.project.sequences[index].markers.getFirstMarker()
app.project.rootItem.children[index].getMarkers().getFirstMarker()

**描述**：在给定的项目项或序列上检索第一个标记对象，按时间排序（以秒为单位）。

**参数**：无。

**返回**：[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)或 undefined.

### getLastMarker() 获取最后 1 个标记

app.project.sequences[index].markers.getLastMarker()
app.project.rootItem.children[index].getMarkers().getLastMarker()

**描述**：检索给定项目项或序列上的最后一个标记对象，按时间排序（以秒为单位）。

**参数**：无。

**返回**：[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)或 undefined.

### getNextMarker() 获取下 1 个标记

app.project.sequences[index].markers.getNextMarker(currentMarker)
app.project.rootItem.children[index].getMarkers().getNextMarker(currentMarker)

**描述**：获取下一个可用标记，按秒排序，从给定的标记开始。

**参数**

| 参数          | 类型                                                                           | 描述                               |
| ------------- | ------------------------------------------------------------------------------ | ---------------------------------- |
| currentMarker | [标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker) | 一个起始标记对象，从中获取下一个。 |

**返回**：[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)或 undefined.

### getPrevMarker() 获取上 1 个标记

app.project.sequences[index].markers.getPrevMarker(currentMarker)
app.project.rootItem.children[index].getMarkers().getPrevMarker(currentMarker)

**描述**：获取前一个可用标记，按秒排序，从给定的标记开始。

**参数**

| 参数          | 类型                                                                           | 描述                               |
| ------------- | ------------------------------------------------------------------------------ | ---------------------------------- |
| currentMarker | [标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker) | 一个起始标记对象，从中获取前一个。 |

**返回**：[标记对象](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)或 undefined.
