---
title: Calling-Sequence
order: 3
---

# 调用序列

与所有的 AEGP 一样，在启动过程中，插件的 PiPL 中导出的入口函数函数被调用。在这个函数中，AEIO 必须提供所需函数的函数指针，并描述它们的能力，然后将适当的结构传递给[AEGP_RegisterIO()](../aegps/aegp-suites.html)(#aegps-aegp-suites-aegp-registersuites)。

## 导入

当用户在文件导入对话框中选择一个由你的 AEIO 处理的类型的文件时，它的[AEIO_VerifyFileImportable()](new-kids-on-function-block.html) 函数将被调用；它对用户导入的每个此类文件再次被调用。[AEIO_InitInSpecFromFile()](new-kids-on-the-function-block.html) 将为每个文件调用；解析文件，并使用各种设置函数向 After Effects 描述它。同时，构建与文件相关的任何选项数据，并使用[AEGP_SetInSpecOptionsHandle()](new-kids-on-the-function-block.html) 保存这些数据。

然后 After Effects 调用插件的[AEIO_GetInSpecInfo()](new-kids-on-the-function-block.html)函数，以获得关于文件的描述性文本，显示在项目窗口。正如这个函数的描述中所指出的，它也可以为文件夹调用；我们建议，如果文件没有有效的选项数据，你什么都不做，不返回错误(这就是我们的 AEIO 所做的)。

然后发送[AEIO_CountUserData()](new-kids-on-the-function-block.html) ；如果 AEIO 表明有用户数据存在，[AEIO_GetUserData()](new-kids-on-the-function-block.html) 将随之而来。然后 After Effects 将通过发送[AEIO_DrawSparseFrame()](new-kids-on-the-function-block.html)要求插件绘制一帧视频(用于项目窗口缩略图)。

一旦支持的文件被添加到一个合成中，用户互动将产生对`AEIO_DrawSparseFrame()`和[AEIO_GetSound()](new-kids-on-the-function-block.html) 的调用。

当项目被保存时，如果有与 AEIO_InSpec 相关的选项数据，After Effects 将发送[AEIO_FlattenOptions()](new-kids-on-the-function-block.html) ，在此期间，AEIO 解析选项数据，并创建一个不包含引用外部内存的表示。同样，任何 AEIO_OutSpec 选项数据的存在将导致[AEIO_GetFlatOutputOptions()](new-kids-on-the-function-block.html) 被发送。

## 导出

如果用户在渲染队列中添加了一个项目并选择了 AEIO 支持的输出格式，[AEIO_InitOutputSpec()](new-kids-on-the-function-block.html) 将被发送。使用各种获取函数来获取关于输出设置的信息，并使用[AEGP_SetOutSpecOptionsHandle()](new-kids-on-the-function-block.html)存储任何相关的信息，随后使用`AEIO_GetFlatOutputOptions()`。[AEIO_GetDepths()](new-kids-on-the-function-block.html) 被发送，这样 After Effects 可以确定 AEIO 支持哪些输出像素位深度。[AEIO_GetOutputInfo()](new-kids-on-the-function-block.html) 被发送，以便文件名、类型和子类型信息可以显示在输出模块的细节中。

当用户点击格式选项按钮时，在渲染队列中，[AEIO_UserOptionsDialog()](new-kids-on-the-function-block.html) 被调用。

当用户实际点击 "渲染 "按钮时，[AEIO_SetOutputFile()](new-kids-on-the-function-block.html) 将被调用，接着是[AEIO_GetSizes()](new-kids-on-the-function-block.html) (你的 AEIO 负责判断目的地是否有足够的磁盘空间可用)。

在视频帧被发送之前，[AEIO_StartAdding()](new-kids-on-the-function-block.html) 被发送，以便 AEIO 打开文件柄并写出文件头。如果 AEIO 支持视频或音频格式，[AEIO_AddSoundChunk()](new-kids-on-the-function-block.html) 将为每个音频块发送，以及每个视频帧的[AEIO_AddFrame()](new-kids-on-the-function-block.html)

如果 AEIO 支持静止图像序列，则重复调用 `AEIO_OutputFrame()`。 After Effects 发送要输出帧的 `PF_EffectWorld` 表示。

AEIO_WriteLabels() 被调用（对于每一帧），让插件有机会写出字段和 alpha 解释信息。 当没有更多帧（或音频）要输出时发送 `AEIO_EndAdding()`。 关闭输出文件。
