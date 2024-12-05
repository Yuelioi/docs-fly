---
title: Other-Integration-Possibilities
order: 10
---

# 其他整合的可能性

虽然这个 SDK 描述了大部分与 AE 整合的可能性，但还有其他的可能性。

## 脚本

脚本是一种轻量级选择，可用于 AE 执行自动任务。ScriptUI 提供自定义对话框和面板 UI 集成(参见[HTML5 面板])。脚本可以和插件开发一起使用，此时，某个功能可以通过脚本而不是 C 语言 API 来实现。

After Effects 中的脚本基于 JavaScript 的 ExtendScript。After Effects 包括 ExtendScript ToolKit，这是一个创建和测试自己脚本的方便接口。脚本可以被编译成.jsxbin 二进制文件，以保护知识产权。

After Effects 脚本指南以及论坛 <https://www.adobe.io/apis/creativecloud/aftereffects.html>。

可以通过命令行来驱动 After Effects。进行打开项目等动作。比如直接从命令行运行:

`AfterFX -s "app.quit()"`

或者执行一个.jsx 脚本，该脚本包含 `app.quit()` 这段代码。

AfterFX -r 该jsx文件路径

在 Windows 上，AfterFX.com 可以获得控制台反馈，因为它是一个命令行应用程序。

## HTML5 面板

CC 2014 及以后的版本，AE 开始支持 HTML5 面板。窗口 > 拓展 >(你的面板名称)访问拓展。面板可以像 AE 中的其他面板一样调整大小和停靠。使用 HTML5、After Effects Scripting 和 JavaScript 构建。从 Adobe I/O 网站上下载 AE Panel SDK。<https://www.adobe.io/apis/creativecloud/aftereffects.html>

## AERender

与脚本密切相关的是 aerender 提供的命令行界面。aerender 主要用于自动渲染，也可用于命令行执行的任何序列脚本命令。在 AE 的帮助文档中可以看到一个概述。[https://helpx.adobe.com/after-effects/using/automated-rendering-](https://helpx.adobe.com/after-effects/using/automated-rendering-network-rendering.html) [network-rendering.html](https://helpx.adobe.com/after-effects/using/automated-rendering-network-rendering.html)

## Premiere Pro 导入器

Premiere Pro 导入器支持将媒体导入到 Adobe Creative Cloud 的大多数应用程序中，包括 Premiere Pro、Media Encoder、Prelude 和 Audition。由于这种更广泛的兼容性，除非你需要与 AE 非常具体的整合，只能通过本 SDK 中的 AEIO API 来实现，否则我们建议开发一个 Premiere Pro 导入器。Premiere Pro 的 SDK 可以在以下网站找到。[https://www.adobe.io/apis/](https://www.adobe.io/apis/creativecloud/premierepro.html)  [creativecloud/premierepro.html](https://www.adobe.io/apis/creativecloud/premierepro.html)

与 AEIOs 相比，MediaCore 导入器插件的一个优势是它的优先级系统。最高优先级的进口商可以首先导入一个文件，如果特定的进口文件不被支持，那么下一个最高优先级的进口商将有机会尝试导入该文件，以此类推。

## Mercury 传输

Mercury 传输插件用于发送视频到输出硬件，以实现广播质量的监控。传输器在 Adobe Creative Cloud 的大多数应用程序中都得到支持，包括 Premiere Pro、After Effects、Prelude 和 Character Animator。Mercury 传输 API 记录在 Premiere Pro SDK 中，可在以下网站找到。[https://](https://www.adobe.io/apis/creativecloud/premierepro.html) () [www.adobe.io/apis/creativecloud/premierepro.html](https://www.adobe.io/apis/creativecloud/premierepro.html)
