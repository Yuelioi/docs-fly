---
title: Accessing-the-After-Effects-Function-Suites
order: 2
---

# 访问AE函数套件

如果您正在编写C++代码，则应使用AEFX_SuiteScoper访问套件，该会在需要时自动获取套件，并在完成时将其处理掉。

下面是一个使用AEFX_SuiteScope访问PF_GPUDeviceSuite1套件的示例：

```cpp
AEFX_SuiteScoper<PF_GPUDeviceSuite1> gpu_suite = AEFX_SuiteScoper<PF_GPUDeviceSuite1>(
 in_dataP,
 kPFGPUDeviceSuite,
 kPFGPUDeviceSuiteVersion1,
 out_dataP);
```

::: tip
如果请求的套件不能被获取，并且可选的第二个模板参数，`ALLOW_NO_SUITE`，被设置为 false, AEFX_SuiteScoper 将抛出一个`A_Err_MISSING_SUITE`异常。如果你将`ALLOW_NO_SUITE`设置为false，请确保你用一个try/catch来包装`AEFX_SuiteScoper<>`。如果`ALLOW_NO_SUITE`被设置为 true，那么你应该在使用前检查返回的指针是否为 NULL。
:::

一旦你有了这个套件，就可以对套件列表中的函数进行调用，例如:

```cpp
gpu_suite->GetDeviceInfo(in_dataP->effect_ref, extraP->input->device_index, &device_info);
```

如果必须使用C代码，则使用`PF_Suite_Helper`实用程序文件手动获取和释放套件，如Checkout示例项目中所示。

在幕后，这两种方法都使用`AcquireSuite`获取 PICA 函数套件，这是`PF_InData`中指向的`SPBasicSuite`的一个成员函数。

## 套件版本

WhizBangSuite1 可以提供一个 Foobar()函数，它需要两个参数，而 WhizBangSuite2>Foobar()可能需要三个。虽然一个套件的每个新版本都会取代旧版本，但请随意获取同一套件的多个版本；我们从不删除或改变以前版本的套件。

当不确定插件主机的功能时(除了 Premiere 之外，没有第三方主机支持 PICA)，尝试获取最新的版本，并 "回退"到以前的版本。如果你需要的功能不可用，就向用户发出警告，并返回一个错误(或者在更"原始"的插件主机中运行时,会返回到其他行为)。请注意，在其他 After Effects 插件主机中对这些套件的支持都不太顺利。

## 线程

除非另有记录，套件提供的任何函数都不是线程安全的。例如，只有你插件的主线程应该做任何修改用户界面的事情。
