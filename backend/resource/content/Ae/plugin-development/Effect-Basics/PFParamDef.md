---
order: 8
title: PF_ParamDef
---

# PF_ParamDef

After Effects 在每个入口指令中传递给插件一个 PF_ParamDefs 数组，描述插件在当前时间的参数。params 数组中的值只在某些入口指令中有效(这在[入口指令描述](command-selectors.html)中已经注明)。

## Param Zero

第一个参数，即params[0]，是你的插件接收到的输入图像(一个[PF_EffectWorld / PF_LayerDef](PF_EffectWorld.html))

## 剩下的参数

所有的参数类型都由 PF_ParamDef 表示。使用联合体，里面有共用体(Unions)，所以只需要填充PF_ParamDef中的这部分信息。

## PF_ParamDef Members

| 数据类型  | 名称  | 描述 |
| ---- | ---- | ------- |
| `A_long`   | `id`   | 参数的ID . 如果跨版本维护参数的ID，可以在插件的未来版本中重新排序参数，而不会导致用户重新应用您的效果。   |
| `PF_ChangeFlags`   | `change_flags` | 设置是否更改了参数值。仅在拖动（不是单击！）事件期间有效,<br />[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 或者是[PF_Cmd_UPDATE_PARAMS_UI](../effect-basics/command-selectors.html). |
| [PF_ParamUIFlags](../effect-basics/PF_ParamDef.html) | `ui_flags`   | 在添加之前指定参数的UI行为；在事件处理期间只能设置`PF_PUI_DISABLED`   |
| `A_short`   | `ui_width`   | 用户交互界面的宽度（仅适用于非标准参数）。   |
| `A_short`   | `ui_height`   | 用户交互界面的高度（仅适用于非标准参数）。   |
| [PF_ParamType](../effect-basics/parameters.html)   | `param_type`   | 参数的类型   |
| `A_char[32]`   | `name`   | 参数名. 可以在事件处理时更改.<br />自从After Effects 1.0以后, 我们需要一些长名称命名，31个字符内都可以，就像俳句（haiku）一样。   |
| [PF_ParamFlags](../effect-basics/PF_ParamDef.html)   | `flags`   | 添加前指定参数的UI行为；可能在事件处理期间仅设置 `PF_ParamFlag_COLLAPSE_TWIRLY`.   |
| `PF_ParamDefUnion`   | `u`   | 一个包含几乎所有（Ae中）常用的[参数类型](../effect-basics/parameters.html)的储存共同体（union）。<br />只有被param_type定义的类型才可使用。   |

## 参数 UI 开关

用这些开关来控制一个参数的用户界面。

不要把用户界面开关和行为开关混为一谈；它们位于你的参数定义中的不同区域，如果应用不当，会导致不可预知的行为。

(当然，如果你不想纠结于这些乱七八糟的开关和定义，只想专注于图形算法或者得到某种效果，你完全可以使用SDK里面提供的宏定义来安装控件，SDK里面的宏定义会替你掌控这些开关)

| **Flag**   | **Description**      |
| --- | -------------- |
| `PF_PUI_TOPIC`   | 当你要实现“topic”功能时设置该开关<br /> “topic”是效果控件窗口（ECW）UI的一部分，可以展开和折叠，就像P粒子中常用的那一排可以展开折叠的栏状UI）的时候就打开这个开关。同时也需要你在PF_Cmd_GLOBAL_SETUP期间设置 `PF_OutFlag_CUSTOM_UI`     |
| `PF_PUI_CONTROL`   | 当你要为实现控制区域(control area)功能的时候就打开这个开关。<br />同时也需要你在  `PF_Cmd_GLOBAL_SETUP` 时设置`PF_OutFlag_CUSTOM_UI`<br />详情见 [Effect UI & Events](../effect-ui-events/effect-ui-events.html)    |
| `PF_PUI_STD_CONTROL_ONLY`   | S如果你想要仅标准控制模式(这将导致没有数据流与控件相关联，因此时间轴上的关键帧信息将无法获得)就打开这个开关。你一般会用他来控制你自己的数据，或者自定义UI，或者一组复合控制。同时你不能把这个开关用在 <br /> _PF_Param_CUSTOM_ ,<br /> _PF_Param_NO_DATA_  , <br /> _PF_Param_LAYER_ ,<br /> _PF_Param_ARBITRARY_DATA_ ,<br /> _PF_Param_PATH_ .<br />在这个开关打开之前你需要打开 `PF_ParamFlag_SUPERVISE` (否则你将永远得不到数据变化)。<br />这个开关不需要你打开[PF_OutFlag_CUSTOM_UI](../effect-basics/PF_OutData.html) <br />如果只想要PF_Param_ARBITRARY_DATA的标准控件，你只需要加入一个或多个PF_PUI_STD_CONTROL_ONLY，然后在[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 期间可以修改您的数据。  |
| `PF_PUI_NO_ECW_UI`   | 支持该模式的控件并打开这个开关，<br />不过这样的话，我们推测你应该想通过一些别的方式来修改你这个控件的数值(比如，通过时间轴窗口来更改，或者你通过在`PF_Cmd_USER_CHANGED_PARAM` 的时候开启了`PF_ParamFlag_SUPERVISE` 的控件来更改)。.<br />在Ae中，这个开关不会让这个控件在时间轴窗口中消失，也就是你仍能看到关键帧，但在PPro中，连时间轴窗口中的控件也会被隐藏掉。    |
| `PF_PUI_ECW_SEPARATOR`   | Ae中没用，Pr中有用。 如果您想在效果控制窗口中的该参数上方设置一条粗线，请设置此开关。提供此功能以便在需要时可以对参数进行可视化分组（无需添加组）。此开关可以在运行时通过 `PF_UpdateParamUI()`方法更改。     |
| `PF_PUI_DISABLED`   | 使控件无效化(会变灰色)，经常结合[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html)使用.     |
| `PF_PUI_DONT_ERASE_TOPIC`   | Ae不会抹去“主题组”。      |
| `PF_PUI_DONT_ERASE_CONTROL` | Ae不会抹去控件控制。      |
| `PF_PUI_RADIO_BUTTON`   | 未在After Effects中使用，但在Premiere中使用。将参数显示为单选按钮组。仅对  `PF_Param_POPUP`生效.     |
| `PF_PUI_INVISIBLE`   | 一开始只支持Pr，自Ae cs6之后Ae也支持了，这个会把控件在插件面板和时间轴窗口同时隐藏掉。
Pr中这个开关是动态的，可以在[PF_UpdateParamUI](../effect-details/parameter-supervision.html) 中随时更改。    |

除了这些开关之外，还可以通过使用[AEGP_GetDynamicStreamFlags](../aegps/aegp-suites.html)来隐藏或显示一个效果参数。

## Parameter Flags

行为开关和UI行为开关控制控件的不同特质。在[PF_Cmd_PARAM_SETUP](../effect-basics/command-selectors.html)期间,在安装控件前设置这些开关。对于一些需要在运行期间修改的开关，会有相关的提示。

| **Flag**   | **Meaning**   |
| --- | -------- |
|   |   |
| `PF_ParamFlag_CANNOT_TIME_VARY`   | 控件不会随时间的改变而改变。在时间轴面板上将不会提供关键帧码表(k帧器)   |
|   |   |
| `PF_ParamFlag_CANNOT_INTERP`   | 数值不会被线性插值并将使用非连续插值。对于一些布尔类型控件有用，并能加快渲染。   |
| `PF_ParamFlag_COLLAPSE_TWIRLY`   | 在[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html)期间设置.<br />现在可以在 [PF_Cmd_UPDATE_PARAMS_UI](../effect-basics/command-selectors.html) and [PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 打开和关闭，也就是你可以随意打开或收起你的主题组。 |
| `PF_ParamFlag_SUPERVISE`   | 可通过这个开关来在[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 期间获取信息。<br />详情见[Parameter Supervision](../effect-details/parameter-supervision.html)   |
| `PF_ParamFlag_START_COLLAPSED`   | 控制主题组的下拉列表的收合，不仅是在[PF_Cmd_PARAM_SETUP](../effect-basics/command-selectors.html)期间,在插件管理(Parameter Supervision)期间就可以对这个开关进行修改，想要让这个开关发挥作用就必须提前设定[PF_OutFlag2_PARAM_GROUP_START_COLLAPSED](../effect-basics/PF_OutData.html)   |
| `PF_ParamFlag_USE_VALUE_FOR_OLD_PROJECTS`   | 这只会影响使用旧版本效果保存的项目的加载，该效果缺少稍后添加的参数。<br />设置时, `PF_ADD_PARAM()` 中的 `PF_ParamDef.value`将用于初始化缺失的参数，但当新应用或重置效果时，dehault字段仍将用于参数的初始值。<br />当您希望参数默认为一个值, 而且需要将其设置为其他值以保留旧项目的呈现行为时，这很有用。   |
| `PF_ParamFlag_LAYER_PARAM_IS_TRACKMATTE`   | 仅适用于Premiere Pro：仅对图层参数有效。表示图层参数被用作带有应用滤波器的跟踪遮罩。<br/>在AE中无效   |
| `PF_ParamFlag_EXCLUDE_FROM_HAVE_INPUTS_CHANGED` | 只有效果在设置[PF_OutFlag2_AUTOMATIC_WIDE_TIME_INPUT](../effect-basics/PF_OutData.html)时,而且调用 [PF_AreStatesIdentical](../effect-details/parameter-supervision.html) 或 [PF_HaveInputsChangedOverTimeSpan](../effect-details/parameter-supervision.html)   |
| `PF_ParamFlag_SKIP_REVEAL_WHEN_UNHIDDEN`   | CS6 新出.如果此参数未隐藏，则此开关告诉After Effects不要打开任何父级，也不要将参数滚动到效果控制面板和时间轴面板中的视图中。<br />当绘制描边时，After Effects在内部使用此方法，以免因为显示参数来分散用户的注意力。但在另一种情况下，打开时间重新映射时，该参数会被显示。<br/>因此，我们为您提供了对自己效果中参数的相同控制。   |

## PF_ValueDisplayFlags

在 PF_ParamDefUnion 中，PF_FloatSliderDef 和 PF_FixedSliderDef 都有一个成员变量 PF_ValueDisplayFlags，它允许它们响应用户的像素值显示偏好(他们在信息调色板中设置)。如果设置了这个，参数的值将显示为 0-1、0-255、0-32768，或者 0.0 到 1.0，取决于偏好。你也可以设置第一个位(PF_ValueDisplayFlag_PERCENT)，在参数的显示值上附加一个百分号。

我们知道你永远不会做这样的事情，但是如果你创建了一个显示为百分比的参数，不要通过允许 0 到 100 以外的任何范围来混淆用户。请注意。百分比的意思是 "100 分之 1"。
