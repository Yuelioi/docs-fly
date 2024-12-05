---
title: RenderQueueItem
order: 4
---

# RenderQueueItem 渲染序列项目

## 渲染序列项目

app.project.renderQueue.item(index)

描述：渲染序列项目（RenderQueueItem）对象表示渲染队列中的单个项目。可以访问要渲染项目的特定设置。通过将带有 RQItemCollection 对象的合成添加到“渲染队列”中，创建对象;
参见 RQItemCollection.add()。

## 属性

### comp 将渲染的合成

app.project.renderQueue.item(index).comp

描述：渲染队列项中将渲染的合成。若要更改合成，必须删除此渲染队列项，并创建一个新项。

类型：CompItem 对象; 只读。

### elapsedSeconds 渲染时长

app.project.renderQueue.item(index).elapsedSeconds

描述：渲染此项目花费的秒数。

类型：整数；如果尚未渲染项目，则为 null；只读。

### logType 渲染日志

app.project.renderQueue.item(index).logType

描述：此项目的日志类型，指示在呈现此项目时应记录哪些事件。

类型：LogType，枚举值; (读/写)：

- LogType.ERRORS_ONLY：仅错误
- LogType.ERRORS_AND_SETTINGS：增加设置
- LogType.ERRORS_AND_PER_FRAME_INFO：增加每帧信息

### numOutputModules 输出模块数量

app.project.renderQueue.item(index).numOutputModules

描述：分配给该项目的输出模块的总数量。

类型：整数; 只读。

### onStatusChanged 监控渲染回调函数

app.project.renderQueue.item(index).onStatusChanged

描述：每当 RenderQueueItem.status 属性的值更改时调用的回调函数。

在渲染进行中或暂停时，您无法更改渲染队列项或应用程序；但是，您可以使用此回调来暂停或停止渲染过程。请参阅 RenderQueue.pauseRendering()和 RenderQueue.stopRendering()。另请参阅 app.onError。

类型：函数名称，字符串，如果未分配任何函数，则为 null。

示例：

```javascript
function myStatusChanged() {
  alert(app.project.renderQueue.item(1).status);
}
app.project.renderQueue.item(1).onStatusChanged = myStatusChanged();
app.project.renderQueue.item(1).render = false; // 更改状态并显示日志
```

### outputModules 输出模块

app.project.renderQueue.item(index).outputModules

描述：该项目的输出模块的集合。

类型：OMCollection 对象; 只读。

### render 项目渲染

app.project.renderQueue.item(index).render

描述：如果为 true，则在启动渲染队列时渲染该项目，此时，RenderQueueItem.status 设置为 RQItemStatus.QUEUED。设置为 false 时，status 设置为
RQItemStatus.UNQUEUED。

类型：布尔值 读/写。

### skipFrames 跳帧

app.project.renderQueue.item(index).skipFrames

描述：渲染此项目时要跳过的帧数。用于渲染测试。值为 0，将不跳帧进行常规渲染。值为 1，每隔一帧跳 1 次。这等效于“双重渲染”。总时间长度保持不变。例如，如果值为 1，输出原有帧数的一半，但每个帧持续时间翻倍。

类型：范围内的整数 0~99; 读/写。

### startTime 起始渲染时间

app.project.renderQueue.item(index).startTime

描述：此项目开始渲染的日期和时间。

类型：Date 对象；如果该项尚未开始呈现，则为 null；只读。

### status 渲染状态

app.project.renderQueue.item(index).status

描述：项目的当前渲染状态。

类型：RQItemStatus，枚举值; 只读：

- RQItemStatus.WILL_CONTINUE：渲染过程已暂停。
- RQItemStatus.NEEDS_OUTPUT：项目缺少有效的输出路径。
- RQItemStatus.UNQUEUED：项目在“渲染队列”面板中列出，但合成尚未准备好进行渲染。
- RQItemStatus.QUEUED：合成准备就绪。
- RQItemStatus.RENDERING：合成正在渲染
- RQItemStatus.USER_STOPPED：渲染过程已被用户或脚本停止。
- RQItemStatus.ERR_STOPPED：由于错误，渲染过程已停止。
- RQItemStatus.DONE：项目的渲染过程已完成。

### templates 渲染模板

app.project.renderQueue.item(index).templates

描述：该项目可用的所有“渲染设置”模板的名称。另请参见 RenderQueueItem.saveAsTemplate()。

类型：字符串数组；只读。

### timeSpanDuration 渲染合成时长

app.project.renderQueue.item(index).timeSpanDuration

描述：要渲染的合成的持续时间(秒)。通过从结束时间中减去开始时间来确定持续时间。与“渲染设置>自定义>结束时间"相同。

类型：浮点值；读/写。

### timeSpanStart 开始渲染时间

app.project.renderQueue.item(index).timeSpanStart

描述：合成中开始渲染的时间(秒)。与“渲染设置>自定义>开始时间"相同。

类型：浮点值；读/写。

## 方法

### applyTemplate() 应用渲染模板

app.project.renderQueue.item(index).applyTemplate(templateName)

描述：将“渲染设置”模板应用于该项目。另请参见 RenderQueueItem.saveAsTemplate()和 RenderQueueItem.templates。

参数：templateName，字符串，要应用的模板名称。

返回：无。

### duplicate() 复制渲染项目

app.project.renderQueue.item(index).duplicate()

描述：创建此项的副本并将其添加到此渲染队列。

::: info 提示
复制状态为“完成”的项目，会将新项目的状态设置为“已排队”。
:::

参数：无。

返回：RenderQueueItem 对象。

### getSetting() 获取渲染设置

app.project.renderQueue.item(index).getSetting()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：获取特定的“渲染队列项目”设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment =dva

存档版本：https://web.archive.org/web/20200622100656/https:
//blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment
=dva

示例

```javascript
/获取渲染设置的“代理使用”的当前值
//关键键和值均为英文字符串。
var rqItem1_proxyUse = app.project.renderQueue.item(1).getSetting("Proxy Use");

// 获取相同设置的字符串版本，在字符串的末尾添加“ -str”
var rqItem1_proxyUse_str = app.project.renderQueue.item(1).getSetting("Proxy Use-str");
```

### getSettings() 获取渲染设置

app.project.renderQueue.item(index).getSettings()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：获取给定渲染队列项的所有设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment =dva

存档版本：https://web.archive.org/web/20200622100656/https:
//blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment
=dva

示例

//获取包含以下对象的所有渲染设置的所有可能值的对象  
//呈现队列项目 1 并转换为 JSON 格式。

```javascript
var rqItem1_spec_str = app.project.renderQueue.item(1).getSettings(GetSettingsFormat.SPEC);
var rqItem1_spec_str_json = rqItem1_spec_str.toSource();
```

### outputModule() 输出模块

app.project.renderQueue.item(index).outputModule(index)

描述：获取具有指定索引位置的输出模块。

参数：index，输出模块的位置索引。

返回：OutputModule 对象。

### remove() 移除渲染项目

app.project.renderQueue.item(index).remove()

描述：从渲染队列中删除此项目。

参数：无。

返回：无。

### saveAsTemplate() 另存渲染模板

app.project.renderQueue.item(index).saveAsTemplate(name)

描述：将项目的当前渲染设置，另存为具有指定名称的新模板。

参数：name 包含新模板名称的字符串。

返回：无。

### setSetting() 渲染设置

app.project.renderQueue.item(index).setSetting()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：为给定的“渲染队列项目”设置特定设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment =dva

存档版本：https://web.archive.org/web/20200622100656/https:
//blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment
=dva

示例：将“代理使用”的值设置为“使用所有代理”

```javascript
app.project.renderQueue.item(1).setSetting("Proxy Use", "Use ll Proxies");

//您也可以使用数字。
//这两个示例效果相同。

app.project.renderQueue.item(1).setSetting("Proxy Use", 1);
```

### setSettings() 渲染设置

app.project.renderQueue.item(index).setSettings()

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：为给定的“渲染队列项目”设置多个设置。

折旧资源：https://blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment =dva

存档版本：https://web.archive.org/web/20200622100656/https:
//blogs.adobe.com/creativecloud/new-changed-after-effects-cc-2014/? segment
=dva

示例

```javascript
//获取一个包含可设置渲染设置的字符串版本的对象
//呈现队列项目 1 的值。
//要获取数字格式的值，请使用 GetSettingsFormat.NUMBER_SETTABLE 作为参数。

var rqItem1_settable_str = app.project.renderQueue.item(1).getSettings(GetSettingsFormat.STRING_SETTABLE);

// Set render queue item 2 with values that you got from render
//queue item 1.

app.project.renderQueue.item(2).setSettings(rqItem1_settable_str);

// Set render queue item 3 with values you create.

var my_renderSettings = {
  "Color Depth": "32 bits per channel",
  Quality: "Best",
  Effects: "All On",
  "Time Span Duration": "1.0",
  "Time Span Start": "2.0",
};

app.project.renderQueue.item(2).setSettings(my_renderSettings);
```
