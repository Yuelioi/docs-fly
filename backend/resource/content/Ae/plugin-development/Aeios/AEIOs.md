---
order: 2
title: AEIOs
---

# AEIOs

AEIOs 是执行媒体文件导入(Import)和/或导出(Output)的 AEGPs。AEIOs 为一个特定类型的文件做 After Effects(或 After Effects 的插件)通常会做的一切。在导入方面，AEIOs 可以打开现有的文件，管理文件特定的解释选项，并以 AEGP_SoundWorld 和 PF_EffectWorld 格式向 After Effects 提供文件中的音频和帧。此外，AEIOs 可以交互式地创建文件，要求用户提供他们想要的设置，而不是从文件中读取它们。在输出方面，AEIOs 可以创建和管理渲染队列项目的输出选项，创建输出文件并将帧(由 After Effects 提供的 PF_EffectWorlds)保存到这些文件中。

AEIOs 可以处理未压缩的视频，像素按 ARGB 顺序从低字节到高字节。每个通道的像素可以是 8 位、16 位或 32 位浮点。AEIOs 必须处理他们自己对任何支持的编解码器的压缩/解压缩。

## AEIO, 还是 AEGP?

AEIO 向 After Effects 提供像素和音频数据。

如果你要为代表时间轴或项目格式的文件格式(参考 After Effects 或其他已安装的 AEIO 支持的文件格式)编写一个导入器/导出器，并将其命令添加到导入/导出子菜单中。

## 使用 AEIO 导入, 还是使用 MediaCore 导入器?

After Effects 支持 MediaCore 导入器插件。MediaCore 是一套从 Premiere Pro 发展而来的共享库，因此 MediaCore 的 API 在[Premiere Pro SDK](http://ppro-plugin-sdk.aenhancers.com/)中有所描述。

只有 MediaCore 导入器插件支持导入器优先系统。优先级最高的进口商首先获得进口文件的机会，如果特定的进口文件不被支持，那么下一个优先级最高的进口商将有机会尝试进口该文件，以此类推。MediaCore 进口商不能将文件进口推迟到 AEIO。因此，如果你的目标是接管 After Effects 已经提供插件的任何文件类型的文件处理，你需要开发一个 MediaCore 导入器插件。

另一方面，只有 AEIO 可以在 Interpret Footage < Main > More Options 对话框中显示设置对话框。

如果上述制约因素还没有回答你是需要建立一个 AEIO 还是 MediaCore 导入器，那么你很可能想建立一个 MediaCore 导入器，它可以用于整个视频和音频应用程序，包括 Premiere Pro、Media Encoder、Prelude、SpeedGrade 和 Audition。

## 如何工作

在其入口函数函数中，AEIO 填充了一个函数指针的结构，其中包含它希望在响应某些事件时调用的函数名称。许多这些函数钩子是可选的。

## AE 会作何反应

对于许多 AEIO 钩子函数，你可以要求 After Effects 执行默认处理(这种能力在每个钩子的描述中都有注明)。

除非你有令人信服的理由，否则就从函数中返回`AEIO_Err_USE_DFLT_CALLBACK`，让 After Effects 来做这些工作。

这也是在开始执行之前学习调用顺序的好方法。

## 注册 AEIO

在你的插件的入口函数函数中，填充一个 AEIO_ModuleInfo，描述 AEIO 支持的文件类型，和一个 AEIO_FunctionBlock 结构，指向你的文件处理函数。对于其中一些函数，你可以通过返回 AEIO_Err_USE_DFLT_CALLBACK 来依赖 After Effects 的默认行为。但是，你仍然必须提供一个符合要求的签名的函数，这样做。一旦你填好了这两个结构，从[AEGP_RegisterSuites5](../aegps/aegp-suites.html)调用`AEGP_RegisterIO()`。

在你传递给注册调用的 AEIO_ModuleInfo 中，你提供了文件类型和描述信息，After Effects 在导入对话框中使用这些信息，用于 Windows 的 "类型文件 "下拉，或 MacOS 的启用下拉。从 CS6 开始，文件扩展名不能超过三个字符，尽管我们有一些内置的导入器，扩展名更长。

## InSpec, OutSpec

在大多数与导入有关的函数中，会传递一个`AEIO_InSpecH`。在大多数与输出相关的函数中，会传递一个`AEIO_OutSpecH`。

这些神秘的句柄是什么？这些不透明的数据句柄可以与[AEGP_IOInSuite5](new-kids-on-the-function-block.html) 和[AEGPIOOutSuite4](new-kids-on-the-function-block.html) 一起使用，来设置或查询关于导入或输出的信息。

例如，在导入时，当调用 AEGP_IOInSuite 中的`AEGP_SetInSpecDimensions`时，你会使用`AEIO_InSpecH`。

而在出口时，当在`AEGP_IOOutSuite`中调用`AEGP_GetOutSpecDimensions`时，你将使用`AEIO_OutSpecH`。所以使用这些句柄与 After Effects 交换关于输入或输出的细节信息。
