---
title: Accessing-Audio-Data
order: 1
---

# 访问音频数据

使用`PF_CHECKOUT_LAYER_AUDIO`从[Interaction Callbacks](../effect-details/interaction-callback-functions.html) 检索一个音频层。

这个层是不透明的；使用`PF_GET_AUDIO_DATA`来获取该音频的具体细节。

和像素数据一样，你必须尽快检查音频。

如果你的效果要求输入的时间跨度与输出的时间跨度不同，在 "PF_Cmd_AUDIO_SETUP "中更新`startsampL`和`endsampL`字段，从[Frame Selectors](../effect-basics/command-selectors.html) 。

## 扩展音频剪辑

你不能通过 API 延长一个音频片段的长度。

然而，在应用你的效果之前，用户延长片段的长度是一个相对简单的事情。对该层应用时间重映射，并简单地延长输出点。

如果你要给一个声音片段添加延迟效果，你要让它有时间渐渐消失，而不是在原来的结束点截断声音。

记录用户在应用你的效果时应该采取的步骤。
