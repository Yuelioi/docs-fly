---
title: How-To-Start-Creating-Plug-ins
order: 6
---

# 如何开始创建插件

## 开整

在写代码前，花些时间学习 After Effects，以及[示例项目](sample-projects.html) 。把插件安在正确的文件夹里。设置断点，阅读有趣且丰富的注释。

关于构建效果的快速入门视频(在 macOS 上)。[adobe.ly/2sjMDwM](https://adobe.ly/2sjMDwM)

## 计划

想清楚想做什么样的插件。

## 模板

在研究示例项目后，找一个功能类似的插件模板。从头开始做听起来很牛, 但对于效果来说，使用 Skeleton 模板项目。通过将代码追加到模板中，避免了重建项目的麻烦(包括 Windows PiPL 资源生成的麻烦的自定义构建步骤)。

## 魔改

要使 Skeleton 模板成为你自己的，复制整个\Skeleton 目录，(比如)把它重命名为\WhizBang。使用文本编辑器，搜索\WhizBang*.*(包括.NET 和 Xcode 项目文件)中出现的 Skeleton 和 SKELETON，并用 WhizBang 和 WHIZBANG 替换。

现在编译并运行该插件，它可以响应常见的命令，处理 8 和 16-bpc 的颜色，使用 AEGP_SuiteHandler 实用代码，并响应 3D 灯光和摄像机信息。就是这么简单

AEGP 开发者最好从 Projector(支持 After Effects 项目的创建)开始，Easy Cheese 是个关键帧助手，IO 是媒体文件格式拓展，Persisto 处理简单的菜单命令和首选项。

## 测试

如果只是为了测试方便，应该保存一个项目，应用你的效果，并将其所有参数的关键帧改为奇怪的值。

## 错误

如果遇到BUG，看看是否用原始样本项目来重现。来确定是你的原因，还是一开始就存在的, 这可以节省很多时间.

## 开发者事宜

第三方开发者推动了 API 和 SDK 的改进和扩展。你们的产品使 After Effects 内容更加丰富。你们的努力使 After Effects 变得更好。

我们在 SDK 上努力工作，并欢迎你的意见和反馈。我们对 API 所做的每一个改动几乎都是由像你这样的开发者提出的。[给我们反馈](https://community.adobe.com/t5/after-effects/bd-p/after-effects?page=1&sort=latest_replies&filter=all&topics=label-sdkcom)
