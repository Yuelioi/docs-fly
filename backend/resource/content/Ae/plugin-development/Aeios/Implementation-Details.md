---
title: Implementation-Details
order: 4
---

# 实施细节

## 导出 Bit-Depth

在输出模块设置中，用户可以根据 AEIO 在[AEIO_FunctionBlock4](new-kids-on-the-function-block.html) 的`AEIO_GetDepths()`中声明支持的选项选择一个深度。

如果一个插件支持更高的比特深度导出，它应该能够处理这些在`AEIO_AddFrame()`或`AEIO_OutputFrame()`中传递的更高比特深度的 PF_EffectWorlds，即使导出设置没有设置为相同的深度。

交付给 AEIO 的帧，和最终的输出不一定是相同的深度。

如果 After Effects 认为最终输出的质量更高，你可能会得到以项目比特深度传递的帧而不是最终输出的帧。

## 用户数据 vs 选项

可以使用用户数据分配或选项句柄来存储文件的元数据。

我们用用户数据来存储要嵌入文件中的信息(假定文件格式支持这种信息)；标记数据、字段标签等。

我们使用选项句柄来存储关于文件的信息；输出设置、尺寸、所用压缩设置的细节。
