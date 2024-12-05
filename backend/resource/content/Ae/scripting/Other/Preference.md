---
title: Preference
order: 5
---

# Preference 首选项

## app.preferences

描述：首选项对象，保存在首选项文件中（C:\Users\你的用户名\AppData\Roaming\Adobe\After
Effects\版本号），并且在 AE 关闭之前保持不变。

或者菜单栏 - 编辑 - 首选项 - 常规：在资源管理器中显示首选项

![](https://cdn.yuelili.com/20211021115157.png)

![](https://cdn.yuelili.com/20211021115314.png)

::: info 提示
首选项文件不同 AE 版本、不同 AE 软件语言，均不相同。
:::

文件中的节（section）和键（key）标识，并且每个键名都与一个值相关联。可以使用该对象创建新的首选项，也可以访问现有首选项。

从 12 / CC 版本开始，如果 Section /
Key 不在"Prefs.txt"文件中，则使用第三个参数来指定首选项文件。如果没有第三个参数，则使用默认值(PREFType.PREF_Type_MACHINE_SPECIFIC)
"Prefs.txt"文件

第三个参数：PREFType，枚举值：

| 名称                                      | 文件名                      | 主要功能（重要程度排序）                                 |
| ----------------------------------------- | --------------------------- | -------------------------------------------------------- |
| PREF_Type_MACHINE_SPECIFIC                | Prefs.txt                   | 磁盘缓存、表达式（设置、关键词、高亮）\*属于核心预设内容 |
| PREF_Type_MACHINE_INDEPENDENT             | Prefs-indep-general.txt     | 自动保存与标签等                                         |
| PREF_Type_MACHINE_INDEPENDENT_RENDER      | Prefs-indep-render.txt      | 渲染导出设置                                             |
| PREF_Type_MACHINE_INDEPENDENT_OUTPUT      | Prefs-indep-output.txt      | 渲染导出设置                                             |
| PREF_Type_MACHINE_INDEPENDENT_COMPOSITION | Prefs-indep-composition.txt | 创建合成                                                 |
| PREF_Type_MACHINE_SPECIFIC_TEXT           | Prefs-text.txt              |                                                          |
| PREF_Type_MACHINE_SPECIFIC_PAINT          | Prefs-paint.txt             | Clone                                                    |

#### 如何查找

软件使用的：File Seek（其他能搜 txt 的都行）

**1.搜索大类名/Selection**

一般大类用 XX Selection：比如 General Selection

![](https://cdn.yuelili.com/20210930125042.png)

**2.搜索组名或者关键词**

以允许脚本读取网络与调试脚本为例

![](https://cdn.yuelili.com/20210930125514.png)

搜索 Scripting

![](https://cdn.yuelili.com/20210930125344.png)

搜索 Debugger

![](https://cdn.yuelili.com/20210930125449.png)

## 方法

### deletePref()

app.preferences.deletePref(sectionName, keyName[, prefType])

描述：从首选项文件中删除首选项。

::: danger
通常 **不建议** 删除内部 AE 首选项，但是可以删除保存的 Settings 对象。请注意，需要先添加"Settings\_"节名
:::

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：无。

示例："Precomp Cropper"节保存键名为"trimPrecomps"的设置，则可以通过以下方式删除该设置：

```javascript
app.preferences.deletePref("Settings_Precomp Cropper", "trimPrecomps");
```

### getPrefAsBool()

app.preferences.getPrefAsBool(sectionName, keyName[, prefType])

描述：从首选项文件中检索首选项值，并解析为布尔值。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：布尔值。

示例 1："默认情况下扩展流程图比较"流程图的值

```javascript
var expandByDefault = app.preferences.getPrefAsBool("Flowchart Settings", "Expand Flowchart Comps by Default");
alert("The setting is: " + expandByDefault);
```

示例 2："已启用 Javascript 调试器"：

```javascript
var debuggerEnabled = app.preferences.getPrefAsBool("Main Pref Section v2", "Pref_JAVASCRIPT_DEBUGGER", PREFType.PREF_Type_MACHINE_INDEPENDENT);
alert("当前设置为: " + debuggerEnabled);
```

### getPrefAsFloat()

app.preferences.getPrefAsFloat(sectionName, keyName[, prefType])

描述：从首选项文件中检索首选项值，并解析为浮点数。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：浮点值。

示例：是否允许脚本连接网络

```javascript
var ff = app.preferences.getPrefAsFloat("Main Pref Section v2", "Pref_SCRIPTING_FILE_NETWORK_SECURITY", PREFType.PREF_Type_MACHINE_SPECIFIC);
alert(ff); // 返回true
```

### getPrefAsLong()

app.preferences.getPrefAsLong(sectionName, keyName[, prefType])

描述：从首选项文件中检索首选项值，并解析为长整型(数字)。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：长整型。

### getPrefAsString()

app.preferences.getPrefAsString(sectionName, keyName[, prefType])

描述：从首选项文件中检索首选项值，并将其解析为字符串。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：字符串。

示例：查找导入静止素材默认持续时间

![](https://cdn.yuelili.com/20210930130927.png)

```javascript
var pree = app.preferences.getPrefAsString("Main Pref Section v2", "Pref_DEFAULT_STILL_OUT_POINT v2", PREFType.PREF_Type_MACHINE_INDEPENDENT);
alert(pree); // 返回 57/30
```

### havePref()

app.preferences.havePref(sectionName, keyName[, prefType])

描述：如果指定的首选项存在并且有值，则返回 true。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- prefType：可选，枚举值，首选项文件

返回：布尔值。

```javascript
var boo = app.preferences.havePref("Main Pref Section v2", "Pref_DEFAULT_STILL_OUT_POINT v2", PREFType.PREF_Type_MACHINE_INDEPENDENT);
alert(boo); // 有则返回true
```

### reload()

app.preferences.reload()

描述：手动重载首选项文件。否则只有在重启 AE 后，才能通过脚本访问对首选项所做的更改。

参数：无。

返回：无。

### savePrefAsBool()

app.preferences.savePrefAsBool(sectionName, keyName, value[, prefType])

描述：将首选项保存为布尔值。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- value：新值，布尔值
- prefType：可选，枚举值，首选项文件

返回：无。

### savePrefAsFloat()

app.preferences.savePrefAsFloat(sectionName, keyName, value[, prefType])

描述：将首选项保存为浮点值。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- value：新值，浮点值
- prefType：可选，枚举值，首选项文件

返回：无。

### savePrefAsLong()

app.preferences.savePrefAsLong(sectionName, keyName, value[, prefType])

描述：将首选项保存为长整型数字。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- value：新值，长整型，数字
- prefType：可选，枚举值，首选项文件

返回：无。

### savePrefAsString()

app.preferences.savePrefAsString(sectionName, keyName, value[, prefType])

描述：将首选项保存为字符串。

参数：

- sectionName：节名，字符串
- keyName ：键名，字符串
- value：新值，字符串
- prefType：可选，枚举值，首选项文件

返回：无。

### saveToDisk()

app.preferences.saveToDisk()

描述：手动将首选项保存到磁盘。否则只有在重启 AE 后，才能通过脚本访问对首选项所做的更改。

参数：无。

返回：无。
