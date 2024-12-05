---
title: 卡片擦除
order: 12
---

# Card Wipe - 卡片擦除

## 简述

此效果模拟一组卡片，这组卡片先显示一张图片，然后翻转以显示另一张图片。“卡片擦除”提供对卡片的行数和列数、翻转方向以及过渡方向的控制（包括使用渐变来确定翻转顺序的功能）。您还可以控制随机性和抖动以使效果看起来更逼真。通过改变行和列，您还可以创建百叶窗和灯笼效果。

卡片擦除效果与卡片动画效果共享许多控件。

此效果适用于 8-bpc 颜色。

## 效果展示

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=78&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Card_Wipe.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Transition-Card_Wipe_cn.png)

| Card Wipe             | 卡片擦除   | カードワイプ           |                          |                |                      |                      |                |     |
| --------------------- | ---------- | ---------------------- | ------------------------ | -------------- | -------------------- | -------------------- | -------------- | --- |
| Transition Completion | 过渡完成   | 変換終了               |                          |                |                      |                      |                |     |
| Transition Width      | 过渡宽度   | 変換の幅               |                          |                |                      |                      |                |     |
| Back Layer            | 背面图层   | 背面レイヤー           |                          |                |                      |                      |                |     |
| Rows & Columns        | 行数和列数 | 行と列                 | Independent              | 独立           | 独立                 |                      |                |     |
|                       |            |                        | Columns Follows Rows     | 列数受行数控制 | 列が行に従う         |                      |                |     |
| Rows                  | 行数       | 行数                   |                          |                |                      |                      |                |     |
| Columns               | 列数       | 列                     |                          |                |                      |                      |                |     |
| Card Scale            | 卡片缩放   | カードスケール         |                          |                |                      |                      |                |     |
| Flip Axis             | 翻转轴     | 反転の軸               | X                        | X              | x                    |                      |                |     |
|                       |            |                        | Y                        | Y              | y                    |                      |                |     |
|                       |            |                        | Random                   | 随机           | ランダム             |                      |                |     |
| Flip Direction        | 翻转方向   | 反転の方向             | Positive                 | 正向           | 正                   |                      |                |     |
|                       |            |                        | Negative                 | 反向           | 負                   |                      |                |     |
|                       |            |                        | Random                   | 随机           | ランダム             |                      |                |     |
| Flip Order            | 翻转顺序   | 反転の順序             | Left to Riaht            | 从左到右       | 左から右へ           |                      |                |     |
|                       |            |                        | Riaht to Left            | 从右到左       | 右から左へ           |                      |                |     |
|                       |            |                        | Top to Bottom            | 自上而下       | 上から下             |                      |                |     |
|                       |            |                        | Bottom to Top            | 自下而上       | 下から上             |                      |                |     |
|                       |            |                        | Top Left to Bottom Right | 左上到右下     | 左上から右下         |                      |                |     |
|                       |            |                        | Top Right to Bottom Left | 右上到左下     | 右上から左下         |                      |                |     |
|                       |            |                        | Bottom Left to Top Riaht | 左下到右上     | 左下から右上         |                      |                |     |
|                       |            |                        | Bottom Right to Top Left | 右下到左上     | 右下から左上         |                      |                |     |
|                       |            |                        | Gradient                 | 渐变           | グラデーション       |                      |                |     |
| Gradient Layer        | 渐变图层   | グラデーションレイヤー |                          |                |                      |                      |                |     |
| Timing Randomness     | 随机时间   | タイミングのランダム度 |                          |                |                      |                      |                |     |
| Random Seed           | 随机植入   | ランダムシード         |                          |                |                      |                      |                |     |
| Camera System         | 摄像机系统 | カメラシステム         | Camera Position          | 摄像机位置     | カメラ位置           |                      |                |     |
|                       |            |                        | Corner Pins              | 边角定位       | コーナーピン         |                      |                |     |
|                       |            |                        | Comp Camera              | 合成摄像机     | コンポジションカメラ |                      |                |
| Camera Position       | 摄像机位置 | カメラ位置             |                          |                |                      |                      |                |     |
|                       |            |                        | X Rotation               | X 轴旋转       |                      |                      |                |     |
|                       |            |                        | Y Rotation               | Y 轴旋转       |                      |                      |                |     |
|                       |            |                        | Z Rotation               | Z 轴旋转       |                      |                      |                |     |
|                       |            |                        | XY Position              | X、Y 位置      |                      |                      |                |     |
|                       |            |                        | Z Position               | Z 位置         |                      |                      |                |     |
|                       |            |                        | Focal Length             | 焦距           | フォーカス距離       |                      |                |     |
|                       |            |                        | Transform Order          | 变换顺序       | トランスフォーム順序 | Rotate XYZ, Position | 旋转 XYZ，位置 |

|  
|| | | | | | Rotate XZY, Position | 旋转 XZY，位置 | 回転 XZY, 位置|
|| | | | | | Rotate YXZ, Position | 旋转 YXZ，位置 | 回転 YXZ, 位置|
|| | | | | | Rotate YZX, Position | 旋转 YZX，位置 | 回転 YZX, 位置|
|| | | | | | Rotate ZXY, Position | 旋转 ZXY，位置 | 回転 ZXY, 位置|
|| | | | | | Rotate ZYX, Position | 旋转 ZYX，位置 | 回転 ZYX, 位置|
|| | | | | | Position, Rotate XYZ | 位置，旋转 XYZ | 位置, 回転 XYZ|
|| | | | | | Position, Rotate XZY | 位置，旋转 XZY | 位置, 回転 XZY|
|| | | | | | Position, Rotate YXZ | 位置，旋转 YXZ | 位置, 回転 YXZ|
|| | | | | | Position, Rotate YZX | 位置，旋转 YZX | 位置, 回転 YZX|
|| | | | | | Position, Rotate ZXY | 位置，旋转 ZXY | 位置, 回転 ZXY|
|| | | | | | Position, Rotate ZYX | 位置，旋转 ZYX | 位置, 回転 ZYX|
|Corner Pins | 边角定位 | コーナーピン | | | | | ||
|| | | Upper Left Corner | 左上角 | 左上隅 | | ||
|| | | Upper Right Corner | 右上角 | 右上隅 | | ||
|| | | Lower Left Corne | 左下角 | 左下隅 | | ||
|| | | wwer Riaht Corner | 石下角 | | | ||
|| | | Auto Focal Lenght | 自动焦距 | 自動焦点距離 | | ||
|| | | Focal Length | 焦距 | フォーカス距離 | | ||
|Lighting | 灯光 | 照明 | | | | | ||
|| | | Light Type | 灯光类型 | ライトの種類 | Point Source | 点光源 | ポイントソース|
|| | | | | | Distant Source | 远光源 | 離れたソース|
|| | | | | | First Comp Light | 首选合成灯光 | 最初のコンポジションライト|
|| | | Light Intensity | 灯光强度 | ライトの強さ | | ||
|| | | Light Color | 灯光颜色 | ライトのカラー | | ||
|| | | Light Position | 灯光位置 | ライトの位置 | | ||
|| | | Light Depth | 灯光深度 | ライトの奥行き | | ||
|| | | Ambient Light | 环境光 | アンビエントライト | | ||
|Material | 材质 | マテリアル | | | | | ||
|| | | Diffuse Reflection | 漫反射 | 拡散反射 | | ||
|| | | Specular Reflecton | 镜面反射 | 鏡面反射 | | ||
|| | | Highlight Sharpness | 高光锐度 | ハイライトのシャープネス | ||
|Position Jitter | 位置抖动 | 位置の変位 | | | | | ||
|| | | X Jitter Amount | X 抖动量 | | | ||
|| | | X Jitter Speed | X 抖动速度 | | | ||
|| | | Y Jitter Amount | Y 抖动量 | | | ||
|| | | Y Jitter Speed | Y 抖动速度 | | | ||
|| | | Z Jitter Amount | Z 抖动量 | | | ||
|| | | Z Jitter Speed | Z 抖动速度 | | | ||
|Rotation Jitter | 旋转抖动 | 回転の変位 | | | | | ||
|| | | X Rot Jitter Amount | X 旋转抖动量 | | | ||
|| | | X Rot Jitter Speed | X 旋转抖动速度 | | | ||
|| | | Y Rot Jitter Amount | Y 旋转抖动量 | | | ||
|| | | Y Rot Jitter Speed | Y 旋转抖动速度 | | | ||
|| | | Z Rot Jitter Amount | Z 旋转抖动量 | | | ||
|| | | Z Rot Jitter Speed | Z 旋转抖动速度 | | | ||

## 参数详解

### 过渡宽度

从原始图像更改到新图像的区域的宽度。

### 背面图层

在卡片背面分段显示的图层。您可以使用合成中的任何图层；甚至可以关闭其“视频”开关 。如果图层有效果或蒙版，则先预合成此图层。

### 行数和列数

指定行数和列数的相互关系。“独立”可同时激活“行数”和“列数”滑块。“列数受行数控制”只激活“行数”滑块。如果选择此选项，则列数始终与行数相同。

**-行数/列数**

最多 1000 行

**注意**

行列始终均匀地分布在图层中，因此形状不规则的矩形拼贴不能沿图层边缘显示，除非您使用 Alpha 通道。

### 卡片缩放

卡片的大小。小于 1 的值会按比例缩小卡片，从而显示间隙中的底层图层。大于 1 的值会按比例放大卡片，从而在卡片相互重叠时创建块状的马赛克效果。

### 翻转轴

每个卡片绕其翻转的轴。

### 翻转方向

卡片绕其轴翻转的方向。

### 翻转顺序

过渡发生的方向。您还可以使用渐变来定义自定义翻转顺序：卡片首先翻转渐变为黑色的位置，最后翻转渐变为白色的位置。

### 渐变图层

要用于“翻转顺序”的渐变图层。您可以使用合成中的任何图层。

### 随机时间

使过渡的时间随机化。如果此控件设置为 0，则卡片将按顺序翻转。值越高，卡片翻转顺序的随机性就越大。

### 摄像机系统

使用效果的“摄像机位置”属性、效果的“边角定位”属性，还是默认的合成摄像机和光照位置来渲染卡片的 3D 图像。

**“摄像机位置”控件**

X 轴旋转、Y 轴旋转、Z 轴旋转

围绕相应的轴旋转摄像机。使用这些控件可从上面、侧面、背面或其他任何角度查看卡片。

**-X、Y 位置**

摄像机在 x、y 空间中的位置。

**-Z 位置**

摄像机在 Z 轴上的位置。较小的数值使摄像机更接近卡片，较大的数值使摄像机远离卡片。

**-焦距**

从摄像机到图像的距离。焦距越小视角越大。

**-变换顺序**

摄像机围绕其三个轴旋转的顺序，以及摄像机是在使用其他“摄像机位置”控件定位之前还是之后旋转。

### “边角定位”控件

边角定位是备用的摄像机控制系统。此控件可用作辅助控件，以便将效果的结果合成到相对于帧倾斜的平面上的场景中。

**-左上角、右上角、左下角、右下角**

附加图层每个角的位置。

**-自动焦距**

控制动画期间效果的透视。如果取消选择“自动焦距”，程序将使用您指定的焦距查找摄像机位置和方向，以便在边角固定点放置图层的角（如果可能）。

如果不能完成此操作，则此图层将替换为在固定点之间绘制的轮廓。如果选择“自动焦距”，将在可能的情况下使用匹配边角点所需的焦距。否则，程序将插入附近帧中正确的值。

**-焦距**

如果您已获得的结果不是所需结果，则覆盖其他设置。如果为“焦距”设置的值不等于固定点实际在该配置中时焦距本该使用的值，则图像可能看起来异常（例如，被奇怪地修剪）。但是，如果您知道您试图匹配的焦距，则此控件是获得正确结果的最简单方法。

### “抖动”控件

添加抖动（“位置抖动”和“旋转抖动”）可使该过渡更加逼真。抖动可在过渡发生之前、发生过程中和发生之后对卡片生效。如果您想让抖动仅在过渡期间发生，请从抖动量
0 开始，在过渡期间使其增加到所需的量，然后在过渡完成时使其返回到 0。

**-位置抖动**

指定 x、y 和 z 轴的抖动量和速度。“X 抖动量”、“Y 抖动量”和“Z 抖动量”指定额外运动的量。“X 抖动速度”、“Y 抖动速度”和“Z
抖动速度”值指定每个“抖动量”选项的抖动速度。

**-旋转抖动**

指定围绕 x、y 和 z 轴的旋转抖动的量和速度。“X 旋转抖动量”、“Y 旋转抖动量”和“Z 旋转抖动量”指定沿某个轴旋转抖动的量。值 90°
使卡片可在任意方向旋转最多 90°。“X 旋转抖动速度”、“Y 旋转抖动速度”和“Z 旋转抖动速度”值指定旋转抖动的速度。

### 友情效果

[卡片动画](https://www.yuelili.com/docs/ae-effect/card-dance/) 参数 75%相同
