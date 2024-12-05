---
title: ProjectItemCollection 项目集合
order: 4
---

---
title: ProjectItemCollection 项目集合
order: 6
category:

- PR脚本

---

## ProjectItemCollection 对象

app.project.rootItem.children

ProjectItemCollection 对象代表的集合[的 ProjectItem 对象](https://ppro-scripting.docsforadobe.dev/item/projectitem.html#projectitem)的活动项目。

ProjectItemCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 ProjectItemCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numItems 项目总数

app.project.rootItem.children.numItems

**描述**：活动项目中的项目总数。

**类型**：整数，只读。

![](https://cdn.yuelili.com/20211028175950.png)

```javascript
alert(app.project.rootItem.children.numItems.toString()); // 项目面板共有4个项目
```
