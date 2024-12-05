---
title: FileSource
order: 1
---

# FileSource 文件源

## 文件源对象

app.project.item(index).mainSource  
app.project.item(index).proxySource

描述：文件源（FileSource）对象，导入的文件素材。

父级关系：FileSource 是素材源（footageSource）对象的子类。素材源的方法和属性均可用

## 属性

### file 资产文件

app.project.item(index).mainSource.file  
app.project.item(index).proxySource.file

描述：定义此资产文件的 ExtendScript File 对象。如果要更改值的话：

- 如果 FileSource 是 AVItem 的 proxySource，请调用 setProxy()或 setProxyWithSequence()。
- 如果 FileSource 是 footageItem 的 mainSource，请调用 replace()或 replaceWithSequence()。

类型：文件对象；只读。

示例：查找项目面板的文件（注意返回的是文件对象）

![](https://cdn.yuelili.com/20211012162834.png)

```javascript
var myFootage = app.project.item(2).mainSource;
alert(myFootage.file.toString());
```

### missingFootagePath 丢失素材路径

app.project.item(index).mainSource.missingFootagePath  
app.project.item(index).proxySource.missingFootagePath

描述：该资产缺少的素材的路径和文件名。另请参阅 AVItem.footageMissing。

类型：字符串; 只读。

示例：查找项目面板的丢失文件路径（跟项目面板旁边的值一样）

![](https://cdn.yuelili.com/20211012162311.png)

```javascript
var myFootage = app.project.item(2).mainSource;
alert(myFootage.missingFootagePath); //（注意返回的是字符串）
```

## 方法

### reload() 文件重加载 //用不了 不知道为啥

app.project.item(index).mainSource.file.mainSource.reload()

描述：从文件中重新加载资产。此方法只能调用 mainSource，而不能调用 proxySource。

参数：无。

返回：无。
