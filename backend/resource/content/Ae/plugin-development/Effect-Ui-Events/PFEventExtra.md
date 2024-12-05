---
title: PF_EventExtra
order: 3
---

# PF_EventExtra

这个结构提供了当前事件的上下文信息。After Effects 在[Entry Point](../effect-basics/entry-point.html)函数的额外参数中传递一个指向这个结构的指针。

`PF_EventUnion`(在`PF_EventExtra`中发送)随事件类型而变化，包含该事件的特定信息。

| 成员 | 目的  |
| --------- | --------- |
| `contextH`  | 处理 PF_Context <br />此绘制上下文与 [Drawbot suites](../effect-ui-events/custom-ui-and-drawbot.html) 用于绘制，也用于 [UI回调](../effect-ui-events/ui-callbacks.html).|
| `e_type`  | 哪个 [事件](../effect-ui-events/effect-ui-events.html) 正在发生 |
| `effect_win`  | 如果事件发生在效果窗口中，则是关于`PF_EffectWindowInfo`的事件<br />否则，从AE 5.0开始`effect_win`可以替换为`PF_WindowUnion`.<br />此结构包含 `PF_EffectWindowInfo`和`PF_ItemWindowInfo`, 这（目前）只是项目窗口的端口矩形。<br />只有在编译期间定义了`PF_USE_NEW_WINDOW_UNION`时才会发生替换；否则，它将继续只是一个`PF_EffectWindowInfo`。 |
| `evt_in_flags`  | 事件输入标志。目前只包含一个值，`PF_EI_DONT_DRAW`, 您应该在绘图前检查它！|
| `evt_out_flags` | 以下一项或多项，与按位或运算组合：<br />`PF_EO_NONE`<br />`PF_EO_HANDLED_EVENT`告诉AE你已经处理了这个事件。<br />`PF_EO_ALWAYS_UPDATE`强制AE重新渲染复合以响应每次单击或拖动；这与“alt-scrubing”参数值生成的行为相同。  `PF_EO_NEVER_UPDATE`防止AE重新渲染合成，直到用户停止单击和拖动。<br /> `PF_EO_UPDATE_NOW`告诉AE在调用后事件返回后立即更新视图`PF_InvalidateRect` |

## PF_Context

PF_Context 详细说明了事件的 UI 环境。

| **Member** | **Purpose**  |
| --- | ------- |
| `magic` | Do not change. |

## PF_EffectWindowInfo

如果一个事件发生在 ECP 中，`PF_EffectWindowInfo` 将在 `PF_EventExtra` 中发送。

| **Member**  | **Purpose**  |
| -------- | ----- |
| `index` | 表示效果窗口中的哪个参数受到影响。控件编号从0到控件总数减1。 |
| `area`  | 表示控件标题（`PF_EA_PARAM_TITLE`）或控件本身（`PF_EA_CONTROL`）是否受到影响。<br />标题是参数的主题（“展开”）启动时仍然可见的区域。|  |
| `current_frame` | `PF_Rect`指示控件所占区域的完整帧。 |
| `param_title_frame` | `PF_Rect` 指示控件的标题区域。|
| `horiz_offset`  | 与要绘制到标题中的标题区域左侧的水平偏移量。|
