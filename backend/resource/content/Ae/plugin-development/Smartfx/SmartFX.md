---
title: SmartFX
order: 1
---

# SmartFX

SmartFX(智能插件) API 提供效果和 AE 之间的双向通信，实现了许多性能优化，并提供了以前无法获得的依赖信息。这种效果 API 的扩展是在 After Effects 中实现每通道 32 位支持的方式。

普通的效果插件被赋予一个全尺寸的输入缓冲区，并要求渲染一个全尺寸的输出缓冲区。虽然输出[extent_hint](../effect-basics/PF_InData.html) 指定了必须实际填充的输出缓冲区的部分，但如果效果不需要其整个输入，这种方案仍然是非常低效。另外，许多效果并不使用范围提示。

## 没有智能插件前

比如应用于一个超大图层的模糊效果，该图层大部分在屏幕外，或者只通一小块目标区域(region of interest)查看，或者使用蒙版裁剪到一个小尺寸。只需要渲染输出一小部分，用输出的 extent_hint 来表示效果。也只需要对一小部分输入区域进行模糊处理 --- 输出的 extent_hint 通过模糊半径扩大。然而，使用传统的特效 API，After Effects 没有办法知晓这一点，所以只会传递给插件整个图层。计算这些多余像素非常浪费，特别是在先前的效果或嵌套合成的情况下。

## 有了智能插件后

通过颠倒调用顺序, SmartFX 解决了这个问题。效果可以知道需要多少输出，必须明确地向主机请求所需要的输入。渲染过程被分成两部分：预渲染和渲染。

在预渲染过程中，效果只获取需要的输入像素数据；输入可以根据你喜欢的任何东西(非输入图层参数、非图层参数、来自 in_data 的信息、序列数据中的设置...)变化。效果还必须返回结果输出的范围，如果请求图层有部分空像素，这个范围可能小于请求的大小。

在渲染阶段，效果只检索之前请求的像素。这种两次判定的方法有利于优化。例如，将一个输入与另一个输入相乘或遮罩可能会发现，根本就不需要第一个输入。

如果遮罩不与它相交，那么才需要。After Effects 内部还有一些重要的优化，以确保尽可能少地复制图像缓冲区，而这些优化只有在主机知道缓冲区大小以及所有输入和输出才可以做到。

与 AEGPs 一样，SmartFX 插件永远不会被 After Effects 卸载。

## 内容边界

节点的内容边界是调用预渲染所能返回的最大可能的结果矩形。它绝对不能根据当前的渲染请求或其他情况而变化。应该被仔细计算，而不是松散地计算。

这个计算是非常重要的。它是节点(和它的输入)的内在属性，一旦图被建立，它就被固定了。违反它可以而且可能会在各种代码中造成各种问题。

## 如何智能

设置了 "PF_OutFlag2_SUPPORTS_SMART_RENDER" ( 参见[PF_OutFlags](../effect-basics/PF_OutData.html) 的效果将收到 SmartFX 的调用 "PF_Cmd_SMART_PRE_RENDER "和 "PF_Cmd_SMART_RENDER" ( 参见[帧入口智能](../effect-basics/command-selectors.html))，而不是以前的`PF_Cmd_FRAME_SETUP`/`PF_Cmd_RENDER`/`PF_Cmd_FRAME_SETDOWN`序列。为了保持与非智能主机的兼容性，你可能也想继续支持旧的命令。

## PF_Cmd_SMART_PRE_RENDER

AE 请求效果的输出。效果通过使用回调函数和操纵额外参数中的结构来告诉 AE 它需要什么输入来生成输出。效果不能访问它在*PF_Cmd_SMART_PRE_RENDER*期间未签出的任何图层输入的像素。因此，效果可能需要的所有图层输入都必须使用 checkout_layer 提前签出。如果效果可能需要某些图层输入，则必须现在签出，即使稍后在渲染过程中效果可能会决定不需要图层。另外，由于在 PF_Cmd_SMART_PRE_RENDER\*或`PF_Cmd_SMART_RENDER`过程中，没有参数数组传递给 SmartFX，因此必须使用`PF_CHECKOUT_PARAM`来获取 ( 参见[交互回调](../effect-details/interaction-callback-functions.html)检索所需的任何非图层参数

## PF_PreRenderExtra

### PF_PreRenderInput

描述了 After Effects 需要渲染的内容(在 `PF_RenderRequest`中)，以及要求的比特深度(在命名恰当的 bitdepth 成员中)。

```cpp
typedef struct {
  PF_LRect        rect;
  PF_Field        field;
  PF_ChannelMask  channel_mask;
  PF_Boolean      preserve_rgb_of_zero_alpha;
  char            unused[3];
  long            reserved[4];
} PF_RenderRequest;
```

`rect`是图层坐标，字段也是相对于图层原点的；活动字段是落在输出缓冲区的偶数还是奇数扫描线上，取决于输出缓冲区的原点。

`channel_mask`指定效果应该为哪些通道提供输出。

写入其他通道的数据将不被接受。

以下的一个或多个。

> - `PF_ChannelMask_ALPH`
> - `PF_ChannelMask_RED`
> - `PF_ChannelMask_GREE`
> - `PF_ChannelMask_BLUE`
> - `PF_ChannelMask_ARGB`

如果 `preserve_rgb_of_zero_alpha`像素为 TRUE`，效果必须将透明像素的颜色内容传递到输出。

这与[PF_OutFlag2_REVEALS_ZERO_ALPHA](../effect-basics/PF_OutData.html)有关，但与之不同，后者告诉 AE,效果可能会将这些像素的 alpha 设置为非零值，从而使它们恢复可见性。

### PF_PreRenderOutput

由效果填充，告诉 AE 它计划根据输入生成什么输出。

```cpp
typedef struct {
  PF_LRect                    result_rect;
  PF_LRect                    max_result_rect;
  PF_Boolean                  solid;
  PF_Boolean                  reserved;
  PF_RenderOutputFlags        flags;
  void*                       pre_render_data;
  PF_DeletePreRenderDataFunc  func;
} PF_PreRenderOutput;
```

`pre_render_data`将在[PF_Cmd_SMART_RENDER](../effect-basics/command-selectors.html)时被传回给效果。

目前，唯一的 `PF_RenderOutputFlags`是`PF_RenderOutputFlag_RETURNS_EXTRA_PIXELS`。

### PF_PreRenderCallbacks

目前，只有一个回调 - `checkout_layer`。checkout_idL 由效果选择。

它必须为正数且唯一。After Effects 会填充 `PF_CheckoutResult`。

```cpp
PF_Err checkout_layer(
  PF_ProgPtr              effect_ref,
  PF_ParamIndex           index,
  A_long                  checkout_idL,
  const PF_RenderRequest  *req,
  A_long                  what_time,
  A_long                  time_step,
  A_u_long                time_scale,
PF_CheckoutResult       *result);

typedef struct {
  PF_LRect          result_rect;
  PF_LRect          max_result_rect;
  PF_RationalScale  par;
  long              solid;
  PF_Boolean        reservedB[3];
  A_long            ref_width;
  A_long            ref_height;
} PF_CheckoutResult;
```

`result_rect`可以为空。`max_result_rect`是最大的输出可能，如果主机要求所有的输出。如果 solid 为 TRUE，整个 result_rect 有不透明的 alpha。

`ref_width`和`ref_height`是该图层的原始尺寸，在应用任何效果之前，不考虑任何降频采样因素。这将是折叠图层的合成尺寸。

当一个 SmartFX 插件同时使用 [PF_OutFlag2_AUTOMATIC_WIDE_TIME_INPUT](../effect-basics/PF_OutData.html) 和 [PF_OutFlag_NON_PARAM_VARY](../effect-basics/PF_OutData.html)时, 在 11.0 回存在一个全局性能缓存错误.

在`PF_Cmd_SMART_PRE_RENDER`期间调用`checkout_layer`返回`PF_CheckoutResult`中的空矩形。

解决方法是简单地再次调用。

已在 11.0.1 修复.

### result_rect

渲染请求产生的输出（图层坐标）（可以为空）。

不能大于输入请求矩形（除非设置了`PF_RenderOutputFlag_RETURNS_EXTRA_PIXELS`），但可以比它小。

### max_result_rect

如果 AE 请求所有输出，则输出可能的最大尺寸。

不会因请求的输出尺寸而异。

### solid

如果输出中的每个像素都是完全不透明的，请设置此 TRUE。如果可能，请设置；它会启用某些优化。

### reserved

忽略

### flags

目前，唯一的标志是`PF_RenderOutputFlag_RETURNS_EXTRA_PIXELS`，它告诉 AE 智能效果将返回比 AE 请求更多的像素。

### pre_render_data

将其指向效果在渲染期间想要访问的任何数据。

效果还可以分配句柄并将它们存储在 `out_data>frame_data`, 就像常规（非智能）效果一样。

自从 [PF_Cmd_SMART_PRE_RENDER](../effect-basics/command-selectors.html) 可以在没有对应 [PF_Cmd_SMART_RENDER](../effect-basics/command-selectors.html)的情况下调用,效果不得自行删除此数据;

一旦效果从 [PF_Cmd_SMART_PRE_RENDER](../effect-basics/command-selectors.html)返回, AE 拥有这些数据并将处理它（使用以下函数或标准释放调用）。

### delete_pre_render_data_func

指向一个函数,用于删除 pre_render_data。

## preserve_rgb_of_zero_alpha

`preserve_rgb_of_zero_alpha` 既是效果的输入，用来告诉它要渲染什么，也是效果的输出，用来描述它需要的输入（传递给签出调用）。当 preserve_rgb_of_zero_alpha 在输入请求中设置时，效果必须在签出时递归传递它，否则先前的效果和掩蔽将消除效果将显示的像素。不鼓励使用这种方法，尽管 CS3（8.0）仍然支持。

## Rectangles

效果必须准确设置两个结果矩形。AE 的缓存系统依赖于它们，不正确的值会导致许多问题。如果插件返回的 `result_rect` 小于`request_rect`，这将告诉 AE 像素再`request_rect`内 但`result_rect`外的像素为空。

同样，`max_result_rect` 必须包含所有非零像素；效果永远不会被要求渲染该区域之外的任何内容。如果该矩形之外有像素，它们将永远不会显示。

错误大小的输出矩形也会导致问题。如果这些矩形太大，会导致性能损失。

不仅会缓存许多空像素（抢走应用程序的宝贵内存），效果可能会被不必要地要求渲染无内容的大区域。因此，必须正确计算 `max_result_rect`,而不是随便设置

`result_rect` 和 `max_result_rect` 都可能因效果的参数、当前时间等而异；它们仅对给定的效果调用有效。但是，`max_result_rect`不能取决于特定的渲染请求。无论 AE 请求输出的哪一部分，它都必须相同。

如果 `request_rect` 不与效果的输出像素相交，返回`result_rect`为空是合法的；不需要进行渲染。

After Effects 也可以用一个空的 `request_rect` 来调用效果，这意味着效果只计算`max_result_rect`。

`preserve_rgb_of_zero_alpha`可以影响边界的计算过程(包括 result_rect 和`max_result_rect`)，如果效果的行为因这个设置而不同，则必须尊重这个设置。

## 图层的"尺寸"

与非智能效果一样，每个智能效果都可以任意收缩或扩展其请求的输入。它们不能依赖于一个固定的框架大小，而且输入的大小可能会随着时间的推移而改变。

例如，用户可以给一个图层应用一个动画的阴影，这将在不同的时间给图层的不同边缘增加像素，这取决于阴影投射的方向。

有些效果(例如，那些需要将一个图层与另一个图层对齐的效果)需要一些 "大小 "的概念。这可以用两种方式来定义，每一种都有优点和缺点。

在应用任何效果和降采样之前，原始图层的尺寸被赋予`in_data>width/height`。由于这个数值不受后续效果的影响，它可以作为中心点等事项的绝对参考。

然而，这并不是万无一失的，因为用户可能已经应用了失真或平移效果。而且，这个值只对应用效果的图层可用，对其他图层的参数无效。

…or…

每个图层输入都有一个`max_result_rect`，它包含了所有的像素数据，在某种意义上是一个图层的主 "尺寸"。

它对所有图层都是可用的，但会随着时间的推移根据以前应用的效果而改变，可能是用户没有想到的方式(如上面的水滴阴影例子)。

请注意，输入的 Ref_width/height 和`max_result_rect`可以在不渲染的情况下获得，方法是用空的`request_rect`调用`checkout_layer`。

这是相当有效的，如果首先需要图层的 "尺寸"，以确定渲染所需的确切像素，这将是非常有用的。

这是一个在 pre-render 中请求一个图层，然后从不调用`checkout_layer`的例子(在这种情况下，没有)。

## Flag On The Play

通常情况下，一个给定的`PF_RenderRequest`的`max_result_rect`将被裁剪到任何应用的遮罩的边界。

然而，如果[PF_OutFlag2_REVEALS_ZERO_ALPHA](../effect-basics/PF_OutData.html) 被设置，`max_result_rect`将是该图层的尺寸。

## PF_Cmd_SMART_RENDER

该效果在每次预渲染时最多会收到一次 *PF_Cmd_SMART_RENDER* 调用。

注意，渲染可能根本就不会被调用。After Effects 可能只想进行一些边界计算，或者它后来发现根本不需要效果的输出(例如，如果轨道哑光的预渲染阶段返回一个不与效果的输出相交的矩形，就可能发生这种情况)。

所有的效果都必须能够处理无渲染的预渲染，而不泄漏资源或进入不稳定的状态。

在*PF_Cmd_SMART_RENDER*期间，额外参数指向 PF_SmartRenderExtra。

## PF_SmartRenderExtra

### PF_SmartRenderInput

由一个[PF_RenderRequest](../smartfx/smartfx.html)、位深度和一个指向 `pre_render_data`的指针(在[PF_Cmd_SMART_PRE_RENDER](../effect-basics/command-selectors.html)期间分配)合成。

这个 `PF_SmartRenderInput`与相应的*PF_Cmd_SMART_PRE_RENDER*中传递的相同。

### PF_SmartRenderCallbacks

```cpp

PF_Err checkout_layer_pixels(
  PF_ProgPtr      effect_ref,
  A_long          checkout_idL,
  PF_EffectWorld  **pixels);
```

这是用来实际访问在*PF_Cmd_SMART_PRE_RENDER*过程中检查出来的图层中的像素。

返回的 PF_EffectWorld`在当前命令的持续时间内有效，或者直到检出。

你只能用*PF_Cmd_SMART_PRERENDER*中的 checkout_idL 调用 `checkout_layer_pixels`一次。在*PF_Cmd_SMART_PRERENDER*和*PF_Cmd_SMART_RENDER*中进行的签出数量之间必须有一个一对一的映射。

要在一个图层上多次调用 `checkout_layer_pixels`，应该在*PF_Cmd_SMART_PRERENDER*中用不同的唯一 `checkout_idL`在同一图层上再次调用[checkout_layer](../smartfx/smartfx.html)，然后用这个`checkout_idL`在*PF_Cmd_SMART_RENDER*中做另一个 `checkout_layer_pixels`。

```cpp
PF_Err checkin_layer_pixels(
  PF_ProgPtr  effect_ref,
  A_long      checkout_idL);
```

没有必要调用(当效果从*PF_Cmd_SMART_RENDER*返回时，After Effects 会清理所有这样的检出)，但对释放内存很有用。

```cpp
PF_Err checkout_output(
  PF_ProgPtr      effect_ref,
  PF_EffectWorld  **output);
```

检索输出缓冲区。注意，在至少一个输入被检出之前，效果是不允许检出输出的(除非效果根本没有输入)。

注意：为了优化内存的使用，尽可能晚地请求输出，并尽可能少地一次请求输入。

## 何时访问图层参数

除了图层输入以外的参数可以在任何时候自由检出。图层输入必须在:ref:PF_Cmd_SMART_PRE_RENDER <effect-basics/command-selectors.frame-selectors>期间访问。

然而，你并不需要实际使用每个输入。

如果你在[PF_Cmd_SMART_PRE_RENDER](../effect-basics/command-selectors.html) 中检出一个帧(或其一部分)，随后不在`PF_Cmd_SMART_RENDER`中检出，它就不需要被渲染，大大改善性能。

## Wait, Gimme That Layer Back

`checkout_layer_pixels`只能被调用一次，并使用 PreRender 中的 `checkout_id`。在 PreRender 和 SmartRender 中的签出数量必须有一个一对一的映射。如果你需要对一个图层的像素进行多次检查，也许是因为你的代码结构，只需使用一个以上的 `checkout_id`。在 PreRender 中，在同一个图层上用不同的唯一 checkout_id 调用 checkout_layer。然后在 SmartRender 中，每次在 SmartRender 中调用 `checkout_layer_pixels` 时，使用这些 `checkout_ids` 中不同的一个。
