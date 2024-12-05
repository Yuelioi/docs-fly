---
title: Parameter-Supervision
order: 16
---

# 参数管理

管理(Supervision)是指根据其他参数的值动态地改变一些参数的值。要管理一个参数，在`PF_Cmd_PARAM_SETUP`期间，在添加它之前设置[PF_ParamFlag_SUPERVISE](../effect-basics/PF_ParamDef.html) 。每当它被改变，你会收到[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 。改变的参数的索引(在插件的参数数组中)在`PF_UserChangedParamExtra`(额外)参数中发送。在 `PF_Cmd_USER_CHANGED_PARAM`期间，你可以改变任何参数的值和外观。

## Updating Parameter UI

如果在任何参数上设置了 `PF_ParamFlag_SUPERVISE` ，AE将向你发送 `PF_Cmd_UPDATE_PARAMS_UI` ，就像设置了 `PF_OutFlag_SEND_UPDATE_PARAMS_UI` 一样。

在 `PF_Cmd_UPDATE_PARAMS_UI` 期间，只能改变参数的外观和启用状态。使用[PF_ParamUtilSuite3]中的`PF_UpdateParamUI()`来更新用户界面，把你想修改的参数的 *copy* 传给它。不要试图修改原始参数。没有必要设置`PF_OutFlag_REFRESH_UI`；`PF_UpdateParamUI()`帮你处理。

::: tip

这是更新`PF_PUI_STD_CONTROL_ONLY`参数的用户界面的唯一方法。

:::

## 更新参数值

在[PF_Cmd_USER_CHANGED_PARAM](../effect-basics/command-selectors.html) 和[PF_Cmd_EVENT](../effect-basics/command-selectors.html)(*PF_Event_DO_CLICK*, *PF_Event_DRAG*, & *PF_Event_KEYDOWN*)。After Effects 不会更改其他时间做出的变动。

当改变参数值时(不仅仅是用户界面)，修改原始参数，并将`PF_Paramdef.uu.change_flags`设置为`PF_ChangeFlag_CHANGED_VALUE`。

这个变化也将更新用户界面，并且用户可以撤销。请注意，`PF_ChangeFlag_CHANGED_VALUE`不支持图层参数。

提供该套件是为了让效果插件能够访问它们的参数流，而不需要使用AEGP套件。至少其中一些功能是由几个第三方主机提供的。这些功能对于具有管理参数的效果特别方便。

## PF_ParamUtilSuite3

### PF_UpdateParamUI

```cpp
PF_UpdateParamUI(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
constPF_ParamDef*defP);
```

在效果控制面板中强制 After Effects 刷新参数的用户界面。

从 CC 2014 开始，After Effects 现在会尊重对自定义 UI 高度的改变。只要改变你的自定义 UI PF_ParamDef 的 ui_height，然后调用 PF_UpdateParamUI。效果的自定义用户界面高度将在效果控制窗口中更新。

从 CS6 开始，当一个插件禁用一个参数时，我们现在将该状态保存在 UI 开关中，这样插件就可以在将来检查该开关，看它是否被禁用。

注意：千万不要把 param[0]传给这个函数。

### PF_GetCurrentState

```cpp
PF_GetCurrentState(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
constA_Time*startPT0,
constA_Time*durationPT0,
PF_State*stateP);
```

这个 API，结合下面的 PF_AreStatesIdentical，用于确定在第一次调用 PF_GetCurrentState 和当前调用时，一组输入(图层、其他属性，或两者)是否不同，所以它可以用于缓存。以指定一个时间范围或所有时间。

在 CS6 中更新，增加了 param_index、startPT0 和 durationPT0。param_index 的预定义常量如下:

> - PF_ParamIndex_CHECK_ALL` - 检查每个参数，包括图层参数引用的每个图层。
> - PF_ParamIndex_CHECK_ALL_EXCEPT_LAYER_PARAMS` - 省略所有图层。传递一个特定的图层参数索引，将其作为唯一的图层参数进行测试。
> - PF_ParamIndex_CHECK_ALL_HONOR_EXCLUDE` - 与 CHECK_ALL 类似，但尊重 PF_ParamFlag_EXCLUDE_FROM_HAVE_INPUTS_CHANGED。

在开始和持续时间中传递 NULL 表示所有时间。对于做跨时间模拟并因此设置了 PF_OutFlag2_AUTOMATIC_WIDE_TIME_INPUT 的效果，当你询问时间范围时，它将被扩展到包括产生该范围所需的任何时间。

填充 PF_State，这是一种不透明的数据类型，用作效果参数当前状态的凭据(PF_State 用于我们内部的帧缓存数据库)。

### PF_AreStatesIdentical

```cpp
PF_AreStatesIdentical(
PF_ProgPtreffect_ref,
constPF_State*state1P,
constPF_State*state2P,
A_Boolean*samePB);
```

CS6 中的新功能。比较两个不同的状态，使用上述 PF_GetCurrentState`检索。

### PF_HasParamChanged

PFParamUtilsSuite3`中不再支持。

```cpp
PF_HasParamChanged(
PF_ProgPtreffect_ref,
constPF_State*stateP,
PF_ParamIndexparam_index,
PF_Boolean*changedPB);
```

给定一个 PF_State，如果任何测试的参数与保存的状态不同，则返回 true。与名字相反的是，该调用不提供测试单个参数的方法。至少，所有的非图层级参数将被测试。如果要测试一组特定的参数，对于更细的粒度(finer granularity)来测试一组特定的参数，请使用下面的 PF_HaveInputsChangedOverTimeSpan。

param_index 的预定义常量如下。

> - PF_ParamIndex_CHECK_ALL` - 检查每个参数，包括由层参数提引用每个图层。
> - PF_ParamIndex_CHECK_ALL_EXCEPT_LAYER_PARAMS` - 省略所有层。传递一个特定的层参数索引，将其作为唯一被测试的层参数。

### PF_HaveInputsChangedOverTimeSpan

`PFParamUtilsSuite3`中不再支持。使用`PF_AreStatesIdentical()`代替。

### PF_IsIdenticalCheckout

```cpp
PF_IsIdenticalCheckout(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
A_longwhat_time1,
A_longtime_step1,
A_u_longtime_scale1,
A_longwhat_time2,
A_longtime_step2,
A_u_longtime_scale2,
PF_Boolean*identicalPB);
```

如果一个参数的值在两个传递的时间内是相同的，返回 则TRUE`。注意：时间不需要是连续的；中间可能有不同的值。

### PF_FindKeyframeTime

```cpp
PF_FindKeyframeTime(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
A_longwhat_time,
A_u_longtime_scale,
PF_TimeDirtime_dir,
PF_Boolean*foundPB,
PF_KeyIndex*key_indexP0,
A_long*key_timeP0,
A_u_long*key_timescaleP0);
```

在参数流中搜索(按指定方向)下一个关键帧。后面三个参数是可选的。

### PF_GetKeyframeCount

```cpp
PF_GetKeyframeCount(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
PF_KeyIndex*key_countP);
```

返回参数流中的关键帧的数量。

### PF_CheckoutKeyframe

```cpp
PF_CheckoutKeyframe(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
PF_KeyIndexkey_index,
A_long*key_timeP0,
A_u_long*key_timescaleP0,
PF_ParamDef*paramP0);
```

从我们的关键帧数据库中检查指定参数的关键帧，param_index 为零。你可以请求时间、时间刻度，或者两者都不要求；如果你要执行自己的运动模糊，那就很有用。

### PF_CheckinKeyframe

```cpp
PF_CheckinKeyframe(
PF_ProgPtreffect_ref,
PF_ParamDef*paramP);
```

所有对 PF_CheckoutKeyframe 的调用都必须与这个 Check-in 相平衡，否则会带来痛苦。

### PF_KeyIndexToTime

```cpp
PF_KeyIndexToTime(
PF_ProgPtreffect_ref,
PF_ParamIndexparam_index,
PF_KeyIndexkey_indexP,
A_long*key_timeP,
A_u_long*key_timescaleP);
```

返回指定关键帧的时间(和时间尺度(timescale))。
