---
title: ProjectCollection 项目集合对象
order: 3
---

---
title: ProjectCollection 项目集合对象
order: 5
category:

- PR脚本

---

## 项目集合对象

app.projects
app.production.projects

ProjectCollection 对象表示[Project 对象](https://ppro-scripting.docsforadobe.dev/general/project.html#project)的集合。

ProjectCollection 是[Collection 对象](https://ppro-scripting.docsforadobe.dev/collection/collection.html#collection)的子类。使用 ProjectCollection 时，除了下面列出的方法和属性之外，Collection 的所有方法和属性都可用。

## 属性

### numProjects

app.projects.numProjects
app.production.projects.numProjects

**描述**：在“项目”面板中找到的项目和作品总数。

**类型**：整数，只读。

示例：查看当前 PR 打开了几个项目

```javascript
alert(app.projects.numProjects.toString()); // 2 （打开了俩）
```
