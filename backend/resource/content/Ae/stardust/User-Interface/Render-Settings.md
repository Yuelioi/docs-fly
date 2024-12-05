---
title: Render Settings 渲染设置
order: 2
---

# Render Settings 渲染设置

## Quality Settings 质量设置

渲染质量或速度。

- Medium 中等：
- High 高：
- Very High 很高：
- Extreme 极限：

### Render Output 渲染输出

选择渲染不同的输出，深度通道，环境光遮挡等。此选项仅适用于 3D 模型。

- Render 渲染：
- Diffuse 漫射：
- Depth 深度：
- Ambient Occlusion 环境吸收：

### Preview 预览

仅渲染一些粒子以进行快速预览。这是一个百分比参数，因此您可以通过繁重的设置更快地查看动画。

- On／Off 开关：
- Particle Chance 粒子显示量：
- 3D Preview 3D 预览：
- 3D 预览–将 3D 模型显示为线框/平面/平面阴影以进行快速预览。：
- Wire 线框：
- Flat 平：
- Shaded：

### Helpers

向某些节点添加视觉提示，比如使用样条线构造线，OBJ 轮廓，字段边界框等

- Draw Helpers：
- Color 颜色：
- Opacity 不透明的：

### Environment 环境

设置图像以用于基于图像的照明，如果未选择任何层，则使用默认颜色。此选项仅适用于 3D 模型。  
使用“混合灯光”阴影通过场景灯光影响环境。  
您可以将环境切换为在后台可见。

- Environment Layer 环境图层：
- Color 颜色：
- Orient X X 轴方向：
- Orient Y Y 轴方向：
- Orient Z Z 轴方向：
- Mix Lights Shading 混合灯光阴影：
- Visible 可见：

### Ambient Occlusion 环境吸收

向场景添加环境光遮挡阴影，您可以控制其半径，强度等。您还可以启用“环境光遮挡”双边模糊，以进一步平滑结果。此选项仅适用于 3D 模型。

- On／Off 开关：
- Radius 半径：
- Intensity 强度：
- Camma 伽马：
- AO Blur AO 模糊：
- Color 颜色：

### Motion Blur 运动模糊

对粒子应用运动模糊效果。  
您可以在需要第一次和最后一次采样并对其进行插值的线性样本之间进行选择，也可以在对准确时间进行采样的子帧样本之间进行选择，以给出更准确的结果，这可能会更慢计算。  
在应用运动模糊之后，增益可以调整粒子的不透明度。

- On／Off 开关：
- Shutter Angle 快门角度： 可以理解为模糊程度
- Shutter Phase 快门相位： 可以理解为采样偏移
- Subframe Sample 子取样： 单击查看取样范围
- Samples Per Frame 样品取样：

控制运动模糊的精度

- Gain 增益：增强模糊，增加不透明度：

### Shadows 阴影

向 3D 模型添加阴影。要启用阴影，请在合成中添加投射阴影灯光，并具有投射和接受阴影 3D 模型。  
阈值参数允许调整一些阴影伪像。  
您还可以使用“柔和度”参数调整阴影的整体模糊度。

- On／Off 开关：
- Threshold 阈值：
- Softness 柔和度：

### Subsurface Scattering 次表面散射

创建具有光散射的材料，例如皮肤或蜡。启用此选项可在场景中启用“次表面散射”材料。

- On／Off 开关：
- Radius 半径：

### Fog 雾

- On／Off 开关：
- Opacity 不透明的：
- Color 颜色：
- Start 开始：
- End 结束：

### Shading 2D

打开以后，灯光自己的颜色可以作用于粒子的阴影

### Z Buffer Z 通道缓存

使用此参数通过合成内部的另一个 3D 层或外部 Z 缓冲区遮挡粒子，输入层应为亮度线性贴图。此选项仅适用于 2D 粒子。  
使用外部 Z 时，可以将黑白点设置为特定的 z 值。

- Obstruction 遮挡：
- Z Buffer Z 通道图层：
- Z Map To Black Z 通道黑色：
- Z Map To White Z 通道白色：

### Clipping 剪裁

- Min Distance 最小距离：
- Min Fade 最小距离过度：
- Max Distance 最大距离：
- Max Fade 最大距离过度：
- 2D Renderer 二维渲染器： 当认为“最佳”时，渲染“Circle”和“云”粒子类型时，星尘将自动切换到 GPU。可以手动覆盖并将其设置为 CPU。

- Auto 自动
- CPU

### Volume Cache 体积缓存

- Off 关：
- On 开：
- Freeze 冻结：
- Cache Folder 缓存文件夹：

物理粒子–首次在每个 Stardust 实例上启用“模拟”时，为了加快处理速度，可以缓存。要启用缓存，请打开“缓存”  
Stardust 将在与项目文件相同的位置选择一个文件夹。因此，必须先保存项目，然后才能生成仿真。您可以通过单击“显示”来显示该文件夹，或使用“重置”将其清除（这只会清除名称，而不会清除实际文件）
