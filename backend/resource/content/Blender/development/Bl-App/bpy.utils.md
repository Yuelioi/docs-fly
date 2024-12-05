---
title: 实用程序 (bpy.utils)
order: 2
---

# 实用程序 (bpy.utils)

## bpy.utils

该模块用于与 BL 内部数据无关的实用功能。

### blend_paths()

全名：bpy.utils.blend_paths(absolute=False, packed=False, local=False)

说明：返回一个由加载的.blend 文件引用的外部文件的路径列表。

参数：

- absolute (布尔) ： 为 true 返回绝对路径。
- packed (boolean) ： 为 true 跳过打包数据的文件路径。
- local (boolean) ： 为 true 跳过链接库路径。

返回：路径列表。

返回类型：字符串的列表

### escape_identifier()

全名：bpy.utils.escape_identifier(string)

说明：用于动画路径的简单字符串转义函数。

参数：string (字符串)

返回：转义后的字符串。

返回类型：字符串

### flip_name()

全名：bpy.utils.flip_name(name, strip_digits=False)

说明：在左/右两边翻转一个名字，对镜像骨骼名称很有用。

参数：

- name (string) ： 要翻转的骨骼名称。
- strip_digits (bool) ： 是否去除.###的后缀。

返回：翻转后的名称。

返回类型：字符串

### unescape_identifier()

全名：bpy.utils.unescape_identifier(string)

说明：用于动画路径的简单字符串解密函数。这个函数执行 escape_identifier 的反向操作。

参数：string (字符串) ： 文本

返回：未转义的字符串。

返回类型：字符串

### register_class()

全名：bpy.utils.register_class(cls)

说明：注册一个 Blender 类的子类。

参数：cls (class) – Blender 的类。包括：

- bpy.types.Panel
- bpy.types.UIList
- bpy.types.Menu
- bpy.types.Header
- bpy.types.Operator
- bpy.types.KeyingSetInfo
- bpy.types.RenderEngine

如果该类不是一个可注册的 blender 类的子类，则 ValueError。

注意：如果该类有一个注册类方法，它将在注册前被调用。

### resource_path()

全名：bpy.utils.resource_path(type, major=bpy.app.version[0],minor=bpy.app.version[1])

说明：返回存储系统文件的基本路径。

参数：

- type (string) ： ['USER', 'LOCAL', 'SYSTEM'] 中的字符串。
- major (int) ： 主版本，默认为当前版本。
- minor (string) ： 次要版本，默认为当前版本。
- 返回：资源路径（不一定是现有的）。

返回类型：字符串

### unregister_class()

全名：bpy.utils.unregister_class(cls)

说明：从 blender 中注销 Python 类。

如果该类有一个 unregister 类方法，它将在取消注册之前被调用。

### keyconfig_init()

全名：bpy.utils.keyconfig_init()

说明：

### keyconfig_set()

全名：bpy.utils.keyconfig_set(filepath, \*, report=None)

说明：

### load_scripts()

全名：bpy.utils.load_scripts(\*, reload_scripts=False, refresh_scripts=False)

说明：加载脚本并运行每个模块的注册功能。

参数：

- reload_scripts (bool) ：所有脚本在加载前调用其注销方法。
- refresh_scripts (bool) ： 只加载尚未加载为模块的脚本。

### modules_from_path()

全名：bpy.utils.modules_from_path(path, loaded_modules)

说明：加载一个路径中的所有模块，并以列表形式返回。

参数：

- path (string) ：被扫描的路径。
- loaded_modules (set) ： 已加载的模块名，符合这些名称的文件将被忽略。

返回：所有加载的模块。

返回类型：列表

### preset_find()

全名：bpy.utils.preset_find(name, preset_path, \*, display_name=False, ext='.py')

说明：

### preset_paths()

全名：bpy.utils.preset_paths(subdir)

说明：返回一个特定预设的路径列表。

参数：subdir (string) ： 预设的子目录（必须不是一个绝对路径）。

返回：脚本路径。

返回类型：列表

### refresh_script_paths()

全名：bpy.utils.refresh_script_paths()

说明：在创建新的脚本路径后运行此程序，用于更新 sys.path

### app_template_paths()

全名：bpy.utils.app_template_paths(\*, path=None)

说明：返回有效的应用程序模板路径。

参数：path (string) ： 可选的子目录。

返回：应用程序模板路径。

返回类型：生成器

### register_manual_map()

全名：bpy.utils.register_manual_map(manual_hook)

说明：

### unregister_manual_map()

全名：bpy.utils.unregister_manual_map(manual_hook)

说明：

### register_classes_factory() 批量注册注销类

全名：bpy.utils.register_classes_factory(classes)

说明：可以批量注册和注销类。

```python
classes = (SingleVertex_OT_Operator , SingleVertex_PT_Panel)
register , unregister = bpy.utils.register_classes_factory(classes)
```

### register_submodule_factory() 注册和注销子模块

全名：bpy.utils.register_submodule_factory(module_name, submodule_names)

说明：注册和注销子模块，调用子模块的注册和注销函数。

注意：模块按照给出的顺序注册，注销则相反。

参数：

- module_name (string) ： 模块名称，通常是 **name**。
- submodule_names (字符串列表) ： 要注册和注销子模块名称列表。

返回：注册和解除注册的函数。

返回类型：元组对的函数

### register_tool()

全名：bpy.utils.register_tool(tool_cls, \*, after=None, separator=False,
group=False)

说明：在工具条上注册一个工具。

参数：

- tool (bpy.types.WorkSpaceTool subclass.) ： 一个工具子类。
- space_type (string) ： 空间类型标识符。
- after (string collection of strings or None.) ： 可选标识符，这个工具将被添加到 XX 工具后面。
- separator (bool) ：为 True 则在这个工具之前添加一个分隔符。
- group (bool) ： 为 True 则添加一个新的嵌套工具组。

### make_rna_paths()

全名：bpy.utils.make_rna_paths(struct_name, prop_name, enum_name)

说明：从给定名称创建 RNA "路径"。

参数：

- struct_name (string) ： RNA 结构的名称（例如 "Scene"）。
- prop_name (string) ： RNA 结构的属性名称。
- enum_name (string) ： RNA 枚举标识符的名称。

返回：三个
"RNA 路径"（most_complete_path、"struct.prop"、"struct.prop：'enum'"）的三联体。如果没有给出 enum_name，第三个元素将无效。

返回类型：字符串的元组

### manual_map()

全名：bpy.utils.manual_map()

说明：

### script_path_user()

全名：bpy.utils.script_path_user()

说明：返回环境变量，并返回到主目录或 Null

### script_path_pref()

全名：bpy.utils.script_path_pref()

说明：返回用户偏好或 Null

### script_paths()

全名：bpy.utils.script_paths(\*, subdir=None, user_pref=True, check_all=False,
use_user=True)

说明：返回一个有效的脚本路径列表。

参数：

- subdir (string) ： 可选，子目录。
- user_pref (bool) ： 包括用户偏好的脚本路径。
- check_all (bool) ： 包括本地、用户和系统的路径，而不是只包括 Blender 使用的路径。

返回：脚本路径。

返回类型：列表

### smpte_from_frame()

全名：bpy.utils.smpte_from_frame(frame, \*, fps=None, fps_base=None) 说明：从帧中返回一个 SMPTE 格式的字符串。HH：MM：SS：FF。如果没有给出 fps 和 fps_base，则使用当前场景。

参数：frame，帧数，整数或浮点数。

返回：帧字符串。

返回类型：字符串

### smpte_from_seconds()

全名：bpy.utils.smpte_from_seconds(time, \*, fps=None, fps_base=None) 说明：返回一个由时间组成的 SMPTE 格式字符串。HH：MM：SS：FF。如果没有给出 fps 和 fps_base，则使用当前场景。

参数：time ，时间（秒），int, float or datetime.timedelta。

返回：帧字符串。

返回类型：字符串

### unregister_tool()

全名：bpy.utils.unregister_tool(tool_cls)

说明：

### user_resource()

全名：bpy.utils.user_resource(resource_type, \*, path='', create=False)

说明：返回一个用户资源路径（通常来自用户的主目录）。

参数：

- type (string) ： 资源类型['DATAFILES', 'CONFIG', 'SCRIPTS', 'AUTOSAVE'] 。
- path (string) ： 可选的子目录。
- create (boolean) ： 将路径视为一个目录，如果它不存在则创建它。

返回：路径。

返回类型：字符串

### execfile()

全名：bpy.utils.execfile(filepath, \*, mod=None)

说明：将一个文件路径作为 Python 脚本执行。

参数：

- filepath (string) ： 要执行的脚本的路径。
- mod (Module or None) ： 可选的缓存模块，是之前执行的结果。

返回：可以作为 mod 传回的模块。

返回类型：模块类型

## bpy.utils 子模块 (bpy.utils.previews)

此模块包含处理自定义预览的实用功能。

一个高级的 "缓存 "预览管理器。

允许脚本生成他们自己的预览，并在 UI 部件中使用它们作为图标（UILayout 函数的'icon_value'）。

### 自定义图标示例

报错，日后再试

```python
# 本脚展示如何在一个按钮或菜单项放置一个自定义图标。
#
# 重要提示：如果运行这个示例，按钮上不会有图标。
# 需要替换自己的图片路径。
# 对于可分发的脚本，建议将图标放在 addon 文件夹中
# 并相使用相对路径访问，保证可移植性。
#
#
# UI-previews的其他用例。
# - 提供一个固定的预览列表供用户选择
# - 提供一个动态的预览列表（例如，通过读取目录来计算）。
#
# 对于上述用例，见模板 "ui_previews_dynamic_enum.py"


import os
import bpy


class PreviewsExamplePanel(bpy.types.Panel):
    """在对象属性窗口中创建一个面板"""
    bl_label = "Previews Example Panel"
    bl_idname = "OBJECT_PT_previews"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        pcoll = preview_collections["main"]

        row = layout.row()
        my_icon = pcoll["my_icon"]
        row.operator("render.render", icon_value=my_icon.icon_id)

        # my_icon.icon_id可以在任何接受icon_value的UI函数中使用
        # 也可以尝试设置text=""来获得一个只有图标的操作按钮


# 可以在这里存储多个预览集合。
# 本例只存储 "main"。
preview_collections = {}


def register():

    # 注意，由bpy.utils.previews返回的预览集合是普通的py对象
    # --你可以用它们来存储自定义数据。
    import bpy.utils.previews
    pcoll = bpy.utils.previews.new()

    # 图标所在的文件夹的使用相对路径计算
    my_icons_dir = os.path.join(os.path.dirname(__file__), "icons")

    # 加载一个文件的预览缩略图，并存储在预览集合中
    pcoll.load("my_icon", os.path.join(my_icons_dir, "icon-image.png"), 'IMAGE')

    preview_collections["main"] = pcoll

    bpy.utils.register_class(PreviewsExamplePanel)


def unregister():

    for pcoll in preview_collections.values():
        bpy.utils.previews.remove(pcoll)
    preview_collections.clear()

    bpy.utils.unregister_class(PreviewsExamplePanel)


if __name__ == "__main__":
    register()
```

### previews.new()

全名：bpy.utils.previews.new()

返回：一个新的预览集合。

返回类型：ImagePreviewCollection

### previews.remove()

全名：bpy.utils.previews.remove(pcoll)

移除指定的预览集合。

参数：pcoll (ImagePreviewCollection) - 要关闭的预览集合。

### previews.ImagePreviewCollection

classbpy.utils.previews.ImagePreviewCollection

类似于字典的预览类。

这是 Python 内置的 dict 类型的一个子类，用于存储多个图像预览。

## bpy.utils 子模块 (bpy.utils.units)

### units.categories

全名：bpy.utils.units.categories

常量： bpy.utils.units.categories(NONE=’NONE’, LENGTH=’LENGTH’, AREA=’AREA’,VOLUME=’VOLUME’, MASS=’MASS’, ROTATION=’ROTATION’, TIME=’TIME’,VELOCITY=’VELOCITY’, ACCELERATION=’ACCELERATION’, CAMERA=’CAMERA’,
POWER=’POWER’)

### units.systems

全名：bpy.utils.units.systems

常量： bpy.utils.units.systems(NONE=’NONE’, METRIC=’METRIC’, IMPERIAL=’IMPERIAL’)

### units.to_string()

全名：bpy.utils.units.to_string(unit_system, unit_category, value, precision=3,split_unit=False, compatible_unit=False)

说明：将浮点数转成一个带单位的字符串。

参数

- unit_system (string) ：单位，来自 bpy.utils.units.systems。
- unit_category (string) ： 要转换的数据类别（长度、面积、旋转，等等），来自 bpy.utils.units.categories。
- value (float)：要转换为字符串的值。
- precision (int)： 逗号后的位数。
- split_unit (bool) ：是否使用多个单位（1m1cm），或者始终只使用一个单位（1.01m）。
- compatible compatible_unit (bool) ： 是否使用键盘友好的单位（1m2）或更好的 utf-8 单位（1m²）。

返回：转换后的字符串。

返回类型：str

Raises：如果不能转为有效的 python float 值，则 ValueError 。

### units.to_value

全名：bpy.utils.units.to_value(unit_system, unit_category, str_input,
str_ref_unit=None)

说明：将给字符串转换成浮点值。

参数

- unit_system (string) ：单位系统，来自 bpy.utils.units.systems。
- unit_category (string)：要转换的数据类别（长度、面积、旋转，等等），来自 bpy.utils.units.categories。
- str_input (string) ：要转换为浮点数的字符串。
- str_ref_unit (string or None) -： 一个参考字符串，如果在 str_input 中没有找到，可以从中提取一个默认单位。

返回：转换/解析后的值。

返回类型：浮点数

Raises：如果不能转为有效的 python float 值，则 ValueError 。
