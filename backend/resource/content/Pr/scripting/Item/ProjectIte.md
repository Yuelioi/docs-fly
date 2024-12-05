---
title: ProjectItem 工程项对象
order: 1
---

---
title: ProjectItem 工程项对象
order: 2
category:

- PR脚本

---

## 工程项目对象

app.project.rootItem.children[index]

**描述**：其中每一项都是一个**projectItem**，包括根项目（rootItem）。

示例：遍历项目面板的项目

![](https://cdn.yuelili.com/20211027235014.png)

```javascript
var p = app.project.rootItem.children;

for (var i = 0; i < p.numItems; i++) {
  alert(p[i].name); // 素材箱、hakase.png.....
}
```

## 属性

### children 项目数组

app.project.rootItem.children

**描述**：包含在指定工程中的项目数组。

**类型**：[ProjectItemCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/projectitemcollection.html#projectitemcollection)，只读。

示例：查看项目面板一共有多少项

```javascript
alert(app.project.rootItem.children.length.toString()); // 也可以用numItems替换length
```

### getAudioChannelMapping 音频通道映射

app.project.rootItem.children[index].getAudioChannelMapping

**描述**：当前应用于此**projectItem**的音频通道映射。

**类型**：audioChannelMapping 对象。

### getOverrideColorSpaceList

app.project.rootItem.children[index].getOverrideColorSpaceList

**描述**：_增加一个说明_

返回：一个对象，包含相似的数据

```javascript
{

    value: [

        sRGB,

        BT.601 (NTSC),

        BT.601 (PAL),

        BT.709,

        BT.709 (Scene),

        BT.2020,

        BT.2020 (Scene),

        BT.2100 PQ,

        BT.2100 PQ (Scene),

        BT.2100 HLG,

        BT.2100 HLG (Scene),

        DCDM XYZ,

    ]

};


```

**类型**：Javascript 对象。

### name 项目名称

app.project.rootItem.children[index].name

**描述**：项目的名称。

**类型**：字符串; 读/写。

**示例**：把我滴第一个序列，重命名为 序列 XXX

```javascript
var item = app.project.rootItem.children[0];

if (item) {
  item.name = "序列 XXX";
} else {
  alert("这个项目无法重命名");
}
```

### nodeId 项目 ID

app.project.rootItem.children[index].nodeId

**描述**：项目添加到项目面板时，分配的唯一 ID。

**注意**：区分对相同源媒体的引用。

**类型**：字符串; 只读。

![](https://cdn.yuelili.com/20211027235708.png)

```javascript
alert(app.project.rootItem.children[1].nodeId); // 00f4222 （总之就是一个单独ID）
```

### teamProjectsAssetId 团队项目资产 ID

app.project.rootItem.children[index].teamProjectsAssetId

**描述**：项目的团队项目资产 ID。

**类型**：字符串; 只读。

### treePath 项目路径

app.project.rootItem.children[index].treePath

**描述**：项目的当前项目位置。

**类型**：字符串; 只读。

示例：第 1 个序列的位置

![](https://cdn.yuelili.com/20211027150519.png)

```javascript
var items = app.project.rootItem.children;

alert(items[0].treePath); //
```

### type 项目类型

app.project.rootItem.children[index].type

**描述**：四大类型，**CLIP**、**BIN**、**ROOT**或**FILE**。（剪辑片段、素材箱、根、文件）一般用来判断

**类型**：枚举值；只读。

## 方法

### attachProxy()  附加项目

app.project.rootItem.children[index].attachProxy(mediaPath, isHiRes)

**描述**：将媒体附加 newMediaPath 到项目，作为高分辨率或代理媒体。

**参数**

| 参数      | 类型    | 描述                                     |
| --------- | ------- | ---------------------------------------- |
| mediaPath | String  | 新媒体的路径。                           |
| isHiRes   | Integer | 附加新媒体，0 作为代理，1 高分辨率媒体。 |

**返回**：如果成功则返回**0**。

### canChangeMediaPath() 可以更换媒体路径

app.project.rootItem.children[index].canChangeMediaPath()

**描述**：判断   Premiere Pro 可否更改与此项目关联的路径。

**参数**：无。

**返回**：布尔值；可以更换，则为**true，**否则为**false**。

### canProxy() 能否创建代理

app.project.rootItem.children[index].canProxy()

**描述**：能否将代理附加到此项目。

**参数**：无。

**返回**：如果项目允许附加代理，则返回**true**；如果不是，则为**假**。

### changeMediaPath() 更改媒体路径

app.project.rootItem.children[index].changeMediaPath(newPath)

**描述**：更新项目，指向新的媒体路径。

**参数**

| 参数           | 类型    | 描述               |
| -------------- | ------- | ------------------ |
| newPath        | String  | 媒体文件的新路径。 |
| overrideChecks | Boolean | 忽略任何安全问题。 |

**返回**：如果替换成功，则返回**0**。

### clearOutPoint() 清除出点

app.project.rootItem.children[index].clearOutPoint()

**描述**：清除任何指定的出点；然后项目从 startTime 开始。

**参数**：无

**返回**：成功则返回 0。

### createBin() 创建素材箱

app.project.rootItem.createBin(name)

**描述**：创建一个空的素材箱。等同于右键项目面板 - 创建素材箱（然后自己重命名）

**参数**

| 参数 | 类型   | 描述            |
| ---- | ------ | --------------- |
| name | String | 新 bin 的名称。 |

**返回**：如果成功则返回新 bin 的项目，如果不成功则返回**0**。

示例：创建一个素材箱

```javascript
app.project.rootItem.createBin("素材箱 01");
```

### createSmartBin() 创建搜索箱

app.project.rootItem.createSmartBin(name, queryString)

**描述**：创建一个搜索箱；bin 项目。差不多等同于右键项目面板 - 创建搜索素材箱（然后自己重命名）

**参数**

| 参数        | 类型   | 描述                   |
| ----------- | ------ | ---------------------- |
| name        | String | 新 bin 的名称。        |
| queryString | String | 用于搜索的查询字符串。 |

**返回**：如果成功创建，则返回**0**。

示例：创建一个视频搜索箱

![](https://cdn.yuelili.com/20211028000721.png)

```javascript
app.project.rootItem.createSmartBin("视频", "mp4");
```

### createSubClip() 创建子剪辑

app.project.rootItem.children[index].createSubClip(name, startTime, endTime, hasHardBoundaries, takeAudio, takeVideo)

**描述**：为现有项目的子剪辑创建新项目。

**参数**

| 参数              | 类型    | 描述                                   |
| ----------------- | ------- | -------------------------------------- |
| name              | String  | 新子剪辑的名称。                       |
| startTime         | String  | 子剪辑的开始时间，以**Ticks 为单位**。 |
| endTime           | String  | 子剪辑的结束时间，以**Ticks 为单位**。 |
| hasHardBoundaries | Integer | 如果 1，用户不能设置入点和出点         |
| takeVideo         | Integer | 如果 1，使用源视频。                   |
| takeAudio         | Integer | 如果 1，使用源音频。                   |

**返回**：返回表示新子剪辑的项目，如果创建失败，则返回 0。

示例：基于一个视频，创建一个 1 秒的子剪辑。

```javascript
app.project.rootItem.children[3].createSubClip("1", "0", "254016000000", 1, 1, 1);
```

### deleteBin() 删除素材箱

app.project.rootItem.children[index].deleteBin()

**描述**：从项目中删除一个 bin**及其所有内容**。

**参数**：无。

**返回**：如果删除成功则返回**0**。

### findItemsMatchingMediaPath() 引用相同的媒体路径

app.project.rootItem.children[index].findItemsMatchingMediaPath(pathToMatch, ignoreSubClips)

**描述**：返回一组项目，它们都引用相同的媒体路径。

**参数**

| 参数           | 类型    | 描述                           |
| -------------- | ------- | ------------------------------ |
| pathToMatch    | String  | 匹配的路径。                   |
| ignoreSubClips | Integer | 如果 1，则不会返回任何子剪辑。 |

**返回**：返回一个项目数组，如果没有找到与 匹配的项目路径，则返回**0**。

### getColorLabel() 检索项目颜色标签

app.project.rootItem.children[index].getColorLabel()

**描述**：检索项目的颜色标签。

**参数**：无。

**返回**：整数，以下其中之一

| 标签颜色 | \* 0 = 紫色

- 1 = 虹膜
- 2 = 加勒比海
- 3 = 薰衣草
- 4 = 蔚蓝
- 5 = 森林
- 6 = 玫瑰
- 7 = 芒果
- 8 = 紫色
- 9 = 蓝色
- 10 = 青色
- 11 = 洋红色
- 12 = 棕褐色
- 13 = 绿色
- 14 = 棕色
- 15=黄色
  |

### getFootageInterpretation() 当前解释结构

app.project.rootItem.children[index].getFootageInterpretation()

**描述**：返回描述项目当前解释的结构。

**参数**：无。

**返回**：一个镜头解释结构，如果不成功则为 0。

| alphaUsage |
Alpha，以下之一：

- 0 ALPHACHANNEL_NONE
- 1 ALPHACHANNEL_STRAIGHT
- 2 ALPHACHANNEL_PREMULTIPLIED
- 3 ALPHACHANNEL_IGNORE

|
| fieldType |
字段类型，以下之一：

- -1 FIELDTYPE_DEFAULT
- 0 FIELDTYPE_PROGRESSIVE
- 1 FIELDTYPE_UPPERFIRST
- 2 FIELDTYPE_LOWERFIRST

|
| ignoreAlpha | true 或 false。 |
| invertAlpha | true 或 false。 |
| frameRate | 帧速率，浮点值。 |
| pixelAspectRatio | 像素长宽比，浮点值。 |
| removePulldown | true 或 false。 |
| vrConformProjectionType | 正在使用的投影类型，用于 VR 素材。其中之一：0 VR_CONFORM_PROJECTION_NONE1 VR_CONFORM_PROJECTION_EQUIRECTANGULAR |
| vrLayoutType |
正在使用的素材布局，用于 VR。其中之一：

- 0 VR_LAYOUT_MONOSCOPIC
- 1 VR_LAYOUT_STEREO_OVER_UNDER
- 2 VR_LAYOUT_STEREO_SIDE_BY_SIDE

|
| vrHorizontalView | 正在使用的水平视图，用于 VR 素材。 |
| vrVerticalView | 正在使用的垂直视图，用于 VR 素材。 |

### getInPoint() 当前项目入点

app.project.rootItem.children[index].getInPoint()

**描述**：获取当前项目的入点。

**参数**：无。

**返回**：一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，包含入点。

### getMarkers() 当前项目标记

app.project.rootItem.children[index].getMarkers()

**描述**：检索与此项目关联的[MarkerCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/markercollection.html#markercollection)。

**参数**：无。

**返回**：[MarkerCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/markercollection.html#markercollection)，只读；

### getMediaPath() 获取媒体路径

app.project.rootItem.children[index].getMediaPath()

**描述**：以字符串形式返回与项目的媒体关联的路径。**注意**：这仅适用于原子媒体；此调用无法为没有实际路径的媒体提供有意义的路径（对于由合成导入器生成的任何媒体，例如 Premiere Pro 自己的 Universal Counting Leader，情况都是如此）。此外，对于图像序列，仅返回序列中第一个图像的路径。

**参数**：无。

**返回**：包含与项目关联的媒体路径的字符串。

### getOutPoint() 当前项目出点

app.project.rootItem.children[index].getOutPoint(mediaType)

**描述**：检索指定媒体类型的当前出点。

**参数**

| 参数      | 类型    | 描述                                                                         |
| --------- | ------- | ---------------------------------------------------------------------------- |
| mediaType | Integer | 1 仅限视频，或 2 仅音频。如果没有 mediaType 传递，函数获取所有媒体的输出点。 |

**返回**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)。

### getProjectMetadata() 获取元数据

app.project.rootItem.children[index].getProjectMetadata()

**描述**：检索与项目关联的元数据。与媒体 XMP 不同。

**参数**：无。

**返回**：包含所有 Premiere Pro 私有项目元数据的字符串，已序列化。

### getProxyPath() 检索代理路径

app.project.rootItem.children[index].getProxyPath()

**描述**：检索与此项目关联的代理媒体的路径。

**参数**：无。

**返回**：返回与代理项关联的代理媒体的路径（作为**String**），如果没有找到，则返回**0**。

### getXMPMetadata() 获取 XMP 元数据

app.project.rootItem.children[index].getXMPMetadata()

**描述**：检索与项目关联的 XMP 元数据，作为字符串。

**参数**：无。

**返回**：包含所有 XMP 元数据的字符串，已序列化。

### hasProxy() 是否有代理

app.project.rootItem.children[index].hasProxy()

**描述**：指示代理是否已附加到项目。

**参数**：无。

**返回**：如果项目附加了代理，则返回**true**；如果不是，则为**假**。

### isMergedClip() 是否引用合并剪辑

app.project.rootItem.children[index].isMergedClip()

**描述**：指示项目是否引用合并的剪辑。

**参数**：无。

**返回**：true 如果项目是合并剪辑，返回 true，否则返回 false。

### isMulticamClip() 是否涉及多机位剪辑

app.project.rootItem.children[index].isMulticamClip()

**描述**：项目是否涉及多机位剪辑。

**参数**：无。

**返回**：是则返回 true，不是则为 false。

### isOffline() 是否脱机

app.project.rootItem.children[index].isOffline()

**描述**：项目是否脱机。

**参数**：无。

**返回**：布尔值，如果离线则为 true。

### isSequence() 是否是序列

app.project.rootItem.children[index].isSequence()

**描述**：项目是否引用[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**：无。

**返回**：是多机位剪辑还是合并剪辑则返回 true。如果不是其中任何一个，则返回 false。

### moveBin() 移动到素材箱

app.project.rootItem.children[index].moveBin(newParentBinProjectItem)

**描述**：将 projectItem 移动到新的父 bin 中。

**参数**：

- newParentBinProjectItem：父级素材箱

**返回**：如果移动成功，则返回**0**。

### refreshMedia() 刷新素材

app.project.rootItem.children[index].refreshMedia()

**描述**：强制 Premiere Pro 更新其与项目关联的媒体的表示。如果媒体之前离线，这可能会导致它变为在线（如果之前丢失的媒体已经可用）。

**参数**：无。

**返回**：与项目关联的标记数组，如果没有标记，则为**0**。

### renameBin() 重命名素材箱

app.project.rootItem.children[index].renameBin(newName)

**描述**：更改 bin 的名称。仅适用于素材箱项目。

**参数**

| 参数    | 类型   | 描述                |
| ------- | ------ | ------------------- |
| newName | String | 一个新的 bin 名称。 |

**返回**：如果重命名 成功，则返回**0**。

### select() 选择状态

app.project.rootItem.children[index].select()

**描述**：将项目（必须是 bin）设置为后续导入项目的目标。

**参数**：无。

**返回**：如果项目已成功成为目标，则返回**0**，用于后续导入。

### setColorLabel() 设置颜色标签

app.project.rootItem.children[index].setColorLabel(labelColor)

**描述**：设置项目的颜色标签。

**参数**

| 参数       | 类型    | 描述                                                                                                                                     |
| ---------- | ------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| labelColor | Integer | 标签颜色；请参阅[ProjectItem.getColorLabel()](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem-getcolorlabel)。 |

**返回**：如果成功则返回 0 。

### setFootageInterpretation() 项目解释结构

app.project.rootItem.children[index].setFootageInterpretation(interpretation)

**描述**：返回描述 projectItem 当前解释的结构。

**参数**

| 参数           | 类型 | 描述               |
| -------------- | ---- | ------------------ |
| interpretation |      | 一个镜头解释结构。 |

**返回**：true 如果成功。

### setInPoint() 设置当前入点

app.project.rootItem.children[index].setInPoint(time, mediaType)

**描述**：对于指定的媒体类型，将入点设置为 timeInTicks 。

**参数**

| 参数      | 类型    | 描述                                                                  |
| --------- | ------- | --------------------------------------------------------------------- |
| time      | String  | **Ticks** 时间。也可以是秒（整数）                                    |
| mediaType | Integer | 确定要影响的媒体类型；1 仅限视频，2 仅限音频，或 4 针对所有媒体类型。 |

**返回**：0 成功则返回。

示例：入点设置为 2 秒

```javascript
app.project.rootItem.children[2].setInPoint(2, 4);
```

### setOffline() 项目脱机

app.project.rootItem.children[index].setOffline()

**描述**：使项目脱机。

**参数**：无。

**返回**：如果成功则返回 true 。

### setOutPoint() 设置出点

app.project.rootItem.children[index].setOutPoint(time, mediaType)

**描述**：对于指定的媒体类型，将出点设置为 timeInTicks。

**参数**

| 参数      | 类型    | 描述                                                                  |
| --------- | ------- | --------------------------------------------------------------------- |
| time      | String  | **Ticks**时间。也可以是整数（秒）                                     |
| mediaType | Integer | 确定要影响的媒体类型；1 仅限视频，2 仅限音频，或 4 针对所有媒体类型。 |

**返回**：成功则返回 0。

示例

`app.project.rootItem.children[2].setOutPoint(2,4)`

### setOverrideFrameRate() 设置项目帧速率

app.project.rootItem.children[index].setOverrideFrameRate(newFrameRate)

**描述**：设置项目的帧速率。

**参数**

| 参数         | 类型   | 描述         |
| ------------ | ------ | ------------ |
| newFrameRate | 浮点值 | 新的帧速率。 |

**返回**：如果已成功更改帧速率，则返回**0**。

### setOverridePixelAspectRatio() 设置项目像素长宽比

app.project.rootItem.children[index].setOverridePixelAspectRatio(numerator, denominator)

**描述**：设置项目的像素长宽比。

**参数**

| 参数        | 类型    | 描述   |
| ----------- | ------- | ------ |
| numerator   | Integer | 分子。 |
| denominator | Integer | 分母。 |

**返回**：如果成功更改，则返回**0**。

### setProjectMetadata() 设置项目元数据

app.project.rootItem.children[index].setProjectMetadata(newMetadata, updatedFields)

**描述**：设置与项目关联的私有项目元数据。

**参数**

| 参数          | 类型   | 描述                               |
| ------------- | ------ | ---------------------------------- |
| newMetadata   | String | 一个新的、序列化的私有项目元数据。 |
| updatedFields | Array  | 包含要更新的字段名称的数组。       |

**返回**：如果更新成功则返回 0。

### setScaleToFrameSize() 设置项目帧大小

app.project.rootItem.children[index].setScaleToFrameSize()

**描述**：打开缩放到帧大小，用于将此项目中的媒体插入到序列中时。

**参数**：无。

**返回**：未定义的返回值。

### setStartTime() 设置项目起始时间

app.project.rootItem.children[index].setStartTime(time)

**描述**：为项目分配新的起始时间

**参数**

| 参数 | 类型   | 描述                            |
| ---- | ------ | ------------------------------- |
| time | String | 新的开始时间，用**Ticks**表示。 |

**返回**：0 成功则返回。

### setXMPMetadata() 设置 XMP 元数据

app.project.rootItem.children[index].setXMPMetadata(newXMP)

**描述**：设置与项目关联的 XMP 元数据。

**参数**

| 参数   | 类型   | 描述                        |
| ------ | ------ | --------------------------- |
| newXMP | String | 新的、序列化的 XMP 元数据。 |

**返回**：如果更新成功则返回 0。

### startTime() 项目开始时间

app.project.rootItem.children[index].startTime()

**描述**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，表示开始时间。

**参数**：无。

**返回**：[时间对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)。

### videoComponents() 视频组件

app.project.rootItem.children[index].videoComponents()

**描述**：此项目的“主剪辑”的视频组件。

**类型**：[ComponentCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/componentcollection.html#componentcollection)，只读。
