---
title: KeyMap
order: 7
---

# KeyMap

## Description

[KeyMap(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyMap.html)

[KeyMapItem(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyMapItem.html)

[KeyMapItems(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyMapItems.html)

[KeyMaps(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyMaps.html)

[KeyConfig(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyConfig.html)

[KeyConfigurations(bpy_struct)](https://docs.blender.org/api/master/bpy.types.KeyConfigurations.html)

### 操作项快捷键

在 3D 视图按 ctrl shift K，会调用操作项 test.info，在底部打印 hello world

```python
import bpy


class TEST_OT_print_test(bpy.types.Operator):
    '''一个打印 info 的操作项'''
    bl_idname = "test.info"
    bl_label = "test ops"
    bl_options = {"REGISTER"}

    def execute(self, context):
        self.report({"INFO"}, "hello")
        return {"FINISHED"}


# 按键映射表
addon_keymaps = []


def register():
    # 注册操作项
    bpy.utils.register_class(TEST_OT_print_test)

    # 插件快捷键配置
    kcfg = bpy.context.window_manager.keyconfigs.addon

    if kcfg:
        # 设置快捷键应用环境，name就是首选项快捷键的分类，space_type对应空间名
        km = kcfg.keymaps.new(name='3D View', space_type='VIEW_3D')

        kmi = km.keymap_items.new(
            idname="test.info",    # 操作项标识符
            type="K",              # 按键，可以自己查表
            value="PRESS",         # 按键方式，查表
            shift=True,            # 是否需要同时按xxx键
            ctrl=True,
            alt=False,
            oskey=False
        )

    # 添加到快捷键表
    addon_keymaps.append((km, kmi))


def unregister():
    # 注销操作项
    bpy.utils.unregister_class(TEST_OT_print_test)

    # 注销快捷键
    for km, kmi in addon_keymaps:
        km.keymap_items.remove(kmi)
    addon_keymaps.clear()


if __name__ == "__main__":
    register()

```

### pie 菜单快捷键

ctrl shift k，弹出 pie menu （饼菜单）

```python

import bpy


class TEST_MT_pie_menu(bpy.types.Menu):
    '''简单的饼菜单'''
    bl_label = 'PIE TEST'
    bl_idname = 'TEST_MT_pie_menu'

    def draw(self, context):
        layout = self.layout
        pie = layout.menu_pie()

        pie.operator('mesh.primitive_cube_add', text="南", icon='EVENT_A')
        pie.operator('mesh.primitive_cube_add', text="北")


addon_keymaps = []


def register():
    bpy.utils.register_class(TEST_MT_pie_menu)

    wm = bpy.context.window_manager
    kcfg = bpy.context.window_manager.keyconfigs.addon
    if kcfg:
        km = kcfg.keymaps.new(name='3D View', space_type='VIEW_3D')
        kmi = km.keymap_items.new(
            idname="wm.call_menu_pie",  # 因为这里是操作项，所以要用 wm.call_menu_pie 操作项，用于调用饼菜单
            type="K",
            value="PRESS",
            shift=True,
            ctrl=True,
            alt=False,
            oskey=False
        )
    kmi.properties.name = TEST_MT_pie_menu.bl_idname # 这里写操作项的名称（官方没找到解释，死记吧）
    addon_keymaps.append((km, kmi))


def unregister():
    bpy.utils.unregister_class(TEST_MT_pie_menu)

    for km, kmi in addon_keymaps:
        km.keymap_items.remove(kmi)
    addon_keymaps.clear()


if __name__ == "__main__":
    register()

```

## KeyMap(bpy_struct)

基类： bpy_struct

说明：输入配置，包括键盘映射

### bl_owner_id

说明：Internal owner

类型：字符串， 默认为"", 不会为 None

### is_modal

说明：表示一个按键映射，被用于 translate 操作项 modal 事件

类型：boolean, default False, 只读

### is_user_modified

说明：按键映射由用户定义

类型：boolean, default False

### keymap_items

说明：按键映射中的项目，将一个操作项与一个输入事件联系起来

类型：KeyMapItems bpy_prop_collection of KeyMapItem, 只读

### name

说明：按键映射的名称

类型：字符串， 默认为"", (只读，可选)

### region_type

说明：与之相关的可选的按键映射的区域类型

类型：只读

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

### show_expanded_children

说明：用户界面的子扩展

类型：boolean, default False

### show_expanded_items

说明：在用户界面扩展

类型：boolean, default False

### space_type

说明：可选的空间类型按键映射

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

### active()

全名：active()

说明：激活

返回：Key Map, Active key map

返回类型：KeyMap

### restore_to_default()

全名：restore_to_default()

说明：恢复到默认状态

### restore_item_to_default()

全名：restore_item_to_default(item)

说明：恢复项目为默认值

参数：item (KeyMapItem), 不会为 None。项目

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyMapItem(bpy_struct)

基类： bpy_struct

### classbpy.types.KeyMapItem()

全名：classbpy.types.KeyMapItem(bpy_struct)

Item in a Key Map

### active

说明：Activate or deactivate item

类型：boolean, default False

### alt

说明：Alt key pressed, -1 for any state

类型：整数 [-1, 1], default 0

### alt_ui

说明：Alt key pressed

类型：boolean, default False

### any

说明：Any modifier keys pressed

类型：boolean, default False

### ctrl

说明：Control key pressed, -1 for any state

类型：整数 [-1, 1], default 0

### ctrl_ui

说明：Control key pressed

类型：boolean, default False

### id

说明：ID of the item

类型：整数 [-32768, 32767], default 0, 只读

### idname

说明：Identifier of operator to call on input event

类型：字符串， 默认为"", 不会为 None

### is_user_defined

说明：Is this keymap item user defined (doesn’t just replace a builtin item)

类型：boolean, default False, 只读

### is_user_modified

说明：Is this keymap item modified by the user

类型：boolean, default False, 只读

### key_modifier

说明：Regular key pressed as a modifier

NONE

类型：enum in [‘NON

### map_type

说明：Type of event mapping

类型：enum in [‘KEYBOARD’, ‘TWEAK’, ‘MOUSE’, ‘NDOF’, ‘TEXTINPUT’, ‘TIMER’],
default ‘KEYBOARD’

### name

说明：Name of operator (translated) to call on input event

类型：字符串， 默认为"", (只读，可选)

### oskey

说明：Operating system key pressed, -1 for any state

类型：整数 [-1, 1], default 0

### oskey_ui

说明：Operating system key pressed

类型：boolean, default False

### properties

说明：Properties to set when the operator is called

类型：OperatorProperties, 只读

### propvalue

说明：The value this event translates to in a modal keymap

类型：enum in [‘NONE’], default ‘NONE’

### repeat

说明：Active on key-repeat events (when a key is held)

类型：boolean, default False

### shift

说明：Shift key pressed, -1 for any state

类型：整数 [-1, 1], default 0

### shift_ui

说明：Shift key pressed

类型：boolean, default False

### show_expanded

说明：Show key map event and property details in the user interface

类型：boolean, default False

### type

说明：Type of event

NONE Un

类型：enum in [‘NONE

### value

说明：类型：enum in [‘ANY’, ‘PRESS’, ‘RELEASE’, ‘CLICK’, ‘DOUBLE_CLICK’,‘CLICK_DRAG’, ‘NORTH’, ‘NORTH_EAST’, ‘EAST’, ‘SOUTH_EAST’, ‘SOUTH’,‘SOUTH_WEST’, ‘WEST’, ‘NORTH_WEST’, ‘NOTHING’], default ‘NOTHING’

### compare()

全名：compare(item)

compare

参数：item (KeyMapItem) 。Item

返回：Comparison result

返回类型：boolean

### to_string()

全名：to_string(compact=False)

to_string

参数：compact 布尔值，可选) 。Compact

返回：result

返回类型：字符串， 不会为 None

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyMapItems(bpy_struct)

基类： bpy_struct

说明：按键映射 项目集（Collection of keymap items）

### new()

全名：new(idname, type, value, any=False, shift=0, ctrl=0, alt=0, oskey=0,key_modifier='NONE', repeat=False, head=False)

说明：新建按键映射

参数：

idname ：字符串， 不会为 None。操作项标识

type：[按键类型](https://www.yuelili.com/?p=19528)之一。

value：枚举， ['ANY', 'PRESS', 'RELEASE', 'CLICK', 'DOUBLE_CLICK', 'CLICK_DRAG','NORTH', 'NORTH_EAST', 'EAST', 'SOUTH_EAST', 'SOUTH', 'SOUTH_WEST', 'WEST','NORTH_WEST', 'NOTHING'])

any：布尔值，可选 。

shift：shift 键，整数， [-1, 1]，可选。

ctrl：ctrl 键，整数， [-1, 1]，可选) 。Ctrl

alt：alt 键，整数， [-1, 1]，可选) 。Alt

oskey：oskey 键，整数， [-1, 1]，可选) 。OS Key

key_modifier：[按键类型](https://www.yuelili.com/?p=19528)。

repeat：布尔值，可选。重复，当设置时，接受按键重复事件

head：布尔值，可选 。强制在快捷键表前面（而不是末尾）添加项目，这样就不会被现有快捷键干扰。

返回：Item，添加了按键映射的项目

返回类型：KeyMapItem

### new_modal()

全名：new_modal(propvalue, type, value, any=False, shift=0, ctrl=0, alt=0,oskey=0, key_modifier='NONE', repeat=False)

new_modal

参数：propvalue ：字符串， 不会为 None。Property Value

type：[按键类型](https://www.yuelili.com/?p=19528)之一。

value：(enum in ['ANY', 'PRESS', 'RELEASE', 'CLICK', 'DOUBLE_CLICK','CLICK_DRAG', 'NORTH', 'NORTH_EAST', 'EAST', 'SOUTH_EAST', 'SOUTH','SOUTH_WEST', 'WEST', 'NORTH_WEST', 'NOTHING'])

any：布尔值，可选 。

shift：shift 键，整数， [-1, 1]，可选。

ctrl：ctrl 键，整数， [-1, 1]，可选) 。Ctrl

alt：alt 键，整数， [-1, 1]，可选) 。Alt

oskey：oskey 键，整数， [-1, 1]，可选) 。OS Key

key_modifier：[按键类型](https://www.yuelili.com/?p=19528)之一。

repeat：布尔值，可选。重复，当设置时，接受按键重复的事件

返回：Item，添加了按键映射的项目

返回类型：KeyMapItem

### new_from_item()

全名：new_from_item(item, head=False)

new_from_item

参数：

item (KeyMapItem, 不会为 None。作为参考的项目

head 布尔值，可选) 。At Head

返回：Item, 添加映射的项目

返回类型：KeyMapItem

### remove()

全名：remove(item)

说明：移除项目

参数：item (KeyMapItem, 不会为 None。Item

### from_id()

全名：from_id(id)

from_id

参数：id 整数， [-inf, inf]) 。id, ID of the item

返回：Item

返回类型：KeyMapItem

### find_from_operator()

全名：find_from_operator(idname, properties=None, include={'ACTIONZONE','KEYBOARD', 'MOUSE', 'NDOF', 'TWEAK'}, exclude={})

说明：从操作项查找

参数：

idname ：字符串， 不会为 None。Operator Identifier

include (enum set in {'KEYBOARD_MODIFIER', 'KEYBOARD', 'MOUSE_WHEEL','MOUSE_GESTURE', 'MOUSE_BUTTON', 'MOUSE', 'NDOF', 'TWEAK', 'ACTIONZONE'}，可选)
。Include

exclude (enum set in {'KEYBOARD_MODIFIER', 'KEYBOARD', 'MOUSE_WHEEL','MOUSE_GESTURE', 'MOUSE_BUTTON', 'MOUSE', 'NDOF', 'TWEAK', 'ACTIONZONE'}，可选)
。Exclude

返回类型：KeyMapItem

### match_event()

全名：match_event(event)

match_event

返回类型：KeyMapItem

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyMaps(bpy_struct)

基类： bpy_struct

说明：按键映射的集合

### new()

全名：new(name, space_type='EMPTY', region_type='WINDOW', modal=False,
tool=False)

说明：确保按键映射存在。这将返回基于名称/空间类型/区块类型的映射。没有的话，则创建一个新映射

参数：

name ：名称，字符串， 不会为 None。对应 首选项 - 插件 - 分类

![](https://cdn.yuelili.com/20220121002326.png)

space_type：空间，space 类型，可选

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

region_type：区块，Region 类型，

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

modal：布尔值，可选。Modal, modal 操作项的按键映射

tool：布尔值，可选 。Tool, 激活工具的按键映射

返回：Key Map, Added key map

返回类型：KeyMap

### remove()

全名：remove(keymap)

说明：移除按键映射

参数：keymap (KeyMap, 不会为 None。Key Map, Removed key map

### find()

全名：find(name, space_type='EMPTY', region_type='WINDOW')

find

参数：name ：字符串， 不会为 None。Name

space_type：空间类型，可选

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

region_type：Region 类型，

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

返回：Key Map, Corresponding key map

返回类型：KeyMap

### find_modal()

全名：find_modal(name)

说明：查找 modal

参数：name ：字符串， 不会为 None。Operator Name

返回：Key Map, Corresponding key map

返回类型：KeyMap

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyConfig(bpy_struct)

基类： bpy_struct

### classbpy.types.KeyConfig()

全名：classbpy.types.KeyConfig(bpy_struct)

说明：输入配置，包含按键映射

### is_user_defined

说明：Indicates that a keyconfig was defined by the user

类型：boolean, default False, 只读

### keymaps

说明：Key maps configured as part of this configuration

类型：KeyMaps bpy_prop_collection of KeyMap, 只读

### name

说明：按键配置名称

类型：字符串， 默认为"", 不会为 None

### preferences

说明：类型：KeyConfigPreferences, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyConfigPreferences(bpy_struct)

基类： bpy_struct

### bl_idname

说明：类型：字符串， 默认为"", 不会为 None

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## KeyConfigurations(bpy_struct)

基类： bpy_struct

说明：按键配置集合

### active

说明：Active key configuration (preset)

类型：KeyConfig

### addon

说明：可由插件扩展的快捷键配置，并在处理事件时添加到活动配置中。

类型：KeyConfig, 只读

### default

说明：默认内置的按键配置

类型：KeyConfig, 只读

### user

说明：最终按键配置，结合了已激活和插件的配置，可以被用户修改

类型：KeyConfig, 只读

### new()

全名：new(name)

说明：新建按键配置

参数：name ，字符串， 不会为 None。Name

返回：Key Configuration, Added key configuration

返回类型：KeyConfig

### remove()

全名：remove(keyconfig)

说明：移除按键配置

参数：keyconfig (KeyConfig, 不会为 None。Key Configuration, Removed key configuration

### find_item_from_operator()

全名：find_item_from_operator(idname, context='INVOKE_DEFAULT', properties=None,include={'ACTIONZONE', 'KEYBOARD', 'MOUSE', 'NDOF', 'TWEAK'}, exclude={})

说明：从操作项寻找项目

参数：idname ：字符串， 不会为 None。Operator Identifier

context：上下文

- `INVOKE_DEFAULT` ：只执行 invoke()
- `INVOKE_REGION_WIN`：
- `INVOKE_REGION_CHANNELS`：
- `INVOKE_REGION_PREVIEW`：
- `INVOKE_AREA`：
- `INVOKE_SCREEN`：
- `EXEC_DEFAULT`：只执行 execute()
- `EXEC_REGION_WIN`：
- `EXEC_REGION_CHANNELS`：
- `EXEC_REGION_PREVIEW`：
- `EXEC_AREA`：
- `EXEC_SCREEN`：

include：枚举元组，可选

- {'KEYBOARD_MODIFIER', 'KEYBOARD', 'MOUSE_WHEEL', 'MOUSE_GESTURE', 'MOUSE_BUTTON', 'MOUSE', 'NDOF', 'TWEAK', 'ACTIONZONE'}，

exclude：枚举元组，可选

- {'KEYBOARD_MODIFIER', 'KEYBOARD', 'MOUSE_WHEEL', 'MOUSE_GESTURE', 'MOUSE_BUTTON', 'MOUSE', 'NDOF', 'TWEAK', 'ACTIONZONE'}

返回：(keymap, item)，keymap, KeyMap

item：KeyMapItem

### update()

全名：update()

update

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
