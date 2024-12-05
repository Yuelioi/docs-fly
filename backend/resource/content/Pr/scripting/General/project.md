---
title: project 项目对象
order: 4
---

---
title: project 项目对象
order: 7
category:

- PR脚本

---

## 项目对象

app.project

**描述**：代表 Premiere Pro 项目。从 Premiere Pro 12.0 开始，可以同时打开多个项目。

## 属性

### activeSequence 活动序列

app.project.activeSequence

**描述**：项目中当前活动的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**类型**：一个[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)，或者 0 如果当前没有序列处于活动状态。

### cloudProjectlocalID 云项目 ID

app.project.cloudProjectlocalID

**描述**：云项目 ID。

**类型**：字符串; 只读。

### documentID 项目 ID

app.project.documentID

**描述**：此项目的唯一标识符，格式为 xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.

**类型**：字符串; 只读。

### isCloudProject 是否为云项目

app.project.isCloudProject

**描述**：检查项目是否为云项目。

**类型**：布尔值；只读。

### name 项目名称

app.project.name

**描述**：项目名称。

**类型**：字符串; 只读。

### path 项目路径

app.project.path

**描述**：项目的文件路径。

**类型**：字符串; 只读。

**示例**：获取当前活动项目的路径

```javascript
app.project.path; // /Users/USERNAME/Desktop/Project.prproj
```

### rootItem 根项目

app.project.rootItem

**描述**：一个[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)表示该项目的“根”。

**类型**：一个[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)；这将始终是类型 ProjectItemType_BIN。

### sequences 序列

app.project.sequences

**描述**：项目中的序列。

**类型**：[SequenceCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/sequencecollection.html#sequencecollection)，只读。

## 方法

### addPropertyToProjectMetadataSchema() 私有项目元数据架构

app.project.addPropertyToProjectMetadataSchema(propertyName, propertyLabel, propertyType)

**描述**：将指定类型的新字段添加到 Premiere Pro 的私有项目元数据架构。

**参数**

| 参数          | 类型   | 描述                         |
| ------------- | ------ | ---------------------------- |
| propertyName  | String | 要添加的属性的名称。         |
| propertyLabel | String | 要添加的属性标签。           |
| propertyType  |        | 必须是以下之一：\* 0 Integer |

- 1 Real
- 2 String
- 3 Boolean
  |

**返回**：如果成功则返回**true**，如果不成功则返回**undefined**。

### closeDocument() 关闭此项目

app.project.closeDocument(saveFirst, promptIfDirty)

**描述**：关闭此项目。

**参数**

| 参数          | 类型    | 描述                                 |
| ------------- | ------- | ------------------------------------ |
| saveFirst     | Integer | 如果 1，项目将在关闭前保存。         |
| promptIfDirty | Integer | 如果 1，将询问用户是否要先保存更改。 |

**返回**：如果成功则返回**0**。

### consolidateDuplicates() 合并重复素材

app.project.consolidateDuplicates()

**描述**：调用 Premiere Pro 的“合并重复素材”功能。

**参数**：无。

**返回**：如果成功则返回 **0**。

### createNewSequence() 创建新序列

app.project.createNewSequence(sequenceName, sequenceID)

**描述**：使用指定的 ID 创建一个新的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**

| 参数         | 类型   | 描述                  |
| ------------ | ------ | --------------------- |
| sequenceName | String | 序列的名称。          |
| sequenceID   | String | 新序列的唯一标识 ID。 |

**返回**：如果创建成功则返回一个[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)，如果不成功则返回**0**。

### createNewSequenceFromClips() 从剪辑创建新序列

app.project.createNewSequenceFromClips(sequenceName, arrayOfProjectItems, destinationBin);

**描述**：在指定的目标 bin 中创建一个具有给定名称的新[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)，并按顺序将项目项插入其中。等同于项目面板右键 从剪辑创建新序列

![](https://cdn.yuelili.com/20211029004236.png)

**参数**

| 参数                | 类型                                                                                                 | 描述                         |
| ------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------- |
| sequenceName        | String                                                                                               | 可选的。新序列的名称。       |
| arrayOfProjectItems | Array，[project Item](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)对象 | 要插入序列的项目项数组。     |
| destinationBin      | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)                | 可选的。一个包含序列的 bin。 |

**返回**：如果成功，则返回新创建的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)；如果不成功则返回 0。

### deleteSequence() 删除序列

app.project.deleteSequence(sequence)

**描述**：从项目中删除指定的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**

| 参数     | 类型                                                                                | 描述           |
| -------- | ----------------------------------------------------------------------------------- | -------------- |
| sequence | [序列对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence) | 要删除的序列。 |

**返回**：如果成功则返回 0。

### exportAAF() 导出 AAF

app.project.exportAAF(sequenceToExport, outputPath, mixdownVideo, explodeToMono, sampleRate, bitsPerSample, embedAudio, audioFileFormat, trimSources, handleFrames, presetPath, renderAudioEffects, includeClipCopies, preserveParentFolder)

**描述**：使用指定的设置导出指定[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)的 AAF 文件。

**参数**

| 参数                 | 类型                                                                                | 描述                                 |
| -------------------- | ----------------------------------------------------------------------------------- | ------------------------------------ |
| sequence             | [序列对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence) | 要导出的序列。                       |
| filePath             | String                                                                              | .aaf 文件的输出路径。                |
| mixdownVideo         | Integer                                                                             | 如果 1，则在导出前渲染视频。         |
| explodeToMono        | Integer                                                                             | 如果 1，则将立体声轨道分解为单声道。 |
| sampleRate           |                                                                                     | 输出音频的采样率。                   |
| bitsPerSample        |                                                                                     | 每个音频输出样本的位数。             |
| embedAudio           | Integer                                                                             | 如果 1，则嵌入音频，如果 0，则外部。 |
| audioFileFormat      | Integer                                                                             | 0 是 AIFF，1 是 WAV。                |
| trimSources          | Integer                                                                             | 如果 1，则在导出前修剪音频文件。     |
| handleFrames         | Integer                                                                             | 句柄帧数（从 0 到 1000）。           |
| presetPath           | String                                                                              | 导出预设 (.epr) 文件的路径。         |
| renderAudioEffects   | Integer                                                                             | 如果 1，则在导出前渲染音频效果。     |
| includeClipCopies    | Integer                                                                             | 如果 1，则包括剪辑的每个副本。       |
| preserveParentFolder | Integer                                                                             | 如果 1，则在输出中保留父文件夹。     |

**返回**：如果成功则返回**0**。

### exportFinalCutProXML() 导出 XML

app.project.exportFinalCutProXML(outputPath, suppressUI)

**描述**：将整个项目的 FCP XML 表示形式导出到指定的输出路径。

**参数**

| 参数       | 类型    | 描述                                     |
| ---------- | ------- | ---------------------------------------- |
| outputPath | String  | .xml 文件的输出路径。                    |
| suppressUI | Integer | 如果 1，则在导出期间不会显示警告或警报。 |

**返回**：如果成功则返回 0。

### exportOMF() 导出 OMF

app.project.exportOMF(sequence, outputPath, omfTitle, sampleRate, bitsPerSample, audioEncapsulated, audioFileFormat, trimAudioFiles, handleFrames, includePan)

**描述**：使用指定的设置导出指定[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)的 OMF 文件。

**参数**

| 参数              | 类型                                                                                | 描述                                 |
| ----------------- | ----------------------------------------------------------------------------------- | ------------------------------------ |
| sequence          | [序列对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence) | 要输出的序列。                       |
| filePath          | String                                                                              | .omf 文件的输出路径。                |
| omfTitle          | String                                                                              | OMF 的标题。                         |
| sampleRate        |                                                                                     | 输出音频的采样率。                   |
| bitsPerSample     |                                                                                     | 每个音频输出样本的位数。             |
| audioEncapsulated | Integer                                                                             | 如果 1，则嵌入音频，如果 0，则外部。 |
| audioFileFormat   | Integer                                                                             | 0 是 AIFF，1 是 WAV。                |
| trimAudioFiles    | Integer                                                                             | 1 表示是的，修剪音频文件。           |
| handleFrames      | Integer                                                                             | 句柄帧数（从 0 到 1000）。           |
| includePan        | Integer                                                                             | 1 意味着包括泛信息；0 意思是不要。   |

**返回**：如果成功则返回**0**。

### exportTimeline() 导出控制器插件导出序列

app.project.exportTimeline(exportControllerName)

**描述**：使用具有指定名称的导出控制器插件导出当前活动的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**

| 参数                 | 类型   | 描述                                                                                             |
| -------------------- | ------ | ------------------------------------------------------------------------------------------------ |
| exportControllerName | String | 要使用的导出控制器插件的名称。要使用 Premiere Pro SDK 示例导出控制器，该值应为“SDK 导出控制器”。 |

**返回**：如果成功则返回**0**，否则返回错误代码。

### getGraphicsWhiteLuminance() 检索项目当前图形白色亮度值

app.project.getGraphicsWhiteLuminance();

**描述**：检索此项目的当前图形白色亮度值。

**参数**：无。

**返回**：返回当前选定的图形白色值。

### getInsertionBin() 获取要导入素材箱项目

app.project.getInsertionBin()

**描述**：返回引用将发生导入 bin 的[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)。

**参数**：无。

**返回**：如果成功则返回[ProjectItem 对象，](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)否则返回**0**。

### getProjectPanelMetadata() 获取项目面板的当前布局

app.project.getProjectPanelMetadata()

**描述**：返回“项目”面板的当前布局。

**参数**：无。

**返回**：返回表示当前项目面板布局的**字符串**，如果不成功则返回**0**。

### getSharedLocation() 共享文件要复制到的路径

app.project.getSharedLocation()

**描述**：返回共享文件要复制到的位置的路径。

**参数**：无。

**返回**：返回一个包含路径的**字符串**。

### getSupportedGraphicsWhiteLuminances() 检索此项目支持的图形白色亮度值

app.project.getSupportedGraphicsWhiteLuminances();

**描述**：检索此项目支持的图形白色亮度值。

**参数**：无。

**返回**：返回项目支持的一组图形白色设置；目前它返回 (100, 203, 300)

### importAEComps() 导入 AE 合成

app.project.importAEComps(path, compNames, targetBin)

**描述**：从包含的 After Effects .aep 项目文件中导入指定的合成（按名称）。您可以在包含项目中指定一个目标 bin；否则，作品将出现在此项目中最近的目标箱中。

**参数**

| 参数      | 类型                                                                                  | 描述                                |
| --------- | ------------------------------------------------------------------------------------- | ----------------------------------- |
| path      | String                                                                                | After Effects .aep 项目文件的路径。 |
| compNames | Array                                                                                 | 要导入的指定项目中的合成名称。      |
| targetBin | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 可选的。此导入的目标 bin。          |

**返回**：如果成功则返回**0**。

### importAllAEComps() 导入 AE 合成

app.project.importAllAEComps(path, targetBin)

**描述**：从包含的 After Effects .aep 项目文件中导入指定的合成（按名称）。您可以在包含项目中指定一个目标 bin；否则，作品将出现在此项目中最近的目标箱中。

**参数**

| 参数      | 类型                                                                                  | 描述                                |
| --------- | ------------------------------------------------------------------------------------- | ----------------------------------- |
| path      | String                                                                                | After Effects .aep 项目文件的路径。 |
| targetBin | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 可选的。此导入的目标 bin。          |

**返回**：如果成功则返回**0**。

### importFiles() 导入文件

app.project.importFiles(filePaths, suppressUI, targetBin, importAsNumberedStills)

**描述**：从指定的文件路径导入媒体。

**参数**

| 参数                   | 类型                                                                                  | 描述                                       |
| ---------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------ |
| filePaths              | Array                                                                                 | 要导入的文件路径数组。                     |
| suppressUI             | Boolean                                                                               | 是否应抑制警告对话框。                     |
| targetBin              | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 应将文件导入到的 bin。                     |
| importAsNumberedStills | Boolean                                                                               | 文件路径是否应解释为一系列编号的静止图像。 |

**返回**：成功则返回**真，**否则返回**假**。

### importSequences() 导入序列

app.project.importSequences(path, sequenceIDs)

**描述**：将指定项目中的一系列[序列](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)对象（具有指定的序列 ID）导入当前项目。

**参数**

| 参数        | 类型   | 描述                   |
| ----------- | ------ | ---------------------- |
| path        | String | 项目文件的路径。       |
| sequenceIDs | Array  | 要导入的序列 ID 数组。 |

**返回**：如果成功则返回**0**。

### isSharedLocationCopyEnabled() 共享位置

app.project.isSharedLocationCopyEnabled()

**描述**：确定是否为此项目启用复制到共享位置。

**参数**：无。

**返回**：如果启用复制，则返回 **true**；如果不是，则为**假**。

### newBarsAndTone() 创建 Tone

app.project.newBarsAndTone(width, height, timeBase, PARNum, PARDen, audioSampleRate, name)

**描述**：根据指定的预设（.sqpreset 文件）创建具有给定名称的新[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)

**参数**

| 参数            | 类型    | 描述             |
| --------------- | ------- | ---------------- |
| width           | Integer |                  |
| height          | Integer |                  |
| timeBase        |         | 新项目项的时基。 |
| PARNum          | Integer | 像素纵横比分子。 |
| PARDen          | Integer | 像素纵横比分母。 |
| audioSampleRate |         | 音频采样率。     |
| name            | String  | 新项目项的名称。 |

**返回**：为新的小节和音调返回一个[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)，如果不成功则返回**0**。

### newSequence() 创建序列

app.project.newSequence(name, pathToSequencePreset)

**描述**：根据指定的预设（.sqpreset 文件）创建具有给定名称的新[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**

| 参数                 | 类型   | 描述                        |
| -------------------- | ------ | --------------------------- |
| name                 | String | 新序列的名称。              |
| pathToSequencePreset | String | 预设 .sqpreset 文件的路径。 |

**返回**：返回一个[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)，如果不成功则返回**0**。

### openSequence() 打开序列

app.project.openSequence(sequence.sequenceID)

**描述**：使具有提供的序列 ID 的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)处于活动状态。这将在“时间轴”面板中打开序列。

**参数**

| 参数       | 类型                                                                                                    | 描述                  |
| ---------- | ------------------------------------------------------------------------------------------------------- | --------------------- |
| sequenceID | [序列号.sequenceID](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence-sequenceid) | 应打开的有效序列 ID。 |

**返回**：成功则返回**真，**否则返回**假**。

### pauseGrowing() 暂停文件增长

app.project.pauseGrowing(pause)

**描述**：暂停（并恢复）不断增长的文件捕获。

**参数**

| 参数  | 类型    | 描述                           |
| ----- | ------- | ------------------------------ |
| pause | Integer | 如果 1，则启用不断增长的文件。 |

**返回**：如果成功则返回**0**。

### save() 保存工程

app.project.save()

**描述**：以当前路径保存项目。

**参数**：无。

**返回**：如果成功则返回**0**。

### saveAs() 另存为

app.project.saveAs(path)

**描述**：将当前项目导出到新的唯一文件路径，从新位置打开项目，并关闭之前打开的（和相同的）项目。

**参数**

| 参数 | 类型   | 描述           |
| ---- | ------ | -------------- |
| path | String | 新文件的路径。 |

**返回**：如果成功则返回**0**，否则返回错误代码。

### setEnableTranscodeOnIngest() 转码摄取

app.project.setEnableTranscodeOnIngest(state);

**描述**：控制对给定项目启用转码摄取行为。

**参数**

| 参数  | 类型    | 描述         |
| ----- | ------- | ------------ |
| state | Boolean | 想要的状态。 |

**返回**：如果成功则返回**真**。

### setGraphicsWhiteLuminance() 设置图形的白色亮度值

app.project.setGraphicsWhiteLuminance(value)

**描述**：为该项目设置当前图形的白色亮度值。

**参数**

| 参数  | 类型    | 描述                                                                                                                                                                                  |
| ----- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| value | Integer | 要使用的值；必须是[Project.getSupportedGraphicsWhiteLuminances()](https://ppro-scripting.docsforadobe.dev/general/project.html#project-getsupportedgraphicswhiteluminances)提供的值。 |

**返回**：如果成功则返回真。

### setProjectPanelMetadata() 设置布局

app.project.setProjectPanelMetadata(layout)

**描述**：返回“项目”面板的当前布局。

**参数**

| 参数   | 类型   | 描述                                                                                                                                                                                                                              |
| ------ | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| layout | String | 表示所需的项目面板布局。注意：生成有效布局字符串的唯一已知方法是根据需要设置项目面板，然后使用[Project.getProjectPanelMetadata()](https://ppro-scripting.docsforadobe.dev/general/project.html#project-getprojectpanelmetadata)。 |

**返回**：如果不成功则返回 **0**。

### setScratchDiskPath() 设置暂存盘

app.project.setScratchDiskPath(newPath, whichScratchDiskPath)

**描述**：将指定的暂存盘路径更改为新路径。

**参数**

| 参数            | 类型   | 描述                                                               |
| --------------- | ------ | ------------------------------------------------------------------ |
| newPath         | String | 新路径                                                             |
| scratchDiskType | Enum   | 枚举值，必须是以下之一：\* ScratchDiskType.FirstVideoCaptureFolder |

- ScratchDiskType.FirstAudioCaptureFolder
- ScratchDiskType.FirstVideoPreviewFolder
- ScratchDiskType.FirstAudioPreviewFolder
- ScratchDiskType.FirstAutoSaveFolder
- ScratchDiskType.FirstCCLibrariesFolder
- ScratchDiskType.FirstCapsuleMediaFolder
  |

- ScratchDiskType.FirstAudioCaptureFolder
- ScratchDiskType.FirstVideoPreviewFolder
- ScratchDiskType.FirstAudioPreviewFolder
- ScratchDiskType.FirstAutoSaveFolder
- ScratchDiskType.FirstCCLibrariesFolder
- ScratchDiskType.FirstCapsuleMediaFolder

**返回**：如果不成功则返回 **0**。
