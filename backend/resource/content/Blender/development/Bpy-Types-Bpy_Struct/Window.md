---
title: Window
order: 15
---

# Window

## Description

[Window(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Window.html)

基类： bpy_struct

说明：Open window

### height

说明：Window height

类型：int in [0, 32767], default 0, 只读

### parent

说明：Active workspace and scene follow this window

类型：Window, 只读

### scene

说明：Active scene to be edited in the window

类型：Scene, 不会为 None

### screen

说明：Active workspace screen showing in the window

类型：Screen, 不会为 None

### stereo_3d_display

说明：Settings for stereo 3D display

类型：Stereo3dDisplay, (只读，可选)

### view_layer

说明：The active workspace view layer showing in the window

类型：ViewLayer, 不会为 None

### width

说明：Window width

类型：int in [0, 32767], default 0, 只读

### workspace

说明：Active workspace showing in the window

类型：WorkSpace, 不会为 None

### x

说明：Horizontal location of the window

类型：int in [-32768, 32767], default 0, 只读

### y

说明：Vertical location of the window

类型：int in [-32768, 32767], default 0, 只读

### cursor_warp()

全名：cursor_warp(x, y)

Set the cursor position

### cursor_set()

全名：cursor_set(cursor)

Set the cursor

参数：cursor ：[鼠标状态](https://www.yuelili.com/?p=19525) 列表里任选一个

### cursor_modal_set()

全名：cursor_modal_set(cursor)

Restore the previous cursor after calling cursor_modal_set

参数：cursor ：[鼠标状态](https://www.yuelili.com/?p=19525) 列表里任选一个

### cursor_modal_restore()

全名：cursor_modal_restore()

cursor_modal_restore

### event_simulate()

全名：event_simulate(type, value, unicode='', x=0, y=0, shift=False, ctrl=False,
alt=False, oskey=False)

event_simulate

参数：

- type：[按键类型](https://www.yuelili.com/?p=19528)之一。
- value：列表之一。['ANY', 'PRESS', 'RELEASE', 'CLICK', 'DOUBLE_CLICK', 'CLICK_DRAG', 'NORTH', 'NORTH_EAST', 'EAST', 'SOUTH_EAST', 'SOUTH', 'SOUTH_WEST', 'WEST', 'NORTH_WEST', 'NOTHING']
- shift (boolean，可选) 。 Shift
- ctrl (boolean，可选) 。 Ctrl
- alt (boolean，可选) 。 Alt
- oskey (boolean，可选) 。 OS Key

返回：Item, Added key map item

返回类型：Event

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
