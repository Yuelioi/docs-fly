---
title: Artisan-Data-Types
order: 1
---

# Artisan 数据类型

以下是 Artisan API 中最常用的数据类型。

## Data Types Used In The Artisan API

| 类型                   | 描述                                                                        |
| --------------------------- | --------------------------------------------------------------------------------------- |
| `AEGP_RenderLayerContextH` | State information at the time of a render request, sent to an Artisan by After Effects. |
| `AEGP_SoundDataH`          | The audio settings used for a given layer.                                              |
| `AEGP_RenderOptionsH`       | The settings associated with a render queue item.                                       |

### AEGP_RenderLayerContextH

渲染请求时的状态信息，由 After Effects 发送至 Artisan。

### PR_RenderContextH

设置的集合，定义了要渲染的内容和方式。

### AEGP_SoundDataH

用于给定图层的音频设置。

"AEGP_RenderReceiptH"(渲染接收)。

`AEGP_FrameReceiptH`在渲染时由Artisan们使用。

### AEGP_WorldH

一个像素的框架。

### AEGP_RenderOptionsH

与一个渲染队列项目相关的设置。

## 行还是列?

After Effects 的矩阵是基于行的，OpenGL 的是基于列的。这意味着你要做更多的工作。耶，计费时间!

## 实施与设计

一个Artisan几乎是一个独立的应用程序。因为我们在 After Effects 5.0 的早期就意识到，有很多方法可以解决 3D 渲染中固有的问题；例如，交叉点和阴影等。

我们提供了一个 API，我们和第三方(是的，我们真的使用自己的 API)可以实现任何需要的 3D 渲染方案。

## 3D 合成，而非建模

After Effects 不是一个 3D 建模应用程序。用户在响应模式下工作，只有在打样或最终输出时才切换到更高的质量。考虑提供至少两种质量模式，一种用于布局，另一种用于最终输出。在低质量模式下要注意渲染时间。

## 注册 Artisan

Artisan是一个 AEGP，有单一的入口函数。Artisan也必须注册自己的函数入口函数，并为此有一个特殊的回调。参见[AEGP_RegisterSuites5](../aegps/aegp-suites.html#aegps-aegp-suites-aegp-registersuites)中的`AEGP_RegisterArtisan()`。

这个表格显示了由`PR_ArtisanEntryPoints`定义的 Artisans 可以支持的功能：只有`render_func`是必须的。

### Artisan 入口函数

**PR_ArtisanEntryPoints**

#### global_setup_func0

只调用一次，就在`GP_Main`之后。全局数据在插件的所有实例中是通用的。如果你在全局设置期间分配了内存，你必须在`global_setdown_func`期间释放它。

```cpp

PR_GlobalSetupFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_GlobalDataH *global_dataPH);
```

#### global_setdown_func0

处置你分配的任何全局数据。

```cpp

PR_GlobalSetdownFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_GlobalDataH global_dataH);
```

#### global_do_about_func0

告诉世界你的情况! 使用`in_dataP>msg_func`来显示你的对话框。

```cpp

PR_GlobalDoAboutFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_GlobalDataH global_dataH);
```

#### setup_instance_func0

分配和实例化任何特定于你的 Artisan 实例的数据。

```cpp

PR_InstanceSetupFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceFlags flags,

 PR_FlatHandle flat_dataH0,

 PR_InstanceDataH *instance_dataPH);
```

#### setdown_instance_func0

卸载并释放你 Artisan 实例的任何特定数据。

```cpp

PR_InstanceSetdownFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH);
```

#### flatten_instance_func0

扁平化你的数据，准备写入磁盘。(确保它是独立于操作系统的，如果你的 Artisan 是这样的话)。

```cpp

PR_FlattenInstanceFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH,

 PR_FlatHandle *flatH);
```

#### do_instance_dialog_func0

如果你的 Artisan 有一个额外的参数(通过它的选项对话框访问)，这个函数将被调用以获得和设置它们。

```cpp

PR_DoInstanceDialogFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH,

 PR_DialogResult *resultP);
```

`PR_DialogResultis`是`PR_DialogResult_NO_CHANGE`或`PR_DialogResult_CHANGE_MADE`。

#### frame_setup_func0

执行渲染一帧所需的任何设置(在渲染前立即调用)。

```cpp

PR_FrameSetupFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH

 PR_RenderContextH render_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH,

 PR_RenderDataH *render_dataPH);
```

#### frame_setdown_func0

处理在`frame_setup`期间分配的任何设置数据(在渲染后立即发送)。

```cpp

PR_FrameSetdownFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH

 PR_RenderContextH render_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH,

 PR_RenderDataH render_dataH);
```

#### render_func

渲染场景。

```cpp

PR_FrameRenderFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH

 PR_RenderContextH render_contextH,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH,

 PR_RenderDataH render_dataH);
```

#### query_func0

如果有需要，美工人员可以绘制他们自己的投影轴。

After Effects 将调用这个函数来获得合成世界和这些轴之间的变换，以及其他一些与屏上和屏下预览绘制有关的功能(前者只与交互式Artisan有关)。

```cpp

PR_QueryFunc(

 const PR_InData *in_dataP,

 PR_GlobalContextH global_contextH,

 PR_InstanceContextH instance_contextH

 PR_QueryContextH query_contextH,

 PR_QueryType query_type,

 PR_GlobalDataH global_dataH,

 PR_InstanceDataH instance_dataH);
```

`PR_QueryType`可以是以下之一。

- `PR_QueryType_NONE = 0`,
- `PR_QueryType_TRANSFORM`,
- `PR_QueryType_INTERACTIVE_WINDOW_DISPOSE`,
- `PR_QueryType_INTERACTIVE_WINDOW_CLEAR`,
- `PR_QueryType_INTERACTIVE_WINDOW_FROZEN_PROXY`,
- `PR_QueryType_INTERACTIVE_SWAP_BUFFER`,
- `PR_QueryType_INTERACTIVE_DRAW_PROCS`,
- `PR_QueryType_PREPARE_FOR_LINE_DRAWING`,
- `PR_QueryType_UNPREPARE_FOR_LINE_DRAWING`,
- `PR_QueryType_GET_CURRENT_CONTEXT_SAFE_FOR_LINE_DRAWING`,
- `PR_QueryType_GET_ARTISAN_QUALITY`

CS6 中的新功能。

## 世界是你的画布

`AEGP_RenderTexture()`提供一个图层的原始像素，未经转换，进入一个任意大小的缓冲区。

`AEGP_RenderLayer()`调用整个 After Effects 的渲染管线，包括变换、遮罩等，提供图层在其合成中的样子，在一个合成大小的缓冲区中。

如果被渲染的图层是 3D 的，就会调用默认的(标准 3D)Artisan 来执行任何 3D 几何图形。

你的Artisan可以用它来渲染跟踪哑光层，并且只在严格的二维意义上将它们应用到转换后的三维层。

在渲染之前，After Effects 附带的 Artisans 应用一个反变换来获得方形像素，然后在显示之前重新应用变换。

例如，如果像素的长宽比是 10/11(DV NTSC)，我们就乘以 11/10 来得到方形像素。我们对三维图层进行处理和合成，然后重新分割，以恢复到原始的像素长宽比。

以下套件提供了图层、合成、纹理和目标缓冲区。这是所有Artisan的一个重要套件。

### AEGP_CanvasSuite8

#### AEGP_GetCompToRender

鉴于在渲染时提供给 Artisan 的渲染上下文，返回一个指向合成的句柄。

```cpp

AEGP_GetCompToRender(

 PR_RenderContextH render_contextH,

 AEGP_CompH *compPH)
```

#### AEGP_GetNumLayersToRender

给定渲染上下文，返回 Artisan 需要渲染的层数。

```cpp

AEGP_GetNumLayersToRender(

 PR_RenderContextH render_contextH,

 A_long *num_to_renderPL)
```

#### AEGP_GetNthLayerContextToRender

在确定需要由 Artisan 渲染的总层数后，用于建立一个要渲染的层的列表。

```cpp

AEGP_GetNthLayerContextToRender(

 PR_RenderContextH render_contextH,

 A_long n,

 AEGP_RenderLayerContextH *layer_indexPH)
```

#### AEGP_GetLayerFromLayerContext

给出一个`AEGP_RenderLayerContextH`，检索相关的`AEGP_LayerH`(许多套件函数需要)。

```cpp

AEGP_GetLayerFromLayerContext(

 const PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 AEGP_LayerH *layerPH);
```

#### AEGP_GetLayerAndSubLayerFromLayerContext

允许渲染子层(如 Photoshop 文件中的子图层)。

```cpp

AEGP_GetLayerAndSubLayerFromLayerContext(

 const PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 AEGP_LayerH *layerPH,

 AEGP_SubLayerIndex *sublayerP);
```

#### AEGP_GetTopLayerFromLayerContext

在折叠的几何图形 "打开 "的情况下，它给出了包含图层上下文的根合成中的图层。

关闭折叠的几何图形时，这与`AEGP_GetLayerFromLayerContext`相同。

```cpp

AEGP_GetTopLayerFromLayerContext(

 const PR_RenderContextH r_contextH,

 AEGP_RenderLayerContextH l_contextH,

 AEGP_LayerH *layerPH);
```

#### AEGP_GetCompRenderTime

给定渲染上下文，返回当前的(合成)渲染时间点。

```cpp

AEGP_GetNthLayerIndexToRender(

 PR_RenderContextH render_contextH,

 A_long *time,

 A_long *time_step)
```

#### AEGP_GetCompDestinationBuffer

给定渲染上下文，返回一个缓冲区，用于放置最终渲染的输出。

```cpp

AEGP_GetCompToRender(

 PR_RenderContextH render_contextH,

 AEGP_CompH compH,

 PF_EffectWorld *dst);
```

#### AEGP_GetROI

给出在渲染时提供给 Artisan 的渲染上下文，返回一个指向合成的句柄。

```cpp

AEGP_GetROI(

 PR_RenderContextH render_contextH,

 A_LegacyRect *roiPR);
```

#### AEGP_RenderTexture

给出渲染上下文和图层，返回图层纹理。

所有带 "0 "的参数是可选的；返回的 "PF_EffectWorld "可以是 NULL。

```cpp

AEGP_RenderTexture(

 PR_RenderContextH render_contextH,

 AEGP_LayerH layerH,

 AEGP_RenderHints render_hints,

 A_FloatPoint *suggested_scaleP0,

 A_FloatRect *suggsted_src_rectP0,

 A_Matrix3 *src_matrixP0,

 PF_EffectWorld *render_bufferP);
```

`AEGP_RenderHints`包含一个或多个以下内容。

- `AEGP_RenderHints_NONE`
- `AEGP_RenderHints_IGNORE_EXTENTS`
- `AEGP_RenderHints_NO_TRANSFER_MODE`

`AEGP_RenderHints_NO_TRANSFER_MODE`防止应用不透明和混合模式；用于调用`RenderLayer`。

#### AEGP_DisposeTexture

处置一个获得的图层纹理。

```cpp

AEGP_DisposeTexture(

 PR_RenderContextH render_contextH,

 AEGP_LayerH layerH,

 AEGP_WorldH *dst0);
```

#### AEGP_GetFieldRender

返回给定的`PR_RenderContextH`的字段设置。

```cpp

AEGP_GetFieldRender(

 PR_RenderContextH render_contextH,

 PF_Field *field);
```

#### AEGP_ReportArtisanProgress

给出在渲染时提供给 Artisan 的渲染上下文，返回一个句柄给该合成。

注意：这在 macOS 上不是线程安全的；只有在当前线程 ID 为 0 时才使用这个函数。

```cpp

AEGP_ReportArtisanProgress(

 PR_RenderContextH render_contextH,

 A_long countL,

 A_long totalL);
```

#### AEGP_GetRenderDownsampleFactor

返回 "PR_RenderContextH "的降采样系数。

```cpp

AEGP_GetRenderDownsampleFactor(

 PR_RenderContextH render_contextH,

 AEGP_DownsampleFactor *dsfP);
```

#### AEGP_IsBlankCanvas

确定`PR_RenderContextH`是否为空白(空)。

```cpp

AEGP_IsBlankCanvas(

 PR_RenderContextH render_contextH,

 A_Boolean *is_blankPB);
```

#### AEGP_GetRenderLayerToWorldXform

给出一个渲染环境和一个图层(在给定的时间)，检索 4 乘 4 的变换，在它们的坐标空间之间移动。

```cpp

AEGP_GetRenderLayerToWorldXform(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timeP,

 A_Matrix4 *transform);
```

#### AEGP_GetRenderLayerBounds

检索 layer_contextH(在给定时间)在 render_contextH 中的边界矩形。

```cpp

AEGP_GetRenderLayerBounds(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timeP,

 A_LegacyRect *boundsP);
```

#### AEGP_GetRenderOpacity

返回给定的图层上下文在给定的时间，在渲染上下文中的不透明度。

```cpp

AEGP_GetRenderOpacity(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timePT,

 A_FpLong *opacityPF);
```

#### AEGP_IsRenderLayerActive

返回在给定时间内，渲染上下文中给定的图层上下文是否处于活动状态。

```cpp

AEGP_IsRenderLayerActive(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timePT,

 A_Boolean *activePB);
```

#### AEGP_SetArtisanLayerProgress

CountL 是完成的图层数。

`num_layersL`是 Artisan 正在渲染的总层数。

```cpp

AEGP_SetArtisanLayerProgress(

 PR_RenderContextH render_contextH,

 A_long countL,

 A_long num_layersL);
```

#### AEGP_RenderLayerPlus

与`AEGP_RenderLayer`类似，但考虑到`AEGP_RenderLayerContextH`。

```cpp

AEGP_RenderLayerPlus(

 PR_RenderContextH r_contextH,

 AEGP_LayerH layerH,

 AEGP_RenderLayerContextH l_contextH,

 AEGP_RenderHints render_hints,

 AEGP_WorldH *bufferP);
```

#### AEGP_GetTrackMatteContext

为指定的渲染和填充上下文检索`AEGP_RenderLayerContextH`。

```cpp

AEGP_GetTrackMatteContext(

 PR_RenderContextH rnder_contextH,

 AEGP_RenderLayerContextH fill_contextH,

 AEGP_RenderLayerContextH *mattePH);
```

#### AEGP_RenderTextureWithReceipt

渲染纹理到一个`AEGP_WorldH`，并为该操作提供一个`AEGP_RenderReceiptH`。

返回的`receiptPH`必须用`AEGP_DisposeRenderReceipt`处理。

```cpp

AEGP_RenderTextureWithReceipt(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 AEGP_RenderHints render_hints,

 A_FloatPoint *suggested_scaleP0,

 A_FloatRect *suggest_src_rectP0,

 A_Matrix3 *src_matrixP0,

 AEGP_RenderReceiptH *receiptPH,

 AEGP_WorldH *dstPH);
```

#### AEGP_GetNumberOfSoftwareEffects

返回在给定的`AEGP_RenderLayerContextH`中应用的软件效果的数量。

```cpp

AEGP_GetNumberOfSoftwareEffects(

 PR_RenderContextH ren_contextH,

 AEGP_RenderLayerContextH lyr_contextH,

 A_short *num_sft_FXPS);
```

#### AEGP_RenderLayerPlusWithReceipt

作为对`AEGP_RenderLayerPlus`的改进，该函数还提供了一个`AEGP_RenderReceiptH`用于缓存目的。

```cpp

AEGP_RenderLayerPlusWithReceipt(

 PR_RenderContextH render_contextH,

 AEGP_LayerH layerH,

 AEGP_RenderLayerContextH layer_contextH,

 AEGP_RenderHints render_hints,

 AEGP_NumEffectsToRenderType num_effectsS,

 AEGP_RenderReceiptH *receiptPH,

 AEGP_WorldH *bufferPH);
```

#### AEGP_DisposeRenderReceipt

释放一个`AEGP_RenderReceiptH`。

```cpp

AEGP_DisposeRenderReceipt(

 AEGP_RenderReceiptH receiptH);
```

#### AEGP_CheckRenderReceipt

检查 After Effects 的内部缓存，以确定一个给定的`AEGP_RenderReceiptH`是否仍然有效。

```cpp

AEGP_CheckRenderReceipt(

 PR_RenderContextH current_contextH,

 AEGP_RenderLayerContextH current_lyr_ctxtH,

 AEGP_RenderReceiptH old_receiptH,

 A_Boolean check_aceB,

 AEGP_NumEffectsToRenderType num_effectsS,

 AEGP_RenderReceiptStatus *receipt_statusP);
```

#### AEGP_GenerateRenderReceipt

为一个图层生成一个`AEGP_RenderReceiptH`，就像第一个`num_effectsS`已经被渲染了。

```cpp

AEGP_GenerateRenderReceipt(

 PR_RenderContextH current_contextH,

 AEGP_RenderLayerContextH current_lyr_contextH,

 AEGP_NumEffectsToRenderType num_effectsS,

 AEGP_RenderReceiptH *render_receiptPH);
```

#### AEGP_GetNumBinsToRender

返回 After Effects 希望Artisan渲染的 bins 数量。

```cpp

AEGP_GetNumBinsToRender(

 const PR_RenderContextH contextH,

 A_long *num_binsPL);
```

#### AEGP_SetNthBin

将给定的渲染上下文设置为 After Effects 要渲染的第 n 个 bin。

```cpp

AEGP_SetNthBin(

 const PR_RenderContextH contextH,

 A_long n);
```

#### AEGP_GetBinType

检索给定仓的类型。

```cpp

AEGP_GetBinType(

 const PR_RenderContextH contextH,

 AEGP_BinType *bin_typeP);
```

`AEGP_BinType`将是以下之一。

- `AEGP_BinType_NONE`
- `AEGP_BinType_2D`
- `AEGP_BinType_3D`

#### AEGP_GetRenderLayerToWorldXform2D3D

检索转换，以正确确定被渲染层与输出世界的方向。

为`only_2dB`传递`TRUE`，将变换限制在二维。

```cpp

AEGP_GetRenderLayerToWorldXform2D3D(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timeP,

 A_Boolean only_2dB,

 A_Matrix4 *transformP);
```

::: tip

以下函数仅适用于交互式Artisan。
:::

#### AEGP_GetPlatformWindowRef

检索特定平台的窗口环境，以绘制给定的`PR_RenderContextH`。

```cpp

AEGP_GetPlatformWindowRef(

 const PR_RenderContextH contextH,

 AEGP_PlatformWindowRef *window_refP);
```

#### AEGP_GetViewportScale

检索给定的`PR_RenderContextH`的源到帧的降采样系数。

```cpp

AEGP_GetViewportScale(

 const PR_RenderContextH contextH,

 A_FpLong *scale_xPF,

 A_FpLong *scale_yPF);
```

#### AEGP_GetViewportOrigin

为给定的`PR_RenderContextH`检索源的原点，在帧内(在两者之间平移所需)。

```cpp

AEGP_GetViewportOrigin(

 const PR_RenderContextH contextH,

 A_long *origin_xPL,

 A_long *origin_yPL);
```

#### AEGP_GetViewportRect

为给定的`PR_RenderContextH`检索要绘制的区域的边界矩形。

```cpp

AEGP_GetViewportRect(

 const PR_RenderContextH contextH,

 A_LegacyRect *v_rectPR);
```

#### AEGP_GetFallowColor

检索给定的`PR_RenderContextH`中的休眠区域的颜色。

```cpp

AEGP_GetFallowColor(

 const PR_RenderContextH contextH,

 PF_Pixel8 *fallow_colorP);
```

#### AEGP_GetInteractiveCheckerboard

读取当前是否为给定的`PR_RenderContextH`激活棋盘。

```cpp

AEGP_GetInteractiveCheckerboard(

 const PR_RenderContextH contextH,

 A_Boolean *cboard_onPB);
```

#### AEGP_GetInteractiveCheckerboardColors

读取棋盘上使用的颜色。

```cpp

AEGP_GetInteractiveCheckerboardColors(

 const PR_RenderContextH contextH,

 PF_Pixel *color1P,

 PF_Pixel *color2P);
```

#### AEGP_GetInteractiveCheckerboardSize

检索一个棋盘的宽度和高度。

```cpp

AEGP_GetInteractiveCheckerboardSize(

 const PR_RenderContextH contextH,

 A_u_long *cbd_widthPLu,

 A_u_long *cbd_heightPLu);
```

#### AEGP_GetInteractiveCachedBuffer

检索上次用于`PR_RenderContextH`的缓存 AEGP_WorldH。

```cpp

AEGP_GetInteractiveCachedBuffer(

 const PR_RenderContextH contextH,

 AEGP_WorldH *buffer);
```

#### AEGP_ArtisanMustRenderAsLayer

确定Artisan是否必须将当前的`AEGP_RenderLayerContextH`渲染成一个图层。

```cpp

AEGP_ArtisanMustRenderAsLayer(

 const PR_RenderContextH contextH,

 AEGP_RenderLayerContextH layer_contextH,

 A_Boolean *use_txturePB);
```

#### AEGP_GetInteractiveDisplayChannel

返回哪些通道应该由交互式Artisan显示。

```cpp

AEGP_GetInteractiveDisplayChannel(

 const PR_RenderContextH contextH,

 AEGP_DisplayChannelType *channelP);
```

`AEGP_DisplayChannelType`为以下之一。

- `AEGP_DisplayChannel_NONE`
- `AEGP_DisplayChannel_RED`
- `AEGP_DisplayChannel_GREEN`
- `AEGP_DisplayChannel_BLUE`
- `AEGP_DisplayChannel_ALPHA`
- `AEGP_DisplayChannel_RED_ALT`
- `AEGP_DisplayChannel_GREEN_ALT`
- `AEGP_DisplayChannel_BLUE_ALT`
- `AEGP_DisplayChannel_ALPHA_ALT`

#### AEGP_GetInteractiveExposure

返回给定`PR_RenderContextH`的曝光量，以浮点数表示。

```cpp

AEGP_GetInteractiveExposure(

 const PR_RenderContextH rcH,

 A_FpLong *exposurePF);
```

#### AEGP_GetColorTransform

返回给定的`PR_RenderContextH`的颜色变换。

```cpp

AEGP_GetColorTransform)(

 const PR_RenderContextH render_contextH,

 A_Boolean *cms_onB,

 A_u_long *xform_keyLu,

 void *xformP);
```

#### AEGP_GetCompShutterTime

返回给定`PR_RenderContextH`的快门角度。

```cpp

AEGP_GetCompShutterTime)(

 PR_RenderContextH render_contextH,

 A_Time *shutter_time,

 A_Time *shutter_dur);
```

#### AEGP_MapCompToLayerTime

CC 中的新内容。与[AEGP_ConvertCompToLayerTime](../aegps/aegp-suites.html#aegps-aegp-suites-aegp-layersuite)不同，它可以处理折叠或嵌套的合成的时间重映射。

```cpp

AEGP_MapCompToLayerTime(

 PR_RenderContextH render_contextH,

 AEGP_RenderLayerContextH layer_contextH,

 const A_Time *comp_timePT,

 A_Time *layer_timePT);
```

## Convert Between Different Contexts

在渲染和实例上下文之间进行转换，并管理特定于Artisan的全局数据。

### AEGP_ArtisanUtilSuite1

#### AEGP_GetGlobalContextFromInstanceContext

给定一个实例上下文，返回一个到全局上下文的句柄。

```cpp

AEGP_GetGlobalContextFromInstanceContext(

 const PR_InstanceContextH instance_contextH,

 PR_GlobalContextH *global_contextPH);
```

#### AEGP_GetInstanceContextFromRenderContext

给定渲染上下文，返回到实例上下文的句柄。

```cpp

AEGP_GetInstanceContextFromRenderContext(

 const PR_RenderContextH render_contextH,

 PR_InstanceContextH *instnc_ctextPH);
```

#### AEGP_GetInstanceContextFromQueryContext

给定一个查询上下文，返回一个到实例上下文的句柄。

```cpp

AEGP_GetInstanceContextFromQueryContext(

 const PR_QueryContextH query_contextH,

 PR_InstanceContextH *instnce_contextPH);
```

#### AEGP_GetGlobalData

给定全局上下文，返回全局数据的句柄。

```cpp

AEGP_GetGlobalData(

 const PR_GlobalContextH global_contextH,

 PR_GlobalDataH *global_dataPH);
```

#### AEGP_GetInstanceData

给定一个实例上下文，返回相关的实例数据。

```cpp

AEGP_GetInstanceData(

 const PR_InstanceContextH instance_contextH,

 PR_InstanceDataH *instance_dataPH);
```

#### AEGP_GetRenderData

给定一个渲染上下文，返回相关的渲染数据。

```cpp

AEGP_GetRenderData(

 const PR_RenderContextH render_contextH,

 PR_RenderDataH *render_dataPH);
```

## Smile! Cameras

获取相机的几何形状，包括相机的属性(类型、镜头、景深、焦距、光圈等等)。

### AEGP_CameraSuite2

#### AEGP_GetCamera

给定一个图层句柄和时间，返回当前的摄像机图层句柄。

```cpp

AEGP_GetCamera(

 PR_RenderContextH render_contextH,

 const A_Time *comp_timeP,

 AEGP_LayerH *camera_layerPH);
```

#### AEGP_GetCameraType

给定一个图层，返回该图层的摄像机类型。

```cpp

AEGP_GetCameraType(

 AEGP_LayerH aegp_layerH,

 AEGP_CameraType *camera_typeP;
```

相机类型可以是以下几种。

- `AEGP_CameraType_NONE = -1`
- `AEGP_CameraType_PERSPECTIVE`
- `AEGP_CameraType_ORTHOGRAPHIC`

#### AEGP_GetDefaultCameraDistanceToImagePlane

给定一个合成句柄，返回相机到图像平面的距离。

```cpp

AEGP_GetDefaultCamera DistanceToImagePlane(

 AEGP_CompH compH,

 A_FpLong *dist_to_planePF)
```

#### AEGP_GetCameraFilmSize

检索指定相机使用的胶片的尺寸(以及用于测量该尺寸的单位)。

```cpp

AEGP_GetCameraFilmSize(

 AEGP_LayerH camera_layerH,

 AEGP_FilmSizeUnits *film_size_unitsP,

 A_FpLong *film_sizePF0);
```

#### AEGP_SetCameraFilmSize

设置指定相机使用的胶片的尺寸(和用于测量该尺寸的单位)。

```cpp

AEGP_SetCameraFilmSize)(

 AEGP_LayerH camera_layerH,

 AEGP_FilmSizeUnits film_size_units,

 A_FpLong *film_sizePF0);
```

## 关于相机行为的注意事项

摄像机的方向是以合成坐标为单位，旋转是以层(摄像机的层)坐标为单位。

如果摄像机层有一个父层，那么位置就是相对于父层的坐标空间。

## 正交相机矩阵

在内部，我们使用合成宽度和高度来设置 OpenGL 规范所描述的矩阵为

```cpp

glOrtho(-width/2, width/2, -height/2, height/2, -1, 100);
```

正射矩阵描述了投影。相机的位置由另一个按比例的矩阵描述。摄像机位置矩阵的倒数提供 "眼睛 "坐标。

## 专注于焦距

记住，焦距(Focal)会影响视野；焦距只影响景深。

## Film 尺寸

在现实世界中，胶片的尺寸是以毫米为单位的。在 After Effects 中，它是以像素来衡量的。乘以 72，再除以 25.4，就可以从毫米变成像素。

视野是比较复杂的。

ϴ=1/2 视场

tan(ϴ) = 1/2 合成高度/焦距

焦距 = 2 tan(ϴ) / 合成高度

## 打灯

获取和设置合成中的灯光类型。

### AEGP_LightSuite2

#### AEGP_GetLightType

检索指定相机层的`AEGP_LightType`。

```cpp

AEGP_GetLightType(

 AEGP_LayerH light_layerH,

 AEGP_LightType *light_typeP);
```

`AEGP_LightType`将是以下其中之一。

- `AEGP_LightType_PARALLEL`
- `AEGP_LightType_SPOT`
- `AEGP_LightType_POINT`
- `AEGP_LightType_AMBIENT`

#### AEGP_SetLightType

设定指定摄影机层的`AEGP_LightType`。

```cpp

AEGP_SetLightType(

 AEGP_LayerH light_layerH,

 AEGP_LightType light_type);
```

### 关于灯光行为的注意事项

平行光的公式可以在 Foley 和 Van Dam 的《计算机图形学导论》(ISBN 0-201-60921-5)中找到，点光源的公式也是如此。

我们使用 Jim Blinn 提出的半角变体来代替。

假设我们在一个图层上有一个点，想用灯光对其进行遮挡。

让 V 为从图层点到眼睛点的单位向量。

让 L 是到光的单位向量(在平行光的情况下这是常数)。让 H 是(V+L)/2(归一化)。

让 N 为该层的单位法向量。

镜面反射光的数量是 S \* power(H Dot N, shine)，其中 S 是镜面反射系数。

## 我应该如何绘制

After Effects 依靠 Artisans 来绘制 3D 图层手柄。如果你的美工选择不响应这个调用，默认的美工将为你绘制 3D 图层手柄。查询变换对于优化 After Effects 的缓存非常重要。

坐标系统是正 X 向右，正 Y 向下，正 Z 向屏幕。原点是左上角。旋转是先 x 后 y 再 z。对于矩阵来说，平移是最下面一行，方向是四元数(先应用)，然后是任何 x-y-z 的旋转。作为一般规则，使用方向或旋转，但不能同时使用。如果你需要控制角速度，也可以使用旋转。

## 查询变换函数

这些函数给Artisan提供了他们需要的变换(Transform)信息，以便正确地将图层放置在一个合成中，并对 After Effects 发送的各种查询作出适当的反应，这些查询将发送到他们的`PR_QueryFunc`入口函数函数。

由于该入口函数是可选的，所以你的Artisan对查询的响应也是可选的；然而，如果你不这样做，你的用户可能会失望，(在进行交互式预览绘图时)所有的相机和灯光指示器都消失了，直到他们停止移动 Artisan是复杂的野兽；如果你有任何问题，请联系我们。

### AEGP_QueryXFormSuite2

#### AEGP_QueryXformGetSrcType

给定一个查询上下文，返回当前正在修改的 trasnsform 源。

```cpp

AEGP_QueryXformGetSrcType(

 PR_QueryContextH query_contextH,

 AEGP_QueryXformType *src_type);
```

查询上下文将是以下内容之一。

- `AEGP_Query_Xform_LAYER`,
- `AEGP_Query_Xform_WORLD`,
- `AEGP_Query_Xform_VIEW`,
- `AEGP_Query_Xform_SCREEN`

#### AEGP_QueryXformGetDstType

给定一个查询上下文，返回当前请求的转换目的地。

```cpp

AEGP_QueryXformGetDstType(

 PR_QueryContextH query_contextH,

 AEGP_QueryXformType *dst_type);
```

#### AEGP_QueryXformGetLayer

如果源或目的地类型是一个层，则使用。给定一个查询上下文，返回图层句柄。

```cpp

AEGP_QueryXformGetLayer(

 PR_QueryContextH query_contextH,

 AEGP_LayerH *layerPH);
```

#### AEGP_QueryXformGetComp

给定一个查询上下文，返回当前的合成句柄。

```cpp

AEGP_QueryXformGetComp(

 PR_QueryContextH query_contextH,

 AEGP_CompH *compPH);
```

#### AEGP_QueryXformGetTransformTime

给定一个查询上下文，返回转换的时间。

```cpp

AEGP_QueryXformGetTransformTime(

 PR_QueryContextH query_contextH,

 A_Time *time);
```

#### AEGP_QueryXformGetViewTime

给定一个查询上下文，返回相关视图的时间。

```cpp

AEGP_QueryXformGetViewTime(

 PR_QueryContextH query_contextH,

 A_Time *time);
```

#### AEGP_QueryXformGetCamera

给定一个查询上下文，返回当前的相机层句柄。

```cpp

AEGP_QueryXformGetCamera(

 PR_QueryContextH query_contextH,

 AEGP_LayerH *camera_layerPH);
```

#### AEGP_QueryXformGetXform

给定一个查询上下文，返回当前的矩阵变换。

```cpp

AEGP_QueryXformGetXform(

 PR_QueryContextH query_contextH,

 A_Matrix4 *xform);
```

#### AEGP_QueryXformSetXform

给定一个查询上下文，返回你在`xform`中计算的矩阵变换。

```cpp

AEGP_QueryXformSetXform(

 PR_QueryContextH query_contextH,

 A_Matrix4 *xform);
```

#### AEGP_QueryWindowRef

为给定的`PR_QueryContextH`设置窗口参考(由 After Effects)。

```cpp

AEGP_QueryWindowRef(

 PR_QueryContextH q_contextH,

 AEGP_PlatformWindowRef *window_refP);
```

#### AEGP_QueryWindowClear

返回要清除的`AEGP_PlatformWindowRef`(和`A_Rect`)，为给定的`PR_QueryContextH`。

```cpp

AEGP_QueryWindowClear(

 PR_QueryContextH q_contextH,

 AEGP_PlatformWindowRef *window_refP,

 A_LegacyRect *boundsPR);
```

#### AEGP_QueryFrozenProxy

返回在给定的`PR_QueryContextH`中使用的纹理是否应该被冻结。

```cpp

AEGP_QueryFrozenProxy(

 PR_QueryContextH q_contextH,

 A_Boolean *onPB);
```

#### AEGP_QuerySwapBuffer

在渲染和摄像机/灯光手柄绘制完成后发送；After Effects 返回Artisan应该绘制其输出的缓冲区。

```cpp

AEGP_QuerySwapBuffer(

 PR_QueryContextH q_contextH,

 AEGP_PlatformWindowRef *window_refP,

 AEGP_WorldH *dest_bufferp);
```

#### AEGP_QueryDrawProcs

设置 After Effects 在绘制相机和灯光手柄到Artisan提供的上下文时将调用的交互式绘图函数。

```cpp

AEGP_QueryDrawProcs(

 PR_QueryContextH query_contextH,

 PR_InteractiveDrawProcs *window_refP);
```

#### AEGP_QueryPrepareForLineDrawing

告知 After Effects 它将绘制的上下文。

```cpp

AEGP_QueryPrepareForLineDrawing(

 PR_QueryContextH query_contextH,

 AEGP_PlatformWindowRef *window_refP,

 A_LegacyRect *viewportP,

 A_LPoint *originP,

 A_FloatPoint *scaleP);
```

#### AEGP_QueryUnprepareForLineDrawing

就 After Effects 而言，Artisan已经完成了画线。

```cpp

AEGP_QueryUnprepareForLineDrawing(

 PR_QueryContextH query_contextH,

 AEGP_PlatformWindowRef *window_refP);
```

## 交互式绘图功能

我们增加了Artisan提供 After Effects 可以用来做基本绘图功能的功能，以便在预览期间更新合成窗口的显示，包括相机、灯光和线框预览建模。

### PR_InteractiveDrawProcs

#### PR_Draw_MoveToFunc

```cpp

PR_Draw_MoveToFunc(

 short x,

 short y);
```

#### PR_Draw_LineToFunc

```cpp

PR_Draw_LineToFunc(

 short x,

 short y);
```

#### PR_Draw_ForeColorFunc

```cpp

PR_Draw_ForeColorFunc(

 const A_Color *fore_colo
```

#### PR_Draw_FrameRectFunc

```cpp

PR_Draw_FrameRectFunc(

 const A_Rect *rectPR );
```

#### PR_Draw_PaintRectFunc

```cpp

PR_Draw_PaintRectFunc(

 const A_Rect *rectPR );
```

## 查询时间函数注意事项

`AEGP_QueryXformGetTransformTime()`和`AEGP_QueryXformGetViewTime()`都是Artisan建立要渲染的场景表示所必需的。

`AEGP_QueryXformGetTransformTime()`获得转换的时间，然后传递给[AEGP_CompSuite11](../aegps/aegp-suites.html#aegps-aegp-suites-aegp-compsuite)的`AEGP_GetCompShutterFrameRange()`。

`AEGP_QueryXformGetViewTime()`获得视图的时间，用于调用[AEGP_LayerSuite8](./aegps/aegp-suites.html#aegps-aegp-suites-aegp-layersuite)中的`AEGP_GetLayerToWorldXformFromView()`。
