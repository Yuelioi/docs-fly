---
title: AVLayer
order: 1
---

# AVLayer 对象

## AVLayer object

app.project.item(index).layer(index)

描述：AVLayer 对象为包含 AVItem 对象的那些层(合成，素材，纯色层，文本层和声音层)提供接口。

父级关系：AVLayer 是 Layer 对象的子类。除下面列出的方法外，Layer 的所有方法和属性均可用。

子级关系：AVLayer 是 TextLayer 对象的基类，因此在使用 TextLayer 对象时，可以使用 AVLayer 属性和方法。

AE 属性（ 2019 已启用光线追踪）

不同类型的图层具有不同的 AE 属性。AVLayer 具有以下属性和属性组：

- 标记 Marker
- 时间重映射 Time Remap
- 运动追踪器 Motion Trackers
- 蒙版 Masks
- 效果 Effects
- 变换 Transform
  - 锚点 Anchor Point
  - 位置 Position
  - 缩放 Scale
  - 方向 Orientation
  - X 旋转 X Rotation
  - Y 旋转 Y Rotation
  - 旋转 Rotation
  - 不透明度 Opacity
- 图层样式 Layer Styles
- 几何选项//光线追踪 3D Geometry Options // Ray-traced 3D
- 材质选项 Material Options
  - 投影 Casts Shadows
  - 透光率 Light Transmission
  - 接受阴影 Accepts Shadows
  - 接受灯光 Accepts Lights
  - 出现在反射中//光线追踪的 3D Appears in Reflections // Ray-traced 3D
  - 环境 Ambient
  - 漫射 Diffuse
  - 镜面强度 Specular Intensity
  - 镜面反光度 Specular Shininess
  - 金属质感 Metal
  - 反射强度//光线追踪 3D Reflection Intensity // Ray-traced 3D
  - 反射锐度//光线追踪 3D Reflection Sharpness // Ray-traced 3D
  - 反射衰减//光线追踪 3D Reflection Rolloff // Ray-traced 3D
  - 透明度//光线追踪 3D Transparency // Ray-traced 3D
  - 透明度下降//光线追踪 3D Transparency Rolloff // Ray-traced 3D
  - 折射率//光线追踪 3D Index of Refraction // Ray-traced 3D
- 音频 Audio
  - 音频电平 AudioLevels

示例：将活动合成的第 1 个图层，startTime 设置为 2

```javascript
var firstLayer = app.project.activeItem.layer(1);

firstLayer.startTime = 2;
```

# 属性

## adjustmentLayer 设置调整图层

app.project.item(index).layer(index).adjustmentLayer

描述：是调整图层，则为 true。

类型：布尔值，读/写。

示例：打开图层的调节开关

![](https://mir.yuelili.com/2021/07/d9037981753847964c32064b7957434d.png)

```javascript
app.project.activeItem.layer(1).adjustmentLayer = true;
```

## audioActive 判断音频是否激活

app.project.item(index).layer(index).audioActive

描述：如果音频在当前时间处于活动状态，则为 True。

活动状态：音频小喇叭打开，且当前时间处于入点和出点内

类型：布尔值，只读。

## audioEnabled 激活音频

app.project.item(index).layer(index).audioEnabled

描述：如果为 true，则启用图层的音频。对应于“时间轴”面板中的音频切换开关（小喇叭）。

类型：布尔值，读/写。

## autoOrient 图层自动定向

app.project.item(index).layer(index).autoOrient

描述：图层自动定向。

类型：AutoOrientType，枚举值。读/写：

- AutoOrientType.ALONG_PATH：沿路径定向。

- AutoOrientType.CAMERA_OR_POINT_OF_INTEREST：定位于相机。

- AutoOrientType.CHARACTERS_TOWARD_CAMERA：独立定向每个字符（需开启逐字 3D 化）。

- AutoOrientType.NO_AUTO_ORIENT：关。

## blendingMode 图层混合模式

app.project.item(index).layer(index).blendingMode

描述：图层的混合模式。

类型：BlendingMode，枚举值；读/写：

| 枚举值                           | 英文版叫法          | 中文版叫法   |
| -------------------------------- | ------------------- | ------------ |
| BlendingMode.NORMAL              | Normal              | 正常         |
| BlendingMode.DISSOLVE            | Dissolve            | 溶解         |
| BlendingMode.DANCING_DISSOLVE    | Dancing Dissolve    | 动态抖动溶解 |
| BlendingMode.DARKEN              | Darken              | 变暗         |
| BlendingMode.MULTIPLY            | Multiply            | 相乘         |
| BlendingMode.COLOR_BURN          | Color burn          | 颜色加深     |
| BlendingMode.CLASSIC_COLOR_BURN  | Classic Color Burn  | 经典颜色加深 |
| BlendingMode.LINEAR_BURN         | Linear Burn         | 线性加深     |
| BlendingMode.DARKER_COLOR        | Darker Color        | 较深的颜色   |
| BlendingMode.ADD                 | Add                 | 相加         |
| BlendingMode.LIGHTEN             | Lighten             | 变亮         |
| BlendingMode.SCREEN              | Screen              | 屏幕         |
| BlendingMode.COLOR_DODGE         | Color Dodge         | 颜色减淡     |
| BlendingMode.CLASSIC_COLOR_DODGE | Classic Color Dodge | 经典颜色减淡 |
| BlendingMode.LINEAR_DODGE        | Linear Dodge        | 线性减淡     |
| BlendingMode.LIGHTER_COLOR       | Lighter Color       | 较浅的颜色   |
| BlendingMode.OVERLAY             | Overlay             | 叠加         |
| BlendingMode.SILHOUETTE_LUMA     | Soft Light          | 柔光         |
| BlendingMode.HARD_LIGHT          | Hard Light          | 强光         |
| BlendingMode.LINEAR_LIGHT        | Linear Light        | 线性光       |
| BlendingMode.VIVID_LIGHT         | Vivid Light         | 亮光         |
| BlendingMode.PIN_LIGHT           | Pin Light           | 点光         |
| BlendingMode.HARD_MIX            | Hard Mix            | 纯色混合     |
| BlendingMode.DIFFERENCE          | Difference          | 差值         |
| BlendingMode.CLASSIC_DIFFERENCE  | Classic Difference  | 经典差值     |
| BlendingMode.EXCLUSION           | Exclusion           | 排除         |
| BlendingMode.STENCIL_LUMA        | Subtract            | 相减         |
| BlendingMode.DIVIDE              | Divide              | 相除         |
| BlendingMode.HUE                 | Hue                 | 色相         |
| BlendingMode.SATURATION          | Saturation          | 饱和度       |
| BlendingMode.COLOR               | Color               | 颜色         |
| BlendingMode.LUMINOSITY          | Luminosity          | 发光度       |
| BlendingMode.SOFT_LIGHT          | Stencil Alpha       | 模板 Apha    |
| BlendingMode.STENCIL_ALPHA       | Stencil Luma        | 模板亮度     |
| BlendingMode.SUBTRACT            | Silhouette Alpha    | 轮廓 Apha    |
| BlendingMode.SILHOUETE_ALPHA     | Silhouette luma     | 轮廓亮度     |
| BlendingMode.ALPHA_ADD           | Alpha Add           | Alpha 添加   |
| BlendingMode.LUMINESCENT_PREMUL  | Luminescent premul  | 冷光预乘     |

```javascript
app.project.activeItem.layer(1).blendingMode = BlendingMode.ADD;

// 将活动合成的第 1 个图层，混合模式改为 Add
```

## canSetCollapseTransformation 可否塌陷

app.project.item(index).layer(index).canSetCollapseTransformation

描述：如果有塌陷属性（或者叫折叠、连续栅格化、collapseTransformation），则为 True 。比如 AI 图层、合成都可以开塌陷

类型：布尔值，只读。

```javascript
alert(app.project.activeItem.layer(1).canSetCollapseTransformation);
```

## canSetTimeRemapEnabled 可否时间重映射

app.project.item(index).layer(index).canSetTimeRemapEnabled

描述：如果可以更改时间重映射（timeRemapEnabled），则为 True 。

类型：布尔值，只读。

## collapseTransformation 是否塌陷

app.project.item(index).layer(index).collapseTransformation

描述：如果开启了塌陷，则为 true。

类型：布尔值，读/写。

## effectsActive 效果是否活动

app.project.item(index).layer(index).effectsActive

描述：如果该图层的效果处于活动状态，则为 True。也就是打开效果开关。

类型：布尔值，读/写。

## environmentLayer 是否环境图层

app.project.item(index).layer(index).environmentLayer

描述：如果这是射线追踪 3D 合成中的环境图层，则为 true。将此属性设置为 true 会自动变成 3D 图层(.threeDLayer 变为 true)。

类型：布尔值，读/写。

## frameBlending 是否帧混合

app.project.item(index).layer(index).frameBlending

描述：如果图层启用了帧混合，则为 True。

类型：布尔值，只读。

## frameBlendingType 帧混合类型

app.project.item(index).layer(index).frameBlendingType

描述：图层帧混合类型。

类型：FrameBlendingType 枚举值；读/写：

- FrameBlendingType.FRAME_MIX （帧混合）

- FrameBlendingType.NO_FRAME_BLEND（不混合）

- FrameBlendingType.PIXEL_MOTION（像素混合）

## guideLayer 引导层设置

app.project.item(index).layer(index).guideLayer

描述：如果该图层是引导层，则为 true。

类型：布尔值，读/写。

## hasAudio 是否包含音频

app.project.item(index).layer(index).hasAudio

描述：如果图层包含音频，则为 true。

类型：布尔值，只读。

## hasTrackMatte 是否开启轨道遮罩

app.project.item(index).layer(index).hasTrackMatte

描述：如果该图层开启轨道遮罩，则为 true。（trackMatteType 控制遮罩类型）

类型：布尔值，只读。

## height 图层高度

app.project.item(index).layer(index).height

描述：图层高度(像素)。

类型：浮点，只读。

## isNameFromSource 是否基于图层源命名

app.project.item(index).layer(index).isNameFromSource

描述：未单独设置名称，但包含源命名，则为 True。此时 layer.name = layer.source.name。否则为 False。

类型：布尔值，只读。

示例：新建一个纯色图层，会根据颜色自动命名，如：灰色 纯色 1

## isTrackMatte 是否轨道遮罩

app.project.item(index)layer(index).isTrackMatte

描述：如果该图层下面图层开启轨道遮罩，则为 true。

类型：布尔值 只读。

## motionBlur 运动模糊设置

app.project.item(index).layer(index).motionBlur

描述：如果图层启用运动模糊，则为 True。

类型：布尔值，读/写。

## preserveTransparency 保留透明度设置

app.project.item(index).layer(index).preserveTransparency

![](https://mir.yuelili.com/2021/07/2efcf8f573e4423ff5f05ae0ed7efad4.png)

描述：如果为图层启用了保留透明度，则为 True。

类型：布尔值，读/写。

## quality 图层质量与采样

app.project.item(index).layer(index).quality

描述：图层的质量与采样。右键图层 - 品质

类型：LayerQuality，枚举值; 读/写：

- LayerQuality.BEST

- LayerQuality.DRAFT

- LayerQuality.WIREFRAME

## samplingQuality 图层采样方法

app.project.item(index).layer(index).samplingQuality

::: info 提示
此功能于 After Effects 12.0(CC)添加
:::

描述：设置/获取图层采样方法(双三次或双线性)。右键图层 - 品质

类型：LayerSamplingQuality，枚举值; 读/写：

- LayerSamplingQuality.BICUBIC

- LayerSamplingQuality.BILINEAR

## source 图层源

app.project.item(index).layer(index).source

描述：此 AVItem 层的源。如果是 Text 层则为 null。使用 AVLayer.replaceSource()更改数值。

类型：AVItem 对象；只读。

## threeDLayer 3D 图层设置

app.project.item(index).layer(index).threeDLayer

描述：如果是 3D 图层，则为 true。

类型：布尔值，读/写。

## threeDPerChar 逐字 3D 化设置

app.project.item(index).layer(index).threeDPerChar

描述：如果开启“动画 - 逐字 3D 化”，则为 true。仅适用于文本层。

类型：布尔值，读/写。

## timeRemapEnabled 时间重新映射设置

app.project.item(index).layer(index).timeRemapEnabled

描述：如果开启时间重新映射，则为 True。

类型：布尔值，读/写。

## trackMatteType 轨迹遮罩设置

app.project.item(index).layer(index).trackMatteType

描述：如果此图层具有轨迹遮罩，指轨迹遮罩方式。

类型：TrackMatteType，枚举值; 读/写：

- TrackMatteType.ALPHA

- TrackMatteType.ALPHA_INVERTED

- TrackMatteType.LUMA

- TrackMatteType.LUMA_INVERTED

- TrackMatteType.NO_TRACK_MATTE

## width 图层宽度

app.project.item(index).layer(index).width

描述：图层宽度(像素)。

类型：浮点; 只读。

##

# 方法

## addToMotionGraphicsTemplate() 图层属性添加到动态图形

app.project.item(index).layer(index).addToMotionGraphicsTemplate(comp)

::: info 提示
:::

此功能已添加到 After Effects 18.0(2021)中

描述：将图层属性添加到指定合成的“基本图形面板”（EGP）。

如果成功添加图层，则返回 true，否则为 false。

如果无法添加，AE 将弹出警告对话框。可能已经添加过了，或者该图层不是可以替换其媒体的图层类型(称为“媒体替换图层”)。

使用 AVLayer.canAddToMotionGraphicsTemplate()测试是否可以将图层添加到运动图形模板。

参数：

- comp，合成对象；添加到 EGP 的合成，必需。

返回：布尔值。

案例：给当前合成当前选择的图层位置属性，添加到基本图形模板

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
var pos = selLayer.position; // 它的位置
pos.addToMotionGraphicsTemplate(myComp); //添加属性到 EGP;
```

## addToMotionGraphicsTemplateAs() 将图层添加到动态图形

app.project.item(index).layer(index).addToMotionGraphicsTemplateAs(comp, name)

::: info 提示
此功能已添加到 After Effects 18.0(2021)中
:::

描述：将图层添加到“基本图形面板”。

如果成功添加图层，则返回 true，否则为 false。

如果无法添加，AE 将弹出警告对话框。可能已经添加过了，或者该图层不是可以替换其媒体的图层类型(称为“媒体替换图层”)。

使用 AVLayer.canAddToMotionGraphicsTemplate()测试是否可以将图层添加到运动图形模板。

参数：

- comp：合成对象；添加到 EGP 的合成，必需。

- name：自定义属性名称。

返回：布尔值。

案例：给当前合成当前选择的图层位置属性，添加到基本图形模板，并且命名该属性

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
var pos = selLayer.position; // 它的位置
pos.addToMotionGraphicsTemplate(myComp); //添加属性到 EGP;
```

## audioActiveAtTime() 音频是否活动

app.project.item(index).layer(index).audioActiveAtTime(time)

描述：音频在指定的时间处于活动状态，则返回 true。

活动状态：音频小喇叭打开，且当前时间处于入点和出点内

参数：

- time 时间(以秒为单位)。浮点值。

返回：布尔值。

## calculateTransformFromPoints() 基于坐标生成 3D 图形

app.project.item(index).layer(index).calculateTransformFromPoints(pointTopLeft, pointTopRight, pointBottomRight)

描述：根据此层中的一组点计算转换。

参数：

- pointTopLeft：左上角坐标，[x，y，z]。

- pointTopRight：右上角坐标，[x，y，z]。

- pointBottomRight：右下角坐标，[x，y，z]。

返回：图层对象。

示例：基于 3 个空间点坐标，生成对应图层。会自动更改缩放、旋转

![](https://cdn.yuelili.com/20210914014358.png)![](https://cdn.yuelili.com/20210914014457.png)

```javascript
var newLayer = app.project.item(1).layer(1); //随便选择一个图层
newLayer.threeDLayer = true; //改为3D图形
var tl = [0, 0, 0];
var tr = [500, 0, 100];
var bl = [200, 300, 300];
var transform = newLayer.calculateTransformFromPoints(tl, tr, bl);
for (var sel in transform) {
  newLayer.transform[sel].setValue(transform[sel]);
}
```

## canAddToMotionGraphicsTemplate() 图层能否添加到基本图形

app.project.item(index).layer(index).canAddToMotionGraphicsTemplate(comp)

::: info 提示
此功能已添加到 After Effects 18.0(2021)中
:::

描述：测试是否可以将图层的属性添加到指定合成的“基本图形面板”中。如果成功添加图层，则返回 true，否则为 false。

如果无法添加，AE 将弹出警告对话框。可能已经添加过了，或者该图层不是可以替换其媒体的图层类型（媒体替换图层）。

媒体替换层被识别为 AVLayers，并且将 AVLayer.source 设置为 FootageItem 对象(具有特定的源类型)或 CompItem 对象。

为了将其视为媒体替换层，AVLayer 需要遵守以下限制：

- Layer.hasVideo 应该返回 true。

- AVLayer.adjustmentLayer 应该返回 false。

- Layer.nullLayer 应该返回 false。

- 如果 AVLayer.source 是 footageItem 对象，则 FootageItem.FootageSource 不应是 SolidSource 对象。

- 如果 AVLayer.source 是 footageItem 对象，而 FootageItem.FootageSource 是 FileSource 对象，则该 FileSource 不应指向非媒体文件，例如 JSX 脚本文件。

说人话：必须有视频（ 可视），不能是调整图层或空对象、不能是纯色图层...Source 没学 晚点校正

参数：

- comp ，合成对象；要将属性添加到 EGP 的合成。必需。

返回：布尔值。

示例：判断当前合成下，选择的第 1 个图层的位置属性，能否添加到合成基本图形

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
var pos = selLayer.position; // 它的位置
alert(pos.canAddToMotionGraphicsTemplate(myComp));
```

## compPointToSource() 合成坐标转图层坐标

app.project.item(index).layer(index).compPointToSource(pos)

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：将合成坐标(例如 sourcePointToComp)转换为图层坐标。详见表达式[空间转换](https://www.yuelili.com/docs/ae-expression/exp-layer/#48-toc-title)

:::danger
对于文本层，该值仅反映第一个字符。
:::

参数：

- pos， ([X，Y])格式的图层位置坐标数组。

返回：([X，Y])位置坐标的数组；只读。

示例：将[100,100]转为指定图层的图层坐标

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
alert(selLayer.compPointToSource([100, 100]));
```

## openInViewer() 打开图层查看器

app.project.item(index).layer(index).openInViewer()

描述：在“图层”面板中打开该图层，然后将“图层”面板移到最前面并为其赋予焦点。

![](https://cdn.yuelili.com/20210914015944.png)

参数：无。

返回：“图层”面板的 Viewer 对象，如果无法打开该图层，则为 null(无法在“图层”面板中打开文本或形状图层)。

## replaceSource() 替换图层源

app.project.item(index).layer(index).replaceSource(newSource, fixExpressions)

描述：替换此层的源。

参数：

- newSource：新的源 AVItem 对象。

- fixExpressions：true，调整新来源的表达式，否则为 false。

注意：此功能可能会占用大量资源；如果要替换大量素材，请仅在操作结束时执行此操作。另请参阅 Project.autoFixExpressions()。

返回：无。

::: danger
如果在空图层上执行此方法，则 layersisNull 属性不会从更改 true。这将导致该图层在 comp 查看器和渲染中不可见。
:::

## sourcePointToComp() 图层坐标转合成坐标

app.project.item(index).layer(index).sourcePointToComp(pos)

::: info 提示
此功能已添加到 After Effects 13.2(CC 2014.2)中
:::

描述：将图层坐标转换为合成坐标。详见表达式[空间转换](https://www.yuelili.com/docs/ae-expression/exp-layer/#48-toc-title)

::: danger
对于文本层，该值仅反映第一个字符。
:::

参数：

- pos：([X，Y])格式的图层位置坐标数组。

返回：([X，Y])位置坐标的数组；只读。

示例：把[100,100]转为指定合成坐标

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
alert(selLayer.sourcePointToComp([100, 100]));
```

## sourceRectAtTime() 图形边界宽高

app.project.item(index).layer(index).sourceRectAtTime(timeT, extents)

描述：检索指定时间，图层的矩形边界，文本或形状图层结果不同。详见表达式[sourceRectAtTime()](https://www.yuelili.com/?p=10443)

参数：

- timeT：时间索引，以秒为单位。浮点值。

- extents：true：包括范围。false，不包括范围。范围参数仅适用于形状图层，size 可增加图层的边界大小。

返回：四个属性，JavaScript 对象[ top，left，width，height]。

示例：返回选择图层的上边界

```javascript
var myComp = app.project.activeItem; //myComp = 当前合成
selLayer = myComp.selectedLayers[0]; //选择图层集里的第 1 个
alert(selLayer.sourceRectAtTime(selLayer.time, false).top);
```
