---
title: 上下文访问(bpy.context)
order: 1
---

# 上下文访问(bpy.context)

## 上下文是什么？

[此块文档](https://wiki.blender.org/wiki/Source/Architecture/Context)

需要注意的是，不是谁在调用谁。而是不同类型的代码，可以在上下文呈现不同数据。

它是 blender 数据管理中的一个全局变量集合，它包含了目前被选择的和激活的（物体的两种属性状态）的物体的信息，还记录了当前物体的编辑模式（data 中没有这方面的信息）。当然使用 context 对物体对象的访问以及属性的操作和 data 子模块中相似。//此句引用[CSDN](https://blog.csdn.net/honyniu/article/details/46573505)

![](https://cdn.yuelili.com/20220112004703.png)

所有的代码都可以访问用户偏好（User
Preference），这是一个全局性的东西，并不是真正的上下文。此外，还有 Main，它是一个.blend 文件，有自己的数据块和来自其他.blend 文件的数据块链接。我们应该尽可能 地假设可能有多个这样的文件打开，尽管这还没有实现。

接下来有两个分支。绘图（Draw）和交互式编辑代码可以假设是在窗口管理器（Window Manager）下运行。

在窗口管理器下运行的绘图和操作代码总是在某一层次上工作，在场景（Scene）、区域（Area）或地区（Region）。伴随着这一层次的是某些数据，场景还有一个活动场景，而区域和地区可以另外把其他数据放到上  
下文中，具体是什么取决于区域和地区的类型。

此外，还有些代码不能认为是在窗口管理器下运行。并不是说不能从窗口管理器中调用，而是说不应该假定有一个窗口管理器并访问这样的代码。

然后，这段代码可以在没有窗口选项的情况下以后台模式运行，通常用于渲染。Evaluation（修改器、约束......）和渲染通常会被传递到一个场景和其他数据，但不应该访问 bContext 结构。

文件读写、各种内核函数、窗口管理器代码等都不在此图中。一般来说，这些功能上下文很少或根本没有上下文指针的情况下工作。

### 设置上下文

重要的是要理解，上下文不是持久的。也就是说，如果你想改变活动对象，这不是通过在上下文中设置它来完成的。相反，应该设置场景中的适当属性，然后在上下文查找中使用。当运行一个操作者或绘图时，上下文是一个临时的、局部的状态。

### 回调

上下文主要由回调来定义。屏幕、空间类型和区域类型有一个 context()回调函数。这个函数得到一个上下文成员的名字，然后检查它是否拥有这个成员，如果拥有，则返回和 RNA 指针或集合。此外，每个上下文回调还应该提供一个它所提供的上下文成员的列表。

### UI 布局

此外，还可以为 UI 布局指定一个更具体的上下文。这是为像修改器或约束这样的东西准备的。在每个修改器 UI 布局框中，"修改器"被设置在上下文中，所有从该框中的按钮调用的操作者将在他们的上下文中收到这个。

### 查询

当一个成员从上下文中被请求时，UI 布局上下文将被首先查看，然后是区域回调、区域回调和屏幕回调。一旦找到它，就会返回结果。如果没有找到，将返回一个 NULL 指针或一个空集合。

![](https://cdn.yuelili.com/20220112232552.png)

### 获取上下文

#### 窗口管理器

窗口管理器的上下文是最简单的，这些只是指向屏幕（screen）、区域（area）、空间数据（space data）、地区（region）和地区数据（region data）的指针。首先要确保在上下文中。

对于绘图函数来说，何时在上下文中非常明确，而对于操作项来说，要在 poll()函数检查，或者在运行时验证。否则可能导致 Blender 崩溃

#### 数据

数据上下文比较复杂，重要的是要理解它是基于 RNA 指针和集合的。这使得 Python 脚本也可以自动使用上下文。

在 C 语言中，有两种类型的函数可以调用以从上下文中获取数据。对于某些东西，已经定义了一个访问器函数。注意，在内部这仍然会使用 "edit_object
"字符串来获取数值，这主要是为了方便。

## 上下文介绍

[官方文档链接](https://docs.blender.org/api/current/bpy.context.html)

上下文成员是否可用，取决于当前正在访问的 Blender 区域。

请注意，所有上下文值都是的，但可以通过数据 API 或运行运算符进行修改

全局上下文

这些属性在全局可用。

### area 区域？

全名：bpy.context.area  
说明：  
类型：bpy.types.Area, (readonly)

### asset_file_handle

全名：bpy.context.asset_file_handle 说明：活动资产的文件。不要使用，它会被一个适当的 AssetHandle 设计所取代。

类型：bpy.types.FileSelectEntry, (readonly)

### blend_data

全名：bpy.context.blend_data  
说明：blender 数据，具体能干嘛未知  
类型：bpy.types.BlendData, (readonly)

### collection 集合

全名：bpy.context.collection  
说明：blender 活动项目集合  
类型：bpy.types.Collection, (readonly)

示例：

![](https://cdn.yuelili.com/20220112233402.png)

```python
>>>bpy.context.collection
bpy.data.collections['Collection 2']

```

### engine 引擎

全名：bpy.context.engine

说明：Blender 使用的引擎

类型：string, default “”, (readonly, never None)

```python
>>>bpy.context.collection
'BLENDER_EEVEE'
```

### gizmo_group

全名：bpy.context.gizmo_group  
说明：  
类型：bpy.types.GizmoGroup, (readonly)

### layer_collection

全名：bpy.context.layer_collection  
说明：  
类型：bpy.types.LayerCollection, (readonly)

### mode 模式

![](https://cdn.yuelili.com/20220112234155.png)

全名：bpy.context.mode

说明：当前模式

类型：,默认为 EDIT_MESH，只读

- EDIT_MESH
- EDIT_CURVE
- EDIT_SURFACE
- EDIT_TEXT
- EDIT_ARMATURE
- EDIT_METABALL
- EDIT_LATTICE
- POSE
- SCULPT
- PAINT_WEIGHT
- PAINT_VERTEX
- PAINT_TEXTURE
- PARTICLE
- OBJECT
- PAINT_GPENCIL
- EDIT_GPENCIL
- SCULPT_GPENCIL
- WEIGHT_GPENCIL
- VERTEX_GPENCIL

### preferences

全名：bpy.context.preferences  
说明：首选项  
类型：bpy.types.Preferences, (readonly)

### region 区块？

![](https://cdn.yuelili.com/20220112234658.png)

全名：bpy.context.region

说明：当前所处区块

类型：bpy.types.Region, 只读

示例：

```python
 >>>bpy.context.region
 bpy.data.screens['Layout']...Region

```

### region_data

全名：bpy.context.region_data  
说明：  
类型：bpy.types.RegionView3D, (readonly)

### scene 场景？

全名：bpy.context.scene  
说明：  
类型：bpy.types.Scene, (readonly)

### screen 屏幕

全名：bpy.context.screen  
说明：当前所处屏幕  
类型：bpy.types.Screen, (readonly)

示例：我在雕刻时输入，则...

![](https://cdn.yuelili.com/20220112234919.png)

```python
>>> bpy.context.screen
bpy.data.screens['Sculpting']

```

## space_data

全名：bpy.context.space_data

说明：

类型：bpy.types.Space, (readonly)

## tool_settings

全名：bpy.context.tool_settings

说明：

类型：bpy.types.ToolSettings, (readonly)

## view_layer

全名：bpy.context.view_layer

说明：

类型：bpy.types.ViewLayer, (readonly)

## window

全名：bpy.context.window

说明：

类型：bpy.types.Window, (readonly)

## window_manager

全名：bpy.context.window_manager

说明：

类型：bpy.types.WindowManager, (readonly)

## workspace 工作空间

![](https://cdn.yuelili.com/20220112234658.png)

全名：bpy.context.workspace

说明：工作空间

类型：bpy.types.WorkSpace, (readonly)

# Screen Context

## scene

全名：bpy.context.scene

说明：

类型：bpy.types.Scene

## view_layer

全名：bpy.context.view_layer

说明：

类型：bpy.types.ViewLayer

## visible_objects 可视对象

全名：bpy.context.visible_objects

说明：可视对象

类型：sequence of bpy.types.Object

```python
>>> bpy.context.visible_objects

[bpy.data.objects['Light'], bpy.data.objects['Camera'], bpy.data.objects['Cube.001']]


```

## selectable_objects 可选的对象

![](https://cdn.yuelili.com/20220112235424.png)

全名：bpy.context.selectable_objects

说明：可选择的对象。如果关掉，那么就是无法选择

类型：sequence of bpy.types.Object

## selected_objects 选中的对象

全名：bpy.context.selected_objects

说明：当前选择的对象

类型：sequence of bpy.types.Object，列表，如果没有则为[]

## editable_objects 可编辑对象

全名：bpy.context.editable_objects

说明：可编辑对象。

类型：sequence of bpy.types.Object

## selected_editable_objects 选中可编辑对象

全名：bpy.context.selected_editable_objects

说明：当前选中的，可编辑对象

类型：sequence of bpy.types.Object

## objects_in_mode

全名：bpy.context.objects_in_mode

说明：

类型：sequence of bpy.types.Object

## objects_in_mode_unique_data

全名：bpy.context.objects_in_mode_unique_data

说明：

类型：sequence of bpy.types.Object

## visible_bones

全名：bpy.context.visible_bones

说明：

类型：sequence of bpy.types.EditBone

## editable_bones

全名：bpy.context.editable_bones

说明：

类型：sequence of bpy.types.EditBone

## selected_bones

全名：bpy.context.selected_bones

说明：

类型：sequence of bpy.types.EditBone

## selected_editable_bones

全名：bpy.context.selected_editable_bones

说明：

类型：sequence of bpy.types.EditBone

## visible_pose_bones

全名：bpy.context.visible_pose_bones

说明：

类型：sequence of bpy.types.PoseBone

## selected_pose_bones

全名：bpy.context.selected_pose_bones

说明：

类型：sequence of bpy.types.PoseBone

## selected_pose_bones_from_active_object

全名：bpy.context.selected_pose_bones_from_active_object

说明：

类型：sequence of bpy.types.PoseBone

## active_bone

全名：bpy.context.active_bone

说明：

类型：bpy.types.EditBone

## active_pose_bone

全名：bpy.context.active_pose_bone

说明：

类型：bpy.types.PoseBone

## active_object 激活对象

![](https://cdn.yuelili.com/20220113000807.png)

全名：bpy.context.active_object

说明：带黄点的就是激活对象，即使没选中

类型：bpy.types.Object

## object 对象

全名：bpy.context.object

说明：对象，不知道与激活对象有啥区别

类型：bpy.types.Object

## edit_object 正在编辑的对象

全名：bpy.context.edit_object

说明：正在编辑的对象，需要进入编辑模式

类型：bpy.types.Object

## sculpt_object

全名：bpy.context.sculpt_object

说明：

类型：bpy.types.Object

## vertex_paint_object

全名：bpy.context.vertex_paint_object

说明：

类型：bpy.types.Object

## weight_paint_object

全名：bpy.context.weight_paint_object

说明：

类型：bpy.types.Object

## image_paint_object

全名：bpy.context.image_paint_object

说明：

类型：bpy.types.Object

## particle_edit_object

全名：bpy.context.particle_edit_object

说明：

类型：bpy.types.Object

## pose_object

全名：bpy.context.pose_object

说明：

类型：bpy.types.Object

## active_sequence_strip

全名：bpy.context.active_sequence_strip

说明：

类型：bpy.types.Sequence

## sequences

全名：bpy.context.sequences

说明：

类型：sequence of bpy.types.Sequence

## selected_sequences

全名：bpy.context.selected_sequences

说明：

类型：sequence of bpy.types.Sequence

## selected_editable_sequences

全名：bpy.context.selected_editable_sequences

说明：

类型：sequence of bpy.types.Sequence

## active_nla_track

全名：bpy.context.active_nla_track

说明：

类型：bpy.types.NlaTrack

## active_nla_strip

全名：bpy.context.active_nla_strip

说明：

类型：bpy.types.NlaStrip

## selected_nla_strips

全名：bpy.context.selected_nla_strips

说明：

类型：sequence of bpy.types.NlaStrip

## selected_movieclip_tracks

全名：bpy.context.selected_movieclip_tracks

说明：

类型：sequence of bpy.types.MovieTrackingTrack

## gpencil_data

全名：bpy.context.gpencil_data

说明：

类型：bpy.types.GreasePencil

## gpencil_data_owner

全名：bpy.context.gpencil_data_owner

说明：

类型：bpy.types.ID

## annotation_data

全名：bpy.context.annotation_data

说明：

类型：bpy.types.GreasePencil

## annotation_data_owner

全名：bpy.context.annotation_data_owner

说明：

类型：bpy.types.ID

## visible_gpencil_layers

全名：bpy.context.visible_gpencil_layers

说明：

类型：sequence of bpy.types.GPencilLayer

## editable_gpencil_layers

全名：bpy.context.editable_gpencil_layers

说明：

类型：sequence of bpy.types.GPencilLayer

## editable_gpencil_strokes

全名：bpy.context.editable_gpencil_strokes

说明：

类型：sequence of bpy.types.GPencilStroke

## active_gpencil_layer

全名：bpy.context.active_gpencil_layer

说明：

类型：sequence of bpy.types.GPencilLayer

## active_gpencil_frame

全名：bpy.context.active_gpencil_frame

说明：

类型：sequence of bpy.types.GreasePencilLayer

## active_annotation_layer

全名：bpy.context.active_annotation_layer

说明：

类型：bpy.types.GPencilLayer

## active_operator

全名：bpy.context.active_operator

说明：

类型：bpy.types.Operator

## selected_visible_actions

全名：bpy.context.selected_visible_actions

说明：

类型：sequence of bpy.types.Action

## selected_editable_actions

全名：bpy.context.selected_editable_actions

说明：

类型：sequence of bpy.types.Action

## visible_fcurves

全名：bpy.context.visible_fcurves

说明：

类型：sequence of bpy.types.FCurve

## editable_fcurves

全名：bpy.context.editable_fcurves

说明：

类型：sequence of bpy.types.FCurve

## selected_visible_fcurves

全名：bpy.context.selected_visible_fcurves

说明：

类型：sequence of bpy.types.FCurve

## selected_editable_fcurves

全名：bpy.context.selected_editable_fcurves

说明：

类型：sequence of bpy.types.FCurve

## active_editable_fcurve

全名：bpy.context.active_editable_fcurve

说明：

类型：bpy.types.FCurve

## selected_editable_keyframes

全名：bpy.context.selected_editable_keyframes

说明：

类型：sequence of bpy.types.Keyframe

## ui_list

全名：bpy.context.ui_list

说明：

类型：bpy.types.UIList

## asset_library_ref

全名：bpy.context.asset_library_ref

说明：

类型：bpy.types.AssetLibraryReference

# 3D 视图 View3D

## active_object

全名：bpy.context.active_object

说明：

类型：bpy.types.Object

## selected_ids

全名：bpy.context.selected_ids

说明：

类型：sequence of bpy.types.ID

## 按钮上下文 Buttons Context

## texture_slot

全名：bpy.context.texture_slot

说明：

类型：bpy.types.TextureSlot

## scene

全名：bpy.context.scene

说明：

类型：bpy.types.Scene

## world

全名：bpy.context.world

说明：

类型：bpy.types.World

## object

全名：bpy.context.object

说明：

类型：bpy.types.Object

## mesh

全名：bpy.context.mesh

说明：

类型：bpy.types.Mesh

## armature

全名：bpy.context.armature

说明：

类型：bpy.types.Armature

## lattice

全名：bpy.context.lattice

说明：

类型：bpy.types.Lattice

## curve

全名：bpy.context.curve

说明：

类型：bpy.types.Curve

## meta_ball

全名：bpy.context.meta_ball

说明：

类型：bpy.types.MetaBall

## light

全名：bpy.context.light

说明：

类型：bpy.types.Light

## speaker

全名：bpy.context.speaker

说明：

类型：bpy.types.Speaker

## lightprobe

全名：bpy.context.lightprobe

说明：

类型：bpy.types.LightProbe

## camera

全名：bpy.context.camera

说明：

类型：bpy.types.Camera

## material

全名：bpy.context.material

说明：

类型：bpy.types.Material

## material_slot

全名：bpy.context.material_slot

说明：

类型：bpy.types.MaterialSlot

## texture

全名：bpy.context.texture

说明：

类型：bpy.types.Texture

## texture_user

全名：bpy.context.texture_user

说明：

类型：bpy.types.ID

## texture_user_property

全名：bpy.context.texture_user_property

说明：

类型：bpy.types.Property

## bone

全名：bpy.context.bone

说明：

类型：bpy.types.Bone

## edit_bone

全名：bpy.context.edit_bone

说明：

类型：bpy.types.EditBone

## pose_bone

全名：bpy.context.pose_bone

说明：

类型：bpy.types.PoseBone

## particle_system

全名：bpy.context.particle_system

说明：

类型：bpy.types.ParticleSystem

## particle_system_editable

全名：bpy.context.particle_system_editable

说明：

类型：bpy.types.ParticleSystem

## particle_settings

全名：bpy.context.particle_settings

说明：

类型：bpy.types.ParticleSettings

## cloth

全名：bpy.context.cloth

说明：

类型：bpy.types.ClothModifier

## soft_body

全名：bpy.context.soft_body

说明：

类型：bpy.types.SoftBodyModifier

## fluid

全名：bpy.context.fluid

说明：

类型：bpy.types.FluidSimulationModifier

## collision

全名：bpy.context.collision

说明：

类型：bpy.types.CollisionModifier

## brush

全名：bpy.context.brush

说明：

类型：bpy.types.Brush

## dynamic_paint

全名：bpy.context.dynamic_paint

说明：

类型：bpy.types.DynamicPaintModifier

## line_style

全名：bpy.context.line_style

说明：

类型：bpy.types.FreestyleLineStyle

## collection

全名：bpy.context.collection

说明：

类型：bpy.types.LayerCollection

## gpencil

全名：bpy.context.gpencil

说明：

类型：bpy.types.GreasePencil

## hair

全名：bpy.context.hair

说明：

类型：bpy.types.Hair

## pointcloud

全名：bpy.context.pointcloud

说明：

类型：bpy.types.PointCloud

## volume

全名：bpy.context.volume

说明：

类型：bpy.types.Volume

# 图片上下文 Image Context

## edit_image

全名：bpy.context.edit_image

说明：

类型：bpy.types.Image

## edit_mask

全名：bpy.context.edit_mask

说明：

类型：bpy.types.Mask

## 节点上下文 Node Context

## selected_nodes

全名：bpy.context.selected_nodes

说明：

类型：sequence of bpy.types.Node

## active_node

全名：bpy.context.active_node

说明：

类型：bpy.types.Node

## light

全名：bpy.context.light

说明：

类型：bpy.types.Light

## material

全名：bpy.context.material

说明：

类型：bpy.types.Material

## world

全名：bpy.context.world

说明：

类型：bpy.types.World

## 文字上下文 Text Context

## edit_text

全名：bpy.context.edit_text

说明：

类型：bpy.types.Text

Clip Context

## edit_movieclip

全名：bpy.context.edit_movieclip

说明：

类型：bpy.types.MovieClip

## edit_mask

全名：bpy.context.edit_mask

说明：

类型：bpy.types.Mask

## Sequencer Context

## edit_mask

全名：bpy.context.edit_mask

说明：

类型：bpy.types.Mask

## 文件上下文 File Context

## active_file

全名：bpy.context.active_file

说明：

类型：bpy.types.FileSelectEntry

## selected_files

全名：bpy.context.selected_files

说明：

类型：sequence of bpy.types.FileSelectEntry

## asset_library_ref

全名：bpy.context.asset_library_ref

说明：

类型：bpy.types.AssetLibraryReference

## selected_asset_files

全名：bpy.context.selected_asset_files

说明：

类型：sequence of bpy.types.FileSelectEntry

## id

全名：bpy.context.id

说明：

类型：bpy.types.ID
