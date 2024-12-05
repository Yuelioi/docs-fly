---
title: ItemCollection
order: 6
---

# ItemCollection 项目集

## app.project.items

描述：ItemCollection 对象表示项目(item)的集合。

说明：

- Project 对象下的 ItemCollection，包含项目中所有 Item 对象。
- FolderItem 对象下的 ItemCollection，包含该文件夹中所有 Item 对象。

关系：ItemCollection 是 Collection 对象的子类。使用 ItemCollection 时，除下面列出的方法和属性外，Collection 的所有方法和属性均可用。

简写：ItemCollection → items

## 方法

### addComp() 创建合成

app.project.items.addComp(name, width, height, pixelAspect, duration,
frameRate)

描述：创建新合成。并返回一个新的合成对象（CompItem），并将其添加到此集合中。

如果 itesm 属于项目或根文件夹，则新项目的父文件夹（parentFolder）为根文件夹。

如果 itesm 属于任何其他文件夹，则新项目的父文件夹为该文件夹（ FolderItem）。

参数：

- name：名称，字符串。
- width：宽度(像素)，[4~30000]
- height：高度(像素)，[4~30000]
- pixelAspect：像素长宽比，[0.01~100.0]
- duration：持续时间(秒)，[0.0~10800.0]
- frameRate：帧速率， [1.0~99.0]

返回：CompItem，合成对象。

示例：新建一个名为合成 1，宽 1902，高 1080,100 秒，25fps 的新合成。

```javascript
app.project.items.addComp("合成1", 1920, 1080, 1, 100, 25);
```

### addFolder() 创建文件夹

app.project.items.addFolder(name)

描述：创建文件夹。并返回一个文件夹对象（FolderItem）。

如果 itesm 属于项目或根文件夹，则新项目的父文件夹（parentFolder）为根文件夹。

如果 itesm 属于任何其他文件夹，则新项目的父文件夹为该文件夹（ FolderItem）。

要将项目放在文件夹中，请使用 Item.parentFolder 进行设置

参数：

- name：文件夹名称，字符串。

返回：FolderItem，文件夹对象。

示例：“项目”面板中创建文件夹，并将合成移入其中。

```javascript
//创建文件夹对象，命名为comps
var compFolder = app.project.items.addFolder("comps");
//把所有合成已到comps文件夹中 //只需把项目的 parentFolder 设置为comps即可
var order = true for (var i = 1; i <= app.project.numItems; i++) {
  if (app.project.item(i).name == 'comps') { order = false } if
    (app.project.item(i) instanceof CompItem) {
      app.project.item(i).parentFolder = compFolder;
    //找到合成并添加到文件夹后，如果合成在文件夹前，当前循环数-1，文件总数-1，否则直接循环
    if (order == true) {
      i--;
    }
  }
}
```
