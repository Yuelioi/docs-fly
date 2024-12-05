---
title: Global-Sequence-Frame-Data
order: 8
---

# 全局、序列和帧数据

After Effects 允许插件在三个范围内存储数据：全局、序列和帧。仔细考虑你存储信息的位置；选择不当会影响性能，或使你的插件对用户产生混淆。

使用全局数据来存储效果的所有实例共有的信息：静态变量和数据、位图、指向其他 DLLs 或外部应用程序的指针。如果你的效果支持多帧渲染，任何静态变量或全局变量都必须没有竞争关系(更多信息请参见[什么是效果的线程安全？](multi-frame-rendering-in-ae.html) )。

在序列数据或新的[多帧渲染计算缓存](multi-frame-rendering-in-ae.html)中存储任何仅限你插件实例的信息(UI 设置、文本字符串和任何不存储在参数中的自定义数据)。

帧数据用于渲染一个特定的帧的信息。现已弃用，因为大多数机器都能够一次将整个帧加载到内存中。当然，你的 IMAX 生成用户仍然会感谢你所做的任何优化。

## 持久性

AE将序列数据保存在项目文件中，但不保存全局或帧数据。序列数据中指向外部数据的指针，在重新打开项目时很可能是无效的，必须重新连接。我们把这个过程称为 "展平化"和 "非展平化 "序列数据。(“flattening” and “unflattening” the sequence data)

::: tip

计算缓存并不将其内容存储到项目文件中。存储在缓存中的数据必须在渲染时重新创建。

:::

## 验证序列数据

仔细的序列数据验证对于跨时间模拟的效果很重要，在这种情况下，帧 N 依赖于帧 N-1，并且您在序列数据中使用计算数据缓存。如果参数发生更改，某些计算数据可能不再有效，但每次更改后盲目重新计算也是浪费。

当被要求渲染第 N 帧时，假设你的缓存数据已经计算到了第 N-1 帧，调用`PF_GetCurrentState()/PF_AreStatesIdentical()`中的[PF_ParamUtilSuite3](Parameter-Supervision.html)，查看当前参数设置下缓存的计算数据是否仍然有效。

在经过的时间跨度内检查所有参数(除了设置[PF_ParamFlag_EXCLUDE_FROM_HAVE_INPUTS_CHANGED](../effect-basics/PF_ParamDef.html))，包括图层参数(包括[param[0]](../effect-basics/PF_ParamDef.html))。

这可以有效地完成，因为更改跟踪是使用时间戳完成的。

如果输入没有变化，你可以安全地使用缓存，而且内部缓存系统会认为你对传递的范围有时间上的依赖。因此，如果上游有什么变化，主机缓存将自动失效。

验证是否有效，可以在每帧上用参数关键帧应用你的效果。用 RAM 预览来填充缓存，然后改变其中一个关键帧。相关的帧和所有从属的帧(例如，在模拟的情况下，后面的关键帧)应该失去它们的缓存标记，需要重新渲染。同样，对图层参数来源的上游改变应该导致缓存的时间选择性失效。

## 展平化与非展平化序列数据

如果你的序列数据引用了外部内存(在指针或句柄中)，你必须对你的数据进行展平化与非展平化(Flattened And Unflattened)，以实现磁盘安全存储。这类似于创建自己的微型文件格式。

在收到[PF_Cmd_SEQUENCE_FLATTEN](../effect-basics/command-selectors.html)后，将指针引用的数据放入一个连续的块中，稍后您可以从中恢复旧结构。

如果你的序列数据包含一个指向 long 的指针，请分配4个字节来存储扁展平化的数据。您必须处理特定于平台的字节排序。

记住，你的用户(反正是买了两份插件的人)可能希望一个项目能同时在 macOS 和 Windows 上运行。

对于展平化与非展平化数据, After Effects 在重新加载数据时发送[PF_Cmd_SEQUENCE_RESETUP](../effect-basics/command-selectors.html)。

在两个结构中使用一个共同偏移量的标志来表示数据的状态。

```cpp
typedef struct {
 A_char* messageZ;
 PF_FpLong big_numF;
 void* temp_storage;
} non_flat_data;

typedef struct {
 char message[256];
 PF_FpLong big_numF;
 A_Boolean big_endianB;
} flat_data;

```

## 调整序列数据大小

在[PF_Cmd_SEQUENCE_SETUP](../effect-basics/command-selectors.html) 期间，为特定于此效果实例的数据分配一个句柄。

你可以在任何入口指令中修改序列数据的内容，但不能修改其大小。

你可以只在以下入口指令中调整序列数据句柄的大小。

> - `PF_Cmd_AUDIO_SETUP`
> - `PF_Cmd_AUDIO_SETDOWN`
> - `PF_Cmd_FRAME_SETUP`
> - `PF_Cmd_FRAME_SETDOWN`
> - `PF_Cmd_AUDIO_RENDER`
> - `PF_Cmd_RENDER`
> - `PF_Cmd_SEQUENCE_SETUP`
> - `PF_Cmd_SEQUENCE_SETDOWN`
> - `PF_Cmd_SEQUENCE_FLATTEN`
> - `PF_Cmd_SEQUENCE_RESETUP`
> - `PF_Cmd_DO_DIALOG`

## 使用多帧渲染在渲染时间访问sequence_data

当在一个效果上启用多帧渲染时，`sequence_data`对象在渲染过程中是只读/常量的，并且可以通过`PF_EffectSequenceDataSuite1`套件在每个渲染线程上访问。

### PF_EffectSequenceDataSuite1

#### PF_GetConstSequenceData

当一个效果启用多帧渲染时，为一个渲染线程检索只读的 const_data 对象。

```cpp
PF_Err(*PF_GetConstSequenceData)(
PF_ProgPtreffect_ref,
PF_ConstHandle*sequence_data);
```

```cpp
static PF_Err Render(
   PF_InData   *in_dataP,
   PF_OutData  *out_dataP,
   PF_ParamDef *params[],
   PF_LayerDef *output )
{
    PF_ConstHandle seq_handle;

    AEFX_SuiteScoper<PF_EffectSequenceDataSuite1> seqdata_suite =
      AEFX_SuiteScoper<PF_EffectSequenceDataSuite1>(
        in_dataP,
        kPFEffectSequenceDataSuite,
        kPFEffectSequenceDataSuiteVersion1,
        out_dataP);

    PF_ConstHandle const_seq;
    seqdata_suite->PF_GetConstSequenceData(in_data->effect_ref, &const_seq);

    // cast const_seq to the type used when storing to sequence_data

    // rest of render function code...
}
```
