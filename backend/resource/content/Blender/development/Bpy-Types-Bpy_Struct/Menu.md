---
title: Menu
order: 8
---

# Menu

## Description

[bpy.types.MenuMenu(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Menu.html)

基类： bpy_struct

说明：菜单

### 基础菜单示例

这里个简单的菜单示例。菜单（Menu）与面板（Panle）不同，它们必须从标题、面板或其他菜单引用。

注意：Menu.bl*idname 中的'CATEGORY_MT_name'，这是菜单命名惯例。（大写分类\_MT*小写菜单功能）

::: tip
Menu 子类在调用前必须注册
:::

::: tip
菜单拥有自己的`Layout.operator_context` 作为‘EXEC_REGION_WIN’初始化，而不是‘INVOKE_DEFAULT’
:::

(参见： [Execution Context](https://docs.blender.org/api/master/bpy.ops.html#operator-execution-context)).
如果操作项上下文需要来自[`Operator.invoke`](https://docs.blender.org/api/master/bpy.types.Operator.html#bpy.types.Operator.invoke"bpy.types.Operator.invoke") 的函数输入, 这需要明确设置

![](https://cdn.yuelili.com/20220115215216.png)

```python
import bpy


class BasicMenu(bpy.types.Menu):
    bl_idname = "OBJECT_MT_select_test"
    bl_label = "Select"

    def draw(self, context):
        layout = self.layout

        layout.operator("object.select_all", text="Select/Deselect All").action = 'TOGGLE'
        layout.operator("object.select_all", text="Inverse").action = 'INVERT'
        layout.operator("object.select_random", text="Random")


# 注册
bpy.utils.register_class(BasicMenu)

# 直接使用
bpy.ops.wm.call_menu(name="OBJECT_MT_select_test")
```

### 子菜单

这种菜单展示了不同的函数

![](https://cdn.yuelili.com/20220115215116.png)

```python
import bpy


class SubMenu(bpy.types.Menu):
    bl_idname = "OBJECT_MT_select_submenu"
    bl_label = "Select"

    def draw(self, context):
        layout = self.layout

        # 1.加入3个自带的操作项
        layout.operator("object.select_all", text="Select/Deselect All").action = 'TOGGLE'
        layout.operator("object.select_all", text="Inverse").action = 'INVERT'
        layout.operator("object.select_random", text="Random")

        # 2.作为子菜单访问操作项
        layout.operator_menu_enum("object.select_by_type", "type", text="Select All by Type...")

        # 3.插入分割符
        layout.separator()

        # 4.每个操作项逐个添加到该菜单
        layout.operator_enum("object.light_add", "type")

        # 5.插入分割符
        layout.separator()

        # 6.插入现有菜单
        layout.menu("VIEW3D_MT_transform")


bpy.utils.register_class(SubMenu)

# 立即调用它
bpy.ops.wm.call_menu(name="OBJECT_MT_select_submenu")

```

### 扩展菜单

在为插件创建菜单时，不能在 Blender 的默认脚本中引用菜单。相反，插件可以将菜单项添加到现有的菜单中。

函数 menu_draw 的作用类似于 Menu.draw。

![](https://cdn.yuelili.com/20220115215444.png)

```python
import bpy


def menu_draw(self, context):
    self.layout.operator("wm.save_homefile")

# 在菜单 - 文件 最下面加入系统的操作项
bpy.types.TOPBAR_MT_file.append(menu_draw)
```

### 预设菜单

预设菜单只是一个惯例，它使用一个菜单子类来执行管理预设的普通任务。

此例展示如何添加一个预设菜单。

此例使用了“对象显示”选项，你也可以使用自己的自定义属性。

![](https://cdn.yuelili.com/20220115225437.png)

```python
import bpy
from bpy.types import Operator, Menu
from bl_operators.presets import AddPresetBase


class OBJECT_MT_display_presets(Menu):
    bl_label = "对象显示预设"
    preset_subdir = "object/display"
    preset_operator = "script.execute_preset"
    draw = Menu.draw_preset


class AddPresetObjectDisplay(AddPresetBase, Operator):
    '''添加一个对象显示预设'''
    bl_idname = "camera.object_display_preset_add"
    bl_label = "Add Object Display Preset"
    preset_menu = "OBJECT_MT_display_presets"

    # 用于所有预设值的变量
    preset_defines = [
        "obj = bpy.context.object"
    ]

    # 在预设中存储的属性
    preset_values = [
        "obj.display_type",
        "obj.show_bounds",
        "obj.display_bounds_type",
        "obj.show_name",
        "obj.show_axis",
        "obj.show_wire",
    ]

    # 在哪里存储预设
    preset_subdir = "object/display"


# 在已存在的面板中显示
def panel_func(self, context):
    layout = self.layout

    row = layout.row(align=True)
    row.menu(OBJECT_MT_display_presets.__name__, text=OBJECT_MT_display_presets.bl_label)
    row.operator(AddPresetObjectDisplay.bl_idname, text="", icon='ZOOM_IN')
    row.operator(AddPresetObjectDisplay.bl_idname, text="", icon='ZOOM_OUT').remove_active = True


classes = (
    OBJECT_MT_display_presets,
    AddPresetObjectDisplay,
)


# 注册
def register():
    for cls in classes:
        bpy.utils.register_class(cls)
    bpy.types.OBJECT_PT_display.prepend(panel_func)


def unregister():
    for cls in classes:
        bpy.utils.unregister_class(cls)
    bpy.types.OBJECT_PT_display.remove(panel_func)


if __name__ == "__main__":
    register()
```

### 拓展按钮的上下文菜单

此例可以在右键菜单中插入自己的菜单条目，当你把鼠标悬停在一个值、颜色、字符串等地方时，该菜单就会显示。

必须先选择一个对象，然后右击一个用户界面元素（也许是材料属性中的一个颜色），选择执行自定义操作。

执行：打印所有值。

在对象属性面板的一个值上右键，可以看到增加了上下文菜单

![](https://cdn.yuelili.com/20220115233101.png)

```python
import bpy
from bpy.types import Menu


def dump(obj, text):
    for attr in dir(obj):
        print("%r.%s = %s" % (obj, attr, getattr(obj, attr)))


class WM_OT_button_context_test(bpy.types.Operator):
    """ 右键测试 """
    bl_idname = "wm.button_context_test"
    bl_label = "测试上下文"

    @classmethod
    def poll(cls, context):
        return context.active_object is not None

    def execute(self, context):
        value = getattr(context, "button_pointer", None)
        if value is not None:
            dump(value, "button_pointer")

        value = getattr(context, "button_prop", None)
        if value is not None:
            dump(value, "button_prop")

        value = getattr(context, "button_operator", None)
        if value is not None:
            dump(value, "button_operator")

        return {'FINISHED'}


# 这个类必须这样准确命名，以便在右键菜单中插入一个条目
class WM_MT_button_context(Menu):
    bl_label = "Unused"

    def draw(self, context):
        pass


def menu_func(self, context):
    layout = self.layout
    layout.separator()
    layout.operator(WM_OT_button_context_test.bl_idname)


classes = (
    WM_OT_button_context_test,
    WM_MT_button_context,
)


def register():
    for cls in classes:
        bpy.utils.register_class(cls)
    bpy.types.WM_MT_button_context.append(menu_func)


def unregister():
    for cls in classes:
        bpy.utils.unregister_class(cls)
    bpy.types.WM_MT_button_context.remove(menu_func)


if __name__ == "__main__":
    register()


```

### bl_description

说明：菜单的描述

类型：字符串，默认为""

### bl_idname

说明：菜单 ID，如果设置则使用自定义 ID，否则根据类名（例如类名是"OBJECT_MT_hello"，bl_idname 为"OBJECT_MT_hello"）

类型：字符串，默认为"", 不会为 None

### bl_label

说明：菜单的标签

类型：string, default “”, 不会为 None

### bl_owner_id

说明：

类型：字符串，默认为"", 不会为 None

### bl_translation_context

说明：

类型：字符串，默认为"", 不会为 None

### layout

说明：Defines the structure of the menu in the UI

类型：UILayout, 只读

### classmethod poll()

全名：classmethod poll(context)

说明：如果返回不为空，则会绘制菜单

返回类型：boolean

### draw()

全名：draw(context)

说明：将 UI 元素绘制到菜单 UI 布局中

### draw_preset()

全名：draw_preset(\_context)

在子类上定义:

\- preset_operator：(字符串)

\- preset_subdir：(字符串)

可选:

\- preset_add_operator：(字符串)

\- preset_extensions：(set of strings)

\- preset_operator_defaults (dict of keyword args)

### path_menu()

全名：path_menu(searchpaths, operator, \*, props_default=None,prop_filepath='filepath', filter_ext=None, filter_path=None,display_name=None, add_operator=None)

说明：从一个路径列表填充菜单。

参数：

- searchpaths：字符串序列。 要扫描的路径。
- operator：(字符串)。 在每个文件中使用的操作项 ID。
- prop_filepath：(字符串)。 可选的操作项文件路径属性（默认为 "filepath"）。
- props_default：字典。 分配给每个操作项的属性。
- filter_ext：可回调，接收一个字符串并返回 bool。
- 可选：可回调，获取文件扩展名。
- 返回：false，将文件从列表中排除。
- display_name：Callable，接受一个字符串并返回一个字符串。 可选回调：接受完整的路径，返回要显示的名称。

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id(字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
