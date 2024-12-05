---
title: Apple-Silicon-Support
order: 1
---

# Apple Silicon 支持

Adobe 现在支持一些在苹果上原生运行的产品中的效果插件。例如，After Effects 的效果插件在 Adobe Premiere Pro 和 Adobe Media Encoder 中也可用。

并非所有的 Adobe 产品都有原生的 Apple Silicon 版本，但在那些产品中，只有具有 Apple Silicon 实现的效果插件才可用。我们建议尽快添加 Apple Silicon 目标，以期待这些新的 M1 机器的快速采用。

::: tip

为了建立一个 Mac 通用二进制文件，需要 Xcode 12.2 或更高版本。Adobe 目前正在使用 Xcode 12.4。
:::

要了解更多关于通用二进制文件的信息，请访问<https://developer.apple.com/documentation/apple-silicon/building-a-universal-macos-binary>。

## How to add Universal Binary Support for your Plugins

- 1.在 12.2 或以上版本中打开您的插件 Xcode 项目，Xcode 将自动为您添加一个 Apple Silicon 目标。

![/images/mac_universal_build.png](/images/mac_universal_build.png)

- 2.告诉 After Effects 的苹果硅构建的主要入口函数是什么。

找到你的插件的.r 资源文件。

在现有的 Intel Mac 入口函数定义旁边添加`CodeMacARM64 {"EffectMain"}`。

```cpp
#if defined(AE_OS_MAC)
CodeMacARM64 {"EffectMain"},
CodeMacIntel64 {"EffectMain"},
#endif
```

如果由于某种原因，在 x64 和 ARM 上需要不同的入口函数，只需提供不同的入口函数名称和字符串。

- 3.通过为任何 Mac(Apple Silicon, Intel)目标构建，或通过使用产品->存档，编译通用二进制文件。

假设 Apple Silicon 的编译时间没有问题，你现在可以将单一的通用二进制文件用于 Intel 和 Apple Silicon 应用程序。

## 苹果跨 “C” 函数的异常行为

在 Apple Silicon 上使用异常时应特别小心。在许多环境中，抛出通过传统 "C "函数传播的异常，效果很好。这是不好的做法，有未定义的行为，但通常是 "有效的"。

在 Apple Silicon 上，ABI 已经改变了，所以当这种情况发生时，会调用 terminate()，而不是未定义的行为。

由于一个插件的主要入口函数总是一个外部 "C "的调用约定，这段代码应该被包裹在一个 try/catch 块中，以防止程序终止。比如。

```cpp
PF_Err EffectMain ( PF_Cmd cmd,
 PF_InData *in_data,
 PF_OutData *out_data,
 PF_ParamDef *params[],
 PF_LayerDef *output )
{
 try
 {
 /* 你的代码 */
 }
 catch
 {
 /* 返回最合适的 PF_Err */
 }
}

```
