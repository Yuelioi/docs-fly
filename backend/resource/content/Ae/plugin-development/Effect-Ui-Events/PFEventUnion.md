---
title: PF_EventUnion
order: 4
---

# PF_EventUnion

PF_EventExtra 中的 PF_EventUnion 是以下四个结构体的联合。

## 点击

鼠标点击或拖动发生在自定义用户界面的区域内。

### PF_DoClickEventInfo

| 成员  | 目的 |
| --- | --- |
| `when`  | 发生单击事件的（操作系统级）时间。   |
| `screen_point`  | 在屏幕坐标中发生了单击。对于自定义Comp UI，可以使用[UI Callbacks](ui-callbacks.html) .有关示例，请参阅CCU示例项目。|  |
| `num_clicks`  | 点击次数。  |
| `modifiers`  | 单击时按住了哪些修改键（如果有）。  |
| `continue_refcon[4]`  | 插件可以使用`4A_intptr_t`数组在"单击-拖动-拖动"序列中存储信息。  |
| `send_drag`  |将此标志设置为`TRUE`以指示继续拖动。下一个单击事件将实际上是拖动事件。|
| `last_time`  | 设置拖动事件何时结束（用户已释放鼠标按钮）。  |

## 绘制

After Effects 需要你的自定义 UI 来刷新。

注意：在处理绘制请求时，使用[PF_InData](../effect-basics/PF_InData.html)中提供的图像尺寸(而不是你的输入层的尺寸，就像在[PF_Cmd_RENDER](../effect-basics/command-selectors.html))。

### PF_DrawEventInfo

| 成员  | 目的 |
| --- | --- |
| `update_rect`  | 在上下文窗口的坐标系中绘制的矩形。这些坐标可以使用 [UI Callbacks](ui-callbacks.html) .<br/> 有关示例，请参阅CCU示例项目。 |
| `depth`  | 绘图上下文的像素深度。 |

## Keydown

用户按了一个键，效果的用户界面被激活。

使用 AE_EffectUI.h 中的宏来访问和操作收到的按键代码。

为了在 Premiere Pro 中接收按键事件，插件必须在.CustomUIInfo.events 中设置 PF_CustomEFlag_COMP。\PF_PF_Cmd_PARAM_SETUP。
PF_KeyDownEvent

### when

点击发生的时间。

### screen_point

按键时鼠标指针的屏幕坐标。对于 Custom Comp UI，这些坐标可以使用[UI Callbacks](../effect-ui-events/ui-callbacks.html)转换为合成坐标。请看 CCU 示例项目的例子。

### key_code

要么是一个字符代码(对于可打印的字符，我们使用未移位(unshifted)的大写版本；A 不是 a，7 不是&)，要么是一个控制代码。

> - `PF_ControlCode_Unknown`
> - `PF_ControlCode_Space`
> - `PF_ControlCode_Backspace`
> - `PF_ControlCode_Tab`
> - `PF_ControlCode_Return`
> - `PF_ControlCode_Enter`
> - `PF_ControlCode_Escape`
> - `PF_ControlCode_F1`

…

> - `PF_ControlCode_F24`
> - `PF_ControlCode_PrintScreen`
> - `PF_ControlCode_ScrollLock`
> - `PF_ControlCode_Pause`
> - `PF_ControlCode_Insert`
> - `PF_ControlCode_Delete`
> - `PF_ControlCode_Home`
> - `PF_ControlCode_End`
> - `PF_ControlCode_PageUp`
> - `PF_ControlCode_PageDown`
> - `PF_ControlCode_Help`
> - `PF_ControlCode_Clear`
> - `PF_ControlCode_Left's`
> - `PF_ControlCode_Right`
> - `PF_ControlCode_Up`
> - `PF_ControlCode_Down`
> - `PF_ControlCode_NumLock's`
> - `PF_ControlCode_Command`
> - `PF_ControlCode_Option`
> - `PF_ControlCode_Alt`
> - `PF_ControlCode_Control`
> - `PF_ControlCode_Shift`
> - `PF_ControlCode_CapsLock`
> - `PF_ControlCode_ContextMenu`

### modifiers

在按键过程中，哪些(如果有的话)修改键被按下。

> - `PF_Mod_NONE`
> - `PF_Mod_CMD_CTRL_KEY` (Mac 的 cmd，Windows 的 ctrl)
> - `PF_Mod_SHIFT_KEY` (Mac 上的 cmd，Windows 上的 ctrl)
> - `PF_Mod_CAPS_LOCK_KEY` (Mac 上的 cmd，Windows 上的 ctrl)
> - `PF_Mod_OPT_ALT_KEY` (option on Mac, alt on Windows)
> - `PF_Mod_MAC_CONTROL_KEY` (Mac 上的选项，Windows 上的 alt)

## AdjustCursor

光标已经移动到(但没有离开)效果的自定义用户界面上，以允许效果改变光标。

### PF_AdjustCursorEventInfo

| 成员  | 目的 |
| ---- | --- |
| `screen_point`| 鼠标指针的屏幕坐标。对于自定义Comp UI，可以使用[UI Callbacks](../effect-ui-events/ui-callbacks.html).有关示例，请参阅CCU示例项目。 |
| `modifiers` | 发送消息时按住了哪些修改键（如果有的话）。|
| `set_cursor`| 将此设置为所需的光标，如果使用特定于操作系统的调用自己设置了光标，则设置为`PF_Cursor_CUSTOM`。有关内置光标的完整枚举，请参阅AE_EffectUI. h。如果不想覆盖光标，请将其设置为`PF_Cursor_NONE`，或者直接忽略此消息。  |

## 任意参数事件

After Effects 需要你的插件来管理它的任意数据参数。

虽然任意数据类型不需要自定义用户界面的支持，但是`PF_ArbParamsExtra`遵循 EventInfo 模型。

### PF_ArbParamsExtra

#### which_function

一个 `PF_FunctionSelector`, 表示哪个函数被调用。

#### id

由 After Effects 使用；将与在*PF_Cmd_PARAM_SETUP*期间分配给任意数据类型的 ID 相匹配。

#### padding

用来进行字节对齐

#### u

(其中一个将被传递；见[任意数据参数](../effect-details/arbitrary-data-parameters.html))

```cpp
u {
  new_func_params
  dispose_func_params
  copy_func_params
  flat_size_func_params
  flatten_func_params
  unflatten_func_params
  interp_func_params
  compare_func_params
  print_size_func_params
  print_func_params
  scan_func_params
}
```
