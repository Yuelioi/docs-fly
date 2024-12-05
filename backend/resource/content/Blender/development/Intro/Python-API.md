---
title: Python API 概述
order: 1
---

# 【BL开发】Python API 概述

## 介绍

Blender 将其 Python 模块（例如`bpy`和`mathutils`）提供给嵌入式解释器，这样导入脚本就可以访问 Blender
的数据、类和函数。

处理 Blender 数据的脚本需要导入 bpy 模块才能工作。

示例：移动一个“Cube”的对象上的顶点：

```python
import bpy
bpy.data.objects["Cube"].data.vertices[0].co.x += 1.0

```

这会直接修改 Blender 内部数据。运行后，视窗直接更新。

### 默认环境

Blender 启动会扫描安装目录 `scripts/startup/`Python 模块的目录并导入。

### 脚本加载

要注意直接执行脚本和将脚本作为模块导入之间的区别。

通过直接执行脚本来扩展 Blender 意味着脚本定义的类在脚本完成执行后在 Blender 中仍然可用。与将脚本作为模块导入相比，以这种方式使用脚本使得未来访问它们的类（例如取消注册它们）更加困难。当脚本作为模块导入时，其类实例将保留在模块内，稍后可以通过再次导入该模块进行访问。

出于这个原因，最好避免通过注册类直接执行扩展 Blender 的脚本。

以下是在 Blender 中直接运行脚本的一些方法：

- 在文本编辑器中加载并按 _Run Script_ 。
- 键入或粘贴到交互式控制台中。
- 使用 Blender 从命令行执行 Python 文件，例如：>>> blender --python /home/me/my_script.py

作为模块运行：

- 文本编辑器或交互式控制台的命令。`import 你的模块`
- 作为文本数据块打开并检查注册选项，这将与混合文件一起加载。
- 复制到其中一个目录中`scripts/startup`，会启动时自动导入。
- 定义为附加组件（插件），启用附加组件会将其作为 Python 模块加载。

#### 附加组件（Add-on）

保存在`scripts/addons`，如果从用户首选项中选择，它们只会在启动时加载。

附加组件和内置 Python 模块唯一区别是它必须包含一个`bl_info`变量，Blender 使用该变量读取元数据，例如名称、作者、类别和项目链接。

用户首选项加载项列表用于`bl_info`显示有关每个加载项的信息。

### 通过类集成

文本编辑器测试还行，一般还是会集成到 BL 自身 UI 界面里。

Blender Python API 允许集成到：

- `bpy.types.Panel`
- `bpy.types.Menu`
- `bpy.types.Operator`
- `bpy.types.PropertyGroup`
- `bpy.types.KeyingSet`
- `bpy.types.RenderEngine`

至于更高级的功能，例如网格修改器、对象类型或着色器节点，必须使用 C/C++。

对于 Python 集成，Blender 定义了所有类型通用的方法。这是通过创建 Blender 类的 Python 子类来实现的。其实就是自己的 Class
继承父级 Blender 的类

示例：

```python
import bpy
class SimpleOperator(bpy.types.Operator):
    bl_idname = "object.simple_operator"
    bl_label = "工具名称"

    def execute(self, context):
        print("Hello World")
        return {'FINISHED'}

bpy.utils.register_class(SimpleOperator)

```

首先子类继承自`bpy.types.Operator`。类属性以`bl_`作为前缀。这是用于区分 Blender 属性和开发者自定义属性。

接下来是 execute 函数（接受操作项的实例和当前上下文），函数不使用公共前缀(bl)。最后调用 register（注册） 函数，获取类并将其加载到
Blender 中。

关于继承，Blender 不限制使用的类继承的种类， register 将使用在父类中定义的属性和函数。

类混合示例：

```python
import bpy
class BaseOperator:
    def execute(self, context):
        print("Hello World BaseClass")
        return {'FINISHED'}

class SimpleOperator(bpy.types.Operator, BaseOperator):
    bl_idname = "object.simple_operator"
    bl_label = "工具名称"

bpy.utils.register_class(SimpleOperator)

```

这些类没有定义`__init__(self)`函数。虽然`__init__()`和`__del__()`将在定义时被调用，但这些类实例的生命周期短暂。例如，面板每次重绘都会产生新实例，因此变量一般不会存储在面板实例中。相反，持久变量会存储在 Blender 的数据中，以便在 Blender 重新启动时可以恢复。

::: tip
模态操作项是例外，在 Blender 运行时保持它们的实例变量，请参阅模态运算符模板。
:::

因此，一旦类在 Blender 中注册，Blender 会组织实例化类和调用函数 。

事实上，跟常规 Python API 不同，一般很难实例化类，可以直接用 Python 的操作项 API 直接调用它们，例如：

```python
import bpy
bpy.ops.object.simple_operator() // 刚才注册的类
```

用户界面类被赋予了绘制的上下文、按钮、窗口、文件标题、工具栏等，然后在显示该区域时绘制，因此它们永远不会被 Python 脚本直接调用。

### 注册

#### 模块注册

启动时加载的 Blender 模块需要`register()`和`unregister()`功能（有点像功能的绑定与解绑）。

只有注册过的函数，才可以从 Blender 中调用。

简单的 Blender Python 模块：

```python
import bpy

class SimpleOperator(bpy.types.Operator):
    """ 示例见上 """

# 注册
def register():
    bpy.utils.register_class(SimpleOperator)

# 注销
def unregister():
    bpy.utils.unregister_class(SimpleOperator)

if __name__ == "__main__":
    register()
```

### 类注册

向 Blender 注册类，这样该类会被加载到 Blender 中，与 Blender 现有功能一起使用。接着可以从`bpy.types`使用`bl_idname`访问它。

::: tip
如果类名不唯一，可以使用：[`bpy.types.Struct.bl_rna_get_subclass_py()`](bpy.types.Struct.html#bpy.types.Struct.bl_rna_get_subclass_py"bpy.types.Struct.bl_rna_get_subclass_py")。

:::

#### 类间依赖

在定制 Blender 时，可以自定义设置分组，毕竟要与其他脚本共存。为了将这些属性分组，需要定义类，对于组中的子组或组中的集合，要处理注册/不注册的顺序。

自定义属性组本身就是需要被注册的类。

例如，如果你想为一个自定义引擎存储材料设置。

```python
# 创建一个新属性,创建后，可以在控制台测试
# >>> bpy.data.materials[0].my_custom_props.my_float

import bpy

class MyMaterialProps(bpy.types.PropertyGroup):
    my_float: bpy.props.FloatProperty()

def register():
    bpy.utils.register_class(MyMaterialProps)
    bpy.types.Material.my_custom_props: bpy.props.PointerProperty(type=MyMaterialProps)

def unregister():
    del bpy.types.Material.my_custom_props
    bpy.utils.unregister_class(MyMaterialProps)

if __name__ == "__main__":
    register()

```

::: tip
该类 **必须** 先注册再调用属性，否则将引发错误：
:::

```python
# 使用子属性创建组
# bpy.data.materials[0].my_custom_props.sub_group.my_float
import bpy

class MyMaterialSubProps(bpy.types.PropertyGroup):
    my_float: bpy.props.FloatProperty()

class MyMaterialGroupProps(bpy.types.PropertyGroup):
    sub_group: bpy.props.PointerProperty(type=MyMaterialSubProps)

def register():
    bpy.utils.register_class(MyMaterialSubProps)
    bpy.utils.register_class(MyMaterialGroupProps)
    bpy.types.Material.my_custom_props: bpy.props.PointerProperty(type=MyMaterialGroupProps)

def unregister():
    del bpy.types.Material.my_custom_props
    bpy.utils.unregister_class(MyMaterialGroupProps)
    bpy.utils.unregister_class(MyMaterialSubProps)

if __name__ == "__main__":
    register()

```

#### 操作类

属性可以在 Blender 运行时添加和删除，通常在注册或取消注册时完成，但对于某些特殊情况，在脚本运行时修改类型可能很有用。

例如：

```python
# 给属性添加已存在的类型
bpy.types.Object.my_float: bpy.props.FloatProperty()
# 移除
del bpy.types.Object.my_float
```

也可以自定义

```python
class MyPropGroup(bpy.types.PropertyGroup):
    pass
MyPropGroup.my_float: bpy.props.FloatProperty()
```

相当于

```python
class MyPropGroup(bpy.types.PropertyGroup):
  my_float:bpy.props.FloatProperty()
```

#### 动态类定义（高级）

某些情况下，数据分类可能不在 Blender 中，例如外部渲染引擎着色器，最好将它们定义为类型并动态删除。

```python
for i in range(10):
    idname = "object.operator_%d" % i

    def func(self, context):
        print("Hello World", self.bl_idname)
        return {'FINISHED'}

    opclass = type("DynOp%d" % i,
                   (bpy.types.Operator, ),
                   {"bl_idname": idname, "bl_label": "Test", "execute": func},
                   )
    bpy.utils.register_class(opclass)


```

::: tip
type()被调用来定义类。这是在 Python 中创建类的另一种语法，更适合动态构造类。
:::

要调用上个示中的操作项：

```python
>>> bpy.ops.object.operator_1()
Hello World OBJECT_OT_operator_1
{'FINISHED'}

```

```python
>>> bpy.ops.object.operator_2()
Hello World OBJECT_OT_operator_2
{'FINISHED'}

```
