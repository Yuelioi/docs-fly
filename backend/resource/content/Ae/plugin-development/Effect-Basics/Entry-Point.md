---
order: 3
title: Entry-Point
---

# 入口函数

After Effects 和效果插件之间的所有通信都是由 After Effects 发起的，而且都是由主机(After Effects)调用一个入口函数函数来完成。

对于所有的效果插件来说，入口函数函数必须有以下特征。

```cpp
PF_Err main (
 PF_Cmd cmd,
 PF_InData *in_data,
 PF_OutData *out_data,
 PF_ParamDef *params[],
 PF_LayerDef *output,
 void *extra)

```

上述入口函数函数的名称是 "main"，但也可以是[PiPL 资源](../intro/pipl-resources.html)中指定的任何内容。

在每次调用入口函数之前，After Effects 都会更新[PF_InData](PF_InData.html) 和插件的参数数组 PF_ParamDef\[\](除非另有说明)。

在插件从被调用中反馈后，After Effects 会检查[PF_OutData](PF_OutData.html) 是否有变化，并在适当的时候使用效果所渲染的 PF_LayerDef。

## 入口函数参数

|参数列表|用途|
| ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [cmd](../effect-basics/command-selectors.html)  | After Effects 设置 [入口指令](../effect-basics/command-selectors.html) 来控制插件想做什么                                                                                                                                                                                                                                                                                                                                                                            |
| [in_data](../effect-basics/PF_InData.html)      | 应用程序状态的信息以及插件需要处理的数据. 许多界面和图像处理函数的指针。                                                                                                                                                                                                                                                                                                                                                                       |
| [out_data](../effect-basics/PF_OutData.html)    | 通过在out_data中设置字段，将信息传回After Effects。                                                                                                                                                                                                                                                                                                                                                                                                |
| [params](../effect-basics/parameters.html)      | 在 in_data>current_time 中提供的插件参数数组。 `params[0]` 是应用效果的输入图像( 一个 [PF_EffectWorld / PF_LayerDef](../effect-basics/PF_EffectWorld.html)) 这些值仅在某些入口指令期间有效 (详见 [入口指令描述](../effect-basics/command-selectors.html)). 这里详细讨论了参数: [PF_ParamDef](../effect-basics/PF_ParamDef.html). |
| [output](../effect-basics/PF_EffectWorld.html)  | 输出图像，由效果插件渲染并传回After Effects。仅在某些入口指令期间有效。                                                                                                                                                                                                                                                                                                                                             |
| [extra](../effect-ui-events/PF_EventExtra.html) | 额外的参数, 随发送命令而变化或 (在[PF_Cmd_EVENT](../effect-basics/command-selectors.html)) 的 [事件类型](../effect-ui-events/effect-ui-events.html).主要用于事件管理和  [参数监控](../effect-details/parameter-supervision.html).                                                                                                                                                                   |
