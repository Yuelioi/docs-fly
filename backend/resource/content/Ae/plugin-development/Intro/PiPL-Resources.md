---
title: PiPL-Resources
order: 11
---

# PiPL 资源

起源于 20 多年前的 Adobe Photoshop，Plug-In Property Lists(插件属性列表)，或称作 PiPLs，是在不执行插件的情况下，提供有关插件行为的基本信息的资源。PiPLs 在 After Effects 中已被`PF_Cmd_GLOBAL_SETUP`(见[全局入口指令](../effect-basics/command-selectors.html)(#effect-basics-command-selectors-global-selectors)和动态 outflags 所取代。然而，出于历史原因，在 `PF_Cmd_GLOBAL_SETUP`过程中指示的行为必须与 PiPL 中的行为一致。

PiPL 指定了一个插件的入口函数，显示名称，以及插件的匹配名称。匹配名称是一个唯一的、不变的标识符，与插件的显示名称不同，显示名称可以动态地改变。从 CC 开始，显示名称的长度可达 47 个字符。以前则是 31 个字符。

为了实现跨平台的兼容性，请为插件的 macOS 和 Windows 版本使用一个.r 文件，就像范例那样。PiPL 属性必须始终以 macOS 特有的字节顺序排列。在 Windows 上，PiPLs 是通过 pipltool.exe 处理.r 文件来编译的，它将.r 文件转换为二进制.rc 文件。Windows 示例项目都包含自定义构建步骤，使用跨平台的.r 文件和我们的 cnvtpipl.exe 命令行工具，生成一个.rc 文件。以现有的样本插件为基础进行开发，才能正确实现构建步骤。

## 入口函数

插件入口函数是通过 PiPL 在 Windows 和 macOS 上输出的。如果该插件支持多个平台，那么必须在 PiPL 中定义多个入口函数。不需要 Windows .def 文件或手动导出，除非你还指定了一些其他操作系统的入口函数。

entry.h 中定义的宏(在\SDKExamplesHeaders 目录下)负责导出每个样本的入口函数函数。所有样本项目的入口函数函数对效果来说是 EffectMain()，对 AEGP 来说是 EntryPointFunc()。

## PiPL 资源以及 Microsoft Visual Studio

要用 pipltool 生成的资源来使用来自 Microsoft Visual Studio .NET 的资源。

将自定义构建步骤的输出包含到 Microsoft 生成的. rc 文件中。

```cpp
// 在文件WhizBang.rc中，由. NET生成。
#include "WhizBang_PiPL_temp.rc" // pipltool.exe's output

```

如果修改一个样本插件，请将 pipltool.exe 生成的文件名改为 WhizBang_PiPL_temp.rc 这样的名字，否则每次构建时都会覆盖微软的资源；这不是好事。

## 复数个 PiPLs

在同一个文件中包含多个插件(包括 AEGP 和特效)，使用多个 PiPLs 是可行的，但不推荐。如果在同一个文件中同时有 AEGP 和特效的 PiPLs，那么 AEGPs 必须排在第一位。

没有其他主机(甚至 Premiere Pro)支持在同一个.dll 或代码片段中的多个 PiPLs 指向多个效果。另外，如果你需要更新一个插件，你真的想把你所有的插件都换成新的版本吗？我们建议每个代码片段使用一个 PiPL 和一个插件。

## 超级秘密 PiPL Bit

对于那些使用 C++的人来说，必须一直保持你的插件被加载(以避免你的 v-table 被破坏，以及其他危害)，将 PiPL 的 AE_Reserved_Info 成员设置为 8。多年来，我们一直很严格，坚持要求插件作为良好的内存公民，并对卸载作出优雅的反应。我们知道在有些情况下，卸载时没有任何警告会毁掉一个插件(和 v-tables)，所以我们提供了这个变通办法。善待自己，严格管理内存，只用你的力量去做好事。

## 为什么需要了解这些

你不需要，After Effects 需要。如果你听从我们的建议，以 SDK 样本为基础，可以简单地修改包含 PiPL 定义的.r 文件，插件资源将在你下次构建时自动更新。感受爱。或者，如果你曾经修改过自定义的构建步骤，请感受一下痛苦。
