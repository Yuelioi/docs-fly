---
title: Exporting-Symbols-in-Effects
order: 5
---

# 在效果中导出符号

After Effects 团队最近意识到了一个符号冲突的问题，它违反了 C++语言的一个定义规则(ODR)。

2021 年初，After Effects 使用的 Boost 库的版本升级到 1.74。在过去的几个月里，我们发现有一些插件也在使用 Boost，但它们输出符号的方式让 After Effects 或插件最终调用了不正确的 Boost 版本，导致用户挂起和崩溃。我们还发现一些 AE SDK 样本被设置为默认导出所有符号，这可能是造成问题的原因，假设它们被用作其他插件的起点。这些问题已经在 2021 年 3 月的 SDK 中得到修复。

After Effects 要求导出的唯一符号是插件的入口函数。

可以在 SDK 样本中的 entry.h 中找到示例。

```cpp
#ifdef AE_OS_WIN
 #define DllExport __declspec( dllexport )
#elif defined AE_OS_MAC
 #define DllExport __attribute__ ((visibility ("default")))
#endif

```

然后将其应用到入口函数函数中，例如。

```cpp
extern "C" DllExport
PF_Err PluginDataEntryFunction(
 PF_PluginDataPtr inPtr,
 PF_PluginDataCB inPluginDataCallBackPtr,
 SPBasicSuite* inSPBasicSuitePtr,
 const char* inHostName,
 const char* inHostVersion)
{
 PF_Err result = PF_Err_INVALID_CALLBACK;

 result = PF_REGISTER_EFFECT(
 inPtr,
 inPluginDataCallBackPtr,
 "ColorGrid", // Name
 "ADBE ColorGrid", // Match Name
 "Sample Plug-ins", // Category
 AE_RESERVED_INFO); // Reserved Info

 return result;
}

```

## 禁用Xcode符号导出

要在 Xcode 中禁用符号导出。

1. 在你项目的**Build**设置中找到**Apple Clang - Code Generation**部分。
2. 将**符号默认隐藏**设置为**YES**。

![images/appleclang-symbols.png](/images/appleclang-symbols.png)

对于任何必须公开的特定符号，在代码中使用`__attribute__((visibility("default")))`。

更多的信息可以在苹果的 Xcode 文档<https://help.apple.com/xcode/mac/11.4/#/itcaec37c2a6>中找到(摘录如下)。

```cpp
Symbols Hidden by Default (GCC_SYMBOLS_PRIVATE_EXTERN)

When enabled, all symbols are declared private extern unless explicitly marked to be exported using __attribute__((visibility("default"))) in code. If not enabled, all symbols are exported unless explicitly marked as private extern.

```

## 禁用 Visual Studio 导出符号

默认情况下，来自 Visual Studio 的构建会自动禁用符号导出。要导出符号，你必须提供一个模块定义文件或在函数定义中设置 declspec(dllexport) 关键字。

更多信息可以在微软的 Visual Studio 文档<https://docs.microsoft.com/en-us/cpp/build/exporting-from-a-dll?view=msvc-160>中找到(摘录如下)。

```cpp
You can export functions from a DLL using two methods:
  1. Create a module definition (.def) file and use the .def file when building the DLL. Use this approach if you want to export functions from your DLL by ordinal rather than by name.
  2. Use the keyword __declspec(dllexport) in the function's definition.

When exporting functions with either method, make sure to use the __stdcall calling convention.

```
