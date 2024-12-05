---
title: Overview
order: 7
---

# 预览

AEGPs 使用插件组件架构(PICA)功能套件来访问所有功能。

他们也可以发布自己的功能套件，供特效插件使用(由于插件的加载顺序不同，AEGPs 不能依赖 After Effects 没有提供的套件)。

AEGPs 也可以请求一个套件，如果它不存在，就自己提供替代功能。

## AEGP 与 AE 通信

对于效果插件来说，所有与 After Effects 的通信都是通过一个入口函数进行的。而 AEGP 则不是这样的。

虽然 After Effects 会调用 AEGP 的 PiPL 中指定的入口函数函数(这仍然是必需的)，但 After Effects 和 AEGP 之间的所有后续通信都由 AEGP 注册的钩子函数处理。

这种注册必须在插件的入口函数中进行，使用[注册套件](aegp-suites.html) 。

## 不同工作方式, 同样的API

AEGP 的工作方式是一样的，不分专业。

它们可以很简单，只是[添加一个菜单项](implementation.html) 来触发一个外部应用程序，也可以像 Artisans 那样复杂。

虽然插件可以访问任何函数套件，但只有适当类型的插件才可以访问所有需要的参数。

只有 Artisans 会有渲染上下文，只有 AEIO 插件会收到输入和输出规范；消息传递取决于哪些钩子函数被注册。
