---
title: Unsupported-Features
order: 10
---

# 不支持的特性

目前已知 Premiere Pro 不支持 After Effects API 的以下功能。

(如果你想要一个带"-"的功能，请给[Premiere Pro API Engineering](mailto:bbb%40adobe.com)发邮件，提出功能请求。前面有 "F "的数字是功能请求编号，其他的是错误编号)

- F7233 - 支持 extent_hint
- F7835 - 一个插件中的多个 PiPLs
- F7836 - 支持 AEGP
- F7517 - 音频支持 - 如果一个插件在 PF_Cmd_GLOBAL_SETUP 中设置了 PF_OutFlag_I_USE_AUDIO，它将完全不被加载
- F9355 - 支持 PF_ParamFlag_COLLAPSE_TWIRLY
- PF World Transform Suite
- PF AE 通道套件
- AE 对高位色深支持的实现
- SmartFX
- 3D 支持
- pf_subpixel_sample(), pf_get_pixel_data16()

## 为何要加载而不是运行它

Premiere Pro 会尝试加载 AEGP 插件。为了检测并避免任何问题行为，你的命令钩(command hook)函数可以访问一个只由 After Effects 提供的套件；`AEGP_CanvasSuite` 是一个很好的候选。

如果该套件不存在，则返回错误。该插件将被放在 Premiere Pro 的 "不加载这些 "列表中。
