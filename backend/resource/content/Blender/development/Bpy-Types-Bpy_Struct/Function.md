---
order: 6
title: Function
---

# Function

## Description

[Function(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Function.html)

基类：bpy_struct

说明：RNA 函数定义

### description 描述

说明：Description of the Function’s purpose

类型：string, default “”, (readonly，可选)

### identifier ID

说明：Unique name used in the code and scripting

类型：string, default “”, (readonly，可选)

### is_registered 判断注册

说明：Function is registered as callback as part of type registration

类型：boolean, default False, 只读

### is_registered_optional

说明：Function is 字符串 ly registered as callback part of type registration

类型：boolean, default False, 只读

### parameters 参数

说明：Parameters for the function

类型：bpy_prop_collection of Property, 只读

### use_self

说明：Function does not pass its self as an argument (becomes a static method in
python)

类型：boolean, default False, 只读

### use_self_type

说明：Function passes its self type as an argument (becomes a class method inpython if use_self is false)

类型：boolean, default False, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id(字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
