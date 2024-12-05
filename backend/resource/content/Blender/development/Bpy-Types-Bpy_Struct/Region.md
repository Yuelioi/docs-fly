---
title: Region
order: 10
---

# Region

## Description

[Region(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Region.html)

基类： bpy_struct

说明：Region in a subdivided screen area

### alignment

说明：在区域(area)中区块(Region)的对齐方式

- NONE：不要使用任何固定的对齐方式，填补可用空间。
- TOP：顶部对齐
- BOTTOM：底部对齐
- LEFT：左对齐
- RIGHT：右对齐
- HORIZONTAL_SPLIT：水平分割
- VERTICAL_SPLIT：垂直分割
- FLOAT：区域在屏幕上浮动，不使用任何固定对齐方式。
- QUAD_SPLIT：区域被水平和垂直分割。

类型：8 选 1，默认 NONE，只读

### data

说明：区块特定数据（类型取决于区域类型）。

类型：AnyType, 只读

### height

说明：区块高度

类型：int in [0, 32767], default 0, 只读

### type

说明：区块类型

类型：默认 WINDOW，只读

- WINDOW
- HEADER
- CHANNELS：文件选择窗口
- TEMPORARY’
- UI：N 面板
- TOOLS：
- TOOL_PROPS
- PREVIEW
- HUD
- NAVIGATION_BAR
- EXECUTE
- FOOTER
- TOOL_HEADER
- XR

### view2d

说明：2D view of the region

类型：View2D, (只读，可选)

### width

说明：区块宽度

类型：int in [0, 32767], default 0, 只读

### x

说明：该区块的窗口相对垂直位置

类型：int in [-inf, inf], default 0, 只读

### y

说明：该区块的窗口相对水平位置

类型：int in [-inf, inf], default 0, 只读

### tag_redraw()

全名：tag_redraw()

tag_redraw

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id(字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
