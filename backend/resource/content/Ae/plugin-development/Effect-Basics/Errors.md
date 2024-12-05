---
title: Errors
order: 4
---

# Errors

总是，总是，_总是_(总是！)从`main()`返回一个`PF_Err`。插件必须把所有的错误传回给 After Effects。

你必须把任何错误(由回调和 PICA 套件返回给你)传递给 After Effects，除非你已经处理了它们，这是非常重要的。

要注意返回正确的错误代码，并处理掉你分配的任何内存。

真的。我们是认真的。

## Error Codes

| **Error**                            | **Meaning**                                                                                            |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `PF_Err_NONE`                       | Success.                                                                                               |
| `PF_Err_INTERNAL_STRUCT_DAMAGED`    | Problems using a data structure.                                                                       |
| `PF_Err_UNRECOGNIZED_PARAM_TYPE`    | Problem with parameter data.                                                                           |
| `PF_Err_BAD_CALLBACK_PARAM`         | Problems using a parameter passed to a callback.                                                       |
| `PF_Err_CANNOT_PARSE_KEYFRAME_TEXT` | Return this from PF_Arbitrary_SCAN_FUNC` when problems occur parsing the clipboard into keyframe data. |

## Error Reporting Policy

After Effects 有一个统一的错误处理政策，请遵循它。

如果你在插件的代码中遇到了错误，请立即向用户报告，然后再从插件返回到 After Effects。

After Effects 认为在你的插件执行过程中遇到的来自操作系统的错误是你的。

如果你从我们的回调函数中得到一个错误代码，请把它传回给 After Effects；我们已经报告了它。

After Effects 不会报告内存不足的错误。在 RAM 预览期间，以及 After Effects 在 noui 模式下运行时，错误报告总是被抑制。

要在插件中报告错误，请设置`PF_OutFlag_DISPLAY_ERROR_MESSAGE`，并在[PF_OutData>return_msg](PF_OutData.html) 中描述错误。

这样做会把你的错误输入到渲染日志中，并防止由渲染引擎或脚本驱动的渲染中出现系统挂起。

## Dig In

现在你已经对特效插件有了基本的了解，并准备开始尝试一些真正的代码。来吧，开始吧!

在掌握了插件的基本设置后，你可能会有一些关于可重用代码、高级功能以及如何优化代码以使其更快的问题。

为此，After Effects 通过函数套件暴露了大量的内部功能。

通过依赖 After Effects 代码的实用功能，你应该能够快速实现你的图像处理算法。

这将在[Effect Details](../effect-details/effect-details.html)(#effect-details-effect-details)中讨论。
