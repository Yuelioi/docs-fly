---
title: RenderQueue
order: 3
---

# RenderQueue 渲染队列

## 渲染队列对象 这节暂时跳过

app.project.renderQueue

描述：渲染队列（RenderQueue）对象表示渲染自动化，从“渲染队列”面板获取的数据和方法。

- 通过属性可以访问渲染队列中的项目及其渲染状态。
- 通过方法可以启动，暂停和停止渲染过程。渲染队列项目（RenderQueueItem）对象可以访问要渲染的项目的特定设置。

## 属性

### canQueueInAME 是否有排队渲染项

app.project.renderQueue.canQueueInAME

::: info 提示
在 After Effects 14.0(CC 2017)中添加了此功能
:::

描述：指示 After Effects 渲染队列中是否有排队的渲染项。只能将排队的项目添加到 AME 队列中。

类型：布尔值，只读。

### items 渲染项目集

app.project.renderQueue.items

描述：渲染队列中所有项目的集合。请参见 RenderQueueItem 对象。

类型：RQItemCollection 对象; 只读。

### numItems 渲染项目总数

app.project.renderQueue.numItems

描述：渲染队列中的项目总数。

类型：整数; 只读。

### rendering 正在渲染

app.project.renderQueue.rendering

描述：如果为 true，则渲染正在进行或暂停。如果为 false，则停止。

类型：布尔值，只读。

## 方法

### item() 查找渲染项目

app.project.renderQueue.item(index)

描述：从 items 集合中获取指定的项目。

参数：index，项目的位置索引。

返回：RenderQueueItem 对象。

### pauseRendering() 暂停渲染

app.project.renderQueue.pauseRendering(pause)

描述：暂停当前渲染过程，或继续暂停的渲染过程。这与渲染队列面板单击“暂停”相同。可以从 RenderQueueItem.onStatusChanged 或 app.onError 回调此方法。

参数：pause 为 true，暂停当前渲染过程。为 false，则继续暂停的渲染。

返回：无。

### render() 开始渲染

app.project.renderQueue.render()

描述：开始渲染过程。这与在“渲染队列”面板中单击“渲染”相同。在渲染过程完成之前，该方法不会返回。要暂停或停止渲染过程，请从或回调调用 RenderQueue.pauseRendering()或 RenderQueue.stopRendering()。onErroronStatusChanged

要在渲染过程中响应错误，请在 app.onError 中定义一个回调函数。  
要在渲染进行过程中响应特定项目状态的变化，请在关联的 RenderQueueItem 对象的 RenderQueueItem.onStatusChanged 中定义一个回调函数。

参数：无。

返回：无。

### showWindow() 显示渲染面板

app.project.renderQueue.showWindow(doShow)

描述：显示或隐藏“渲染队列”面板。

参数：doShow 如果为 true，则显示“渲染队列”面板。如果为 false，则隐藏。

返回：无。

### stopRendering() 停止渲染

app.project.renderQueue.stopRendering()

描述：停止渲染过程。这与“渲染队列”面板单击“停止”相同。可以从 RenderQueueItem.onStatusChanged 或 app.onError 回调此方法。

参数：无。

返回：无。

### queueInAME() 调用 AME 队列

app.project.renderQueue.queueInAME(render_immediately_in_AME)

::: info 提示
在 After Effects 14.0(CC 2017)中添加了此功能
:::

描述：调用 AME 中的 Queue 命令。此方法需要传递一个布尔值，告诉 AME 是进行排队(false)还是开始处理队列(true)。

:::tip
当 AME 收到排队的项目时，它将应用最近使用的编码预设。如果 render_immediately_in_AME 设置为 true，则更改编码设置。
:::

参数：render_immediately_in_AME 告诉 AME 是只对渲染项目排队(false)还是 AME 也应该开始处理其队列(true)。

返回：无。

示例：检查渲染队列中是否有排队的项目，如果有，则将它们发送到 AME 中排队，但不会立即开始渲染：

```javascript
// 在AME中添加渲染队列，需要AME11.0以上
if (app.project.renderQueue.canQueueInAME === true) {
  // 把序列发送到AME排队，但是不开始渲染
  app.project.renderQueue.queueInAME(false);
} else {
  alert("There are no queued item in the Render Queue.");
}
```
