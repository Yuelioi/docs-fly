---
title: Localization
order: 8
---

# 本地化

从 CC 开始，PF 应用套件([实用功能](../effect-details/useful-utility-functions.html) )增加了`PF_AppGetLanguage()`来查询当前语言，这样就可以使用正确的语言字符串。

向 AE 传递字符串时，API 的某些部分接受 Unicode。在其他地方，例如在`PF_Cmd_PARAM_SETUP`期间指定效果参数名称时，需要用 char 字符串传递名称。对于非 Unicode 字符串，AE 将字符串解释为使用应用程序的当前区域码进行多字节编码。要建立这些字符串，在 Windows 上可以使用`WideCharToMultiByte()`函数，指定`CP_OEMCP`作为第一个参数。在 macOS 上，使用由`GetApplicationTextEncoding()`返回的编码。

在 AE 中用不同的语言测试不需要重新安装操作系统，但需要重新安装 AE。

Win 系统可将系统区域设置改为目标语言(控制面板>区域和语言>管理标签>改变系统区域设置)，重新启动机器，然后用相应的语言安装 AE。

Mac 可将目标语言设置为首选语言列表中的主要语言，然后用相应的语言安装 AE。
