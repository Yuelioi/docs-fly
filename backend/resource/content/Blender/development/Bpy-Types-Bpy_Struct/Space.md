---
title: Space
order: 11
---

# Space

## Description

Space(bpy_struct)

基类： bpy_struct

子类：SpaceClipEditor, SpaceConsole, SpaceDopeSheetEditor, SpaceFileBrowser,SpaceGraphEditor, SpaceImageEditor, SpaceInfo, SpaceNLA, SpaceNodeEditor,SpaceOutliner, SpacePreferences, SpaceProperties, SpaceSequenceEditor,SpaceSpreadsheet, SpaceTextEditor, SpaceView3D

说明：场景的空间数据（Space data for a screen area）

### show_locked_time

说明：Synchronize the visible timeline range with other time-based editors

类型：boolean, default False

### show_region_header

说明：类型：boolean, default False

### type

说明：Space data type

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

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

### draw_handler_add()

全名：draw_handler_add(callback, args, region_type, draw_type)

说明：为这个空间类型添加一个新的绘制处理程序（drawhandler）。它将在每次空间类型中的指定区块（region）被绘制时被调用。注意：目前所有的参数都只是位置性的。

参数：

- callback：函数 。 一个当区域被绘制时将被调用的函数。它获得指定的参数作为输入。
- args：元组 。 将被传递给回调的参数。
- region_type：字符串。回调绘制的区块类型；通常是 WINDOW。(bpy.types.Region.type) 。
- draw_type：字符串。通常是 POST_PIXEL 用于 2D 绘图，POST_VIEW 用于 3D 绘图。在某些情况下可以使用 PRE_VIEW。BACKDROP 可以用于节点编辑器中的背景图。

返回：处理程序，可以在之后删除。

返回类型：object

### draw_handler_remove()

全名：draw_handler_remove(handler, region_type)

说明：移除之前添加的绘图处理程序。

参数：handler (object) 。 The draw handler that should be removed.

region_type (str) 。 Region type the callback was added to.
