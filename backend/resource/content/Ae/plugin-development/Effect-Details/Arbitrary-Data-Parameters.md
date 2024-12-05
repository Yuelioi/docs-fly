---
title: Arbitrary-Data-Parameters
order: 3
---

# 任意数据参数

AE现有参数类型不能充分表示某些值。可以通过创建任意数据类型的参数或“arb数据”来创建和注册任何用于AE插值的数据。可以依赖我们的插值引擎和参数管理，而无需强制将数据转换为预定义的参数类型。

我们已经为自定义数据类型创建了一个新的信息传递结构，它很容易被概念化为 C++类的成员(和友元(friend))函数。如果你使用 arb 数据，你必须对这里详述的所有入口指令做出回应。

这些函数处理自定义数据结构的管理。你的 arb 数据将在用户的要求下被卸载和重新加载；提供磁盘安全的 展平 和 非展平 函数。

## 任意数据类型入口指令

| **入口指令** | **响应** |
| ------------------------------ |-------------------------------------------------- |
| `PF_Arbitrary_NEW_FUNC` | 为arb数据的新实例分配、填充和返回句柄。|  
| `PF_Arbitrary_DISPOSE_FUNC`| 释放并销毁任意数据类型的实例。|
| `PF_Arbitrary_COPY_FUNC`  | M创建现有实例的副本。将向您传递两个句柄，但只有源句柄包含有效实例。您必须创建一个新实例，从源句柄复制值，并将其放入目标句柄。如果向您传递了NULL句柄，请创建arb数据的默认实例。|  | PF_Arbitrary_FLAT_SIZE_FUNC` | 您将被传递一个句柄到您的数据类型的实例，以及一个变量，您可以在其中返回该实例的扁平化版本的大小。 |
| `PF_Arbitrary_FLATTEN_FUNC` | 展平您传递的实例，并将其放在提供的缓冲区中。缓冲区将是您响应PF_Arbitrary_FLAT_SIZE_FUNC时报告的大小。 |  | PF_Arbitrary_UNFLATTEN_FUNC` | 将缓冲区解压到任意数据类型的实例中，并放入传递给您的句柄中。|
| `PF_Arbitrary_INTERP_FUNC`  | 您的插值函数将三个句柄传递给任意数据类型的实例；一个包含初始值（0），一个包含最终值（1），第三个用于保存您的插值数据（介于0和1之间）。您还会被传递一个浮点数，指示您的解释值应该在0和1之间的位置.分配一个实例并用插值数据填充它。然后将插值实例放入您传递的句柄中。当计算归一化时间值时，速度曲线已经考虑在内。<br />注意：如果[in_data>effect_ref](../effect-basics/PF_InData.html)为NULL，则永远不要签出参数。 |  | PF_Arbitrary_COMPARE_FUNC`| 将向您传递任意数据的两个实例和一个指向比较结果的指针。使用PF_ArbCompareResult（参见AE_Effect.h）值之一填充结果，以指示第一个值是否等于、小于、大于或不等于第二个值。|
| `PF_Arbitrary_PRINT_SIZE_FUNC` | 通过设置print_sizePLu（print_size_func_params成员，PF_ArbParamsExtra结构的一部分）来指示打印参数当前值所需的缓冲区大小。 |
| `PF_Arbitrary_PRINT_FUNC`  | 格式化任意数据以进行基于文本的导出，并将结果复制到缓冲区。这可以像您希望的那样详细。您的插件应该模拟剪切和粘贴行为，以粘贴由AE附带的插件显示的参数设置的文本表示（例如，到Microsoft Excel电子表格中）。您在如何格式化输出方面有很大的灵活性。 |
| `PF_Arbitrary_SCAN_FUNC`  | 给定一个文本数据缓冲区（通常来自系统剪贴板），将其解析为任意数据格式。|

## 实现任意数据

除了正常的命令和事件入口指令外，arb 数据还需要另一套主机交互。这对其他参数类型来说是透明的，因为 After Effects 管理着它们的数据。编写arb数据插件将使您深入了解AE执行的大量参数管理，以及这些管理发生的顺序。甚至可能导致重新思考您的实现，并使用AE为您管理的参数类型。

实例化你的 arb 数据(当然，使用 After Effects 的内存分配功能)，并将 `ParamDef.u.arb_d.dephault` 指向它。用适当的默认值来填充它。设置参数时不需要数值变量，为安全起见，将其归零。

在插件的入口函数中，包括一个处理[PF_Cmd_ARBITRARY_CALLBACK](../effect-basics/command-selectors.html) 的案例。

调用一个二级事件处理程序，`HandleArbitrary`。它接收一个额外的`PF_ArbParamsExtra`，它又包含一个`PF_FunctionSelector`，用来识别发送的命令。

也许 After Effects 发送了`PF_Cmd_ARBITRARY_CALLBACK`，`PF_FunctionSelector`是`PF_Arbitrary_COPY_FUNC`。在`PF_ArbParamsExtra.copy_func_params`中提供了指向源和目标 Arb 的指针。分配一个新的 Arb，并将`dest_arbPH`指向它。如果`src_arbH`是 NULL，为`dest_arbPH`创建一个默认的 Arb。

用户可以在时间轴面板中选择 arb 的关键帧数据，复制它，然后切换到另一个应用程序。你将收到一个`PF_Arbitrary_PRINT_SIZE_FUNC`；通过在`PF_ArbParamsExtra`中设置`print_sizePLu`来设置输出缓冲器的大小。然后你会收到`PF_Arbitrary_PRINT_FUNC`；用有关Arb的文本表示来填充`print_bufferPC`输出缓冲区。

用户可以将关键帧数据粘贴到你的 Arb 的时间轴上。你将收到`PF_Arbitrary_SCAN_FUNC`。根据交给你的字符缓冲区的内容创建一个 Arb(其大小在`print_sizeLu`中显示)。

## 任意数据？重入

你的插件代码必须是递归可重入(Re-Entrancy)的，以支持自定义的数据类型，因为AE可以出于多种原因调用它。您的插件可以检出一个图层，而图层反过来又依赖于您的效果的另一个实例。您尝试检出一个（看似）不相关的图层将触发插件的任意数据处理代码。注意对依赖于通过全局变量访问的静态值的C运行时库的调用。如果您没有为这种可能性做好准备，此时会挂起AE，用户会疯掉并狂拍他们的显示器。

## 何时不访问任意参数

如果`in_data>effect_ref`是`NULL`，就不要检出任意的参数。

## 在绘画期间修改

After Effects 忽略了在`PF_Cmd_DO_DIALOG`期间对任意数据参数的任何改变。

这是设计好的；在显示选项对话框期间的改变会影响整个效果流，而不仅仅是某一时刻的任意参数。

如果你必须根据这些变化改变你的 Arb 的行为，请将该信息保存在序列数据中并在以后应用，通常是在`PF_Cmd_USER_CHANGED_PARAM`期间。
