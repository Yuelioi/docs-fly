---
title: AME 编码器对象
order: 1
---

---
title: AME 编码器对象
order: 3
category:

- PR脚本

---

app.encoder

**描述**：**编码器**对象代表 Adobe Media Encoder，以及用于 Premiere Pro 之外本地渲染。

## 属性

无。

## 方法

### encodeFile() AME 渲染文件

app.encoder.encodeFile(filePath, outputPath, presetPath, workArea, removeUponCompletion)

**描述**：使用指定设置使用 Adob​​e Media Encoder 渲染指定文件。

**参数**

| 参数       | 类型   | 描述                                           |
| ---------- | ------ | ---------------------------------------------- |
| filePath   | 字符串 | 要渲染的文件的路径。                           |
| outputPath | 字符串 | 输出文件的路径。                               |
| presetPath | 字符串 | 预设 (.epr) 文件的路径。                       |
| workArea   | 整数   | 要使用的渲染区间：\* 0 - ENCODE_ENTIRE（整个） |

- 1 - ENCODE_IN_TO_OUT（入点到出点）
- 2 - ENCODE_WORK_AREA（工作区）
  |
  | removeUponCompletion | 整数 | 如果为 1，作业将在完成后被删除。 |
  | inPoint | 时间 | 一个 **Time** 对象，用于新文件的入点。 |
  | outPoint | 时间 | 一个 **Time** 对象，用于新文件的出点。 |

**返回**：对于添加到 AME 队列的渲染作业，将作业 ID 作为**String**返回，如果不成功，则返回**0**。

示例：使用指定预设渲染一个文件

```javascript
app.encoder.launchEncoder(); // 启动AME

// 设置要渲染的文件以及预设文件

var video = File("C:\\Users\\YL\\Desktop\\test.mp4");

var preset = File("C:\\Users\\YL\\Documents\\Adobe\\Adobe Media Encoder\\14.0\\Presets\\IG.epr");

//添加文件到渲染队列

app.encoder.encodeFile(video.fsNmae, preset.fsNmae, 0, true);

app.encoder.startBatch(); // 开始渲染
```

### encodeProjectItem() AME 渲染项目

app.encoder.encodeProjectItem(projectItem, outputPath, presetPath, workArea, removeUponCompletion)

**描述**：使用指定设置使 Adob​​e Media Encoder 呈现（可选地，指定范围内）指定的[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)。

**参数**

| 参数        | 类型                                                                                  | 描述                                           |
| ----------- | ------------------------------------------------------------------------------------- | ---------------------------------------------- |
| projectItem | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 要渲染的项目。                                 |
| outputPath  | String                                                                                | 输出文件路径。                                 |
| presetPath  | String                                                                                |  (.epr) 预设文件路径。                         |
| workArea    | Integer                                                                               | 要使用的渲染区间：\* 0 - ENCODE_ENTIRE（整个） |

- 1 - ENCODE_IN_TO_OUT（入点到出点）
- 2 - ENCODE_WORK_AREA（工作区）
  |
  | removeUponCompletion | Integer | 如果 1，作业将在完成后被删除。0 则保留队列 |

**返回**：返回工作 ID ，**String**，如果不成功，则返回**0**。

示例：渲染项目面板第 2 个视频文件（不能是序列）

![](https://cdn.yuelili.com/20211028151608.png)

```javascript
var proj = app.project.rootItem.children[1]; // 定义要渲染的项目

// alert(proj.name) // 123.mp4

app.encoder.encodeProjectItem(
  proj,
  "C:\\Users\\YL\\Desktop\\test.mp4",
  "C:\\Users\\YL\\Documents\\Adobe\\Adobe Media Encoder\\14.0\\Presets\\我的预设.epr",
  0,
  1
);
```

### encodeSequence() AME 渲染序列

app.encoder.encodeSequence(sequence, outputPath, presetPath, workArea, removeUponCompletion)

**描述**：使用 Adob​​e Media Encoder 指定设置呈现指定的[Sequence 对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**参数**

| 参数       | 类型                                                                                | 描述                                           |
| ---------- | ----------------------------------------------------------------------------------- | ---------------------------------------------- |
| sequence   | [序列对象](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence) | 要渲染的序列。                                 |
| outputPath | String                                                                              | 输出文件的路径。                               |
| presetPath | String                                                                              | 预设 (.epr) 文件的路径。                       |
| workArea   | Integer                                                                             | 要使用的渲染区间：\* 0 - ENCODE_ENTIRE（整个） |

- 1 - ENCODE_IN_TO_OUT（入点到出点）
- 2 - ENCODE_WORK_AREA（工作区）
  |
  | removeUponCompletion | Integer | 如果 1，作业将在完成后被删除。 |

**返回**：对于添加到 AME 队列的渲染作业，将作业 ID 作为**String**返回，如果不成功，则返回**0**。

示例：渲染活动序列

```javascript
var sequence = app.project.activeSequence; // 活动序列

app.encoder.launchEncoder(); // 启动AME

// 添加系列到渲染队列

app.encoder.encodeSequence(
  sequence,
  "C:\\Users\\YL\\Desktop\\test.mp4",
  "C:\\Users\\YL\\Documents\\Adobe\\Adobe Media Encoder\\14.0\\Presets\\MY Preset.epr",
  0,
  true
);

app.encoder.startBatch(); // 开始渲染
```

### launchEncoder() 启动 AME

app.encoder.launchEncoder()

**描述**：启动 Adob​​e Media Encoder。

**参数**：无。

**返回**：如果成功则返回**0**。

### setEmbeddedXMPEnabled() 没用过

app.encoder.setEmbeddedXMPEnabled(enabled)

**描述**：确定是否将输出嵌入的 XMP 元数据。

**参数**

| 参数    | 类型    | 描述                               |
| ------- | ------- | ---------------------------------- |
| enabled | Integer | 通过 1 启用 sidecar 输出，0 禁用。 |

**返回**：如果成功则返回**0**。

注意：Premiere Pro 和 Adob​​e Media Encoder 将输出某些文件格式的 sidecar XMP，并为大多数文件嵌入 XMP。应用程序基于多种因素做出此决定，对于通常使用“其他方法”的格式，没有 API 控制来“强制”sidecar 或嵌入式输出。

### setSidecarXMPEnabled() 没用过

app.encoder.setSidecarXMPEnabled(enabled)

**描述**：确定是否输出包含 XMP 元数据的 sidecar 文件。

**参数**

| 参数    | 类型    | 描述                               |
| ------- | ------- | ---------------------------------- |
| enabled | Integer | 通过 1 启用 sidecar 输出，0 禁用。 |

**返回**：如果成功则返回**0**。

### startBatch() AME 开始渲染

app.encoder.startBatch()

**描述**：Adob​​e Media Encoder 开始渲染其渲染队列。

**参数**：无。

**返回**：如果成功则返回**0**。
