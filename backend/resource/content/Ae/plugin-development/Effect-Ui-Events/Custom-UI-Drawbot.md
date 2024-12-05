---
title: Custom-UI-Drawbot
order: 1
---

# 自定义UI与Drawbot

自定义用户界面使用 Drawbot 的合成绘图模型。Drawbot 套件可用于。

1. 基本的二维路径绘制。线条、矩形、弧线、贝塞尔曲线
2. 描边/填充/着色路径
3. 图像绘制。将 ARGB/BGRA 缓冲器合成到表面上
4. 推送/弹出表面状态
5. 文本绘制，如果供应商支持的话(客户应该在实际绘制之前首先检查是否支持文本绘制)。

绘图只能在`PF_Event_DRAW`期间发生(而不是在`PF_Event_DRAG`或`PF_Event_DO_CLICK`期间)。

要使用 Drawbot，首先通过传递 PF_Context 到一个新的套件调用`PF_GetDrawingReference`获得绘图参考。

如果返回一个非 NULL 的绘图参考，使用它从[DRAWBOT_DrawbotSuite]获得供应商和表面参考。

Drawbot 套件包括`DRAWBOT_DrawbotSuite`, `DRAWBOT_SupplierSuite`, `DRAWBOT_SurfaceSuite`, `DRAWBOT_PathSuite`。

## 让您的自定义UI看起来不是那么“自定义”

使用新的`PF_EffectCustomUIOverlayThemeSuite`来匹配主机应用程序的用户界面。你的用户会感谢你的。

## 重绘

为了重新绘制一个窗口(pane)的特定区域，我们推荐使用以下方法。

1. 从效果中调用`PF_InvalidateRect` ( 参见[PF_AppSuite](../effect-details/useful-utility-functions.html))。这将会惰性显示重绘，并将在下一个可用的空闲时刻更新。矩形的坐标与相关窗口有关。使用一个 NULL 矩形将更新整个窗口。
2. 将[事件输出标志](PF_EventExtra.html)设置为`PF_EO_UPDATE_NOW`，这将在当前事件返回时为指定的窗口时,进行即时绘制事件。

如果一个效果需要一次更新多个窗口，应该设置`PF_OutFlag_REFRESH_UI` ( 参见[PF_OutFlags](./effect-basics/PF_OutData.html))，这将导致整个 ECW、合成和图层窗口的重绘。

## HiDPI和Retina显示支持

为了支持 HiDPI 和 Retina 显示器，可以使用两倍大小的屏幕外图像，然后使用`Drawbot_SurfaceSuite`的`Transform`函数，在绘制图像前将其缩小一半。

## PF_EffectCustomUISuite

启用一个效果来获得绘图参考。这是使用 Drawbot 需要的第一个调用。

### PF_EffectCustomUISuite1

#### PF_GetDrawingReference

获取绘图参考。

```cpp
PF_GetDrawingReference(
 const PF_ContextH effect_contextH,
 DRAWBOT_DrawRef *referenceP0);
```

## Drawbot_DrawbotSuite

使用 Drawbot 的参考，得到供应商和表面的参考。

### Drawbot_DrawbotSuite1

#### GetSupplier

获取供应商参考。
需要使用`Drawbot_SupplierSuite`。

```cpp
GetSupplier(
 DRAWBOT_DrawRef in_drawbot_ref,
 DRAWBOT_SupplierRef *out_supplierP);
```

#### GetSurface

获取外表参考。
需要使用`Drawbot_SurfaceSuite`。

```cpp
GetSurface(
 DRAWBOT_DrawRef in_drawbot_ref,
 DRAWBOT_SurfaceRef *out_surfaceP);
```

## Drawbot_SupplierSuite

调用创建和释放绘图工具，获得默认设置，并查询绘图能力。

### Drawbot_SupplierSuite1

#### NewPen

创建一个新的笔。使用`Drawbot_SupplierSuite`的`ReleaseObject`释放它。

```cpp
NewPen(
 DRAWBOT_SupplierRef in_supplier_ref,
 const DRAWBOT_ColorRGBA *in_colorP,
 float in_size,
 DRAWBOT_PenRef *out_penP);
```

#### NewBrush

创建一个新的画笔。使用`Drawbot_SupplierSuite`中的`ReleaseObject`释放这个笔刷。

```cpp
NewBrush(
 DRAWBOT_SupplierRef in_supplier_ref,
 const DRAWBOT_ColorRGBA *in_colorP,
 DRAWBOT_BrushRef *out_brushP);
```

#### SupportsText

检查当前供应商是否支持文本。

```cpp
SupportsText(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_Boolean *out_supports_textB);
```

#### GetDefaultFontSize

获取默认的字体大小。

```cpp
GetDefaultFontSize(
 DRAWBOT_SupplierRef in_supplier_ref,
 float *out_font_sizeF);
```

#### NewDefaultFont

用默认设置创建一个新的字体。
你可以从`GetDefaultFontSize`传递默认的字体大小。
使用`Drawbot_SupplierSuite`中的`ReleaseObject`来释放。

```cpp
NewDefaultFont(
 DRAWBOT_SupplierRef in_supplier_ref,
 float in_font_sizeF,
 DRAWBOT_FontRef *out_fontP);
```

#### NewImageFromBuffer

从传递给 in_dataP 的缓冲区创建一个新的图像。
使用[Drawbot_SupplierSuite]中的`ReleaseObject`释放这个图像。

```cpp
NewImageFromBuffer(
 DRAWBOT_SupplierRef in_supplier_ref,
 int in_width,
 int in_height,
 int in_row_bytes,
 DRAWBOT_PixelLayout in_pl,
 const void *in_dataP,
 DRAWBOT_ImageRef *out_imageP);
```

`DRAWBOT_PixelLayout`可以是以下的一种。

- `kDRAWBOT_PixelLayout_24RGB`,
- `kDRAWBOT_PixelLayout_24BGR`,
- `kDRAWBOT_PixelLayout_32RGB`,
- `ARGB` (A is ignored),
- `kDRAWBOT_PixelLayout_32BGR`,
- `BGRA` (A is ignored),
- `kDRAWBOT_PixelLayout_32ARGB_Straight`,
- `kDRAWBOT_PixelLayout_32ARGB_Premul`,
- `kDRAWBOT_PixelLayout_32BGRA_Straight`,
- `kDRAWBOT_PixelLayout_32BGRA_Premul`

#### NewPath

创建一个新的路径。使用`Drawbot_SupplierSuite`中的`ReleaseObject`将其释放。

```cpp
NewPath(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_PathRef *out_pathP);
```

#### SupportsPixelLayoutBGRA

一个给定的 Drawbot 实现可以支持多个通道命令，但很可能会偏向于其中一个。
使用以下四个回调，为任何需要 "DRAWBOT_PixelLayout "的 API(例如："NewImageFromBuffer")获得首选通道顺序。

```cpp
SupportsPixelLayoutBGRA(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_Boolean *out_supports_bgraPB);
```

#### PrefersPixelLayoutBGRA

```cpp
PrefersPixelLayoutBGRA(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_Boolean *out_prefers_bgraPB);
```

#### SupportsPixelLayoutARGB

```cpp
SupportsPixelLayoutARGB(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_Boolean *out_supports_argbPB);
```

#### PrefersPixelLayoutARGB

```cpp
PrefersPixelLayoutARGB(
 DRAWBOT_SupplierRef in_supplier_ref,
 DRAWBOT_Boolean *out_prefers_argbPB);
```

#### RetainObject

保留(增加引用次数)任何对象(笔、刷、路径等)。例如，当任何对象被复制时，应该使用它，而且复制的对象应该被保留。

```cpp
RetainObject(
 DRAWBOT_ObjectRef in_obj_ref);
```

#### ReleaseObject

释放(减少引用次数)任何对象(笔、笔刷、路径等)。这个函数必须被调用，用于任何使用`NewXYZ()`从这个套件创建的对象。
不要对`DRAWBOT_SupplierRef`和`DRAWBOT_SupplierRef`调用此函数，因为这些不是由插件创建的。

```cpp
ReleaseObject(
 DRAWBOT_ObjectRef in_obj_ref);
```

## Drawbot_SurfaceSuite

在曲面上绘图，以及查询和设置绘图设置。

### Drawbot_SurfaceSuite1

#### PushStateStack

把当前的曲面状态推到堆栈中。它应该被弹出以检索旧的状态。
如果你要对曲面进行剪裁或转换，或改变插值或抗锯齿的策略，就需要恢复状态。

```cpp
PushStateStack(
 DRAWBOT_SurfaceRef in_surface_ref);
```

#### PopStateStack

把上次推送的曲面状态从堆栈中弹出。

```cpp
PopStateStack(
 DRAWBOT_SurfaceRef in_surface_ref);
```

#### PaintRect

在曲面上画一个颜色矩形。

```cpp
PaintRect(
 DRAWBOT_SurfaceRef in_surface_ref,
 const DRAWBOT_ColorRGBA *in_colorP,
 const DRAWBOT_RectF32 *in_rectPR);
```

#### FillPath

用画笔和填充类型填充一个路径。

```cpp
FillPath(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_BrushRef in_brush_ref,
 DRAWBOT_PathRef in_path_ref,
 DRAWBOT_FillType in_fill_type);
```

`DRAWBOT_FillType`是以下其中一个。

- `kDRAWBOT_FillType_EvenOdd`,
- `kDRAWBOT_FillType_Winding`

#### StrokePath

用笔画出路径。

```cpp
StrokePath(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_PenRef in_pen_ref,
 DRAWBOT_PathRef in_path_ref);
```

#### Clip

夹住表面。

```cpp
Clip(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_SupplierRef in_supplier_ref,
 const DRAWBOT_Rect32 *in_rectPR);
```

#### GetClipBounds

获取夹子的边界。

```cpp
GetClipBounds(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_Rect32 *out_rectPR);
```

#### IsWithinClipBounds

检查矩形是否在剪裁的范围内。

```cpp
IsWithinClipBounds(
 DRAWBOT_SurfaceRef in_surface_ref,
 const DRAWBOT_Rect32 *in_rectPR,
 DRAWBOT_Boolean *out_withinPB);
```

#### Transform

转换最后的曲面状态。

```cpp
Transform(
 DRAWBOT_SurfaceRef in_surface_ref,
 const DRAWBOT_MatrixF32 *in_matrixP);
```

#### DrawString

画一个字符串。

```cpp
DrawString(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_BrushRef in_brush_ref,
 DRAWBOT_FontRef in_font_ref,
 const DRAWBOT_UTF16Char *in_stringP,
 const DRAWBOT_PointF32 *in_originP,
 DRAWBOT_TextAlignment in_alignment_style,
 DRAWBOT_TextTruncation in_truncation_style,
 float in_truncation_width);
```

`DRAWBOT_TextAlignment`是以下其中之一。

- `kDRAWBOT_TextAlignment_Left`,
- `kDRAWBOT_TextAlignment_Center`,
- `kDRAWBOT_TextAlignment_Right`

`DRAWBOT_TextTruncation`是下列其中之一。

- `kDRAWBOT_TextTruncation_None`,
- `kDRAWBOT_TextTruncation_End`,
- `kDRAWBOT_TextTruncation_EndEllipsis`,
- `kDRAWBOT_TextTruncation_PathEllipsis`

#### DrawImage

在曲面上绘制一个用`NewImageFromBuffer()`创建的图像。Alpha = [0.0f, 1.0f ].

```cpp
DrawImage(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_ImageRef in_image_ref,
 const DRAWBOT_PointF32 *in_originP,
 float in_alpha);
```

#### SetInterpolationPolicy

```cpp
SetInterpolationPolicy(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_InterpolationPolicy in_interp);
```

`DRAWBOT_InterpolationPolicy`是以下之一。

- `kDRAWBOT_InterpolationPolicy_None`,
- `kDRAWBOT_InterpolationPolicy_Med`,
- `kDRAWBOT_InterpolationPolicy_High`

#### GetInterpolationPolicy

```cpp
GetInterpolationPolicy(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_InterpolationPolicy *out_interpP);
```

#### SetAntiAliasPolicy

```cpp
SetAntiAliasPolicy(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_AntiAliasPolicy in_policy);
```

`DRAWBOT_AntiAliasPolicy`是以下其中一个。

- `kDRAWBOT_AntiAliasPolicy_None`,
- `kDRAWBOT_AntiAliasPolicy_Med`,
- `kDRAWBOT_AntiAliasPolicy_High`

#### GetAntiAliasPolicy

```cpp
GetAntiAliasPolicy(
 DRAWBOT_SurfaceRef in_surface_ref,
 DRAWBOT_AntiAliasPolicy *out_policyP);
```

#### Flush

Flush 绘图。这并不总是需要的，如果过度使用，可能会导致过度重绘和闪烁。

```cpp
Flush(
 DRAWBOT_SurfaceRef in_surface_ref);
```

## Drawbot_PathSuite

调用绘制路径。

### Drawbot_PathSuite1

#### MoveTo

移动到一个点。

```cpp
MoveTo(
 DRAWBOT_PathRef in_path_ref,
 float in_x,
 float in_y);
```

#### LineTo

在路径上添加一条线。

```cpp
LineTo(
 DRAWBOT_PathRef in_path_ref,
 float in_x,
 float in_y);
```

#### BezierTo

在路径上添加一个立方体贝塞尔。

```cpp
BezierTo(
 DRAWBOT_PathRef in_path_ref,
 const DRAWBOT_PointF32 *in_pt1P,
 const DRAWBOT_PointF32 *in_pt2P,
 const DRAWBOT_PointF32 *in_pt3P);
```

#### AddRect

在路径上添加一个矩形。

```cpp
AddRect(
 DRAWBOT_PathRef in_path_ref,
 const DRAWBOT_RectF32 *in_rectPR);
```

#### AddArc

在路径上添加一个圆弧。零起点度数==3 点钟方向。
顺时针扫过。角度的单位是度。

```cpp
AddArc(
 DRAWBOT_PathRef in_path_ref,
 const DRAWBOT_PointF32 *in_centerP,
 float in_radius,
 float in_start_angle,
 float in_sweep);
```

#### Close

关闭路径。

```cpp
Close(
 DRAWBOT_PathRef in_path_ref);
```

## PF_EffectCustomUIOverlayThemeSuite

这个套件应该用于在合成和图层窗口中对路径和顶点进行描画和填充。After Effects 内部正在使用这个套件，我们将其提供给大家，以使自定义的用户界面在不同的效果中看起来一致。前景/阴影的颜色是根据应用程序的亮度水平来计算的，因此，无论应用程序在偏好设置中的亮度如何，自定义用户界面总是可见的。

### PF_EffectCustomUIOverlayThemeSuite1

#### PF_GetPreferredForegroundColor

获取首选的前景颜色。

```cpp
PF_GetPreferredForegroundColor(
 DRAWBOT_ColorRGBA *foreground_colorP);
```

#### PF_GetPreferredShadowColor

获取首选的阴影颜色。

```cpp
PF_GetPreferredShadowColor(
 DRAWBOT_ColorRGBA *shadow_colorP);
```

#### PF_GetPreferredStrokeWidth

获取首选的前景和阴影笔刷宽度。

```cpp
PF_GetPreferredStrokeWidth(
 float *stroke_widthPF);
```

#### PF_GetPreferredVertexSize

获取首选的顶点尺寸。

```cpp
PF_GetPreferredVertexSize(
 float *vertex_sizePF);
```

#### PF_GetPreferredShadowOffset

获得首选的阴影偏移。

```cpp
PF_GetPreferredShadowOffset(
 A_LPoint *shadow_offsetP);
```

#### PF_StrokePath

用叠加主题的前景色描画路径。
可以选择使用叠加主题的阴影颜色绘制阴影。
使用叠加主题的笔刷宽度来绘制前景和阴影。

```cpp
PF_StrokePath(
 const DRAWBOT_DrawRef drawbot_ref,
 const DRAWBOT_PathRef path_ref
 PF_Boolean draw_shadowB);
```

#### PF_FillPath

用叠加主题的前景色填充路径。
可以选择使用叠加主题的阴影颜色来绘制阴影。

```cpp
PF_FillPath(
 const DRAWBOT_DrawRef drawbot_ref,
 const DRAWBOT_PathRef path_ref
 PF_Boolean draw_shadowB);
```

#### PF_FillVertex

使用覆盖主题的前景颜色和顶点大小，在中心点周围填充一个方形顶点。

```cpp
PF_FillVertex(
 const DRAWBOT_DrawRef drawbot_ref,
 const A_FloatPoint *center_pointP
 PF_Boolean draw_shadowB);
```
