---
title: 项目管理器对象 projectManager
order: 9
---

---
title: 项目管理器对象 projectManager
order: 8
category:

- PR脚本

---

## 项目管理器对象

app.projectManager.options

**描述**：ProjectManager 对象公开 Premiere Pro 的项目管理器，用于项目合并、传输和转码。

## 属性

### affectedSequences

app.projectManager.options.affectedSequences

**描述**：一个阵列的[顺序](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)的对象，要导出。

**类型**：数组; 读/写。

### clipTranscoderOption

app.projectManager.options.clipTranscoderOption

**描述**：剪辑转码的指定设置。值将是以下之一：

| CLIP_TRANSCODE_MATCH_PRESET | 使用指定的预设进行转码。 |
| CLIP_TRANSCODE_MATCH_CLIPS | 匹配剪辑 |
| CLIP_TRANSCODE_MATCH_SEQUENCE| 必须是以下之一： |

**类型**：字符串; 读/写。

### clipTransferOption

app.projectManager.options.clipTransferOption

**描述**：剪辑传输的指定设置。值将是以下之一：

| CLIP_TRANSFER_COPY | 复制整个源媒体。 |
| CLIP_TRANSFER_TRANSCODE | 转码为默认输出格式。 |

### convertAECompsToClips

app.projectManager.options.convertAECompsToClips

**描述**：如果为 true，则将动态链接的 After Effects 合成渲染到新媒体（使用指定的输出预设）。

**类型**：布尔值；读/写。

### convertImageSequencesToClips

app.projectManager.options.convertImageSequencesToClips

**描述**：如果为 true，则将图像序列转码为新媒体（使用指定的输出预设）。

**类型**：布尔值；读/写。

### convertSyntheticsToClips

app.projectManager.options.convertSyntheticsToClips

**描述**：如果为 true，则将合成导入器中的剪辑转码为新媒体（使用指定的输出预设）。

**类型**：布尔值；读/写。

### copyToPreventAlphaLoss

app.projectManager.options.copyToPreventAlphaLoss

**描述**：如果为 true，则将任何可用的 alpha 信息包含到转码媒体中。

**类型**：布尔值；读/写。

### destinationPath

app.projectManager.options.destinationPath

**描述**：导出项目和媒体的路径。

**类型**：字符串; 读/写。

### encoderPresetFilePath

app.projectManager.options.encoderPresetFilePath

**描述**：要使用的输出预设（.epr 文件）的路径。

**类型**：字符串; 读写。

### excludeUnused

app.projectManager.options.excludeUnused

**描述**：如果非零，则从导出的项目中排除未使用的项目项。

**类型**：布尔值；读/写。

### handleFrameCount

app.projectManager.options.handleFrameCount

**描述**：要包括多少帧“处理”素材（进出点前后）。

**类型**：整数; 读/写。

### includeAllSequences

app.projectManager.options.includeAllSequences

**描述**：如果为 true，则导出导出项目中的所有[序列](https://ppro-scripting.docsforadobe.dev/sequence/sequence.html#sequence)。

**类型**：布尔值；读/写。

### includeConformedAudio

app.projectManager.options.includeConformedAudio

**描述**：如果为 true，则在导出的项目中包含符合要求的音频文件。

**类型**：布尔值；读/写。

### includePreviews

app.projectManager.options.includePreviews

**描述**：如果为 true，则在导出的项目中包含渲染的预览文件。

**类型**：布尔值；读/写。

### renameMedia

app.projectManager.options.renameMedia

**描述**：如果为 true，则在导出过程中执行重命名。

**类型**：布尔值；读/写。
