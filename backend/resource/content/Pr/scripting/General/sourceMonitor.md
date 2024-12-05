---
title: 源监视器对象 sourceMonitor
order: 7
---

---
title: 源监视器对象 sourceMonitor
order: 10
category:

- PR脚本

---

## 源监视器对象

app.sourceMonitor

**描述**：Source 对象代表 Premiere Pro 的源监视器。

## 属性

无。

## 方法

### closeAllClips()

app.sourceMonitor.closeAllClips()

**描述**：关闭源监视器中的所有剪辑。

**参数**：无。

**返回**：如果成功则返回**0**。

### closeClip()

app.sourceMonitor.closeClip()

**描述**：关闭源监视器中最前面的剪辑。

**参数**：无。

**返回**：如果成功则返回**0**。

### getPosition()

app.sourceMonitor.getPosition()

**描述**：检索源监视器的当前时间指示器的位置。

**参数**：无。

**返回**：返回一个[Time 对象](https://ppro-scripting.docsforadobe.dev/other/time.html#time)，其中包含源监视器的当前时间指示器的位置。

### openFilePath()

app.sourceMonitor.openFilePath(path)

**描述**：在源监视器中打开一个文件。

**参数**

| 参数 | 类型   | 描述                 |
| ---- | ------ | -------------------- |
| path | String | 要打开的文件的路径。 |

**返回**：true 成功则返回。

### openProjectItem()

app.sourceMonitor.openProjectItem(projectItem)

**描述**：在源监视器中打开一个项目项。

**参数**

| 参数        | 类型                                                                                  | 描述             |
| ----------- | ------------------------------------------------------------------------------------- | ---------------- |
| projectItem | [项目对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem) | 要打开的项目项。 |

**返回**：如果成功则返回 0。

### play()

app.sourceMonitor.play(playbackSpeed)

**描述**：以指定的播放速度开始播放源监视器。

**参数**

| 参数          | 类型  | 描述       |
| ------------- | ----- | ---------- |
| playbackSpeed | Float | 播放速度。 |

**返回**：如果成功则返回 0。
