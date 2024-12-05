---
title: FolderItem
order: 3
---

# FolderItem 文件夹

## FolderItem object

app.project.FolderItem

下面可能有点绕，只要记住，这些全都是针对于“文件夹”这一大类说的。

描述：文件夹对象（FolderItem），对应“项目”面板中的文件夹项。它可以包含各种类型的项目(素材，合成，纯色层)以及其他文件夹。

示例：假定项目中的第二个项目是文件夹（FolderItem），下面代码会遍历文件夹内所有项目，并弹出其名称。

```javascript
var secondItem = app.project.item(2);
if (!(secondItem instanceof FolderItem)) {
  alert("报错: 第2项不是文件夹，请重新选择");
} else {
  for (var i = 1; i <= secondItem.numItems; i++) {
    alert("第" + i + "项的名称为：" + secondItem.item(i).name);
  }
}
```

## 属性篇

### items 项目集

app.project.item(index).items

描述：包含 Item 的 Item 集（ItemCollection），代表此文件夹的顶级内容。这个集合仅包含文件夹中的子级项目（不包含孙级），而工程对象(Prpject)中的 ItemCollection 集则是包括所有级别（子孙）。只有在根文件夹中处于顶层的项目才在项目中也处于顶层。

类型：ItemCollection 对象；只读。

示例：查看当前文件夹子级数量 vs 项目面板所有项目数量

![](https://mir.yuelili.com/2021/07/aebece4f63f2413b362989fbe0f3a450.png)

```javascript
var myItems = app.project.item(1).items;
alert(myItems.length); //返回2，myIterm共有2个子项（不包含子项的子项） var myItems =
app.project.items;
alert(myItems.length); //返回7，项目面板一共7个（包含所有）
```

### numItems 项目个数

app.project.item(index).numItems

描述：项目集合包含的项目数。如果包含另一个文件夹，不会计算另一个文件夹的子项。

类型：整数; 只读。

案例：返回当前项目的子项数量

```javascript
alert(app.project.activeItem.numItems);
```

## 方法

### FolderItem.item()

app.project.item(index).item(index)

描述：返回此文件夹中指定索引位置的顶级项目。顶级意味着不往下层检索

注意：这里官方文档给的是错的（https://ae-scripting.docsforadobe.dev/items/folderitem/#folderitem-item）

参数：

- index：要检索的项目的位置索引，整数。起始索引为 1。

返回：Item 对象。

![](https://cdn.yuelili.com/20210912205548.png)

示例：父文件夹中第 4 个项目的名称（父文件夹在项目第 2 位），Comp 5/7/8 不计数

```javascript
alert(app.project.item(2).item(4).name); //弹窗：子文件夹2
```
