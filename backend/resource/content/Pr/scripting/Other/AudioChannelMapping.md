---
order: 4
title: 音频通道映射对象 AudioChannelMapping
---

---
title: 音频通道映射对象 AudioChannelMapping
order: 2
category:

- PR脚本

---

## 音频通道映射对象

app.project.rootItem.children[index].getAudioChannelMapping

**描述**：AudioChannelMapping 对象定义应用于给定[ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)的音频通道映射。

## 属性

### AudioChannelMapping.audioChannelsType

app.project.rootItem.children[index].getAudioChannelMapping.audioChannelsType

**描述**：此频道中包含的音频类型。将是 0、1 或 2，对应于 AUDIOCHANNELTYPE_Mono、AUDIOCHANNELTYPE_Stereo、 或 AUDIOCHANNELTYPE_51。

### AudioChannelMapping.audioClipsNumber

app.project.rootItem.children[index].getAudioChannelMapping.audioClipsNumber

**描述**：与此音频通道关联的音频剪辑的数量。

## 方法

### AudioChannelMapping.setMappingForChannel()

app.project.rootItem.children[index].setMappingForChannel(channelIndex, sourceChannelIndex)

**描述**：将源通道映射到指定的通道索引。

**参数**

| 参数               | 类型    | 描述                   |
| ------------------ | ------- | ---------------------- |
| channelIndex       | Integer | 要映射的通道的索引。   |
| sourceChannelIndex | Integer | 要映射的源通道的索引。 |

**返回**：如果成功则返回**true**，如果该映射不受支持则返回**false**。
