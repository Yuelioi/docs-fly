---
title: Where-Installers-Should-Put-Plug-ins
order: 17
---

# 应该在哪安装插件

将插件安装在公共位置，可以防止被 Premiere Pro 加载。

Windows 可以在以下注册表项中找到公共插件文件夹(作为一个明确的路径)。`HKLM\SOFTWARE\After Effects\[version]\CommonPluginInstallPath`。

Mac 常用插件文件夹在 `/Library/Application Support/Adobe/Common/Plug-ins/[版本]/MediaCore`。

所有 CC 版本的版本被锁定在 7.0，或早期版本的 CSx。例如：`/Library/Application Support/Adobe/Common/Plug-ins/7.0/MediaCore/`。

不要使用 macOS 的别名或 Windows 的快捷方式，因为 Premiere Pro 不会越过这些地方。

## 必须安装在 Common 文件夹么?

你可能有充分的理由只为 After Effects 安装你的插件，例如，如果你的插件依赖于 Premiere Pro 中没有的套件和功能。我们强烈建议你尽可能地使用通用文件夹，但对于某些情况，AE 专用的插件文件夹仍然可用。

Windows 可以在以下注册表项中找到应用程序特定的插件文件夹(作为明确的路径) `\HKEY_LOCAL_MACHINE\SOFTWARE\AdobeAfter Effects\(version)\PluginInstallPath`。

macOS 应用程序特定的插件文件夹在 `/Applications/Adobe After Effects [版本]/Plug-ins/`。

启动 After Effects 时 会递归子目录的 10 层处。以括号结尾的目录或前面有 ¬(macOS)或~(Windows)的目录不会被扫描。

尽管你想在 AE 和 Premiere Pro 之间建立一道隔阂，但用户仍然可以使用可爱的整合功能的方法 -- 你的插件仍然可以被 Premiere Pro 用户使用，只需创建一个动态链接的 AE 合成，并将其放入 Premiere Pro 序列。
