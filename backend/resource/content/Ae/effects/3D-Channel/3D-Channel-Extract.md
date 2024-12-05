---
title: 3D通道提取
order: 1
---
# 3D Channel Extract - 3D通道提取

## 简述

用 RLA、RPF 文件的通道制作蒙版

使用 3D
CG 软件生成的 RLA、RPF 文件中，包含的 Z 深度、对象 ID 等特有的通道，分别显示并用于蒙版。在进行焦点处理、图像内每个对象的色彩调整等时，效果非常便利。

例如，在 3D 通道图像文件中提取深度信息，然后在粒子运动场效果中将其用作影响图，或从非固定 RGB 通道提取值来生成遮罩，用以生成夺目的高光。

## 参数中英日对照


| 3D Chanel Extract | 3D 通道提取  | 3D チャンネル抽出 |                 |            |                  |
| ----------------- | ------------ | ----------------- | --------------- | ---------- | ---------------- |
| 3D channel        | 3D 通道      | 3D チャンネル     | Z-Depth         | Z 深度     | Z 深度           |
|                   |              |                   | Object ID       | 对象 ID    | オブジェクト ID  |
|                   |              |                   | Texture UV      | 纹理 UV    | テクスチャ UV    |
|                   |              |                   | Surface Normals | 曲面法线   | サーフェイス法線 |
|                   |              |                   | Coverage        | 覆盖范围   | 範囲             |
|                   |              |                   | Background RGB  | 背景 RGB   | 背景 RGB         |
|                   |              |                   | Unclamped RGB   | 非固定 RGB | 非固定 RGB       |
|                   |              |                   | Meterial ID     | 材质 ID    | マテリアル ID    |
| Black Point       | 黑场         | ブラックポイント  |                 |            |                  |
| White Point       | 白场         | ホワイトポイント  |                 |            |                  |
| Anti-alias        | 消除锯齿     | アンチエイリアス  |                 |            |                  |
| Clamp Output      | 固定输出     | クランプ出力      |                 |            |                  |
| Invert Depth Map  | 反转景深映射 | 深度マップを反転  |                 |            |                  |

## 参数图片参考

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-3D_Chanel_Extract.png)![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-3D_Chanel_Extract_cn.png)

![](https://mir.yuelili.com/user/AE/effects/list0/3D-Channel-3D_Chanel_Extract-00.png)![](https://mir.yuelili.com/user/AE/effects/list0/3D-Channel-3D_Chanel_Extract-01_cn.png)
