---
title: Interaction-Callback-Functions
order: 11
---

# 交互回调函数

虽然在[PF_InData](../effect-basics/PF_InData.html)中提供了未使用宏的函数指针，但请使用所提供的宏来访问它们。看看看我们对弃用宏有多严格？这是我们的小秘密哦。

## 交互回调

### PF_ADD_PARAM

在[PF_Cmd_PARAM_SETUP](../effect-basics/command-selectors.html)期间枚举插件参数给 After Effects，对该函数进行多次调用。。

注意：在 `PF_ADD_PARAM()`之前没有清除 `PF_ParamDef` 会出问题。因此在添加参数之前一定要使用 `AEFX_CLR_STRUCT`。

```cpp
PF_ErrPF_ADD_PARAM(
PF_InData*in_data,
PF_ParamIndexindex,
PF_ParamDefPtrdef);
```

我们在Utils/Param_Utils.h中为特定的参数类型提供了方便的宏：

> - `PF_ADD_COLOR`,
> - `PF_ADD_ARBITRARY`,
> - `PF_ADD_SLIDER`,
> - `PF_ADD_FIXED`,
> - `PF_ADD_FLOAT_SLIDERX`,
> - `PF_ADD_CHECKBOXX`,
> - `PF_ADD_BUTTON`,
> - `PF_ADD_ANGLE`,
> - `PF_ADD_NULL`,
> - `PF_ADD_LAYER`,
> - `PF_ADD_255_SLIDER`,
> - `PF_ADD_PERCENT`,
> - `PF_ADD_POINT`,
> - `PF_ADD_POINT_3D`,
> - `PF_ADD_TOPICX`,
> - `PF_END_TOPIC`,
> - `PF_ADD_POPUPX`,
> - `PF_ADD_FLOAT_SLIDERX_DISABLED`

### PF_ABORT

如果用户取消，返回非零值；并将该值返回给 After Effects。将你的渲染程序包裹在一个 "没有中止请求" 的 while 循环中。

```cpp
PF_ErrPF_ABORT(PF_InData*in_data);
```

### PF_PROGRESS

在处理过程中显示一个进度条；当前完成值和总数的百分比。如果暂停或中止当前处理，则返回非零值；将该值返回给 After Effects。每个扫描线(scanline)调用一次，除非你的效果非常慢。

如果总数为 0，则用 PF_ABORT 代替(给用户提供不同的选择)。

```cpp
PF_ErrPF_PROGRESS(
PF_InData*in_data,
A_longcurrent,
A_longtotal);
```

### PF_CHECKOUT_PARAM

在指定时间获取参数值，或源视频层。After Effects 会根据参数的签出状态做出缓存决定。

分配一个新的[PF_ParamDef](../effect-basics/PF_ParamDef.html)来保存结果；传递给插件的是只读的。如果你签出的图层参数被设置为`<none>`，返回的图层将被零(zeros)填充。蒙版不包括在签出图层中。

在 UI 事件处理过程中不要签出图层参数。

```cpp
PF_ErrPF_CHECKOUT_PARAM(
PF_InData*in_data,
PF_ParamIndexindex,
A_longwhat_time,
A_longstep,
A_longtime_scale,
PF_ParamDef*param);
```

如果签出源图层，将返回一个去隔行帧。如果你要求引用上层字段(upper field)的时间，将收到带有用于生成额外扫描线的过滤器的上层字段。例如，假设第 0 行和第 2 行是上场，第 1 行是下场，如果你签出上场，第 0 行和第 2 行将直接从源素材传回，第1行将通过第0行和第2行平均计算。如果您想在两个字段都存在的情况下重新组装全分辨率源帧, 可以调用 PF_CHECKOUT_PARAM 两次来获得两个字段，然后重新交错素材。

在帧不对齐的时间签出图层时会发生什么？所有的项目基本上都有无限的时间分辨率，所有项目本质上都有无限的时间分辨率，因此当以任何值询问时间时，AE会在该时间渲染项目。对于一个合成，将涉及到所有关键帧值插值到子帧时间。对于素材，AE返回与查询时间相对应的完整图像，即最接近左的帧。如果用户在该图层上进行帧混合，则会生成插值帧。

### PF_CHECKIN_PARAM

请使用`PF_CHECKIN_PARAM`来平衡每一个`PF_CHECKOUT_PARAM`。

否则会导致糟糕的性能和泄漏内存。一旦签入，[PF_ParamDef](../effect-basics/PF_ParamDef.html)中的字段将不再有效。

```cpp
PF_ErrPF_CHECKIN_PARAM(
PF_InData*in_data,
PF_ParamDef*param);
```

### PF_REGISTER_UI

注册一个自定义的用户界面元素。见[Effect UI & Events](../effect-ui-events/effect-ui-events.html)。注意：PF_UIAlignment 的标志不被认可。

```cpp
PF_ErrPF_REGISTER_UI(
PF_InData*in_data,
PF_CustomUIInfo*cust_info);
```

### PF_CHECKOUT_LAYER_AUDIO

给定一个索引、start_time、duration、time_scale、rate、byte_per_sample、num_channels 和 fmt_signed，After Effects 将返回一个相应的 PF_LayerAudio。After Effects 将执行任何必要的重采样。

```cpp
PF_ErrPF_CHECKOUT_LAYER_AUDIO(
PF_InData*in_data,
PF_ParamIndexindex,
A_longstart_time,
A_longduration,
A_u_longtime_scale,
PF_UFixedrate,
A_longbytes_per_sample,
A_longnum_channels,
A_longfmt_signed,
PF_LayerAudio*audio);
```

### PF_CHECKIN_LAYER_AUDIO

平衡所有对 PF_CHECKOUT_LAYER_AUDIO 的调用，无论错误情况如何，都要匹配对 PF_CHECKIN_LAYER_AUDIO 的调用。

```cpp
PF_ErrPF_CHECKIN_LAYER_AUDIO(
PF_InData*in_data,
PF_LayerAudioaudio);
```

### PF_GET_AUDIO_DATA

返回关于 PF_LayerAudio 的信息。

音频之后的所有参数都是可选的；用不到参数的请传 0。rate0 是无符号的，fmt_signed0 是有符号的非零，零是无符号的。此回调用于读取音频信息的视觉效果。要更改音频，请编写音频过滤器(filter)。

```cpp
PF_ErrPF_GET_AUDIO_DATA(
PF_InData*in_data,
PF_LayerAudioaudio,
PF_SndSamplePtr*data0,
A_long*num_samples0,
PF_UFixed*rate0,
A_long*bytes_per_sample0,
A_long*num_channels0,
A_long*fmt_signed0);
```

## Parameter 签出 vs Param Zero

在效果控制(和合成)面板中，效果是按照从 0 到 n 的顺序应用于图像的。

第n-1效果的输出是第n个效果的输入([param[0]](../effect-basics/PF_ParamDef.html) )。

另一方面，当一个普通的效果使用 "PF_CHECKOUT_PARAM "签出一个图层时，不管顺序如何, 始终接收原始(未受影响的)源图层，。

然而，当[SmartFX](../smartfx/smartfx.html)效果检出其输入参数(params[0])时，先前的效果会被应用。

## Parameter 签出行为

无论图层进出点是否被修剪，都会得到从源素材开始到结束的有效帧，然后在这之前和之后都是透明的。

帧率低于所签出的合成帧的图层参数，则仅在低帧率需要的时候刷新。

一个10fps图层在一个30fps合成签出, 只需要每三帧刷新一次.如果效果想在每一帧改变它的输出，尽管输入层是静态的，也需要设置[PF_Outflag_NON_PARAM_VARY](../effect-basics/PF_OutData.html)。

当效果检出连续光栅化的 Adobe Illustrator 图层时，After Effects 会在一个合成大小的缓冲区中渲染带有几何图形的 Illustrator 图层。

## Parameter 签出 And Re-Entrancy(重入)

在不同时间签出图层的插件会产生重入行为。比如：签出示例插件应用于合成B中的图层，并且B被预合成为A，其中签出也应用于它。

当合成 A 被渲染时，`Checkout[A]` 将被发送*PF_Cmd_RENDER*，在此期间，在此期间，它从当前时间以外的时间签出图层（合成B）。

为了提供签出图层，AE将发送*PF_Cmd_RENDER*到`Checkout[B]`。

很快，递归！

如果您要签出参数，插件必须适当地处理可重入渲染请求。

不要使用全局变量，也不要读取或写入静态变量
......但你反正也不打算这么做，对吗？

## 迭代过程中的进度

AE努力尽可能地响应用户交互，即使在渲染时也是如此。通过适当使用PF_ITERATE()也可以做到这一点。例如在对`PF_Cmd_RENDER`的响应过程中，使用了三次 `PF_ITERATE`函数。

在这种情况下，应该先用:

```cpp
lines_per_iterateL = in_data>extent_hint.top - in_data>extent_hint.bottom;
total_linesL = 3 * lines_per_iterateL;
lines_so_farL = 0;

```

每次迭代后，将已完成的行添加到当前位置：

```cpp
suites.iterate8suite()>iterate( lines_so_farL,
 total_linesL,
 input_worldP,
 &output>extent_hint,
 refcon,
 WhizBangPreProcessFun,
 output_worldP);

lines_so_farL += lines_per_iterateL;

ERR(PF_PROGRESS(lines_so_farL, total_linesL));

suites.iterate8suite()>iterate( lines_so_farL,
 total_linesL,
 input_worldP,
 &output>extent_hint,
 refcon,
 WhizBangRenderFunc,
 output_worldP);

lines_so_far += lines_per_iterateL;

ERR(PF_PROGRESS(lines_so_farL, total_linesL));

suites.iterate8suite()>iterate( lines_so_farL,
 total_linesL,
 input_worldP,
 &output>extent_hint,
 refcon,
 WhizBangPostProcessFunc,
 output_worldP);

ERR(PF_PROGRESS(lines_so_farL, total_linesL));

```
