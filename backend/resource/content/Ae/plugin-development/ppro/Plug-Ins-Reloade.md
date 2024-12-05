---
order: 7
title: Plug-Ins-Reloaded
---

# 插件重载

## 插件重载

在第一次启动时，Premiere Pro 会加载所有的插件，读取 PiPL，并发送 `PF_Cmd_GLOBAL_SETUP` 来确定插件的能力。为了节省以后启动应用程序的时间，它将这些功能中的一部分保存在我们所说的插件缓存中(Windows 的注册表，macOS 的属性列表文件)。下次启动应用程序时，尽可能使用缓存的信息，而不是重新加载插件。

调试时，你可以在启动 Premiere Pro 时按住 Shift 键，强制重新加载所有的插件。

如果效果每次都需要重新加载，有一种方法可以禁用插件缓存。可以使用 AE_CacheOnLoadSuite.h ( 参见[Premiere Pro SDK](http://ppro-plugin-sdk.aenhancers.com/)头文件)中的 `PF Cache On Load` 套件，在 `PF_Cmd_GLOBAL_SETUP` 期间调用 `PF_SetNoCacheOnLoad()`。对于该函数的第二个参数，如果想让你的效果显示在用户界面上，请传一个非零值。如果加载失败，但仍然希望 Premiere Pro 在下一次重新启动时尝试再次加载，则传0。
