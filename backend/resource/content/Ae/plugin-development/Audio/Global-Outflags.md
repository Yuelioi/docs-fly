---
title: Global-Outflags
order: 6
---

# 全局 Outflags

所有的音频效果必须设置`PF_OutFlag_AUDIO_EFFECT_TOO`或`PF_OutFlag_AUDIO_EFFECT_ONLY`。

`PF_OutFlag_I_USE_AUDIO`是用于检查音频数据的视觉效果，但不修改它。

`PF_OutFlag_AUDIO_FLOAT_ONLY`，`PF_OutFlag_AUDIO_IIR`和`PF_OutFlag_I_SYNTHESIZE_AUDIO`提供了对音频输出的更大控制(详见[PF_OutFlags](./effect-basics/PF_OutData.html) )。
