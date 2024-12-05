---
title: "【中英对照】AE蓝宝石插件 -说明+效果预览"
date: "2021-07-24"
categories: 
  - "ae"
  - "ae-post"
---

官方文档：[Support](https://borisfx.com/support/documentation/)

腾讯文档：[戳我](https://docs.qq.com/sheet/DQVRnZmNBTHBYYlJx?groupUin=9BSaAFcBFRVXbL9pdSyoLw%253D%253D&ADUIN=435826135&ADSESSION=1629774748&ADTAG=CLIENT.QQ.5827_.0&ADPUBNO=27151&tab=BB08J2)

全参数汉化文档：尽情期待 说明：＂描述＂由Deepl生成，部分没图，官方也没。有几个插件带＂\_2＂ 是因为说明太长了，分了2个格子

| 英文名 | 分类 | 预览 | 简单描述 |
| --- | --- | --- | --- |
| S\_ChannelSwitcher | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_ChannelSwitcher.jpg) | 重新排列源素材的RGBA通道。 允许将任何源通道映射到任何输出通道，并对每个输出通道进行缩放和抵消。 |
| S\_ClampChroma | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_ClampChroma.jpg) | 必要时降低输入素材的色度，使其不超过指定的最大值。 这个效果可以用来制作 广播安全 的颜色。 它还可以用来缩放色度，钳制亮度，或缩放亮度。 |
| S\_DuoTone | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_DuoTone.jpg) | 使用源素材的亮度在两个指定的颜色之间进行插值。 |
| S\_Gamma | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_Gamma.jpg) | 对输入的素材进行伽玛校正。 红色、绿色和蓝色通道可以独立调整。 从Gamma只是引起调整Gamma的反向效果。 |
| S\_Hotspots | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_Hotspots.jpg) | 生成一个热点图像，包含源剪辑中比给定阈值更亮的区域。 热点的颜色应该与原始源头相匹配。 这可用于增加对比度或寻找剪辑中的明亮区域，但不改变结果的色彩饱和度或色调。 |
| S\_HueSatBright | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_HueSatBright.jpg) | 调整输入素材的色调、饱和度、亮度和/或偏移。 |
| S\_Invert | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_Invert.jpg) | 反转源素材的颜色，所以黑色变成白色，白色变成黑色。 这也可以选择独立地反转luma、chroma、RGB和alpha通道，并对反转后的结果做一些基本的颜色校正。 |
| S\_Monochrome | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_Monochrome.jpg) | 使用可调整的红、绿、蓝通道的权重，生成源素材的单色版本。 这可以模拟在黑白相机的镜头上使用彩色滤镜的情况。 例如，使用更多的红色权重，使输入的蓝色天空区域变暗。 在调整权重时，权重会被缩放，所以它们的总和为1，以减少整体亮度的变化。 |
| S\_OCIOTransform | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_OCIOTransform.jpg) | 使用OpenColorIO将输入从一个颜色空间转换到另一个颜色空间。 |
| S\_QuadTone | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_QuadTone.jpg) | 利用源素材的亮度在四个指定的颜色之间进行插值。 |
| S\_ShowBadColors | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_ShowBadColors.jpg) | 识别所有不在给定颜色范围内的像素，并用相同的颜色标记它们，这样它们就很容易被看到。 |
| S\_Threshold | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_TimeAverage.jpg) | 使用给定的柔和度和阈值将源素材的颜色通道设置为全开或全关。 这可以用来独立增加每个颜色通道的对比度。 |
| S\_Tint | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_TrailsDiffuse.jpg) | 将输入片段的深色和浅色区域染成指定的颜色。 黑暗的颜色用 Tint Dark 颜色，明亮的颜色则用 Tint Lights 颜色。 |
| S\_TriTone | Adjust | ![](https://cdn.yuelili.com/2021/sapphire_TVDamage.jpg) | 使用源素材的亮度在三种指定的颜色之间执行插值。 |
| S\_AutoPaint | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_AutoPaint.jpg) | 生成源素材的 画笔 版本。 使用频率和笔触长度参数来调整笔触的密度和形状。 如果你想为每一帧重新随机化画笔笔触模式，你可以把抖动帧数设置为1。 |
| S\_Cartoon | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_Cartoon.jpg) | 生成一个具有卡通外观的源片段的版本。 找到图像中的边缘，为这些边缘绘制新的轮廓。 平滑边缘之间区域的颜色，并可选择将颜色海报化为更少的颜色值。 |
| S\_CartoonPaint | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_CartoonPaint.jpg) | 自动生成一个具有卡通刷色外观的源素材版本。 找到图像中的边缘，为这些边缘绘制新的轮廓。 用画笔形状取代边缘之间区域的颜色。 |
| S\_Crosshatch | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_Crosshatch.jpg) | 用重叠的笔触模拟钢笔画的交叉阴影效果。 信号源根据luma分为四个波段；每个波段从暗到亮都有不同的笔触模式。 |
| S\_Etching | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_Etching.jpg) | 使用两组不同厚度的黑白线条生成源素材的版本，以呈现 蚀刻 或 石版画 的效果。 使用平滑源参数，去除一些细节，使线条的形状更均匀。 使用线条频率参数来调整所有线条的密度。 |
| S\_HalfTone | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_HalfTone.jpg) | 使用黑白点的模式生成源素材的半色调版本。 使用平滑源参数去除一些细节，使点的圆度更一致。 |
| S\_HalfToneColor | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_HalfToneColor.jpg) | 使用彩色的点状图案生成源素材的版本。 使用平滑源参数去除一些细节，使圆点更一致。 你可以使用 点 菜单将点的图案从CMY反转为RGB。 |
| S\_HalfToneRings | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_HalfToneRings.jpg) | 使用同心圆的重复图案生成源素材的双色调版本。 使用平滑源参数去除一些细节，使圆点的形状更加一致。 |
| S\_Posterize | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_Posterize.jpg) | 通过限制源图像的颜色数量，用纯色取代详细的纹理和噪音，生成输入的海报版本。 |
| S\_Sketch | AutoPaint & HalfTone | ![](https://cdn.yuelili.com/2021/sapphire_Sketch.jpg) | 生成一个具有手绘草图外观的输入版本。 这种效果的结果可能取决于图像的分辨率，因此建议在处理一个片段之前测试你的最终分辨率。 |
| S\_Beauty | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Beauty.jpg) | 对皮肤区域进行平滑、色彩校正、柔焦和发光处理。 皮肤区域是基于启用皮肤检测的值和是否提供第二个输入。 如果启用皮肤检测功能，该插件将从皮肤颜色、亮度和色度范围参数中生成一个内部哑光。 第二个输入是一个可选的哑光，如果连接起来，效果只适用于哑光的明亮区域（如果启用皮肤检测功能并提供哑光，两者相加）。 |
| S\_Blur | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Blur.jpg) | 使用高斯滤波器、三角滤波器或盒式滤波器对源素材进行任意数量的模糊处理。 这种效果即使在非常大的宽度值下也能迅速呈现。 使用Blur Rel X和Y参数可以获得更多的水平或垂直模糊方向。 |
| S\_BlurChannels | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_BlurChannels.jpg) | 使用高斯、三角或盒式滤波器对源素材的每个通道进行任意数量的模糊。 这种效果即使在非常大的宽度值下也能快速渲染。 使用Blur Rel X和Y参数获得更多水平或垂直的模糊方向。 |
| S\_BlurChroma | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_BlurChroma.jpg) | 将信号源分离成亮度和色度成分，独立地模糊色度和/或亮度，然后将它们重新组合。 你也可以独立地缩放亮度和色度，以增强或删除其中一个。 |
| S\_BlurDirectional | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_BlurDirectional.jpg) | 使用高斯、三角或盒式滤波器在给定的方向上模糊源素材。它还可以对每个通道进行不同程度的模糊。 |
| S\_BlurMoCurves | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_BlurMoCurves.jpg) | 执行运动模糊，并可选择使用Z Dist、Rotate和Shift参数的动画曲线对源素材进行变换。 如果这些参数是恒定的，就不会发生运动模糊。 |
| S\_BlurMotion | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_BlurMotion.jpg) | 在指定的 从 和 到 转换之间，对源素材进行运动模糊处理。 这可以用来执行径向变焦模糊、旋转模糊、定向模糊，或这些的任何组合。 From和To参数并不是指时间。 它们描述了空间中的两个变换，决定了应用于每一帧的模糊的风格。 |
| S\_Convolve | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Convolve.jpg) | 用内核对源图像进行卷积。 卷积是一种数学运算，它使用一个图像，即内核，作为另一个图像（源）的过滤形状。 卷积有效地在源图像的每个点上印上一个内核的副本，使用源图像在该点的亮度。 其效果是，核心的副本将出现在源的所有亮点上。 一个像圆形或多边形的内核图像会产生类似于RackDefocus的效果；一个像星芒的内核图像会产生类似于Glare的效果。 |
| S\_ConvolveComp | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_ConvolveComp.jpg) | 用内核卷积正面和背面的图像，并使用哑光进行合成。 卷积是一个数学运算，它使用一个图像，即内核，作为另一个图像（源）的过滤器形状。 卷积有效地在源的每个点上印上一个内核的副本，使用源在该点的亮度。 其效果是，核心的副本将出现在源的所有亮点上。 核心图像的形状像一个圆形或多边形，会产生类似于RackDefocusComp的效果；核心图像的形状像一个星芒，会产生类似于GlareComp的效果。 |
| S\_Deband | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Deband.jpg) | 通过扩散带状区域的像素来去除剪辑中的带状伪影，同时保持原始边缘的完整。 要使用这个效果，首先选择Show:Edges并调整边缘阈值，直到带状边缘消失，只留下所需的真实边缘。 然后选择 显示：结果 来查看结果。 如果你仍然看到一些条带，增加漫反射阈值和/或漫反射半径。 |
| S\_DefocusPrism | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_DefocusPrism.jpg) | 将源素材的颜色通道散焦为不同宽度的环。 |
| S\_EdgeAwareBlur | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_EdgeAwareBlur.jpg) | 模糊相似颜色的区域，同时保留不同颜色区域之间的边缘。 |
| S\_EdgeBlur | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_EdgeBlur.jpg) | 找到哑光素材的边缘，并在这些边缘模糊源素材。 使用 显示边缘 选项来查看哪些区域将接受模糊处理，同时调整边缘参数。 然后调整模糊宽度来控制模糊的数量。 |
| S\_FreeLens | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_FreeLens.jpg) | 生成源素材的扭曲、失焦和漏光的版本，以模拟摄影机内的技术，即在摄影机前拿着一个分离的镜头，并移动它来创造焦点和光线效果。 |
| S\_GrainRemove | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_GrainRemove.jpg) | 平滑源素材，同时保留其边缘。 要调整这个效果的参数，首先使用Show:Edges选项来检查哪些边缘将被保留，并调整Edges Threshold、Edges Width和Edges Scale，直到重要的边缘相当锐利和明亮但不锯齿。 然后返回Show:Result（结果），调整平滑参数以去除适当的颗粒。 |
| S\_Median | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Median.jpg) | 在源图像上应用中值滤镜。中值滤镜对于清理孤立的斑点和噪声很有用。 |
| S\_RackDefocus | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_RackDefocus.jpg) | 使用 混乱圈 卷积，生成源素材的失焦版本。 这种效果通常比高斯模糊更适合用来模拟真实的失焦相机镜头，因为亮点可以被失焦成干净的形状，而不是被平滑掉。 虹膜的形状可以用 点、点 和 旋转 来控制，而 使用伽玛 参数可以调整模糊的亮点的相对亮度。 |
| S\_RackDfComp | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_RackDfComp.jpg) | 在背景上合成前景，同时以不同的数量对两层进行散焦。 前景的alpha通道被用作哑光。如果提供 中间 输入，它将在前景和背景之间合成。 |
| S\_Sharpen | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Sharpen.jpg) | 放大源素材中的高频部分，如边缘和细节。 增加Sharpen Width（锐化宽度）参数以锐化更多的中段频率，并调整Sharpen Amp（锐化放大器）以控制应用的锐化量。 |
| S\_SoftFocus | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_SoftFocus.jpg) | 将源文件的模糊版本与原始文件结合起来，产生 软聚焦 的效果。 调整宽度和混合参数以提供不同类型的外观。 |
| S\_ZBlur | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_Zap.jpg) | 使用ZBuffer输入的深度值对源素材的区域进行不同程度的模糊。 将输入的深度分成若干层，并根据每层的深度以不同的数量模糊它们。 线性雾化也可以混合到结果中。 要使用这个效果，首先根据你的Z缓冲区设置ZBuffer:Black Is Near或White Is Near，然后调整焦点深度和景深参数以获得你想要的效果。 为了帮助设置聚焦深度，你可以使用Show: 在焦点区。 |
| S\_ZConvolve | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_ZapTo.jpg) | 使用一个内核对源剪辑的区域进行卷积，该内核使用ZBuffer输入的深度值进行放大或缩小。将输入分成若干层，根据与焦点深度的距离和景深，应用不同大小的卷积模糊。 这类似于ZDefocus，但有一个来自剪辑的虹膜形状（或内核）。 |
| S\_ZDefocus | Blur+Sharpen | ![](https://cdn.yuelili.com/2021/sapphire_ZBlur.jpg) | 使用来自ZBuffer输入的深度值，以不同的数量对源素材的区域进行散焦。将输入分成若干深度层，根据每层的深度应用不同的散焦量。 要使用这个效果，首先根据你的Z缓冲区设置ZBuffer:Black Is Near或White Is Near，然后调整Focus Depth和Depth Of Field参数，以获得你想要的效果。 为了帮助设置聚焦深度，你可以使用Show: 在焦点区。 |
| S\_BrushChalk | BrushChalk | ![](https://cdn.yuelili.com/2021/sapphire_BrushChalk.jpg) | 通过将不同大小和方向的笔触分层来模拟粉笔画的外观。这种效果可以使用以下画笔之一：毡尖、飞溅、水彩、点画、铅笔、粉笔、海绵、飞溅、圆形或立方体。此外，还有调整形状、大小、方向、密度、照明和阴影的控制。 |
| S\_BrushOil | BrushChalk | ![](https://cdn.yuelili.com/2021/sapphire_BrushOil.jpg) | 通过将不同大小和方向的笔触分层，模拟出油画的外观。这种效果可以使用以下画笔之一：毡尖、飞溅、水彩、点画、铅笔、粉彩、海绵、飞溅、圆形或立方体。此外，还有用于调整形状、大小、方向、密度、照明和阴影的控制。 |
| S\_EdgeFlash | Composite | ![](https://cdn.yuelili.com/2021/sapphire_EdgeFlash.jpg) | 将正面素材的光晕添加到背面素材上，反之亦然，然后将正面素材合成到背面素材上。 这可以用来使合成物看起来更自然，光在各层之间闪动，就像在胶片上一起曝光一样。 |
| S\_Layer | Composite | ![](https://cdn.yuelili.com/2021/sapphire_Layer.jpg) | 使用各种混合操作之一将前景图像分层在背景上。 每个输入的颜色也可以用光、暗和饱和度参数来调整。 |
| S\_MathOps | Composite | ![](https://cdn.yuelili.com/2021/sapphire_MathOps.jpg) | 使用各种数学运算之一来合并两个片段。 每个输入的颜色也可以用光、暗和饱和度参数来调整。 |
| S\_MatteOps | Composite | ![](https://cdn.yuelili.com/2021/sapphire_MatteOps.jpg) | 增长、缩小或增加源输入的alpha通道的噪音。 这对消除色度键的蓝色或绿色溢出很有用。 |
| S\_MatteOpsComp | Composite | ![](https://cdn.yuelili.com/2021/sapphire_MatteOpsComp.jpg) | 在前景阿尔法通道的边缘增加、缩小或添加噪声，然后用这个结果在背景上合成前景。 这对去除色度键中的蓝色或绿色溢出很有用。 |
| S\_ZComp | Composite | ![](https://cdn.yuelili.com/2021/sapphire_ZapFrom.jpg) | 根据两个深度图像的差异，将一个源输入分层在另一个源输入之上或之下。 DepthA输入应该是对应于第一个输入的物体的 z 深度图像，DepthB应该是对应于第二个输入的物体的 z 深度图像。 |
| S\_Dissolve | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_Dissolve.jpg) | 两个输入片段之间的基本交叉淡出。 |
| S\_DissolveAutoPaint | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveAutoPaint.jpg) | 淡入开始片段的 涂刷 版本。 减少绘画的复杂性，直到它只有几种颜色，然后过渡到第二个片段的 油漆刷 版本，然后在颜色和复杂性上增长，直到第二个片段淡入。 |
| S\_DissolveBlur | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveBlur.jpg) | 在两个输入片段之间过渡，同时对每个片段进行模糊处理。 第一个片段被模糊并淡出，而第二个片段未被模糊并淡入。 Dissolve Percent参数应该是动画的，以控制过渡速度。 |
| S\_DissolveBubble | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveBubble.jpg) | 使用气泡翘曲功能在两个输入片段之间进行过渡。 第一个片段被翘起并淡出，而第二个片段被解除翘起并淡入。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveDefocus | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveDefocus.jpg) | 在两个输入片段之间过渡，同时对每个片段进行散焦。 第一个片段被散焦并淡出，而第二个片段被带入焦点并淡入。 Dissolve Percent参数应该是动态的，以控制转换速度。 |
| S\_DissolveDiffuse | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveDiffuse.jpg) | 在两个输入片段之间的过渡，通过在由最大数量决定的区域内扰乱输入片段的像素。 第一个片段被扩散掉，而第二个片段被扩散到位。 Dissolve Percent（溶解百分比）参数应该被动画化以控制过渡速度。 这种效果的像素化外观取决于图像的分辨率，所以建议在处理前测试你的最终分辨率。 |
| S\_DissolveDigitalDamage | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveDigitalDamage.jpg) | 在两个输入片段之间使用溶镜进行过渡，同时对被溶镜的片段应用DigitalDamage。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveDistort | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveDistort.jpg) | 两个输入片段之间的转换，同时使用另一个片段的梯度来扭曲每个片段。 第一个片段被扭曲并淡出，而第二个片段被解除扭曲并淡入。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 注意，必须提供背景输入，否则这个效果将只是执行一个简单的溶解，没有任何变形。 |
| S\_DissolveEdgeRays | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveEdgeRays.jpg) | 使用动画边缘射线在两个输入片段之间进行过渡。 片段相互溶解，边缘射线被添加到结果中。 边缘射线在效果的持续时间内上升和下降。 边缘射线的动画是通过在屏幕上沿一条线移动边缘射线的原点来实现的。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveFilm | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveFilm.jpg) | 使用可选择伽玛的胶片溶解在两个输入片段之间进行过渡。 与普通的溶镜相比，胶片溶镜能更长时间地保留片段中的高光部分。 溶解百分比参数应该是动画化的，以控制过渡速度。 |
| S\_DissolveFlashbulbs | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveFlashbulbs.jpg) | 在两个片段之间溶入时，模拟大量的闪光灯响起。 如果有很多小闪光灯，看起来就像一个体育场的场景。 如果有几个大闪光灯，在名人红地毯的片段中效果很好。 |
| S\_DissolveGlare | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveGlare.jpg) | 使用动画眩光在两个输入片段之间进行转换。 这些片段相互溶入，而眩光被添加到结果中。 眩光的大小和亮度在效果的持续时间内不断上升和下降。 |
| S\_DissolveGlint | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveGlint.jpg) | 使用明亮的闪光在两个输入片段之间进行转换。 这些片段相互溶解，而每个片段都有一个闪光，在效果的持续时间内上升和下降。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveGlintRainbow | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveGlintRainbow.jpg) | 使用明亮的闪光在两个输入片段之间进行过渡。 这些片段相互溶解，而每个片段都有一个闪光点，在效果的持续时间内不断上升和下降。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveGlow | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveGlow.jpg) | 使用一个明亮的闪光在两个输入片段之间过渡。 这些片段相互溶入，而每个片段都有一个闪光，在效果的持续时间内不断上升和下降。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveLensFlare | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveLensFlare.jpg) | 使用动画镜头闪光在两个输入片段之间进行过渡。 两个片段相互溶入，而镜头光斑沿着一条直线移动。 镜头耀斑在效果的持续时间内增长和缩小。 溶解百分比参数应该是动画的，以控制过渡速度。 |
| S\_DissolveLuma | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveLuma.jpg) | 两个输入片段之间的过渡，使用从它们的亮度得出的模式。 一个片段经常出现在另一个片段中。 Dissolve Percent（溶入百分比）参数应该是动态的，以控制转换速度。 |
| S\_DissolvePixelSort | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolvePixelSort.jpg) | 在两个输入片段之间的过渡，同时对溶解的结果进行排序。 |
| S\_DissolvePuddle | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolvePuddle.jpg) | 两个输入片段之间的过渡，同时以圆形的波浪模式进行扭曲。 第一个片段被翘起并渐渐消失，而第二个片段则被解除翘起并渐渐进入。 Dissolve Percent参数应该是动画的，以控制过渡速度。 |
| S\_DissolveRays | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveRays.jpg) | 使用动画射线在两个输入片段之间进行转换。 这些片段相互溶合，射线被添加到结果中。 射线在效果的持续时间内上升和下降。 射线的动画是通过在屏幕上沿着一条线移动射线的原点来实现的。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveShake | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveShake.jpg) | 在两个片段之间的转场是通过对它们应用摇晃的动作，以及快速的溶入。 晃动使用平移、缩放和/或旋转。 它是随机的，但也是可重复的，所以在参数相同的情况下，每次都会产生相同的摇晃动作。 打开Motion Blur，调整Mo Blur Length以获得不同的模糊量。 调整振幅和频率以获得不同的摇动速度和数量。 Rand参数对随机的非周期性摇动进行详细控制，Wave参数调整有规律的周期性摇动。 X、Y、Z和倾斜参数分别控制水平、垂直、缩放和旋转的摇动量。 |
| S\_DissolveSpeckle | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveSpeckle.jpg) | 使用斑点噪声模式在两个输入片段之间进行过渡。 Dissolve Percent参数应该是动画的，以控制过渡速度。 |
| S\_DissolveStatic | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveStatic.jpg) | 在两个输入片段之间使用随机像素静态过渡。 Dissolve Percent（溶解百分比）参数应该被动画化以控制过渡速度。 这种效果的像素化外观取决于图像的分辨率，因此建议在处理前测试你的最终分辨率。 |
| S\_DissolveTiles | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveTiles.jpg) | 在两个输入片段之间进行过渡，同时将每个片段分解成瓦片并扰乱它们。第一个片段破裂并散开，而第二个片段在第一个片段后面凝聚。 Dissolve Percent参数应该是动画的，以控制过渡速度。 |
| S\_DissolveVortex | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveVortex.jpg) | 使用旋涡翘曲功能在两个输入片段之间进行过渡。 第一个片段被翘起并渐渐消失，而第二个片段被解除翘起并渐渐进入。 Dissolve Percent（溶解百分比）参数应该被动画化，以控制过渡速度。 |
| S\_DissolveWaves | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveWaves.jpg) | 在两个输入片段之间使用波浪式翘曲功能进行过渡。 第一个片段被翘起并淡出，而第二个片段被取消翘起并淡入。 Dissolve Percent参数应该是动画的，以控制过渡速度。 |
| S\_DissolveZap | Dissolve Transitions | ![](https://cdn.yuelili.com/2021/sapphire_DissolveZap.jpg) | 使用动画的闪电在两个片段之间过渡。 片段相互溶入，而闪电不断增长。 Dissolve Percent参数应该是动态的，以控制过渡速度。 |
| S\_Distort | Distort Effects | ![](https://cdn.yuelili.com/2021/sapphire_Distort.jpg) | 使用Lens输入片段的梯度对源输入片段进行扭曲。 这可以产生类似光学玻璃的效果，就像通过一个任意形状的镜头观看源素材一样。 当镜头图像只包含一些大胆的形状或一个简单的纹理时，它是最好的证明。 |
| S\_DistortBlur | Distort Effects | ![](https://cdn.yuelili.com/2021/sapphire_DistortBlur.jpg) | 在Lens输入片段的梯度方向上模糊源输入片段。 当镜头图像只包含一些简单的形状时，它的效果最好。 |
| S\_DistortChroma | Distort Effects | ![](https://cdn.yuelili.com/2021/sapphire_DistortChroma.jpg) | 使用Lens输入素材的梯度将源输入的色度进行不同程度的扭曲。 这可以产生类似光学玻璃的效果，就像通过一个任意形状或纹理的棱镜观看源素材一样。 当镜头图像只包含一些简单的粗体形状时，它的效果最好。 |
| S\_DistortRGB | Distort Effects | ![](https://cdn.yuelili.com/2021/sapphire_DistortRGB.jpg) | 使用镜头输入素材的梯度，将源输入的红、绿、蓝三色通道进行不同程度的扭曲。 当镜头图像只包含一些简单的粗体图形时，它是最好的示范。 |
| S\_EdgeColorize | Edge Effects | ![](https://cdn.yuelili.com/2021/sapphire_EdgeColorize.jpg) | 根据源素材的方向，为其边缘分配不同的颜色。 增加 边缘平滑 参数以获得更厚的边缘。 |
| S\_EdgeDetect | Edge Effects | ![](https://cdn.yuelili.com/2021/sapphire_EdgeDetect.jpg) | 在源素材中寻找边缘。 增加 边缘平滑 参数以获得更厚的边缘。 选择Mono或Chroma模式，只显示亮度或彩度的边缘。 |
| S\_EdgeDetectDouble | Edge Effects | ![](https://cdn.yuelili.com/2021/sapphire_EdgeDetectDouble.jpg) | 执行两次边缘检测操作，产生双绞线的边缘效果。 增加边缘平滑参数以获得更厚的边缘。 |
| S\_EdgesInDirection | Edge Effects | ![](https://cdn.yuelili.com/2021/sapphire_EdgesInDirection.jpg) | 查找在指定方向上对齐的源输入的边缘。 增加 边缘平滑 参数以获得更厚的边缘。 |
| S\_Effect | Effect Builder | ![](https://cdn.yuelili.com/2021/sapphire_Effect.jpg) | 一个效果生成器，让你结合多个蓝宝石效果，并从任何效果中加载预设。点击加载预设或编辑效果就可以开始了。 |
| S\_Transition | Effect Builder | ![](https://cdn.yuelili.com/2021/sapphire_TVChannelChange.jpg) | 转场生成器，可以让你结合多个蓝宝石效果来创建新的转场，并可以从任何转场中加载预设。点击加载预设或编辑效果就可以开始了。 |
| S\_Emboss | Embossing | ![](https://cdn.yuelili.com/2021/sapphire_Emboss.jpg) | 使用Bumps输入的亮度作为浮雕图对源素材进行压花。 增加Bumps Smooth参数以获得更大的凸起，并调整Light Dir以从不同角度照亮凸起。 |
| S\_EmbossDistort | Embossing | ![](https://cdn.yuelili.com/2021/sapphire_EmbossDistort.jpg) | 使用Bumps输入作为浮雕图对源素材进行压印和扭曲，并使用Bumps作为 镜头 图像对结果进行扭曲。 增加Bumps Smooth参数以获得更大的凸起，并调整Light Dir以从不同角度照亮凸起。 |
| S\_EmbossGlass | Embossing | ![](https://cdn.yuelili.com/2021/sapphire_EmbossGlass.jpg) | 资源被压印和扭曲，使用Bumps输入作为浮雕地图和镜头图像。 还会进行色度失真，将光谱分离成一个 棱镜 的样子。 增加Bumps Smooth参数以获得更大的凸起，并调整Light Dir以从不同角度照亮凸起。 |
| S\_EmbossShiny | Embossing | ![](https://cdn.yuelili.com/2021/sapphire_EmbossShiny.jpg) | 使用凸点输入作为浮雕地图对源素材进行压花。 使用一个照明模型，其中包括镜面反射的高光。 增加Bumps Smooth参数以获得更大的凸起，并调整Light Dir以从不同角度照亮凸起。 |
| S\_BleachBypass | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_BleachBypass.jpg) | 模拟一个胶片过程，其中银没有从底片中去除。 其结果是增加了对比度，降低了色彩饱和度。 |
| S\_Diffuse | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_Diffuse.jpg) | 在漫反射量确定的区域内，对源输入的像素进行扰乱。 使用Blur Rel X和Y参数来获得更多的水平或垂直漫射方向。 这种效果的像素化外观取决于图像的分辨率，所以建议在处理前测试你的最终分辨率。 |
| S\_FilmDamage | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_FilmDamage.jpg) | 模拟受损的胶片，有许多选项，包括灰尘、毛发、污渍、划痕、失焦、闪烁和晃动。 每个选项都有一个主控制和一组详细的控制，用于调整该类型的损坏的外观。 |
| S\_FilmEffect | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_FilmEffect.jpg) | 提供一个物理上准确的胶片曝光和处理模型，使你的视频片段看起来像用特定的胶片拍摄的。 它可以去除现场伪影，为特定的胶片类型进行颜色校正，添加胶片颗粒，并应用发光或软聚焦效果。 可以使用Scale CC和Grain Amp参数选择性地禁用颜色校正和颗粒。 |
| S\_Grain | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_Grain.jpg) | 在源素材上添加颜色和/或单色颗粒。 振幅和频率参数允许独立调整所有颜色的纹理，每个颜色通道，或黑白纹理。 |
| S\_GrainStatic | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_GrainStatic.jpg) | 在源素材的每个像素上添加给定振幅的彩色和/或单色随机噪声。 与其他颗粒效果不同的是，像素之间的颗粒没有连贯性，所以产生的外观会因不同的输出分辨率而不同。 |
| S\_Vignette | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_Vignette.jpg) | 使源素材的边界区域变暗，以创造一个晕染效果。 使用平方度、半径和边缘软度参数来影响晕染的形状。 使用不透明度和颜色参数来调整其强度和颜色。 |
| S\_VintageColor2Strip | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_VintageColor2Strip.jpg) | 模拟20世纪20年代的老式彩色2条胶片工艺。 场景被曝光两次，通过红色和绿色滤镜，在单色胶片条的交替框架上。 然后，红色拷贝被染成红色染料，绿色拷贝被染成青色。 这两条带子被背靠背地粘在一起，形成最终的印刷品。 其结果主要包含红色和绿色，还有一些来自染料中蓝色成分的合成蓝色。 |
| S\_VintageColor2Strip-2 | Film Effects & Grain |  | 这种效果模拟了两种滤色片的颜色和两种染料的颜色，也可以添加颗粒和颜色校正。 |
| S\_VintageColor3Strip | Film Effects & Grain | ![](https://cdn.yuelili.com/2021/sapphire_VintageColor3Strip.jpg) | 模拟1935年至1955年的彩色三条线胶片工艺。 三条色彩是一种减法工艺，通过彩色滤光片曝光三条独立的胶片，然后根据原始记录的密度在拷贝上应用互补的彩色染料。 这种工艺被用于许多电影，如《绿野仙踪》、《幻想曲》和《飘》。 |
| S\_VintageColor3Strip-2 | Film Effects & Grain |  | 现代彩色胶片在乳剂层有更广泛的滤色，所以这种效果模拟了那个时代更窄的滤色和更清晰的彩色染料，使其具有特有的活力。 这个效果还可以添加颗粒和颜色校正。 |
| S\_Flicker | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_Flicker.jpg) | 随着时间的推移，源素材的颜色会以不同的数量进行缩放，以达到闪烁的效果。 闪烁的模式可以是随机的，也可以是周期性的波，或者是两者的组合。 |
| S\_FlickerMatch | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerMatch.jpg) | 使用第二个匹配剪辑的闪烁，将闪烁添加到源剪辑中。 例如，一个片段可以与另一个片段中的闪光灯同步变亮。 要使用这个效果，首先要把矩形的角放在匹配剪辑的一个区域上，这个区域有你想要复制的亮度变化。 一个中间或浅灰色的区域是最好的。 然后选择一个你希望源亮度不变的帧，并点击Set Match Level按钮。 当其他帧被处理时，源亮度将被矩形内的平均匹配亮度所缩放，相对于匹配级别。 |
| S\_FlickerMatchColor | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerMatchColor.jpg) | 使用第二个匹配片段的颜色变化来增加源片段的颜色变化。 类似于FlickerMatch，但这个过程是应用于每个颜色通道。 要使用这个效果，首先要把矩形的角放在匹配素材的一个区域上，这个区域有你想要复制的颜色变化。 一个中间或浅灰色的区域是最好的。 然后选择一个你希望源颜色不变的帧，并点击Set Match Level按钮。 当你处理其他帧时，源颜色将被矩形内的平均匹配颜色所缩放，相对于匹配颜色。 |
| S\_FlickerMatchMatte | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerMatch.jpg) | 在Matte指定的区域内，使用第二个Match素材的闪烁，将闪烁添加到Source素材中。要使用这个效果，选择一个你希望源亮度不变的帧，然后点击Set Match Level按钮。 当其他帧被处理时，源亮度将被相对于匹配级别的哑光内的平均匹配亮度所缩放。 |
| S\_FlickerMchMatteColor | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerMatchColor.jpg) | 在Matte指定的区域内，使用第二个Match剪辑的颜色变化来为Source剪辑添加颜色变化。 要使用这个效果，选择一个你希望源颜色不变的帧，然后点击Set Match Color按钮。 当其他帧被处理时，源的颜色将被相对于匹配色的哑光内的平均匹配色所缩放。 |
| S\_FlickerRemove | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerRemove.jpg) | 去除源素材的时间性闪烁。例如，具有不均匀曝光时间的旧片段可以用这个效果来平滑处理。 要使用这个效果，首先将矩形的四角放在平均亮度应保持不变的区域上。 一个中间或浅灰色的区域是最好的。 然后在矩形内选择一个具有所需亮度的源帧，并点击设置保持水平按钮。当其他帧被处理时，它们的亮度将被缩放，使矩形内的平均亮度等于保持水平。 你可以将不同的保持水平值作为时间的关键帧，以考虑到理想的亮度变化。 |
| S\_FlickerRemoveColor | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerRemoveColor.jpg) | 移除源素材中的时间性色彩变化。 类似于FlickerRemove，但这个过程是应用于每个颜色通道。 要使用这个效果，首先将矩形的角定位在一个平均颜色应该保持不变的区域。一个中间或浅灰色的区域是最好的。 然后在矩形内选择一个具有所需颜色的源帧，并点击设置保持颜色按钮。 当其他帧被处理时，它们的颜色将被缩放，使矩形内的平均颜色等于保持颜色。 |
| S\_FlickerRemoveMatte | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerRemove.jpg) | 移除源素材的时间性闪烁，使用哑光素材来指定平均亮度应保持不变的区域。 要使用这个效果，选择一个在Matte中具有所需亮度的源帧，然后点击Set Hold Level按钮。 当其他帧被处理时，它们的亮度将被缩放，使哑光区域的平均亮度等于保持水平。 你可以将不同的保持水平值作为时间的关键帧，以考虑到理想的亮度变化。 |
| S\_FlickerRmMatteColor | Flicker Tools | ![](https://cdn.yuelili.com/2021/sapphire_FlickerRemoveColor.jpg) | 移除源素材中的时间性颜色变化，使用哑光素材来指定平均颜色应该保持不变的区域。 要使用这个效果，选择一个在Matte中具有所需颜色的源帧，然后点击Set Hold Color按钮。 当其他帧被处理时，它们的颜色将被缩放，以便在Matte内的平均颜色等于保持颜色。 |
| S\_Aurora | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_Aurora.jpg) | 沿着用户控制的花键生成一个双色的光漩涡，让人想起北极光（Northern Lights）。 |
| S\_Clouds | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_Clouds.jpg) | 生成一个程序性的噪声纹理。 使用频率参数来放大和缩小纹理。 Shift Speed参数使纹理随时间自动平移。 |
| S\_CloudsColorSmooth | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_CloudsColorSmooth.jpg) | 生成一个全彩的云彩纹理。 程序性噪声纹理是为每个红、绿、蓝输出通道独立生成的。 移位速度参数使纹理随时间自动平移。 |
| S\_CloudsMultColor | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_CloudsMultColor.jpg) | 生成一个像S\_Clouds一样的程序噪声纹理，并使用一个额外的彩色噪声纹理对颜色进行着色。 移位速度参数会使纹理随时间自动平移。 |
| S\_CloudsPerspective | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_CloudsPerspective.jpg) | 生成一个程序性噪声纹理，并将其转换到一个具有透视功能的三维平面上。 调整Latitude（纬度）、Swing（摆动）和Roll（滚动）参数，使图像在不同的轴上旋转，每个轴，使用Frequency（频率）参数来放大和缩小纹理。 Shift Speed使纹理随时间自动平移。 |
| S\_CloudsPsyko | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_CloudsPsyko.jpg) | 生成程序性噪声纹理，并将其通过着色处理。 移位速度参数使图案随时间自动平移，而相位速度使颜色随时间旋转。 |
| S\_CloudsVortex | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_CloudsVortex.jpg) | 生成一个扭曲成涡旋的程序性噪音纹理。 漩涡速度参数使漩涡的旋转量随时间自动生成动画。 |
| S\_LaserBeam | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_LaserBeam.jpg) | 模拟科幻小说风格的激光枪的光束。光束在若干帧内从源点移动到目标点。也可以添加一个透视效果。 |
| S\_Luna | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_Luna.jpg) | 渲染地球上的月亮；你可以调整相位和颜色，并添加大气效果。 |
| S\_MuzzleFlash | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_MuzzleFlash.jpg) | 模拟枪支发射时产生的闪光和烟雾。可以模拟几种类型的枪的闪光。所有的枪都有主闪光，带有压制器的枪可能有次级闪光。枪支可以很容易地被反复发射。 注意：这个效果中的所有时间单位都是帧。 |
| S\_NightSky | Fractal Clouds | ![](https://cdn.yuelili.com/2021/sapphire_NightSky.jpg) | 从一个主要城市或指定的经纬度看，生成一个逼真的星空。 星星是使用一个星星数据库生成的，因此主要的星座在预期的地方是可见的。 调整星等限制以看到更多的星星。对Minute参数进行动画处理，使星星随着时间真实地移动。 |
| S\_Glare | Glints & Glares | ![](https://cdn.yuelili.com/2021/sapphire_Glare.jpg) | 在源素材比阈值更亮的地方，合成彩虹光晕和/或闪电般的光线。 降低阈值参数可以在更多的区域产生眩光。 使用风格菜单来选择不同的眩光类型。 将Glare Res参数设置为1/2，以获得更快的渲染和略微柔和的眩光。 使用Convolve选项可以获得更平滑的结果。 眩光在有少数亮点的黑暗图像上观察效果最好。 |
| S\_Glint | Glints & Glares | ![](https://cdn.yuelili.com/2021/sapphire_Glint.jpg) | 在源素材比阈值更亮的地方产生星形的闪光。 降低阈值参数以在更多的区域产生闪光。 调整大小和亮度参数以产生不同类型的闪光。 在有几个亮点的暗色图像上观察到的闪光效果最好。 |
| S\_GlintRainbow | Glints & Glares | ![](https://cdn.yuelili.com/2021/sapphire_GlintRainbow.jpg) | 在源剪辑比阈值更亮的地方产生星形彩虹色的闪光。 降低阈值参数可以在更多的区域产生闪光。 调整移出、大小和亮度参数以产生不同类型的闪光。 在有几个亮点的黑暗图像上观察到的闪光是最好的。 |
| S\_Glow | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_Glow.jpg) | 从源素材中比给定阈值更亮的区域生成闪光。 提高阈值参数以在更少的区域产生闪光。 调整Width RGB参数以产生具有不同颜色落差的发光，调整Width XY参数以产生水平或垂直的发光。 |
| S\_GlowAura | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowAura.jpg) | 按照源素材的梯度生成放射状的彩色光环线。 提高阈值参数以在更少的区域产生光晕。 调整宽度、频率、相位和扭曲参数，产生具有不同光环模式的发光。 |
| S\_GlowDarks | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowDarks.jpg) | 源素材中比给定阈值更暗的区域会被模糊化，并与输入的素材结合起来，形成深色的烟熏效果。 调整暗度、宽度和阈值参数，以获得不同类型的外观。 |
| S\_GlowDist | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowDist.jpg) | 根据与源输入的边缘的距离生成辉光。 输入图像中的任何边缘，只要亮度超过给定的阈值，就会在边缘的暗处产生同样明亮的光晕。 这在用于具有黑暗背景的图像上时观察效果最好。 |
| S\_GlowEdges | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowEdges.jpg) | 从源素材的边缘创建发光。 这与默认的发光不同，小的或薄的物体在其边缘产生的发光与大物体一样多。 另外，发光的颜色不受源素材颜色的影响。 |
| S\_GlowNoise | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowNoise.jpg) | 从源素材中比给定阈值更亮的区域生成发光。 辉光也被固体噪声纹理所衰减，使其具有噪声或颗粒效果。 如果Jitter Frames参数是正数，噪声将在每一帧中重新生成，以获得晕眩的外观。 如果Jitter Frames为零，两个噪声纹理将被组合起来，并以取决于Spread Speed的速率相互滑动。 |
| S\_GlowOrthicon | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowOrthicon.jpg) | 源素材在比给定阈值更亮的部分周围区域变暗，以产生 正交 或 暗光 的外观。 降低阈值参数，在更多的区域产生正字形效果。 调整暗度和宽度参数，以提供不同类型的外观。 |
| S\_GlowRainbow | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowRainbow.jpg) | 根据与源输入的边缘的距离，生成彩虹色的光芒。 输入图像中的任何边缘，只要亮度超过给定的阈值，就会在边缘的暗部产生同样的光亮。 当使用在有黑暗背景的图像上时，这是最好的观察。 |
| S\_GlowRings | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_GlowRings.jpg) | 在源素材中比给定阈值更亮的区域周围产生彩色环状的发光。 提高阈值参数可以在更少的区域产生光晕。 调整RGB的宽度和厚度参数以产生不同颜色模式的光芒，调整XY的宽度参数以产生水平或垂直的光芒。 |
| S\_UltraGlow | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_UltraGlow.jpg) | 从源素材的明亮区域生成各种发光的光线。 提高阈值参数以在较少的区域产生发光。 调整Width RGB参数以产生具有不同颜色衰减的辉光，调整Width XY参数以产生水平或垂直辉光。 调整Glow Falloff和Glow Bias参数来控制falloff距离和最热区域的延伸范围。 调整After Glow参数，在主辉光的结果上产生一个二次辉光。 可以选择增强边缘，或在源素材上添加高光，或将结果与大气噪声结合起来。 |
| S\_ZGlow | Glow Effects | ![](https://cdn.yuelili.com/2021/sapphire_ZDefocus.jpg) | 根据ZBuffer输入的深度值，以不同的宽度发光源片段的区域。 将输入分成若干层，并根据近宽度、远宽度、近亮度和远亮度参数应用不同的发光量。 |
| S\_Gradient | Gradients, Grid & Shape | ![](https://cdn.yuelili.com/2021/sapphire_Gradient.jpg) | 使用给定的开始和结束位置及颜色在屏幕上制作一个平滑的颜色渐变，然后可选择将渐变与背景剪辑结合起来。 增加 添加噪音，以减少由于颜色量化造成的梯度中的带状伪影。 |
| S\_GradientMulti | Gradients, Grid & Shape | ![](https://cdn.yuelili.com/2021/sapphire_GradientMulti.jpg) | 使用多个控制点在屏幕上生成一个平滑的多色渐变，并可选择将渐变与背景剪辑相结合。 |
| S\_GradientRadial | Gradients, Grid & Shape | ![](https://cdn.yuelili.com/2021/sapphire_GradientRadial.jpg) | 给予中心、内半径和外半径参数，在一个椭圆形状中制作一个平滑的径向颜色渐变，并可选择将渐变与背景素材结合起来。 增加 添加噪音，以减少由于颜色量化而在渐变中出现的带状伪影。 |
| S\_Grid | Gradients, Grid & Shape | ![](https://cdn.yuelili.com/2021/sapphire_Grid.jpg) | 生成一个线条的网格，并将其与背景素材相结合。 调整Latitude（纬度）、Swing（摇摆）和Roll（滚动）参数，在不同的轴上旋转网格，并调整Shift和Z Dist来进行平移和缩放。 |
| S\_Shape | Gradients, Grid & Shape | ![](https://cdn.yuelili.com/2021/sapphire_Shape.jpg) | 在图像中画出一个形状。 它可以给出各种各样的形状，从多边形和圆形到星星、花朵形状和旋转的海星形状。 要看的主要参数是点数、尖度、圆度和漩涡。 |
| S\_KaleidoDiamonds | KaleidoDiamonds | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoDiamonds.jpg) | 将源素材反射成一个钻石的图案。 内部 参数在源图像被反射到图案之前对其进行转换。 Center（中心）和Z Dist（Z）改变整个结果，包括反射图案，而Rotate（旋转）只影响反射的 镜面。 |
| S\_KaleidoOct | KaleidoDiamonds | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoOct.jpg) | 将源素材反射成一个八角形的直角三角形图案。 内部 参数在源图像被反射到图案之前对其进行转换。 中心 和 Z 参数对包括反射图案在内的整个结果进行转换，而 旋转 参数只影响反射的 镜子。 |
| S\_KaleidoSquares | KaleidoDiamonds | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoPolar.jpg) | 将源素材反射成一个正方形的图案。 内部 参数在源图像被反射到图案之前对其进行转换。 Center（中心）和Z Dist（Z）对包括反射图案在内的整个结果进行转换，而Rotate（旋转）只影响反射的 镜子。 |
| S\_KaleidoTriangles | KaleidoDiamonds | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoRadial.jpg) | 将源素材反射成一个等边三角形的图案。 内部 参数在源图像被反射到图案之前对其进行转换。 Center（中心）和Z Dist（Z）对整个结果进行转换，包括反射图案，而Rotate（旋转）只影响反射的 镜面。 |
| S\_FlysEyeCircles | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_FlysEyeCircles.jpg) | 将图像分解成圆圈状的瓦片，并在每个形状中转换图像，以创造一个飞眼的效果。 重叠 选项允许圆圈在重叠的地方以不同方式组合。 内部 参数在源图像被平铺到图案之前对其进行转换，而 平铺 参数则对整个苍蝇眼图案进行转换。 |
| S\_FlysEyeHex | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_FlysEyeHex.jpg) | 将图像分解成六边形的瓦片，并在每个形状内变换图像，以创造苍蝇眼的效果。 增加边缘柔和度，使瓦片之间的重叠更加平滑。 内部 参数在源图像被贴入图案前对其进行变换，贴图 参数对整个苍蝇眼图案进行变换。 |
| S\_FlysEyeRect | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_FlysEyeRect.jpg) | 将图像分解成矩形瓦片，并在每个形状中变换图像，以创造苍蝇眼的效果。 内部 参数在源图像被贴入图案之前对其进行变换，贴图 参数对整个苍蝇眼图案进行变换。 |
| S\_KaleidoPolar | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoSquares.jpg) | 将源片段扭曲成圆盘状，并在径向上进行反射，就像通过一个反射圆柱体观看一样。 |
| S\_KaleidoRadial | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_KaleidoTriangles.jpg) | 模拟一个传统的2或3个镜面的万花筒。 你通过镜子之间的角度看到一个饼状的源片，而在图像的其他部分则是镜子反射的副本。 使用 切片 参数来控制你在中心点周围看到多少份源的饼片。 |
| S\_TileScramble | Kaleidoscopes | ![](https://cdn.yuelili.com/2021/sapphire_TimeDisplace.jpg) | 将图像分解成矩形瓦片，并在每个瓦片中移动图像，以创造出像一堵随机取向的小镜子反映源图像的效果。 移位的数量和方向是可控的。 |
| S\_LensFlare | LensFlares | ![](https://cdn.yuelili.com/2021/sapphire_LensFlare.jpg) | 在背景素材上渲染一个镜头炫光图像，在热点和支点位置之间对齐各种炫光元素。 使用镜头菜单来选择不同类型的镜头炫光。 |
| S\_LensFlareAutoTrack | LensFlares | ![](https://cdn.yuelili.com/2021/sapphire_LensFlareTrack.jpg) | 在背景素材上渲染一个或多个镜头炫光图像，在热点和支点位置之间对齐各种炫光元素。 在LensFlare的这个自动跟踪版本中，热点会自动定位在背景素材中最亮的地方。增加 自动模糊 将使输入在找到最亮的位置之前就被平滑，并有助于消除次要亮点的影响。 |
| S\_LightLeak | LensFlares | ![](https://cdn.yuelili.com/2021/sapphire_LightLeak.jpg) | 渲染抽象的颜色模式，模拟光线从相机机身的缝隙中漏出。漏光由三个不同的元素组成，可以单独调整。 |
| S\_EdgeRays | Light Beams | ![](https://cdn.yuelili.com/2021/sapphire_EdgeRays.jpg) | 生成从输入素材的边缘发出的光束。 你可以提供一个哑光输入来选择性地调整光线的颜色。 如果哑光类型被设置为颜色，你也可以使用哑光输入来对不同区域的光线进行不同的着色。 将Rays Res参数设置为1/2，可以用稍软的射线快速渲染。 |
| S\_Rays | Light Beams | ![](https://cdn.yuelili.com/2021/sapphire_Rays.jpg) | 产生从源剪辑的明亮区域发出的光束。 降低阈值参数以产生更多区域的光线，或提高阈值以只产生最亮区域的光线。 将Rays Res参数设置为1/2，以便在光线略微柔和的情况下快速渲染。 |
| S\_Streaks | Light Beams | ![](https://cdn.yuelili.com/2021/sapphire_Streaks.jpg) | 运动使光源的明亮区域在 从 和 到 的转换之间模糊成条纹状。 这可以用来创造一个扩展的电影曝光效果，或模拟柔和的光束。 从和到的参数不是指时间。 它们描述的是空间中的两个变换，决定了应用于每一帧的模糊风格。 |
| S\_UltraZap | Lightning Zaps | ![](https://cdn.yuelili.com/2021/sapphire_UltraZap.jpg) | 沿着花键生成闪电，并在背景上渲染它们。 增加Vary Endpoint来分散闪电的末端。 调整发光颜色以获得不同颜色的结果。 调整循环速度参数，使其随着时间的推移穿越花键。 |
| S\_Zap | Lightning Zaps | ![](https://cdn.yuelili.com/2021/sapphire_Zebrafy.jpg) | 在两点之间产生闪电，并在背景上进行渲染。 增加闪电的数量以产生电浆的效果。 增加Vary Endpoint来分散闪电的末端。 调整发光颜色以获得不同颜色的效果。 摇摆速度参数使闪电随时间自动起伏。 |
| S\_ZapFrom | Lightning Zaps | ![](https://cdn.yuelili.com/2021/sapphire_ZebrafyColor.jpg) | 从FromObj输入片段中的物体边缘向外生成多个闪电，并在背景输入上渲染。 使用Show:Edges选项来查看源边缘，同时调整Threshold和Blur From Obj参数。 |
| S\_ZapTo | Lightning Zaps | ![](https://cdn.yuelili.com/2021/sapphire_ZFogExponential.jpg) | 从给定的点到ToObject输入素材中的物体边缘生成一个分叉的闪电，并在背景输入上进行渲染。 使用Show:Edges选项来查看目标边缘，同时调整Threshold和Blur To Obj参数。 |
| S\_BokehLights | Other Lighting | ![](https://cdn.yuelili.com/2021/sapphire_BokehLights.jpg) | 产生随机的、失焦的光线，在屏幕上移动。 |
| S\_DropShadow | Other Lighting | ![](https://cdn.yuelili.com/2021/sapphire_DropShadow.jpg) | 使用前景的alpha通道或可选的Matte在背景素材上生成一个阴影，然后将前景与背景合成，得到最终结果。 |
| S\_Flashbulbs | Other Lighting | ![](https://cdn.yuelili.com/2021/sapphire_Flashbulbs.jpg) | 模拟大量的闪光灯响起。 用许多小闪光灯，可以看起来像一个体育场的场景。 用几个大的闪光灯，在名人红地毯的片段上效果很好。 |
| S\_Light3D | Other Lighting | ![](https://cdn.yuelili.com/2021/sapphire_Light3D.jpg) | 用最多4个单独控制的光源执行3D重新照明。 光源输入通常是来自三维渲染器的环境或漫反射通道，显示表面的颜色。 法线矢量输入决定了每个像素的表面方向。 光源和法线应该由3D程序一起生成，这样它们才会匹配。 |
| S\_SpotLight | Other Lighting | ![](https://cdn.yuelili.com/2021/sapphire_SpotLight.jpg) | 使用一个或两个聚光灯点亮输入片段。对于每个启用的灯光，使用给定的光源位置、瞄准位置和光束角度来计算三维光锥与图像平面的交点。 环境光也可以被应用来均匀地影响整个源图像。 通过调整所提供的参数，可以创建各种各样的照明形状。 |
| S\_BandPass | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_BandPass.jpg) | 使用带通滤波器生成一个类似X射线的效果。 以不同的宽度进行两次模糊，其结果是以灰度值缩放和偏移的差异。 高于和低于截止点的频率被衰减，只留下中间的频率带。 |
| S\_DigitalDamage | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_DigitalDamage.jpg) | 模拟不良的数字电视传输，有许多选项，包括定格、移动和流动的块、各种块状的噪声和像素化。 可以给出类似于MPEG流错误、数字掉线和卫星馈电数据损坏的外观。 |
| S\_DogVision | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_DogVision.jpg) | 生成输入图像的双色通道版本，就像狗的有限色觉系统可能感知的那样。 人类有三个颜色受体（红色、绿色和蓝色），而狗只有两个受体（黄色和蓝色）。 |
| S\_Grunge | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_Grunge.jpg) | 模拟许多不同种类的grunge，包括灰尘、污渍、斑点、污垢、划痕和油漆。最多可以组合三种不同的污垢。 有用于调整所有灰度的主控制，以及一套用于调整每个灰度系列外观的详细控制。 |
| S\_JpegDamage | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_JpegDamage.jpg) | 创建一个受Jpeg压缩假象和错误影响的源输入版本。 这可以用来赋予低质量数字传输的各种外观。 提供了三种操纵图像的方法：可以调整Jpeg质量，可以缩放各种内部频率，可以引入随机解压错误。 在所有情况下，降低分辨率系数以创建更大、更明显的Jpeg块也是有用的。 |
| S\_Mosaic | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_Mosaic.jpg) | 生成源片段的像素化版本。 使用 像素频率 和 像素相对高度 参数调整块的大小和形状。 增加平滑颜色参数，使附近的像素块的颜色更加一致，并且随着时间的推移减少闪烁。 |
| S\_PixelSort | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_PixelSort.jpg) | 将超过阈值的像素沿着各种模式排列的线进行排序。 模式包括平行线、从中心点辐射出来的线和圆线。 |
| S\_PseudoColor | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_PseudoColor.jpg) | 对源图像进行着色。 色调是根据源图像的亮度计算出来的。 |
| S\_PsykoBlobs | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_PsykoBlobs.jpg) | 将源素材与 blob 形状的区域结合起来，然后将它们通过一个着色过程。 相位速度参数使颜色随时间自动旋转。 |
| S\_PsykoStripes | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_PsykoStripes.jpg) | 将源素材与条纹图案结合起来，然后通过一个着色过程。 相位速度参数使颜色随时间自动旋转。 |
| S\_RomanTile | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_RomanTile.jpg) | 基于源片段生成一个马赛克图案。 调整 边缘吸引 参数，使瓷砖的角落偏向源文件中的边缘。 增加Vary Shape参数，以获得一个不太规则的瓷砖图案。 |
| S\_ScanLines | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ScanLines.jpg) | 创建一个带有扫描线模式的源素材版本，类似于彩色电视显示器。 增加 添加噪音 参数，也可以为结果添加颗粒状的效果。 |
| S\_ScanLinesMono | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ScanLinesMono.jpg) | ScanLines的一个单色版本。 创建一个具有扫描线模式的源素材版本，类似于黑白电视监视器。 增加 添加噪音 参数，也可以为结果添加颗粒效果。 |
| S\_Solarize | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_Solarize.jpg) | 将输入素材中比阈值更亮的颜色插入，以产生 日晒 效果。 |
| S\_StripSlide | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_StripSlide.jpg) | 将一个片段分解成条状，并一次将它们滑出屏幕以显示背景。 |
| S\_TVDamage | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_TextureChromaSpiral.jpg) | 模拟有传输和接收问题的电视，录像机问题和电视硬件困难。 模拟静态、干扰、重影、水平和垂直保持、嗡嗡声、彩色条纹、可见扫描线、录像机快进、掉线、晕染、正视、鱼眼和关闭。 |
| S\_ZFogExponential | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ZComp.jpg) | 使用ZBuffer输入的深度值将雾的颜色混合到源素材中。 雾从Z近处开始，以指数形式增加到Z远处，其速度取决于雾密度。 如果不提供ZBuffer输入将是纯黑色的，所以你应该指定这个输入，这样效果才有用。 |
| S\_ZFogLinear | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ZConvolve.jpg) | 使用来自ZBuffer输入的深度值将雾的颜色混入源素材。 雾量在雾近和雾远之间线性变化，因为深度在Z近和Z远之间变化。 如果不提供ZBuffer输入将是纯黑色的，所以你应该指定这个输入，以使这个效果有任何作用。 |
| S\_Zebrafy | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ZFogLinear.jpg) | 用正弦波调节源素材的亮度，使其具有黑白相间的外观。 |
| S\_ZebrafyColor | Other Stylize | ![](https://cdn.yuelili.com/2021/sapphire_ZGlow.jpg) | 用每个颜色通道的正弦波来调节源素材的亮度，以获得彩色条纹的效果。 |
| S\_CardFlip | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_CardFlip.jpg) | 通过滑动或旋转流出的片段，在两个片段之间进行过渡，以显示其后面的流入片段。数量参数应该是动画的，以控制过渡速度。调整Revolutions和Shift可以得到不同的过渡效果。 |
| S\_FilmRoll | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_FilmRoll.jpg) | 通过垂直滚动一个片段离开屏幕，同时滚动另一个片段，在两个片段之间进行转换，同时应用各种胶片损坏效果，如摇晃、污点、划痕和闪烁。 |
| S\_FlutterCut | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_FlutterCut.jpg) | 通过在两个片段之间快速来回切割来实现过渡，也可以选择插入纯色或倒置的帧。每个片段的剪辑可以在过渡的长度上变长或变短。 |
| S\_HyperPull | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_HyperPull.jpg) | 在溶入背景之前，将前景拉离Z空间，并以一些额外的颜色和抖动效果来增加风格。 |
| S\_HyperPush | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_HyperPush.jpg) | 在溶入背景之前，将前景在Z空间中推近，并具有一些额外的颜色和抖动效果，以增加风格。 |
| S\_ParallaxStripsTransition | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_ParallaxStripsTransition.jpg) | 应用3D折射玻璃条的集合来分解图像。图像在每个条带内移动，条带随时间移动。这些条带逐渐淡入或淡出，因此与源头的过渡是无缝的。 |
| S\_ParallaxStripsTransition-2 | Other Transitions |  | 注意：因为你可以控制条带的大小，所以有可能设置一个过渡，不会完全覆盖二级（传入或传出）片段，所以在开始或结束时有一个 流行。 为了确保平滑过渡，请到过渡的末端，即条带最大的地方，选择显示。条带模式，并调整条带大小以确保它们完全覆盖图像。这样一来，后面的剪辑就不会在那一帧上漏掉了。 或者，点击 确保全覆盖 按钮，条带的数量将自动调整到完全覆盖图像所需的最小值。 |
| S\_StripSlideTransition | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_StripSlideTransition.jpg) | 在两个片段之间进行转换，把它们分成条状，然后一次一次地滑出屏幕，显示出输入的片段。 |
| S\_Swish3D | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_Swish3D.jpg) | 在两个输入片段之间溶入，同时在每个片段上执行3D动作。 在转换过程中，从片段被Zdist、Rotate、Swivel、Tilt、Shift、Scale和Shear参数改变，而到片段则被这些数值的反面改变。 每个图像的整体运动量可以通过Rel Amp From和Rel Amp To参数进行缩放。 |
| S\_SwishPan | Other Transitions | ![](https://cdn.yuelili.com/2021/sapphire_SwishPan.jpg) | 在两个输入片段之间的转换，是通过将一个片段滑出画面，另一个片段滑入画面，并添加运动模糊，使其看起来像一个快速平移。 当转换的持续时间很短时，这种方法效果最好。 |
| S\_Caustics | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_Caustics.jpg) | 模拟光线被弯曲的表面反射或折射时产生的图案。在明亮的阳光下，在游泳池的底部或在水下观察的物体上，经常可以看到苛责的现象。 |
| S\_Sparkles | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_Sparkles.jpg) | 产生一个闪闪发光的场面效果。 调整频率、密度和大小参数以获得不同类型的闪光图案。 使用Matte输入，只在指定区域产生闪光。 |
| S\_SparklesColor | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_SparklesColor.jpg) | 生成一个具有不同颜色的闪光效果的场。 为不同类型的闪光图案调整频率、密度和大小参数。 使用 哑光 输入，只在指定区域产生闪光。 |
| S\_TextureCells | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureFlux.jpg) | 生成程序性细胞形状的图像。 旋转速度参数使细胞中心在每个细胞内随时间旋转。 |
| S\_TextureChromaSpiral | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureFolded.jpg) | 将WarpChroma效果应用于程序生成的噪声纹理，从而创建一个抽象的纹理。 |
| S\_TextureFlux | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureLoops.jpg) | 创建波动的液体或细胞图案的抽象纹理。 Morph Speed参数使图案随时间自动起伏。 |
| S\_TextureFolded | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureMicro.jpg) | 创建类似于折叠的布或液体的抽象纹理，可以通过动画来产生动态的湍流效果。 折叠速度参数使图案随时间自动起伏。 |
| S\_TextureLoops | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureMoire.jpg) | 创建一个重叠的循环形状的抽象纹理。 三组形状可以分别调整，上色，然后组合在一起。 相位速度参数使图案随时间自动变化。 |
| S\_TextureMicro | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureNeurons.jpg) | 生成一个程序性纹理，看起来有点像电子显微镜下粗糙物体的表面 |
| S\_TextureMoire | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureNoiseEmboss.jpg) | 通过把两个同心圆的图案加在一起，创建一个抽象的莫尔纹理。 相位速度和摩尔速度参数使环状物随着时间的推移而自动产生变化。 |
| S\_TextureNeurons | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureNoisePaint.jpg) | 创建一个抽象的纹理，类似于移动的神经细胞卷须。 相位速度和变形速度参数使图案随时间自动变化。 |
| S\_TextureNoiseEmboss | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TexturePlasma.jpg) | 通过将EmbossShiny效果应用于程序生成的噪声纹理，创建一个抽象的纹理。 调整Light Dir来从不同角度照亮图案。 |
| S\_TextureNoisePaint | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureSpots.jpg) | 通过对程序生成的噪声纹理应用AutoPaint效果来创建一个抽象的纹理。 |
| S\_TexturePlasma | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureTiles.jpg) | 创建一个类似于电浆效果的抽象纹理。 相位速度参数使图案随时间自动起伏。 |
| S\_TextureSpots | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TextureWeave.jpg) | 创建一个可以被扭曲和动画化的斑点场。 翘曲速度参数使斑点随着时间的推移被随机的翘曲模式所扭曲。 |
| S\_TextureTiles | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_Threshold.jpg) | 瓷砖（TextureTiles）绘制一个重复的瓷砖图案。 形状可以是六边形、三角形、菱形、星形或它们的变化，取决于Morph参数。 |
| S\_TextureWeave | Procedural Textures | ![](https://cdn.yuelili.com/2021/sapphire_TileScramble.jpg) | 创建一个抽象的纹理，类似于垂直编织的股线。 两组股线，水平和垂直，可以使用频率、八度和速度参数独立调整。 |
| S\_CutToDissolve | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_CutToDissolve.jpg) | 将单个片段中的剪切变成溶入。 不需要头或尾；只需设置剪切点（帧），CutToDissolve就会围绕该点创建一个溶镜。 请注意，这种效果不需要两个片段；只需要一个已经包含剪切的单一片段。 剪切点（Cut Point）参数是实现该效果的关键；该参数两侧的任何帧都将被视为剪切点，而溶镜将围绕它们进行创建。 |
| S\_Feedback | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_Feedback.jpg) | 输入片段的前几帧被转换并与当前帧相结合，以获得受视频反馈启发的各种效果。 每个处理过的帧的输出被储存起来，然后与下一个帧结合。 每当处理任何非连续的帧时，反馈就会被重新初始化：要么是第一帧，要么是重新处理一个给定的帧，要么是跳到另一个帧。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_FeedbackBubble | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_FeedbackBubble.jpg) | 与 反馈 类似，以前的帧与当前的帧结合在一起，同时通过气泡模式进行扭曲。 每当处理任何非连续的帧时，反馈都会被重新初始化：要么是第一帧，要么是重新处理给定的帧，要么是跳转到另一帧。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_FeedbackDistort | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_FeedbackDistort.jpg) | 输入片段的前几帧被给定的Lens输入片段的梯度所扭曲，并与当前帧相结合，从而得到各种可能的效果。 每个处理过的帧的输出被储存起来，然后与下一帧结合。 每当处理任何非连续的帧时，反馈都会被重新初始化，要么是第一帧，要么是重新处理一个给定的帧，要么是跳到另一个帧。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_FieldRemove | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_FieldRemove.jpg) | 自适应地去除运动区域的视频场交错伪影，而不模糊图像的静止部分。 内部产生一个 运动磨砂，运动区域被去隔行，通常会损失垂直分辨率，但静止区域不会被去隔行，应该保持清晰。 |
| S\_FreezeFrame | Time Effects |  | 冻结每一个冻结帧的持续时间的运动。 例如，如果冻结帧数为5，源帧数为。1 2 3 4 5 6 7 8 9 10 11......输出帧将是。1 1 1 1 1 6 6 6 6 6 11... |
| S\_GetFrame | Time Effects |  | 为每个目标帧从源剪辑中检索一个指定的帧。 这是指通过动画化Get Frame的值，以任意的方式加快、减慢或逆转输入片段的速度。 |
| S\_JitterFrames | Time Effects |  | 每个输出帧都会收到当前帧加减抖动帧距之间的随机帧，抖动是随机的，但可重复。 |
| S\_MotionDetect | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_MotionDetect.jpg) | 对于每一帧，找出该帧与之前的一帧之间的差异。 |
| S\_NearestColor | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_NearestColor.jpg) | 从输入片段的帧中收集最接近给定匹配颜色的像素颜色。 例如，这可以从一个有物体在蓝色或绿色屏幕背景上移动的片段中创建一个纯背景的图像。 它也可以用来积累一个移动物体在非彩色背景上的颜色。 每当处理任何非连续的帧时，收集的颜色都会被重新初始化，无论是第一帧、重新处理给定的帧，还是跳转到另一帧。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_RandomEdits | Time Effects |  | 随机地重新编辑整个源片段。 洗牌是随机的，但可以重复。 |
| S\_RepairFrames | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_RepairFrames.jpg) | 用周围帧的时间扭曲版本替换片段中的一个或多个帧，以修复这些帧。 |
| S\_Retime | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_Retime.jpg) | 使用基于光学流的运动估计和帧插值对剪辑进行重新计时。 |
| S\_ReverseClip | Time Effects |  | 颠倒剪辑的帧序，也可选择颠倒每一帧的场。 |
| S\_ReverseEdits | Time Effects |  | 独立地反转输入片段。 例如，如果编辑帧长度是5，而输入的剪辑帧是。1 2 3 4 5 6 7 8 9 10 11......输出的帧将是。5 4 3 2 1 10 9 8 7 6 15... |
| S\_TimeAverage | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_TimeSlice.jpg) | 每个输出帧是多个输入帧的平均值：从当前帧开始，回溯到给定数量的先前帧。这类似于轨迹效应，只是范围内的所有帧都被平均加权，而不是渐渐消失，所以轨迹的终点是突然的。 每一帧只占总亮度的1/n，所以在黑暗背景下快速移动的物体可能看起来很暗。 每当处理任何非连续的帧时，平均数都会被重新初始化：要么是第一帧，要么是重新处理一个给定的帧，要么是跳到另一个帧。 你必须连续处理一个片段的多个帧来观察这个效果，在渲染前清除图像缓存有时是必要的。如果在场上处理，这种效果就不能正常工作。 请确保在帧上进行处理。 |
| S\_TimeDisplace | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_TimeWarpRGB.jpg) | 根据Matte输入的亮度值，将源素材在时间上进行不同程度的置换。 |
| S\_TimeSlice | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_Tint.jpg) | 将输出帧分成片断，每个片断从源素材中接收一个不同的帧。 这个效果的一个例子是让一个转动的物体扭曲成螺旋状，而不是僵硬地旋转。 这些片断的方向取决于片断方向，并接收片断号的正负一半之间的相对帧数。 例如，如果当前帧数是30，切片方向是-90度，切片数是12，帧偏移是0，结果将由水平切片组成，从下到上大约包含30-6到30+6帧。 |
| S\_TimeWarpRGB | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_Trails.jpg) | 将红、绿、蓝三个通道按不同的量进行时间偏移，以产生时间性的色度失真。 |
| S\_Trails | Time Effects | ![](https://cdn.yuelili.com/2021/sapphire_Transition.jpg) | 输入片段的前几帧与当前帧相结合，产生各种 时间痕迹 效果。 每个处理过的帧的输出被储存起来，然后与下一个帧结合。 每当处理任何非连续的帧，无论是第一帧、重新处理一个给定的帧，还是跳到另一个帧，都会重新初始化这些轨迹。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_TrailsDiffuse | TrailsDiffuse | ![](https://cdn.yuelili.com/2021/sapphire_TriTone.jpg) | 输入片段的前几帧用像素扩散处理，然后与当前帧结合。 每个处理过的帧的输出被储存起来，然后与下一帧结合。 每当处理一个非连续的帧时，轨迹就会被重新初始化，要么是第一帧，要么是重新处理一个给定的帧，要么是跳转到另一个帧。 你必须连续处理一个片段的多个帧来观察效果，在渲染前清除你的图像缓存有时可能是必要的。 |
| S\_TVChannelChange | TVChannelChange | ![](https://cdn.yuelili.com/2021/sapphire_TextureCells.jpg) | 通过模拟老式电视机的频道变化在两个输入片段之间进行过渡。 第一个片段因接收不良而变黑，接着是第二个片段，接收不良。 随着时间的推移，接收情况有所改善，直到只剩下第二个片段。 |
| S\_InfiniteZoom | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_InfiniteZoom.jpg) | 缩放到一个图像的无尽重复的副本，让人想起M.C.Escher的某些画。 对具有透明边缘的剪辑效果最好，如时钟或盘子；或透明中心，如画框。 透明度可以来自源素材的alpha或遮罩。 动画化缩放参数以获得完整的效果。 |
| S\_ParallaxStrips | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_ParallaxStrips.jpg) | 应用3D折射玻璃条的集合来分解图像。图像在每个条带中被移动，条带随时间移动。这些条带逐渐淡入或淡出，因此，与源头的过渡是无缝的。 |
| S\_Shake | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_Shake.jpg) | 通过平移、缩放和/或旋转，在源素材上应用摇晃动作。 晃动是随机的，但可以重复，所以在相同的参数下，每次都会产生相同的晃动。 打开运动模糊，调整Mo Blur Length以获得不同程度的模糊。 调整振幅和频率以获得不同的摇动速度和数量。 Rand参数对随机的非周期性摇动进行详细控制，Wave参数调整有规律的周期性摇动。 X、Y、Z和倾斜参数分别控制水平、垂直、缩放和旋转的摇动量。 |
| S\_StretchFrameEdges | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_StretchFrameEdges.jpg) | 拉伸4X3图像的边缘，同时保留中心，以隐藏16X9画面中的黑柱。 这个效果是将源素材的中间部分进行挤压，因为在16x9的编译器中观看4x3的图像通常会将其拉长以适应。 边缘没有被挤压，所以图像一直延伸到边缘。 图像的左侧和右侧边缘部分会出现水平拉伸。 虽然是为4x3转换而设计的，但它可以适用于任何长宽比。 |
| S\_WarpBubble | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpBubble.jpg) | 通过一个平滑的噪音功能使源素材变形。 这可以用来创造热扩散或水底类型的效果。 移位速度参数使噪声图案随时间自动平移。 调整振幅和频率参数以提供不同类型的扭曲。 |
| S\_WarpBubble2 | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpBubble2.jpg) | 使用两套重叠的气泡图案对源素材进行扭曲。 这可以用来创造热扩散或水底类型的效果。 移位速度参数使噪音图案随时间自动平移。 调整振幅和频率参数可以产生不同类型的扭曲。 |
| S\_WarpChroma | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpChroma.jpg) | 将源素材分离成光谱带，并以不同的数量对其进行扭曲。 红色被 从 转换扭曲，蓝色被 到 转换扭曲，光谱的其他颜色在两者之间。 From和To参数并不是指时间。 它们描述了空间中的两个变换，决定了应用于每种颜色的扭曲顺序。 |
| S\_WarpCornerPin | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpCornerPin.jpg) | 对源图像进行三维透视翘曲，使其四角与指定的四个点对齐。 这对定位源图像在另一个素材中的物体上很有用，如广告牌或电脑屏幕。 |
| S\_WarpDrops | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpDrops.jpg) | 通过从多个中心位置发出的多个同心波的模式来扭曲源剪辑。 Centers（中心）输入素材中的每个区域都比Threshold Cntrs（阈值）的值要亮，会产生一个独立的同心波模式，每个区域的总亮度会影响到这些波的扭曲程度。 如果Centers图像很复杂，产生的中心的数量和位置会对阈值相当敏感。 试试只用纯黑加几个白点作为中心的输入。 如果你只需要一组波浪，你可以用WarpPuddle效果代替。 |
| S\_WarpFishEye | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpFishEye.jpg) | 将源素材的中心扩大，就像通过鱼眼镜头看一样。 调整数量参数，以提供更多或更少的失真。 关掉 包裹 选项，使输入素材的边界以外的部分透明，而不是反射的副本。 |
| S\_WarpMagnify | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpMagnify.jpg) | 放大源文件的一个椭圆区域，创造一个玻璃镜片的折射效果。 |
| S\_WarpPerspective | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpPerspective.jpg) | 用透视法将源素材转换到一个三维平面上。 调整Latitude（纬度）、Swing（摆动）和Roll（滚动）参数，在不同的轴上旋转图像，并调整Shift（移位）和Z Dist（缩放）来平移和缩放。 关闭 包裹 选项，给图像一个单一的不重复的拷贝。 |
| S\_WarpPolar | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpPolar.jpg) | 将源素材翘曲成一个圆盘状。 源图像的垂直方向被映射到内半径和外半径之间，水平方向根据角度重复的数量围绕中心旋转，并以角度抵消。 |
| S\_WarpPuddle | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpPuddle.jpg) | 通过同心波的模式对源素材进行扭曲。 相位速度参数使波浪随着时间自动从中心向外移动。 调整内部和外部半径参数来限制波浪出现的区域。 增加 内 和 外 的软度，使波浪出现和不出现之间的过渡更加平滑。 |
| S\_WarpPuff | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpPuff.jpg) | 根据源素材的梯度对其进行扭曲。 默认情况下，较亮的区域被膨大，较暗的区域被缩小。 这类似于用自身作为镜头对图像应用扭曲效果。 |
| S\_WarpRepeat | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpRepeat.jpg) | 对源输入进行多次转换，并对结果进行平均。 From和To参数不是指时间。 它们描述了空间中的两个变换，决定了应用于每一帧的重复扭曲的序列。 |
| S\_WarpTransform | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpTransform.jpg) | 通过线性变换的组合对源素材进行扭曲，包括缩放、剪切、缩放、旋转和翻译。 |
| S\_WarpVortex | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpVortex.jpg) | 将源素材扭曲成一个漩涡，围绕一个给定的中心位置。 使用涡旋开始参数来调整涡旋的数量，并使用角度偏移来应用正常的旋转。 涡流速度可以用来自动调整涡流的数量。 |
| S\_WarpWaves | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpWaves.jpg) | 通过波浪模式对源素材进行扭曲。 你可以通过增加相位速度参数，或通过动画相位开始的值，使波浪随时间移动。 |
| S\_WarpWaves2 | Warps & Shake | ![](https://cdn.yuelili.com/2021/sapphire_WarpWaves2.jpg) | 用两组重叠的波浪图案使源素材变形。 你可以通过增加相位速度参数，或通过对相位开始参数值的动画化来使波浪随时间移动。 |
| S\_WhipLash | WhipLash | ![](https://cdn.yuelili.com/2021/sapphire_WhipLash.jpg) | 仅有2D版本的Swish3D，有可选的鞭打动作和RGB分割。 在两个输入片段之间进行切分，同时对每个片段进行二维运动。 在转换过程中，剪辑被旋转、移动和缩放参数所改变。 |
| S\_WipeBlobs | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeBlobs.jpg) | 在两个输入片段之间执行擦拭过渡，使用由噪声函数生成的斑点图案。 擦拭百分比参数应该被动画化，以控制过渡速度。 增加Grad Add参数，使Blobs图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边框宽度）参数，在擦拭过渡边缘画出一个边框。 |
| S\_WipeBubble | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeBubble.jpg) | 在两个输入片段之间进行擦拭，在过渡区域内进行气泡擦拭过程。 擦拭百分比参数应该是动画的，以控制过渡速度。 |
| S\_WipeCells | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeCells.jpg) | 在两个输入片段之间执行擦拭过渡，使用程序生成的蜂窝状图案。 擦拭百分比参数应该是动态的，以控制过渡速度。 增加Grad Add参数，在擦拭过程中使细胞图案的时间在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡的边缘画一个边界。 |
| S\_WipeChecker | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeChecker.jpg) | 在两个输入片段之间执行擦拭过渡，使用一个增长或缩小的格子。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Grad Add参数，使格子图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边框宽度）参数，在擦拭过渡边缘画一个边框。 |
| S\_WipeCircle | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeCircle.jpg) | 在两个输入片段之间使用一个增长或缩小的圆圈执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Border Width（边框宽度）参数，在擦拭过渡边缘画一个边框。 |
| S\_WipeClock | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeClock.jpg) | 在两个输入片段之间执行时钟擦拭过渡。 擦拭百分比参数应该是动态的，以控制过渡速度。 增加边框宽度参数，在擦拭过渡边缘画一个边框。 |
| S\_WipeClouds | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeClouds.jpg) | 使用移动云纹理从第一个片段过渡到第二个片段。 擦拭百分比参数应该是动态的，以控制过渡速度。 |
| S\_WipeDiffuse | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeDiffuse.jpg) | 在两个输入片段之间进行擦拭，在过渡区域内进行像素扩散处理。 擦拭百分比参数应该被动画化以控制过渡速度。 这种效果的像素化外观取决于图像的分辨率，因此建议在处理前测试你的最终分辨率。 |
| S\_WipeDots | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeDots.jpg) | 在两个输入片段之间用一个不断增长或缩小的点的网格进行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Grad Add参数，以便在擦拭过程中使小点图案的时间在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡边缘画出边界。 |
| S\_WipeDoubleWedge | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeDoubleWedge.jpg) | 使用两个楔形在两个输入片段之间执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加边界宽度参数，在擦拭过渡边缘画出边界。 |
| S\_WipeFlux | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeFlux.jpg) | 在两个输入片段之间执行擦拭过渡，使用大部分为圆形单元的通量纹理。 擦拭百分比参数应该被动画化，以控制过渡速度。 增加Grad Add参数，使助焊剂图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡的边缘画一个边界。 |
| S\_WipeFourWedges | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeFourWedges.jpg) | 在两个输入片段之间执行擦拭过渡，使用四个楔形合并成 X 形的图案。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Border Width（边框宽度）参数，在擦拭过渡的边缘画一个边框。 |
| S\_WipeLine | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeLine.jpg) | 在两个输入片段之间执行一个简单的线擦拭过渡。 擦拭百分比参数应该是动态的，以控制过渡速度。 增加边界宽度参数，在擦拭过渡的边缘画出边界。 |
| S\_WipeMoire | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeMoire.jpg) | 在两个输入片段之间进行擦拭过渡，使用组合的同心圆模式。 擦拭百分比参数应该被动画化以控制过渡速度。 相位速度和摩尔速度参数使环状物随着时间的推移自动产生动画。 增加Grad Add参数，使图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡边缘画出边界。 |
| S\_WipePixelate | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipePixelate.jpg) | 通过将一个片段的像素块以半随机的顺序添加到另一个片段上，实现两个输入片段之间的过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 调整 边缘宽度 和 块状 参数以获得不同的像素化模式。 |
| S\_WipePlasma | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipePlasma.jpg) | 使用带有移动卷须的等离子体纹理在两个输入片段之间执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Grad Add参数，使等离子体图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡边缘画出边界。 |
| S\_WipePointalize | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipePointalize.jpg) | 在两个输入片段之间的过渡，以半随机的顺序将一个片段中的刷状多边形添加到另一个片段上。 擦拭百分比参数应该是动画的，以控制过渡速度。 调整频率以改变形状的大小，并调整边缘宽度和大块参数以获得不同的模式。 |
| S\_WipeRectangle | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeRectangle.jpg) | 使用一个增长或缩小的矩形在两个输入片段之间执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Border Width（边框宽度）参数，在擦拭过渡的边缘画一个边框。 |
| S\_WipeRings | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeRings.jpg) | 在两个输入片段之间使用同心圆的模式进行擦拭过渡。 擦拭百分比参数应该是动态的，以控制过渡速度。 增加Grad Add参数，使环形图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡边缘画出边界。 |
| S\_WipeStar | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeStar.jpg) | 使用星形在两个输入片段之间执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加边界宽度参数，在擦拭过渡的边缘画出边界。 |
| S\_WipeStripes | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeStripes.jpg) | 在两个输入片段之间用一系列的条纹进行擦拭过渡。 擦拭百分比参数应该是动态的，以控制过渡速度。 增加Grad Add参数，使条纹图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边框宽度）参数，在擦拭过渡边缘画出一个边框。 |
| S\_WipeTiles | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeTiles.jpg) | 在两个输入片段之间执行擦拭过渡，使用增长或缩小的六边形、三角形、钻石或星星的图案。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Grad Add参数，使瓦片图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边框宽度）参数，在擦拭过渡的边缘画一个边框。 |
| S\_WipeWeave | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeWeave.jpg) | 在两个输入片段之间执行擦拭过渡，使用的纹理类似于垂直的编织线。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加Grad Add参数，使编织图案的时间在擦拭过程中在屏幕上移动。 增加Border Width（边界宽度）参数，在擦拭过渡边缘画出边界。 |
| S\_WipeWedge | Wipe Transitions | ![](https://cdn.yuelili.com/2021/sapphire_WipeWedge.jpg) | 使用楔形在两个输入片段之间执行擦拭过渡。 擦拭百分比参数应该是动画的，以控制过渡速度。 增加边界宽度参数，在擦拭过渡的边缘画一个边界。 |
