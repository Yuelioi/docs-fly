---
title: ComponentParam 组件参数对象
order: 1
---

---
title: ComponentParam 组件参数对象
order: 3
category:

- PR脚本

---

## ComponentParam 组件参数对象

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index]
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index]

**描述**：组件**参数**对象表示与组件相关联的应用到[TrackItem 对象](https://ppro-scripting.docsforadobe.dev/item/trackitem.html#trackitem)的参数。

## 属性

### displayName 显示名称

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].displayName
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].displayName

**描述**：组件参数的名称，本地化显示给用户。

**类型**：字符串; 只读。

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

var first_clip = sequence.videoTracks[0].clips[0]; // 视频轨第1个片段

alert(first_clip.components[0].properties[0].displayName); // 不透明度
```

## 方法

### addKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].addKey(time)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].addKey(time)

**描述**：在指定时间将关键帧添加到组件参数流。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数 | 类型                                                                     | 描述               |
| ---- | ------------------------------------------------------------------------ | ------------------ |
| time | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 何时应添加关键帧。 |

**返回**：如果成功则返回**0**。

### areKeyframesSupported()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].areKeyframesSupported()
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].areKeyframesSupported()

**描述**：检索此组件参数是否支持关键帧。

**参数**：无。

**返回**：true 如果选择了 trackItem 则返回；false 如果不。

### findNearestKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].findNearestKey(timeToCheck, threshold)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].findNearestKey(timeToCheck, threshold)

**描述**：设置组件参数是否随时间变化。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数        | 类型 | 描述                                       |
| ----------- | ---- | ------------------------------------------ |
| timeToCheck |      | 从给定时间开始搜索                         |
| threshold   |      | 在任一方向上的时间距离，以**刻度为单位**。 |

**返回**：返回**时间**值，指示最接近的关键帧的时间。

### findNextKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].findNextKey(timeToCheck)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].findNextKey(timeToCheck)

**描述**：在提供的 之后暂时返回关键帧 timeToCheck。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数        | 类型 | 描述                 |
| ----------- | ---- | -------------------- |
| timeToCheck |      | 从给定时间开始搜索。 |

**返回**：返回**时间**值，指示最接近的关键帧的**时间**，如果没有可用的后续关键帧，则返回**0**。

### findPreviousKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].findPreviousKey(timeToCheck)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].findPreviousKey(timeToCheck)

**描述**：返回时间上提供的关键帧 timeToCheck。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数        | 类型 | 描述                 |
| ----------- | ---- | -------------------- |
| timeToCheck |      | 从给定时间开始搜索。 |

**返回**：返回**时间**值，指示最近的关键帧的**时间**，如果没有可用的前一个关键帧，则返回**0**。

### getColorValue()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].getColorValue()
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].getColorValue()

**描述**：获取组件参数流的值。注意：这仅适用于非时变的参数。

**参数**：无。

**返回**：返回包含在组件参数流中找到的值的**颜色**，如果不成功则返回**0**。

### getKeys()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].getKeys()
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].getKeys()

**描述**：返回 timeToCheck 组件参数上所有关键帧的数组。注意：这只能在支持关键帧的参数上设置。

**参数**：无。

**返回**：返回一个**数组**的**时间**值，表明发生在什么时候每个关键帧，或者**0**，如果没有关键帧可用。

### getValue()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].getValue()
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].getValue()

**描述**：获取组件参数流的值。注意：这仅适用于非时变的参数。

**参数**：无。

**返回**：返回组件参数流的值；回报因流类型而异。

### getValueAtKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].getValueAtKey(time)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].getValueAtKey(time)

**描述**：在指定的关键帧时间检索组件参数流的值。注意：只能与关键帧参数流一起使用。

**参数**

| 参数 | 类型                                                                     | 描述                       |
| ---- | ------------------------------------------------------------------------ | -------------------------- |
| time | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 应从中检索关键帧值的时间。 |

**返回**：返回组件参数流的值在 time，或**0**如果不成功。

### getValueAtTime()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].getValueAtTime(time)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].getValueAtTime(time)

**描述**：在指定时间检索组件参数流的值。如果该值介于两个关键帧之间，则会进行插值。

**参数**

| 参数 | 类型                                                                     | 描述                       |
| ---- | ------------------------------------------------------------------------ | -------------------------- |
| time | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 应从中检索关键帧值的时间。 |

**返回**：返回组件参数流的值在 time，或**0**如果不成功。

### isTimeVarying()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].isTimeVarying()
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].isTimeVarying()

**描述**：检索组件参数是否随时间变化。

**参数**：无。

**返回**：true 如果参数随时间变化，则返回；false 如果不。

### removeKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].removeKey(time)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].removeKey(time)

**描述**：在指定时间移除组件参数流上的关键帧。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数 | 类型                                                                     | 描述                           |
| ---- | ------------------------------------------------------------------------ | ------------------------------ |
| time | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 时间值，指示何时应删除关键帧。 |

**返回**：如果成功则返回**0**。

### removeKeyRange()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].removeKeyRange(startTime, endTime)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].removeKeyRange(startTime, endTime)

**描述**：在指定时间之间从组件参数流中删除所有关键帧。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数      | 类型                                                                     | 描述                             |
| --------- | ------------------------------------------------------------------------ | -------------------------------- |
| startTime | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 在什么时间（含）开始删除关键帧。 |
| endTime   | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 在什么时候结束删除关键帧。       |

**返回**：如果成功则返回**0**。

### setColorValue()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].setColorValue(alpha, red, green, blue, updateUI)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].setColorValue(alpha, red, green, blue, updateUI)

**描述**：设置组件参数流中的值，表示颜色。

**参数**

| 参数     | 类型    | 描述                      |
| -------- | ------- | ------------------------- |
| alpha    | Integer | 阿尔法值。                |
| red      | Integer | 红色值。                  |
| green    | Integer | 绿色价值。                |
| blue     | Integer | 蓝色值。                  |
| updateUI | Integer | 更新流的值后强制更新 UI。 |

**返回**：如果成功则返回**0**。

### setInterpolationTypeAtKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].setInterpolationTypeAtKey(time, interpretationType)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].setInterpolationTypeAtKey(time, interpretationType)

**描述**：指定在指定时间分配给关键帧的插值类型。注意：只能与关键帧参数流一起使用。

**参数**

| 参数               | 类型                                                                     | 描述                                     |
| ------------------ | ------------------------------------------------------------------------ | ---------------------------------------- |
| time               | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 要修改的关键帧时间。                     |
| interpretationType | type                                                                     | 必须是以下之一：\* 0 kfInterpMode_Linear |

- 1 kfInterpMode_EaseIn_Obsolete
- 2 kfInterpMode_EaseOut_Obsolete
- 3 kfInterpMode_EaseInEaseOut_Obsolete
- 4 kfInterpMode_Hold
- 5 kfInterpMode_Bezier
- 6 kfInterpMode_Time
- 7 kfInterpMode_TimeTransitionStart
- 8 kfInterpMode_TimeTransitionEnd
  |
  | updateUI | boolean | 之后是否更新 UI。 |

**返回**：如果成功则返回**0**。

### setTimeVarying()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].setTimeVarying(varying)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].setTimeVarying(varying)

**描述**：设置组件参数是否随时间变化。注意：这只能在支持关键帧的参数上设置。

**参数**

| 参数    | 类型    | 描述                                                  |
| ------- | ------- | ----------------------------------------------------- |
| varying | Boolean | 如果 true，组件参数将随时间变化；如果 false，它不会。 |

**返回**：如果成功则返回**0**。

### setValue()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].setValue(value, updateUI)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].setValue(value, updateUI)

**描述**：获取组件参数流的值。注意：这仅适用于非时变的参数。

**参数**

| 参数     | 类型    | 描述                                                |
| -------- | ------- | --------------------------------------------------- |
| value    |         | 必须是组件参数流的适当类型。                        |
| updateUI | Integer | 如果 1，将在更新流的值后强制 Premiere Pro 更新 UI。 |

**返回**：如果成功则返回**0**。

### setValueAtKey()

app.project.sequences[index].audioTracks[index].clips[index].components[index].properties[index].setValueAtKey(time, value, updateUI)
app.project.sequences[index].videoTracks[index].clips[index].components[index].properties[index].setValueAtKey(time, value, updateUI)

**描述**：在指定的关键帧时间设置组件参数流的值。注意：只能与关键帧参数流一起使用。

**参数**

| 参数     | 类型                                                                     | 描述                                                |
| -------- | ------------------------------------------------------------------------ | --------------------------------------------------- |
| time     | [时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time) | 应设置关键帧值的时间。                              |
| value    |                                                                          | 要设置的值。                                        |
| updateUI | Integer                                                                  | 如果 1，将在更新流的值后强制 Premiere Pro 更新 UI。 |

**返回**：如果成功则返回**0**。
