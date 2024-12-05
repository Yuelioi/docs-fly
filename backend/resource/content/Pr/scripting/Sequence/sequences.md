---
title: sequences 序列对象
order: 4
---

---
title: sequences 序列对象
order: 4
category:

- PR脚本

---

## 序列对象

app.project.sequences[index]

**描述**：**序列**对象表示在 Premiere Pro 中的媒体序列（也叫做“时间轴”）。

## 属性

### audioDisplayFormat 音频显示模式

app.project.sequences[index].audioDisplayFormat

**描述**：音频显示模式，具体不清楚

**类型**：数字。

示例：查看当前轨道的音频显示模式

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.audioDisplayFormat.toString()); //200
```

### audioTracks 音频轨道

app.project.sequences[index].audioTracks

**描述**：序列中的一系列音轨。

**类型**：[TrackCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/trackcollection.html#trackcollection)，只读；

示例：获取第 1 个音轨的名称

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.audioTracks[0].name); // 音频 1
```

### end 序列结束

app.project.sequences[index].end

**描述**：序列结束的时间，以 Ticks 为单位。

**类型**：字符串; 只读。

![](https://cdn.yuelili.com/20211027124456.png)

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.ticks = sequence.end;

alert(newTime.seconds.toString()); // 200（秒）
```

### frameSizeHorizo​​ntal 宽度

app.project.sequences[index].frameSizeHorizontal

**描述**：序列中帧的水平宽度。

**类型**：整数; 只读。

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.frameSizeHorizontal.toString()); // 1920
```

### frameSizeVertical 高度

app.project.sequences[index].frameSizeVertical

**描述**：序列中帧的垂直高度。

**类型**：整数; 只读。

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.frameSizeVertical.toString()); // 1080
```

### id 序列序号

app.project.sequences[index].id

**描述**：创建序列的序数。Premiere Pro 会话期间在项目中创建的第 33 个序列，则该值为“33”。

**类型**：整数，只读。

示例：查看第一个序列的 ID

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.id.toString()); // 1
```

### markers 序列标记

app.project.sequences[index].markers

**描述**：与此序列关联的[Marker](https://ppro-scripting.docsforadobe.dev/general/marker.html#marker)对象。

**类型**：[MarkerCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/markercollection.html#markercollection)，只读；

### name 序列名

app.project.sequences[index].name

**描述**：序列的名称。

**类型**：字符串; 读/写。

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.name); // 序列 01
```

### projectItem 项目对象

app.project.sequences[index].projectItem

**描述**：与此序列关联的[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)。

**类型**：**项目项目**；只读。

### sequenceID 序列 ID

app.project.sequences[index].sequenceID

**描述**：在创建时分配给此序列的唯一标识符，格式为 xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。

**类型**：字符串; 只读。

![](https://cdn.yuelili.com/20211027125317.png)

### timebase 单帧持续时间

app.project.sequences[index].timebase

**描述**：序列中每帧的 Tick 数。

**类型**：字符串; 只读。

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.ticks = sequence.timebase;

alert(newTime.seconds.toString()); // 0.04（秒）
```

### videoDisplayFormat 视频显示模式

app.project.sequences[index].videoDisplayFormat

**描述**：视频显示模式，具体不清楚

**类型**：数字。

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoDisplayFormat.toString()); // 101
```

### videoTracks 视频轨道

app.project.sequences[index].videoTracks

**描述**：序列中的一系列视频轨道。

**类型**：[TrackCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/trackcollection.html#trackcollection)，只读；

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.videoTracks[0].name); // 视频1
```

### zeroPoint 序列零点

app.project.sequences[index].zeroPoint

**描述**：序列的开始时间，以 Ticks 为单位。

**类型**：字符串; 只读。

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.ticks = sequence.zeroPoint;

alert(newTime.seconds.toString()); // 0（秒）
```

## 方法

### autoReframeSequence() 重构序列

app.project.sequences[index].autoReframeSequence(numerator, denominator, motionPreset, newName, useNestedSequences)

**描述**：生成一个新的、重构的序列。

**参数**

| 参数         | 类型    | 描述           |
| ------------ | ------- | -------------- |
| numerator    | Integer | 帧长宽比分子   |
| denominator  | Integer | 帧长宽比分母。 |
| motionPreset | String  | \* “slower”    |

- “default”
- “faster”
  |
  | newName | String | 新序列名称。 |
  | useNestedSequences | Boolean | 是否要嵌套序列。 |

**返回**：如果成功，则返回新的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)；如果不成功则为 0。

|        |        |
| ------ | ------ |
| 创建前 | 创建后 |
|        |        |

示例：重构序列，并且把素材重新嵌套

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

var sequence = app.project.activeSequence;

if (sequence) {
  var numerator = 1;

  var denominator = 1;

  var motionPreset = "default"; // 'default', 'faster', 'slower'

  var newName = sequence.name + "-重构";

  var useNestedSequences = true;

  var newSequence = sequence.autoReframeSequence(numerator, denominator, motionPreset, newName, useNestedSequences);

  // 判断是否创建成功

  if (newSequence) {
    alert("创建成功：" + newName);
  } else {
    alert("创建 " + newName + "失败！");
  }
} else {
  alert("请激活一个序列");
}
```

### clone() 克隆序列

app.project.sequences[index].clone()

**描述**：创建给定序列的克隆。新序列名称结尾会自动加 “-副本”

**参数**：无。

**返回**：布尔值，克隆是否成功。

### createSubsequence() 创建子序列

app.project.sequences[index].createSubsequence(ignoreChannelMapping)

**描述**：创建一个新序列，它是现有序列的子序列。等同于 序列 - 制作子序列

**参数**

| 参数                 | 类型    | 描述                                       |
| -------------------- | ------- | ------------------------------------------ |
| ignoreChannelMapping | Boolean | 新序列是否应忽略原始序列中存在的通道映射。 |

**返回**：如果成功则返回 0。

![](https://cdn.yuelili.com/20211027135747.png)

示例：连续创建 2 次

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

sequence.createSubsequence(true);
```

### exportAsFinalCutProXML() 导出 XML

app.project.sequences[index].exportAsFinalCutProXML(outputPath)

**描述**：创建序列及其组成媒体的新 FCP XML 表示。

**参数**

| 参数       | 类型   | 描述                        |
| ---------- | ------ | --------------------------- |
| outputPath | String | 新 FCP XML 文件的输出路径。 |

**返回**：如果成功则返回 0。

### exportAsMediaDirect() 导出序列

app.project.sequences[index].exportAsMediaDirect(outputPath, presetPath, workAreaType)

**描述**：使用指定的输出预设（.epr 文件）将序列渲染到指定的输出路径，并遵循指定的工作区类型。

**参数**

| 参数         | 类型   | 描述                               |
| ------------ | ------ | ---------------------------------- |
| outputPath   | String | 将媒体渲染到的输出路径。           |
| presetPath   | String | ？？？                             |
| workAreaType |        | 必须是以下之一：\* 0 ENCODE_ENTIRE |

- 1 ENCODE_IN_TO_OUT
- 2 ENCODE_WORK_AREA
  |

**返回**：如果成功则返回 0。

### exportAsProject() 导出工程

app.project.sequences[index].exportAsProject(outputPath)

**描述**：创建一个仅包含给定序列及其组成媒体的新[项目对象](https://ppro-scripting.docsforadobe.dev/general/project.html#project)。

**参数**

| 参数       | 类型   | 描述               |
| ---------- | ------ | ------------------ |
| outputPath | String | 新项目的输出路径。 |

**返回**：如果成功则返回 0。

### getExportFileExtension() 检索拓展名

app.project.sequences[index].getExportFileExtension(outputPresetPath)

**描述**：检索与当前序列关联的文件扩展名。

**参数**

| 参数             | 类型   | 描述               |
| ---------------- | ------ | ------------------ |
| outputPresetPath | String | 要使用的输出预设。 |

**返回**：返回包含输出文件扩展名的**字符串**，如果不成功则返回**0**。

### getInPoint() 获取序列入点

app.project.sequences[index].getInPoint()

**描述**：以点为单位检索当前序列，以秒为单位。

**参数**：无。

**返回**：返回一个 Real 表示入点，以秒为单位。

示例：

![](https://cdn.yuelili.com/20211028024217.png)

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.getInPoint().toString()); // 20
```

### getInPointAsTime() 获取序列入点

app.project.sequences[index].getInPointAsTime()

**描述**：检索点中的当前序列。

**参数**：无。

**返回**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，以秒为单位表示入点。

### getOutPoint() 获取序列出点

app.project.sequences[index].getOutPoint()

**描述**：以秒为单位检索当前序列输出点。

**参数**：无。

**返回**：返回代表出点的 Real，以秒为单位。

![](https://cdn.yuelili.com/20211028024217.png)

示例：

```javascript
var sequence = app.project.activeSequence; // 获取当前序列

alert(sequence.getOutPoint().toString()); // 50（秒）
```

### getOutPointAsTime() 获取序列出点

app.project.sequences[index].getOutPointAsTime()

**描述**：检索当前序列输出点。

**参数**：无。

**返回**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，以秒为单位表示出点。

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 20;

alert(sequence.getOutPointAsTime().seconds.toString()); // 50（秒）
```

### getPlayerPosition() 获取时间轴位置

app.project.sequences[index].getPlayerPosition()

**描述**：检索当前时间轴位置，以 Ticks 为单位。

**参数**：无。

**返回**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，表示当前玩家的位置。

![](https://cdn.yuelili.com/20211028024217.png)

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 20;

alert(sequence.getPlayerPosition().seconds.toString()); // 36（秒）
```

### getSettings() 检索设置

app.project.sequences[index].getSettings()

**描述**：检索当前序列的设置。

**参数**：无。

**返回**：返回序列设置结构。

| audioChannelCount | 序列中的音频通道数。 |
| audioChannelType |
正在使用的音频通道类型。以下之一：

- 0 AUDIOCHANNELTYPE_Mono
- 1 AUDIOCHANNELTYPE_Stereo
- 2 AUDIOCHANNELTYPE_51
- 3 AUDIOCHANNELTYPE\_多声道
- 4 AUDIOCHANNELTYPE_4Channel
- 5 AUDIOCHANNELTYPE_8Channel

|
| audioDisplayFormat |
音频时间码显示格式。以下之一：

- 100 TIMEDISPLAY_24 时间码
- 101 TIMEDISPLAY_25 时间码
- 102 TIMEDISPLAY_2997DropTimecode
- 103 TIMEDISPLAY_2997NonDropTimecode
- 104 TIMEDISPLAY_30 时间码
- 105 TIMEDISPLAY_50 时间码
- 106 TIMEDISPLAY_5994DropTimecode
- 107 TIMEDISPLAY_5994NonDropTimecode
- 108 TIMEDISPLAY_60 时间码
- 109 TIMEDISPLAY_Frames
- 110 TIMEDISPLAY_23976 时间码
- 111 TIMEDISPLAY_16mmFeetFrames
- 112 TIMEDISPLAY_35mmFeetFrames
- 113 TIMEDISPLAY_48 时间码
- 200 TIMEDISPLAY_AudioSamplesTimecode
- 第 201 话

|
| audioSampleRate | 序列中的音频采样率，作为 int. |
| compositeLinearColor | 序列是否以线性颜色合成。1 如果为真。 |
| editingMode | 正在使用的编辑模式的 GUID。 |
| maximumBitDepth | 序列是否在最大深度合成；1 如果为真。 |
| maximumRenderQuality | 序列是否以最高质量呈现；1 如果为真。 |
| previewCodec | 正在使用的预览编解码器的四字符代码。 |
| previewFrameWidth | 预览框的宽度。 |
| previewFrameHeight | 预览框的高度。 |
| previewFileFormat | 用于预览文件渲染的输出预设（.epr 文件）的路径。 |
| videoDisplayFormat |
视频时间显示格式。以下之一：

- 100 TIMEDISPLAY_24 时间码
- 101 TIMEDISPLAY_25 时间码
- 102 TIMEDISPLAY_2997DropTimecode
- 103 TIMEDISPLAY_2997NonDropTimecode
- 104 TIMEDISPLAY_30 时间码
- 105 TIMEDISPLAY_50 时间码
- 106 TIMEDISPLAY_5994DropTimecode
- 107 TIMEDISPLAY_5994NonDropTimecode
- 108 TIMEDISPLAY_60 时间码
- 109 TIMEDISPLAY_Frames
- 110 TIMEDISPLAY_23976 时间码
- 111 TIMEDISPLAY_16mmFeetFrames
- 112 TIMEDISPLAY_35mmFeetFrames
- 113 TIMEDISPLAY_48 时间码
- 200 TIMEDISPLAY_AudioSamplesTimecode
- 第 201 话

|
| videoFieldType |
按顺序使用的视频字段类型。其中之一：

- -1 FIELDTYPE_DEFAULT
- 0 FIELDTYPE_PROGRESSIVE
- 1 ALPHACHANNEL_UPPERFIRST
- 2 ALPHACHANNEL_LOWERFIRST

|
| videoFrameHeight | 序列视频帧的高度。 |
| videoFrameWidth | 序列视频帧的高度。 |
| videoPixelAspectRatio | 像素纵横比，作为一个比率，作为一个字符串。 |
| vrHorzCapturedView | |
| vrVertCapturedView | |
| vrLayout |
正在使用的素材布局，用于 VR。其中之一：

- 0 VR_LAYOUT_MONOSCOPIC
- 1 VR_LAYOUT_STEREO_OVER_UNDER
- 2 VR_LAYOUT_STEREO_SIDE_BY_SIDE

|
| vrProjection |
正在使用的投影类型，用于 VR 素材。其中之一：

- 0 VR_LAYOUT_MONOSCOPIC
- 1 VR_LAYOUT_STEREO_OVER_UNDER
- 2 VR_LAYOUT_STEREO_SIDE_BY_SIDE

|
| videoFieldType |
按顺序输入字段。以下之一：

- -1 FIELDTYPE_DEFAULT
- 0 FIELDTYPE_PROGRESSIVE
- 1 ALPHACHANNEL_UPPERFIRST
- 2 ALPHACHANNEL_LOWERFIRST

|

### isDoneAnalyzingForVideoEffects() 视频效果分析

app.project.sequences[index].isDoneAnalyzingForVideoEffects()

**描述**：返回序列是否已完成视频效果分析。

**参数**：无。

**返回**：如果分析完成则返回 true。

### performSceneEditDetectionOnSelection() 剪切检测

app.project.sequences[index].performSceneEditDetectionOnSelection(actionDesired, applyCutsToLinkedAudio, sensitivity)

**描述**：对序列选择执行剪切检测。

**参数**

| 参数          | 类型   | 描述                |
| ------------- | ------ | ------------------- |
| actionDesired | String | 之一：\* “创建标记” |

- “ApplyCuts”
  |
  | applyCutsToLinkedAudio | Boolean | |
  | sensitivity | String | 之一：\* “低灵敏度”
- “中等灵敏度”
- “高灵敏度”
  |

**返回**：如果成功则返回真。

### setInPoint() 设置序列入点

app.project.sequences[index].setInPoint(time)

**描述**：指定一个新的序列入点。

**参数**

| 参数 | 类型   | 描述                  |
| ---- | ------ | --------------------- |
| time | String | **ticks**类型新时间。 |

**返回**：如果成功则返回**0**。

示例：入点设置为 20 秒

![](https://cdn.yuelili.com/20211028023953.png)

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 20;

sequence.setInPoint(newTime.ticks);
```

### setOutPoint() 设置序列出点

app.project.sequences[index].setOutPoint(time)

**描述**：指定一个新的序列出点。

**参数**

| 参数 | 类型   | 描述                       |
| ---- | ------ | -------------------------- |
| time | String | **ticks 类型的**的新时间。 |

**返回**：如果成功则返回**0**。

示例：出点设置为 50 秒

![](https://cdn.yuelili.com/20211028024020.png)

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 50;

sequence.setOutPoint(newTime.ticks);
```

### setPlayerPosition() 设置时间轴位置

app.project.sequences[index].setPlayerPosition(time)

**描述**：指定一个新的玩家位置，在 Ticks 中，作为一个字符串。

**参数**

| 参数 | 类型   | 描述                     |
| ---- | ------ | ------------------------ |
| time | String | **ticks 类型**的新时间。 |

**返回**：如果成功则返回**0**。

示例：把时间指示器设置为 5 秒

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 5;

sequence.setPlayerPosition(newTime.ticks);
```

### setSettings() 当前序列设置

app.project.sequences[index].setSettings(sequenceSettings)

**描述**：设置当前序列的设置。_[社论：我为任何被察觉的迂腐致歉；有时，明显的文档需要显而易见。-bbb]_

**参数**

| 参数             | 类型 | 描述                                                                                                                                  |
| ---------------- | ---- | ------------------------------------------------------------------------------------------------------------------------------------- |
| sequenceSettings |      | 通过[Sequence.getSettings()](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence-getsettings)获得的序列设置结构。 |

**返回**：如果成功则返回 0。

### setZeroPoint() 设置序列开始时间

app.project.sequences[index].setZeroPoint(newZeroPoint)

**描述**：设置序列的开始时间。

**参数**

| 参数         | 类型   | 描述                       |
| ------------ | ------ | -------------------------- |
| newZeroPoint | String | **ticks 类型的**的新零点。 |

**类型**：整数; 只读。

**返回**：如果成功则返回**0**。

示例：序列开始时间设置为 0

```javascript
var newTime = new Time(); // 定义时间对象

var sequence = app.project.activeSequence; // 获取当前序列

newTime.seconds = 0;

sequence.setZeroPoin(newTime.ticks);
```
