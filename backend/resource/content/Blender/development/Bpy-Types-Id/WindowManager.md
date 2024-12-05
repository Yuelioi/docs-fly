---
title: WindowManager 窗口管理器
order: 3
---

# WindowManager 窗口管理器

## Description

[WindowManager(ID)](https://docs.blender.org/api/master/bpy.types.WindowManager.html)

基类：bpy_struct, ID

说明：窗口管理器（WindowManager）数据块，定义打开的窗口和其他用户界面数据

### addon_filter

说明：按类别过滤插件

类型：enum in [], 默认为""

### addon_search

说明：在选定的过滤器内搜索

类型：字符串，默认为"", 不会为 None

### addon_support

说明：插件支持程度

- OFFICIAL：官方支持。
- COMMUNITY：由社区开发者维护。
- TESTING：测试。新贡献的脚本（不包括在发布的构建中）。

类型：默认为 {"COMMUNITY,OFFICIAL"}

### asset_path_dummy

说明：Full path to the Blender file containing the active asset

类型：字符串，默认为"", 只读，不会为 None

### is_interface_locked

说明：如果为 True，该接口目前被一个正在运行的作业锁定，数据不应该从应用计时器中被修改。否则，正在运行的作业可能与处理程序冲突，导致  
意外结果，甚至崩溃

类型：布尔值，默认为 False，只读

### keyconfigs

说明：已注册的密钥配置

类型：KeyConfigurations bpy_prop_collection of KeyConfig, (readonly)

### operators

说明：操作项注册

类型：bpy_prop_collection of Operator, (readonly)

### pose_assets

说明：类型：bpy_prop_collection of AssetHandle, (readonly)

### poselib_flipped

说明：类型：boolean, default False

### poselib_previous_action

说明：类型：Action

### preset_name

说明：新预设的名称

类型：字符串，默认为“New Preset”, 不会为 None

### windows

说明：打开窗口

类型：bpy_prop_collection of Window, (readonly)

### xr_session_settings

说明：类型：XrSessionSettings, (readonly, never None)

### xr_session_state

说明：关于 VR 会话的运行时状态信息

类型：XrSessionState, (readonly)

### classmethod fileselect_add()

全名：classmethodfileselect_add(operator)

说明：用一个操作符打开一个文件选择器。

在操作项中出现时，分配字符串属性 "filepath,filename,directory "和 "files "集合。

参数：operator (Operator) – Operator to call

### classmethod modal_handler_add(operator)

全名：classmethodmodal_handler_add(operator)

说明：为窗口管理器添加一个模态 Modal 处理程序，用于给定的模态操作项。（在 return{'RUNNING_MODAL'}之前，由 invoke()与 self 调用）。

参数：：operator ，要调用的操作项

返回：处理程序是否添加成功

返回类型：布尔值

### event_timer_add()

全名：event_timer_add(time_step, window=None)

说明：为给定的窗口添加计时器，用于定期产生 "TIMER "事件。

参数：

- time_step：整型， [0, inf]) ，两次 timer events 之间的时间间隔，
- window ，窗口，可选。计时器要绑定的窗口，可以为 None

返回类型：Timer

### event_timer_remove()

全名：event_timer_remove(timer)

说明：event_timer_remove

### classmethodgizmo_group_type_ensure()

全名：classmethodgizmo_group_type_ensure(identifier)

说明：激活一个现有的插件组（widget group）（设置 persistent 选项时）。

参数：identifier (字符串 不会为 None) – Gizmo group type name

### classmethodgizmo_group_type_unlink_delayed()

全名：classmethodgizmo_group_type_unlink_delayed(identifier)

说明：解除插件组的链接（设置 persistent 选项时）。

参数：identifier ，字符串 不会为 None。 – Gizmo group type name

### progress_begin()

全名：progress_begin(min, max)

说明：开始进程 report

参数：min (float in [-inf, inf]) – min, any value in range [0,9999]

max (float in [-inf, inf]) – max, any value in range [min+1,9998]

### progress_update()

全名：progress_update(value)

说明：Update the progress feedback

参数：value (float in [-inf, inf]) – value, Any value between min and max as set
in progress_begin()

### progress_end()

全名：progress_end()

说明：Terminate progress report

### classmethod invoke_props_popup()

全名：classmethod invoke_props_popup(operator, event)

说明：弹出式调用操作项（显示操作项属性，只要变化就会自动执行）。

参数：

- operator：要调用的操作项
- event：事件

返回：result

- RUNNING_MODAL：保持运行模式，操作项与 blender 一起运行。
- CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。
- FINISHED：完成。操作项在完成其动作后退出。
- PASS_THROUGH：跳过。不做任何事情，将事件传递下去。
- INTERFACE：界面。 处理但不执行（弹出式菜单）。

返回类型：以上之一

示例：

![](https://cdn.yuelili.com/20220118001203.png)

F3，输入 Test My UI，然后弹出“操作项对话框”

此时你可以更改属性，此时更改属性时，会直接执行 execute()，也就是会打印 mStr 的内容

```python
import bpy
from bpy.props import StringProperty, BoolProperty


class Test_OT_test_ui(bpy.types.Operator):
    bl_idname = "test.ui"
    bl_label = "Test My UI"
    bl_options = {"REGISTER", "UNDO"}

    # 创建2个RNA属性
    mStr: StringProperty(
        default="我是字符串",
        description="我是属性说明",
    )

    mBool: BoolProperty(
        name="我是布尔值",
        description="我是属性说明",
        default=True,
    )

    def execute(self, context):
        # 控制台打印字符串，可以不要
        self.report({"INFO"}, self.mStr)
        return {"FINISHED"}

    def invoke(self, context, event):
        return context.window_manager.invoke_props_popup(self,event)


bpy.utils.register_class(Test_OT_test_ui)
```

### classmethod invoke_props_dialog()

全名：classmethod invoke_props_dialog(operator, width=300)

说明：操作项对话框（非自动执行弹出式）调用（显示操作项属性，点击确定按钮后执行）。

参数：

- operator：要调用的操作项
- width ：弹窗宽度，整数，可选

返回：result

- RUNNING_MODAL：保持运行模式，操作项与 blender 一起运行。
- CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。
- FINISHED：完成。操作项在完成其动作后退出。
- PASS_THROUGH：跳过。不做任何事情，将事件传递下去。
- INTERFACE：界面。 处理但不执行（弹出式菜单）。

返回类型：以上之一

示例：

![](https://cdn.yuelili.com/20220118001203.png)

F3，输入 Test My UI，然后弹出“操作项对话框”

此时更改属性，单击 OK 确定后，会执行 execute()，也就是会打印 mStr 的内容

```python
import bpy
from bpy.props import StringProperty, BoolProperty


class Test_OT_test_ui(bpy.types.Operator):
    bl_idname = "test.ui"
    bl_label = "Test My UI"
    bl_options = {"REGISTER", "UNDO"}

    # 创建2个RNA属性
    mStr: StringProperty(
        default="我是字符串",
        description="我是属性说明",
    )

    mBool: BoolProperty(
        name="我是布尔值",
        description="我是属性说明",
        default=True,
    )

    def execute(self, context):
        # 控制台打印字符串，可以不要
        self.report({"INFO"}, self.mStr)
        return {"FINISHED"}

    def invoke(self, context, event):
        return context.window_manager.invoke_props_dialog(self,width = 200)


bpy.utils.register_class(Test_OT_test_ui)

```

### classmethod invoke_search_popup() 弹出搜索式操作项

全名：classmethodinvoke_search_popup(operator)

说明：弹出调用搜索式操作项，搜索操作项的 bpy.types.Operator.bl_property（必须是 EnumProperty），在确认后执行。

参数：operator，要调用的操作项

### classmethod invoke_popup()

全名：classmethodinvoke_popup(operator, width=300)

说明：操作项弹出式调用（只显示操作项属性，不执行）。

参数：

- operator ：要执行的操作项
- width，整数，[0, inf]，可选，弹窗的宽度

返回：result

- RUNNING_MODAL：保持运行模式，操作项与 blender 一起运行。
- CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。
- FINISHED：完成。操作项在完成其动作后退出。
- PASS_THROUGH：跳过。不做任何事情，将事件传递下去。
- INTERFACE：界面。 处理但不执行（弹出式菜单）。

返回类型：以上之一

示例：

![](https://cdn.yuelili.com/20220118001203.png)

F3，输入 Test My UI，然后弹出“操作项对话框”

此时你可以更改属性，单击 OK 确定后，啥也不干（不会进入 invoke），但是属性值会更改

```python
import bpy
from bpy.props import StringProperty, BoolProperty


class Test_OT_test_ui(bpy.types.Operator):
    bl_idname = "test.ui"
    bl_label = "Test My UI"
    bl_options = {"REGISTER", "UNDO"}

    # 创建2个RNA属性
    mStr: StringProperty(
        default="我是字符串",
        description="我是属性说明",
    )

    mBool: BoolProperty(
        name="我是布尔值",
        description="我是属性说明",
        default=True,
    )

    def execute(self, context):
        # 控制台打印字符串，可以不要
        self.report({"INFO"}, self.mStr)
        return {"FINISHED"}

    def invoke(self, context, event):
        return context.window_manager.invoke_popup(self,width = 200)


bpy.utils.register_class(Test_OT_test_ui)

```

### classmethod invoke_confirm()

全名：classmethodinvoke_confirm(operator, event)

说明：确认操作项弹窗（只让用户确认执行，不显示操作项属性）

参数：

- operator ：要执行的操作项
- event：事件

返回：result

- RUNNING_MODAL：保持运行模式，操作项与 blender 一起运行。
- CANCELLED：取消。 操作项不做任何事情退出，所以不应该推送撤销条目。
- FINISHED：完成。操作项在完成其动作后退出。
- PASS_THROUGH：跳过。不做任何事情，将事件传递下去。
- INTERFACE：界面。 处理但不执行（弹出式菜单）。

返回类型：以上之一

示例：

![](https://cdn.yuelili.com/20220118002743.png)

F3，输入 Test My UI，然后弹出“确认对话框”，此时不会显示属性

单击确定后，会进入 invoke 函数

```python
import bpy
from bpy.props import StringProperty, BoolProperty


class Test_OT_test_ui(bpy.types.Operator):
    bl_idname = "test.ui"
    bl_label = "Test My UI"
    bl_options = {"REGISTER", "UNDO"}

    # 创建2个RNA属性
    mStr: StringProperty(
        default="我是字符串",
        description="我是属性说明",
    )

    mBool: BoolProperty(
        name="我是布尔值",
        description="我是属性说明",
        default=True,
    )

    def execute(self, context):
        # 控制台打印字符串，可以不要
        self.report({"INFO"}, self.mStr)
        return {"FINISHED"}

    def invoke(self, context, event):
        return context.window_manager.invoke_confirm(self,event)


bpy.utils.register_class(Test_OT_test_ui)

```

### classmethod popmenu_begin\_\_internal()

全名：classmethodpopmenu_begin\_\_internal(title, icon='NONE')

说明：popmenu_begin\_\_internal

参数：icon，字符串，[系统图标](https://www.yuelili.com/blender-development-use-blenders-own-icon/)列表里任选一个

### classmethod popmenu_end\_\_internal()

全名：classmethod popmenu_end\_\_internal(menu)

popmenu_end\_\_internal

### classmethod popover_begin\_\_internal()

全名：classmethod popover_begin\_\_internal(ui_units_x=0, from_active_button=False)

说明：popover_begin\_\_internal

参数：

- ui_units_x ，整数，[0, inf]，可选，ui_units_x
- from_active_button ，布尔值，可选，使用按钮, 使用活动按钮进行定位

返回类型：UIPopover, 不会为 None

### classmethod popover_end\_\_internal()

全名：classmethodpopover_end\_\_internal(menu, keymap=None)

说明：popover_end\_\_internal

参数：keymap (KeyMap, (optional)) – Key Map, Active key map

### classmethod piemenu_begin\_\_internal()

全名：classmethodpiemenu_begin\_\_internal(title, icon='NONE', event=None)

参数：icon，字符串，[系统图标](https://www.yuelili.com/blender-development-use-blenders-own-icon/)列表里任选一个

说明：piemenu_begin\_\_internal

返回类型：UIPieMenu, 不会为 None

### classmethod piemenu_end\_\_internal()

全名：classmethodpiemenu_end\_\_internal(menu)

说明：piemenu_end\_\_internal

### classmethodoperator_properties_last()

全名：classmethodoperator_properties_last(operator)

说明：operator_properties_last

返回类型：OperatorProperties, 不会为 None

### print_undo_steps()

说明：print_undo_steps

### classmethodtag_script_reload()

说明：脚本重新加载后刷新界面的标签

### popover()

全名：popover(draw_func, \*, ui_units_x=0, keymap=None, from_active_button=False)

说明：popup_menu(draw_func, \*, title='', icon='NONE')

参数：icon，字符串，[系统图标](https://www.yuelili.com/blender-development-use-blenders-own-icon/)列表里任选一个

说明：直接弹出式菜单，不需要注册菜单类。

::: tip
请注意，它们不会阻止脚本执行，所以没法等待用户输入。
:::

```python
import bpy


def draw(self, context):
    self.layout.label(text="Hello World")


bpy.context.window_manager.popup_menu(draw, title="Greeting", icon='INFO')

```

### popup_menu_pie()

全名：popup_menu_pie(event, draw_func, \*, title='', icon='NONE')

说明：classmethodbl_rna_get_subclass(id, default=None)

说明：Parameters：d (string) – The RNA type identifier.

返回：RNA 类型，未找到则为默认

返回类型：bpy.types.Struct subclass

### classmethodbl_rna_get_subclass_py()

全名：classmethodbl_rna_get_subclass_py(id, default=None)

说明：参数：id (string) – The RNA type identifier.

返回：类，未找到则为默认

返回类型：type

### draw_cursor_add()

说明：Undocumented, consider contributing.

### draw_cursor_remove()

说明：Undocumented, consider contributing.
