---
order: 11
title: ppro
---

# Premiere Pro & 其他主机

Adobe Premiere Pro 和 Adobe Premiere Elements 都支持第 2、3、5 章中描述的 After Effects 插件API。

它们提供了一个完整的主机实现，一些关键的遗漏是 3D 相关的调用(辅助通道信息、摄像机和灯光)、16 位和 SmartFX 支持，以及 After Effects 的 AEGP API 提供的其他实用功能。

Premiere Pro 和 Premiere Elements 都将`PF_InData>appl_id`设置为'PrMr'。

在本章中，我们将介绍 Premiere Pro 中的 AE API 支持，但一般来说，Premiere Elements 的相应版本中也存在同样的支持。

如果你需要区分 Premiere Pro 和 Premiere Elements，你可以使用 Premiere 特定的 App Info Suite，可从[Premiere Pro SDK](http://ppro-plugin-sdk.aenhancers.com)标题中获得。

| 应用版本 | PF_InData> version.major | PF_InData> version.minor |
| -------------------------------------------- | ---------------------------- | ---------------------------- |
| Premiere Pro CC through Premiere Pro CC 2019 | 13 | 4  |
| Premiere Pro CS6 | 13 | 2  |
| Premiere Pro CS5.5 | 13 | 1  |
| Premiere Pro CS5, Premiere Elements 9  | 13 | 0  |
| Premiere Pro CS4, Premiere Elements 8  | 12 | 5  |
| Premiere Pro CS3, Premiere Elements 4 and 7  | 12 | 4  |
| Premiere Pro 2.0, Premiere Elements 3  | 12 | 3  |
| Premiere Pro 1.5, Premiere Elements 2  | 12 | 2  |
| Premiere Pro 1.0, Premiere Elements 1  | 12 | 1  |

请注意，Premiere Pro 和 Premiere Elements 使用的版本划分并不意味着它们支持 After Effects 在同一版本中的相同 API 功能。只是为了区分不同的版本。
