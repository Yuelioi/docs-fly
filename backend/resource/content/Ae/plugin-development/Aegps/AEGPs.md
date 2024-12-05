---
title: AEGPs
order: 3
---

# AEGPs

After Effects General Plug-in (AEGP)的 API 强大而广泛，提供的功能超出了效果插件范围。

对用户来说，AEGPs 似乎是 After Effects 的一部分。

它们可以添加、拦截和触发菜单命令，访问关键帧数据库，并将功能注册为 After Effects 内部用于信息传递的一部分。

AEGPs 可以在项目和合成中添加和删除项目，添加和删除过滤器(filters)和关键帧。

一旦其命令被触发，AEGPs 就会使用众多的 PICA 功能套件(在本节中描述)来处理每一个 After Effects 项目。

AEGPs 可以为插件发布功能套件，操纵所有项目元素，改变 interpretations，替换文件，并确定哪些外部文件用于渲染项目。

有几种专门的 AEGP 类型；Keyframers、Artisans 和 I/O 模块(AEIOs)。它们都是 AEGP，但可以访问专门的信息流，为此在 After Effects 中单独注册。
