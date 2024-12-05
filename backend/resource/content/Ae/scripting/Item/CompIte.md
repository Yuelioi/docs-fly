---
order: 2
title: CompItem
---
# CompItem 合成

## CompItem object

app.project.item(index)app.project.items[index]

描述：CompItem 对象代表一个合成。通过项目和项目索引进行访问此对象。

父级关系：CompItem 是 AVItem 对象的子类，AVItem 对象是 Item 111 对象的子类。因此 AVItem 和 Item 的所有方法和属性均可在 CompItem 中使用。

示例：假定项目面板第 1 个是合成（CompItem），下面代码将弹出两个对话框。第一个显示合成里的总图层数，第二个显示合成中的最后一个图层的名称。

```javascript
var firstComp = app.project.item(1);
alert("一共有 " + firstComp.numLayers + "个图层");
alert("最后一个图层名为： " + firstComp.layer(firstComp.numLayers).name);
```

## 属性篇

### activeCamera 活动摄像机

app.project.item(index).activeCamera

描述：活动摄像机对象，当前启用的，最前面的摄像机图层。如果不包含活动摄相机，则返回 null。

类型：CameraLayer 对象；只读。

### bgColor 背景色

app.project.item(index).bgColor

描述：合成的背景色。红色、绿色、蓝色

类型：浮点数的数组，读/写。[R,G,B] ，范围均为 0.0~1.0

示例：把第 1 个项目的背景色改为中灰

```javascript
app.project.item(1).bgColor = [0.5, 0.5, 0.5];
```

### counters 不懂

app.project.item(index).counters

::: info 提示
在 After Effects 13.2(CC2014)中添加了此功能。
:::

::: danger
此方法无正式记录！可能会失效。如果有更多信息，请联系！
:::

描述：此属性适用于整个应用程序：如果在一个 CompItem 上进行了更改，它将为项目中的每个 CompItem 进行更改。该值保持不变，直到重新启动 AE。重新启动后，它将恢复为 false。

类型：布尔值 读/写。

### displayStartFrame 合成起始帧

app.project.item(index).displayStartFrame

描述：合成起始帧数。

![](https://mir.yuelili.com/2021/07/d12e0ae32bfc7be309d897e0f25cb9f4.png)

该值是使用 CompItem.displayStartTime 和 CompItem.frameDuration 计算起始帧以补偿浮点问题的替代方法。

::: info 提示
此功能在 After Effects 17.1 中添加。
:::

类型：整数; 读/写。如果不是合成，返回 undefined

示例：查看第 1 个项目当前起始帧数：返回 50

```javascript
alert(app.project.item(1).displayStartFrame);
```

### displayStartTime 合成起始秒数

app.project.item(index).displayStartTime

![](https://mir.yuelili.com/2021/07/d12e0ae32bfc7be309d897e0f25cb9f4.png)

描述：设置为合成开始时间。等效于“合成设置” - “开始时间码”或“开始帧”设置。

::: info 提示
从 After Effects 17.1 开始，最小值为-10800.0。在 17.1 之前，最小值为 0.0
:::

类型：浮点值，范围为-10800.0~86339.0，也就是(-3：00：00：00 到 23：59：00：00)；读/写。如果不是合成，可能返回 undefined

示例：查看第 1 个项目当前起始秒数：返回 2（秒）

```javascript
alert(app.project.item(1).displayStartTime);
```

### draft3d 草图 3D

app.project.item(index).draft3d

![](https://mir.yuelili.com/2021/07/ed86ca76ff107bc6c1de9c9042739a42.png)

描述：“合成”面板的“草图 3D”模式。

类型：布尔值，读/写。

### dropFrame 合成丢帧设置

app.project.item(index).dropFrame

![](https://mir.yuelili.com/2021/07/766e6f6df6a57fbf43a488505f68e440.png)

描述：丢帧/非丢帧时间码。可以在“合成设置”中进行设置。

类型：布尔值，读/写。

### frameBlending 帧混合

app.project.item(index).frameBlending

![](https://mir.yuelili.com/2021/07/0940861c89d0c1fecea12683369a990e.png)

描述：合成是否启用帧混合。

类型：布尔值；如果为 true，则启用帧混合；读/写。

### frameDuration 单帧持续时间

app.project.item(index).frameDuration

描述：单帧持续时间，（秒）。是 frameRate(每秒帧数)的倒数。

类型：浮点; 读/写。

### hideShyLayers 消隐

app.project.item(index).hideShyLayers

![](https://cdn.yuelili.com/20210912200507.png)

描述：合成的 shy 开关（消隐开关）

类型：布尔值，读/写。

### layers 图层集

app.project.item(index).layers

描述：一个图层对象集（LayerCollection），包含此合成中所有图层对象。一般配合其他参数使用，比如遍历图层。

类型：LayerCollection 对象；只读。

### markerProperty 标记属性

app.project.item(index).markerProperty

::: info 提示
在 After Effects 14.0(CC 2017)中添加了此功能
:::

描述：包含所有合成标记的对象集（PropertyGroup）。合成标记与图层标记具有相同的功能。参见 MarkerValue 对象

类型：属性对象组（PropertyGroup），或 null。只读。

示例：创建一个新的项目、新合成、新图层，然后创建两个合成标记

![](https://mir.yuelili.com/2021/07/9f845d259dd3dbac2f07489ece284641.png)

```javascript
// comp.markerProperty 允许添加标记到合成，与layer.property("Marker")类似
var currentProj = app.newProject();
var comp = currentProj.items.addComp("要标记的合成", 1920, 1080, 1.0, 5, 29.97);
var solidLayer = comp.layers.addSolid([1, 1, 1], "要标记的图层", 1920, 1080, 1.0);
var compMarker = new MarkerValue("我是合成标记!");
compMarker.duration = 1;
var compMarker2 = new MarkerValue("另一个合成标记");
compMarker2.duration = 1;
comp.markerProperty.setValueAtTime(1, compMarker);
comp.markerProperty.setValueAtTime(3, compMarker2);
```

### motionBlur 运动模糊

app.project.item(index).motionBlur

![](https://mir.yuelili.com/2021/07/01e76b852dabb7dbee37d8fa5edf5a33.png)

描述：对应于“合成面板” - “运动模糊”。

类型：布尔值，读/写。

### motionBlurAdaptiveSampleLimit 运动模糊采样限制

app.project.item(index).motionBlurAdaptiveSampleLimit

![](https://mir.yuelili.com/2021/07/37deeb5e048f86fcb702e1d8cbc84a06.png)

描述：2D 图层运动的运动模糊采样的最大数量。对应于“合成设置” - “高级” - “自适应采样限制”。

类型：整数(16 到 256 之间)；读/写。

### motionBlurSamplesPerFrame

app.project.item(index).motionBlurSamplesPerFrame

![](https://mir.yuelili.com/2021/07/ba2fbe4325d81eeb30e28abc8aac8dcc.png)

描述：经典 3D 图层，形状图层和某些效果的每帧最小运动模糊采样数。对应于“合成设置”-“高级”-“每帧样本数”。

类型：整数(2 到 64 之间)；读/写。

### motionGraphicsTemplateControllerCount 基本图形属性个数

app.project.item(index).motionGraphicsTemplateControllerCount

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：合成的“基本图形”面板中的属性个数。

类型：整数; 只读。

示例：当前为 2 个属性

![](https://mir.yuelili.com/2021/07/c41b00b232b5062c65df8527dbcbd0a2.png)

app.project.item(1).motionGraphicsTemplateControllerCount // 返回 2

### motionGraphicsTemplateName 基本图形模板名称

app.project.item(index).motionGraphicsTemplateName

::: info 提示
在 After Effects 15.0(CC 2018)中添加了此功能
:::

描述：在“基本图形”面板中读取或写入该合成的模板名。

类型：字符串; 读/写。

示例：将设置活动合成的模板名称，然后弹出对话框显示模板名

```javascript
app.project.activeItem.motionGraphicsTemplateName = "我是模板名";
alert(app.project.activeItem.motionGraphicsTemplateName);
```

![](https://mir.yuelili.com/2021/07/e9e90c533a27f087dddf1c5e570f7412.png)

### numLayers 图层数量

app.project.item(index).numLayers

描述：合成中的图层数。

类型：整数; 只读。

### preserveNestedFrameRate 预合成时保留帧速率

app.project.item(index).preserveNestedFrameRate

![](https://mir.yuelili.com/2021/07/7c8925fa63970651ea1c114c990e9485.png)

描述：对应于“合成设置” - “高级” - “在预合成时或在渲染队列中时保留帧速率。

类型：布尔值，读/写。

### preserveNestedResolution 预合成时保留分辨率

app.project.item(index).preserveNestedResolution

![](https://mir.yuelili.com/2021/07/e3799982ac02dd66c58647b9b8bf2667.png)

描述：对应于“合成设置” - “高级” - “预合成时保留分辨率”。

类型：布尔值，读/写。

### renderer 当前 3D 渲染器

app.project.item(index).renderer

![](https://cdn.yuelili.com/20210912201532.png)

描述：当前的渲染模块，等同于“合成设置” - “3D 渲染器”所示。允许的值为 CompItem.renderers 的成员。

类型：字符串; 读/写。

### renderers 可用 3D 渲染器

app.project.item(index).renderers

描述：可用的渲染模块。等同于“合成设置” - “3D 渲染器”所示。

类型：字符串数组；只读。

我这里返回：[ADBE Advanced 3d, ADBE Standard 3d， ADBE Ernst]

### resolutionFactor 预览分辨率

app.project.item(index).resolutionFactor

描述：采样时跳过水平和竖直的像素数[x,y]；第一个水平，第二个垂直。

full 分辨率为[1, 1]，1/2 分辨率为[2, 2]，1/4 分辨率为[4, 4]。默认值为[1, 1]

类型：在范围内的两个整数的数组[1~99]; 读/写。

示例：将合成预览改为 1/4 分辨率

```javascript
app.project.activeItem.resolutionFactor = [4, 4];
```

### selectedLayers 选定图层集

app.project.item(index).selectedLayers

描述：此合成中的所有选定图层集合。起始索引为 0。

类型：图层对象数组；只读。

示例：先获取所有选定图层，然后再把所有图层的标签改成红色

```javascript
// 定义选择的图层组以及选择的图层总数
var selLayers = app.project.activeItem.selectedLayers;
var selNums = app.project.activeItem.selectedLayers.length;

// 遍历，把所有选择图层的标签改色
for (var i = 0; i < selNums; i++) {
  selLayers[i].label = 6;
}
```

### selectedProperties 选定属性集

app.project.item(index).selectedProperties

描述：所有选定属性组(Property 和 PropertyGroup 对象)。起始索引为 0。

类型：Property 和 PropertyGroup 对象的数组；只读。

案例：弹窗显示所有选择的属性，当前弹窗顺序为：Rotation、Positon、Scale

```javascript
// 定义选择的图层属性组以及选择的属性总数
var selLayers = app.project.activeItem.selectedProperties;
var selNums = app.project.activeItem.selectedProperties.length;

// 遍历，弹窗所有属性的名称
for (var i = 0; i < selNums; i++) {
  alert(selLayers[i].name);
}
```

![](https://mir.yuelili.com/2021/07/2e48d400d540721e45e180da1d22e4a9.png)

### shutterAngle 快门角度

app.project.item(index).shutterAngle

![](https://cdn.yuelili.com/20210912202003.png)

描述：合成的快门角度设置。对应“合成设置-高级-快门角度”的设置。

类型：整数，范围为 0~720; 读/写。

### shutterPhase 快门相位

app.project.item(index).shutterPhase

![](https://cdn.yuelili.com/20210912202003.png)

描述：合成的快门相位设置。对应“合成设置-高级-快门相位”设置。

类型：整数，范围为–360~360; 读/写。

### workAreaDuration 工作区的持续时间

app.project.item(index).workAreaDuration

![](https://mir.yuelili.com/2021/07/d3b6ebc437001e34eeeaae85524bcffe.png)

描述：工作区的持续时间(秒)。也就是工作区的起点时间和终点时间之差。

类型：浮点数; 读/写。

### workAreaStart 工作区起始时间

app.project.item(index).workAreaStart

![](https://mir.yuelili.com/2021/07/d3b6ebc437001e34eeeaae85524bcffe.png)

描述：合成工作区开始的时间(秒)。

类型：浮点; 读/写。

## 方法

### duplicate() 复制

app.project.item(index).duplicate()

描述：复制。等同于编辑 - 重复（Ctrl D）

参数：无。

返回：CompItem 对象。

### exportAsMotionGraphicsTemplate() 导出动态图形模板

app.project.item(index).exportAsMotionGraphicsTemplate(doOverWriteFileIfExisting,
file_path)

::: info 提示
在 After Effects 15.0(CC 2018)中添加了此功能
:::

描述：将合成导出为动态图形模板。如果动态图形模板已成功导出，则返回 true，否则返回 false。

“基本图形”面板中的名称就是动态图形模板的文件名(例如，“我的模板.mogrt”)。使用 motionGraphicsTemplateName 为其设置名称。

指定保存运动图形模板文件的文件夹的路径(可选)。如果未指定，则文件将保存在当前用户的“基本图形”文件夹中：

- macOS: /Users//Library/Application Support/Adobe/Common/Essential Graphics/
- Windows: C:\Users\你的用户名\AppData\Roaming\Adobe\Common\Essential Graphics\

如果项目自上次保存以来已更改，After Effects 将提示用户保存项目。为避免这种情况，请在导出运动图形模板之前使用 project 的 save()方法。

参数：

- doOverWriteFileIfExisting ：布尔值，是否覆盖同名的现有文件。必填。
- file_path：要保存的文件夹路径。选填。

返回：布尔值。

### getMotionGraphicsTemplateControllerName() 获取基本图形属性名

app.project.item(index).getMotionGraphicsTemplateControllerName(index)

![](https://mir.yuelili.com/2021/07/c41b00b232b5062c65df8527dbcbd0a2.png)

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：获取“基本图形”面板中单个属性的名称。

参数：

- index：EGP 属性的索引，整数;

返回：字符串; 只读。

弹窗第二个属性的名称：scale

```javascript
alert(app.project.activeItem.getMotionGraphicsTemplateControllerName(2));
```

### setMotionGraphicsControllerName() 设置基本图形属性名

app.project.item(index).setMotionGraphicsControllerName(index,newName)

::: info 提示
此功能已添加到 After Effects 16.1(CC 2019)中
:::

描述：在“基本图形”面板中设置单个属性的名称。

::: info 提示
要在添加到 EGP 中的属性重命名，请参见 Property.addToMotionGraphicsTemplateAs()。
:::

参数：

- index：整数; 要重命名的 EGP 属性的索引。
- newName：字符串; EGP 属性的新名称。

返回：字符串; 只读。

![](https://mir.yuelili.com/2021/07/34976752ee4f23fe2e686c5f4324722b.png)

```javascript
app.project.activeItem.setMotionGraphicsControllerName(2, "sccccale");
```

### layer() 图层

app.project.item(index).layer(index)
app.project.item(index).layer(otherLayer, relIndex)
app.project.item(index).layer(name)

描述：返回一个 Layer 对象，可以通过名称、索引或相对于另一层的索引位置来确定。

参数：

- index ：图层索引号。整数，起始值为 1
- otherLayer：该合成中的其他图层对象。
- relIndex：目标图层相对于的当前图层的索引差值，整数。比如当前图层索引为 5，目标图层索引为 3，那么 relIndex 就是-2
- name：目标图层的名称，字符串。

返回：图层对象。

示例：花式找到图层 1 的名称，先找图层 3，然后索引-2

```javascript
var srcLayer = app.project.activeItem.layer(3);
alert(app.project.activeItem.layer(srcLayer, -2).name);
```

### openInEssentialGraphics() 在基本图形面板中打开合成

app.project.item(index).openInEssentialGraphics()

::: info 提示
此功能已添加到 After Effects 15.0(CC 2018)中
:::

描述：在“基本图形”面板中打开合成。

参数：无。

返回：无。

示例：在基本图形面板打开活动合成

```javascript
app.project.activeItem.openInEssentialGraphics();
```

### openInViewer() 在查看器中打开合成

app.project.item(index).openInViewer()

描述：在“合成查看器”面板中打开目标合成，然后将“合成查看器”聚焦。

参数：无。

返回：“合成”面板的查看器对象（Viewer）；如果无法打开合成，则为 null。

案例：把工程面板第一个合成放在查看器上

```javascript
app.project.item(1).openInViewer();
```

### saveFrameToPng() 单帧保存为图片

app.project.item.saveFrameToPng(time, fileObject)

描述：合成保存单帧到 png

参数：

- time：合成的时间（秒）
- fileObject：保存的文件对象

返回：不知道啥对象

案例：保存当前合成第 1 秒内容到桌面

```javascript
app.project.activeItem.saveFrameToPng(1, File("~/Desktop/1.png"));
```

### ramPreviewTest() RAM 预览

app.project.item.ramPreviewTest(gridshow,multi,exposure)

参数

- gridshow：布尔值：查看器的透明网格。true->显示透明网格，false->显示背景色。
- multi：查看器放大倍数。1 是 100%。0.5 是 50%，依次类推。
- exposure：曝光值。

案例：当前合成打开透明网格，并以 66%倍率预览

![](https://cdn.yuelili.com/20211014001317.png)

```javascript
var comp = app.project.activeItem;
alert(comp.ramPreviewTest(true, 0.66, 0));
```
