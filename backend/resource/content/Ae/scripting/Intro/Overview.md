---
order: 4
title: Overview
---

# 脚本概述

## 版本说明

ES5（CEP) → ES3(脚本） → ES6（表达式）

## After Effects 脚本介绍

脚本是执行一系列命令。使用脚本可执行自动、重复、复杂的计算，甚至使用某些隐藏功能。例如，在 After
Effects 合成中重新排列图层，在文本图层中查找和替换源文本，或在渲染完成后发送电子邮件。

尽管 After
Effects 表达式和脚本都是基于 JavaScript，但表达式和脚本截然不同。表达式无法访问脚本中的信息（例如变量和函数）。脚本是我能干嘛，表达式则是我这个属性能干嘛。但表达式和脚本都基于 JS，因此熟悉一种对理解另一种颇有帮助。

编写脚本的核心是对象。使用 Adobe After
Effects 时，您将创建项目，合成和渲染队列以及它们所包含的元素：素材、图像、solid、图层、蒙版、效果和属性。用脚本术语来说，每个项目都是一个对象。本指南描述了为 After
Effects 项目定义的 ExtendScript 对象。

After Effects 对象由工程，项目，合成，图层和渲染队列组成。每个对象都有独特属性，都有自己的标识（尽管并非所有脚本都可访问）。唯有熟悉 After
Effects 对象才能创建脚本。

::: info 提示
在本指南中，“属性（properties）”称为“属性（attributes）”，避免与 After Effects 自己的属性（properties）混淆。
:::

几乎脚本能完成的工作在 AE 用户界面也能完成。因此编写脚本需要对程序本身及用户界面了解透彻。

## ExtendScript 语言

AE 脚本使用 Adobe
ExtendScript 语言，是 JavaScript 的扩展形式。ExtendScript 根据 ECMA-262 规范实现 JavaScript 语言。After
Effects 脚本引擎支持 ECMA-262 标准的第三版，包括其符号和词汇约定，类型，对象，表达式和语句。ExtendScript 还实现了 E4X
ECMA-357 规范，该规范定义了对 XML 格式的数据的访问。

ExtendScript`$`为 ExtendScript 元素定义了一个全局调试对象，dollar()对象以及一个报告实用工具，即 ExtendScript
Reflection 接口。

**文件和文件夹对象：** 由于路径名的语法在不同的操作系统中区别很大，因此 Adobe
ExtendScript 定义了文件和文件夹对象，以提供与平台无关的对底层文件系统的访问。

**ScriptUI 用户界面模块：** ExtendScript Script
UI 模块提供创建用户界面元素并进行交互的功能。ScriptUI 为 Windows 和 UI 控件元素提供了一个对象模型，可用于为脚本创建用户界面。

**工具和实用程序：**
此外，ExtendScript 还提供了工具和功能，例如本地化实用程序，用于提供不同语言的用户界面字符串值，以及用于在对话框中显示短消息（警报，确认和提示）的全局功能。

**外部通信：** ExtendScript 提供了一个 Socket 对象，可以使 After Effects 脚本与远程系统进行通信。

应用程序间 **通信：** ExtendScript 为所有 Adobe 应用程序提供通用的脚本环境，允许通过脚本进行应用程序间通信。

---

## ExtendScript 工具包（ESTK）

After Effects 包括脚本编辑器和调试器 ExtendScript Toolkit（ESTK），它为创建和测试您自己的脚本提供了方便的界面。

要启动 ESTK，请选择“文件”>“脚本”>“打开脚本编辑器”。

如果选择使用其他文本编辑器来创建，编辑和保存脚本，请确保选择一个应用程序，该应用程序在保存文件时不会自动添加标头信息，并且使用 Unicode（UTF-8）编码进行保存。在许多文本编辑器中，您可以设置首选项以使用 UTF-8 编码进行保存。默认情况下，某些应用程序（例如 Microsoft
Word）将标头信息添加到文件中，这可能导致脚本中出现“第 0 行”错误，从而导致它们失败。

有关 ExtendScript 工具包的详细信息，请参阅《[JavaScript 工具指南》](http://estk.aenhancers.com/)。

---

## .jsx 和.jsxbin 文件扩展名

ExtendScript 脚本文件的区别在于`.jsx`文件名扩展名，是 JavaScript 文件使用的标准扩展名`.js`的变体。After
Effects 脚本必须包含`.jsx`文件扩展名才能被应用程序正确识别。任何带有`.jsx`扩展名的 UTF-8 编码的文本文件都将被识别为 ExtendScript 文件。

您可以使用 ExtendScript 工具包导出扩展名为.jsxbin 的 ExtendScript 文件的二进制版本。这样的二进制文件可能无法与 After
Effects 中的所有脚本集成功能一起使用。

---

## 激活完整的脚本功能

默认是不允许脚本通过网络写入文件或发送或接收通信。

要允许脚本编写文件并通过网络进行通信，请选择“编辑”>“首选项”>“常规”（Windows）或“ After Effects”>“首选项”>“常规”（Mac
OS），然后选择“允许脚本编写文件和访问网络”。

After Effects 脚本会从应用程序生成错误消息。包括错误的性质以及错在哪一行。应用程序遇到脚本错误时，ExtendScript
Toolkit（ESTK）调试器会自动打开。默认情况此功能是关闭的，以便普通用户不会遇到。要激活此功能，请选择“首选项”>“常规”，然后选择“启用 JavaScript 调试器”。

---

## 加载和运行脚本

### 文件>脚本

After Effects 启动后，会在“scripts”文件夹中搜索要加载的脚本。可从“文件”>“脚本”菜单中获取已加载的脚本。

要运行 **已加载** 的脚本，请选择“文件”>“脚本”>“ [脚本名称]”。

如果在 After Effects 运行时编辑脚本，必须保存方可生效。如果在 After
Effects 运行后才把脚本放在 Scripts 文件夹中，必须重新启动 AE，使该脚本出现在“脚本”菜单中。（或者不重启，使用下面的运行脚本文件）

### 文件>脚本>运行脚本文件

要运行尚未加载的脚本，请选择“文件”>“脚本”>“运行脚本文件”，找到脚本，单击“打开”。

### 从命令行，批处理文件或 AppleScript 中运行脚本

如果您熟悉如何从 Windows 中的命令行或 AppleScript，则可以将脚本直接发送到打开的 After Effects 应用程序，从而自动运行该脚本。

要从命令行运行脚本，请从命令行调用 afterfx.exe。使用`-r`开关和脚本的完整路径作为参数运行。该命令不会打开 After
Effects 应用程序的新实例。它在现有实例中运行脚本。

示例（对于 Windows）：

```javascript
afterfx -r c:\script_path\我的脚本.jsx
```

可以使用命令行将脚本的开关绑定到键盘快捷方式。

以下是 Windows 命令行示例，这些示例会将 After Effects 脚本发送到应用程序，而无需使用 After Effects 用户界面执行脚本。

在第一个示例中，您直接在命令行上复制并粘贴 After Effects 脚本，然后运行它。脚本文本在 afterfx.exe -s 命令之后用引号引起来：

```javascript
afterfx.exe -s "alert("向 AE 发送一个警告对话框")"
```

![](https://mir.yuelili.com/2021/07/598c2136064c0d659d8d3b359d16d731.png)![](https://mir.yuelili.com/2021/07/ce93a6485de6e6683731a528181258b2.png)

或者，您可以指定要执行的 JSX 文件的位置。例如：

```javascript
afterfx.exe -r c:\myDocuments\Scripts\我的脚本.jsx afterfx.exe -r "c:\myDocuments\Scripts\我 的 脚 本.jsx" //可以带空格
```

## 如何在 AppleScript 中测试 After Effects 脚本（Mac OS）

\*\*MAC 没用过，不校正

以下是三个示例，将 AE 脚本 JSX 文件发送到应用程序，而不用在 AE 中打开。

第一个示例，将 After Effects 脚本直接复制到“脚本编辑器”中，然后运行。脚本文本出现在 DoScript 命令后的引号内，因此必须使用反斜杠对脚本中的内部引号进行转义，如下所示

```javascript
tell application "Adobe After Effects CS6"
DoScript "alert(\"向AE发送一个警告对话框\")"
end tell
```

或者可以显示一个对话框，询问要执行的 JSX 文件的位置，如下所示：

```javascript
set theFile to choose file
tell application "Adobe After Effects CS6"
DoScript theFile
end tell
```

::: info 提示

该文档不正确，此实例中的正确调用是 DoScriptFile

:::

最后，当您直接编辑 JSX 脚本并将其发送到 AE 进行测试或运行时，该脚本可能最有用。为了有效地使用它，您必须输入包含打开的 JSX 文件的应用程序（在本示例中为 TextEdit）。如果您不知道应用程序的正确名称，请输入最佳猜测以替换“ TextEdit”，AppleScript 会提示您找到它。

只需突出显示要运行的脚本文本，然后激活此 AppleScript：

```javascript
(\*
This script sends the current selection to After Effects as a script.
*)


tell application "TextEdit"
    set the_script to text of front document
end tell

tell application "Adobe After Effects CS6" activate
    DoScript the_script
end tell
```

## 在应用程序启动或关闭期间自动运行脚本

在 Scripts 文件夹中有名为 Startup 和 Shutdown 的文件夹。After Effects 分别在启动和退出时按字母顺序在这些文件夹中自动运行脚本。

在启动（Startup）文件夹中，可以放置 ​​ 在 AE 启动时执行的脚本。会在应用程序初始化并加载插件之后执行。

脚本共享一个 **全局** 环境，因此在启动时执行的任何脚本，都可以定义用于所有脚本的变量和函数。在任何情况下，一旦运行包含变量和函数的脚本，进行变量和函数定义，便会 After Effects 运行期间保留在后续脚本中。退出 AE 后，才会清除此类全局定义的变量和函数。\*\* 确保为脚本中的变量赋予唯一名称，以便脚本不会无意间重新分配 AE 运行中的全局变量。

还可以将属性添加到现有对象（例如[Application 对象](https://ae-scripting.docsforadobe.dev/general/application/#application)），以将应用程序扩展为其他脚本。

程序退出时，将执行 Shutdown 文件夹中的脚本。这是在项目关闭之后但在其他应用程序关闭之前发生。

## 从“窗口”菜单运行脚本

可以从“窗口”菜单的底部获得 ScriptUI Panels 文件夹中的脚本。如果编写脚本的可停靠面板用于提供用户界面，则应将脚本放在 ScriptUI 文件夹中。ScriptUI 面板的工作方式与 After Effects 用户界面中的默认面板相同。

ScriptUI Panels 的脚本使用 this 代表面板对象，而不是创建 Window 对象并向其添加控件。例如，以下代码将按钮添加到面板：

```javascript
var myPanel = this;
myPanel.add("button", [10, 10, 100, 30], "Tool #1");
```

如果脚本在函数中创建用户界面，则不能使用 this，因为它将引用函数本身，而不是面板。在这种情况下，应将 this 对象作为参数传递给函数。例如：

```javascript
function createUI(thisObj) {
  var myPanel = thisObj;
  myPanel.add("button", [10, 10, 100, 30], "Tool #1");
  return myPanel;
}

var myToolsPanel = createUI(this);
```

不能从“文件”>“脚本”>“运行脚本文件”运行带 UI 的脚本。要使脚本与 Window 对象（可从 File> Scripts 菜单访问）或 Mac 面板（可从窗口菜单访问）一起使用，要先检查是否为 Panel 对象。例如：

```javascript
function createUI(thisObj) {
  var myPanel =
    thisObj instanceof Panel
      ? thisObj
      : new Window(
          "palette",
          "My Tools",

          [100, 100, 300, 300]
        );

  myPanel.add("button", [10, 10, 100, 30], "Tool #1");

  return myPanel;
}

var myToolsPanel = createUI(this);
```

## 停止正在运行的脚本

当 After Effects 或脚本的用户界面具有焦点时，可以通过按 Esc 或 Cmd +句号（Mac OS）来停止脚本。如果脚本正在处理大量数据，可能响应不及时。
