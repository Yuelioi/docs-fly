---
title: What’s-New
order: 16
---

# 新鲜东西

如果这是你第一次开发 After Effects 插件，可以跳过新内容部分，直接进入[如何开始创建插件](how-to-start-creating-plug-ins.html) 。

## AE 2022 新变换

After Effects 2022 包含了第一个支持多帧渲染的完整公开版本。2021 年 10 月发布的相关 AE Effects SDK 包括一个增加 PF_Iterate 线程最大数量的改变。

## AE 2021 SDK 变换

### 多帧渲染

1. `PF_OutFlag2_SUPPORTS_THREADED_RENDERING`标志现已上线。设置这个标志就可以多帧渲染，但存储在`sequence_data`中的数据在渲染时强制要求是常量/只读，现在通过一个套件来访问`sequence_data`，`PF_EffectSequenceDataSuite1`。
2. 如果你的插件无法与新的 sequence_data 一起工作，现有一个新标志，`PF_OutFlag2_MUTABLE_RENDER_SEQUENCE_DATA_SLOWER`可以和`PF_OutFlag2_SUPPORTS_THREADED_RENDERING`一起设置。After Effects 将无法应用尽可能多的渲染并发性，因此无法对设置该标志的效果进行性能改进(因此，标志名称为 \_SLOWER)。
3. 一个新的套件--Compute Cache 计算缓存(以前被称为 3-way checkout cache)现已可用。该套件提供了一个线程安全的缓存，插件可以将其作为 sequence_data 的替代或补充，以支持多渲染线程计算和缓存渲染帧所需的数据。

由于这些变化，你必须更新到 2021 年 3 月的 SDK 并进行编译，以保持多帧渲染与 AE Beta 构建的兼容性。使用 2020 年 6 月的 SDK 编译的插件将不再支持多帧渲染，即使设置了`PF_OutFlag2_SUPPORTS_THREADED_RENDERING`，从 AE 22.0x6(2021 年 6 月 29 日发布)开始。

详情请参阅[Multi-Frame Rendering in AE](../effect-details/multi-frame-rendering-in-ae.html)。

### Apple Silicon 支持

- AE SDK 现在支持为 Apple Silicon 原生构建特效。虽然 After Effects 本身还不能在 Apple Silicon 上运行，但 Adobe 作为一家公司，正在推进对我们许多产品的原生支持。像 Premiere Pro 这样的应用程序现在有原生版本，你的效果可以通过 Motion Graphic Templates 等功能加载到 Premiere Pro 中。当运行 Premiere Pro 的原生版本时，只有原生编译的特效才能工作，所以尽快更新你的特效，支持 Apple Silicon 很重要。请参阅[Apple Silicon Support](apple-silicon-support.html)部分了解更多信息。

### 从效果导出符号

- SDK 样本已经更新，在 MacOS 上默认不导出符号。请参阅[Exporting Symbols in Effects](symbol-export.html) 了解更多信息。

### 下载 March 2021 SDK

SDK 可以从 Adobe Developer Console 下载，地址是<https://adobe.io/after-effects/>。

### After Effects Beta 构建

为了获得这个 SDK 的 AE 主机端变化，你需要从 Creative Cloud Desktop App 下载一个新的 After Effects beta 版本。2021 年 3 月的 SDK 支持 18.2x11 及以上版本的构建。

## June 2020 新特性

AE(目前仅在 Beta 版本中)现在支持多帧渲染。请参阅[Multi-Frame Rendering in AE](../effect-details/multi-frame-rendering-in-ae.html) 了解更多细节。

## CC 2019 (16.0) 新特性

我们对处理 GPU 特效的方式做了一些改变。详情请参见 "GPU 特效的变化"。

## 15.0 新特性

After Effects 现在支持以前在 Premiere Pro 中支持的*GPU 效果渲染*。请注意，匹配名称包含 "ADBE "的未知效果将被排除在 GPU 渲染之外，因此请确保您的任何 GPU 效果都有自己的自定义匹配名称。支持 GPU 渲染的效果将在效果面板中获得 GPU 标志。

Premiere Pro SDK 中的 GPU 效果示例项目已经更新，可以在 AE 中注册为 GPU 效果，不过渲染输出仍需努力。

一个新的入口函数已经被定义，以允许特效在运行时向主机注册基本信息，而不需要依赖传统的 PiPL 资源。一个特效可以通过这种方式在一个二进制文件中注册多个入口函数。Premiere Pro 是第一个支持这个入口函数的主机，而 After Effects 将在未来的版本中支持这个入口函数。

效果示例项目已被更新，以使用这种方法，同时保留 PiPL，以便向后兼容。

`AEGP_StreamSuite`现在是第 5 版，其中[AEGP_GetExpression()](../aegps/aegp-suites.html) 和[AEGP_SetExpression()](../aegps/aegp-suites.html) 已升级为支持统一码。

`PF_AdvTimeSuite`现在是第 4 版，有一个新的调用[PF_TimeCountFrames()](../effect-details/useful-utility-functions.html) ，它返回当前编译中的帧的索引。

新的 AEGP 数学套件提供了对矩阵乘法的有用调用。

应用程序的字体现在是 Adobe Clean。以前，After Effects 的用户界面中使用的字体在 Windows 上是 Tahoma，在 macOS X 上是 Lucida Grande。这是一种专有字体，我们不能让它在你的用户界面中使用。

## What’s New In CC 2017.1 (14.2)?

- 图层参数可以包括面罩和效果

使用图层作为输入的效果，如 Set Matte 和 Displacement Map，现在可以针对输入图层的遮罩和效果，而不是只针对图层的源。这意味着，没有必要为了让它们能够被效果引用而预先编排图层。

当一个效果包括一个图层参数时，图层选择器右边的一个新菜单允许你选择是否从输入图层的源头、遮罩或效果来瞄准它。

- 源：只针对图层的源。面具和效果被忽略。
- 面具：针对应用了面具的图层。效果被忽略。
- 效果和遮罩：在应用了遮罩和效果后，针对该层。

这个控制类似于图层查看器面板底部的查看菜单，它允许你从渲染顺序中的不同位置来渲染图层：从它的源，从它的蒙版，或从它的单个效果。

由于这是一个面向用户的选项，设计的目的是对效果透明。从效果的角度来看，输入只是包括上游的效果和遮罩，而没有对效果进行任何改变。对于任何使用层参数的效果，这里有一些测试建议。

- 效果继续按预期工作。
- 在 Source/Mask/Effects 的层参数中使用新的控制，可以与效果一起工作。
- 打开旧项目或保存回以前版本的项目不会破坏效果。
- 确认 effect 不能自我引用；也就是说不能把图层上的效果作为同一图层的输入。
- 套装的增强

PF_AdvTimeSuite 现在是第 3 版，提供了一个修订的[PF_GetTimeDisplayPref()](../effect-details/useful-utility-functions.html) 调用，使用修订的`PF_TimeDisplayPrefVersion`参数，支持更高帧率。
如果数值超过结构支持的范围有问题，以前的第二版调用现在可以返回错误。

Comp Suite 现在是第 11 版，有一个新的调用，[AEGP_ReorderCompSelection()](../aegps/aegp-suites.html) ，可以将一个选择移到某个层的索引。
它应该与`AEGP_SetSelection()`一起使用。

## What’s New In CC 2017 (14.1)?

对[AEGP 项目套件](../aegps/aegp-suites.html) 和[AEGP 渲染队列项目套件](../aegps/aegp-suites.html) 提供 Unicode 支持。

## What’s New In CC 2017 (14.0)?

GLator 样本回来了! 它已被更新，以展示效果插件中正确的 OpenGL 上下文管理。

## What’s New In CC 2015.3 (13.8)?

PF_OutFlag_I_AM_OBSOLETE is now supported in Premiere Pro. Also, effect custom UI in Premiere Pro now supports high DPI displays, such as Retina Displays.

## What’s New In CC 2015 (13.6)?

New AEGP Item View Suite. This provides a way to get playback time for item view. Only the composition case is implemented in this release. The time passed back should be the playback time of the view when playing and the current (needle) time otherwise.

AEGP_RenderNewItemSoundData() has been reworked and provides functionality similar to 13.2.

## What’s New In CC 2015 (13.5.1)?

This release fixes some audio APIs that broke in 13.5 due to threading changes. In 13.5, when called on the UI thread, AEGP_RenderNewItemSoundData() would return A_Err_GENERIC. This restores the functionality when called on the UI thread.

To avoid a deadlock, in PF_Cmd_UPDATE_PARAMS_UI only, AEGP_RenderNewItemSoundData() will now return silence. This will no longer function as before in this context, but it will continue to work properly elsewhere.

## What’s New In CC 2015 (13.5)?

- Separate UI and Render Threads

This release of After Effects includes major architectural changes to separate the UI (main) thread from the render thread. The render thread sends selectors such as PF_Cmd_RENDER, PF_Cmd_SMART_PRERENDER, and PF_Cmd_SMART_RENDER to effect plug-ins. The UI thread sends selectors such as PF_Cmd_SEQUENCE_SETUP, PF_Cmd_USER_CHANGED_PARAM, PF_Cmd_DO_DIALOG, and PF_EVENT_DRAW. PF_Cmd_SEQUENCE_RESETUP is

sent on both render and UI threads.

These changes are to improve interactive performance and responsiveness. At the same time, the new design introduces some new requirements and may break assumptions that existing plug-ins relied on. Here are some of the major changes:

1. The project can no longer be modified by the render thread (and in fact the render thread now has its own local copy of the project)
2. Rendering cannot pass modified sequence data back to the UI thread for custom UI updates
3. In general the UI thread should no longer do time-consuming operations such as synchronously rendering frames

Is your plug-in affected? Test for these problems:

1. Render not updating after UI parameter change because it depends on sequence_data, which may not be currently copied to render
2. Render not updating during click/drag in the Composition Window (similar reasons)
3. Custom Effect UI not updating because it depends on sequence_data generated in render (which is no longer available to the UI because it is in a different project, the render project is immutable, and cache contains previously-rendered frames)
4. Errors telling you an operation on the render thread (or UI thread) is not expected

Generally, calculations that will persist or update the UI will now have to be pulled from the UI thread rather than pushed from the render thread. These cases can require use of new

13.5 APIs or different solutions than in past releases.

- The Need For More Efficient Sequence Data Handling

PF_OutFlag2_SUPPORTS_GET_FLATTENED_SEQUENCE_DATA

PF_Cmd_GET_FLATTENED_SEQUENCE_DATA

Up to version 13.2, serializing/flattening sequence_data always involved deallocating and reallocating any data structures. Starting in 13.5, as effect changes are made, serializing/ flattening happens even more often. Why? AE needs to serialize/flatten project changes to send from the UI thread to the render thread, to keep them both synchronized.

To make this process more efficient, starting in 13.5, AE can send PF_Cmd_GET_FLATTENED_SEQUENCE_DATA to request sequence data without requiring the existing data to be deallocated and reallocated. The main difference between this selector and PF_Cmd_SEQUENCE_FLATTEN is that a copy of the correct flattened state is returned without disposing the original structure(s) the effect is currently using. For a code example, refer to the PathMaster sample project.

This will eventually become required for plug-ins that are rebuilt to be thread-safe (see PF_OutFlag2_AE13_5_THREADSAFE below). The venerable PF_Cmd_SEQUENCE_FLATTEN will eventually be unsupported in future versions.

- PF_OutFlag_FORCE_RERENDER 更改

在可能的情况下，我们建议使用以下方法之一来触发重做。GuidMixInPtr()(在下一节中描述)，arb 数据，或 PF_ChangeFlag_CHANGED_VALUE。所有这些都允许在撤销后重新使用缓存的帧。

注意：从 14.0 开始，为图层或路径参数设置 PF_ChangeFlag_CHANGED_VALUE 不会触发重新渲染。相反，你可以使用 AEGP_StreamSuit->AEGP_SetStreamValue 来改变设置值。

> AEGP_SetStreamValue()。

当序列数据需要从 UI 线程复制到渲染项目/效果克隆中以保持它们的匹配时，仍然需要 FORCE_RERENDER。

无论渲染请求是否使用缓存，FORCE_RERENDER 都是这种情况的触发器。一旦我们有了管理渲染状态所需的全套 API，我们就可以废除 FORCE_RERENDER 了。

FORCE_RERENDER 并不是在以前的每一种情况下都能工作，因为它需要将 UI 的 sequence_data 副本与渲染线程的副本同步。

FORCE_RERENDER 在 PF_Cmd_USER_CHANGED_PARAM 中设置时可以工作。它在点击和拖动事件中也起作用，但只有在 PF_Cmd_GET_FLATTENED_SEQUENCE_DATA 被实现的情况下。这是必要的，以防止在鼠标操作过程中出现扁平化和 UI 状态的丢失。如果没有 GET_FLATTENED，新的 FORCE_RERENDER 行为将不会被打开。

- 缓存帧的 GUIDs

PF_OutFlag2_I_MIX_GUID_DEPENDENCIES

GuidMixInPtr()

仅由 SmartFX 使用。如果自定义 UI 或 PF_Cmd_DO_DIALOG 改变了序列数据，或者渲染结果取决于其他没有考虑的东西，可能需要重新渲染，请使用这个。在 PF_Cmd_SMART_PRERENDER 过程中，效果可以调用 GuidMixInPtr()来将任何影响渲染的额外状态混合到我们的内部 GUID 中，用于缓存帧。使用这个 GUID，AE 可以知道这个帧是否已经存在，或者是否需要被渲染。请看 SmartyPants 示例项目中的一个例子。

这是对旧机制 PF_OutFlag_FORCE_RERENDER 和 PF_Cmd_DO_DIALOG 的改进，后者会从缓存中删除帧，因为主机不知道插件在渲染中还考虑了什么因素。这也可以用来代替 PF_OutFlag2_OUTPUT_IS_WATERMARKED。

- 异步请求帧而不阻塞 UI

PF_OutFlag2_CUSTOM_UI_ASYNC_MANAGER

PF_GetContextAsyncManager() AEGP_CheckoutOrRender_ItemFrame_AsyncManager() AEGP_CheckoutOrRender_LayerFrame_AsyncManager()

对于这种渲染以前是由副作用或取消的隐含性触发的情况下

(如自定义 UI 直方图的绘制)，并且在插件内部的寿命不太明确，使用新的 "Async Manager"，它可以处理效果自定义 UI 的多个同时的异步请求，并将自动支持与其他 AE UI 行为的交互。

注意：在处理被动绘图的情况下，异步检索帧是首选，但当用户行为将更新项目状态时，则不适用。如果你(1)对特定的用户点击做出反应，并且(2)你需要因此而更新项目，建议使用同步的 AEGP_RenderAndCheckoutLayerFrame()。

SDK 中新的 HistoGrid 示例展示了当需要 1 个或多个帧渲染时，如何在 UI 线程上进行完全异步的自定义 UI DRAW 事件处理，例如计算显示在效果窗格中的直方图。请注意仍有一个已知的错误，即拖动改变上游参数可能不会刷新直方图的绘制，直到鼠标悬停在它上面。

- 从效果的用户界面获取效果的渲染输出

诸如抠像器或那些绘制后处理视频直方图的效果，可以使用 AEGP_LayerRenderOptionsSuite 中的新函数 AEGP_NewFromDownstreamOfEffect()检索所需的 AEGP_LayerRenderOptionsH。这个函数只能从 UI 线程中调用。

- 渲染线程上的 AEGP 使用方法

我们加强了对 AEGP 调用可能被危险使用的验证(例如从错误的线程或在渲染中对项目状态进行更改)。如果代码碰到了这种情况，你可能会看到新的错误。例如，在渲染线程中进行这些调用会导致错误。

suites.UtilitySuite5()->AEGP_StartUndoGroup() suites.StreamSuite2()->AEGP_GetStreamName() suites.StreamSuite2()->AEGP_SetExpressionState() suites.StreamSuite2()->AEGP_SetExpression() suites. 流套件 2()->AEGP_GetNewLayerStream() suites.StreamSuite2()->AEGP_DisposeStream() suites.EffectSuite3()->AEGP_DisposeEffect() suites.UtilitySuite5()->AEGP_EndUndoGroup()

解决方案是将这些调用转移到 UI 线程。用于被动 UI 更新的选择器(如 PF_EVENT_DRAW)不是一个对项目状态进行更改的地方。

另一个有更严格要求的例子是 AEGP_RegisterWithAEGP()。文档中一直指出，这个函数必须在 PF_Cmd_GLOBAL_SETUP 上调用。然而在以前的版本中，插件可以在其他时间调用这个函数而不会遇到麻烦。在 13.5 中不再是这样了! 在其他时间调用这个函数会导致崩溃!

- PF_Cmd_SEQUENCE_RESETUP 在 UI 或渲染线程上被调用？

现在有一个 PF_InFlag_PROJECT_IS_RENDER_ONLY 标志，只在 PF_Cmd_SEQUENCE_RESETUP 中有效，它将告诉你这个效果实例是否只用于渲染目的。如果是这样，项目应该被视为完全只读，你将不会在该效果实例上收到与 UI 相关的选择器。这可以用来优化任何渲染不需要的 UI-only 初始化。如果这个标志是假的，你应该正常地设置 UI。这不应该被用来避免报告渲染中的错误。渲染中的错误应该像往常一样通过现有的 SDK 机制报告。

- 为避免死锁的变化

在开发过程中，人们注意到在特定的调用用法中可能会出现死锁。为了避免这种情况，我们引入了安全带。这些情况发生在 PF_Cmd_UPDATE_PARAMS_UI 中，当使用特定的调用时，因为这些调用在 UI 中使用时被废弃的同步行为。

仅在 PF_Cmd_UPDATE_PARAMS_UI 中，PF_PARAM_CHECKOUT()对于图层参数的行为将和以前一样，只是它将返回一个相同大小的黑框，等等，而不是实际渲染的像素。用于启用/禁用参数检测的代码应该仍然像以前一样工作。在 PF_Cmd_UPDATE_PARAMS_UI 之外使用这个来获取分析帧等的代码将和以前一样工作。

仅在 PF_Cmd_UPDATE_PARAMS_UI 中，PF_GetCurrentState()现在将返回一个随机的 GUID。在这种情况下，这将不再像以前那样起作用，但它在其他地方将继续正常工作。

上述用途应该是很少见的，但如果这影响到你，请联系我们了解解决方法。

- 废弃的

AEGP_RenderAndCheckoutFrame()(在 UI 线程上)。这个调用一般不应该在 UI 线程上使用，因为同步渲染会阻止交互性。

在渲染线程中使用就可以了。有一种情况下，在 UI 线程上可能仍然有用，那就是像一个 UI 按钮，需要一个框架来计算一个参数，然后更新 AE 项目。

例如，一个 "自动颜色 "按钮，需要一个帧，然后调整效果参数的结果。

如果这个阻塞操作很慢，已经实现了一个进度对话框的测试版，但是在 UI 线程上使用这个调用应该仅限于这种特殊情况。对话框的设计还没有最终完成。

- 线程安全效应的标志

PF_OutFlag2_AE13_5_THREADSAFE

为线程更新的插件应该使用这个标志来告诉 AE，该插件应该是 UI 线程<>渲染线程安全的。

这个标志告诉 AE，不同 AE 项目副本上的不同线程可以同时进入效果，但不能访问同一个实例。虽然多渲染线程还没有被使用，但这在未来的版本中会很有用。

- 支持大于 7 的效果版本(新的最大版本是 MAJOR 127)。

如果使用当前的 SDK 头文件，大于 7 版的效果现在可以在 13.5 版中正常报告。在 13.5 之前的 AE 版本中可以使用这些重新编译的效果，但在内部，版本号将以 8 为模数(例如，AE 在内部将效果版本 8 视为版本 0)。

这可能会影响老的 AE 在错误对话框中显示的版本，并影响使用报告。

由于许多老的插件在 AE 中无法加载，随着向 64 位的转变，这种包装不太可能对当前使用的实际插件造成歧义(除非这些插件在过去几年中迅速增加了版本号)。

然而，用旧的 SDK 构建并使用 8 或更高的版本将导致插件向 AE 报告一个不正确的版本，这将导致与 PiPL 版本检查的效果不匹配，这将有更高的位设置。这是不被支持的。

如果使用旧的 SDK，你将需要把效果的版本保持在 7 或以下。最大版本的增加是通过在版本中增加 4 个新的更高的有效位来实现的，只有 AE 13.5 和更高的版本才能 "看到"。这些新的高版本位与原来的、预先存在的主要版本位不相邻--只是忽略了中间位。新的版本布局在十六进制或二进制中看起来像这样。

0x 3C38 0000

^^ 原有的主要版本位作为一个十六进制掩码 0-7

^^ 新的高位扩展了原来的 MAJOR 版本位 8-127

0b 0011 1100 0011 1000 0000 0000 0000 0000

^^ 原有的 MAJOR 版本位作为一个十六进制掩码 0-7

^^ 忽略/不使用

^^^ 新的高位扩展了原来的 MAJOR 版本位 8-127。

这些位在 13.5 以上的 AE 版本中被忽略。

- 针对 macOS 的新安装程序提示

开发者可以在一个新的 plist 文件中找到 macOS X 上插件、脚本和预置的默认位置的路径(与 Windows 注册表中的路径相同)。/Library/Preferences/ com.Adobe.After Effects.paths.plist

你可以使用这个 plist 中的值来指导你的安装程序或脚本写文件的位置，就像你在 Windows 上使用注册表中的 paths 键一样。HKEY_LOCAL_MACHINESOFTWAREAdobeAfter Effects13.5

- 正在进行的工作

AEGP_RenderAndCheckoutLayerFrame_Async() AEGP_CancelAsyncRequest()

这些 API 正在进行中，还不应该使用。

## What’s New In CC 2014.1 (13.1)?

PF_CreateNewAppProgressDialog()

它不会打开对话框，除非它检测到缓慢的渲染。(2 秒超时)。

## What’s New In CC 2014 (13.0)?

从 CC 2014 开始，After Effects 现在会尊重使用[PF_UpdateParamUI](../effect-details/parameter-supervision.html)对自定义 UI 高度进行的更改。

[AEGP 效果套件](../aegps/aegp-suites.html) 现在是第 4 版，增加了与效果遮罩有关的新功能。[AEGP_RenderSuite](../aegps/aegp-suites.html) 现在是第 4 版，增加了一个新函数`AEGP_RenderAndCheckoutLayerFrame`，它允许对当前图层进行帧检出。

它允许在非渲染时间内对当前图层的效果进行帧检查。这对需要框架的操作很有用，例如，当一个按钮被点击时，在渲染时等待一会儿是可以接受的。

::: tip

由于它不是异步的，所以它不能解决自定义 UI 需要根据框架来绘制的一般问题。
:::
使用新的[AEGP_LayerRenderOptionsSuite](../aegps/aegp-suites.html)(#aegps-aegp-suites-aegp-layerrenderoptionssuite)指定图层渲染选项。

现在支持 [Mercury Transmit](other-integration-possibilities.html) 插件和 [HTML5 Panels](other-integration-possibilities.html) 。

## What’s New In CC (12.0)?

效果名称的长度现在可以达到 47 个字符，而之前是 31 个字符。

我们增加了[PF_AngleParamSuite](../effect-details/parameters-floating point-values.html) ，提供了一种获得角度参数浮点值的方法。[PF 应用套件](../effect-details/useful-utility-functions.html) 第五版增加了`PF_AppGetLanguage`来查询当前的语言，以便插件可以使用正确的语言字符串，以及几个新的 PF_App_ColorType 枚举值，可以查询到新元素的颜色。

[AEGP 持久性数据套件](../aegps/aegp-suites.html) 现在是第 4 版，为 AEGP_GetApplicationBlob 增加了一个新参数，可以在检索几个不同的应用程序 blob 之间选择。还有一些新的函数来获取/设置时间和 ARGB 值。

[AEGP Composition Suite](../aegps/aegp-suites.html) 现在是第 10 版，增加了新函数来检查/修改是否显示层名或源名，以及是否显示混合模式列。另外还增加了新的函数来获取和设置运动模糊的自适应采样限制。

[AEGP Layer Suite](../aegps/aegp-suites.html) 现在是第 8 版，增加了设置/获取图层采样质量的新功能。[AEGP_CanvasSuite](../artisans/artisan-data-types.html) 现在也是第 8 版。新函数`AEGP_MapCompToLayerTime`处理折叠或嵌套合成的时间重映射，与 AEGP_ConvertCompToLayerTime 不同。

[AEGP_UtilitySuite](../aegps/aegp-suites.html) 现在是第 6 版，增加了一个新的 Unicode 感知函数：`AEGP_ReportInfoUnicode'。另一个新函数，`AEGP_GetPluginPaths`，提供了一些与插件和 After Effects 可执行文件本身有关的有用路径。

更新了`AEGP_NewPlaceholderFootageWithPath`的行为，所以现在应该正确设置 file_type，否则会出现警告。

`AEGP_InsertMenuCommand`现在可以在文件>新建子菜单中插入菜单项。

[AEGP_IOInSuite](../aeios/new-kids-on-the-function-block.html) 现在是第 5 版，增加了获取/设置/清除本地开始时间，以及获取/设置镜头下降帧设置的新函数。

## What’s New In CS6.0.1 (11.0.1)?

11.0.1 中的新功能，AE 效果的 API 版本已经增加到了 13.3。

这使得效果可以区分 11.0 和 11.0.1。

11.0 中全局性能缓存有一个错误，当一个 SmartFX 效果同时使用`PF_OutFlag2_AUTOMATIC_WIDE_TIME_INPUT`和`PF_OutFlag_NON_PARAM_VARY`。

在`PF_Cmd_SMART_PRE_RENDER`过程中调用`checkout_layer`会在`PF_CheckoutResult`中返回空矩形。

解决方法是简单地再次进行调用。这个解决方法在 11.0.1 中不再需要了。

## What’s New In CS6 (11.0)?

我们已经做了一些改进，以更好地处理参数 UI。`PF_PUI_INVISIBLE`参数 UI 标志现在在 After Effects 中被支持，如果你的插件需要影响渲染的隐藏参数，这很有用。现在，当插件使用[PF_UpdateParamUI](../effect-details/parameter-supervision.html) 禁用一个参数时，我们现在将该状态保存在 UI 标志中，以便插件在未来可以检查该标志是否被禁用。一个新的标志，`PF_ParamFlag_SKIP_REVEAL_WHEN_UNHIDDEN'，允许一个参数被取消隐藏，而不需要旋转打开任何父类，也不需要滚动参数到效果控制面板和时间轴面板中。

当插件处于试用模式时，在输出上渲染水印的效果，现在可以使用新的`PF_OutFlag2_OUTPUT_IS_WATERMARKED'告诉 After Effects 水印渲染模式是开还是关。

新的全局性能缓存意味着你必须告诉 After Effects 丢弃旧的缓存帧[当改变你的效果的渲染时](../effect-details/tips-tricks.html) 。

我们删除了`PF_HasParamChanged`和`PF_HaveInputsChangedOverTimeSpan`，改为提供[PF_AreStatesIdentical](../effect-details/parameter-supervision.html) 。

提供自定义用户界面的效果现在可以接收`PF_Event_MOUSE_EXITED`，以获得鼠标退出层或 comp 面板的通知。`PF_ParamUtilsSuite`现在是第 3 个版本。

`PF_GET_PLATFORM_DATA`现在有新的选择器，用于获取可执行文件和资源文件的宽字符路径。`PF_PlatData_EXE_FILE_PATH_W`和`PF_PlatData_RES_FILE_PATH_W`。以前的非范围选择器现在已经被废弃了。

3D 是 AE CS6 的一个主要主题。增加了一个新的`AEGP_LayerFlag_ENVIRONMENT_LAYER`。增加了许多新的[层流](../aegps/aegp-suites.html) 。

此外，"AEGP_LayerStream_SPECULAR_COEFF "更名为 "AEGP_LayerStream_SPECULAR_INTENSITY"，"AEGP_LayerStream_SHININESS_COEFF "更名为 "AEGP_LayerStream_SPECULAR_SHININESS"，而 "AEGP_LayerStream_METAL_COEFF "则直接更名为 "AEGP_LayerStream_METAL"。

一个新的套件，[AEGP_RenderQueueMonitorSuite](../aegps/aegp-suites.html) ，提供了渲染队列管理器需要的所有信息，以弄清在渲染的任何时刻正在发生什么。

[AEGP Mask Suite](../aegps/aegp-suites.html) 现在是第 6 版，提供了获取和设置遮罩羽化类型的函数。[AEGP Mask Outline Suite](../aegps/aegp-suites.html) 现在是第 3 版，提供了获取和设置遮罩轮廓羽化信息的权限。

依赖于遮罩的效果现在有了一个新的标志，`PF_OutFlag2_DEPENDS_ON_UNREFERENCED_MASKS`。

[AEGP Composition Suite](../aegps/aegp-suites.html) 现在是第 9 版。AEGP_CreateTextLayerInComp 和

AEGP_CreateBoxTextLayerInComp 现在有了一个新的参数：select_new_layerB。

[AEGP Render Suite](../aegps/aegp-suites.html) 现在是第 3 版，增加了一个新的函数来获取渲染收据的 GUID。

最后，我们增加了两个新的只读[Dynamic Stream](../aegps/aegp-suites.html) 标志。`AEGP_DynStreamFlag_SHOWN_WHEN_EMPTY`和`AEGP_DynStreamFlag_SKIP_REVEAL_WHEN_UNHIDDEN`。

对于在 Premiere Pro CS6 中运行的效果，我们增加了从`PF_CHECKOUT_PARAM`中获得 32 位浮点和 YUV 帧的能力。

## …and what was new before CS6?

关于这么久以前的历史，请看过时的 SDK 拷贝(我们不提供；如果有人想让你为古董软件做开发，他们最好提供 SDK)。
