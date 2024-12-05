---
title: 设置遮罩
order: 11
---

# Set Matte - 设置遮罩

## 简述

设置遮罩效果可将某图层的 Alpha 通道（遮罩）替换为该图层上面的其他图层的通道，以此创建移动遮罩效果。

设置遮罩效果仅提供与如下项目的兼容性：在 After Effects 早期版本中创建的使用设置遮罩效果的项目。

此效果适用于 8-bpc 和 16-bpc 颜色。在 After Effects CS6 或更高版本中，此效果在 32 位颜色中有效。

## 参数中英日对照

| Set Matte             | 设置遮罩                       | マット設定                   |                               |                      |                    |
| --------------------- | ------------------------------ | ---------------------------- | ----------------------------- | -------------------- | ------------------ |
| Take Matte From Layer | 从图层获取遮罩                 | レイヤーからマットを取り込む |                               |                      |
| Use For Matte         | 用于遮罩                       | マットに使用                 | Red channe                    | 红色通道             | 赤チャンネル       |
|                       |                                |                              | Green Channel                 | 绿色通道             | 緑チャンネル       |
|                       |                                |                              | Blue channel                  | 蓝色通道             | 青チャンネル       |
|                       |                                |                              | Alpha Channel                 | Alpha 通道           | Alpha チャンネル   |
|                       |                                |                              | Luminance                     | 明亮度               | ルミナンス         |
|                       |                                |                              | Hue                           | 色相                 | 色相               |
|                       |                                |                              | Lightness                     | 亮度                 | 輝度               |
|                       |                                |                              | Saturation                    | 饱和度               | 彩度               |
|                       |                                |                              | Full                          | 完全                 |                    |
|                       |                                |                              | Off                           | 关闭                 | 閉じる             |
| Invert Matte          | 反转遮罩                       | 反転マット                   |                               |                      |                    |
| If Layer Sizes Differ | 如果图层大小不同               | レイヤーサイズが異なる場合   | Stretch Matte to Fit          |                      |
| 伸缩遮罩以适合        | マットを伸縮してフィットさせる |
|                       |                                |                              | Composite Matte with Original | 将遮罩与原始图像合成 | 元のマットを合成   |
|                       |                                |                              | Premultiply Matte Layer       | 预乘遮罩图层         | 合成マットレイヤー |

## 参数图片参考

![](http://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Set_Matte.png)
![](http://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Set_Matte_cn.png)

## 效果展示

以下部分图无版权！仅供学习参考！

![](http://mir.yuelili.com/user/AE/effects/ext/image00422.jpg)

![](http://mir.yuelili.com/user/AE/effects/ext/image00423-1.jpg)

### 从图层获取遮罩

用作替换遮罩的图层。

### 用于遮罩

用于遮罩的通道。

### 反转遮罩

反转遮罩的透明度值。

### 伸缩遮罩以适合

缩放所选图层以匹配当前图层的大小。如果取消选择“伸缩遮罩以适合”，则将指定为遮罩的图层居中放置在第一个图层中。

### 将遮罩与原始图像合成

将新遮罩与当前图层合成，而不是替换它。仅当当前遮罩和新遮罩都有些不透明时，生成的遮罩才使图像可隐约显出。

### 预乘遮罩图层

将新遮罩图层预乘当前图层。
