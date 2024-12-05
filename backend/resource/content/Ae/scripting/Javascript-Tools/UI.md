---
title: 用户UI
order: 6
---

# 用户 UI

    在线构建：[https://scriptui.joonas.me](https://scriptui.joonas.me/)

(一般都是在线构建，然后查对应控件的用法)

UI 构建原型：[UI Events](https://www.w3.org/TR/uievents/#ui-events-overview)（深入学习构建控件来源，可以看看）

本站只提供说明，以及属性、函数参考，也就是快速入门使用。

## 窗体

#### new Window

new Window (type [, title, bounds, {creation_properties}]);

参数

- type：窗口类型
  - dialog - 对话框，不关闭无法进行其他操作。
  - palette - 浮动对话框，可以贴到其他面板里. （Photoshop CC 不支持。）
  - window- 简单窗口，可用作应用程序的主窗口。（Photoshop CC 不支持。）
- title：窗口标题。
- bounds：窗口的位置和大小。

示例：创建窗体

![](https://cdn.yuelili.com/20211015142542.png)

示例 1：直接创建

```javascript
// 创建一个dialog类型窗口，左上角名称为：用户提示
var dlg = new Window("dialog", "用户提示");
// 给窗口添加一个静态文字
var myMessage = dlg.add("statictext");
// 给静态文字添加内容
myMessage.text = "Hello, world!";
// 显示窗体
dlg.show();
```

示例 2：函数包裹

```javascript
// 创建UI 函数 function
var win = createWindow();
win.show();
createWindow() {
  var w = new Window('palette',"用户窗口"); // ← ← ← 创建一个palette类型窗口，左上角名称为：用户提示
  var m = w.add('statictext');
  m.text = 'Hello,world!';
  return w;
  }
```

## 窗体与控件类型

窗口类型介绍（英文版）：[窗口对象](https://extendscript.docsforadobe.dev/user-interface-tools/window-object.html?highlight=palette)

控件类型介绍（英文版） ：[控件对象](https://extendscript.docsforadobe.dev/user-interface-tools/control-objects.html?highlight=palette)

![](https://cdn.yuelili.com/20211015145803.png)

| 窗口名称 | 类型     | 说明                                       |
| -------- | -------- | ------------------------------------------ |
| dialog   | 对话框   | 在关闭之前，没法干其他操作                 |
| palette  | 面板     | 可以与其他面板合并（需要放在 UI Panel 里） |
| window   | 软件面板 | 不太常用                                   |

| 控件名称     | 类型       | 说明                                                         |
| ------------ | ---------- | ------------------------------------------------------------ |
| statictext   | 文字类控件 | 静态文字。说明信息，或者只展示内容                           |
| edittext     | 文字类控件 | 文本编辑框。可以输入文字                                     |
| button       | 按钮类控件 | 按钮。用于执行某些工作，比如单击运行                         |
| iconbutton   | 按钮类控件 | 图片按钮。跟按钮一样，不过可以使用图片，并且多了些参数       |
| radiobutton  | 按钮类控件 | 单选按钮。使用 group 包裹，则可以多选 1。                    |
| checkbox     | 按钮类控件 | 复选框。提供选择/未选择反馈                                  |
| dropdownlist | 多元素控件 | 下拉菜单。只显示一项，单击弹出其他选择项                     |
| listbox      | 多元素控件 | 列表框。直接以列表的形式显示多个项目                         |
| slider       | 个性化控件 | 滑块。跟 AE 滑块控件差不多                                   |
| progressbar  | 个性化控件 | 进度条。                                                     |
| image        | 个性化控件 | 图片。可以显示图片                                           |
| group        | 容器       | 分组。把其他控件放一起管理                                   |
| panel        | 容器       | 分割面板区（一个小方框）。用于分割控件                       |
| tabbedpanel  | 容器       | 切换面板。可以把其他组件分类，放置于切换面板子项，单击切换。 |
| tab          | 容器       | 切换面板的单面板                                             |
| treeview     | 容器       | 树状视图。                                                   |

## 窗体与控件属性

全部读/写。很多都是通用的，自己看看就行
| 通用属性 | 说明 | 类型 | 参数 | 参数说明 | 图例 |
| ----------------------- | ------------------------------------ | ---------------------------- | -------------------------------- | ------------------------------------------ | -------------------------------------------------- |
| alignChildren | 子组件/子项目的对齐方式 | 数组[水平参数,竖直参数] | center | 居中(水平方向) | ![](https://cdn.yuelili.com/20211015152742.png) |
| | | | left | 居左(水平方向) | ![](https://cdn.yuelili.com/20211015152648.png) |
| | | | right | 居右(水平方向) | ![](https://cdn.yuelili.com/20211015152804.png) |
| | | | fill | 自动填充：填满容器(水平方向) | 根据设置的参数，水平自动填充 |
| | | | center | 居中(竖直方向) | 跟上面类似，不上图了 |
| | | | top | 居上(竖直方向) | |
| | | | bottom | 居下(竖直方向) | |
| | | | fill | 自动填充：填满容器(竖直方向) | |
| margins | 组件与窗体之间的间距 | 数值,[top,right,bottom,left] | 如果是数值 | 20 | ![](https://cdn.yuelili.com/20211015153822.png) |
| | | | 如果是数组[左，右，下，上] | [10,20,30,40] | ![](https://cdn.yuelili.com/20211015153706.png) |
| text | 窗口/控件显示名称（如果有的话） | 字符串 | | | ![](https://cdn.yuelili.com/20211015151534.png) |
| orientation | 窗体对齐方式 | 字符串 | row | 横向对齐 | ![](https://cdn.yuelili.com/20211015152041.png) |
| | | | column | 竖向对齐 | ![](https://cdn.yuelili.com/20211015151643.png) |
| spacing | 组件与组件空隙 | 数值 | 数值 | | ![](https://cdn.yuelili.com/20211015154209.png) |
| 面板属性 | 说明 | 类型 | 其他说明 | | 图示 |
| borderless | 无边框 | 布尔 | 如果开启，按钮也没了 | | ![](https://cdn.yuelili.com/20211015163732.png) |
| independent | 独立窗口 | 布尔 | 仅 windows。独立于 AE 界面的窗口 | | ![](https://cdn.yuelili.com/20211015163656.png) |
| closeButton | 关闭按钮 | 布尔 | | | ![](https://cdn.yuelili.com/20211015163816.png) |
| maximizeButton | 最大化按钮 | 布尔 | | | |
| minimizeButton | 最小化按钮 | 布尔 | | | |
| resizeable | 允许重新调整窗体大小 | 布尔 | 可以拽来拽去 | | |
| su1PanelCoordinates | 子面板坐标 | 布尔 | 仅 Photoshop。是否覆盖父面板设置 | | |
| enabled | 激活面板 | 布尔 | | | |
| 控件特有属性 | 说明 | 类型 | 参数 | 参数说明 | 图示 |
| active | 激活。比如刚打开面板，聚焦于此 | 布尔 | | | ![](https://cdn.yuelili.com/20211016105239.png) |
| alignment | 对齐方式。 | 字符串 | left、center、right、fill | 居左、居中、居右、自动填充 | ![](https://cdn.yuelili.com/20211015231838.png) |
| helpTip | 小提示 | 字符串 | | | |
| index | | | | | |
| onChange | | | | | |
| preferredSize.height | 组件高度。 | 数值 | | | |
| preferredSize.width | 组件宽度。 | 数值 | | | |
| text | 组件当前内容的文字 | 字符串 | | | |
| value | 当前值 | 数值、布尔 | | silder、progressbar、checkbox、radiobutton | |
| visible | 可见性。 | 布尔 | | | |
| 特殊控件属性 | 说明 | 类型 | 参数 | 参数说明 | |
| characters | 控件长度。以字符个数表示 | 数值 | | statictext、edittext | |
| truncate | 截断。超出的文字部分用省略号表示 | 字符串 | middle：中间截断 | statictext | ![](https://cdn.yuelili.com/20211017111344.png) |
| | | | end：结尾截断 | | |
| | | | none：不截断 | | |
| multiline | 多行。允许多行显示 | 布尔 | | statictext、edittext | ![](https://cdn.yuelili.com/20211016114835.png) |
| scrolling | 可滚动。超出显示区域会出现滚动条 | 布尔 | | statictext | ![](https://cdn.yuelili.com/20211016114625.png) |
| scrollable | 可滚动。超出显示区域会出现滚动条 | 布尔 | | edittext | ![](https://cdn.yuelili.com/20211016114625.png) |
| wantReturn | 允许回车。可以使用回车换行 | 布尔 | | | |
| borderless | 没有标题栏和边框。相关属性也会失效 | 布尔 | | edittext | |
| readonly | 只读。无法编辑 | 布尔 | | edittext | |
| noecho | 密码显示。跟输入密码一样 | 布尔 | | edittext | |
| enterKeySignalsOnChange | 失焦。文本更改且失焦，会发送一个信号 | 布尔 | 为 true 则只接受回车的 onChange | edittext | |
| justify | 文字对齐方式 | 字符串 | left、center、right | statictext、edittext、button | left：![](https://cdn.yuelili.com/20211016112822.png) |
| | | | | | center：![](https://cdn.yuelili.com/20211016112838.png) |
| | | | | | rigth：![](https://cdn.yuelili.com/20211016112851.png) |
| maxvalue | 最大值 | | | silder、progressbar | |
| minvalue | 最小值 | | | silder、progressbar | |
| items | 项目集 | | | dropdownlist、listbox | |
| multiselect | 允许多选项目 | | | listbox | |
| columnTitles | | | | listbox | |
| showHeaders | | | | listbox | |
| selection | | | | dropdownlis、listbox、vtable、tpanel | |
| columnTitles | | | | listbox | |
| showHeaders | | | | listbox | |
| selection | | | | dropdownlis、listbox、vtable、tpanel | |

### borderStyle

说明：边框样式

```javascript
w = new Window("dialog");
w.grp1 = w.add("group");
w.grp1.add("panel", [0, 0, 100, 100], "None", { borderStyle: "none" });
w.grp1.add("panel", [0, 0, 100, 100], "Gray", { borderStyle: "gray" });
w.grp1.add("panel", [0, 0, 100, 100], "Black", { borderStyle: "black" });
w.grp1.add("panel", [0, 0, 100, 100], "White", { borderStyle: "white" });
w.grp2 = w.add("group");
w.grp2.add("panel", [0, 0, 100, 100], "Etched", {
  borderStyle: "etched",
});
w.grp2.add("panel", [0, 0, 100, 100], "Sunken", {
  borderStyle: "sunken",
});
w.grp2.add("panel", [0, 0, 100, 100], "Raised", {
  borderStyle: "raised",
});
w.show();
```

![](https://cdn.yuelili.com/20211016111255.png)

### add() 添加组件或子项

window.add(ui)

说明：向窗体添加 UI 组件，或者向组件添加子项

参数：ui： UI 组件信息

第一个参数：必填。一般为组件名称，比如 statictext：静态文字

第二个参数：选填。一般为组件尺寸，[top,left,width,hight]。不填则自适应，使用 undefined 占位即可。

第三个参数：选填。一般为组件的变量名，当然也可以直接 var xxx = 组件信息

示例：

```javascript
var myWindow = new Window("dialog", "用户提示"); // 定义窗体
// 直接添加（参数单独分开）
var myText = myWindow.add("statictext", undefined, "Hello World"); // ← ← ← 添加组件方法1 // 直接添加（参数放在字典里）
var myText = myWindow.add("statictext {text: 'Hello, world!'}"); // ← ← ← 添加组件方法2 // 先添加 然后给参数赋值 var
myText = myWindow.add("statictext"); // ← ← ← 添加组件方法3
myText.text = "Hello World";
```

### show() 显示窗体

window.show

说明：显示窗体

案例：显示窗体

```javascript
var myWindow = new Window("dialog", "用户提示"); // 创建一个dialog类型窗口，左上角名称为：用户提示
var button1 = myWindow.add("button", undefined, undefined, { name: "button1" });
button1.text = "我是按钮" myWindow.show(); // ← ← ← 显示窗体。
```

### onShow() 窗体显示时

window.onShow

说明：窗体显示时

案例：窗体显示时，弹窗

```javascript
var myWindow = new Window("dialog", "窗体");
myWindow.onShow =function () {
  // ← ← ← 显示窗体时运行的函数。
  alert("窗体显示啦！")
  } myWindow.show();

```

### layout.layout()

### remove()

### UI 组件

英文原版：https://extendscript.docsforadobe.dev/user-interface-tools/control-objects.html

### Icon buttons 图片按钮

图片按钮

示例：按钮构建与切换状态

![](https://cdn.yuelili.com/20211016120247.png)

```javascript
var w = new Window("dialog");
var f = File("/d/test/icon.png");
//本地图标文件
w.add("iconbutton", undefined, f);
// a：按钮里嵌入一个图片
w.add("iconbutton", undefined, f, { style: "toolbutton" });
// b：只显示图片，不显示按钮
var t1 = w.add("iconbutton", undefined, f, { style: "toolbutton", toggle: true });
// c：只显示图片，未切换
var t2 = w.add("iconbutton", undefined, f, { style: "toolbutton", toggle: true });
// d：只显示图片，切换
t2.value = true; // ← ← ← 切换时 这里的value为true就行
w.show();
```

案例：不同状态显示不同图标

![](https://cdn.yuelili.com/20211016121210.png)

```javascript
var dir = "/d/scriptui/fig/";

var icons = {
  a: File(dir + "icon-a.png"),
  b: File(dir + "icon-b.png"),

  c: File(dir + "icon-c.png"),
  d: File(dir + "icon-d.png"),
};

var w = new Window("dialog");

b = w.add("iconbutton", undefined, ScriptUI.newImage(icons.a, icons.b, icons.c, icons.d));

w.show();
```

## radiobutton 单选按钮

示例：单击 OK，提示你选择了哪一项

![](https://cdn.yuelili.com/20211016123340.png)

```javascript
var w = new Window("dialog");

var radio_group = w.add("panel");

radio_group.alignChildren = "left";

radio_group.add("radiobutton", undefined, "AE");

radio_group.add("radiobutton", undefined, "PR");

radio_group.add("radiobutton", undefined, "PS");

radio_group.add("radiobutton", undefined, "AME");

w.add("button", undefined, "OK");

radio_group.children[0].value = true; // 初始选择第1个按钮

function selected_rbutton(rbuttons) {
  for (var i = 0; i < rbuttons.children.length; i++) {
    if (rbuttons.children[i].value == true) {
      return rbuttons.children[i].text;
    }
  }
}

//

if (w.show() == 1) {
  alert("你选择了 " + selected_rbutton(radio_group));
}
```

## listbox 列表框

[英文文档](https://extendscript.docsforadobe.dev/user-interface-tools/control-objects.html?highlight=progressbar#listbox)

### 初始化

```javascript
// 直接添加

var w = new Window("dialog");
var myList = w.add("listbox", undefined, ["one", "two", "three"]);
w.show();

// 用 add 添加子项

var w = new Window("dialog");
var myList = w.add("listbox");
myList.add("item", "one");
myList.add("item", "three");
myList.add("item", "two", 1); // 1代表插入到第2位
w.show();

// 添加参数（允许多选）

var w = new Window("dialog");
var myList = w.add("listbox", undefined, ["one", "owo", "three"], { multiselect: true });
w.show();
```

### 选择与查找

```javascript
myList.selection = [0, 1]; // 列表当前选择项
myList.selection = myList.find("two"); // 查找一个子项，并选择（注意 find只能找一个）
myList.revealItem("two"); // 滚动到子项。项目太多的时候，方便查找。参数可以是字符串/index
```

### 个性化

![](https://cdn.yuelili.com/20211016135446.png)![](https://cdn.yuelili.com/20211016135459.png)

```javascript
myList.items[i].image = File(一个图片路径); // 项目前加图片
myList.items[1].checked = true; // 项目前加一个 √
```

### 多列列表

![](https://cdn.yuelili.com/20211016135702.png)

```javascript
var w = new Window("dialog");

var myList = w.add("listbox", undefined, " ", {
  numberOfColumns: 3,
  showHeaders: true, // 列数与标题化
  columnTitles: ["English", "French", "Dutch"], // 标题内容
  columnWidths: [30, 30, 80], // 每列宽度
});

with (myList.add("item", "One")) {
  subItems[0].text = "Un";
  subItems[1].text = "Een";
}

with (myList.add("item", "Two")) {
  subItems[0].text = "Deux";
  subItems[1].text = "Twee";
}

with (myList.add("item", "Three")) {
  subItems[0].text = "Trois";
  subItems[1].text = "Drie";
}

w.show();
```

## dropdownlist

下拉列表跟列表框差不太多，增删改查看上面文档

分隔符

![](https://cdn.yuelili.com/20211016213209.png)

可以直接["one", "two", "-","three"] 在列表里加"-" 插入，也可以后面自己加

```javascript
var w = new Window("dialog");
var myDropdown = w.add("dropdownlist", undefined, ["one", "two", "three"]);
myDropdown.add("separator", undefined, 1); // ← ← ← 后面自己加分隔符
myDropdown.selection = 0;
w.show();
```

## tabbedpanel 切换面板

示例

![](https://cdn.yuelili.com/20211016220130.png)

创建 tab，tab 再添加子控件，所以 tabbedpanel 相当于一个壳，没啥好说的

```javascript
var w = new Window("dialog", "tab 面板演示", undefined, { closeButton: false });
w.alignChildren = "right";
var tpanel = w.add("tabbedpanel");
tpanel.alignChildren = ["fill", "fill"];
tpanel.preferredSize = [350, 300];

// 创建第1个切换面板
var tab1 = tpanel.add("tab", undefined, "tab1");
tab1.alignChildren = "fill";
var group1 = tab1.add("panel", undefined, "组1");
group1.alignChildren = "left";
group1.dtd_decl = group1.add("checkbox", undefined, "选项1");
group1.view_XML = group1.add("checkbox", undefined, "选项2");
group1.export_sel = group1.add("checkbox", undefined, "选项3");

// 创建第2个切换面板

var tab2 = tpanel.add("tab", undefined, "tab2");
tab2.alignChildren = "left";
var group2 = tab2.add("panel", undefined, "组2");
group2.btn = group2.add("button", undefined, "选项1");

// 主面板按钮

var buttons = w.add("group");
buttons.add("button", undefined, "Export", { name: "ok" });
buttons.add("button", undefined, "Cancel");
w.show();
```

## slider

![](https://cdn.yuelili.com/20211017002135.png)

示例：拖动滑块数字改变，改变数字，滑块也会动

```javascript
var w = new Window("dialog");
var myText = w.add('edittext {text: 50, characters: 3, justify: "center", active: true}');
var slider = w.add("slider { minvalue: 0, maxvalue: 100, value: 50 }");

slider.onChanging = function () {
  myText.text = slider.value;
};

myText.onChanging = function () {
  slider.value = Number(myText.text);
};

w.show();
```
