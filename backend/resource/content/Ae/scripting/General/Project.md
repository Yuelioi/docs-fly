---
title: Project
order: 3
---

# Project 工程

## 说明

描述：app.project 代表一个 After Effects 工程对象。以下是它的属性与方法

属性：通过属性可以访问工程中的特定对象(例如导入的文件、素材和合成)，还可以访问项目设置(例如时间码)。

方法：使用方法可以导入素材，创建纯色层，合成和文件夹，并保存更改。

## 属性篇

### activeItem 活动项

app.project.activeItem

描述：当前处于活动状态的项目；如果未选择或选择了多个项目，则为 null。一般配合属性或者方法使用

类型：Item 对象或 null。只读。

### bitsPerChannel 颜色深度

app.project.bitsPerChannel

描述：当前项目的颜色深度，可以是 8 位，16 位或 32 位。

类型：整数；读/写。

```javascript
alert(app.project.bitsPerChannel); // 弹窗：8
```

### compensateForSceneReferredProfiles 补偿场景配置文件

app.project.compensateForSceneReferredProfiles

::: info 提示
在 After Effects 16.0(CC 2019)中添加了此功能
:::

![](https://cdn.yuelili.com/20210912023150.png)

描述：启用 项目设置 - 颜色-补偿场景引用的配置文件，则为 true；否则为 false。

类型：布尔值；读/写。

### dirty 工程保存状态

app.project.dirty

::: info 提示
在 After Effects 17.5(CC2020)中添加了此功能。
:::

::: danger
此方法/属性无正式记录，是通过研究发现的。这里的信息可能不准确，整个方法/属性可能会消失或在某些程度上停止工作。
:::

![](https://cdn.yuelili.com/20210912023244.png)

描述：如果当前工程未保存（后面带\*），则为 true；否则为 false。

类型：布尔值 只读。

### displayStartFrame 起始帧计数

app.project.displayStartFrame  
![](https://mir.yuelili.com/2021/07/326c026389e1a179c7a5910efb32b471.png)
![](https://mir.yuelili.com/2021/07/4df6c5ca7377aa98be38224e4473d798.png)

描述：在“项目设置”中“帧计数”设置为 0 或 1 的另一种方法，等效于使用 frameCountType 的 FramesCountType.FC_START_0 或 FramesCountType.FC_START_1。

类型：整数(0 或 1)；读/写。

::: info 提示
在 CS5 以及更高版本，“时间显示格式”已发生很大变化，因此老版本可能无法使用。更多信息，请参见 After Effects CS5.5 。
:::

```javascript
app.project.displayStartFrame = 1; //把起始帧设置为1
```

### expressionEngine 表达式引擎

app.project.expressionEngine

::: info 提示
在 After Effects 16.0(CC 2019)中添加了此功能
:::

描述：“项目设置” - “表达式引擎”设置。

设置项

- extendscript
- javascript-1.0

类型：字符串; 读/写。

案例：测试当前表达式引擎是否为 js，弹窗：true

```javascript
alert(app.project.expressionEngine == "javascript-1.0");
```

### feetFramesFilmType

app.project.feetFramesFilmType  
![](https://mir.yuelili.com/2021/07/0a9ddf7bf9ad07b00b94396f2bb1cb53.png)
描述：“项目设置”中的“使用 feet+frame”菜单设置。使用此属性代替旧 timecodeFilmType 属性。

类型：FeetFramesFilmType 枚举值; 读/写：

- FeetFramesFilmType.MM16
- FeetFramesFilmType.MM35

示例：查看当前设置是不是 16 帧

```javascript
alert(app.project.feetFramesFilmType == FeetFramesFilmType.MM16);
```

### file 项目文件

app.project.file

描述：包含当前打开的项目的文件的 ExtendScript File 对象。

类型：文件对象；如果尚未保存工程，则为 null。只读。

示例：直接查看当前已保存文件的路径

```javascript
alert(app.project.file);
```

### footageTimecodeDisplayStartType 起始时间码

app.project.footageTimecodeDisplayStartType  
![](https://mir.yuelili.com/2021/07/095ba3ae0f23a14ce0c895835c36f222.png)
描述：“项目设置”中的“开始时间显示”设置，需先选择“时间码”作为显示样式。

类型：FootageTimecodeDisplayStartType，枚举值; 读/写：

- FootageTimecodeDisplayStartType.FTCS_START_0
- FootageTimecodeDisplayStartType.FTCS_USE_SOURCE_MEDIA

示例：使用的默认 000，因此弹窗为 true，

```javascript
alert(app.project.footageTimecodeDisplayStartType == FootageTimecodeDisplayStartType.FTCS_START_0);
```

### framesCountType 帧计数类型

app.project.framesCountType

![](https://mir.yuelili.com/2021/07/326c026389e1a179c7a5910efb32b471.png)
![](https://mir.yuelili.com/2021/07/4df6c5ca7377aa98be38224e4473d798.png)

描述：项目设置-时间显示样式-帧数。

类型：FramesCountType，枚举值; 读/写：

- FramesCountType.FC_START_1
- FramesCountType.FC_START_0
- FramesCountType.FC_TIMECODE_CONVERSION

### framesUseFeetFrames

app.project.framesUseFeetFrames

![](https://mir.yuelili.com/2021/07/78084e2bbcd7fac3e750d93bb669511c.png)

描述：“项目设置”中，如果使用“feet+frame”，则为 true；否则为 false

类型：布尔值 读/写。

### gpuAccelType GUP 加速类型

app.project.gpuAccelType

::: info 提示
此功能已添加到 After Effects 13.8(CC 2015.3)中
:::

描述：获取或设置当前项目的 GPU 加速选项。请参阅 app.availableGPUAccelTypes

类型：GpuAccelType 枚举值; 读/写：

- GpuAccelType.CUDA
- GpuAccelType.Metal
- GpuAccelType.OPENCL
- GpuAccelType.SOFTWARE

示例

```javascript
// 项目设置 -> 视频渲染与效果 -> 使用 var currentGPUSettings =app.project.gpuAccelType; // returns the current value var type_str = "";
//1.检查当前选项，并弹窗
switch (currentGPUSettings) {
  case GpuAccelType.CUDA: type_str ="CUDA":
  break;
  case GpuAccelType.METAL: type_str = "Metal":
  break;
  case GpuAccelType.OPENCL: type_str = "OpenCL":
  break;
  case GpuAccelType.SOFTWARE:  type_str = "Software";
  break;
  default: type_str = "UNKNOWN";
}
alert("Your current setting is " + type_str);
// 2.设置加速选项
app.project.gpuAccelType =GpuAccelType.METAL;
```

### items 项目

app.project.items

描述：项目中的所有项目。一般配合其他参数使用

类型：ItemCollection 对象; 只读。

### linearBlending 线性混合

app.project.linearBlending

描述：如果项目使用线性混合，则为 true；否则为 false。

![](https://cdn.yuelili.com/20210912023401.png)

类型：布尔值 读/写。

### linearizeWorkingSpace 线性工作空间

app.project.linearizeWorkingSpace

::: info 提示
在 After Effects 16.0(CC 2019)中添加了此功能
:::

![](https://cdn.yuelili.com/20210912023401.png)

描述：如果项目启用“线性化工作空间”，则为 true；否则为 false。

类型：布尔值 读/写。

### numItems 项目总数

app.project.numItems

描述：项目中包含的项目总数，包括文件夹和所有类型的素材。

类型：整数; 只读。

示例：显示当前有多少个项目

```javascript
var numItems = app.project.numItems;
alert("有 " + numItems + " 个项目.");
```

### renderQueue 渲染队列

app.project.renderQueue

描述：项目的渲染队列。详见渲染队列 RenderQueue

类型：RenderQueue 对象; 只读。

### revision 修订次数

app.project.revision

描述：项目的当前修订次数。用户的每个操作都会增加修订数。初始为 1。

其他：新建一个纯色层大概能加 4，仅供娱乐。

返回： 项目的当前修订版本；整数；只读。

### rootFolder 根文件夹

app.project.rootFolder

描述：包含项目内容的根文件夹对象；这是一个虚拟文件夹，其中包含“项目”面板中的所有项目，但不包含“项目”面板中其他文件夹中包含的项目。

类型：FolderItem 对象; 只读。

### selection 项目面板所有选择项

app.project.selection

描述：在“项目”面板中选择的所有项目，以“项目”面板中显示的排序顺序。

类型：Item 对象数组; 只读。

### timeDisplayType 时间显示样式

app.project.timeDisplayType

![](https://mir.yuelili.com/2021/07/a203b92ad633e077f454a77e28b7a5c5.png)

描述：时间显示样式，与“项目设置”对话框中的“时间显示样式”部分相对应。

类型：TimeDisplayType，枚举值; 读/写。

- TimeDisplayType.FRAMES
- TimeDisplayType.TIMECODE

### toolType 工具菜单

app.project.toolType

::: info 提示
在 After Effects 14.0(CC 2017)中添加了此功能
:::

描述：在“工具”面板中获取并设置活动工具。

类型：ToolType 枚举值; 读/写：

- ToolType.Tool_Arrow：选择工具
- ToolType.Tool_Rotate：旋转工具
- ToolType.Tool_CameraMaya：统一摄像头工具
- ToolType.Tool_CameraOrbit：轨道摄像头工具
- ToolType.Tool_CameraTrackXY：跟踪 XY 相机工具
- ToolType.Tool_CameraTrackZ：Track Z 摄影机工具
- ToolType.Tool_Paintbrush： 画笔工具
- ToolType.Tool_CloneStamp：仿制图章工具
- ToolType.Tool_Eraser：橡皮擦工具
- ToolType.Tool_Hand： 手形工具
- ToolType.Tool_Magnify：缩放工具
- ToolType.Tool_PanBehind：向后平移(锚点)工具
- ToolType.Tool_Rect：矩形工具
- ToolType.Tool_RoundedRect：圆角矩形工具
- ToolType.Tool_Oval：椭圆工具
- ToolType.Tool_Polygon：多边形工具
- ToolType.Tool_Star：星形工具
- ToolType.Tool_TextH：横排文字工具
- ToolType.Tool_TextV：竖排文字工具
- ToolType.Tool_Pen：钢笔工具
- ToolType.Tool_Feather：蒙版羽化工具
- ToolType.Tool_PenPlus：添加顶点工具
- ToolType.Tool_PenMinus：删除顶点工具
- ToolType.Tool_PenConvert：转换顶点工具
- ToolType.Tool_Pin：木偶位置控点工具
- ToolType.Tool_PinStarch：木偶固化控点工具
- ToolType.Tool_PinDepth：人偶重叠控点工具
- ToolType.Tool_Quickselect：ROTO 笔刷工具
- ToolType.Tool_Hairbrush：调整边缘工具

示例：检查当前工具，如果不是统一摄像机工具，并设置为：

```javascript
// 检查当前工具，并设置为同一摄像机工具 Unified Camera Tool (UCT).
// 假设已经选则了合成
var comp = app.project.activeItem;
if (comp instanceof CompItem) {
  // Add a camera to the current comp. (Requirement for UCT)
  var cameraLayer = comp.layers.addCamera("Test Camera", [comp.width / 2, comp.height / 2]);
  comp.openInViewer();

  // If the currently selected tool is not one of the camera tools, set it to UCT.
  if (
    app.project.toolType !== ToolType.Tool_CameraMaya &&
    app.project.toolType !== ToolType.Tool_CameraOrbit &&
    app.project.toolType !== ToolType.Tool_CameraTrackXY &&
    app.project.toolType !== ToolType.Tool_CameraTrackZ
  ) {
    app.project.toolType = ToolType.Tool_CameraMaya;
  }
}
```

示例（没试）：使用新的 app.project.toolType 属性从“项目”面板中选择的素材片段或合成创建 360 度合成(环境层和相机)。该脚本是从等角镜头中构建 VR 合成的良好起点：

```javascript
// Create a 360 VR comp from a footage item or comp selected in the Project panel.

var item = app.project.activeItem;
if (item !== null && (item.typeName === "Footage" || item.typeName === "Composition")) {
  // Create a comp with the footage.
  var comp = app.project.items.addComp(item.name, item.width, item.height, item.pixelAspect, item.duration, item.frameRate);
  var layers = comp.layers;
  var footageLayer = layers.add(item);

  // Apply the CC Environment effect and create a camera.
  var effect = footageLayer.Effects.addProperty("CC Environment");
  var camera = layers.addCamera("360 Camera", [item.width / 2, item.height / 2]);
  comp.openInViewer();
  app.project.toolType = ToolType.Tool_CameraMaya;
} else {
  alert("Select a single footage item or composition in the Project panel.");
}
```

### transparencyGridThumbnails 缩略图透明网格

app.project.transparencyGridThumbnails

描述：如果为 true，项目面板缩略图将使用透明网格。

类型：布尔值，读/写。

![](https://mir.yuelili.com/2021/07/32bce6fcb2a9291575439bfb53b8d334.png)![](https://mir.yuelili.com/2021/07/4007c57727b6fb4c112a42763c790d53.png)

### workingGamma 工作区灰度系数

app.project.workingGamma

![](https://mir.yuelili.com/2021/07/c4dc2d485eaf00ae2648af1ececa40c9.png)

描述：当前项目的工作区灰度系数值为 2.2 或 2.4。设置 2.2 或 2.4 以外的值将导致脚本错误。请注意，设置项目的颜色工作空间后，After
Effects 会忽略工作灰度系数值。

类型：数字; 读/写。

示例：

要将工作灰度系数设置为 2.4(建议 709)，请执行以下操作：

```javascript
app.project.workingGamma = 2.4;
```

```javascript
要获取当前的工作区灰度系数，请执行以下操作： var currentGamma = app.project.workingGamma;
```

### workingSpace 工作区

app.project.workingSpace

描述：一个字符串，它是项目颜色工作空间的颜色配置文件描述。要将工作空间设置为“无”，请设置 workingSpace 为空字符串。

使用 app.project.listColorProfiles()返回可用于设置颜色的工作空间可用颜色配置文件描述的数组。

类型：字符串; 读/写。

示例

```javascript
//将工作空间设置为Rec.709 Gamma 2.4：
app.project.workingSpace = "Rec.709 Gamma 2.4";
//将工作空间设置为“无”（清空设置）：
app.project.workingSpace = "";
//获得当前的工作空间：
var currentSpace = app.project.workingSpace;
```

### xmpPacket XMP 元数据

app.project.xmpPacket

描述：项目的 XMP 元数据，存储为 RDF(基于 XML)。有关 XMP 的更多信息，请参见《JavaScript 工具指南》。

类型：字符串; 读/写。

示例

以下示例代码访问当前项目的 XMP 元数据，并修改 Label 项目元数据字段。

```javascript
var proj = app.project;

// load the XMPlibrary as an ExtendScript ExternalObject
if (ExternalObject.AdobeXMPScript === undefined) {
  ExternalObject.AdobeXMPScript = new ExternalObject("lib:AdobeXMPScript");
}
var mdata = new XMPMeta(app.project.xmpPacket); //get the project's XMPmetadata
// update the Label project metadata's value
var schemaNS = XMPMeta.getNamespaceURI("xmp");
var propName = "xmp:Label";
try {
  mdata.setProperty(schemaNS, propName, "finalversion...no, really!");
} catch (e) {
  alert(e);
}

app.project.xmpPacket = mdata.serialize();
```

## 方法篇

### autoFixExpressions() 自动修复表达式

app.project.autoFixExpressions(oldText, newText)

描述：如果修好后的内容能修复表达式，则自动替换错误表达式中的文本。

参数：

- oldText：要替换的文字。
- newText： 新文本。

返回：无。

示例：修复滑块名称错误。把错误的“滑块控制 1”改为“滑块控制”

![](https://cdn.yuelili.com/20210912025228.png)

![](https://cdn.yuelili.com/20210912025351.png)

```javascript
app.project.autoFixExpressions("滑块控制1", "滑块控制");
```

### close() 关闭项目选项

app.project.close(closeOptions)

描述：更改关闭项目的选项。

参数：closeOptions，关闭时要执行的操作。枚举值：

- CloseOptions.DO_NOT_SAVE_CHANGES：关闭而不保存。
- CloseOptions.PROMPT_TO_SAVE_CHANGES：关闭前提示是否保存更改。
- CloseOptions.SAVE_CHANGES：关闭时自动保存。

返回：布尔值。如果以前未保存文件，则为 False，将提示用户，并且用户取消保存。

![](https://cdn.yuelili.com/20210912025639.png)

示例：关闭项目，此时会提示是否保存

```javascript
app.project.close(CloseOptions.PROMPT_TO_SAVE_CHANGES);
```

### consolidateFootage() 合并项目素材

app.project.consolidateFootage()

描述：合并项目中的所有素材。等同于“文件”>“合并所有素材”。

参数：无。

返回：整数; 删除的素材项目总数。

### importFile() 导入文件

app.project.importFile(importOptions)

描述：使用指定的选项，导入文件。等同于“文件”>“导入文件”。

参数：

- importOptions ：导入对象，指定要导入的文件以及操作选项。

返回：footageItem 对象。

示例：导入同级目录的视频文件（1.mp4）

```javascript
app.project.importFile(new ImportOptions(new File("1.mp4")));
```

### setDefaultImportFolder() 默认导入文件夹

app.project.setDefaultImportFolder(folder)

描述：文件 - 导入 对话框中的显示文件夹。设置以后，除非找不到该文件夹或退出 After Effects 为止。

参数：folder，ExtendScript 文件夹对象。

返回：布尔值 指示操作是否成功。

![](https://cdn.yuelili.com/20210912030209.png)

示例

把导入默认文件夹设置为 “C:/我的文件夹”的 3 种方法：

```javascript
var myFolder = new Folder("C:/我的文件夹");
app.project.setDefaultImportFolder(myFolder);
app.project.setDefaultImportFolder(new Folder("C:/我的文件夹"));
app.project.setDefaultImportFolder(Folder("C:/我的文件夹"));
```

::: danger
如果路径引写的是某个文件，则 Folder 函数将返回 File 对象而不是 Folder 对象，这将导致 setDefaultImportFolder()返回 false。
:::

默认导入文件夹为桌面： app.project.setDefaultImportFolder(Folder.desktop);

禁用默认文件夹，参数为空即可：app.project.setDefaultImportFolder();

### importFileWithDialog() 导入文件

app.project.importFileWithDialog()

描述：显示“导入文件”对话框。与“文件”>“导入”>“文件”命令相同。

返回：Item 对象的数组；如果用户取消对话框，则返回 null。

### importPlaceholder() 导入占位符

app.project.importPlaceholder(name, width, height, frameRate, duration)

![](https://mir.yuelili.com/2021/07/f720a64f2aec639c71cc63581370ae56.png)

描述：创建并返回一个新的占位符（PlaceholderItem）并将其添加到项目的 items 数组中。与 文件>导入>占位符一样。

参数：

- name ：名称，字符串。
- width ：宽度(像素)，范围为 [4~30000]
- height ：高度(像素)，[4~30000]
- frameRate ：帧速率，[1.0~99.0]
- duration ：持续时间(秒)，[0.0~10800.0]。

返回：占位符对象。

### item() 项目

app.project.item(index)

描述：在指定的索引位置检索项目。

参数：

- index ：项目的索引位置，整数。首项索引为 1。

返回：项目对象。

### itemByID() 基于 ID 检索项目

app.project.itemByID(id)

::: info 提示
此功能已添加到 After Effects 13.0(CC 2014)中
:::

描述：通过项目 ID 检索项目

其他：使用 app.project.item(index).id 可以获得 Item 的 ID

参数：

- id：整数。

返回：项目对象。一般配合其他参数使用

### reduceProject() 减少项目

app.project.reduceProject(array_of_items)

描述：从项目中删除除指定项目外的所有项目。与“文件”>“减少项目”命令相同（新版本是 文件-整理工程-减少文件）。

参数：

- array_of_items ：包含要保留的 Item 对象的数组。

返回：整数; 移除的项目总数。

示例：只保留第 1 项和第 3 项

![](https://mir.yuelili.com/2021/07/9ffb4041b5fa459b317c1e7c98a0a751.png)![](https://mir.yuelili.com/2021/07/0770cf7486514cf198e9b3d646875a56.png)

```javascript
var items = [];
items[items.length] = app.project.item(1);
items[items.length] = app.project.item(3);
app.project.reduceProject(items);
```

### removeUnusedFootage() 删除未使用素材

app.project.removeUnusedFootage()

描述：从项目中删除未使用的素材。与“文件”>“删除未使用的素材”命令相同。

参数：无。

返回：被移除的素材对象的数量，整数 。

### save() 保存工程

app.project.save([file])

描述：保存工程。与“文件”>“保存”或“文件”>“另存为”命令相同。如果从未保存过该项目，则提示用户输入文件名。使用 File 参数，则将项目保存到新文件，而无需提示。

参数：

- file：选填。一个 ExtendScript File 对象，用于保存文件。

返回：无。

示例：保存当前文件到指定路径

```javascript
app.project.save(new File("F:/1.aep"));
```

### saveWithDialog() 保存工程对话框

app.project.saveWithDialog()

描述：弹出“保存”对话框。用户可以保存项目，或单击“取消”退出对话框。

参数：无。

返回：布尔值，如果项目保存成功，则返回 true。

### showWindow() 显示项目面板

app.project.showWindow(doShow)

描述：显示或隐藏“项目”面板。

参数：

- doShow ：为 true，显示“项目”面板。为 false，则隐藏“项目”面板。

返回：无。

### listColorProfiles() 颜色配置文件数组

app.project.listColorProfiles()

描述：返回颜色配置文件描述的数组，可以将其设置为项目的颜色工作空间。

参数：无。

返回：字符串数组。

## 团队项目篇（没用过，不校正）

### newTeamProject() 新团队项目

app.project.newTeamProject(teamProjectName, description)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：创建一个新的团队项目。

参数：

- teamProjectName ：名称，字符串。
- description：选填。描述，字符串。

返回：布尔值。如果成功创建了团队项目，返回 True，否则为 false。

### openTeamProject() 打开团队项目

app.project.openTeamProject(teamProjectName)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：打开一个团队项目。

参数：

- teamProjectName ：项目名称，字符串。

返回：布尔值。成功打开则返回 True，否则为 false。

### shareTeamProject() 共享团队项目

app.project.shareTeamProject(comment)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：共享当前开放的团队项目。

参数：

- comment：注释，选填，字符串。

返回：布尔值。成功共享团队项目则为 True，否则为 false。

### syncTeamProject() 同步团队项目

app.project.syncTeamProject()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：同步当前打开的团队项目。

返回：布尔值。如果成功同步返回 True，否则为 false。

### closeTeamProject() 关闭团队项目

app.project.closeTeamProject()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：关闭当前打开的团队项目。

返回：布尔值。如果团队项目成功关闭返回 True，否则为 false。

### convertTeamProjectToProject()

app.project.convertTeamProjectToProject(project_file)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：将团队项目转换为本地磁盘上的 After Effects 项目。

参数：

- project_file：本地 After Effects 项目的文件对象。文件扩展名应为.aep 或.aet(不支持.aepx)。

返回：布尔值。团队项目是否成功转换返回 True，否则为 false。

### listTeamProjects()

app.project.listTeamProjects()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：返回一个数组，包含当前用户可用的所有团队项目的名称。不包括存档的团队项目。

返回：字符串数组。

### isTeamProjectOpen() 检查团队项目打开

app.project.isTeamProjectOpen(teamProjectName)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查指定的团队项目是否打开。

参数：

- teamProjectName：团队项目名称，字符串。

返回：布尔值。如果指定的团队项目当前处于打开状态返回 True，否则为 false。

### isAnyTeamProjectOpen() 检查团队项目打开

app.project.isAnyTeamProjectOpen()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查当前是否打开任何团队项目。

返回：布尔值。当前是否有任何团队项目处于打开状态返回 True，否则为 false。

### isTeamProjectEnabled() 是否启用团队项目

app.project.isTeamProjectEnabled()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查是否为 After Effects 启用了团队项目。(这几乎总是返回 true。)

返回：布尔值。如果当前启用了团队项目返回 True，否则为 false。

### isLoggedInToTeamProject() 是否登录团队项目

app.project.isLoggedInToTeamProject()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查客户端(After Effects)当前是否已登录到团队项目服务器。

返回：布尔值。如果客户端(After Effects)当前已登录到团队项目服务器返回 True，否则为 false。

### isSyncCommandEnabled() 是否同步团队项目

app.project.isSyncCommandEnabled()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查是否启用了“同步”命令。

返回：布尔值。如果团队项目启用了同步命令返回 True，否则为 false。

### isShareCommandEnabled() 是否共享团队项目

app.project.isShareCommandEnabled()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查是否启用了“共享”命令。

返回：布尔值。如果启用了团队项目“共享”命令返回 True，否则为 false。

### isResolveCommandEnabled() 是否解析团队项目

app.project.isResolveCommandEnabled()

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：检查是否启用了“解析”命令。

返回：布尔值。如果团队项目已启用“解决”命令返回 True，否则为 false。

### resolveConflict() 团队项目版本冲突

app.project.resolveConflict(ResolveType)

::: info 提示
此功能已添加到 After Effects 14.2(CC 2017.1)中
:::

描述：使用指定的解决方法，解决开放团队项目和团队项目服务器上的版本之间的冲突。

参数：ResolveType，要使用的冲突解决方案的类型：

- ResolveType.ACCEPT_THEIRS：获取共享版本。共享版本将替换您的版本。
- ResolveType.ACCEPT_YOURS：保留项目的版本。共享版本未使用。
- ResolveType.ACCEPT_THEIRS_AND_COPY：复制并重命名您的版本，然后获取共享版本。共享版本替换了您的原始版本

返回：布尔值。如果解析成功返回 True，否则为 false。
