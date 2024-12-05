---
title: 属性定义（bpy.props）
order: 3
---

# 属性定义（bpy.props）

## Description

[bpy.props](https://docs.blender.org/api/master/bpy.props.html)

该模块用于定义属性，用于扩展 Blender 的内部数据。这些函数的结果用于在 Blender 给注册的类分配属性，不能直接使用。

::: tip
这些函数的所有参数必须作为关键字传递。
:::

### 赋值给现有类

自定义属性可以被添加到`ID`, `Bone` and `PoseBone`的任何子类中。

这些属性可以像 Blender 现有的属性一样被动画化，被用户界面和 python 访问。

```python
import bpy

# 为现有类分配一个自定义属性。
bpy.types.Material.custom_float = bpy.props.FloatProperty(name="Test Property")

# 测试该属性是否存在
bpy.data.materials[0].custom_float = 5.0

```

### 操作项实例

自定义属性经常基于 python 的操作项使用。

在文本编辑器中运行这段代码，或者点击三维视窗 N 面板的按钮来测试。后者会在重做面板中显示这些属性，并允许改变。

![](https://cdn.yuelili.com/20220119232454.png)

![](https://cdn.yuelili.com/20220119232521.png)

```python
import bpy

# 操作项
class OBJECT_OT_property_example(bpy.types.Operator):
    bl_idname = "object.property_example"
    bl_label = "属性测试"
    bl_options = {'REGISTER', 'UNDO'}

    # 定义RNA属性
    my_float: bpy.props.FloatProperty(name="一些浮点值")
    my_bool: bpy.props.BoolProperty(name="开/关")
    my_string: bpy.props.StringProperty(name="字符串值")

    def execute(self, context):
        self.report(
            {'INFO'}, 'F: %.2f  B: %s  S: %r' %
            (self.my_float, self.my_bool, self.my_string)
        )
        print('浮点为:', self.my_float)
        print('布尔为:', self.my_bool)
        print('字符串为:', self.my_string)
        return {'FINISHED'}

# N面板
class OBJECT_PT_property_example(bpy.types.Panel):
    bl_idname = "object_PT_property_example"
    bl_label = "Property Example"
    bl_space_type = 'VIEW_3D'
    bl_region_type = 'UI'
    bl_category = "Tool"

    def draw(self, context):
        # 当用户单击按钮时 可以更改属性值
        props = self.layout.operator('object.property_example')
        props.my_bool = True
        props.my_string = "我是47?"

        # 可以动态更改属性
        if context.object:
            props.my_float = context.object.location.x
        else:
            props.my_float = 327


bpy.utils.register_class(OBJECT_OT_property_example)
bpy.utils.register_class(OBJECT_PT_property_example)

# 赋值（确保在3D视窗测试）
bpy.ops.object.property_example(
    my_float=47,
    my_bool=True,
    my_string="我是327?",
)


```

### 属性组测试

属性组(PropertyGroup)可将自定义设置收集到一个值里，省得单独设置一堆

按 F3，输入 test ops，下方会显示已设置的字符串信息

![](https://cdn.yuelili.com/20220119235057.png)

```python
import bpy

# 单纯用于打印信息的操作项
class Test_Ops(bpy.types.Operator):
    bl_idname = "test.info"
    bl_label = "test ops"

    def execute(self, context):
        material = bpy.data.materials[0]
        # 打印信息
        self.report({"INFO"},material.my_settings.my_string)
        return {"FINISHED"}

class MaterialSettings(bpy.types.PropertyGroup):
    # 定义仨属性
    my_int: bpy.props.IntProperty()
    my_float: bpy.props.FloatProperty()
    my_string: bpy.props.StringProperty()


# 材质里追加设置
bpy.types.Material.my_settings = bpy.props.PointerProperty(type=MaterialSettings)

# 测试设置是否生效
material = bpy.data.materials[0]

material.my_settings.my_int = 5
material.my_settings.my_float = 3.0
material.my_settings.my_string = "Foo"

# 注册
bpy.utils.register_class(Test_Ops)
bpy.utils.register_class(MaterialSettings)

```

### 集合属性

自定义属性可以添加到 `ID`, `Bone` and `PoseBone`子类中

![](https://cdn.yuelili.com/20220120002856.png)

```python
import bpy

# 定义一个集合类
class SceneSettingItem(bpy.types.PropertyGroup):
    name: bpy.props.StringProperty(name="Test Property", default="Unknown")
    value: bpy.props.IntProperty(name="Test Property", default=22)

bpy.utils.register_class(SceneSettingItem)

# 实例化
bpy.types.Scene.my_settings = bpy.props.CollectionProperty(type=SceneSettingItem)

# 对设置进行追加属性（假设选择了obj）
my_item = bpy.context.scene.my_settings.add()
my_item.name = "小明"
my_item.value = 100

my_item = bpy.context.scene.my_settings.add()
my_item.name = "小王"
my_item.value = 98

for my_item in bpy.context.scene.my_settings:
    print(my_item.name, my_item.value)

```

### 更新示例

当一个属性被改变时，可以执行一些命令，用来更新其他属性或与外部数据同步。

除了 CollectionProperty，所有的属性都有更新(update)函数。

```python
import bpy

# 定义更新函数
def update_func(self, context):
    print("my test function", self)

# 定义个属性
bpy.types.Scene.testprop = bpy.props.FloatProperty(update=update_func)

# 赋值
bpy.context.scene.testprop = 11.0

# 执行更新函数
# >>> my test function <bpy_struct, Scene("Scene")>

# 查看新值
bpy.context.scene.testprop
# >>> 11.0
```

### Getter/Setter 示例

Getter/setter 函数可用于布尔、int、float、string 和 enum 属性。

如果定义了这些函数，将不会被自动存储属性在 ID 属性中。相反，从 API 读取或写入属性时，会调用 get 和 set 函数。

```python
import bpy

# ---- 从ID属性读写属性 ---- #
# 这就是RNA内部会做的事情。
def get_float(self):
    return self["testprop"]


def set_float(self, value):
    self["testprop"] = value


bpy.types.Scene.test_float = bpy.props.FloatProperty(get=get_float, set=set_float)

# ---- 字符串 ---- #
# 只读，返回当前时间（因为没加写入set函数）
def get_date(self):
    import datetime
    return str(datetime.datetime.now())


bpy.types.Scene.test_date = bpy.props.StringProperty(get=get_date)

# ---- 布尔数组 ---- #
# Set函数存储一个布尔值，作为第二部分返回。
# 数组获取器(get)必须返回一个列表或元组
# 数组长度必须与属性向量长度匹配
def get_array(self):
    return (True, self["somebool"])


def set_array(self, values):
    self["somebool"] = values[0] and values[1]


bpy.types.Scene.test_array = bpy.props.BoolVectorProperty(size=2, get=get_array, set=set_array)

# ---- 枚举属性 ---- #
# 注意：getter/setter回调必须使用整数标识符!
test_items = [
    ("RED", "Red", "", 1),
    ("GREEN", "Green", "", 2),
    ("BLUE", "Blue", "", 3),
    ("YELLOW", "Yellow", "", 4),
]


def get_enum(self):
    import random
    return random.randint(1, 4)


def set_enum(self, value):
    print("setting value", value)


bpy.types.Scene.test_enum = bpy.props.EnumProperty(items=test_items, get=get_enum, set=set_enum)


# 测试属性
scene = bpy.context.scene
scene.test_float = 12.34
scene.test_array = (True, False)
# scene.test_date = "blah"   # 写入字符串属性会报错
scene.test_enum = 'BLUE'


print('test_float:', scene.test_float)
print('test_array:', tuple(scene.test_array))
print('test_date:', scene.test_date)
print('test_enum:', scene.test_enum)

# 理论输出:
# test_float: 12.34000015258789
# test_array: (True, False)
# test_date: 2018-03-14 11:36:53.158653
# setting value 3
# test_enum: GREEN

```

### 一口气学 15 个参数

返回：一个新的 XX 属性

## 继承

type：类型(class)。继承 bpy.types.PropertyGroup 的一个子类。

## 属性定义

name：名称(字符串) 。用户界面显示名称

description：描述(字符串) 。用于工具提示和 api 文档的文本。

default：默认(sequence) 。布尔序列的长度。

## 用户界面参数

min：硬性最小值(float) 。低于此值会被裁剪

max：硬性最大值(float) 。高于此值会被裁剪

soft_min：软性最小值(float) 。用户界面拖拽滑块不会低于此值

soft_max：软性最大值(float) 。用户界面拖拽滑块不会高于此值

step ：步长(int) 。UI 中的增量/减量的步长，单位为[1, 100]，默认为 3（警告：实际值为/100）。

## 属性选项

options：选项(set) 。枚举

override：覆盖(set) 。枚举

tags：标签(set) 。枚举，由父类定义的标签决定。

subtype：子类型(字符串) 。枚举，

precision：精度(int) 。显示小数点位数，在[0,6]中。对于单位为'NONE'或'TIME'（帧数）且步长可被 100 整除的精确整数值，分数会自动隐藏。

size：维度 (int or int sequence) 。一个整数序列，用于来定义多维数组。向量尺寸为[1, 32]。

unit：单位(字符串) 。枚举，

## 属性自带函数（可覆盖）

update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。

get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。

set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

poll ：检测(function) 。检测函数，确定项目对该属性是否有效。必须接受 2 个值（self, object）并返回 Bool。

### BoolProperty()

全名：bpy.props.BoolProperty(name='', description='', default=False,options={'ANIMATABLE'}, override=set(), tags=set(), subtype='NONE',update=None, get=None, set=None)

返回：一个新的布尔属性

参数：

- name：名称(字符串)。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘PIXEL’, ‘UNSIGNED’, ‘PERCENTAGE’, ‘FACTOR’, ‘ANGLE’, ‘TIME’, ‘DISTANCE’, ‘DISTANCE_CAMERA’, ‘POWER’, ‘TEMPERATURE’, ‘NONE’].
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### BoolVectorProperty()

全名：bpy.props.BoolVectorProperty(name='', description='', default=(False,False, False), options={'ANIMATABLE'}, override=set(), tags=set(),subtype='NONE', size=3, update=None, get=None, set=None)

返回：一个新的布尔向量属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- default：默认(sequence) 。布尔序列的长度。
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘COLOR’, ‘TRANSLATION’, ‘DIRECTION’, ‘VELOCITY’, ‘ACCELERATION’, ‘MATRIX’, ‘EULER’, ‘QUATERNION’, ‘AXISANGLE’, ‘XYZ’, ‘XYZ_LENGTH’, ‘COLOR_GAMMA’, ‘COORDINATES’, ‘LAYER’, ‘LAYER_MEMBER’, ‘NONE’].
- size：维度 (int or int sequence) 。一个整数序列，用于来定义多维数组。向量尺寸为[1, 32]。
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### bpy.props.CollectionProperty()

全名：bpy.props.CollectionProperty(type=None, name='', description='',options={'ANIMATABLE'}, override=set(), tags=set())

返回：一个新的集合属性

参数：

- type：类型(class)。继承 bpy.types.PropertyGroup 的一个子类。
- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’, ‘NO_PROPERTY_NAME’, ‘USE_INSERTION’].
- tags：标签(set) 。枚举，由父类定义的标签决定。

### bpy.props.EnumProperty()

全名：bpy.props.EnumProperty(items, name='', description='', default=None,options={'ANIMATABLE'}, override=set(), tags=set(), update=None, get=None,
set=None)

返回：一个新的枚举属性

参数：

- items：项目(字符串序列或一个函数)。枚举项的序列，格式为：[（标识符，名称，描述，图标，数字），...] ，前三个参数必填
  - identifier：标识符。用于 Python 访问。
  - name：名称。接口的名称。
  - description：描述。用于文档和工具提示。
  - icon：图标。一个图标字符串标识符或整数的图标值（例如，由 bpy.types.UILayout.icon 返回）。
  - number：数字。作为这个项目的标识符的独特值（存储在文件数据中）。当标识符可能需要改变时使用。如果使用 ENUM_FLAG 选项，值是位掩码，应该是 2 的幂。

::: tip
当一个项目只包含他们定义的 4 个项目（标识符、名称、描述、数字）。
:::

分隔符可以使用 None 而不是元组来添加。对于动态值，可以传递一个回调，返回一个与静态列表相同格式的列表。该函数必须接受 2 个参数（self，context），context 可以是 None。

::: danger
使用回调的一个已知错误，Python 必须保持对回调返回字符串引用，否则 Blender 会崩溃。
:::

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- default ：字符串， integer or set) 。The default value for this enum, a string from the identifiers used in items, or integer matching an item number. If the ENUM_FLAG option is used this must be a set of such string identifiers instead. WARNING: Strings can not be specified for dynamic enums (i.e. if a callback function is given as items parameter).
- default ：字符串， 整数或元组 。默认枚举值，来自项目中使用的标识符字符串，或者匹配项目编号的整数。如果使用 ENUM_FLAG 选项，这必须是一组这样的字符串标识符来代替。

::: danger
警告：不能为动态枚举指定字符串（即使用回调函数作为项目参数）。
:::

- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘ENUM_FLAG’, ‘LIBRARY_EDITABLE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘PIXEL’, ‘UNSIGNED’, ‘PERCENTAGE’, ‘FACTOR’, ‘ANGLE’, ‘TIME’, ‘DISTANCE’, ‘DISTANCE_CAMERA’, ‘POWER’, ‘TEMPERATURE’, ‘NONE’].
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### FloatProperty()

全名：bpy.props.FloatProperty(name='', description='', default=0.0, min=-3.402823e+38, max=3.402823e+38, soft_min=- 3.402823e+38,soft_max=3.402823e+38, step=3, precision=2, options={'ANIMATABLE'},override=set(), tags=set(), subtype='NONE', unit='NONE', update=None,
get=None, set=None)

Returns a new float (single precision) property definition.

返回：一个单精度浮点属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- min：硬性最小值(float) 。低于此值会被裁剪
- max：硬性最大值(float) 。高于此值会被裁剪
- soft_min：软性最小值(float) 。用户界面拖拽滑块不会低于此值
- soft_max：软性最大值(float) 。用户界面拖拽滑块不会高于此值
- step ：步长(int) 。UI 中的增量/减量的步长，单位为[1, 100]，默认为 3（警告：实际值为/100）。
- precision：精度(int) 。显示小数点位数，在[0, 6]中。对于单位为'NONE'或'TIME'（帧数）且步长可被 100 整除的精确整数值，分数会自动隐藏。
- options (set) 。Enumerator in [‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- options：选项(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘PIXEL’, ‘UNSIGNED’, ‘PERCENTAGE’, ‘FACTOR’, ‘ANGLE’, ‘TIME’, ‘DISTANCE’, ‘DISTANCE_CAMERA’, ‘POWER’, ‘TEMPERATURE’, ‘NONE’].
- unit：单位(字符串) 。枚举，[‘NONE’, ‘LENGTH’, ‘AREA’, ‘VOLUME’, ‘ROTATION’, ‘TIME’, ‘VELOCITY’, ‘ACCELERATION’, ‘MASS’, ‘CAMERA’, ‘POWER’].
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### FloatVectorProperty()

全名：bpy.props.FloatVectorProperty(name='', description='', default=(0.0, 0.0,0.0), min=sys.float_info.min, max=sys.float_info.max,soft_min=sys.float_info.min, soft_max=sys.float_info.max, step=3, precision=2,options={'ANIMATABLE'}, override=set(), tags=set(), subtype='NONE',unit='NONE', size=3, update=None, get=None, set=None)

返回：一个新的浮点向量属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- default：默认(sequence) 。浮点序列的长度。
- min：硬性最小值(float) 。低于此值会被裁剪
- max：硬性最大值(float) 。高于此值会被裁剪
- soft_min：软性最小值(float) 。用户界面拖拽滑块不会低于此值
- soft_max：软性最大值(float) 。用户界面拖拽滑块不会高于此值
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- step ：步长(int) 。UI 中的增量/减量的步长，单位为[1, 100]，默认为 3（警告：实际值为/100）。
- precision：精度(int) 。显示小数点位数，在[0, 6]中。对于单位为'NONE'或'TIME'（帧数）且步长可被 100 整除的精确整数值，分数会自动隐藏。
- subtype：子类型(字符串) 。枚举，[‘COLOR’, ‘TRANSLATION’, ‘DIRECTION’, ‘VELOCITY’, ‘ACCELERATION’, ‘MATRIX’, ‘EULER’, ‘QUATERNION’, ‘AXISANGLE’, ‘XYZ’, ‘XYZ_LENGTH’, ‘COLOR_GAMMA’, ‘COORDINATES’, ‘LAYER’, ‘LAYER_MEMBER’, ‘NONE’].
- unit：单位(字符串) 。枚举，[‘NONE’, ‘LENGTH’, ‘AREA’, ‘VOLUME’, ‘ROTATION’, ‘TIME’, ‘VELOCITY’, ‘ACCELERATION’, ‘MASS’, ‘CAMERA’, ‘POWER’].
- size：维度 (int or int sequence) 。一个整数序列，用于来定义多维数组。向量尺寸为[1, 32]。
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### IntProperty()

全名：bpy.props.IntProperty(name='', description='', default=0, min=- 2 ** 31,max=2 ** 31 - 1, soft_min=- 2 ** 31, soft_max=2 ** 31 - 1, step=1,options={'ANIMATABLE'}, override=set(), tags=set(), subtype='NONE',update=None, get=None, set=None)

返回：一个新的整数属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- min：硬性最小值(float) 。低于此值会被裁剪
- max：硬性最大值(float) 。高于此值会被裁剪
- soft_min：软性最小值(float) 。用户界面拖拽滑块不会低于此值
- soft_max：软性最大值(float) 。用户界面拖拽滑块不会高于此值
- step (int) 。Step of increment/decrement in UI, in [1, 100], defaults to 1 (WARNING: unused currently!).
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘PIXEL’, ‘UNSIGNED’, ‘PERCENTAGE’, ‘FACTOR’, ‘ANGLE’, ‘TIME’, ‘DISTANCE’, ‘DISTANCE_CAMERA’, ‘POWER’, ‘TEMPERATURE’, ‘NONE’].
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### IntVectorProperty()

全名：bpy.props.IntVectorProperty(name='', description='', default=(0, 0, 0),min=- 2 ** 31, max=2 ** 31 - 1, soft_min=- 2 ** 31, soft_max=2 ** 31 - 1,step=1, options={'ANIMATABLE'}, override=set(), tags=set(), subtype='NONE',size=3, update=None, get=None, set=None)

返回：一个新的整数向量属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- default：默认(sequence) 。整数序列的长度。
- min：硬性最小值(float) 。低于此值会被裁剪
- max：硬性最大值(float) 。高于此值会被裁剪
- soft_min：软性最小值(float) 。用户界面拖拽滑块不会低于此值
- soft_max：软性最大值(float) 。用户界面拖拽滑块不会高于此值
- step ：步长(int) 。UI 中的增量/减量的步长，单位为[1, 100]，默认为 1（警告：当前无法使用）。
- options (set) 。Enumerator in [‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- options：选项(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘COLOR’, ‘TRANSLATION’, ‘DIRECTION’, ‘VELOCITY’, ‘ACCELERATION’, ‘MATRIX’, ‘EULER’, ‘QUATERNION’, ‘AXISANGLE’, ‘XYZ’, ‘XYZ_LENGTH’, ‘COLOR_GAMMA’, ‘COORDINATES’, ‘LAYER’, ‘LAYER_MEMBER’, ‘NONE’].
- size：维度 (int or int sequence) 。一个整数序列，用于来定义多维数组。向量尺寸为[1, 32]。
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。

### PointerProperty()

全名：bpy.props.PointerProperty(type=None, name='', description='',options={'ANIMATABLE'}, override=set(), tags=set(), poll=None, update=None)

返回：一个新的 pointer 属性

参数：

- type：类型(class)。继承 bpy.types.PropertyGroup 的一个子类。
- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- poll ：检测(function) 。检测函数，确定项目对该属性是否有效。必须接受 2 个值（self, object）并返回 Bool。
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。

### RemoveProperty()

全名：bpy.props.RemoveProperty(cls, attr)

说明：动态移除定义的属性

参数：

- cls：类(type) 。包含该属性的类（必须是一个位置 positional 参数）。
- attr：属性字符串 。属性名称（必须作为一个关键词传递）。

注意：通常情况下，不需要直接访问这个函数。而是使用 del cls.attr

### bpy.props.StringProperty()

全名：bpy.props.StringProperty(name='', description='', default='', maxlen=0,options={'ANIMATABLE'}, override=set(), tags=set(), subtype='NONE',update=None, get=None, set=None)

返回：一个新的字符串属性

参数：

- name：名称(字符串) 。用户界面显示名称
- description：描述(字符串) 。用于工具提示和 api 文档的文本。
- default：默认值(字符串) 。初始化字符串
- maxlen：最大长度(int) 。字符串的最大长度
- options：选项(set) 。枚举，[‘HIDDEN’, ‘SKIP_SAVE’, ‘ANIMATABLE’, ‘LIBRARY_EDITABLE’, ‘PROPORTIONAL’,’TEXTEDIT_UPDATE’].
- override：覆盖(set) 。枚举，[‘LIBRARY_OVERRIDABLE’].
- tags：标签(set) 。枚举，由父类定义的标签决定。
- subtype：子类型(字符串) 。枚举，[‘FILE_PATH’, ‘DIR_PATH’, ‘FILE_NAME’, ‘BYTE_STRING’, ‘PASSWORD’, ‘NONE’].
- update：更新(function) 。覆盖原生更新函数，必须接收 2 个值（self, context）并返回 None。警告：没有自带无限递归检查。
- get：获取(function) 。覆盖原生获取函数，必须接收一个值（self）并返回属性的值。
- set：设置(function) 。覆盖原生设置函数，必须接收两个值（self, value）并返回 None。
