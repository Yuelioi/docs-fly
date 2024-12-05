---
title: Cheating-Effect-Usage-of-AEGP-Suites
order: 4
---

# AEGP套件的作弊效果使用

当我们刚向开发者展示了 AEGP 套件时，有人就想 "作弊"，从效果中使用它们。这当然可以，但请记住，依赖效果 API 以外的因素(即从 AEGP APIs 获得的任何信息)也许会产生麻烦。如果 After Effects 认为一个效果已经有了它所需要的所有信息，它就不会(例如)根据 AEGP 函数产生的变动来更新其参数。我们正在为未来版本积极解决这个依赖性问题，但在编写 "伪装成 "AEGP 的效果时要记住这一点。

效果可以使用一些 AEGP 套件来利用相机和灯光信息，以及 "AEGP*GetLayerParentComp "和 "AEGP_GetCompBGColor "函数，来自[AEGP_CompSuite11](aegp-suites.html)(#aegps-aegp-suites-aegp-compsuite)。这不应该被解释为效果可以使用*任何AEGP 套件调用。另外，请参阅[Effect UI &amp; Events](../effect-ui-events/effect-ui-events.html) 以了解更多关于特效添加关键帧的信息。

[AEGP_PFInterfaceSuite](aegp-suites.html) 是一个起点。这个套件中的函数允许你检索应用效果的层的 AEGP_LayerH，以及你的效果实例的 AEGP_EffectRefH。`AEGP_RegisterWithAEGP`来自[AEGP_UtilitySuite6](aegp-suites.html) 允许你获得一个 AEGP_PluginID，许多 AEGP 调用都需要它。

## 基于 AEGP 查询

一个字。别。效果不能让 AEGP 的查询结果来控制渲染的内容，而不适当地存储这些查询结果(通常是在序列数据中)，取消自己的渲染，并强制使用查询信息进行重新渲染。

这是很棘手的。

如果失败，将导致讨厌的、微妙的缓存错误，别跟头发过不去。
