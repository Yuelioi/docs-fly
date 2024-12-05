---
title: Data-Types
order: 5
---

# 数据类型

只要有可能，After Effects 就会为插件提供不透明的数据类型，并提供操作这些数据的访问函数。例如，视频帧用不透明的 AEGP_WorldH 表示。虽然在某些情况下，简单地修改底层结构可能更有效，但通过保持数据类型的不透明性，我们允许改变我们的实现，而不会让你重新编译(和重新发布)你的插件。

## AEGP API 数据类型

|类型 | 描述 | 管理用法  |
| --- | --- | --- |
| `AEGP_MemHandle` | 这个结构包含的不仅仅是引用的内存。所以它不应该被直接取消引用。使用 `AEGP_LockMemHandle` 在AEGP内存套件中获取指向AEGP_MemHandle引用的内存的指针。在完成后解锁它。  | [AEGP Memory Suite](../aegps/aegp-suites.html)  |
| `AEGP_ProjectH`  | 当前AE工程. 项目是在树中分层排列的一组元素，以保留语义关系。树的内部节点是文件夹。从CS6开始，将只有一个开放项目。 | [AEGP Project Suite](../aegps/aegp-suites.html)  |
| `AEGP_ItemH` |描述项目的任何元素的抽象，包括文件夹。项目是可以选择的任何东西。由于可以选择多种对象类型，我们将它们视为`AEGP_ItemHs`,除非有其他需求。 | [AEGP Item Suite](../aegps/aegp-suites.html)  |
| `AEGP_Collection2H`  |一组选择的项目| [AEGP Collection Suite](../aegps/aegp-suites.html)  |
| `AEGP_CompH` | 合成是可渲染项目的序列，它们一起产生输出。合成存在于一个时间间隔内。一个项目中可以存在多个合成。 | [AEGP Composition Suite](../aegps/aegp-suites.html)  |
| `AEGP_FootageH`  | 可以渲染的项目。文件夹和合成是唯一不是素材的项目。  | [AEGP Footage Suite](../aegps/aegp-suites.html)  |
| `AEGP_LayerH`  | 合成的元素。图层按顺序呈现，允许遮挡。纯色、文本、绘图(paint)、相机、灯光、图像和图像序列都表示为图层。图层可以在合成时间间隔的子间隔上定义。   | [AEGP Layer Suite](../aegps/aegp-suites.html)  |
| `AEGP_WorldH`  | 一帧像素. | [AEGP World Suite](../aegps/aegp-suites.html)  |
| `AEGP_EffectRefH`  | 应用于图层的效果。效果是一个函数，它将图层（可能还有其他参数）作为参数，并返回修改后的图层版本以进行渲染。   | [AEGP Effect Suite](../aegps/aegp-suites.html)  |
| `AEGP_StreamRefH`  | 任何[参数流](../aegps/aegp-suites.html)附加到图层，合成。参见来自[AEGP_StreamSuite5](../aegps/aegp-suites.html)的`AEGP_GetNewLayerStream`描述,来获取流类型的完整列表。| [AEGP Stream Suite](../aegps/aegp-suites.html), [AEGP Dynamic Stream Suite](../aegps/aegp-suites.html) [AEGP Keyframe Suite](../aegps/aegp-suites.html) |
| `AEGP_MaskRefH`  | 应用于图层的蒙版。`AEGP`的`_MaskRefH`用于访问有关蒙版流的详细信息，而不是构成蒙版的特定点(specific points)。蒙版是一种光栅化路径（顶点序列），它将图层分成两部分，允许每一部分以不同的方式呈现。 | [AEGP Mask Suite](../aegps/aegp-suites.html)  |
| `AEGP_MaskOutlineValH` | 构成蒙版的具体点。蒙版大纲中的点是有序的，蒙版无需关闭。| [AEGP Mask Outline Suite](../aegps/aegp-suites.html)  |
| `AEGP_TextDocumentH` | 表示与文本图层关联的实际文本。 | [AEGP Text Document Suite](../aegps/aegp-suites.html)  |
| `AEGP_TextOutlinesH` | 给定文本图层轮廓所有路径的引用。  | [AEGP Text Layer Suite](../aegps/aegp-suites.html)  |
| `AEGP_MarkerVal` | 与给定时间轴标记相关联的数据。  | [AEGP Marker Suite](../aegps/aegp-suites.html)  |
| `AEGP_PersistentBlobH` | 包含当前首选项的数据“blob”。 | [AEGP Persistent Data Suite](../aegps/aegp-suites.html)  |
| `AEGP_RenderOptionsH`  | 与渲染请求关联的设置。 | [AEGP Render Options Suite](../aegps/aegp-suites.html)  |
| `AEGP_LayerRenderOptionsH` | 与图层渲染请求关联的设置。 | [AEGP Layer Render Options Suite](../aegps/aegp-suites.html)  |
| `AEGP_FrameReceiptH` | 对渲染帧的引用。  | [AEGP Render Suite](../aegps/aegp-suites.html)  |
| `AEGP_RQItemRefH`  | 渲染队列中的一项。 | [AEGP Render Queue Suite](../aegps/aegp-suites.html) [AEGP Render Queue Item Suite](../aegps/aegp-suites.html) |
| `AEGP_OutputModuleRefH`  | 一个输出模块，附加到渲染队列中的特定`AEGP_RQItemRef`。 | [AEGP Output Module Suite](../aegps/aegp-suites.html)  |
| `AEGP_SoundDataH`  | 给定图层的[音频设置](../aegps/aegp-suites.html) | [AEGP Sound Data Suite](../aegps/aegp-suites.html)  |
| `AEGP_RenderLayerContextH` | 呈现请求时的状态信息，由AE发送给Artisan  | [AEGP Canvas Suite](../artisans/artisan-data-types.html)  |
| `AEGP_RenderReceiptH`  | Artisans在渲染时使用它。| [AEGP Canvas Suite](../artisans/artisan-data-types.html)  |

## Nasty, Brutish, and Short

关于图层、流和许多其他项目的信息不会存活很久；它经常被用户活动所废止。

任何修改项目数量(而不是质量)的行为都会使其对这些项目的引用失效；向流添加关键帧会使对该流的引用失效，但强制渲染一个图层并不会使对它的引用失效。不要缓存图层像素。

不建议在调用插件中的特定钩子函数之间缓存引用；当你需要时获取信息，并尽快忘记(释放)它。

## 你打算乱放数据吗？

当你要求 After Effects 填充和返回数据结构的句柄时，重要的是你要自己清理。对于以下数据类型，必须调用相应的处理程序。

## Data Types Requiring Disposal

| **Data Type**  | **Disposal function**  |
| --- | --- |
| `AEGP_Collection2H`  | `AEGP_DisposeCollection`, 来自于[AEGP_CollectionSuite2](../aegps/aegp-suites.html)  |
| `AEGP_FootageH`  | `AEGP_DisposeFootage`, 来自于[AEGP_FootageSuite5](../aegps/aegp-suites.html) |
| `AEGP_WorldH`  | `AEGP_Dispose`, 来自于 [AEGP_WorldSuite3](../aegps/aegp-suites.html) 或 `AEGP_DisposeTexture`, 来自于 [AEGP_CanvasSuite8](../artisans/artisan-data-types.html), if layer texture created using AEGP_RenderTexture`) |
| `AEGP_EffectRefH`  | `AEGP_DisposeEffect`, 来自于[AEGP_EffectSuite4](../aegps/aegp-suites.html)  |
| `AEGP_MaskRefH`  | `AEGP_DisposeMask`, 来自于[AEGP_MaskSuite6](../aegps/aegp-suites.html)  |
| `AEGP_RenderOptionsH`  | `AEGP_Dispose`, 来自于[AEGP_RenderQueueMonitorSuite1](../aegps/aegp-suites.html)  |
| `AEGP_LayerRenderOptionsH` | `AEGP_Dispose`, 来自于[AEGP_LayerRenderOptionsSuite1](../aegps/aegp-suites.html)  |
| `AEGP_RenderReceiptH`  | `AEGP_DisposeRenderReceipt`, 来自于[AEGP_CanvasSuite8](../artisans/artisan-data-types.html) |
