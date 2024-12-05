---
title: Bigger-Differences
order: 2
---

# 更大的差异

只要一个效果只支持 After Effects 所支持的基本 ARGB_8u 像素格式，Premiere Pro 就会尝试模仿 After Effects 的托管行为，并隐藏各种差异，因为渲染管道架构不同。但是如果一个效果想要支持额外的像素格式，比如 32 位 RGB，就要准备好处理进一步的分歧行为。

## 像素格式

Premiere Pro 提供了函数套件，用于支持 After Effects 使用的 8 位 RGB 格式--ARGB_8u 以外的像素格式。这些像素格式包括 Premiere Pro 本地的 8 位 RGB 格式--BGRA_8u，以及 YUV、32 位格式等等。关于各种像素格式的详细讨论，请参见《Premiere Pro SDK 指南》中的 [像素格式和色彩空间](http://ppro-plugin-sdk.aenhancers.com/universals/pixel-formats-and-color-spaces.html)。

使用 PF 像素格式套件(定义在 PRAESDKSupport.h 中)来注册其他像素格式的[PF_EffectWorld / PF_LayerDef](../effect-basics/PF_EffectWorld.html)。使用 Premiere 像素格式套件(定义在恰如其名的 PrSDKPixelFormatSuite.h)来获得这些像素格式的黑白值。

像 PF_BLEND()这样的 After Effects 函数还没有被增强到可以处理 8 位 RGB 以上的像素格式。

## 32-Bit 浮点支持

Premiere Pro 不支持 After Effects 的 16 位渲染或 SmartFX。对于 Premiere Pro 中的 32 位渲染，你需要声明对 32 位像素格式的支持(见上一节)，然后为 PF_Cmd_RENDER 实现 32 位渲染。你可以通过这种方式支持多种渲染深度。请看 SDK Noise 示例项目中的一个例子。

根据应用效果的片段，32 位处理并不总是必要的，以保持源输入的质量。但是有一些设置可以强制进行 32 位渲染，如果需要的话，可以给效果处理提供更精细的颗粒度和更多的空间。进入设置>序列设置>视频预览>最大比特深度，以控制从时间轴上的预览。对于导出文件，使用导出设置>视频>基本设置>以最大深度渲染。

## PF_CHECKOUT_PARAM 与像素格式

在 CS6 之前，PF_CHECKOUT_PARAM()只返回 8 位 ARGB 缓冲区，而不考虑当前用于渲染的像素格式。从 CS6 开始，一个效果可以选择以与渲染请求相同的格式获取帧，无论它是 32 位浮点、YUV 等。

插件可以请求这种行为，但现有的插件将继续工作，接收 8 位 ARGB 帧。这个调用是 `EffectWantsCheckedOutFramesToMatch` `RenderPixelFormat()`，在 PF Utility Suite 中，定义在 PrSDKAESupport.h 中。这个调用应该在 PF_Cmd_GLOBAL_SETUP 中进行，同样的入口指令，一个效果已经使用 `AddSupportedPixelFormat()`来支持超过 8 位 RGB。
