---
title: New-Kids-On-The-Function-Block
order: 5
---

# 新功能块

在其主入口函数函数中，每个 AEIO 插件必须填写`AEIO_FunctionBlock`，提供指向 After Effects 将为不同文件相关任务调用的函数的指针。

下表显示了哪些函数是需要输入的，哪些是需要输出的。对于一个简单的实现，从右栏中标明 "需要 "的函数开始。你通常可以通过让 After Effects 为你处理调用(通过返回`AEIO_Err_USE_DFLT_CALLBACK`)来调用 "最佳情况 "行为。

对于一个只用于视频输入的原始 AEIO，实现以下函数。`AEIO_InitInSpecFromFile`或`AEIO_InitInSpecInteractive`(取决于源是文件还是交互生成的)，`AEIO_DisposeInSpec`。`AEIO_GetInSpecInfo`, `AEIO_DrawSparseFrame`, `AEIO_CloseSourceFiles`, 和`AEIO_InqNextFrameTime` (使用`AEIO_Err_USE_DFLT_CALLBACK`也可以)。

从 IO 样本开始，最好也留下其他函数的定义，并根据需要进一步填充。

## AEIO_FunctionBlock4

### AEIO_InitInSpecFromFile

给定一个文件路径，在提供的`AEIO_InSpecH`中向 After Effects 描述其内容。
使用[AEGP_IOInSuite]中所有适当的 "设置 "调用来完成。
如果有图像数据，设置其深度、尺寸和阿尔法解释。
如果有音频，描述其通道和采样率。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。

```cpp
AEIO_InitInSpecFromFile(
 AEIO_BasicData *basic_dataP,
 const A_UTF16Char *file_pathZ,
 AEIO_InSpecH inH);
```

IO：输入

需要。是的，对于基于文件的媒体

### AEIO_InitInSpecInteractive

使用某种形式的用户互动(而不是由 After Effects 提供的文件路径)。
描述你生成的 AEIO_InSpecH 包含的音频和视频。

```cpp
AEIO_InitInSpecInteractive(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH);
```

IO: 输入

需要。是的，对于互动生成的媒体

### AEIO_DisposeInSpec

释放一个`AEIO_InSpecH`。

```cpp
AEIO_DisposeInSpec(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH);
```

IO: 输入

需要。是

### AEIO_FlattenOptions

对于给定的`AEIO_InSpecH`，返回其选项句柄中包含的数据的扁平化版本。
使用`AEGP_GetInSpecOptionsHandle'获取未扁平化的选项句柄。

```cpp
AEIO_FlattenOptions(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 AEIO_Handle *flat_optionsPH);
```

IO: 输入

需要。不需要

### AEIO_InflateOptions

对于给定的`AEIO_InSpecH`，创建(使用`AEGP_SetInSpecOptionsHandle`)它的非扁平化版本。
创建其扁平化的选项数据的非扁平化版本。

```cpp
AEIO_InflateOptions(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 AEIO_Handle flat_optionsH);
```

IO: 输入

需要。不需要

### AEIO_SynchInSpec

允许`AEIO_Err_USE_DFLT_CALLBACK`。
检查`AEIO_InSpecH`，如果需要的话，更新它的选项)，并指出你是否做了改变。

```cpp
AEIO_SynchInSpec(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_Boolean *changed0);
```

IO: 输入

需要。不需要

### AEIO_GetActiveExtent

`AEIO_Err_USE_DFLT_CALLBACK'允许。 在提供的`A_LRect`中填入给定时间内文件像素的活动范围。

```cpp
AEIO_GetActiveExtent(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 const A_Time *tr,
 A_LRect *extent);
```

IO: 输入

需要。是

### AEIO_GetInSpecInfo

在`AEIO_Verbiage`中提供一些字符串来描述文件，这些字符串将出现在项目面板中。
这包括用于描述文件类型和子类型(编解码器)的字符串。

```cpp
AEIO_GetInSpecInfo(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 AEIO_Verbiage *verbiageP);
```

这个函数经常被调用；每次我们刷新项目面板的时候。
保持分配到最低限度。
在 After Effects 的 AEIOs 中，我们检查一个有效的 "optionsH"(使用 "AEGP*GetInSpecOptionsHandle")。
如果我们找到一个，我们就使用其中的信息。如果没有，我们什么都不做。
这很重要；如果你的 AEIO 处理静止图像，这个函数*会被调用到包含静止图像的文件夹。
希望不会有一个与之相关的 optionsH(除非你在写一个真正奇怪的 AEIO)。

IO：输入

需要。是

### AEIO_DrawSparseFrame

从`AEIO_InSpecH`中绘制一个帧。
`PF_EffectWorld*`包含要使用的宽度和高度，但要确保你考虑到 required_region0，如果它不是 NULL。

```cpp
AEIO_DrawSparseFrame(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 AEIO_Quality qual,
 const AEIO_RationalScale *rs0,
 const A_Time *tr,
 const A_Time *duration0,
 const A_Rect *required_region0,
 PF_EffectWorld *wP,
 A_long* originx,
 A_long* originy,
 AEIO_DrawingFlags *draw_flagsP);
```

注意：返回数据为线性光(1.0)，After Effects
将执行任何必要的转换，将镜头带入工作的色彩空间。

IO：输入

需要。是

### AEIO_GetDimensions

AEIO_Err_USE_DFLT_CALLBACK 允许。在 AEIO_InSpecH 中提供视频的尺寸(如有必要，还有缩放系数)。

```cpp
AEIO_GetDimensions(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 const AEIO_RationalScale *rs0,
 A_long *width0,
 A_long *height0);
```

输入：输入

需要。不需要

### AEIO_GetDuration

允许 "AEIO_Err_USE_DFLT_CALLBACK"。提供 "AEIO_InSpecH "的持续时间，单位是秒。

```cpp
AEIO_GetDuration(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_Time *trP);
```

输入：输入

需要。不需要

### AEIO_GetTime

允许 "AEIO_Err_USE_DFLT_CALLBACK"。提供 "AEIO_InSpecH "的时间基数。

```cpp
AEIO_GetTime(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_Time *tr);
```

以下是我们内部使用的常用时间基的值。
29.97 fps: scale = 100; value= 2997;
59.94 fps: scale = 50; value = 2997;
23.976 fps: scale = 125; value = 2997;
30 帧：比例=1；值=30。
25 fps: scale = 1; value = 25;

输入：输入

需要。没有

### AEIO_GetSound

允许使用`AEIO_Err_USE_DFLT_CALLBACK`。提供来自`AEIO_InSpecH'的声音，以描述的质量。

```cpp
AEIO_GetSound(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 AEIO_SndQuality quality,
 const AEIO_InterruptFuncs *interrupt_funcsP0,
 const A_Time *startPT,
 const A_Time *durPT,
 A_u_long start_sampLu,
 A_u_long num_samplesLu,
 void *dataPV);
```

`AEIO_SndQuality`可以是。

- `AEIO_SndQuality_APPROX`, (this quality is used to draw the audio waveform)
- `AEIO_SndQuality_LO`,
- `AEIO_SndQuality_HI`

IO: 输入

需要。不需要

### AEIO_InqNextFrameTime

`AEIO_Err_USE_DFLT_CALLBACK`允许。
在`AEIO_InSpecH`内提供下一帧的时间(在源录像的时间基础上)。

```cpp
AEIO_InqNextFrameTime(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 const A_Time *base_time_tr,
 AEIO_TimeDir time_dir,
 A_Boolean *found0,
 A_Time *key_time_tr0);
```

IO: 输入

需要。是

### AEIO_InitOutputSpec

允许 "AEIO_Err_USE_DFLT_CALLBACK"。
对新的`AEIO_OutSpecH`执行任何必要的初始化，并指出你是否做了改变。

```cpp
AEIO_InitOutputSpec(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_Boolean *user_interacted);
```

注意：第一次使用你的 AEIO 时，After Effects 会在它的偏好中缓存最后一个已知好的选项 H。
在测试这个功能时，请经常[删除你的偏好](../intro/debugging-plug-ins.html#intro-debugging-plug-ins-deleting-preferences)。

IO。输出

需要。是

### AEIO_GetFlatOutputOptions

描述(在一个`AEIO_Handle`中)一个`AEIO_OutSpecH`的输出选项。
在一个磁盘安全的平面数据结构中(不引用外部内存)。
注意，你的输出选项必须是跨平台的，所以要注意字节排序问题。

```cpp
AEIO_GetFlatOutputOptions(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 AEIO_Handle *optionsH);
```

IO。输出

需要。是

### AEIO_DisposeOutputOptions

`AEIO_Err_USE_DFLT_CALLBACK`允许。为传入的输出选项释放内存。

```cpp
AEIO_DisposeOutputOptions(
 AEIO_BasicData *basic_dataP,
 void *optionsPV);
```

IO。输出

需要。不需要

### AEIO_UserOptionsDialog

显示一个输出设置对话框(在 After Effects 中选择 TIFF 输出以查看何时会出现这个对话框)。
使用`AEGP_SetInSpecOptionsHandle`将此信息存储在一个选项句柄中。

```cpp
AEIO_UserOptionsDialog(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 PF_EffectWorld *sample0,
 A_Boolean *interacted0);
```

IO。输出

需要。不需要

### AEIO_GetOutputInfo

在`AEIO_OutSpecH`中描述(文本)输出选项。

```cpp
AEIO_GetOutputInfo(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 AEIO_Verbiage *verbiage);
```

### AEIO_OutputInfoChanged

根据当前的设置更新`AEIO_OutSpecH'(使用各种 Get 函数来获取)。

```cpp
AEIO_OutputInfoChanged(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH);
```

IO: 输出

需要。不需要

### AEIO_SetOutputFile

`AEIO_Err_USE_DFLT_CALLBACK`允许。设置`AEIO_OutSpecH'的输出文件路径。 返回`AEIO_Err_USE_DEFAULT_CALLBACK`，除非你改变了路径。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。

```cpp
AEIO_SetOutputFile(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_UTF16Char *file_pathZ);
```

IO。输出

需要。是

### AEIO_StartAdding

准备向输出文件添加帧。
这是一个在磁盘上创建输出文件的好时机，并将任何标题信息写入这些文件中。
这也是你第一次有机会根据有效的输出规格值来分配像素缓冲区。

```cpp
AEIO_StartAdding(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long flags);
```

IO。输出

需要。是的，用于写入支持多帧的格式

### AEIO_AddFrame

将帧添加到输出文件。你可以传递一个指针到一个你希望在用户中断渲染时调用的函数。

```cpp
AEIO_AddFrame(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long frame_index,
 A_long frames,
 PF_EffectWorld *wP,
 const A_LPoint *origin0,
 A_Boolean was_compressedB,
 AEIO_InterruptFuncs *inter0);
```

IO。输出

需要。是的，用于编写支持多帧的格式

### AEIO_EndAdding

执行与添加帧相关的任何清理工作。

```cpp
AEIO_EndAdding(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long flags);
```

IO.输出 输出

需要。是的，对于支持多帧的写入格式

### AEIO_OutputFrame

输出一个单帧。

```cpp
AEIO_OutputFrame(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 PF_EffectWorld *wP);
```

IO: 输出

需要。是的，用于写入支持单帧的格式

### AEIO_WriteLabels

`AEIO_Err_USE_DFLT_CALLBACK`允许。为`AEIO_OutSpecH`设置字母解释和字段使用信息。
在`AEIO_LabelFlags`中指出你写了哪些标志。

```cpp
AEIO_WriteLabels(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 AEIO_LabelFlags *written);
```

IO。输出

需要。是

### AEIO_GetSizes

`AEIO_Err_USE_DFLT_CALLBACK`允许。提供关于文件大小和输出卷上剩余自由空间的信息。

```cpp
AEIO_GetSizes(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_u_longlong *free_space,
 A_u_longlong *file_size);
```

IO。输出

需要。是

### AEIO_Flush

销毁与`OutSpecH`相关的任何选项或用户数据。

```cpp
AEIO_Flush(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH);
```

### AEIO_AddSoundChunk

将给定的声音添加到输出文件中。

```cpp
AEIO_AddSoundChunk(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 const A_Time *start,
 AEIO_SndWorldH swH);
```

IO。输出

需要。是的，用于写入带有音频的格式

### AEIO_Idle

可选的。对空闲时间做一些处理。不支持`AEIO_Err_USE_DFLT_CALLBACK`。

```cpp
AEIO_Idle(
 AEIO_BasicData *basic_dataP,
 AEIO_ModuleSignature sig,
 AEIO_IdleFlags *idle_flags0);
```

IO: 输出

需要。不需要

### AEIO_GetDepths

设置`AEIO_OptionsFlags'以指示哪些像素和颜色深度对你的输出格式有效。
见[Export Bit-Depth]的讨论(implementation-details.html#aeios-implementation-details)。

```cpp
AEIO_GetDepths(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 AEIO_OptionsFlags *which);
```

IO。输出

需要。是

### AEIO_GetOutputSuffix

允许使用`AEIO_Err_USE_DFLT_CALLBACK`。描述输出文件的三个字符扩展名。

```cpp
AEIO_GetOutputSuffix(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_char *suffix);
```

IO。输出

需要。是

### AEIO_SeqOptionsDlg

显示一个脚注选项对话框，并指出用户是否做了任何改变。

```cpp
AEIO_SeqOptionsDlg(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_Boolean *interactedPB);
```

IO: 输入

需要。不需要

### AEIO_GetNumAuxChannels

枚举一个`AEIO_InSpecH'中包含的数据的辅助通道(超出红、绿、蓝和 α)。

```cpp
AEIO_GetNumAuxChannels(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_long *num_channelsPL);
```

输入：输入

需要。不需要

### AEIO_GetAuxChannelDesc

描述一个辅助数据通道的数据类型、名称、通道和尺寸。

```cpp
AEIO_GetAuxChannelDesc(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 long chan_indexL,
 PF_ChannelDesc *descP);
```

输入：输入

需要。不需要

### AEIO_DrawAuxChannel

从一个`AEIO_InSpecH'中提取辅助通道。

```cpp
AEIO_DrawAuxChannel(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_long chan_indexL,
 const AEIO_DrawFramePB *pbP,
 PF_ChannelChunk *chunkP);
```

### AEIO_FreeAuxChannel

释放与一个辅助通道相关的数据。

```cpp
AEIO_FreeAuxChannel(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 PF_ChannelChunk *chunkP);
```

输入：输入

需要。不需要

`AEIO_Num` AuxFiles 枚举渲染给定的`AEIO_InSpecH`需要的文件。
这个函数和`AEIO_GetNthAuxFileSpec`将在用户选择`文件>依赖性>收集文件...`时被调用。

这里你的 AEIO 告诉 AE 相关的文件是什么。

```cpp
AEIO_NumAuxFiles(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH seqH,
 A_long *files_per_framePL);
```

IO：输入

需要。不需要

### AEIO_GetNthAuxFileSpec

从第 n 个辅助文件检索数据，为指定的框架。
路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处置。

```cpp
AEIO_GetNthAuxFileSpec(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH seqH,
 A_long frame_numL,
 A_long n,
 AEGP_MemHandle *pathPH);
```

IO: 输入

需要。不需要，如果没有辅助文件

### AEIO_CloseSourceFiles

关闭(或打开，取决于 closeB)一个`AEIO_InSpecH`的源文件。
当用户收集文件时，AEIO 将首先被要求关闭它的源文件，然后重新打开它们。

```cpp
AEIO_CloseSourceFiles(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH seqH,
 A_Boolean closeB);
```

`TRUE`表示关闭，`FALSE`表示打开。

IO: 输入

需要。是

### AEIO_CountUserData

列举与`AEIO_InSpecH'相关的用户数据单位。

```cpp
AEIO_CountUserData(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_u_long typeLu,
 A_u_long max_sizeLu,
 A_u_long *num_of_typePLu);
```

### AEIO_SetUserData

为给定的`AEIO_OutSpecH`设置用户数据(给定的索引和类型)。

```cpp
AEIO_SetUserData(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_u_long typeLu,
 A_u_long indexLu,
 const AEIO_Handle dataH);
```

IO。输出

需要。不需要

### AEIO_GetUserData

描述与`AEIO_InSpecH'相关的用户数据(在给定的索引和类型)。

```cpp
AEIO_GetUserData(
 AEIO_BasicData *basic_dataP,
 AEIO_InSpecH inH,
 A_u_long typeLu,
 A_u_long indexLu,
 A_u_long max_sizeLu,
 AEIO_Handle *dataPH);
```

输入：输入

需要。不需要

### AEIO_AddMarker

在指定的帧上，将指定类型的标记与`AEIO_OutSpecH'相关联。
你可以提供一个中断函数来处理用户对这个动作的取消。

```cpp
AEIO_AddMarker(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long frame_index,
 AEIO_MarkerType marker_type,
 void *marker_dataPV,
 AEIO_InterruptFuncs *inter0);
```

IO。输出

需要。不需要

### AEIO_VerifyFileImportable

指示(通过设置 importablePB)插件是否可以导入该文件。
注意，After Effects 已经做了基本的扩展名检查；你可能希望打开文件并确定它是否有效。
这可能是一个耗时的过程；大多数 After Effects 的 AEIO 只是简单地返回 TRUE。
并在`AEIO_InitInSpecFromFile`中处理坏文件。
文件路径是一个以 NULL 结尾的 UTF-16 字符串，带有平台分隔符。

```cpp
AEIO_VerifyFileImportable(
 AEIO_BasicData *basic_dataP,
 AEIO_ModuleSignature sig,
 const A_UTF16Char *file_pathZ,
 A_Boolean *importablePB);
```

IO：输入

需要。不需要

### AEIO_UserAudioOptionsDialog

显示一个音频选项对话框。

```cpp
AEIO_UserAudioOptionsDialog(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_Boolean *interacted0);
```

IO: 输出

需要。不需要

### AEIO_AddMarker3

添加一个标记，用一个标志来指定这是否是一个合成标记。

```cpp
AEIO_AddMarker3(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long frame_index,
 AEGP_ConstMarkerValP marker_valP,
 AEIO_RenderMarkerFlag marker_flag,
 AEIO_InterruptFuncs *inter0);
```

IO。输出

需要。不需要

### AEIO_GetMimeType

描述输出的 mime 类型。这是用于 XMP 支持。

```cpp
AEIO_GetMimeType(
 AEIO_BasicData *basic_dataP,
 AEIO_OutSpecH outH,
 A_long mime_type_sizeL,
 char *mime_typeZ);
```

IO。输出

需要。不需要

## What Goes In

这些函数管理一个输入规范，After Effects 对从任何来源收集的数据的内部表示。

在 After Effects 中的任何图像或音频数据(固态层除外)都是从输入规格句柄或`AEIO_InSpecH`获得的。

### AEGP_IOInSuite5

### AEGP_GetInSpecOptionsHandle

为给定的`AEIO_InSpecH`检索选项数据(由你的 AEIO 创建)。

```cpp
AEGP_GetInSpecOptionsHandle(
 AEIO_InSpecH inH,
 void **optionsPPV);
```

### AEGP_SetInSpecOptionsHandle

为给定的`AEIO_InSpecH'设置选项数据。
必须使用[Memory Suite](./aegps/aegp-suites.html#aegps-aegp-suites-memory-suite)分配。

```cpp
AEGP_SetInSpecOptionsHandle(
 AEIO_InSpecH inH,
 void *optionsPV,
 void **old_optionsPPV);
```

### AEGP_GetInSpecFilePath

检索`AEIO_InSpecH`的文件路径。
文件路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处理。

```cpp
AEGP_GetInSpecFilePath(
 AEIO_InSpecH inH,
 AEGP_MemHandle *file_nameZ);
```

### AEGP_GetInSpecNativeFPS

检索`AEIO_InSpecH`的帧速率。

```cpp
AEGP_GetInSpecNativeFPS(
 AEIO_InSpecH inH,
 A_Fixed *native_fpsP);
```

### AEGP_SetInSpecNativeFPS

设置 "AEIO_InSpecH "的帧速率。

```cpp
AEGP_SetInSpecNativeFPS(
 AEIO_InSpecH inH,
 A_Fixed native_fpsP);
```

### AEGP_GetInSpecDepth

检索 "AEIO_InSpecH "中图像数据的比特深度。

```cpp
AEGP_GetInSpecDepth(
 AEIO_InSpecH inH,
 A_short *depthPS);
```

### AEGP_SetInSpecDepth

向 After Effects 显示`AEIO_InSpecH'中图像数据的比特深度。

```cpp
AEGP_SetInSpecDepth(
 AEIO_InSpecH inH,
 A_short depthS);
```

### AEGP_GetInSpecSize

检索 "AEIO_InSpecH "所引用的数据的大小(以字节为单位)。

```cpp
AEGP_GetInSpecSize(
 AEIO_InSpecH inH,
 AEIO_FileSize *sizePLLu);
```

### AEGP_SetInSpecSize

向 After Effects 表明`AEIO_InSpecH`所引用的数据的大小(以字节为单位)。

```cpp
AEGP_SetInSpecSize(
 AEIO_InSpecH inH,
 AEIO_FileSize sizeL);
```

### AEGP_GetInSpecInterlaceLabel

检索 "AEIO_InSpecH "的字段信息。

```cpp
AEGP_GetInSpecInterlaceLabel(
 AEIO_InSpecH inH,
 FIEL_Label *interlaceP);
```

### AEGP_SetInSpecInterlaceLabel

指定 "AEIO_InSpecH "的字段信息。

```cpp
AEGP_SetInSpecInterlaceLabel(
 AEIO_InSpecH inH,
 const FIEL_Label *interlaceP);
```

### AEGP_GetInSpecAlphaLabel

检索 "AEIO_InSpecH "的 alpha 通道解释信息。

```cpp
AEGP_GetInSpecAlphaLabel(
 AEIO_InSpecH inH,
 AEIO_AlphaLabel *alphaP);
```

### AEGP_SetInSpecAlphaLabel

设置 "AEIO_InSpecH "的字母通道解释信息。

```cpp
AEGP_SetInSpecAlphaLabel(
 AEIO_InSpecH inH,
 const AEIO_AlphaLabel* alphaP);
```

### AEGP_GetInSpecDuration

检索 "AEIO_InSpecH "的持续时间。

```cpp
AEGP_GetInSpecDuration(
 AEIO_InSpecH inH,
 A_Time *durationP);
```

### AEGP_SetInSpecDuration

设置 "AEIO_InSpecH "的持续时间。
注意：从 5.5 开始，必须调用这个功能，即使是基于帧的文件格式。
如果你不把`A_Time.scale`设置为零以外的东西，你的文件将无法导入。
这个问题将在未来的版本中得到解决。

```cpp
AEGP_SetInSpecDuration(
 AEIO_InSpecH inH,
 const A_Time *durationP);
```

### AEGP_GetInSpecDimensions

检索`AEIO_InSpecH`中图像数据的宽度和高度。

```cpp
AEGP_GetInSpecDimensions(
 AEIO_InSpecH inH,
 A_long *widthPL0,
 A_long *heightPL0);
```

### AEGP_SetInSpecDimensions

向 After Effects 显示 "AEIO_InSpecH "中的图像数据的宽度和高度。

```cpp
AEGP_SetInSpecDimensions(
 AEIO_InSpecH inH,
 A_long widthL,
 A_long heightL);
```

### AEGP_InSpecGetRational

尺寸

检索应用于`AEIO_InSpecH'的宽度、高度、边界矩形和缩放系数。

```cpp
AEGP_InSpecGetRationalDimensions(
 AEIO_InSpecH inH,
 const AEIO_RationalScale *rs0,
 A_long *width0,
 A_long *height0,
 A_Rect *r0);
```

### AEGP_GetInSpecHSF

检索应用于 "AEIO_InSpecH "的水平缩放系数。

```cpp
AEGP_GetInSpecHSF(
 AEIO_InSpecH inH,
 A_Ratio *hsf);
```

### AEGP_SetInSpecHSF

设定 "AEIO_InSpecH "的水平缩放系数。

```cpp
AEGP_SetInSpecHSF(
 AEIO_InSpecH inH,
 const A_Ratio *hsf);
```

### AEGP_GetInSpecSoundRate

获取 "AEIO_InSpecH "所引用的音频数据的采样率(以每秒采样数为单位)。

```cpp
AEGP_GetInSpecSoundRate(
 AEIO_InSpecH inH,
 A_FpLong *ratePF);
```

### AEGP_SetInSpecSoundRate

为 "AEIO_InSpecH "引用的音频数据设置采样率(以每秒采样数计)。

```cpp
AEGP_SetInSpecSoundRate(
 AEIO_InSpecH inH,
 A_FpLong rateF);
```

### AEGP_GetInSpecSoundEncoding

从 AEIO_InSpecH 中获取编码方法(有符号 PCM，无符号 PCM，或浮点)。

```cpp
AEGP_GetInSpecSoundEncoding(
 AEIO_InSpecH inH,
 AEIO_SndEncoding *encodingP);
```

### AEGP_SetInSpecSoundEncoding

设置 AEIO_InSpecH 的编码方法。

```cpp
AEGP_SetInSpecSoundEncoding(
 AEIO_InSpecH inH,
 AEIO_SndEncoding encoding);
```

### AEGP_GetInSpecSoundSampleSize

从`AEIO_InSpecH`中获取每样本的字节数(1,2,或 4)。

```cpp
AEGP_GetInSpecSoundSampleSize(
 AEIO_InSpecH inH,
 AEIO_SndSampleSize *bytes_per_smpP);
```

### AEGP_SetInSpecSoundSampleSize

设置 "AEIO_InSpecH "的每采样字节数。

```cpp
AEGP_SetInSpecSoundSampleSize(
 AEIO_InSpecH inH,
 AEIO_SndSampleSize bytes_per_sample);
```

### AEGP_GetInSpecSoundChannels

决定`AEIO_SndChannels`中的音频是单声道还是立体声。

```cpp
AEGP_GetInSpecSoundChannels(
 AEIO_InSpecH inH,
 AEIO_SndChannels *num_channelsP);
```

### AEGP_SetInSpecSoundChannels

将一个`AEIO_SndChannels`中的音频设置为单声道或立体声。

```cpp
AEGP_SetInSpecSoundChannels(
 AEIO_InSpecH inH,
 AEIO_SndChannels num_channels);
```

### AEGP_AddAuxExtMap

如果你的文件格式有辅助文件，你想阻止用户直接打开。
把它的扩展名、文件类型和创建者传给这个函数，以防止它出现在输入对话框中。

```cpp
AEGP_AddAuxExtMap(
 const A_char *extension,
 A_long file_type,
 A_long creator);
```

### AEGP_SetInSpecEmbeddedColorProfile

如果是 RGB 数据，如果有一个嵌入的 icc 配置文件，使用`AEGP_ColorProfile`从这个 icc 配置文件中建立一个`AEGP_ColorProfile`。
使用[AEGP_ColorSettingsSuite2](../aegps/aegp-suites.html#aegps-aegp-suites-aegp-colorsettingssuite)的`AEGP_GetNewColorProfileFromICCProfile'从这个icc配置文件中建立一个`AEGP_ColorProfile'，并将配置文件描述设置为 null。
在非 RGB 数据的情况下，如果有一个嵌入式的非 RGB icc 配置文件，或者你知道数据所处的色彩空间。
将颜色配置文件设置为 NULL，并提供以 NULL 结尾的 unicode 字符串的描述。
这样做会使色彩管理的用户界面失效，从而使用户能够在应用程序的用户界面上影响配置文件的选择。
如果你将非 RGB 数据直接解压到工作空间(要获得工作空间请使用`AEGP_GetNewWorkingSpaceColorProfile`)，你就完成了。
如果你要将非 RGB 数据解压到特定的 RGB 色彩空间，你必须将描述这个空间的配置文件传递给下面的`AEGP_SetInSpecAssignedColorProfile`。
否则，你的 RGB 数据将被错误地解释为在工作空间。
在这个函数中，颜色配置文件或配置文件描述应该是 NULL。你不能同时使用。

```cpp
AEGP_SetInSpecEmbeddedColorProfile(
 AEIO_InSpecH inH,
 AEGP_ConstColorProfileP color_profileP0,
 const A_UTF16Char *profile_descP0);
```

### AEGP_SetInSpecAssignedColorProfile

为镜头分配一个有效的 RGB 颜色配置文件。

```cpp
AEGP_SetInSpecAssignedColorProfile(
 AEIO_InSpecH inH,
 AEGP_ConstColorProfileP color_profileP);
```

### AEGP_GetInSpecNativeStartTime

CC 中的新功能。检索素材的原始开始时间。

```cpp
AEGP_GetInSpecNativeStartTime(
 AEIO_InSpecH inH,
 A_Time *startTimeP);
```

### AEGP_SetInSpecNativeStartTime

在 CC 中新加入。为素材指定一个本地开始时间。

```cpp
AEGP_SetInSpecNativeStartTime(
 AEIO_InSpecH inH,
 const A_Time *startTimeP);
```

### AEGP_ClearInSpecNativeStartTime

在 CC 中新增 清除镜头的原始开始时间。
使用`AEGP_SetInSpecNativeStartTime`将原生开始时间设置为 0 并不能做到这点。
它仍然意味着有一个特殊的本地开始时间。

```cpp
AEGP_ClearInSpecNativeStartTime(
 AEIO_InSpecH inH);
```

### AEGP_GetInSpecNativeDisplayDropFrame

CC 中的新功能。检索镜头的下降帧设置。

```cpp
AEGP_GetInSpecNativeDisplayDropFrame(
 AEIO_InSpecH inH,
 A_Boolean *displayDropFrameBP);
```

### AEGP_SetInSpecNativeDisplayDropFrame

CC 中的新功能。指派镜头的停格设置。

```cpp
AEGP_SetInSpecNativeDisplayDropFrame(
 AEIO_InSpecH inH,
 A_Boolean displayDropFrameB);
```

## What Goes Out

这些功能管理与 After Effects 渲染队列中的输出规范的所有互动。

### AEGPIOOutSuite4

### AEGP_GetOutSpecOptionsHandle

检索 "AEIO_OutSpecH "的选项。

```cpp
AEGP_GetOutSpecOptionsHandle(
 AEIO_OutSpecH outH,
 void **optionsPPV);
```

### AEGP_SetOutSpecOptionsHandle

设置 "AEIO_OutSpecH "的选项。

```cpp
AEGP_SetOutSpecOptionsHandle(
 AEIO_OutSpecH outH,
 void *optionsPV,
 void **old_optionsPPV);
```

### AEGP_GetOutSpecFilePath

获取 "AEIO_OutSpecH "的路径。文件路径是一个以 NULL 结尾的 A_UTF16Char 字符串的句柄，必须用`AEGP_FreeMemHandle`处理。
如果`file_rsrvdPB`返回`TRUE`，插件不应该覆盖它。
(After Effects 已经创建了一个空文件)；这样做会导致网络渲染失败。

```cpp
AEGP_GetOutSpecFilePath(
 AEIO_OutSpecH outH,
 AEGP_MemHandle *unicode_pathPH,
 A_Boolean *file_rsrvdPB);
```

### AEGP_GetOutSpecFPS

获取`AEIO_OutSpecH`的每秒帧数。

```cpp
AEGP_GetOutSpecFPS(
 AEIO_OutSpecH outH,
 A_Fixed *native_fpsP);
```

### AEGP_SetOutSpecNativeFPS

设置 "AEIO_OutSpecH "的每秒帧数。

```cpp
AEGP_SetOutSpecNativeFPS(
 AEIO_OutSpecH outH,
 A_Fixed native_fpsP);
```

### AEGP_GetOutSpecDepth

获取 "AEIO_OutSpecH "的像素位深度。

```cpp
AEGP_GetOutSpecDepth(
 AEIO_OutSpecH outH,
 A_short *depthPS);
```

### AEGP_SetOutSpecDepth

设置 "AEIO_OutSpecH "的像素位深度。

```cpp
AEGP_SetOutSpecDepth(
 AEIO_OutSpecH outH,
 A_short depthPS);
```

### AEGP_GetOutSpecInterlaceLabel

获取 "AEIO_OutSpecH "的字段信息。

```cpp
AEGP_GetOutSpecInterlaceLabel(
 AEIO_OutSpecH outH,
 FIEL_Label *interlaceP);
```

### AEGP_SetOutSpecInterlaceLabel

设置 "AEIO_OutSpecH "的字段信息。

```cpp
AEGP_SetOutSpecInterlaceLabel(
 AEIO_OutSpecH outH,
 const FIEL_Label *interlaceP);
```

### AEGP_GetOutSpecAlphaLabel

获取 "AEIO_OutSpecH "的字母解释信息。

```cpp
AEGP_GetOutSpecAlphaLabel(
 AEIO_OutSpecH outH,
 AEIO_AlphaLabel *alphaP);
```

### AEGP_SetOutSpecAlphaLabel

为 "AEIO_OutSpecH "设置阿尔法解释。

```cpp
AEGP_SetOutSpecAlphaLabel(
 AEIO_OutSpecH outH,
 const AEIO_AlphaLabel *alphaP);
```

### AEGP_GetOutSpecDuration

获取 "AEIO_OutSpecH "的持续时间。

```cpp
AEGP_GetOutSpecDuration(
 AEIO_OutSpecH outH,
 A_Time *durationP);
```

### AEGP_SetOutSpecDuration

设置 "AEIO_OutSpecH "的持续时间。

```cpp
AEGP_SetOutSpecDuration(
 AEIO_OutSpecH outH,
 const A_Time *durationP);
```

### AEGP_GetOutSpecDimensions

获得 "AEIO_OutSpecH "的尺寸。

```cpp
AEGP_GetOutSpecDimensions(
 AEIO_OutSpecH outH,
 A_long *widthPL0,
 A_long *heightPL0);
```

### AEGP_GetOutSpecHSF

获取 "AEIO_OutSpecH "的水平缩放系数。

```cpp
AEGP_GetOutSpecHSF(
 AEIO_OutSpecH outH,
 A_Ratio *hsf);
```

### AEGP_SetOutSpecHSF

设定 "AEIO_OutSpecH "的水平缩放系数。

```cpp
AEGP_SetOutSpecHSF(
 AEIO_OutSpecH outH,
 const A_Ratio *hsf);
```

### AEGP_GetOutSpecSoundRate

获取 "AEIO_OutSpecH "的采样率。

```cpp
AEGP_GetOutSpecSoundRate(
 AEIO_OutSpecH outH,
 A_FpLong *ratePF);
```

### AEGP_SetOutSpecSoundRate

设置 "AEIO_OutSpecH "的采样率。

```cpp
AEGP_SetOutSpecSoundRate(
 AEIO_OutSpecH outH,
 A_FpLong rateF);
```

### AEGP_GetOutSpecSoundEncoding

获取 "AEIO_OutSpecH "的声音编码格式。

```cpp
AEGP_GetOutSpecSoundEncoding(
 AEIO_OutSpecH outH,
 AEIO_SndEncoding *encodingP);
```

### AEGP_SetOutSpecSoundEncoding

设置 "AEIO_OutSpecH "的声音编码格式。

```cpp
AEGP_SetOutSpecSoundEncoding(
 AEIO_OutSpecH outH,
 AEIO_SndEncoding encoding);
```

### AEGP_GetOutSpecSoundSampleSize

获取 "AEIO_OutSpecH "的每采样字节数。

```cpp
AEGP_GetOutSpecSoundSampleSize(
 AEIO_OutSpecH outH,
 AEIO_SndSampleSize *bpsP);
```

### AEGP_SetOutSpecSoundSampleSize

设置 "AEIO_OutSpecH "的每采样字节数。

```cpp
AEGP_SetOutSpecSoundSampleSize(
 AEIO_OutSpecH outH,
 AEIO_SndSampleSize bpsP);
```

### AEGP_GetOutSpecSoundChannels

获取 "AEIO_OutSpecH "中的声音通道数。

```cpp
AEGP_GetOutSpecSoundChannels(
 AEIO_OutSpecH outH,
 AEIO_SndChannels *channelsP);
```

### AEGP_SetOutSpecSoundChannels

设置 "AEIO_OutSpecH "中的声音通道数。

```cpp
AEGP_SetOutSpecSoundChannels(
 AEIO_OutSpecH outH,
 AEIO_SndChannels channels);
```

### AEGP_GetOutSpecIsStill

确定 "AEIO_OutSpecH "是否是静止的。

```cpp
AEGP_GetOutSpecIsStill(
 AEIO_OutSpecH outH,
 A_Boolean *is_stillPB);
```

### AEGP_GetOutSpecPosterTime

获取 "AEIO_OutSpecH "的海报帧的时间。

```cpp
AEGP_GetOutSpecPosterTime(
 AEIO_OutSpecH outH,
 A_Time *poster_timeP);
```

### AEGP_GetOutSpecStartFrame

获取 "AEIO_OutSpecH "中第一帧的时间。

```cpp
AEGP_GetOutSpecStartFrame(
 AEIO_OutSpecH outH,
 A_long *start_frameP);
```

### AEGP_GetOutSpecPullDown

获取 "AEIO_OutSpecH "的下拉阶段。

```cpp
AEGP_GetOutSpecPullDown(
 AEIO_OutSpecH outH,
 AEIO_Pulldown *pulldownP);
```

### AEGP_GetOutSpecIsMissing

如果没有 "AEIO_OutSpecH"，则传回 TRUE。

```cpp
AEGP_GetOutSpecIsMissing(
 AEIO_OutSpecH outH,
 A_Boolean *missingPB);
```

### AEGP_GetOutSpecShouldEmbedICCProfile

如果 AEIO 应该在输出中嵌入一个颜色配置文件，则返回 TRUE。

```cpp
AEGP_GetOutSpecShouldEmbedICCProfile(
 AEIO_OutSpecH outH,
 A_Boolean *embedPB);
```

### AEGP_GetNewOutSpecColorProfile

返回一个(不透明的)ICC 颜色配置文件，用于嵌入到 AEIO 的输出中。
必须用`AEGP_DisposeColorProfile`来处理。

```cpp
AEGP_GetNewOutSpecColorProfile(
 AEGP_PluginID aegp_plugin_id,
 AEIO_OutSpecH outH,
 AEGP_ColorProfileP *color_profilePP);
```

### AEGP_GetOutSpecOutputModule

返回与给定的`AEIO_OutSpecH`相关的`AEGP_RQItemRefH`和`AEGP_OutputModuleRefH`。
如果没有找到渲染队列项目，或者`AEIO_OutSpecH`不是一个确认的 outH，是一个副本，则失败。
也就是说，如果输出模块设置对话框打开，而用户没有点击 OK。

```cpp
AEGP_GetOutSpecOutputModule(
 AEIO_OutSpecH outH,
 AEGP_RQItemRefH *rq_itemP,
 AEGP_OutputModuleRefH *om_refP);
```
