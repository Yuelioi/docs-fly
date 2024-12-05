---
title: Audio-Data-Structures
order: 3
---

# 音频数据结构体

以下数据类型被 After Effects 用来描述音频数据。

| **Structure**  | **Description** |
| -------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `PF_SoundFormat` | 音频是无符号脉冲编码调制 (PCM)、有符号 PCM 还是浮点格式。|
| `PF_SoundSampleSize` | 采样是 1、2 或 4 字节格式。 |
| `PF_SoundChannels` | 音频是单声道还是立体声。  |
| `PF_SoundFormatInfo` | 包含它所引用的音频的采样率、通道数、样本大小和格式。|
| `PF_SoundWorld`  | 使用`PF_SoundWorlds`来表示音频。除了 `PF_SoundFormatInfo` 之外，它们还包含音频的长度和指向实际音频数据的指针。|

`PF_SoundFormat`, `PF_SoundSampleSize`, 和`PF_SoundChannels`都包含在`PF_SoundFormatInfo`中。

`PF_SoundWorlds`包含一个`PF_SoundFormatInfo`，以及进一步的特定实例信息。
