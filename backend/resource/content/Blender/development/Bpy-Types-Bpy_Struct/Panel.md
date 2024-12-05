---
title: Panel 面板
order: 9
---

# Panel 面板

## Description

[Panel(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Panel.html)

### 基本面板实例

一个简单的面板，它将绘制到对象属性部分。

::: tip
'面板命名规则：CATEGORY_PT_name ，Panel.bl_idname
:::

::: tip
面板子类必须注册，Blender 才能使用。
:::

直接在文本编辑器输就成，不走插件了。

![](https://cdn.yuelili.com/20220112213801.png)

```python
import bpy


class HelloWorldPanel(bpy.types.Panel):
    bl_idname = "OBJECT_PT_hello_world"
    bl_label = "Hello！"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        self.layout.label(text="Hello World")


bpy.utils.register_class(HelloWorldPanel)

```

### 简单对象面板

这个面板有一个`Panel.poll`和`Panel.draw_header`函数，即使内容很基本，它也非常类似于搅拌机面板。

直接在文本编辑器输就成，不走插件了

![](https://cdn.yuelili.com/20220112213514.png)

```python
import bpy

class ObjectSelectPanel(bpy.types.Panel):
    bl_idname = "OBJECT_PT_select"
    bl_label = "Select"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"
    bl_options = {'DEFAULT_CLOSED'}

    @classmethod
    def poll(cls, context):
        return (context.object is not None)

    def draw_header(self, context):
        layout = self.layout
        layout.label(text="我的选择面板")

    def draw(self, context):
        layout = self.layout

        box = layout.box()
        box.label(text="选择小工具")
        box.operator("object.select_all").action = 'TOGGLE'
        row = box.row()
        row.operator("object.select_all").action = 'INVERT'
        row.operator("object.select_random")

bpy.utils.register_class(ObjectSelectPanel)

```

### 混合类

混合父类可用于共享公共属性和 Menu.poll 功能。

![](https://cdn.yuelili.com/20220112214032.png)

```python
import bpy


class View3DPanel:
    bl_space_type = 'VIEW_3D'
    bl_region_type = 'UI'
    bl_category = "我是面板"

    @classmethod
    def poll(cls, context):
        return (context.object is not None)


class PanelOne(View3DPanel, bpy.types.Panel):
    bl_idname = "VIEW3D_PT_test_1"
    bl_label = "面板 One"

    def draw(self, context):
        self.layout.label(text="Small Class")


class PanelTwo(View3DPanel, bpy.types.Panel):
    bl_idname = "VIEW3D_PT_test_2"
    bl_label = "面板 Two"

    def draw(self, context):
        self.layout.label(text="Also Small Class")


bpy.utils.register_class(PanelOne)
bpy.utils.register_class(PanelTwo)

```

### 三大属性搭配

| bl_space_type | bl_region_type | bl_context                 | bl_category | 示例                                            | 说明                     |
| ------------- | -------------- | -------------------------- | ----------- | ----------------------------------------------- | ------------------------ |
| VIEW_3D       | UI             | objectmode（对象编辑模式） | Item        | ![](https://cdn.yuelili.com/20220118223615.png) | N 面板（分类覆盖则追加） |
| PROPERTIES    | WINDOW         | object（对象属性）         |             | ![](https://cdn.yuelili.com/20220112213514.png) | 属性面板-对象分类        |
| FILE_BROWSER  | TOOLS          |                            |             | ![](https://cdn.yuelili.com/20220119224209.png) | 文件浏览器， 左侧工具栏  |

## bpy.types.Panel(bpy_struct)

基类：bpy_struct，包含 UI 元素的面板

### bl_category

![](https://cdn.yuelili.com/20220112225540.png)

说明：在面板 (tab) 要显示的分类名称

类型：string, 默认为空“”, (不会为 None)

其他：如果名称一致，则会追加（因此你可以直接在 Blender 原始面板追加自定义面板）

### bl_context

其他 context，请参考[bpy.context](https://www.yuelili.com/docs/blender-master/bpy-context/)

![](https://cdn.yuelili.com/20220113092438.png)

说明：属于面板的上下文 (TODO: 组合使用 bl_context/bl_region_type/bl_space_type，可以搭配不同位置的面板)

类型：字符串，默认为“”, (不会为 None)

示例：

```python
bl_space_type = "VIEW_3D"
bl_region_type = "TOOLS"
bl_category = "Tools"

bl_space_type = "VIEW_3D"
bl_region_type = "TOOL_PROPS"

bl_space_type = "VIEW_3D"
bl_region_type = "UI"

bl_space_type = "PROPERTIES"
bl_region_type = "WINDOW"
bl_context = "object"

bl_space_type = "PROPERTIES"
bl_region_type = "WINDOW"
bl_context = "material"

```

### bl_description

说明：面板的小提示

类型：字符串，默认为“”

### bl_idname

说明：如果设置，面板会使用自定义 ID, 否则使用类的名，通常是面板. 比如类名为 “OBJECT_PT_hello”, 未设置该项，则 bl_idname
为 “OBJECT_PT_hello”

类型：字符串，默认为“”, (不会为 None)

### bl_label

![](https://cdn.yuelili.com/20220113091436.png)

说明：面板的标签, 当面板折叠时，显示在面板上方

类型：字符串，默认为“”, (不会为 None)

### bl_options

说明：此面板类型的选项

- DEFAULT_CLOSED：默认关闭。面板在创建时打开还是折叠。

![](https://cdn.yuelili.com/20220113091436.png)

- HIDE_HEADER：隐藏 Header 。如果设置为 False，面板会显示页眉，会包含一个可点击的箭头来折叠面板和标签（见 bl_label）。
- INSTANCED：实例化面板，这种类型的多个面板可以作为列表的一部分，取决于 UI 的外部数据。用于为修改器和其他堆栈创建面板。
- HEADER_LAYOUT_EXPAND：扩展 Header 布局 ，允许 Header 中按钮自适应填充布局宽度。

类型：枚举型，默认为{‘DEFAULT_CLOSED’}

### bl_order

说明：值越低，在面板中越靠前。

类型：[0, inf], 默认为 0

### bl_owner_id

说明：显示在面板上的数据 ID，如果有的话

类型：字符串，默认为“”, (不会为 None)

### bl_parent_id

说明：如果设置，则该面板会作为子面板（sub-panel）

类型：字符串，默认为“”, (never None)

示例：panel_2 放在 panel_1 的内部

![](https://cdn.yuelili.com/20220113093951.png)

```python
class TEST_PT_panel_1(bpy.types.Panel):
    ...

class TEST_PT_panel_2(bpy.types.Panel):
    bl_parent_id = "TEST_PT_panel_1" # 父类名
    ...

```

### bl_region_type

说明：面板要用于的区块（region）

类型：

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

参数：上方其中之一, 默认为‘WINDOW’

### bl_space_type 作用空间

![](https://cdn.yuelili.com/20220113090042.png)

说明：面板要用于的空间（space）

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

### bl_translation_context

说明：具体的翻译上下文, 只有存在多个相同标签才需要定义

类型：字符串，默认为“\*”, (不会为 None)

### bl_ui_units_x

说明：设置时，定义弹出面板宽度

类型：整数，[0, inf], 默认为 0

### custom_data

说明：面板数据

类型：约束（Constraint），只读

### is_popover

类型：布尔值，默认为 False，只读

### layout

说明：定义在 UI 中的面板构筑

类型：UILayout, 只读

### text

说明：XXX todo

类型：字符串，默认为“”, (never None)

### use_pin

说明：显示所有选项卡（tabs）中的面板（panel）

类型：布尔值，默认为 False

### classmethodpoll(context)

说明：只要返回非空，则绘制面板

返回类型：boolean

### draw(context)

说明：在 UI 布局中绘制 UI 元素

### draw_header(context)

![](https://cdn.yuelili.com/20220113091047.png)

说明：在 UI 头部布局中绘制 UI 元素

### draw_header_preset(context)

说明：在 UI 头部布局中，使用预设绘制 UI 元素

### classmethod bl_rna_get_subclass(id, default=None)

参数：id (string) – The RNA type identifier.

返回：RNA 类型。如果未找到，则为默认

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (string) – The RNA type identifier.

返回：类。如果未找到，则为默认

返回类型：type？？？
