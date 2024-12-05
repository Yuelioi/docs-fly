---
title: Application
order: 1
---
# APP 应用

## 说明

描述：最顶层的对象，也就是 AE 应用/软件本身。用于访问应用的对象以及应用设置。

属性：AE 程序的对象。比如工程文件、AE 版本号等等

方法：软件层级的操作，比如创建项目、打开现有项目、控制监视文件夹模式、清除内存并退出 AE。当 AE 退出时，将关闭打开的项目，提示用户根据需要保存或放弃更改，并根据需要创建项目文件。

## 属性篇

### activeViewer 监视器

全名：app.activeViewer

描述：活动(“合成”，“图层”或“素材”)预览窗口对象。如果没有打开预览窗口，则返回 null。

返回：监视器对象，一般配合预览窗口对象的其他属性使用；只读。

什么是活动：

- 图一：信息面板（多组面板，处于激活状态，并且是当前选择项）
- 图二：查看器（单个面板：处于激活状态）

![](https://cdn.yuelili.com/2021/20210824221123.png)

```javascript
alert(app.activeViewer); //如果打开查看器了，屏幕会显示[object Viewer]
```

查看器的其他属性，请参考 其他>查看器对象

### availableGPUAccelTypes GPU 加速类型

全名：app.availableGPUAccelTypes

::: info 提示
在 AfterEffects14.0(CC2017)中添加了此功能

:::

描述：与 app.project.gpuAccelType 结合使用，可为“项目设置”>“视频渲染和效果”>“使用范围”设置加速方式。

![](https://cdn.yuelili.com/20210824221417.png)

返回：GpuAccelType 数组；如果未打开任何监视器，则为 null。只读。

- CUDA
- Metal
- OPENCL
- SOFTWARE

个人猜测：使用 alert(app.availableGPUAccelTypes)
返回的是 1813 和 1816，我的电脑能开启 CUDA 和 SOFTWARE，因此自上到下的顺序数字代码为 1813、1814、1815、1816

示例：下面的示例代码检查当前计算机的可用 GPU 加速类型，并将其设置为 Metal(如果可用)。

```javascript
// app.availableGPUAccelTypes ：根据当前系统，返回当前可用的GPU加速类型. //可用于检测GPU加速类型.
var newType = GpuAccelType.METAL; // 设置前，请先检查当前系统上可用的GPU加速类型。
var canSet = false;
var currentOptions = app.availableGPUAccelTypes; // 遍历加速类型
for (var op in currentOptions) {
  if (currentOptions[op] === newType) {
    canSet = true;
  }
}
if (canSet) {
  // 可用，则设置加速类型.
  app.project.gpuAccelType = newType;
} else {
  alert("本机无法使用Metal哦");
}
```

### buildName 内部版本名

全名：app.buildName

描述：正在运行的 After Effects 的内部版本名称，用于 Adobe 内部测试和故障排查。也可以在＂帮助-关于＂中查看

![](https://cdn.yuelili.com/20210824221716.png)![](https://cdn.yuelili.com/20210824221751.png)

类型：字符串；只读。

```javascript
alert(app.buildName); // 本机AE：17.7x45
```

### buildNumber 内部版本号

全名：app.buildNumber

描述：正在运行的 After Effects 的内部版本号，用于 Adobe 内部测试和故障排除。

类型：整数；只读。

```javascript
alert(app.buildNumber); // 本机AE：45 （版本号 17.7x45 后面的数字）
```

### disableRendering 实时渲染

全名：app.disableRendering

::: info 提示
在 After Effects 16.0(CC 2019)中添加了此功能
:::

描述：如果为 false(默认值)，则渲染按常规进行。设置为 true 可禁用渲染，等同于打开 Caps Lock。

类型：布尔值；读/写。

```javascript
alert(app.disableRendering); // 读：返回false，当前未禁用
app.disableRendering = True; // 写：强制禁止渲染
```

### effects 效果

全名：app.effects

描述：应用程序中可用的效果。(包括预设）

类型：数组，数组中每个元素(也就是单个效果)包含以下属性；只读：

- displayName ：本地化名称，字符串。比如在中文版叫"高斯模糊"，在英文版叫"Gaussian Blur"。
- category ：本地化类别。没有分类，则为空。比如在中文版叫""颜色校正"，在英文版叫"Color Correction"。
- matchName： 内部名称（唯一），又称匹配名，字符串。换语言，换版本都不会变。
- version ：效果版本号，字符串。可能与插件供应商在效果的“关于”显示的版本号不同。

![](https://cdn.yuelili.com/20210824222547.png)

示例

```javascript
alert(app.effects[12].displayName); //CC Color Offset;
alert(app.effects.length); //461
```

```javascript
//获取本机所有效果的匹配名，！注意 需要等待20秒左右，而且会布满整个屏幕！
var effects = app.effects;
var fxName = new Array();
for (i = 0; i < app.effects.length; i++) {
  fxName.push(effects[i].matchName);
}
alert(fxName);
```

### exitAfterLaunchAndEval 命令行 AE 退出状态

全名：app.exitAfterLaunchAndEval

描述：Windows 的命令行执行脚本时的特有属性，在 Mac OS 中无效。从命令行启动应用程序时，用–r 或–s 运行脚本(从文件或字符串运行)。

说明：为 true 时，在脚本运行后退出 After Effects；为 False，则 AE 继续保持打开状态。

类型：布尔值 。读/写。

示例：在 AE 的 Support
Files 文件中打开 cmd，然后输入以下代码（大致意思是运行 Labels 脚本），再运行该脚本后，AE 会自动退出（因为 exitAfterLaunchAndEval 为 True）

```javascript
AfterFX.exe -r H:\adobe\Adobe After Effects 2020\Support Files\Scripts\ScriptUI Panels\Labels.jsx
```

![](https://cdn.yuelili.com/20210824223422.png)

### exitCode 脚本状态码

全名：app.exitCode

描述：在外部(即从 windows 命令行或 Apple Script)执行脚本时使用的数字状态代码。

- 在 Windows 中，在命令行上启动 After Effects(使用 afterfx 或 afterfx –m)并使用选项（–r 或–s）指定了脚本时，将在命令行上返回该值。
- 在 Mac OS 中，脚本在 AppleScript 中 DoScript 后，返回该值。

在 Mac OS 和 Windows 中，在每次脚本计算时，默认为 0（EXIT_SUCCESS）。你可以给脚本的这个参数设置为一个正整数，用于标记错误。

类型：整数；读/写。

示例：封装到某个代码块里，只要返回这个值，就是这部分代码报错

```javascript
if (XXXX){
  运行 XX 代码 ;
  app.exitCode = 2;
} // 报错时，该方法返回2
```

### isoLanguage 语言环境

全名：app.isoLanguage

描述：显示当前 After Effects 的语言环境(语言和区域名称)。

::: info 提示
$.locale 返回操作系统语言，而不是 AfterEffects 应用程序的语言。
:::

类型：字符串; 只读。一些常见的值包括：

- en_US 英语(美国)
- zh_CN 中文(中国)
- de_DE 德语(德国)
- es_ES 西班牙语(西班牙)
- fr_FR 法语(法国)
- it_IT 意大利语(意大利)
- ja_JP 日语(日本)
- ko_KR 朝鲜语(韩国)

示例：根据不同语言环境，执行不同代码

```javascript
var lang = app.isoLanguage;
if (lang === "en_US") {
  alert("当前使用的是英文版");
} else if (lang === "fr_FR") {
  alert("当前使用的是法文版");
} else {
  alert("反正不是英文或法文版，具体的自己设置");
}
```

### isRenderEngine 是否仅作为渲染引擎

全名：app.isRenderEngine

描述：如果 After Effects 作为渲染引擎运行，则为 True。

猜测：比如我在用 PR 动态链接 AE，这时候 AE 就是渲染引擎？

类型：布尔值；只读。

### isWatchFolder 监视文件夹

全名：app.isWatchFolder

描述：如果当前显示“监视文件夹”对话框，并且 AE 正在监视要渲染的文件夹，则为 True。

类型：布尔值；只读。

如何打开：文件菜单 - 监视文件夹

![](https://cdn.yuelili.com/20210824224830.png)

### memoryInUse 占用内存

全名：app.memoryInUse

描述：此应用程序当前使用的内存，字节数。

等同于：首选项-内存-显示内存详细信息

类型：数值；只读。

![](https://cdn.yuelili.com/20210825001243.png)

```javascript
alert(app.memoryInUse / 1024 / 1024);
//返回348.1654...（单位是兆的话要除以俩个1024）
```

### onError 脚本错误码

全名：app.onError

描述：发生错误时调用的回调函数的名称。通过创建一个函数并设置该属性，您可以方便地查找错误所在地。例如，您可以关闭并重新启动 AE，并在日志文件中查看错误(如果在渲染过程中发生了该错误)。请参阅 RenderQueue.render()。

个人理解：发生错误了（不是书写错误，是 AE 运行或者其他错误），会调用某个函数，这个函数可以返回错误信息，从而找到错误代码位置

类型：函数名称，字符串，如果未分配任何函数，则为 null。读/写。

示例

```javascript
// 发生错误，调用一个名为err的函数。
function err(errString) {
  alert(errString);
}
app.onError = err;
```

### preferences 首选项

全名：app.preferences

描述：当前加载的 AE 首选项。请参阅首选项对象。一般要配合下属函数/属性使用

类型：首选项对象；只读。

### project 项目

全名：app.project

描述：当前加载的项目。请参阅项目对象。一般要配合下属函数/属性使用

类型：项目对象；只读。

### saveProjectOnCrash 崩溃保存设置

全名：app.saveProjectOnCrash

描述：如果为 true(默认值)，则 AE 会尝试弹出对话框，如果 AE 崩溃，该对话框会提示你保存当前项目。设置为 false 则取消显示此对话框，只退出不保存。

类型：布尔值；读/写。

```javascript
app.saveProjectOnCrash = false; //你完了 AE崩了就炸了
```

### settings 加载设置

全名：app.settings

描述：当前加载的设置。请参阅[设置对象](https://aescript.yuelili.fun/setting-object.html)。一般要配合下属函数/属性使用

类型：设置对象；只读。

```javascript
//只能获得在使用脚本过程添加设置，不能修改After Effects本身的默认设置
app.settings.saveSetting("SettingA", "ValueA", 100);
alert(app.settings.getSetting("SettingA", "ValueA"));
```

### version 版本号

全名：app.version

::: info 提示
此功能于 After Effects 12.0(CC)添加
:::

描述：显示正在运行的 After Effects 版本号

类型：字符串; 只读。

示例

```javascript
var ver = app.version;
alert("正在运行 " + ver + " 版本的AE"); //正在运行 17.5.1x47 版本的AE
```

## 方法篇

### activate() 窗口最大化

全名：app.activate()

描述：窗口最大化，并将窗口置顶。

参数：无。

返回：无。

```javascript
app.activate(); //真的直接弹出来了
```

### beginSuppressDialogs() 禁用脚本错误对话框

全名：app.beginSuppressDialogs()

描述：用户界面不显示脚本报错对话框。使用 app.endSuppressDialogs()恢复错误对话框。

参数：无。

返回：无。

### beginUndoGroup() 撤销组起始

全名：app.beginUndoGroup(undoString)

描述：撤消组的开始标记。撤消组相当于打包好几个步骤，作为一个整体，这样点一次撤销，可以撤销好几步，(“编辑”-“撤消/重做”。app.endUndoGroup()是撤销组的结尾标记。

等同：编辑 - 撤消（ctrl z）

参数：

- undoString：撤销组名称，字符串。

返回：无。

```javascript
//同时创建三个方形，如果撤销的话，就撤销3个 app.beginUndoGroup("AddSquare");
myComp = app.project.item(1);
myComp.layers.addSolid([1.0, 1.0, 0], "square1", 50, 50, 1);
myComp.layers.addSolid([1.0, 1.0, 0], "square2", 50, 50, 1);
myComp.layers.addSolid([1.0, 1.0, 0], "square3", 50, 50, 1);
app.endUndoGroup();
```

### cancelTask() 删除延迟任务

全名：app.cancelTask(taskID)

描述：从计划延迟执行的任务队列中删除指定的任务。详情见 app.scheduleTask()

参数：

- taskID，标识任务名，整数。由 app.scheduleTask()返回 。

返回：无。

### endSuppressDialogs() 显示脚本报错框

全名：app.endSuppressDialogs(alert)

描述：在用户界面中继续显示脚本报错框。默认情况就是显示的(true)；只有先用 app.beginSuppressDialogs()禁用报错框，才可以调用此方法恢复显示。

参数：

- alert ：布尔值。

返回：无。

### endUndoGroup() 撤消组结束

全名：app.endUndoGroup()

描述：撤消组的结束标记。撤消组相当于打包好几个步骤，作为一个整体，这样点一次撤销，可以撤销好几步，(“编辑”-“撤消/重做”。app.beginUndoGrou()是撤销组的开始标记。

等同：编辑 - 撤消（ctrl z）

报错：没有使用 beginUndoGroup()方法时，调用本方法会报错，因为他俩是一体的。

返回：无。

```javascript
//同时创建三个方形，如果撤销的话，就撤销3个
app.beginUndoGroup("AddSquare");
myComp = app.project.item(1);
myComp.layers.addSolid([1.0, 1.0, 0], "square1", 50, 50, 1);
myComp.layers.addSolid([1.0, 1.0, 0], "square2", 50, 50, 1);
myComp.layers.addSolid([1.0, 1.0, 0], "square3", 50, 50, 1);
app.endUndoGroup();
```

### endWatchFolder() 结束监视文件夹模式

全名：app.endWatchFolder()

描述：结束结束监视文件夹模式。文件菜单 - 监视文件夹

参数：无。

返回：无。

以下作为参考

- app.watchFolder()
- app.parseSwatchFile()
- app.isWatchFolder
- app.executeCommand
- app.executeCommand(id)

### executeCommand 调用菜单 ID

app.executeCommand(id)

描述：菜单命令具有单独的 ID 号，可用作该方法的参数。对于 API 中未包含的某些功能，这是访问它们的唯一方法。

说明：app.findMenuCommandId 方法，可以用来查找命令的 ID 号。

其他：https://www.provideocoalition.com/after-effects-menu-command-ids/
有不少已知编号 ID 的列表。

参数：

- id：菜单命令的 ID 号，整数。

返回：无。

示例

```javascript
// 调用“转换为贝塞尔”命令
app.executeCommand(4162);
```

### findMenuCommandId 查找菜单 ID

全名：app.findMenuCommandId(Command)

描述：应用程序中的菜单命令具有单独的 ID 号，可以用作 app.executeCommand 命令的参数。对于 API 中未包含的某些功能，这是访问它们的唯一方法。

访问https://www.provideocoalition.com/AECC2015_MenuIDs_v1_0_1-1.pdf查看更多信息，以及已知编号的列表。

参数：

- Command ：菜单命令的文本，与用户界面中显示的完全相同，字符串。

返回：整数，菜单命令的 ID 号。

示例：打开标尺

![](https://cdn.yuelili.com/20210824231633.png)

```javascript
id = app.findMenuCommandId("显示标尺"); //查找显示标尺的id
app.executeCommand(id); //调用该id命令
```

### newProject() 新建工程

全名：app.newProject()

描述：在 After Effects 中创建一个新项目，等同于“文件 - 新建 -
新建项目”。如果已有项目，则提示用户保存。如果用户点取消，则不会创建新项目，并返回 null。

参数：无。

返回：一个新的 Project 对象；如果未创建新项目，则为 null。

示例

```javascript
//不保存现有工程，直接新建
app.project.close(CloseOptions.DO_NOT_SAVE_CHANGES);
app.newProject();
```

### open() 打开项目

全名：app.open([file])

描述：打开项目。

参数：

- file ：选填， 文件对象。如果未提供，则该方法弹出对话框，让用户自己选择。

返回：新 Project 对象；如果用户取消选择对话框，则为 null。

示例：在当前目录下找 test.aep，如果有，则打开，并且弹出文件名

```javascript
var my_file = new File("./test.aep");
if (my_file.exists) {
  var new_project = app.open(my_file);
  if (new_project) {
    alert(new_project.file.name);
  }
}
```

### parseSwatchFile() 监视文件夹加载颜色样本数据

全名：app.parseSwatchFile(file)

描述：从 Adobe Swatch Exchange(ASE)文件加载颜色样本数据。

说明：文件 - 监视文件夹 没咋用过，不细研究。

参数：

- file ：文件路径，ExtendScriptFile 对象。

返回：样本数据，采用以下格式：

| data.majorVersion data.minorVersion | ASE 版本号。                                              |
| ----------------------------------- | --------------------------------------------------------- |
| data.values                         | 样本值数组。                                              |
| SwatchValue.type                    | “ RGB”，“ CMYK”，“ LAB”，“Gray”，其中之一         |
| SwatchValue.r/g/b                   | type = "RGB"时，颜色值在[0.0~1.0]范围内。0，0，0 是黑色。 |

```javascript
示例：SwatchValue.r //返回 0.5
SwatchValue.c/m/y/k | 当 type=“ CMYK”时，颜色值在[0.0~1.0]范围内。0，0，0，0 是白色。
示例：SwatchValue.c //返回 0.5
SwatchValue.l/a/b | 当 type = "LAB"时，颜色值 L 在[0.0~1.0]范围内。a/b 在[-128.0 .. +
128.0]范围内，0，0，0 是黑色
示例：SwatchValue.a //返回 0.5
SwatchValue.value | 当 type = “Gray”时，颜色值在[0.0~1.0]范围内。0.0 为黑色
示例：SwatchValue.value //返回 0.5
```

### pauseWatchFolder() 监视文件夹继续搜索项目

全名：app.pauseWatchFolder(pause)

描述：在目标监视文件夹中暂停或继续搜索项目。

说明：文件 - 监视文件夹 没咋用过，不细研究。

参数：

- pause ：为 True 时暂停，为 false 时继续。

返回：无。

也可以看看

- app.isWatchFolder
- app.watchFolder()
- app.endWatchFolder()

### purge() 清理缓存

全名：app.purge(target)

描述：清理各种缓存。跟编辑菜单中的＂清除＂选项差不多。如果想清除磁盘缓存，请看脚本案例-《清除磁盘缓存》

参数：

- target：要清除的类型；PurgeTarget，枚举值：
  - PurgeTarget.ALL_CACHES：清除所有物理缓存。
  - PurgeTarget.UNDO_CACHES：清除撤消缓存数据。
  - PurgeTarget.SNAPSHOT_CACHES：清除所有合成/图层快照缓存。
  - PurgeTarget.IMAGE_CACHES ：清除图像缓存。

返回：无。

### quit() 退出应用

全名：app.quit()

描述：退出 After Effects 应用程序。如果没保存，会提示你保存文件

参数：无。

返回：无。

### scheduleTask() 脚本延迟执行

全名：app.scheduleTask(stringToExecute, delay, repeat)

描述：延迟执行某段 JavaScript。

参数：

- stringToExecute：要执行的 JavaScript 字符串。
- delay：要等待的毫秒数。浮点值。
- repeat：如果为 true，重复执行脚本，每次执行之间具有相同延迟。如果为 false，则该脚本仅执行一次。

返回：整数，此任务的唯一标识符，可通过 app.cancelTask()取消。

示例：每 5 秒弹一下窗

```javascript
app.scheduleTask('alert("3点了，饮茶先")', 5000, true);
```

### setMemoryUsageLimits() 内存使用限制

全名：app.setMemoryUsageLimits(imageCachePercentage, maximumMemoryPercentage)

n：32 位 Windows 为 2 GB，64 位 Windows 为 4 GB，Mac OS 为 3.5 GB。

描述：设置内存使用限制，相当于首选项“内存和缓存”。如果当前 RAM 小于 n(默认 n 个 GB)，则为当前 RAM 百分比。否则为 n 的百分比。

参数：

- imageCachePercentage ：缓存的内存百分比，浮点值。
- maximumMemoryPercentage ：最大可用百分比，浮点值。

返回：无。

### setSavePreferencesOnQuit() 退出保存首选项

全名：app.setSavePreferencesOnQuit(doSave)

描述：设置或清除标志，该标志影响关闭 AE 时是否保存首选项。请参阅首选项对象。

参数：

- doSave ：如果为 true，则退出时保存首选项；为 false，则不保存。

返回：无。

### watchFolder() 启动监视文件夹

全名：app.watchFolder(folder_object_to_watch)

描述：启动指向指定文件夹的监视文件夹(网络渲染)过程。

参数：

- folder_object_to_watch：要监视的文件夹路径，ExtendScript Folder 对象。

返回：无。

示例

```javascript
var theFolder = new Folder("c:/tool");
app.watchFolder(theFolder);
```

也可以看看

- app.endWatchFolder()
- app.parseSwatchFile()
- app.isWatchFolder
