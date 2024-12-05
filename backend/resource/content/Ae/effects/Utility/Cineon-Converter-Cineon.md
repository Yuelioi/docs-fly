---
title: Cineon 转换器
order: 3
---

# Cineon Converter - Cineon 转换器

## 简述

[以下来自官方文档](https://helpx.adobe.com/cn/after-effects/using/utility-effects.html)。感觉是我不熟悉的领域，就不妄自评价了

“Cineon 转换器”效果提供对 Cineon 帧的颜色转换的高级别控制，但大多数情况下，您应使用色彩管理功能转换 Cineon 文件中的颜色。（请参阅
[Cineon 和 DPX 素材项目](https://helpx.adobe.com/cn/after-effects/using/preparing-importing-still-images.html#cineon_and_dpx_footage_items)。）

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

要使用“Cineon 转换器”效果，请导入 Cineon 文件并将其保留为默认状态；After Effects 会将颜色压缩为 8 bpc 或将其扩大到
16 bpc 或 32 bpc，具体取决于您使用的模式。然后您可以将“Cineon
转换器”效果应用于文件，并在“合成”面板中以交互方式查看结果时精确地调整颜色。您可以设置关键帧或表达式，随着时间的推移针对色调的变化做出调整 -
使用关键帧插值和缓和手柄以精确地匹配最无规律的光照变化。可考虑改用其他“颜色校正”效果来执行这些校正。（请参阅[颜色校正效果](https://helpx.adobe.com/cn/after-effects/using/color-correction-effects.html#color_correction_effects)。）

每个 Cineon 渠道中可用于每个像素的 10
位数据使您能够更轻松地增强重要的色调范围，同时保持总体色调平衡。通过准确指定范围，您可以创建真实地重现原始图像的图像版本。

**注意**

每个计算机显示器都有独特的显示特征，这些特征会影响您对颜色的感受。要在评估色调平衡时获得最佳效果，请在将指针移动到像素上时使用 After Effects
中的“信息”面板查看像素的真实颜色值。

## 效果展示

![](https://cdn.yuelili.com/20220102214757.png)

## 教程

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Utility-Cineon_Converter.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Utility-Cineon_Converter_cn.png)

| Cineon Converter | Cineon 转换器 | Cineon コンバーター |                      |           |                      |
| ---------------- | ------------- | ------------------- | -------------------- | --------- | -------------------- |
|                  |               |                     | Conversion Type      | 转换类型  | 変換の種類           |
|                  |               |                     |                      |           |                      |
|                  |               |                     |                      |           |                      |
|                  |               |                     | 10 Bit Black Point   | 10 位黑场 |                      |
|                  |               |                     | Internal Black Point | 内部黑场  |                      |
|                  |               |                     | 10 Bit White Point   | 10 位白场 |                      |
|                  |               |                     | Internal White Point | 内部白场  | 内部ホワイトポイント |
|                  |               |                     | Gamma                | 灰度系数  | ガンマ               |
|                  |               |                     | Highlight Rolloff    | 高光滤除  | ハイライトロールオフ |

## 参数详解

### 转换类型

转换 Cineon 文件的方式。“对数到线性”可转换您计划作为 Cineon 序列渲染的 8-bpc 对数非 Cineon 图层。“线性到对数”可将
Cineon 文件的包含 8-bpc 线性代理的图层转换为 8-bpc 对数文件，以便其显示特征与原始 Cineon 文件一致。“对数到对数”可在您计划将
8 bpc 或 10 bpc 对数 Cineon 文件渲染为 8-bpc 对数代理时对 Cineon 文件进行检测。

### 10 位黑场

用于转换 10-bpc 对数 Cineon 图层的黑场（最低密度）。

### 内部黑场

用于 After Effects 中的图层的黑场。

### 10 位白场

用于转换 10-bpc 对数 Cineon 图层的白场（最高密度）。

### 内部白场

用于 After Effects 中的图层的白场。

### 灰度系数

分别增大或减小灰度系数以使中间调变亮或变暗。

### 高光滤除

用于校正明亮高光的滤除值。如果调整最亮区域使图像的其余部分看起来太暗，请使用“高光滤除”来调整这些明亮高光。如果高光显示为白斑，请增大“高光滤除”，直到细节可见。具有高对比度的图像可能需要高滤除值。

## 案例
