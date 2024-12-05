---
title: 参考 API 用法
order: 2
---

# 【BL开发】参考 API 用法

## 参考 API 范围

参考 API 涵盖`bpy.types`，它通过访问`bpy.context`-用户上下文 或`bpy.data`\- Blender **数据** 。

其他模块，例如`bmesh`和`aud`没有使用 Blender 的数据 API，因此本文档不适用于这些模块。

### 数据访问

使用参考 API 的重点在于找出访问 blend 数据。

最好先了解 Blender 中的 ID 数据块（ID data-blocks），因为会经常会找到与它们相关的属性。

#### ID Data

**ID 数据** 在 Blender 中用作顶级 **数据容器** 。从用户界面来看不是很明显，但在开发时需要了解。ID
数据类型包括场景、组、对象、网格、工作区、世界、骨架、图像和纹理。

有关完整列表，请参阅[`bpy.types.ID`](https://docs.blender.org/api/current/bpy.types.ID.html#bpy.types.ID "bpy.types.ID")的子类。

以下是 ID 数据块共享的一些特征：

- ID 是 BL 文件数据，因此加载新的 BL 文件会重新加载一组全新的数据块。
- ID 可以在 Python 中从`bpy.data.*`进行访问.
- 每个数据块都有一个独特的`.name`属性，显示在界面中。
- 动画数据存储在 IDs `.animation_data`中。
- ID 是唯一可以在 blend 文件之间链接的数据类型。
- 可以通过 Python 添加/复制和删除 ID。
- ID 有自己的垃圾收集系统，可以在保存时释放未使用的 ID。
- 当一个数据块引用了一些外部数据时，这通常是一个 ID 数据块。

#### 简单的数据访问

来个简单的示例，使用 Python 脚本调整对象的位置。先来收集数据所在的信息。

首先在界面中找到这个设置。从属性编辑器>对象>变换>位置，右键按钮上下文菜单中选择[_Online PythonReference_](https://docs.blender.org/api/3.0/bpy.types.Object.html#bpy.types.Object.location)，

![](https://cdn.yuelili.com/20220108162346.png)

但是你现在知道你必须使用它并且它是一个由三个浮点数组成的数组。以下来自官方文档：

![](https://cdn.yuelili.com/20220108162641.png)

接下来是找出访问对象的位置，到页面底部的引用部分。对象有很多引用，但访问对象最常见的地方之一是通过上下文。

这时候很容易不知所措，因为`Object`在很多地方都引用了它：修饰符、函数、纹理和约束。

但是，想访问用户选择的任何数据，通常只需参考[`bpy.context`](https://docs.blender.org/api/current/bpy.context.html#module-bpy.context "bpy.context")。

即使这样，在这种情况下还是有很多（函数），不过仔细阅读这些参考，您会发现大多数都是针对特定模式的。

比如编写仅在 **权重绘制** 模式下的工具，`weight_paint_object`是不错之选。

要访问用户选择的最后一个项目，查找`active`成员。访问激活项目是 Blender 的一个惯例，如`active_bone`，`active_pose_bone`，`active_node`等，此时可以使用`active_object`。

所以现在你有足够的信息来找到活动对象的位置。

```python
bpy.context.active_object.location # 控制台：Vector((0.0, 0.0, 0.0))

```

访问引用对象的另一个常见位置参考是：[bpy.types.BlendData.objects](https://docs.blender.org/api/current/bpy.types.BlendData.html#bpy.types.BlendData.objects)

`bpy.data.objects`是一个对象集合，我们可以需要访问它的一个成员：

```python
bpy.data.objects["Cube"].location # 控制台：Vector((0.0, 0.0, 0.0))

```

#### 嵌套属性

前面示例非常简单，因为`Object`的`location`属性可以直接从上下文中访问。

来点更复杂的例子：

```python
# 访问 Cycles render engine样本数量.
bpy.context.scene.cycles.samples

# 访问 paint brush size 权重.
bpy.context.tool_settings.weight_paint.brush.size

# 检查当前是否为全屏模式
bpy.context.window.screen.show_fullscreen
```

如上所述，有时想访问某种方式嵌套的数据，这会导致通过一些间接的方式。这些属性的 **排列方式** 与数据在内部的 **存储方式** （在 Blender 的 C 代码中）相匹配，一般是合乎逻辑的，但并不总是那么“逻辑”。所以需要一些时间来学习，它可以帮助您了解在 Blender
中数据如何组合在一起，在编写脚本时这很重要。

开始编写脚本时，如果不确定如何访问所需数据。有如下窍门：

- 使用 Python 控制台的自动完成功能来检查属性。 可能不完善，但优点是方便查看属性值并交互式查看结果 _。_
- 从用户界面复制数据路径。 在 [Copy Data Path 章节](https://docs.blender.org/api/current/info_api_reference.html#info-data-path-copy)细说。
- 使用文档来遵循参考。 在 [间接数据访问中](https://docs.blender.org/api/current/info_api_reference.html#info-data-path-indirect)进一步解释。

#### 复制数据路径

Blender 具有 Python 工具提示。在工具说明下方，以 Python：... 显示。这样就不用一直翻文档了。

![](https://cdn.yuelili.com/20220108164312.png)

示例：将表面细分修改器添加到 Cube。鼠标悬停在标有视图层级上

![](https://cdn.yuelili.com/20220108165008.png)

工具提示[`bpy.types.SubsurfModifier.levels`](https://docs.blender.org/api/current/bpy.types.SubsurfModifier.html#bpy.types.SubsurfModifier.levels"bpy.types.SubsurfModifier.levels")，您想要从对象到此属性的路径。

请注意，复制的文本不会包含`bpy.data.collection["name"].`组件，因为通常使用上下文查找而不是[`bpy.types.ID`](https://docs.blender.org/api/current/bpy.types.ID.html#bpy.types.ID"bpy.types.ID")按名称访问每个实例（通过集合查找）。

在 Python 控制台中输入 ID
路径[`bpy.context.active_object`](https://docs.blender.org/api/current/bpy.context.html#id2"bpy.context.active_object")。但不要执行。

现在在按钮的上下文菜单（右键）选择 _Copy Data Path_

![](https://cdn.yuelili.com/20220108165453.png)

然后将结果粘贴到控制台中（显示属性值）：

```python
bpy.context.active_object.modifiers["细分"].levels # 中文版 > 2

```

对属性进行修改

```python
bpy.context.active_object.modifiers["Subdivision"].levels = 2 #
英文版
```

![](https://cdn.yuelili.com/20220108165938.png)

#### 间接数据访问

来尝试访问活动雕刻笔刷纹理。例如访问画笔的纹理并调整其`contrast`.

1.默认场景开始，启用雕刻模式。

2.展开画笔设置面板的纹理子面板，添加一个新纹理。 （请注意，纹理数据块菜单本身并没有非常有用的链接（您可以查看工具提示）。）

![](https://cdn.yuelili.com/20220108170502.png)

- 3.侧边栏中未显示对比度设置，因此请在[属性编辑器中](https://docs.blender.org/manual/en/dev/render/materials/legacy_textures/colors.html#bpy-types-texture-contrast "（在 Blender 3.1 手册 v3.1 中）")查看纹理 。

![](https://cdn.yuelili.com/20220108171041.png)

- 4.打开对比度的在线参考。也就是`bpy.types.Texture.contrast`. 现在可以看到`contrast`是纹理的属性。

![](https://cdn.yuelili.com/20220108171224.png)

- 5.关于如何从笔刷访问纹理，请查看页面底部的参考资料。有时候参考很多，但此时需要使用`tool_settings.sculpt.brush.texture`.
- 6.现在你知道了，纹理可以通过`bpy.data.brushes["BrushName"].texture` 访问，但一般会使用 **活动笔刷** 访问，而不是 **笔刷名称**

现在可以使用 Python 控制台来形成访问画笔纹理对比度所需的嵌套属性：

上下文 ‣ 工具设置 ‣ 雕刻 ‣ 画笔 ‣ 纹理 ‣ 对比度（Context ‣ Tool Settings ‣ Sculpt ‣ Brush ‣
Texture ‣ Contrast.）

由于每个属性都是逐层访问，因此可以：

```python
bpy.context.tool_settings.sculpt.brush.texture.contrast
```

或者直接访问笔刷

```python
bpy.data.textures["Texture"].contrast
```

### 操作项

Blender
中的大多数热键和按钮都调用一个操作项，也就是[`bpy.ops`](https://docs.blender.org/api/current/bpy.ops.html#module-bpy.ops "bpy.ops").

要查看 Python 等效项，请将鼠标悬停在按钮上并查看工具提示，例如`Python: bpy.ops.render.render()`。

如果没有工具提示，则未使用操作项，无法从 Python 访问。

如果想在脚本中使用，可以鼠标在按钮悬停（`` Ctrl`-`C ``），将其复制到剪贴板。还可以使用使用上下文菜单查看在线参考，这主要显示参数及其默认值，

但是，用 Python 编写的操作项会显示它们文件和行号，如果对源码感兴趣的话，会很有用。

::: tip
并非所有操作项都可以从 Python 中有效地调用，请参阅使用操作项。

:::

#### 信息编辑器

Blender 记录运行的操作项并显示在信息编辑器中。选择脚本工作区查看其输出。可以执行一些操作并看到它们出现，例如删除一个顶点。

可以选择条目复制，粘贴到文本编辑器或 Python 控制台中。

::: tip
并非所有操作项都注册显示，例如缩放视图不太重要，就没显示。要显示所有运行的操作项，请参阅[显示所有操作项](https://docs.blender.org/api/current/info_tips_and_tricks.html#info-show-all-operators)。

:::
