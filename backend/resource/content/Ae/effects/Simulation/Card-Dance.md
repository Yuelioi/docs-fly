---
title: 卡片动画
order: 1
---

# Card Dance - 卡片动画

## 简述

本效果会将图层分为许多卡片，然后使用置换图层控制卡片的几何形状。

可用于模拟挤压的固定点雕塑、波浪人群或飘浮在池塘表面的字母。

此效果适用于 8-bpc 颜色。

## 效果展示

![](https://cdn.yuelili.com/20220103014701.png) |
![](https://cdn.yuelili.com/20220103013421.png)  
---|---

## 教程

<iframe src="https://player.bilibili.com/player.html?bvid=BV1e34y1X7Vj&page=78&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Simulation-Card_Dance.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-Simulation-Card_Dance_cn.png)

| Card Dance           | 卡片动画   | カードダンス         |                       |                  |                          |                       |                |                            |
| -------------------- | ---------- | -------------------- | --------------------- | ---------------- | ------------------------ | --------------------- | -------------- | -------------------------- |
| Rows & columns       | 行数和列数 | 行と列               | Independent           | 独立             | 独立                     |                       |                |                            |
|                      |            |                      | Columns Follows rows  | 列数受行数控制   | 列が行に従う             |                       |                |                            |
|                      |            |                      |                       |                  |                          |                       |                |                            |
| Rows                 | 行数       | 行数                 |                       |                  |                          |                       |                |                            |
| Columns              | 列数       | 列                   |                       |                  |                          |                       |                |                            |
| Back Layer           | 背面图层   | 背面レイヤー         |                       |                  |                          |                       |                |                            |
| Gradient Layer 1     | 渐变图层 1 |                      |                       |                  |                          |                       |                |                            |
| Gradient Layer 2     | 渐变图层 2 |                      |                       |                  |                          |                       |                |                            |
| Rotation Order       | 旋转顺序   | 回転順序             | XYZ                   | XYZ              | XYZ                      |                       |                |                            |
|                      |            |                      | XZY                   | XZY              | XZY                      |                       |                |                            |
|                      |            |                      | YXZ                   | YXZ              | YXZ                      |                       |                |                            |
|                      |            |                      | YZX                   | YZX              | YZX                      |                       |                |                            |
|                      |            |                      | ZXY                   | ZXY              | ZXY                      |                       |                |                            |
|                      |            |                      | ZYX                   | ZYX              | ZYX                      |                       |                |                            |
| Transformation Order | 变换顺序   | トランスフォーム順序 | Position,Rotate,Scale | 位置，旋转，缩放 |                          |
| 位置, 回転, スケール |            |                      |                       |
|                      |            |                      | Position,Scale,Rotate | 位置，缩放，旋转 | 位置, スケール, 回転     |                       |                |                            |
|                      |            |                      | Rotate,Position,Scale | 旋转，位置，缩放 | 回転, 位置, スケール     |                       |                |                            |
|                      |            |                      | Rotate,Scale,Position | 旋转，缩放，位置 | 回転, スケール, 位置     |                       |                |                            |
|                      |            |                      | Scale,Position.Rotate | 缩放，位置，旋转 | スケール, 位置, 回転     |                       |                |                            |
|                      |            |                      | Scale,Rotate,Position | 缩放，旋转，位置 | スケール, 回転, 位置     |                       |                |                            |
| X Position           | X 位置     |                      | Source                | 源               | ソース                   |                       | 一大堆不写了   |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Y Position           | Y 位置     |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Z Position           | Z 位置     |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| X Rotation           | X 轴旋转   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Y Rotation           | Y 轴旋转   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Z Rotation           | Z 轴旋转   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| X Scale              | X 轴缩放   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Y Scale              | Y 轴缩放   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Z Scale              | Z 轴缩放   |                      | Source                | 源               | ソース                   |                       |                |                            |
|                      |            |                      | Multiplier            | 乘数             | 乗数                     |                       |                |                            |
|                      |            |                      | Offset                | 偏移             | オフセット               |                       |                |                            |
| Camera System        | 摄像机系统 | カメラシステム       | Camera System         | 摄像机系统       | カメラシステム           | Camera                |
| Position             | 摄像机位置 | カメラ位置           |
|                      |            |                      |                       |                  |                          | Corner Pins           | 边角定位       | コーナーピン               |
|                      |            |                      |                       |                  |                          | Comp Camera           | 合成摄像机     | コンポジションカメラ       |
| Camera Position      | 摄像机位置 | カメラ位置           | X Rotation            | X 轴旋转         |                          |                       |                |                            |
|                      |            |                      | Y Rotation            | Y 轴旋转         |                          |                       |                |                            |
|                      |            |                      | Z Rotation            | Z 轴旋转         |                          |                       |                |                            |
|                      |            |                      | XY Position           | X、Y 位置        |                          |                       |                |                            |
|                      |            |                      | Z Position            | Z 位置           |                          |                       |                |                            |
|                      |            |                      | Focal Length          | 焦距             | フォーカス距離           |                       |                |                            |
|                      |            |                      | Transform Order       | 变换顺序         | トランスフォーム順序     | Rotate, XYZ, Position | 旋转 XY        |
| Z，位置              |
|                      |            |                      |                       |                  |                          | Rotate, XZY, Position | 旋转 XZY，位置 | 回転 XZY, 位置             |
|                      |            |                      |                       |                  |                          | Rotate, YXZ, Position | 旋转 YXZ，位置 | 回転 YXZ, 位置             |
|                      |            |                      |                       |                  |                          | Rotate, YZX, Position | 旋转 YZX，位置 | 回転 YZX, 位置             |
|                      |            |                      |                       |                  |                          | Rotate, ZXY, Position | 旋转 ZXY，位置 | 回転 ZXY, 位置             |
|                      |            |                      |                       |                  |                          | Rotate, ZYX, Position | 旋转 ZYX，位置 | 回転 ZYX, 位置             |
|                      |            |                      |                       |                  |                          | Position Rotate, XYZ  | 位置，旋转 XYZ | 位置, 回転 XYZ             |
|                      |            |                      |                       |                  |                          | Position Rotate, XZY  | 位置，旋转 XZY | 位置, 回転 XZY             |
|                      |            |                      |                       |                  |                          | Position Rotate, YXZ  | 位置，旋转 YXZ | 位置, 回転 YXZ             |
|                      |            |                      |                       |                  |                          | Position Rotate, YZX  | 位置，旋转 YZX | 位置, 回転 YZX             |
|                      |            |                      |                       |                  |                          | Position. Rotate, ZXY | 位置，旋转 ZXY | 位置, 回転 ZXY             |
|                      |            |                      |                       |                  |                          | Position Rotate, ZYX  | 位置，旋转 ZYX | 位置, 回転 ZYX             |
| Corner Pins          | 边角定位   | コーナーピン         |                       |                  |                          |                       |                |                            |
|                      |            |                      | Upper Left Comer      | 左上角           | 左上隅                   |                       |                |                            |
|                      |            |                      | Upper Right Corner    | 石上角           |                          |                       |                |                            |
|                      |            |                      | Lower Left comer      | 左下角           | 左下隅                   |                       |                |                            |
|                      |            |                      | Lower Right Corner    | 石下角           |                          |                       |                |                            |
|                      |            |                      | Auto Focal Length     | 自动焦距         | 自動焦点距離             |                       |                |                            |
|                      |            |                      | Focal Length          | 焦距             | フォーカス距離           |                       |                |                            |
| Lighting             | 灯光       | 照明                 |                       |                  |                          |                       |                |                            |
|                      |            |                      | Light Type            | 灯光类型         | ライトの種類             | Point Source          | 点光源         | ポイントソース             |
|                      |            |                      |                       |                  |                          | Distant Source        | 远光源         | 離れたソース               |
|                      |            |                      |                       |                  |                          | First Comp Light      | 首选合成灯光   | 最初のコンポジションライト |
|                      |            |                      | Light Intensity       | 灯光强度         | ライトの強さ             |                       |                |                            |
|                      |            |                      | Light Color           | 灯光颜色         | ライトのカラー           |                       |                |                            |
|                      |            |                      | Light Position        | 灯光位置         | ライトの位置             |                       |                |                            |
|                      |            |                      | Light Depth           | 灯光深度         | ライトの奥行き           |                       |                |                            |
|                      |            |                      | Ambient Light         | 环境光           | アンビエントライト       |                       |                |                            |
| Material             | 材质       | マテリアル           |                       |                  |                          |                       |                |                            |
|                      |            |                      | Diffuse Reflection    | 漫反射           | 拡散反射                 |                       |                |                            |
|                      |            |                      | Specular Reflection   | 镜面反射         | 鏡面反射                 |                       |                |                            |
|                      |            |                      | Highlight Sharpness   | 高光锐度         | ハイライトのシャープネス |                       |                |

### 动画拆解

| MAP1（从左到右的灰度渐变）                      | MAP2（从上到下的灰度渐变） | 动画效果 |
| ----------------------------------------------- | -------------------------- | -------- |
| ![](https://cdn.yuelili.com/20220103015951.png) |
| ![](https://cdn.yuelili.com/20220103020221.png) |

![](https://cdn.yuelili.com/20220103014701.png)

设置好渐变图层

![](https://cdn.yuelili.com/20220103020613.png)

K 动画

这个强度 1，代表渐变图层 1 的强度，也就是 MAP1，同理强度 2 就是基于 MAP2 的强度。

可以看到都是基于亮度（由白到灰）进行翻转。

这就是本效果的最基础的原理了。其他设置请看参数详解。

![](https://cdn.yuelili.com/20220103020337.png)

## 参数详解

### “行数”、“列数”、“图层”和“顺序”控件

**行数和列数**

行数和列数的相互关系。

独立：行数、列数都可以控制。

列数受行数控制：控制行数，列数与行数相同。

**-行数/列数**

最多 1000 行

**注意**

行列始终均匀地分布在图层中，因此形状不规则的矩形拼贴不能沿图层边缘显示，除非使用 Alpha 通道。

### 背面图层

在卡片背面分段显示的图层。如果图层有效果或蒙版，请先预合成此图层。

### 渐变图层 1

用于生成卡片动画效果的第 1 个控件图层。一般使用灰度图层如渐变图层，充当置换图，这样方便预测。

### 渐变图层 2

第 2 个控件图层。

### 旋转顺序

在使用多轴旋转时，卡片围绕多轴旋转的顺序。

**变换顺序**

执行变换（缩放、旋转和位置）的顺序。

### “位置”、“旋转”和“缩放”控件

位置”(X, Y, Z)、旋转”(X, Y, Z) 和缩放”(X, Y) 用于指定要调整的变换属性。

由于卡片动画效果是 3D 效果，因此可以为卡片的每个轴单独控制这些属性。但由于卡片本身是 2D 图像，没有固有深度，因此没有 Z 轴缩放。

**-源**

指定控制变换的渐变图层通道。“强度 2”代表使用“渐变图层 2”的强度。

**-乘数**

应用变换的数量。

**-偏移**

变换开始时使用的基值。这样可以从某些非 0 位置开始变换。

### “摄像机系统”和“摄像机位置”控件

**-摄像机系统**

使用效果的“摄像机位置”属性、效果的“边角定位”属性、或合成摄像机和光照位置来渲染卡片的 3D 图像。

**-X 轴旋转、Y 轴旋转、Z 轴旋转**

围绕相应的轴旋转摄像机。使用这些控件可从上面、侧面、背面或其他任何角度查看卡片。

**-X、Y 位置**

摄像机在 X 轴和 Y 轴上的位置。

**-Z 位置**

摄像机在 Z 轴上的位置。较小的数值使摄像机更接近卡片，较大的数值使摄像机远离卡片。

**-焦距**

缩放系数。焦距越小视角越大。

**-变换顺序**

摄像机围绕其三个轴旋转的顺序，以及摄像机是在使用其他“摄像机位置”控件定位之前还是之后旋转。

### “边角定位”控件

边角定位是备用的摄像机控制系统。此控件可用作辅助控件，以便将效果的结果合成到相对于帧倾斜的平面上的场景中。

**-左上角、右上角、左下角、右下角**

附加图层每个角的位置。

### 自动焦距

控制动画期间效果的透视。取消选择“自动焦距”时，程序将使用您指定的焦距查找摄像机位置和方向，以便在边角固定点放置图层的角（如果可能）。如果不能完成此操作，则此图层将替换为在固定点之间绘制的轮廓。在选择“自动焦距”时，将在可能的情况下使用匹配边角点所需的焦距。否则，程序将插入附近帧中正确的值。

### 焦距

如果您已获得的结果不是所需结果，则覆盖其他设置。如果为“焦距”设置的值不等于固定点实际在该配置中时焦距本该使用的值，则图像可能看起来异常（例如，被奇怪地修剪）。但是，如果知道要试图匹配的焦距，则手动设置“焦距”是获得正确结果的最简单方法。

### 友情效果

[卡片擦除](https://www.yuelili.com/docs/ae-effect/card-wipe/) ：参数 75%相同
