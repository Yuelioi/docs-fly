---
title: Item
order: 5
---

# Item 项目

## Item 对象

app.project.items[index]/app.project.item(index)

描述：Item 对象表示可以出现在“项目”面板中的项目。首项索引为 1。

子级关系：Item 是音视频项目（AVItem）和文件夹项目（FolderItem）的基类，后两者又是其他各种 item 类型的基类。因此本页 Item 属性和方法，这些 item 类型都能用。

示例：

从项目面板中获取第二个项目，并检查它是否为文件夹。然后，从文件夹中删除未选择的项目。

![](https://mir.yuelili.com/2021/07/ae5574d72fcab20f7db365198cf375c5.png)![](https://mir.yuelili.com/2021/07/377413dbc6932dff8dbfe665236d26c8.png)

```javascript
var myFolder = app.project.item(2);
if (!(myFolder instanceof FolderItem)) {
  alert("error: 第二项不是文件夹");
} else {
  var numInFolder = myFolder.numItems;
  //删除内容时始终向后运行循环:
  for (var i = numInFolder; i >= 1; i--) {
    var curItem = myFolder.item(i);
    if (curItem.parentFolder !== myFolder) {
      alert("AE: 父文件夹未正确设置");
    } else {
      if (!curItem.selected) {
        //当前项未被选择，则删除.
        curItem.remove();
      }
    }
  }
}
```

## 属性

### comment 项目注释

app.project.item(index).comment

描述：给项目添加注释，长度最多为 15999 字节。

类型：字符串; 读/写。

![](https://mir.yuelili.com/2021/07/8ebf6e9629a89e5bc2c2afbb491f6396.png)

```javascript
app.project.item(1).comment = "我是项目1";
```

### dynamicLinkGUID 动态链接标识号

app.project.item(index).dynamicLinkGUID

描述：用于动态链接的唯一且持久的标识号，形式为 00000000-0000-0000-0000-000000000000。

类型：字符串; 只读。

### guides 参考线

app.project.item(index).guides

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：参考线对象，包含参考线方向（orientationType），和位置（position）属性。

类型：字符串; 只读。

### id 项目 ID

app.project.item(index).id

描述：在内部用于标识项目的唯一且持久的标识号。重新加载，ID 不变。但是将此项目导入新项目时，会分配新的 ID。该 ID 在用户界面不显示，只能通过脚本获取。

类型：整数; 只读。

### label 标签色号

app.project.item(index).label

描述：标签颜色。颜色由其数字表示(“ 0”代表“无”，或者“ 1 至 16”代表“标签”首选项中的一种预设颜色)。

::: info 提示
自定义标签颜色无法通过编程设置。
:::

类型：整数(0 到 16)；读/写。

示例：显示项目中第 5 个项目的标签号，顺序就是设置的那种顺序。

```javascript
alert(app.project.item(5).label);
```

![](https://mir.yuelili.com/2021/07/790ae3d7d5bddf651001f134e6307a91.png)

### name 项目名称

app.project.item(index).name

描述：在“项目”面板中显示的项目名称。

类型：字符串; 读/写。

### parentFolder 父级文件夹

app.project.item(index).parentFolder

描述：此项的父级文件夹对象。如果位于最顶层，则返回根文件夹(app.project.rootFolder)。可以使用 ItemCollection.addFolder()添加新文件夹，并将项目放入新文件夹中。

类型：文件夹对象(FolderItem)；读/写。

示例：在“项目”面板创建一个新文件夹，并将所有合成移入其中。

```javascript
//创建文件夹对象，命名为comps v
ar compFolder = app.project.items.addFolder("comps");
// order：判断文件夹在合成前还是后（项目面板）
var order = true;
for (var i = 1; i <= app.project.numItems; i++) {
  alert(i);
  alert(app.project.item(i).name);
  if (app.project.item(i).name == 'comps') {
    order = false;
  }
  if (app.project.item(i) instanceof CompItem) {
    app.project.item(i).parentFolder = compFolder;
    //找到合成并添加到文件夹后，如果合成在文件夹前，当前循环数-1，文件总数-1，否则直接循环
    if(order == true){
      i--; }
      }
  }
```

### selected 选择状态

app.project.item(index).selected

描述：如果为 true，则选择该项目。为 false 则取消选择。可以同时选择多项。

类型：布尔值 读/写。

示例：选中项目面板中，第 1/3/5 个项目

```javascript
var arr = Array(1, 3, 5);
for (i in arr) {
  app.project.item(arr[i]).selected = true;
}
```

### typeName 项目类别名

app.project.item(index).typeName

描述：用户项目类型名称；例如“文件夹”，“样式”或“合成”。这些名称取决于语言环境。（英文版本：Folder；中文版：文件夹）

类型：字符串; 只读。

示例：查看项目面板第 6 个项目的类型名称

```javascript
alert(app.project.item(6).typeName);
```

![](https://mir.yuelili.com/2021/07/9f624ec01f0c05d2721c14246adf6d04.png)

## 方法

### addGuide() 创建参考线

app.project.item(index).addGuide(orientationType, position)

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：创建一个新参考线，并将其添加到 guidesItem 的对象中。

参数：

- orientationType：0，水平参考线(默认)；1，垂直参考线。
- position：参考线的 X 或 Y 坐标(以像素为单位)；整数。

返回：新创建参考线的索引；整数 。

示例：创建垂直参考线，且距离起点 500px。

```javascript
app.project.activeItem.addGuide(1, 500);
```

![](https://mir.yuelili.com/2021/07/791b911108a92dac7b779373c7fbcb77.png)

### remove() 删除项目

app.project.item(index).remove()

描述：从项目和“项目”面板中删除此项目。如果该项目是文件夹（FolderItem），则文件夹中包含的所有项目也将从项目中删除。不是从硬盘删除，而是从 AE 项目面板！

参数：无。

返回：无。

### removeGuide() 删除参考线

app.project.item(index).removeGuide(guideIndex)

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：根据索引删除现有参考线。

参数：guideIndex，参考线索引；整数。从 0 开始

返回：无。

示例：删除第一个参考线。

```javascript
app.project.activeItem.removeGuide(0);
```

:::danger：
删除索引 2，那么原来的索引 3 就会变成 2。
:::

### setGuide() 修改参考线

app.project.item(index).setGuide(position,guideIndex)

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：根据参考线索引，修改现有参考线位置。不能更改参考线方向。

参数：

- position：新的 X 或 Y 坐标位置(以像素为单位)；整数
- guideIndex：参考线索引；整数。从 0 开始

返回：无。

示例：将第 1 个参考线的位置更改为 1200 像素。

```javascript
app.project.activeItem.setGuide(1200, 0);
```
