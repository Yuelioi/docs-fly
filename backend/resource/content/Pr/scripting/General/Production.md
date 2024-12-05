---
title: 生产对象 Production
order: 8
---

---
title: 生产对象 Production
order: 6
category:

- PR脚本

---

## 生产对象

app.production

**描述**：Production 对象允许 ExtendScript 访问和操作作品、插入项目、创建新项目和素材箱，以及将现有的 Production 项目移至废纸篓。

## 属性

### name 生产名称

app.production.name

**描述**：产品名称。

**类型**：字符串。

### path 生产路径

app.production.path

**描述**：生产文件夹的路径。

**类型**：字符串。

### projects 生产项目

app.production.projects

**描述**：生产中包含的一系列项目，这些项目当前处于打开状态。不包括非开放项目。

**类型**：[ProjectCollection 对象](https://ppro-scripting.docsforadobe.dev/collection/projectcollection.html#projectcollection)，只读。

## 方法

### addProject()

app.production.addProject(srcProjectPath, destProjectPath)

**描述**：将项目从其他位置复制到 Production 目录中。

**参数**

| 参数            | 类型   | 描述                   |
| --------------- | ------ | ---------------------- |
| srcProjectPath  | String | 源项目的路径。         |
| destProjectPath | String | 添加的项目的目标路径。 |

**返回**：如果成功则返回**真**。

### close()

app.production.close()

**描述**：关闭生产以及该生产中所有打开的项目。

**参数**：无。

**返回**：如果成功则返回 true。

### getLocked()

app.production.getLocked()

**描述**：返回生产的当前锁定状态。

**参数**：无。

**返回**：如果生产被锁定，则返回**true**，如果它被解锁，则返回**false**。

### moveToTrash()

app.production.moveToTrash(projectOrFolderPath, suppressUI, saveProject)

**描述**：将指定的路径（“bin”）或 .prproj 移动到 Production 的 Trash 文件夹中。

**参数**

| 参数                | 类型    | 描述                         |
| ------------------- | ------- | ---------------------------- |
| projectOrFolderPath | String  | 源项目的路径。               |
| suppressUI          | Boolean | 是否抑制任何由此产生的对话。 |
| saveProject         | Boolean | 是否先保存项目。             |

**返回**：如果成功则返回**真**。

### setLocked()

app.production.setLocked(locked)

**描述**：设置生产的锁定状态

**参数**

| 参数   | 类型    | 描述             |
| ------ | ------- | ---------------- |
| locked | Boolean | 所需的锁定状态。 |

**返回**：如果成功则返回**true**。
