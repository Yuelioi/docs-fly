---
title: PF_InData
order: 6
---

# PF_InData

After Effects 使用`PF_InData`与系统、项目、图层和音频通信。在插件接收入口指令之前更新此结构体。

注册只在特定的[PF_Cmds](command-selectors.html)中有效的字段(field)。

另外，不用担心；虽然`PF_InData`大得吓人，但不需要记住每个成员的用途；只会在某些时候使用一些字段。

## PF_InData 成员

| **名称**  | **描述**  |
| --- | ------------------------- |
| `inter`  | 除了当前帧被渲染了，这个回调会被用来做用户交互，添加参数，检查用户是否删除效果，显示进度条，获取源帧和参数值。<br />是 [Interaction Callback Functions](../effect-details/interaction-callback-functions.html)中一个十分有用的函数包。. |
| `utils`  | 图像和数学计算的回调。这个指针会一直存在。  |
| `effect_ref`  | 这是个内容不公开的数据，被用入大量的回调函数<br />Ae依靠这个来识别你的插件。  |
| `quality`  | 当前质量设定, 是 _PF_Quality_HI_ 或者 _PF_Quality_LO_ <br />效果在LO模式下算的快，在HI模式下算的精准。<br />图形工具包在LO和HI模式下不同。你的插件也应如此<br />在所有帧和序列入口指令中都定义该字段。  |
| `version`  | 插件版本,在 `PF_Cmd_GLOBAL_SETUP`的时候申明它  |
| `serial_num`  | 调用程序的序号  |
| `appl_id`  | 调用程序的标识符。<br />‘FXTC’ 代表AE <br />如果是 [Premiere Pro 和其他主机](../ppro/ppro.html)则是‘PrMr’<br />使用它来测试您的插件（用于一个应用程序的许可）是否正在与另一个应用程序一起使用。  |
| `num_params`  | 输入参数数量 |
| `what_cpu`  | 在 macOS 下，它包含 CPU 类型的 Gestalt 值（参见 Inside Macintosh，第 6 卷）。在 Windows 上未定义。  |
| `what_fpu`  | 在 macOS 下，它包含FPU类型的 Gestalt值<br />在 Windows 上未定义。  |
| `current_time`  | 当前帧被渲染的时间。在 [PF_Cmd_RENDER](command-selectors.html) 期间有效 <br />这是图层中的当前时间，而不是合成时间。如果图层不是从0开始或者被拉伸，那么图层时间和合成时间不同 <br />当前帧号可以用 `current_time` 除以 `time_step`<br />用秒来表示的当前时间可以用`current_time`除以`time_scale`<br />为了解决时间拉伸压缩，合成的帧率改变以及时间重置。AE会要求插件渲染一个不完整的时间（在两帧之间）<br />做好心理准备，不要以为只会在帧上。
PS：CS3，插件可能会被要求渲染负数帧  |
| `time_step`  | 当前渲染帧持续的时长<br />在合成嵌套中，该帧时长与这一图层的帧间间隔`local_time_step`不同<br />将这个变量除以`time_scale`可以转化为以秒表示的形式<br />当计算其他帧的时长时，比如 [PF_CHECKOUT_PARAM](../effect-details/interaction-callback-functions.html), 应使用这个值而不是`local_time_step`<br />在图层时间反转(time-reversed)时该值会成为负数. 并且在对嵌套合成进行时间重映射(time remapping)时可以随帧的改变而改变。<br />当原素材在嵌套合成里被拉伸或时间重置时与`local_time_step`不同。<br />例如当一个合成被嵌套在一个帧率不同的合成里或者外面的那一个合成上有时间重置效果时，两者就会不同。<br />若每一帧的`time_step`不是定值，那么在[PF_Cmd_SEQUENCE_SETUP](command-selectors.html)的值就为0。<br />在`PF_Cmd_FRAME_SETUP` and `PF_Cmd_FRAME_SETDOWN` 时可以正确设置这个变量。<br />警告：因为该值可以为0，将它作为除数之前一定要再三确认！ |
| `total_time`  | 图层的持续时间<br />如果该图层经过时间拉伸，长度超过100%，这个变量也会相应地被改变，如果经过时间拉伸，长度小于100%，这个变量不会受影响。<br />如果当前图层被时间重映射过，那么这个变量等于其持续时长(duration)！<br />将本变量除以`time_scale`，可以将其转换为以秒为单位的变量。  |
| `local_time_step`  | 该层两帧之间相差的时间。<br />任何的时间拉伸都会影响这个值。<br />当图层时间反转时，可以为负值。<br />与`time_step`不同的是，这个变量对于任何两帧之间都是恒定不变的。<br />除以`time_scale`将本参数转换为秒。<br />如果整个图层每一帧的步值(step value)都相同，请使用建立在合成帧率和时间拉伸基础上进行计算的`local_time_step`。  |
| `time_scale`  | `current_time`, `time_step`, `local_time_step` 以及 `total_time` 的单位<br />如果该值为30，那么 `current_time`, `time_step`, `local_time_step`和`total_time` 单位为1/30秒<br />这时，若图层的`time_step`为 3, 则说明本合成实际为10fps. 如果图层的`total_time`是 105, 那么合成长度为 3.5 秒  |
| `field`  | 只有在 [PF_Cmd_GLOBAL_SETUP](command-selectors.html)设置为 [PF_OutFlag_PIX_INDEPENDENT](PF_OutData.html)时生效<br />使用本变量来对上下场进行操作（现在几乎没用了）。  |
| `shutter_angle`  | 运动模糊快门角度。 <br />不激活动态模糊时为0，激活并指定目标图层时被启用。<br />`shutter_angle == 180`  表示 `current_time` 和 `current_time + 1/2 time_step` 之间的时间间隔(也就是1/2 time_step)<br />只有在[PF_Cmd_GLOBAL_SETUP](command-selectors.html)时设置 [PF_OutFlag_I_USE_SHUTTER_ANGLE](PF_OutData.html) 才能生效 <br />详情参见[动态模糊](../effect-details/motion-blur.html)  |
| `width`  | 源图层的尺寸，不一定与输入图像参数中的宽度和高度相同<br />缓冲区大小调整效应会导致这种差异。不受降频采样的影响。  |
| `height`  | 同上  |
| `extent_hint`  | 输入的可见部分和输出部分的交集；编入图层坐标系的一个包含合成信息的矩形框。<br />只遍历这个矩形像素可以显著加快效果。关于正确的用法，请参见本章后面的[extent_hint用法]。  |
| `output_origin_x`  | 输入缓冲区中输出缓冲区的中心点。除非特效更改中心点，否则为0。  |
| `output_origin_y`  |  见上  |
| `downsample_x`  | 当用户让AE每n个像素中只渲染一个时，控制点参数(Point control parameters)和图层尺寸参数来补偿分辨率损失。<br />插件需要下采样因子来解析代表图像中像素距离的标量。(比如滑块)<br />比如，四格像素的一个模糊，在下采样因子参数为1/2时，就被解析为一个两个像素的模糊（本变量是以比率表示的）。 仅在以下命令中有效<br /> - [PF_Cmd_SEQUENCE_SETUP](command-selectors.html),<br />[PF_Cmd_SEQUENCE_RESETUP](command-selectors.html),<br />[PF_Cmd_FRAME_SETUP](command-selectors.html),<br />[PF_Cmd_RENDER](command-selectors.html)  |
| `downsample_y`  | 见上 |
| `pixel_aspect_ratio`  | 像素宽高比（宽/高）。  |
| `in_flags`  | 未使用  |
| `global_data`  | 当运行其他入口指令时你的插件所储存的信息。调用插件之前属于锁定状态，调用之后属于解锁状态。  |
| `sequence_data`  |  |
| `frame_data`  |  |
| `start_sampL`  | 开始的采样数，相对于音频图层而言。  |
| `dur_sampL`  | 音频层的持续时长，以采样数表示，仅针对音频层。  |
| `total_sampL`  | 音频图层的总采样数，使用采样表示的total_time  |
| `src_snd`  | PF_SoundWorld 描述了输入音频，仅针对音频层。  |
| `pica_basicP`  |PICA Basic suite的指针，可以用来调用其他suite。  |
| `pre_effect_source_origin_x`  | 源图像在输入缓冲区中的中心点。只有在与一个帧的入口指令一起发送时才有效。<br />在这个插件之前的插件更改过图层大小和中心点时非零。<br />查看变换过后的大小与中心点来明确输出区域。<br />这对于有内隐空间操作(implicit spatial operations)的插件特别有用。比如绕一个图片的中心翻转图片。<br />NOTE: Checked-out 点参数会为当前的预效果进行调整而不是时间。  |
| `pre_effect_source_origin_y`  |  |
| `shutter_phase`  | 帧时间到快门打开时间用一帧持续时间的长短的百分数来表示的结果。  |

## extent_hint 用法

::: tip

对于[SmartFX](../smartfx/smartfx.html)来说，提示矩形框(hint rectangle)更有效...也更复杂...。
:::
使用`extent_hint`只处理那些需要输出的像素；这是最简单的优化。

通过在[PF_Cmd_GLOBAL_SETUP](command-selectors.html)(以及你的 PiPL)期间设置[PF_OutFlag_USE_OUTPUT_EXTENT](PF_OutData.html) ，告诉 After Effects 使用`in_data>extent_hint`

在测试`extent_hint`代码之前，请在首选项禁用缓存，这样 After Effects 就会在合成发生变化时渲染你的效果。否则，缓存机制会覆盖插件的输出(可能不正确)。

在合成中移动图层，对其进行裁剪。这时`output>extent_hint`是图层在合成中可见的部分。

给图层添加一个蒙版并移动它。这将改变`extent_hint`，变成包含所有alpha通道不为零的图像区域

`in_data>extent_hint`是这两个矩形(合成和遮罩)的交集，随这个区域改变而变化。

在改变大小和移动原点之前，会使用原输入图层的坐标系计算范围矩形(Extent rectangles),来简化已经设置过[PF_OutFlag_PIX_INDEPENDENT](PF_OutData.html) 的效果输入范围和输出范围交集运算。

根据`PF_InData>output_origin_x` 和`y`来偏移`extent_hint`,得到最终坐标系中输出范围（output extent）区域。

为了能在计算输出大小时降低采样，用户必须能够在完整分辨率下渲染。如果输出缓存超过30000x30000，请限制最大尺寸并使用警示框提示用户。

代码正确运行后，启用缓存并查看效果多久重新渲染(re-render)一次。

例如一个阴影特效，用户经常将静态阴影应用到静态图层。那么`output>extent_hint`就被忽略，则更频繁调用缓存。

对于扩展缓存(buffer-expanding)效果，请用插件变换范围和`output>extent_hint`相交,并在[PF_Cmd_FRAME_SETUP](command-selectors.html)期间设定尺寸

## 现在多了20%的像素

从 6.0 开始，传递的 `extent_hints` 比图层本身大 20%，可以为我们的渲染提供更高的预见性。

许多特效可以轻松扩展缓存, Ae经常使用这个提示长方形(hint rectangles)。

## 点控制和缓存扩展(BUFFER EXPANSION)

那些扩展输出缓存的特效通过在[PF_Cmd_FRAME_SETUP](command-selectors.html)时设置`output_origin_x/y`来改变原图层的左上角位置。这种变换会被报告至下一个特效中的`pre_effect_source_origin_x/y`。当这种变换发生时，点参数会自动被调整。

在你自己的特效之前应用一个诸如高斯模糊或尺寸重置(Resizer)这样的缓存扩展器(buffer expander)，并使用一个巨大的数值来改变大小。如果你的特效不能正确处理`pre_effect_source_origin_x/y`，打开或者关闭模糊会导致输出位置的改变。

所有的点参数值(在任何时间)都有`pre_effect_source_origin_x/y`描述的偏移值。对于大多数要对含透明度的图像进行处理的特效，如果缓存扩展(buffer expansion)随时间而改变(比如打过关键帧的模糊数量)，那么原图层的偏移会移动没有动画的点。在设计插件时，如果效果会在帧之间缓存点参数值，就要多考虑考虑这方面。
