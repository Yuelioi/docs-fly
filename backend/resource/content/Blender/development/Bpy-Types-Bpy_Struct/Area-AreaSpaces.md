---
title: Area + AreaSpaces
order: 2
---

# Area + AreaSpaces

## Description

[Area(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Area.html)

[AreaSpaces(bpy_struct)](https://docs.blender.org/api/master/bpy.types.AreaSpaces.html)

## Area(bpy_struct)

分割屏幕的区域，包括编辑器在内

### height

说明：区域高度

类型：int in [0, 32767], default 0, 只读

### regions

说明：被分割在的区块

类型：bpy_prop_collection of Region, 只读

### show_menus

说明：在抬头(Header)显示标题

类型：boolean, default False

### spaces

说明：该区域包含的空间，第一个是活动空间（例如，在某一区域恢复以前使用的三维视图空间，以获得旧的视图方向）。

类型：AreaSpaces bpy_prop_collection of Space, 只读

### type

说明：当前区域的编辑类型

类型

- EMPTY：无。
- VIEW_3D：3D 视窗，在 3D 环境中操纵物体。
- IMAGE_EDITOR：UV/图像编辑器，查看和编辑图像和 UV 图。
- NODE_EDITOR：Node Editor，基于节点的着色和合成工具的编辑器。
- SEQUENCE_EDITOR：Video Sequencer，视频编辑工具。
- CLIP_EDITOR：Movie Clip Editor，运动跟踪工具。
- DOPESHEET_EDITOR：Dope Sheet，调整关键帧的时间。
- GRAPH_EDITOR：Graph Editor，编辑驱动程序和关键帧插值。
- NLA_EDITOR：Nonlinear Animation，组合和分层动作。
- TEXT_EDITOR：文本编辑器，编辑脚本和文件中的文档。
- CONSOLE：Python Console，用于高级编辑和脚本开发的交互式编程控制台。
- INFO：信息，操作日志，警告和错误信息。
- TOPBAR：顶部栏，屏幕顶部的全局栏，用于每个窗口的全局设置。
- STATUSBAR：状态栏，屏幕底部的全局栏，用于显示一般状态信息。
- OUTLINER：Outliner，场景图和所有可用数据块的概述。
- PROPERTIES：属性，编辑活动对象和相关数据块的属性。
- FILE_BROWSER：文件浏览器，浏览文件和资产。
- SPREADSHEET：Spreadsheet，浏览表格中的几何体数据。
- PREFERENCES：Preferences，编辑持久的配置设置。

类型：枚举值，上面的其中之一。默认为 ‘EMPTY’

### ui_type

说明：当前区域的编辑类型

类型：enum in [], 默认为""

### width

说明：区域宽度

类型：int in [0, 32767], default 0, 只读

### x

说明：该区域的窗口相对垂直位置

类型：int in [-inf, inf], default 0, 只读

### y

说明：该区域的窗口相对水平位置

类型：int in [-inf, inf], default 0, 只读

### tag_redraw()

全名：tag_redraw()

说明：tag_redraw

### header_text_set()

全名：header_text_set(text)

说明：Set the header status text

参数：text，字符串。Header 的新文字，为空则清空

### classmethodbl_rna_get_subclass()

全名：classmethodbl_rna_get_subclass(id, default=None)

参数：id，字符串。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethodbl_rna_get_subclass_py()

全名：classmethodbl_rna_get_subclass_py(id, default=None)

参数：id (string) – The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## AreaSpaces(bpy_struct)

空间集合

### active

说明：Space currently being displayed in this area

类型：Space, 只读

### classmethod bl_rna_get_subclass()

全名：classmethodbl_rna_get_subclass(id, default=None)

参数：id(字符串)– The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethodbl_rna_get_subclass_py(id, default=None)

参数：id (字符串) – The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
