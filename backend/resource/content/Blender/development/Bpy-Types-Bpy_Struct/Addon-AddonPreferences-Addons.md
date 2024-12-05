---
title: Addon AddonPreferences Addons
order: 1
---

# Addon AddonPreferences Addons

## Description

[Addon(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Addon.html)

[AddonPreferences(bpy_struct)](https://docs.blender.org/api/master/bpy.types.AddonPreferences.html)

[Addons(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Addons.html)

## Addon(bpy_struct)

Python 插件自动加载设置

### module 模块

说明：模块名称

类型：字符串，默认为""，不会为 None

### preferences 首选项

类型：AddonPreferences, (readonly)

### classmethod bl_rna_get_subclass(id, default=None)

参数：id，字符串，The RNA type identifier.

返回：RNA 属性，未找到则使用默认

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py(id, default=None)

参数：id，字符串，RNA 属性 ID

返回：类，未找到则用默认

返回类型：type

## AddonPreferences(bpy_struct)

### 示例

![](https://cdn.yuelili.com/20220119225017.png)

```python
bl_info = {
    "name": "示例插件首选项",
    "author": "YLL",
    "version": (1, 0),
    "blender": (2, 65, 0),
    "location": "SpaceBar Search -> Add-on Preferences Example",
    "description": "示例插件",
    "warning": "",
    "doc_url": "",
    "tracker_url": "",
    "category": "Object", # 分类
}


import bpy
from bpy.types import Operator, AddonPreferences
from bpy.props import StringProperty, IntProperty, BoolProperty


class ExampleAddonPreferences(AddonPreferences):
    # 必须匹配插件名称
    # 作为python模块的子模块时，需要使用'__package__
    bl_idname = __name__

    filepath: StringProperty(
        name="Example File Path",
        subtype='FILE_PATH',
    )
    number: IntProperty(
        name="Example Number",
        default=4,
    )
    boolean: BoolProperty(
        name="Example Boolean",
        default=False,
    )

    def draw(self, context):
        layout = self.layout
        layout.label(text="This is a preferences view for our add-on")
        layout.prop(self, "filepath")
        layout.prop(self, "number")
        layout.prop(self, "boolean")


class OBJECT_OT_addon_prefs_example(Operator):
    """显示示例首选项"""
    bl_idname = "object.addon_prefs_example"
    bl_label = "插件首选项label"
    bl_options = {'REGISTER', 'UNDO'}

    def execute(self, context):
        preferences = context.preferences
        addon_prefs = preferences.addons[__name__].preferences

        info = ("Path: %s, Number: %d, Boolean %r" %
                (addon_prefs.filepath, addon_prefs.number, addon_prefs.boolean))

        self.report({'INFO'}, info)
        print(info)

        return {'FINISHED'}


# 注册/注销
def register():
    bpy.utils.register_class(OBJECT_OT_addon_prefs_example)
    bpy.utils.register_class(ExampleAddonPreferences)


def unregister():
    bpy.utils.unregister_class(OBJECT_OT_addon_prefs_example)
    bpy.utils.unregister_class(ExampleAddonPreferences)
```

### bl_idname

类型：字符串，默认为""， 不会为 None

### classmethod bl_rna_get_subclass(id, default=None)

参数：id，字符串，RNA 属性 ID

返回：RNA 类型，未找到则用默认

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py(id, default=None)

参数：id，字符串，RNA 属性 ID

返回：类，未找到则用默认

返回类型：type

## Addons(bpy_struct)

插件集合

### classmethod new()

说明：添加一个新插件

返回：插件数据

返回类型：Addon

### classmethodremove(addon)

说明：移除插件

参数：addon (Addon, 不会为 None) ，要移除的插件

### classmethod bl_rna_get_subclass(id, default=None)

参数：id，字符串，RNA 属性 ID

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py(id, default=None)

参数：id，字符串，RNA 属性 ID

返回：类，未找到则为默认。

返回类型：type
