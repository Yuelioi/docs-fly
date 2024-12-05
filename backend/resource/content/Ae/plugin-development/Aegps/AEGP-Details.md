---
title: AEGP-Details
order: 1
---

# AEGP 细节

## 拥有 Cookie

在 After Effects 必须保留关于你的 AEGP 调用的函数状态信息的情况下(如当一个 Artisan 正在渲染一帧，或者一个 keyframer 正在从同一个流中添加和删除一系列关键帧)，你将调用 begin()和 end()函数。

通常，begin 函数将返回一个不透明的标识符，或 "cookie"，然后必须将其传递给正在使用的函数。end 函数会处理这个 cookie。参见 `AEGP_StartAddKeyframes()` [AEGP_KeyframeSuite3](aegp-suites.html)。

## 修改渲染队列中的项目

如果你调用 `AEGP_AddCompToRenderQueue` ( 参见[AEGP_RenderQueueSuite1](aegp-suites.html) )，或者如果用户手动从渲染队列中添加或删除一个合成，所有对渲染队列项目的引用都会被废止。同样地，添加或删除输出模块也会使每个渲染队列项目的任何此类引用失效。

## 名称与固态层

固态层在 After Effects UI 中是有名字的，但在它们的 "PF_LayerDef"[PF_EffectWorld / PF_LayerDef](../effect-basics/PF_EffectWorld.html) 中没有。因此，它们的名字不能被 `AEGP_GetItemName`(in [AEGP_ItemSuite9](aegp-suites.html) ) 或 `AEGP_GetLayerName`(in [AEGP_LayerSuite8](aegp-suites.html) 检索出来。)

然而，你可以使用与它们相关的 ItemH 来 `AEGP_GetItemName` ( 参见[AEGP_ItemSuite9](aegp-suites.html) )。

## 报告错误和问题

使用 `AEGP_ItemSuite>AEGP_ReportInfo()`来向用户报告信息，并识别你的插件。AEIO 插件使用他们被传递的 AEIO_BasicData 中包含的 msg_func 指针来代替(与每个函数一起)。

## 转换：首先发生什么？

After Effects 根据自动定向(朝向路径或兴趣点)计算旋转，然后计算定向，再计算 X、Y、Z 旋转。

## 从效果层参数访问像素

使用 `AEGP_GetNewStreamValue`(in [AEGP_StreamSuite5](aegp-suites.html) )来获取图层的 `layer_id`，然后使用新的 `AEGP_GetLayerFromLayerID`(in [AEGP_LayerSuite8](aegp-suites. html)(#aegps-aegp-suites-aegp-layersuite))来获得`AEGP_LayerH'。
