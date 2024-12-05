---
title: Parameters
order: 9
---

# Parameters

参数是随时间变化的数值流；源图像、滑块、角度、点、颜色、路径，以及用户可以操作的任何任意数据类型。

它们以 PF_ParamDefs 数组的形式传递给插件，尽管数组中的值只在某些入口指令中有效。

After Effects 效果API 最好的方面之一是参数插值和管理。

在 29.97 帧的四分之一时间里，快门角度有多大变化？这不是你的问题，放心交给 After Effects 吧。

在[PF_Cmd_PARAM_SETUP](command-selectors.html)期间，使用[PF_ADD_PARAM](../effect-details/interaction-callback-functions.html)描述插件中的参数。

你最多可以有(大约)38亿亿个参数可以使用，或者说你的用户在要求退款之前愿意筛选多少就有多少。明智的选择。

在注册 PF_ParamDefs 之前，用 AEFX_CLR_STRUCT(定义在 AE_Macros.h)清除它们，可以避免无数的问题。

## Parameter 类型

| Parameter 类型| Parameter 类型  | PF_ParamDefUnion 成员 | 参数值数据类型 | 描述  |
| --- | --- | --- | --- | --- |
| `PF_Param_LAYER`   | [PF_LayerDef](../effect-basics/PF_EffectWorld.html) | `ld`   | `A_long`   | 合成中的图像和音频层。所有效果都自动至少有1个图层参数:param[0],它们应用到的图层。<br />当用作效果参数时，这些显示为下拉菜单，用户使用该菜单选择当前合成中的图层。<br/>下拉菜单内容由AE生成。<br/>注意：这是对包含像素和音频样本的图层的引用，而不是实际的像素和音频样本。   |
| `PF_Param_SLIDER`   | `PF_SliderDef`   | `sd`   | `long`   | 不再使用   |
| `PF_Param_FIX_SLIDER`   | `PF_FixedSliderDef`   | `fd`   | `PF_Fixed`   | 已弃用. 多年来，我们一直在推广固定滑块。但我们现在推荐“PF_Param_FLOAT_SLIDERs”。<br/>额外的精度在许多情况下都有帮助，而且不再像以前那样昂贵。另外，我们只是厌倦了低字节/高字节(low byte / high byte)的愚蠢。<br/>"FIX_SLIDERs”提供了比“PF_Param_SLIDER”更高的精度。独立指定用户界面小数点。忽略“PF_Fixed”的低字(low word )以获得必要结果。   |
| `PF_Param_FLOAT_SLIDER`   | `PF_FloatSliderDef`   | `fs_d`   | `PF_FPLong`   | 滑块代表数值。FLOAT_SLIDERs包含相位、精度和曲线公差的值，供音频过滤器使用。<br/>指定最小值和最大值，用户可以移动滑块或键入数字来指定设置。<br/>也响应[Audio Filters](../audio/audio-considerations.html)中讨论的滑块开关.   |
| `PF_Param_ANGLE`   | `PF_AngleDef`   | `ad`   | `PF_Fixed`   | 以（固定点）度为单位的角度，精确到一度。<br/>用户可以指定多个旋转数(revolutions)，从而产生大于360的值。   |
| `PF_Param_CHECKBOX`   | `PF_CheckBoxDef`   | `bd`   | `PF_Boolean`   | `PF_ParamFlag_CANNOT_INTERP`强制打开所有复选框   |
| `PF_Param_COLOR`   | `PF_ColorDef`   | `cd`   | `PF_Pixel`   | 用户可以使用标准颜色选择器或滴管工具选择的RGB值（未使用alpha）。<br/>对于浮点精度，请使用[PF_ColorParamSuite1](../effect-details/parameters-floating-point-values.html)检索值。   |
| `PF_Param_POINT`   | `PF_PointDef`   | `td`   | `PF_Fixed`   | 一个二维点。该点在目标层空间中提供x和y值。<br/>图层的原点是左上角，向右x增加，向下y增加。<br/>从CS5.5开始，为了浮点精度，使用[PF_PointParamSuite1](../effect-details/parameters-floating-point-values.html)检索值。
<br/>接下来的Dusty历史教训：在API规范版本12.1（AE 4.0）之前，点的默认值在固定点0到100之间，基数点在第16位（即标准固定点）。<br/>以固定点指定（50,50）会产生图像的中心。<br/>因此，如果您将（50,50）作为默认位置，并且用户将效果应用于640乘480的图层，您将被发送的默认值将是（320,240）在固定点。<br/>指定12.1之前的API版本的插件仍将获得旧行为。
|
| `PF_Param_POPUP`   | `PF_PopupDef`   | `pd`   | `A_long`   | 选项列表。在namesptr中构建一个字符串，其中包含（只读）弹出条目列表（“Entry 1/Entry 2/Entry 3”）。<br/>AE复制数据并创建弹出菜单。<br/>一旦添加参数，这些条目就无法修改。<br/>“(-”可以绘制一个分隔符。   |
| `PF_Param_ARBITRARY_DATA`   | `PF_ArbitraryDef`   | `arb_d`   | `???`   | 自定义数据类型。<br/>[任意数据参数](../effect-details/arbitrary-data-parameters.html) 包含一个ID（您可以在给定的效果中使用多个自定义数据类型）、一个默认值（因此后效果知道您的数据类型应该以什么开头）和实际parameter.InAE的句柄，必须指定`PF_PUI_TOPIC/PF_PUI_CONTROL`或`PF_PUI_NO_ECW`。<br/>在PPro 8.0及更高版本中，可以不设置这些标志，这允许您在效果控件的右侧查看参数的关键帧轨道，而无需创建自定义控件。   |
| `PF_Param_PATH`   | `PF_PathDef`   | `path_d`   | `PF_PathID`   |路径参数是对应用于与效果相同的图层的蒙版的引用。<br/>路径参数数据不能直接访问；使用[PF_PathQuerySuite1](../effect-details/working-with-paths.html) 和 [PF_PathDataSuite](../effect-details/working-with-paths.html) 管理和查询路径。<br/>包含用户选择蒙版的索引。<br/>可以使用 [AEGP_MaskSuite6](../aegps/aegp-suites.html)中的`AEGP_GetLayerMaskByIndex` 获取 `AEGP_MaskRefH`   |
| `PF_Param_GROUP_START` `PF_Param_GROUP_END` | (none) (none)   |   |   | 参数组（主题）将参数组织成集合。<br/>每个组接收自己的旋转(twirly)，并将相对于相邻的参数或组在ECP中缩进。<br/>一个组可以嵌套在另一个组中。<br/>每个旋转可以由用户旋转打开或关闭，或者由效果以编程方式旋转。<br/>效果可以选择让某些组用旋转旋转打开初始化，而另一些组用旋转旋转关闭。.   |
| `PF_Param_BUTTON`   | `PF_Button`   | `button_d`   | (no value)   | 一个简单的按钮。使用[Parameter Supervision](../effect-details/parameter-supervision.html)检测按钮何时被按下。<br/>CS5.5到After Effects中的新功能。   |
| `PF_Param_POINT_3D`   | `PF_Point3D`   | `point3d_d`   | `PF_FpLong(3)`   | 一个三维点。<br/>CS5.5中的新功能。Premiere Pro中不支持。   |

## 滑块范围问题

如果你的滑块似乎被禁用了，但没有灰掉，请检查 valid_min、slider_min、valid_max 和 slider_max 字段。这个参数是`PF_Param_FIX_SLIDER`吗？如果是，你是否将你的 Mins 和 max 转换为合理的固定值？如果你使用 AE_Macros.h 中提供的宏，它们期望接收的是整数；传递 fixed point 是不行的。

## Point Parameter Origin

After Effects 会修改任何点参数以考虑原点偏移，这些偏移是由修改输出尺寸的 "上游 "效果引入的。由修改输出维度的“上游”效果引入。即使ECP UI指示点参数的值为（0,0），偏移量也已经被考虑在内。
