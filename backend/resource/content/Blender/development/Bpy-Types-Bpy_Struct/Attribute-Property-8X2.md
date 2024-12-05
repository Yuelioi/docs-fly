---
title: Attribute Property + 8X2个属性子类
order: 3
---

# Attribute Property + 8X2个属性子类

## Description

**Attribute 属性（几何属性）**

[Attribute(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Attribute.html)

[AttributeGroup(bpy_struct)](https://docs.blender.org/api/master/bpy.types.AttributeGroup.html)

**Property 属性（RNA 属性）**

[Property(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Property.html)

[PropertyGroup(bpy_struct)](https://docs.blender.org/api/master/bpy.types.PropertyGroup.html)

[PropertyGroupItem(bpy_struct)](https://docs.blender.org/api/master/bpy.types.PropertyGroupItem.html)

**基本属性(Attribute)**

[BoolAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.BoolAttribute.html#bpy.types.BoolAttribute)

[ByteColorAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.ByteColorAttribute.html#bpy.types.ByteColorAttribute)

[Float2Attribute(Attribute)](https://docs.blender.org/api/master/bpy.types.Float2Attribute.html#bpy.types.Float2Attribute)

[FloatAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.FloatAttribute.html#bpy.types.FloatAttribute)

[FloatColorAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.FloatColorAttribute.html#bpy.types.FloatColorAttribute)

[FloatVectorAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.FloatVectorAttribute.html#bpy.types.FloatVectorAttribute)

[IntAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.IntAttribute.html#bpy.types.IntAttribute)

[StringAttribute(Attribute)](https://docs.blender.org/api/master/bpy.types.StringAttribute.html#bpy.types.StringAttribute)

**基本属性(Property)**

[BoolProperty(Property)](https://docs.blender.org/api/master/bpy.types.BoolProperty.html)

[CollectionProperty(Property)](https://docs.blender.org/api/master/bpy.types.CollectionProperty.html)

[EnumProperty(Property)](https://docs.blender.org/api/master/bpy.types.EnumProperty.html)

[FloatProperty(Property)](https://docs.blender.org/api/master/bpy.types.FloatProperty.html)

[IntProperty(Property)](https://docs.blender.org/api/master/bpy.types.IntProperty.html)

[PointerProperty(Property)](https://docs.blender.org/api/master/bpy.types.PointerProperty.html)

[StringProperty(Property)](https://docs.blender.org/api/master/bpy.types.StringProperty.html)

[Struct(bpy_struct)](https://docs.blender.org/api/master/bpy.types.Struct.html)

[EnumPropertyItem(bpy_struct)](https://docs.blender.org/api/master/bpy.types.EnumPropertyItem.html)

**图示：**

![](https://cdn.yuelili.com/20220112013129.png)

## Attribute(bpy_struct)

几何属性

### data_type

说明：储存在属性中的数据类型

- FLOAT，Float - 浮点值。
- INT，Integer - 32 位整数。
- FLOAT_VECTOR，Vector - 具有浮点值的 3D 向量。
- FLOAT_COLOR，Color - 具有浮点值的 RGBA 颜色。
- BYTE_COLOR，Byte Color - 具有 8 位数值的 RGBA 颜色。
- STRING，String - 文本字符串。
- BOOLEAN，Boolean - 真或假。
- FLOAT2，2D Vector - 具有浮点值的 2D 向量。

类型：以上之一，默认为‘FLOAT’, 只读

### domain

说明：属性作用范围

POINT：Point，点的属性

EDGE：Edge，网格边缘的属性

FACE：Face，网格面的属性

CORNER：Face Corner，网格面角属性

CURVE：Spline，spline 属性

INSTANCE：Instance，实例属性

类型：以上之一，默认为‘POINT’, 只读

### name

说明：属性名称

类型：字符串，默认为"", 不会为 None

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) – The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) – The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## AttributeGroup(bpy_struct)

Group of geometry attributes

### active

说明：激活属性

类型：Attribute

### active_index

说明：

类型：int in [-inf, inf], default 0

### new()

全名：new(name, type, domain)

说明：添加一个属性

参数：name (string, 不会为 None) 。 属性名

type：其中之一

- FLOAT，Float - 浮点值。
- INT，Integer - 32 位整数。
- FLOAT_VECTOR，Vector - 具有浮点值的 3D 向量。
- FLOAT_COLOR，Color - 具有浮点值的 RGBA 颜色。
- BYTE_COLOR，Byte Color - 具有 8 位数值的 RGBA 颜色。
- STRING，String - 文本字符串。
- BOOLEAN，Boolean - 真或假。
- FLOAT2，2D Vector - 具有浮点值的 2D 向量。

domain

- Domain：属性所存储的元素的类型。
- POINT：点上的属性。
- EDGE：网格边缘的属性。
- FACE：网格面的属性。
- CORNER：网格面角的属性。
- CURVE Spline：CURVE Spline 的属性。
- INSTANCE：实例的属性。

返回：新的几何属性

返回类型：Attribute

### remove()

全名：remove(attribute)

说明：移除属性

参数：attribute (Attribute, 不会为 None) 。 Geometry Attribute

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## Property(bpy_struct)

基类： bpy_struct

子类：BoolProperty, CollectionProperty, EnumProperty, FloatProperty, IntProperty,PointerProperty, StringProperty

介绍：RNA 属性定义

### description 属性描述

说明：此属性的帮助描述

类型：字符串，默认为""，只读，可选

### icon 图标

说明：项目图标

icon：[系统图标](https://www.yuelili.com/blender-development-use-blenders-own-icon/)列表里任选一个

### identifier 标识符

说明：在代码和脚本中使用的唯一名称

类型：字符串，默认为""，只读，可选

### is_animatable

说明：Property is animatable through RNA

类型：boolean, default False, 只读

### is_argument_optional

说明：在实现 RNA 函数的 Python 函数中，当属性为可选时为 True。

类型：boolean, default False, 只读

### is_enum_flag

说明：True when multiple enums

类型：boolean, default False, 只读

### is_hidden

说明：当属性隐藏时为 True

类型：boolean, default False, 只读

### is_library_editable

说明：Property is editable from linked instances (changes not saved)

类型：boolean, default False, 只读

### is_never_none

说明：当属性不能设置为 None 时为 True

类型：boolean, default False, 只读

### is_output

说明：True when this property is an output value from an RNA function

类型：boolean, default False, 只读

### is_overridable

说明：Property is overridable through RNA

类型：boolean, default False, 只读

### is_readonly

说明：属性可通过 RNA 进行编辑

类型：boolean, default False, 只读

### is_registered

说明：Property is registered as part of type registration

类型：boolean, default False, 只读

### is_registered_optional

说明：Property is 可选 ly registered as part of type registration

类型：boolean, default False, 只读

### is_required

说明：False when this property is an 可选 argument in an RNA function

类型：boolean, default False, 只读

### is_runtime

说明：Property has been dynamically created at runtime

类型：boolean, default False, 只读

### is_skip_save

说明：True when the property is not saved in presets

类型：boolean, default False, 只读

### name

说明：Human readable name

类型：字符串，默认为""，只读，可选

### srna

说明：Struct definition used for properties assigned to this item

类型：Struct, 只读

### subtype 子类型

说明：属性的语义解释

类型：下面之一，默认为 NONE， 只读

- NONE：无
- FILEPATH：文件路径（貌似是 FILE_PATH）
- DIRPATH：目录路径（貌似是 DIR_PATH）
- FILENAME：文件名称（貌似是 FILE_NAME）
- BYTESTRING：Byte String.
- PASSWORD：密码 。(**\*\*\*\***).
- PIXEL：像素
- UNSIGNED：无符号
- PERCENTAGE：百分比
- FACTOR：系数（因子）
- ANGLE：角度
- TIME：时间 (场景相对) 。 以帧为单位，根据场景可以转为秒
- TIME_ABSOLUTE：时间 (场景绝对) 。 以秒为单位。
- DISTANCE：距离
- DISTANCE_CAMERA：摄像机距离
- POWER：Power.
- TEMPERATURE：温度
- COLOR：颜色
- TRANSLATION：翻译
- DIRECTION：方向
- VELOCITY：速度
- ACCELERATION：加速度
- MATRIX：矩阵
- EULER：欧拉角
- QUATERNION：四元数
- AXISANGLE：轴-角
- XYZ：XYZ.
- XYZ_LENGTH：XYZ Length.
- COLOR_GAMM： Color.
- COORDS：坐标
- LAYER：图层
- LAYER_MEMBER：图层成员

### tags

说明：Subset of tags (defined in parent struct) that are set for this property

类型：在字典中间枚举, 默认为{}, 只读

### translation_context

说明：该属性名称的翻译上下文

类型：string, default “”, (只读，可选)

### type

说明：属性的数据类型

类型：默认为"BOOLEAN"，只读

- BOOLEAN
- INT
- FLOAT
- STRING
- ENUM
- POINTER
- COLLECTION

### unit

说明：Type of units for this property

类型：enum in [‘NONE’, ‘LENGTH’, ‘AREA’, ‘VOLUME’, ‘ROTATION’, ‘TIME’,‘TIME_ABSOLUTE’, ‘VELOCITY’, ‘ACCELERATION’, ‘MASS’, ‘CAMERA’, ‘POWER’,‘TEMPERATURE’], default ‘NONE’, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## PropertyGroup(bpy_struct)

基类： bpy_struct

子类：AssetHandle, OperatorFileListElement, OperatorMousePath,OperatorStrokeElement, SelectedUvElement

说明：PropertyGroups 是动态定义的属性集的基类。

可以用自己的类型来扩展现有的 Blender 数据，这些类型可以被动画化，也可以从用户界面和 python 中访问。

**注意：**

1.属性组（PropertyGroups）在分配给 Blender 数据前，必须注册

2.分配给 blender 数据的值保存在磁盘上，但类的定义却没有，所以每当加载 blender 时，类也需要注册。最好插件在启动时加载并注册你的属性。

3.Property 类型都是在 bpy.props 中进行申明

**人话（个人理解）：**

自定义的类，继承 PropertyGroup 后，里面所有的 RNA 定义都可以直接作为属性传入。再用一个变量接收，即可完成多个属性传递。

### 示例

```python
import bpy # 创建（类）属性 class
MyPropertyGroup(bpy.types.PropertyGroup): custom_1:bpy.props.FloatProperty(name="My Float") # 注意RNA属性用 : 进行定义 custom_2:bpy.props.IntProperty(name="My Int") # 注册类bpy.utils.register_class(MyPropertyGroup) # 对Object类，添加一个参数bpy.types.Object.my_prop_grp = bpy.props.PointerProperty(type=MyPropertyGroup)# 给一个object添加自定义属性 bpy.data.objects[0].my_prop_grp.custom_1 = 22.0
```

### classbpy.types.PropertyGroup()

全名：classbpy.types.PropertyGroup(bpy_struct)

说明：Group of ID properties

### name

说明：Unique name used in the code and scripting

类型：string, default “”, 不会为 None

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## PropertyGroupItem(bpy_struct)

基类： bpy_struct

说明：储存任意、用户定义的属性

### collection

说明：

类型：bpy_prop_collection of PropertyGroup, 只读

### double

说明：

类型：float in [-inf, inf], default 0.0

### double_array

说明：

类型：float array of 1 items in [-inf, inf], default (0.0)

### float

说明：

类型：float in [-inf, inf], default 0.0

### float_array

说明：

类型：float array of 1 items in [-inf, inf], default (0.0)

### group

说明：类型：PropertyGroup, 只读

### id

说明：

类型：ID

### idp_array

说明：

类型：bpy_prop_collection of PropertyGroup, 只读

### int

说明：类型：int in [-inf, inf], default 0

### int_array

说明：

类型：int array of 1 items in [-inf, inf], default (0,)

### string

说明：

类型：string, default “”, 不会为 None

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## BoolAttribute(Attribute)

基类： bpy_struct, Attribute

说明：几何属性布尔值

### data

说明：  
类型：bpy_prop_collection of BoolAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## ByteColorAttribute(Attribute)

基类： bpy_struct, Attribute

Color geometry attribute, with 8-bit values

### data

说明：  
类型：bpy_prop_collection of ByteColorAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## Float2Attribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.Float2Attribute()

全名：classbpy.types.Float2Attribute(Attribute)

2D vector geometry attribute, with floating-point values

### data

说明：  
类型：bpy_prop_collection of Float2AttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## FloatAttribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.FloatAttribute()

全名：classbpy.types.FloatAttribute(Attribute)

Geometry attribute with floating-point values

### data

说明：  
类型：bpy_prop_collection of FloatAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## FloatColorAttribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.FloatColorAttribute()

全名：classbpy.types.FloatColorAttribute(Attribute)

Color geometry attribute, with floating-point values

### data

说明：  
类型：bpy_prop_collection of FloatColorAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## FloatVectorAttribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.FloatVectorAttribute()

全名：classbpy.types.FloatVectorAttribute(Attribute)

Vector geometry attribute, with floating-point values

### data

说明：  
类型：bpy_prop_collection of FloatVectorAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## IntAttribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.IntAttribute()

全名：classbpy.types.IntAttribute(Attribute)

Integer geometry attribute

### data

说明：  
类型：bpy_prop_collection of IntAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## StringAttribute(Attribute)

基类： bpy_struct, Attribute

### classbpy.types.StringAttribute()

全名：classbpy.types.StringAttribute(Attribute)

String geometry attribute

### data

说明：  
类型：bpy_prop_collection of StringAttributeValue, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## BoolProperty(Property)

基类：bpy_struct, Property

说明：定义 RNA 的布尔属性

### 示例

name、description、default 都是继承于父类 Property

```python
my_bool: bpy.props.BoolProperty( name="开/关", description="我是布尔类型",
default = False )
```

### array_dimensions

说明：Length of each dimension of the array

类型：int array of 3 items in [0, inf], default (0, 0, 0), 只读

### array_length

说明：Maximum length of the array, 0 means unlimited

类型：int in [0, inf], default 0, 只读

### default

说明：Default value for this number

类型：boolean, default False, 只读

### default_array

说明：Default value for this array

类型：boolean array of 3 items, default (False, False, False), 只读

### is_array

说明：类型：boolean, default False, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## CollectionProperty(Property)

基类：bpy_struct, Property

说明：RNA 集合属性用于定义 lists, arrays and mappings

### fixed_type

说明：Fixed pointer type, empty if variable type

类型：Struct, 只读

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

---

### 示例

```python
from bpy.types import PropertyGroup # 定义 class
ListItem(PropertyGroup): """代表列表中一个项目的属性组""" prop1: StringProperty(name="propp1", description="一个属性", default="属性1") prop2: StringProperty(name="propp2", description="另一个属性", default="属性2") # 使用bpy.types.Scene.my_list = CollectionProperty(type=ListItem) # type继承父类Property？ class XX() def execute(self, context): my_list =context.scene.my_list # 从scene中获取my_lsit my_list.add() # 添加项 index =context.scene.list_index my_list.remove(index) # 删除项 my_list.move(index + 1,index) # 移动项，示例为向下移动1 my_list.clear() #清空项 return{'FINISHED'}
```

以下非官方文档，为查看其它示例以及个人找到的。

![](https://cdn.yuelili.com/20220117000319.png)

### add()

说明：添加项

### clear()

说明：清空所有项

### move()

全名：XX list.move(trg , src)

说明：移动项

参数：

- trg：目标索引
- src：被移动的索引

### remove()

说明：移除项

参数：指定一个索引，没啥好说的

## EnumProperty(Property)

基类：bpy_struct, Property

说明：RNA 枚举属性定义，用于从预设选项中选取

### 示例

        my_enum: EnumProperty(

            name="Dropdown:",

            description="Apply Data to attribute.",

            items=[ ('OP1', "选项 1", ""),

                    ('OP2', "选项 2", ""),

                    ('OP3', "选项 3", ""),

                   ]

            )

```


     




    ## default #




    说明：Default value for this enum




    类型：enum in [‘DUMMY’], default ‘DUMMY’, 只读




    ## default_flag #




    说明：Default value for this enum




    类型：enum set in {‘DUMMY’}, default {}, 只读




    ## enum_items #




    说明：Possible values for the property




    类型：bpy_prop_collection of EnumPropertyItem, 只读




    ## enum_items_static #




    说明：Possible values for the property (never calls 可选 dynamic generation of those)




    类型：bpy_prop_collection of EnumPropertyItem, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




    # EnumPropertyItem(bpy_struct) #




    基类： bpy_struct




    说明：定义在RNA枚举属性中的一个选项




    ## description #




    说明：Description of the item’s purpose




    类型：字符串，默认为""，只读，可选




    ## icon #




    说明：Icon of the item




    类型：Icon图标




    ## identifier #




    说明：Unique name used in the code and scripting




    类型：字符串，默认为""，只读，可选




    ## name #




    说明：Human readable name




    类型：string, default “”, (只读，可选)




    ## value #




    说明：Value of the item




    类型：int in [0, inf], default 0, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




    # FloatProperty(Property) #




    基类：bpy_struct, Property




    说明：RNA floating-point number (single precision) property definition




    ## 示例 #





        my_float: FloatProperty(

            name = "浮点值",

            description = "我是浮点值",

            default = 23.7,

            min = 0.01,

            max = 30.0

            )


```

    ## array_dimensions #




    说明：Length of each dimension of the array




    类型：int array of 3 items in [0, inf], default (0, 0, 0), 只读




    ## array_length #




    说明：Maximum length of the array, 0 means unlimited




    类型：int in [0, inf], default 0, 只读




    ## default #




    说明：Default value for this number




    类型：float in [-inf, inf], default 0.0, 只读




    ## default_array #




    说明：Default value for this array




    类型：float array of 3 items in [-inf, inf], default (0.0, 0.0, 0.0), 只读




    ## hard_max #




    说明：Maximum value used by buttons




    类型：float in [-inf, inf], default 0.0, 只读




    ## hard_min #




    说明：Minimum value used by buttons




    类型：float in [-inf, inf], default 0.0, 只读




    ## is_array #




    说明：类型：boolean, default False, 只读




    ## precision #




    说明：Number of digits after the dot used by buttons




    类型：int in [0, inf], default 0, 只读




    ## soft_max #




    说明：Maximum value used by buttons




    类型：float in [-inf, inf], default 0.0, 只读




    ## soft_min #




    说明：Minimum value used by buttons




    类型：float in [-inf, inf], default 0.0, 只读




    ## step #




    说明：Step size used by number buttons, for floats 1/100th of the step size




    类型：float in [0, inf], default 0.0, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




    # IntProperty(Property) #




    基类：bpy_struct, Property




    说明：RNA integer number property definition




    ## 示例 #


    ```python
    my_int: IntProperty(

            name = "整数类型",

            description="我是整型",

            default = 23,

            min = 10,

            max = 100

            )

````


     




    ## array_dimensions #




    说明：Length of each dimension of the array




    类型：int array of 3 items in [0, inf], default (0, 0, 0), 只读




    ## array_length #




    说明：Maximum length of the array, 0 means unlimited




    类型：int in [0, inf], default 0, 只读




    ## default #




    说明：Default value for this number




    类型：int in [-inf, inf], default 0, 只读




    ## default_array #




    说明：Default value for this array




    类型：int array of 3 items in [-inf, inf], default (0, 0, 0), 只读




    ## hard_max #




    说明：Maximum value used by buttons




    类型：int in [-inf, inf], default 0, 只读




    ## hard_min #




    说明：Minimum value used by buttons




    类型：int in [-inf, inf], default 0, 只读




    ## is_array #




    说明：类型：boolean, default False, 只读




    ## soft_max #




    说明：Maximum value used by buttons




    类型：int in [-inf, inf], default 0, 只读




    ## soft_min #




    说明：Minimum value used by buttons




    类型：int in [-inf, inf], default 0, 只读




    ## step #




    说明：Step size used by number buttons, for floats 1/100th of the step size




    类型：int in [0, inf], default 0, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




     




    # PointerProperty(Property) #




    基类：bpy_struct, Property




    说明：RNA pointer property to point to another RNA struct




    ## fixed_type #




    说明：Fixed pointer type, empty if variable type




    类型：Struct, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




    # StringProperty(Property) #




    基类：bpy_struct, Property




    说明：RNA text string property definition




    ## 示例 #


    ```python
    my_string: StringProperty(

            name="用户输入",

            description=":",

            default="",

            maxlen=1024,

            )


````

    ## default #




    说明：String default value




    类型：string, default “”, (只读，可选)




    ## length_max #




    说明：Maximum length of the string, 0 means unlimited




    类型：int in [0, inf], default 0, 只读




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type




    # Struct(bpy_struct) #




    基类： bpy_struct




    说明：RNA structure definition




    ## base #




    说明：Struct definition this is derived from




    类型：Struct, 只读




    ## description #




    说明：Description of the Struct’s purpose




    类型：string, default “”, (只读，可选)




    ## functions #




    说明：类型：bpy_prop_collection of Function, 只读




    ## identifier #




    说明：Unique name used in the code and scripting




    类型：string, default “”, (只读，可选)




    ## name #




    说明：Human readable name




    类型：string, default “”, (只读，可选)




    ## name_property #




    说明：Property that gives the name of the struct




    类型：StringProperty, 只读




    ## nested #




    说明：Struct in which this struct is always nested, and to which it logically belongs




    类型：Struct, 只读




    ## properties #




    说明：Properties in the struct




    类型：bpy_prop_collection of Property, 只读




    ## property_tags #




    说明：Tags that properties can use to influence behavior




    类型：bpy_prop_collection of EnumPropertyItem, 只读




    ## translation_context #




    说明：Translation context of the struct’s name




    类型：string, default “”, (只读，可选)




    ## classmethod bl_rna_get_subclass() #




    全名：classmethod bl_rna_get_subclass(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：RNA类型，未找到则为默认。




    返回类型：bpy.types.Struct subclass




    ## classmethod bl_rna_get_subclass_py() #




    全名：classmethod bl_rna_get_subclass_py(id, default=None)




    参数：id (字符串) 。 The RNA type identifier.




    返回：类，未找到则为默认。




    返回类型：type
