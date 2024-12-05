---
title: 操作项(bpy.ops)
order: 6
---

# 操作项(bpy.ops)

## 对应官方文档链接

[bpy.ops](https://docs.blender.org/api/master/bpy.ops.html)：ops 类

[其他操作项子类总览](https://www.yuelili.com/docs/blender-master/overview-of-action-items/)：一些其他的操作项类

[bpy.types.Operator](https://docs.blender.org/api/master/bpy.types.Operator.html)：操作项的参数（虽然有点僭越，但我觉得放这里比较好）

### 应用操作项

用 python 访问操作项 ，包括 C、Python 或宏编写的操作项。

只有关键字参数可用于传递操作项属性。

操作项返回一个 set() ，它由以下部分组成：`{'RUNNING_MODAL', 'CANCELLED', 'FINISHED','PASS_THROUGH'}` 。常见的返回值是` {'FINISHED'}``{'CANCELLED'} `

在错误场景调用操作项会引发`RuntimeError`，可以使用 poll() 方法避免。（比如你在物体模式移动“编辑模式”下的点）

#### 关键字和位置参数

关键字对应操作项 **属性** ，位置参数用于定义操作项的 **调用方式** 。

有 3 个可选的位置参数（下面详细记录）。

```python
bpy.ops.test.operator(override_context, execution_context, undo)

```

- override_context ：覆写，`dict`类型。
- execution_context ：执行区域， `str`（枚举）。
- undo ：撤销，`bool`类型。

上面每个都是可选的，但必须按照顺序书写。

```python
import bpy

# 调用操作项
bpy.ops.mesh.subdivide(number_cuts=3, smoothness=0.5)


# 用poll()检查工作区 避免发生意外.
if bpy.ops.object.mode_set.poll():
    bpy.ops.object.mode_set(mode='EDIT')

```

### 覆盖上下文

覆盖操作项的上下文成员。这样可以不仅限于操作当前项。

可以扩充 **影响范围** ，比如从当前选择项 → 全部对象（见示例）。还可以提升作用空间，本来在物体模式运行，现在可以扩展到其他模式。

上下文覆盖作为字典传递，键名要与 bpy.context 中的上下文成员名匹配。比如覆盖 active_object（`bpy.context.active_object`），需要使用`{'active_object': object}`。X[键名] = 新的作用成员

::: tip
大家都希望使用当前上下文的副本作为基础（不然还得自己检索数据）。
:::

```python
# 覆写：删除当前对象 → 删除场景中所有对象
import bpy

# 复制一份context，因为context只能读
override = bpy.context.copy()
# 覆写：把当前选择的对象变成场景中所有对象
override['selected_objects'] = list(bpy.context.scene.objects)
# 执行删除：此时删除当前 → 删除全部
bpy.ops.object.delete(override)

```

### 执行上下文

调用操作项时，执行上下文的哪些内容。

决定操作项的运行环境，以及是调用 invoke() 还是只调用 execute()等等。

默认情况下使用`EXEC_DEFAULT`，仅运行`execute()`方法，如果存在`INVOKE_DEFAULT`将调用 invoke()。

执行上下文范围：

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

```python
## 只执行invoke函数
import bpy
bpy.ops.object.collection_instance_add('INVOKE_DEFAULT')
```

也可以在用户界面的特定位置运行。为此，我们需要传递窗口、屏幕、区域，有时还需要传递一个区域。

```python
# 所有窗口的3D视图都最大化显示
import bpy

for window in bpy.context.window_manager.windows:
    screen = window.screen

    for area in screen.areas:
        if area.type == 'VIEW_3D':
            override = {'window': window, 'screen': screen, 'area': area}
            bpy.ops.screen.screen_full_area(override)
            break
```

### 命名规则

![](https://cdn.yuelili.com/20220113084517.png)

| Property           | 缩写 | 名称    |
| ------------------ | ---- | ------- |
| bpy.types.Operator | OT   | 操作项  |
| ---                | ---  | ---     |
| bpy.types.Panel    | PT   | 面板    |
| bpy.types.Header   | HT   | Header  |
| bpy.types.MENU     | MT   | 菜单    |
| bpy.types.UIList   | UL   | UI 列表 |

另见[bpy.types.Operator](https://docs.blender.org/api/master/bpy.types.Operator.html)

### 最简单的操作项

一个简单的操作项，可以打印一条信息。

由于该操作项只有一个 Operator.execute 函数，所以不需要用户输入。

::: tip
在 blender 访问操作项的子类之前，必须先注册。
:::

![](https://cdn.yuelili.com/20220113132251.png)

_在控制台使用 bl_idname 进行访问_

![](https://cdn.yuelili.com/20220113132225.png)

_用 F3 搜索，使用 bl_label 进行访问_

![](https://cdn.yuelili.com/20220113133343.png)

_添加到视图菜单(view)，可以直接在最下面看到_

代码：

```python
import bpy

class HelloWorldOperator(bpy.types.Operator):
    bl_idname = "wm.hello_world" # id, 可以在控制台使用，或者被其他类调用
    bl_label = "最轻量的操作项"   # 显示名称，使用F3搜索bl_label，可以查找操作项

    # 要执行的函数
    def execute(self, context):
        print("Hello World")
        return {'FINISHED'}

# 如果想加入动态菜单
def menu_func(self, context):
    self.layout.operator(HelloWorldOperator.bl_idname, text="我在这儿")

# 注册（必须注册）
bpy.utils.register_class(HelloWorldOperator)

# 添加到视图菜单中
bpy.types.VIEW3D_MT_view.append(menu_func)


```

### Invoke 调用函数

invoke()通常用于分配属性，然后由 execute()来使用。

有些运算符没有 execute()函数，就没法重复使用这个函数了。

本例展示定义一个操作项，获取鼠标输入来执行一个函数。并且这个操作项可以从 python api 中调用或执行。

这个操作项还定义了自己的属性，这些属性与典型的类属性不同，因为 Blender 将它们与操作符注册在一起，在调用时作为参数使用，为操作符的撤销/重做保存，并自动添加到用户界面。

![](https://cdn.yuelili.com/20220113134937.png)

_操作项提示区_

![](https://cdn.yuelili.com/20220113135317.png)

_report 信息展示_

```python
import bpy


class SimpleMouseOperator(bpy.types.Operator):
    """ 本操作项显示鼠标位置 ,
        这里同时也是api提示区
    """


    bl_idname = "wm.mouse_position"
    bl_label = "Invoke Mouse Operator"

    # 定义俩整型变量
    x: bpy.props.IntProperty()
    y: bpy.props.IntProperty()

    def execute(self, context):
        # 使用report函数显示结果,
        self.report({'INFO'}, "鼠标当前位于 %d %d" % (self.x, self.y))
        return {'FINISHED'}

    def invoke(self, context, event):
        self.x = event.mouse_x
        self.y = event.mouse_y
        return self.execute(context)

# 动态添加函数
def menu_func(self, context):
    self.layout.operator(SimpleMouseOperator.bl_idname, text="Simple Mouse Operator")

# 注册并添加到view面板
bpy.utils.register_class(SimpleMouseOperator)
bpy.types.VIEW3D_MT_view.append(menu_func)


# 使用本函数
bpy.ops.wm.mouse_position('INVOKE_DEFAULT')

# 使用本函数（并自带参数）
# bpy.ops.wm.mouse_position('EXEC_DEFAULT', x=20, y=66)


```

### 调用文件选择器

本例展示如何使用文件选择器。（window_manager.fileselect_add）

invoke 函数调用了一个窗口管理器（window manager）方法，并返回 `{'RUNNING_MODAL'}`,模式。此时文件选择器将持续运行，直到 invoke 结束。

运行文件选择器，用户确认后，将调用 `Operator.execute` 。

Operator.poll 函数可选，用于检查操作项是否可以运行（作用范围）。

![](https://cdn.yuelili.com/20220113140243.png)

![](https://cdn.yuelili.com/20220113140159.png)

```python
import bpy


class ExportSomeData(bpy.types.Operator):
    """测试文件选择器"""
    bl_idname = "export.some_data"
    bl_label = "Export Some Data"

    filepath: bpy.props.StringProperty(subtype="FILE_PATH")

    @classmethod
    def poll(cls, context):
        return context.object is not None

    def execute(self, context):
        file = open(self.filepath, 'w')
        file.write("Hello World " + context.object.name)
        return {'FINISHED'}

    def invoke(self, context, event):
        # 使用一下 window_manager
        context.window_manager.fileselect_add(self)
        return {'RUNNING_MODAL'}



def menu_func(self, context):
    self.layout.operator_context = 'INVOKE_DEFAULT'
    self.layout.operator(ExportSomeData.bl_idname, text="Text Export Operator")


bpy.utils.register_class(ExportSomeData)
bpy.types.TOPBAR_MT_file_export.append(menu_func)


# 测试调用
bpy.ops.export.some_data('INVOKE_DEFAULT')

```

### Dialog Box 对话窗口

使用`Operator.invoke`函数调用弹出窗口。使用了（window_manager.invoke_props_dialog）

```python
import bpy


class DialogOperator(bpy.types.Operator):
    bl_idname = "object.dialog_operator"
    bl_label = "简单的对话框"

    my_float: bpy.props.FloatProperty(name="一个浮点数")
    my_bool: bpy.props.BoolProperty(name="开/关")
    my_string: bpy.props.StringProperty(name="一个字符串")

    def execute(self, context):
        message = (
            "用户输入了: %f, %d, '%s'" %
            (self.my_float, self.my_bool, self.my_string)
        )
        self.report({'INFO'}, message)
        return {'FINISHED'}

    def invoke(self, context, event):
        wm = context.window_manager
        return wm.invoke_props_dialog(self)


def menu_func(self, context):
    self.layout.operator(DialogOperator.bl_idname, text="Dialog Operator")


bpy.utils.register_class(DialogOperator)
bpy.types.VIEW3D_MT_object.append(menu_func)

# 测试调用
bpy.ops.object.dialog_operator('INVOKE_DEFAULT')

```

### 自定义绘制

默认情况下，操作项属性自动使用用户界面布局。如果需要更多控制，可以使用`Operator.draw`函数创建自定义布局 。

这类似于`Panel`和`Menu`绘制函数，也可用于对话框和文件选择器。

示例：弹出自定义框

```python
import bpy


class CustomDrawOperator(bpy.types.Operator):
    bl_idname = "object.custom_draw"
    bl_label = "简单的Modal操作项"

    filepath: bpy.props.StringProperty(subtype="FILE_PATH")

    my_float: bpy.props.FloatProperty(name="浮点X")
    my_bool: bpy.props.BoolProperty(name="开关Y")
    my_string: bpy.props.StringProperty(name="字符串Z")

    def execute(self, context):
        print("Test", self)
        return {'FINISHED'}

    def invoke(self, context, event):
        wm = context.window_manager
        return wm.invoke_props_dialog(self)

    def draw(self, context):
        # 创建布局
        layout = self.layout

        # 第一行
        col = layout.column()
        col.label(text="我的自定义界面")

        # 第二行 row 横向布局
        row = col.row()
        row.prop(self, "my_float")
        row.prop(self, "my_bool")

        # 第三行
        col.prop(self, "my_string")

def menu_func(self, context):
    self.layout.operator(CustomDrawOperator.bl_idname, text="Custom Draw Operator")

bpy.utils.register_class(CustomDrawOperator)
bpy.types.VIEW3D_MT_object.append(menu_func)

# 测试
bpy.ops.object.custom_draw('INVOKE_DEFAULT')

```

### Modal 模态执行

定义一个`Operator.modal`函数，该函数将继续持续处理某个事件，直到返回`{'FINISHED'}`或`{'CANCELLED'}`，方可停止该对话。

每次检测到新事件（例如鼠标单击或按键）时，都会运行模态操作项。相反，没有检测到新事件时，模态操作项不会运行。

这对于交互式工具特别有用，可以有自己的状态，在运行时按键切换选项等等。

Grab、Rotate、Scale 和 Fly-Mode 都是模态运算符的示例。

`Operator.invoke`先返回`{'RUNNING_MODAL'}`来初始化模态循环。

通知`__init__()`并被`__del__()`申明。对于其他操作项类型没有用，但对于模态操作项，它们将在`Operator.invoke`之前和操作项完成之后调用。

示例：检测鼠标移动，并应用给物体

```python
import bpy


class ModalOperator(bpy.types.Operator):
    '''检测鼠标移动，并应用给物体'''
    bl_idname = "object.modal_operator"
    bl_label = "简单模态操作项"

    def __init__(self):
        print("开始")

    def __del__(self):
        print("结束")

    def execute(self, context):
        context.object.location.x = self.value / 100.0
        return {'FINISHED'}

    def modal(self, context, event):
        if event.type == 'MOUSEMOVE':  # 鼠标移动，执行excute
            self.value = event.mouse_x
            self.execute(context)
        elif event.type == 'LEFTMOUSE':  # 鼠标左键，确认位置，并结束
            return {'FINISHED'}
        elif event.type in {'RIGHTMOUSE', 'ESC'}:   # 鼠标右键，位置还原，并结束
            context.object.location.x = self.init_loc_x
            return {'CANCELLED'}

        return {'RUNNING_MODAL'}

    def invoke(self, context, event):
        self.init_loc_x = context.object.location.x
        self.value = event.mouse_x
        self.execute(context)

        context.window_manager.modal_handler_add(self)
        return {'RUNNING_MODAL'}


def menu_func(self, context):
    self.layout.operator(ModalOperator.bl_idname, text="Modal Operator")

bpy.utils.register_class(ModalOperator)
bpy.types.VIEW3D_MT_object.append(menu_func)

# test
bpy.ops.object.modal_operator('INVOKE_DEFAULT')

```

### 枚举搜索弹出窗口 Enum

使用`bpy.types.Operator.invoke_search_popup`，可以提示用户从搜索字段中选择一个项目

```python
import bpy
from bpy.props import EnumProperty


class SearchEnumOperator(bpy.types.Operator):
    bl_idname = "object.search_enum_operator"
    bl_label = "Search Enum Operator"
    bl_property = "my_search"

    # 枚举属性
    my_search: EnumProperty(
        name="My Search",
        items=(
            ('FOO', "Foo", ""),
            ('BAR', "Bar", ""),
            ('BAZ', "Baz", ""),
        ),
    )

    def execute(self, context):
        self.report({'INFO'}, "已选择:" + self.my_search)
        return {'FINISHED'}

    def invoke(self, context, event):
        context.window_manager.invoke_search_popup(self)
        return {'RUNNING_MODAL'}

def menu_func(self, context):
    self.layout.operator(SearchEnumOperator.bl_idname, text="Search Enum Operator")

bpy.utils.register_class(SearchEnumOperator)
bpy.types.VIEW3D_MT_object.append(menu_func)

# test
bpy.ops.object.search_enum_operator('INVOKE_DEFAULT')

```

base class — bpy_struct

classbpy.types.Operator(bpy_struct)

说明：Storage of an operator being executed, or registered after execution

### bl_cursor_pending

说明：用户等待时的鼠标状态（当 bl_options 设置了 DEPENDS_ON_CURSOR 时）。

类型：默认为 "DEFAULT "。

- DEFAULT
- NONE
- WAIT
- CROSSHAIR
- MOVE_X
- MOVE_Y
- KNIFE
- TEXT
- PAINT_BRUSH
- PAINT_CROSS
- DOT
- ERASER
- HAND
- SCROLL_X
- SCROLL_Y
- SCROLL_XY
- EYEDROPPER
- PICK_AREA
- STOP
- COPY
- CROSS
- MUTE
- ZOOM_IN
- ZOOM_OUT

### bl_description

![](https://cdn.yuelili.com/20220113134937.png)

说明：描述，或者某些区域的提示信息，比如控制台按 tab 补全的时候

类型：字符串，默认为""

### bl_idname

![](https://cdn.yuelili.com/20220113132251.png)

说明：ID 名称，加上类名，可以在其他区域调用（控制台、别的类、面板插入之类）

类型：字符串，默认为""

```python
bl_idname = "yll.my_operator" >>> bpy.ops.yll.my_operator
```

### bl_label

![](https://cdn.yuelili.com/20220113152918.png)

说明：显示名称，搜索的时候就能看到

类型：字符串，默认为""

### bl_options

说明：操作项的设置

类型：默认为{"REGISTER"}。

- REGISTER：注册。在信息窗口中显示，支持重做工具条面板。
- UNDO：撤销。推送一个撤销事件（操作项重做需要）。
- UNDO_GROUPED：分组撤销。为该操作项的重复实例推送一个撤销事件。
- BLOCKING：阻止。阻止其他东西使用光标。
- MACRO：宏。用来检查一个操作项是否是宏。
- GRAB_CURSOR Grab Pointer：用来使操作项抓取鼠标焦点，当启用连续抓取时启用包装。
- GRAB_CURSOR_X Grab Pointer X：抓取，只对 X 轴进行扭曲。
- GRAB_CURSOR_Y Grab Pointer Y：抓取，只对 Y 轴进行扭曲。
- DEPENDS_ON_CURSOR Depends on Cursor：。使用初始光标位置，当从菜单或按钮运行时，在开始操作前会提示用户放置光标。
- PRESET：预设。显示具有操作项设置的预设按钮。
- INTERNAL：内部。F3 搜索中移除本操作项。

### bl_translation_context

说明：

类型：字符串，默认为“Operator”, (never None)

### bl_undo_group

说明：

类型：字符串，默认为""

### has_reports

说明：操作项有上次执行的报告（警告和错误）。

类型：boolean, default False, (readonly)

### layout

说明：布局 UI

类型：UILayout, (只读)

```python
layout = self.layout balabala 创建布局
```

### macros

说明：

类型：bpy_prop_collection of Macro, (只读)

### name

说明：名称

类型：字符串，默认为"", (只读，不为 None)

### options

说明：运行时的选项

类型：OperatorOptions, (只读，不为 None)

### properties

说明：属性

类型：OperatorProperties, (只读，不为 None)

### bl_property

说明：用作此操作项主要属性的属性名称。目前只用于在将操作项扩展到菜单时选择默认属性。

类型： string

### report(type, message)

全名：report(type, message)

说明：report

参数：

- type：类型。{类型列表之一}，比如 {'INFO'}
- message：要传递的信息，字符串

类型列表：

- DEBUG：调试
- INFO：信息
- OPERATOR：操作项
- PROPERTY：属性
- WARNING：警告
- ERROR：错误
- ERROR_INVALID_INPUT：非法输入
- ERROR_INVALID_CONTEXT：非法上下文
- ERROR_OUT_OF_MEMORY：内存不足

  self.report({'INFO'}, "hello")

```


    ## is_repeat() #




    说明：is_repeat




    返回：result




    返回类型：boolean




    ## classmethodpoll() #




    全名：classmethodpoll(context)




    说明：测试该操作项是否可以被调用




    返回类型：boolean




    ## execute() #




    全名：execute(context)




    说明：执行操作项




    参数：context，代表要处理的内容，不是bpy.context喔




    返回：result，返回类型列表之一。比如：{"FINISHED"}




    返回类型列表：






      * RUNNING_MODAL：保持运行模式，操作项与blender一起运行。


      * CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。


      * FINISHED：完成。操作项在完成其动作后退出。


      * PASS_THROUGH：跳过。不做任何事情，将事件传递下去。


      * INTERFACE：界面。 处理但不执行（弹出式菜单）。





    ## check() #




    全名：check(context)




    说明：检查操作项的设置，返回True说明重绘。




    返回：result




    返回类型：boolean




    ## invoke() #




    全名：invoke(context, event)




    说明：调用操作项




    参数：






      * context：代表要处理的内容


      * event：传递进来的事件。比如鼠标移动，按键信息之类





    返回：result，返回类型列表之一。比如：{"FINISHED"}




    返回类型列表：






      * RUNNING_MODAL：保持运行模式，操作项与blender一起运行。


      * CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。


      * FINISHED：完成。操作项在完成其动作后退出。


      * PASS_THROUGH：跳过。不做任何事情，将事件传递下去。


      * INTERFACE：界面。 处理但不执行（弹出式菜单）。





    ## modal() #




    全名：modal(context, event)




    说明：模态操作项函数




    参数：






      * context：代表要处理的内容


      * event：传递进来的事件。比如鼠标移动，按键信息之类





    返回：result，返回类型列表之一。比如：{"FINISHED"}




    返回类型列表：






      * RUNNING_MODAL：保持运行模式，操作项与blender一起运行。


      * CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。


      * FINISHED：完成。操作项在完成其动作后退出。


      * PASS_THROUGH：跳过。不做任何事情，将事件传递下去。


      * INTERFACE：界面。 处理但不执行（弹出式菜单）。





    ## draw() #




    全名：draw(context)




    参数：context，代表要处理的内容




    说明：绘制操作项。draw 是实时刷新




    ## cancel() #




    全名：cancel(context)




    说明：当操作项取消，执行的内容




    ## classmethod description() #




    全名：classmethoddescription(context, properties)




    说明：计算一个取决于参数的描述字符串




    返回：result




    返回类型：string




    ## as_keywords() #




    全名：as_keywords(*, ignore=())




    说明：以字典的形式返回属性的副本




    ## classmethod bl_rna_get_subclass() #




    全名：classmethodbl_rna_get_subclass(id, default=None)




    说明：




    参数：id (string) – The RNA type identifier.




    返回：RNA类型，如果没有找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethodbl_rna_get_subclass_py(id, default=None)




    说明：参数：id (string) – The RNA type identifier.




    返回：类，如果没有找到，则为默认。




    返回类型：type




    ## poll_message_set() #




    全名：poll_message_set(message, *args)




    说明：Set the message to show in the tool-tip when poll fails.




    当消息是可调用的，额外的用户定义的位置参数被传递给消息函数。




    参数：message，(string or a callable that returns a string or None.) 。消息或一个返回消息的函数。




     




     





```
