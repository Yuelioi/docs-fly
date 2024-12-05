---
title: 复合运算
order: 6
---

# Compound Arithmetic - 复合运算

## 简述

复合运算效果可以数学方式合并应用效果的图层和控件图层。

复合运算效果仅提供与如下项目的兼容性：在 After Effects 早期版本中创建的使用复合运算效果的项目。使用混合模式通常比使用复合运算效果更有效。

此效果适用于 8-bpc 颜色。

## 效果展示

以下部分图无版权！仅供学习参考！

![](https://mir.yuelili.com/user/AE/effects/ext/image00426.jpg)

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=115&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Compound_Arithmetic.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Channel-Compound_Arithmetic_cn.png)

| Compound Arithmetic | 复合运算       | 合成アリスマチック  |            |        |              |
| ------------------- | -------------- | ------------------- | ---------- | ------ | ------------ |
| Second Source Layer | 第二个源图层   | 第 2 ソースレイヤー |            |        |              |
| Operator            | 运算符         | 演算子              | Copy       | 复制   | コピー       |
|                     |                |                     | Add        | 相加   | 加算         |
|                     |                |                     | Subtract   | 相减   | 減算         |
|                     |                |                     | Multiply   | 相乘   | 乗算         |
|                     |                |                     | Difference | 差值   | 差           |
|                     |                |                     | Xor        | 与或   |              |
|                     |                |                     | Lighter    | 异或   |              |
|                     |                |                     | Darker     | 变亮   | 明るく       |
|                     |                |                     | Minimum    | 变暗   | 暗く         |
|                     |                |                     | Maximum    | 最小值 | 最小         |
|                     |                |                     | creen      | 最大值 | 最大         |
|                     |                |                     | Overla     | 滤色   | スクリーン   |
|                     |                |                     | Overlay    | 叠加   | オーバーレイ |
|                     |                |                     | Hard Light | 强光   | ハードライト |
| Operate on Channels | 在通道上运算   | チャンネルを操作    | RGB        | RGB    | RGB          |
|                     |                |                     | ARGB       | ARGB   | ARGB         |
|                     |                |                     | Alpha      | Alpha  | アルファ     |
| Overflow Behavior   | 溢出特性       | オーバーフロー動作  | Clip       | 剪切   | クリッピング |
|                     |                |                     | Wrap       | 回绕   | ラップ       |
|                     |                |                     | Scale      | 缩放   | スケール     |
| Blend With Original | 与原始图像混合 | 元の画像とブレンド  |            |        |              |

## 参数详解

![](https://mir.yuelili.com/user/AE/effects/ext/image00427-1.jpg)

### 第二个源图层

在给定运算中与当前图层一起使用的图层。

### 运算符

在两个图层之间执行的运算。

### 在通道上运算

向其应用效果的通道。

### 溢出特性

效果重映射超出 0-255 灰度范围的值的方式。

### 剪切

大于 255 的值映射到 255。小于 0 的值映射到 0。

### 回绕

大于 255 和小于 0 的值被回绕到 0-255 范围内。例如，将 258 回绕至 2，将 256 回绕至 0，将 3 回绕至 253。

### 缩放

将最大值和最小值重映射至 255 和 0，将中间值伸展或压缩以适合此范围。

### 伸缩第二个源以适合

缩放第二个图层以匹配当前图层的大小（宽度和高度）。如果取消选择此选项，则使第二个图层与源图层的左上角对齐，根据源图层的当前大小放置第二个图层。

### 与原始图像混合

效果图像的透明度。效果图像与原始图像混合的结果，并在上面合成效果图像结果。此值设置得越高，此效果对图层的影响越小。例如，如果将此值设置为
100%，则此效果在图层上不会产生明显结果；如果将此值设置为 0%，则不会隐约显出原始图像。
