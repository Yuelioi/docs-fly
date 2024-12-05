---
title: 雾3D
order: 6
---

# Fog 3D 雾3D

## 简述

使用深度通道产生雾  
根据包括 Z 深度信道的 RLA 和 RPF 文件的深度，在图像中产生雾和雾。可以调整雾的颜色、浓度、深度等。在强调使用大气远近法的深度感的情况下，效果非常便利。

雾 3D 效果通过模拟穿过空气中的散射介质行为（对象的 z 轴距离越远，弥散效果越强）来模拟烟雾效果。

适用于 8-bpc、16-bpc 和 32-bpc 。

注意:与 3D 通道中的其他效果一样，雾 3D 效果取决于从 3D 应用程序渲染的图像序列文件的深度信息。有关 3D 通道效果的常规信息，请参阅关于 3D
通道效果（包括 ProEXR 效果）

## 效果展示

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-fog_3d1.jpg)

## 中英日对照

![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Fog_3D.png)
![](https://mir.yuelili.com/user/AE/effects/AE-Effects-3D-Channel-Fog_3D_cn.png)

| Fog 3D             | 雾 3D      | フォグ 3D              |     |     |
| ------------------ | ---------- | ---------------------- | --- | --- |
| Fog Color          | 雾颜色     | フォグのカラー         |     |     |
| Fog Start Depth    | 雾开始深度 | フォグの開始深度       |     |     |
| Fog End Depth      | 雾结束深度 | フォグの終了深度       |     |     |
| Fog Opacity        | 雾不透明度 | フォグの不透明度       |     |     |
| Scattering Density | 散布浓度   | 濃度の拡散             |     |     |
| Foggy Background   | 多雾背景   | フォギーな背景         |     |     |
| Gradient Layer     | 渐变图层   | グラデーションレイヤー |     |     |
| Layer Contribution | 图层贡献   | レイヤーの適用         |     |     |

## 参数详解

### FOG START DEPTH 雾开始深度

沿 z 轴开始漫散射的深度。注意:要确定对象的深度，请在选择效果后，使用选择工具单击“合成”面板或“图层”面板中的对象即可。

### FOG END DEPTH 雾结束深度

沿 z 轴弥散达到最大程度的位置。

### FOG OPACITY 雾不透明度

雾的不透明度

### SCATTERING DENSITY 散布浓度

确定散布的速度。值越高，起始点雾越浓。

### FOGGY BACKGROUND 多雾背景

创建多雾背景（默认）。取消选择此选项，则在 3D 场景后面创建透明度，以便在其他图层上面进行合成。

### GRADIENT LAYER 渐变图层

（可选）用作控件图层的灰度图层，其亮度值适用于雾密度。例如，使用湍流杂色效果可为大气层烟雾创建迷离的控件图层。确保渐变图层的尺寸至少与 3D
场景图层的尺寸一样大。

### LAYER CONTRIBUTION 图层贡献

渐变图层对雾浓度的影响。示例:渐变图层为从里到外的黑色圆环,把图层贡献开到最大,效果如下所示

渐变图层

![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Fog_3D-01.png)

最终结果  
![](https://mir.yuelili.com/user/AE/effects/list/3D-Channel-Fog_3D-02.png)

### 示例 2

![](https://mir.yuelili.com/user/AE/effects/ext/3D-Channel-fog_3d2.jpg)
