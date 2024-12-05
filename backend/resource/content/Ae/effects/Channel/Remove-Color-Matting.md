---
order: 9
title: 移除颜色遮罩
---

# Remove Color Matting - 移除颜色遮罩

## 简述

去除在合成图像的边缘产生的颜色边缘。

移除颜色遮罩效果可从带有预乘颜色通道的图层移除色边（色晕）。在部分透明的区域保留原始背景的颜色，并使用其他背景颜色合成到时，通常会出现色晕。将此效果与创建透明度的效果（如抠像效果）结合使用可增强对部分透明的区域外观的控制。（请参阅
[Alpha 通道解释：预乘或直接](https://helpx.adobe.com/cn/after-effects/using//importing-interpreting-footage-items.html#alpha_channel_interpretation_premultiplied_or_straight)。）

此效果适用于 8-bpc、16-bpc 和 32-bpc 颜色。

## 效果展示

去除红色颜色边

![](https://mir.yuelili.com/user/AE/effects/ext/image00412.jpg)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Remove_Color_Matting.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Remove_Color_Matting_cn.png)

| Remove Color Matting | 移除颜色遮罩 | カラーマット削除 |                  |               |                    |
| -------------------- | ------------ | ---------------- | ---------------- | ------------- | ------------------ |
| Background Color     | 背景颜色     | 背景色           |                  |               |                    |
| Clipping             | 剪切         | クリッピング     |                  |               |                    |
|                      |              |                  | Clip HDR Results | 剪切 HDR 结果 | HDR 効果をクリップ |

## 参数详解

The background color：用颜色表或滴定工具指定背景色。

剪切：高动态范围效果：在处理 32bpc 色的图像时，在生成 0.0~1.0 的范围以外的色的情况下，除去[剪切]的 HDR 效果。

## 案例

当史蒂夫需要使用在黑色背景上拍摄的火或烟的片段时，他可以使用[Shift
Channels 效果](http://help.adobe.com/en_US/AfterEffects/9.0/WS3878526689cb91655866c1103a9d3c597-7b8da.html)和[Remove
Color
Matting 效果](http://help.adobe.com/en_US/AfterEffects/9.0/WS3878526689cb91655866c1103a9d3c597-7b90a.html)。使用前者从图像的亮度或红色通道中获取 Alpha 通道值。当然，由于黑色区域既不亮度也不红色，因此将黑色区域中的 alpha 通道值设置为 0。换句话说，它可以消除黑色。使用“消除颜色遮罩”效果可消除部分透明区域（其中有许多起火或冒烟的地方）黑色背景的剩余痕迹。

可以使用多种方法来合成具有黑色背景的剪辑，例如使用 Knoll
Unmult（最适合在具有光效果的剪辑（例如火光和耀斑）中抠出黑色）之类的东西，或使用诸如 Screen
的[混合模式](http://help.adobe.com/en_US/AfterEffects/9.0/WSFB0FE760-71F3-4616-AE88-275D718E7125a.html)。

当您以多种不同方式使用同一资产时，Steve 的技术似乎比使用混合模式更具通用性，因为它实际上会创建 alpha 信息。这是在腰带中使用的好工具。
