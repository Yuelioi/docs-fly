---
order: 3
title: 计算
---

# Calculations - 计算

## 简述

将 RGB 或阿尔法频道和其他层指定的频道进行合成。[转换通道]和[通道设置]也能得到同样的效果，不过，在图像简单地合成分层通道的情况下，该效果更有效。

此效果适用于 8-bpc 和 16-bpc 颜色。

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/image00424.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=115&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Calculations.png)

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Calculations_cn.png)

| Calculations  | 计算     | 計算           |                             |                      |                                        |                     |              |                          |
| ------------- | -------- | -------------- | --------------------------- | -------------------- | -------------------------------------- | ------------------- | ------------ | ------------------------ |
| Input         | 输入     | 入力           |                             |                      |                                        |                     |              |                          |
|               |          |                | Input Channel               | 输入通道             | 入力チャンネル                         |                     |              |                          |
|               |          |                |                             |                      |                                        | RGBA                | RGBA         |                          |
|               |          |                |                             |                      |                                        | Gray                | 灰色         |                          |
|               |          |                |                             |                      |                                        | Red                 | 红色         | red                      |
|               |          |                |                             |                      |                                        | Green               | 绿色         | green                    |
|               |          |                |                             |                      |                                        | Blue                | 蓝色         | blue                     |
|               |          |                |                             |                      |                                        | Alpha               | Alpha        | アルファ                 |
|               |          |                | Invert Input                | 反转输入             | 入力を反転                             |                     |              |                          |
| Second Source | 第二个源 | 2 つめのソース |                             |                      |                                        |                     |              |                          |
|               |          |                | Second layer                | 第二个图层           | 2 つめのレイヤー                       |                     |              |                          |
|               |          |                | Second Layer Channel        | 第二个图层通道       | 2 つめのレイヤーチャンネル             | RGBA                |              |
|               |          |                |                             |                      |                                        |                     | 灰色         |                          |
|               |          |                |                             |                      |                                        |                     | 红色         | red                      |
|               |          |                |                             |                      |                                        |                     | 绿色         | green                    |
|               |          |                |                             |                      |                                        |                     | 蓝色         | blue                     |
|               |          |                |                             |                      |                                        |                     | Alpha        | アルファ                 |
|               |          |                | Second Layer Opacity        | 第二个图层不透明度   | 2 つめのレイヤーの不透明度             |                     |              |
|               |          |                | Invert Second Layer         | 反转第二个图层       | 2 つめのレイヤーを反転                 |                     |              |
|               |          |                | Stretch Second Layer to Fit | 伸缩第二个图层以适合 | 2 つめのレイヤーをサイズに合わせて伸縮 |                     |              |
|               |          |                | Blending Mode               | 混合模式             | 描画モード                             | Normal              | 正常         | 標準                     |
|               |          |                |                             |                      |                                        | Copy                | 复制         | コピー                   |
|               |          |                |                             |                      |                                        | Darken              | 变暗         | 暗く                     |
|               |          |                |                             |                      |                                        | Multiply            | 相乘         | 乗算                     |
|               |          |                |                             |                      |                                        | Color Burn          | 颜色加深     | 焼き込みカラー           |
|               |          |                |                             |                      |                                        | Classic Bolor Burn  | 传统颜色加深 |                          |
|               |          |                |                             |                      |                                        | Add                 | 相加         | 加算                     |
|               |          |                |                             |                      |                                        | Lighten             | 变亮         | 明るく                   |
|               |          |                |                             |                      |                                        | Screen              | 滤色         | スクリーン               |
|               |          |                |                             |                      |                                        | Color Dodge         | 颜色减淡     | 覆い焼きカラー           |
|               |          |                |                             |                      |                                        | Classic Color Dodge | 传统颜色减淡 |                          |
|               |          |                |                             |                      |                                        | Overlay             | 叠加         | オーバーレイ             |
|               |          |                |                             |                      |                                        | Soft Light          | 柔光         | ソフトライト             |
|               |          |                |                             |                      |                                        | Hard Light          | 强光         | ハードライト             |
|               |          |                |                             |                      |                                        | Linear Light        | 线性光       | リニアライト             |
|               |          |                |                             |                      |                                        | Vivid Light         | 亮光         | ビビッドライト           |
|               |          |                |                             |                      |                                        | Pin Light           | 点光         | ピンライト               |
|               |          |                |                             |                      |                                        | Difference          | 差值         | 差                       |
|               |          |                |                             |                      |                                        | Classic Difference  | 传统差值     |                          |
|               |          |                |                             |                      |                                        | Exclussion          | 排除         |                          |
|               |          |                |                             |                      |                                        | Hue                 | 色相         | 色相                     |
|               |          |                |                             |                      |                                        | Saturation          | 饱和度       | 彩度                     |
|               |          |                |                             |                      |                                        | Color               | 颜色         | カラー                   |
|               |          |                |                             |                      |                                        | Luminosity          | 发光度       | 輝度                     |
|               |          |                |                             |                      |                                        | Stencil Alpha       | 横版 Alpha   |                          |
|               |          |                |                             |                      |                                        | Stencil Luma        | 横版亮度     |                          |
|               |          |                |                             |                      |                                        | Sihouette Alpha     | 轮廓 Alpha   |                          |
|               |          |                |                             |                      |                                        | Sihouette Luma      | 轮廓亮度     | シルエットルミナンスキー |
|               |          |                |                             |                      |                                        | Alpha Add           | Alpha 添加   |                          |
|               |          |                |                             |                      |                                        | Luminescent Add     | 冷光添加     |                          |

## 参数详解

![](https://mir.yuelili.com/user/AE/effects/ext/image00425-1.jpg)

结尾看推导过程

### Input Channel 输入通道

要提取的通道，可用作混合运算的输入通道。RGBA
通常会显示所有通道。“灰色”通道用于将像素的所有颜色通道值转换为原始像素的明亮度值。“红色”、“绿色”或“蓝色”通道用于将颜色值按所选颜色通道的值计算。Alpha
通道用于将所有通道转换为原始像素 Alpha 通道的值。

### Invert Input 反转输入

在效果提取指定通道信息之前反转图层（从 1.0 中减去每个通道值）。

### Second layer 第二个图层

与原始图层与混合的控件图层。

### Second Layer Channel 第二个图层通道

与输入通道混合的通道。

Second Layer Opacity 第二个图层不透明度

第二个图层的不透明度。如果为 0% ，则没有任何效果。

### Invert Second Layer 反转第二个图层

提取指定通道信息之前反转第二个图层（从 1.0 中减去每个通道值）。

### Stretch Second Layer to Fit 伸缩第二个图层以适合

在混合前将第二个图层伸缩到原始图层的尺寸。取消选择此选项，则将第二个图层居中放置在原始图层上。

Blending Mode 混合模式

[>混合模式计算公式 <](https://mir.yuelili.com/ae/ae-basic/blend-mode.html)

### Presrve Transparency 保持透明度

确保未修改原始图层的 Alpha 通道。

|     | 源图层 | 第二图层 | RGBA | Gray | R   | G   | B   | Alpha |
| --- | ------ | -------- | ---- | ---- | --- | --- | --- | ----- |
| R   | 60     | 200      | 130  | 139  | 130 | 145 | 130 | 228   |
| G   | 90     | 100      | 95   | 89   | 80  | 95  | 55  | 178   |
| B   | 70     | 50       | 65   | 65   | 56  | 71  | 61  | 153   |

以下只举例一个 R 通道，其他同理

RGBA 模式:各通道加各通道的，再处以 2。 130=(60+200)/2

Gray 模式:Y（亮度） = 0.299*R 源 + 0.587*G 源 + 0.114\*B 源 =79。所以源图层就按
(79,79,79)算。139=(79+200)/2

R 模式(GB 同理): 其他通道的数值全按 R 的算，源图层就变成了 (60,60,60)算。所以 R 没变，G 的 80 = (60+100)/2

Alpha:Alpha 一般都是 255...所以 228=(255+200)/2
