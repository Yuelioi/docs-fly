---
title: RqitemCollection
order: 5
---

# RQItemCollection 渲染项目集

## 渲染队列项目集

app.project.renderQueue.items

描述：渲染队列集（RQItemCollection）包含项目中的所有渲染队列项，如项目的“渲染队列”面板中所示。该集合提供对 RenderQueueItem 对象的访问，并允许您从合成中创建它们。集合中的第一个 RenderQueueItem 对象索引位置为 1。

RQItemCollection 是 Collection 对象的子类。使用 RQItemCollection 时，Collection 的所有方法和属性均可用。

## 方法

### add() 添加到渲染队列

app.project.renderQueue.items.add(comp)

描述：将合成添加到“渲染队列”，创建一个 RenderQueueItem。

参数：comp，要添加的合成，为 CompItem 对象。

返回：RenderQueueItem 对象。
