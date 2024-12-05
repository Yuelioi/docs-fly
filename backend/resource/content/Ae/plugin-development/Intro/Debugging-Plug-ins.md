---
title: Debugging-Plug-ins
order: 3
---

# 调试插件

学习 After Effects 和插件之间交互的最好方法是在调试器中运行样本。花一些时间在编译器的调试器上，研究类似的样本项目上，会有很大的收获。

按照上面的解释把插件直接构建到插件文件夹中，下面是如何指定 After Effects 作为调试过程中运行的应用程序。

在 Windows 上。

1. 在 Visual Studio 解决方案中，在解决方案资源管理器面板中，选择你要调试的项目。
2. 右键单击它，选择设置为启动项目
3. 再次右击它，并选择属性
4. 在 "配置属性">"调试">"命令 "中，提供插件将在主机应用程序中运行的可执行文件的路径(After Effects 或 Premiere Pro)。
5. 接着点击运行按钮，或者你可以启动应用程序，然后在任何时候选择 调试 > 附加到进程......

在 macOS 上。

1. 在 Xcode 中，在项目导航器中，选择你要调试的 xcodeproj
2. 选择产品 > 方案 > 编辑方案...
3. 在 "运行 "下，在 "信息 "选项卡中，为 "可执行 "选择插件将运行的主机程序(After Effects 或 Premiere Pro)
4. 接着点击运行按钮来建立和运行当前的方案，或者你可以启动应用程序，然后在任何时候选择调试>附加到进程。

## 删除首选项

在开发插件的过程中，你的插件可能会将设置信息传递给 After Effects，然后存储在其偏好文件中。

你可以在启动时按住 Ctrl-Alt-Shift/Cmd-Opt-Shift，删除这些偏好，以干净的方式重新启动 After Effects。

在 Windows 上，偏好设置存储在`[用户文件夹]/AppData\Roaming\AdobeAfter Effects\[版本]\Adobe After Effects [版本]-x64 Prefs.txt`。

在 macOS 上，存储在`~/Library/Preferences/Adobe/After Effects/[版本]/Adobe After Effects [版本]-x64 Prefs`。
