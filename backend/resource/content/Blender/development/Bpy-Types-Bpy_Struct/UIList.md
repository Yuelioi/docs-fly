---
title: UIList
order: 13
---

# UIList

## Description

[UIList(bpy_struct)](https://docs.blender.org/api/master/bpy.types.UIList.html)

[UI_UL_list(UIList)](https://docs.blender.org/api/master/bpy.types.UI_UL_list.html)

## UIList(bpy_struct)

基类： bpy_struct

子类：ASSETBROWSER_UL_metadata_tags, CLIP_UL_tracking_objects,FILEBROWSER_UL_dir, GPENCIL_UL_annotation_layer, GPENCIL_UL_layer,GPENCIL_UL_masks, GPENCIL_UL_matslots, GPENCIL_UL_vgroups, HAIR_UL_attributes,IMAGE_UL_render_slots, IMAGE_UL_udim_tiles, MASK_UL_layers,MATERIAL_UL_matslots, MESH_UL_attributes, MESH_UL_fmaps, MESH_UL_shape_keys,MESH_UL_uvmaps, MESH_UL_vcols, MESH_UL_vgroups, NODE_UL_interface_sockets,PARTICLE_UL_particle_systems, PHYSICS_UL_dynapaint_surfaces,POINTCLOUD_UL_attributes, RENDER_UL_renderviews, SCENE_UL_keying_set_paths,TEXTURE_UL_texpaintslots, TEXTURE_UL_texslots, UI_UL_list, VIEWLAYER_UL_aov,VIEWLAYER_UL_linesets, VOLUME_UL_grids

### 基本 UIList 示例

这个脚本是 UIList 子类，显示素材槽，有一堆额外的注释。

注意类的名字，命名规则与面板或菜单的命名规则类似。

::: tip
UIList 子类在使用前必须注册
:::

![](https://cdn.yuelili.com/20220116013512.png)

```python
import bpy


class MATERIAL_UL_matslots_example(bpy.types.UIList):

    # draw_item函数：列表中每个可见集合项目都会被调用。
    # data：集合包含的 RNA 对象。
    # item：集合中当前绘制的项目。
    # icon：该项目的 "计算 "图标(整数，因为有些对象如材料或纹理有自定义的图标ID，这些图标不能作为枚举项)。
    # active_data：包含在集合内活动属性的RNA对象（即指向集合的活动项目的整数）。
    # active_propname：活动属性的名称(使用'getattr(active_data, active_propname)')。
    # index：集合中当前项目的索引。
    # flt_flag：本项目的过滤过程的结果。
    # 注意：index和flt_flag是可选参数

    def draw_item(self, context, layout, data, item, icon, active_data, active_propname):
        ob = data
        slot = item
        ma = slot.material
        # draw_item必须处理三种布局类型... 通常'DEFAULT'和'COMPACT'可以共享同一个代码。

        if self.layout_type in {'DEFAULT', 'COMPACT'}:
            # 开始行布局：以标签（图标+文本）或非浮雕的文本字段。
            # 这样使该行在列表中容易被选择! 后者还可以让通过ctrl点击重命名。
            # 使用标签的icon_value，因为给定的图标是一个整数值，而不是一个枚举ID。
            # 注意："data "名称不能被翻译!
            if ma:
                layout.prop(ma, "name", text="", emboss=False, icon_value=icon)
            else:
                layout.label(text="", translate=False, icon_value=icon)

        # 'GRID'布局类型应尽可能紧凑（通常是一个单一的图标！）。
        elif self.layout_type in {'GRID'}:
            layout.alignment = 'CENTER'
            layout.label(text="", icon_value=icon)


# 现在可以在Blender的任何地方使用这个列表。这有一个小的面板示例。
class UIListPanelExample1(bpy.types.Panel):
    """在对象属性窗口中创建一个面板"""
    bl_label = "UIList 测试 1"
    bl_idname = "OBJECT_PT_ui_list_example_1"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout

        obj = context.object

        # template_list 现在需要两个新参数。
        # 第一个参数：注册UIList的要用的标识符（如果你只要默认列表，没有自定义的绘制代码，使用 "UI_UL_list"）。
        layout.template_list("MATERIAL_UL_matslots_example", "",
                             obj, "material_slots", obj, "active_material_index")

        # 第二个参数：通常可以保留为空字符串。
        # 这是个额外的ID，用于区分列表，以防你在某个区域多次使用同一个列表。
        layout.template_list("MATERIAL_UL_matslots_example", "compact", obj, "material_slots",
                             obj, "active_material_index", type='COMPACT')


def register():
    bpy.utils.register_class(MATERIAL_UL_matslots_example)
    bpy.utils.register_class(UIListPanelExample1)


def unregister():
    bpy.utils.unregister_class(UIListPanelExample1)
    bpy.utils.unregister_class(MATERIAL_UL_matslots_example)


if __name__ == "__main__":
    register()


```

### 进阶 UIList 实例：过滤和重排序

本脚本是 UIList 子类的扩展版本，用于显示顶点组。  
没有 "原封不动 "地使用，因为在 "draw
"函数中遍历所有顶点对 UI 性能来说是一个非常糟糕的主意！但它是个很好示例，方便说明如何创建/使用过滤/排序回调。

![](https://cdn.yuelili.com/20220116015824.png)

```python
import bpy


class MESH_UL_vgroups_slow(bpy.types.UIList):
    # 常量（flags）。
    # 注意不要对 FILTER_ITEM 产生阴影!
    VGROUP_EMPTY = 1 << 0

    # 自定义属性，与.blend文件一起保存。
    use_filter_empty: bpy.props.BoolProperty(
        name="Filter Empty",
        default=False,
        options=set(),
        description="Whether to filter empty vertex groups",
    )
    use_filter_empty_reverse: bpy.props.BoolProperty(
        name="Reverse Empty",
        default=False,
        options=set(),
        description="Reverse empty filtering",
    )
    use_filter_name_reverse: bpy.props.BoolProperty(
        name="Reverse Name",
        default=False,
        options=set(),
        description="Reverse name filtering",
    )

    # 这允许我们有相互排斥的选项，这些选项也可禁用
    def _gen_order_update(name1, name2):
        def _u(self, ctxt):
            if (getattr(self, name1)):
                setattr(self, name2, False)
        return _u
    use_order_name: bpy.props.BoolProperty(
        name="Name", default=False, options=set(),
        description="Sort groups by their name (case-insensitive)",
        update=_gen_order_update("use_order_name", "use_order_importance"),
    )
    use_order_importance: bpy.props.BoolProperty(
        name="Importance",
        default=False,
        options=set(),
        description="Sort groups by their average weight in the mesh",
        update=_gen_order_update("use_order_importance", "use_order_name"),
    )

    # 通常的绘制项目功能。
    def draw_item(self, context, layout, data, item, icon, active_data, active_propname, index, flt_flag):
        # 以防万一，不在这用
        self.use_filter_invert = False

        # 申明(isinstance(item, bpy.types.VertexGroup)
        vgroup = item
        if self.layout_type in {'DEFAULT', 'COMPACT'}:
            # 这里使用新过滤功能的一个特点：它可以通过flt_flag参数向draw_item传递数据，
            # 这个参数正好包含了filter_items在其过滤列表中为这个项目所设置的内容!
            # 在这种情况下，空组显示为灰色。
            if flt_flag & self.VGROUP_EMPTY:
                col = layout.column()
                col.enabled = False
                col.alignment = 'LEFT'
                col.prop(vgroup, "name", text="", emboss=False, icon_value=icon)
            else:
                layout.prop(vgroup, "name", text="", emboss=False, icon_value=icon)
            icon = 'LOCKED' if vgroup.lock_weight else 'UNLOCKED'
            layout.prop(vgroup, "lock_weight", text="", icon=icon, emboss=False)
        elif self.layout_type in {'GRID'}:
            layout.alignment = 'CENTER'
            if flt_flag & self.VGROUP_EMPTY:
                layout.enabled = False
            layout.label(text="", icon_value=icon)

    def draw_filter(self, context, layout):
        # 普通的UI绘制代码
        row = layout.row()

        subrow = row.row(align=True)
        subrow.prop(self, "filter_name", text="")
        icon = 'ZOOM_OUT' if self.use_filter_name_reverse else 'ZOOM_IN'
        subrow.prop(self, "use_filter_name_reverse", text="", icon=icon)

        subrow = row.row(align=True)
        subrow.prop(self, "use_filter_empty", toggle=True)
        icon = 'ZOOM_OUT' if self.use_filter_empty_reverse else 'ZOOM_IN'
        subrow.prop(self, "use_filter_empty_reverse", text="", icon=icon)

        row = layout.row(align=True)
        row.label(text="Order by:")
        row.prop(self, "use_order_name", toggle=True)
        row.prop(self, "use_order_importance", toggle=True)
        icon = 'TRIA_UP' if self.use_filter_orderby_invert else 'TRIA_DOWN'
        row.prop(self, "use_filter_orderby_invert", text="", icon=icon)

    def filter_items_empty_vgroups(self, context, vgroups):
        # 辅助函数：检查vgroup，查找是否为空，以及平均权重是多少。
        # TODO: 实际上应该是RNA助手（一个vgroup prop，如 "raw_data: ((vidx, vweight), 等等.)"）。
        # 对于python来说太慢了!
        obj_data = context.active_object.data
        ret = {vg.index: [True, 0.0] for vg in vgroups}
        if hasattr(obj_data, "vertices"):  # Mesh data
            if obj_data.is_editmode:
                import bmesh
                bm = bmesh.from_edit_mesh(obj_data)
                # 只有一个变形权重图层
                dvert_lay = bm.verts.layers.deform.active
                fact = 1 / len(bm.verts)
                if dvert_lay:
                    for v in bm.verts:
                        for vg_idx, vg_weight in v[dvert_lay].items():
                            ret[vg_idx][0] = False
                            ret[vg_idx][1] += vg_weight * fact
            else:
                fact = 1 / len(obj_data.vertices)
                for v in obj_data.vertices:
                    for vg in v.groups:
                        ret[vg.group][0] = False
                        ret[vg.group][1] += vg.weight * fact
        elif hasattr(obj_data, "points"):  # Lattice 数据
            # XXX 没有访问lattice editdata的权限?
            fact = 1 / len(obj_data.points)
            for v in obj_data.points:
                for vg in v.groups:
                    ret[vg.group][0] = False
                    ret[vg.group][1] += vg.weight * fact
        return ret

    def filter_items(self, context, data, propname):
        # 此函数获得集合属性( 元组(data, propname) )，并且必须返回两个列表。
        # 第一个用来过滤，必须包含32位整数，其中self.bitflag_filter_item标志着匹配的项目被过滤（即被显示），其他31位看需求设置。这里我们用第一个来标记VGROUP_EMPTY。
        # * 第二个用来重新排序，必须返回一个包含项目的新索引的列表（给了我们一个org_idx -> new_idx的映射）。
        # 请注意，默认的UI_UL_list为常见的任务定义了辅助函数(更多信息请看它的文档)。
        # 如果你不做过滤 和/或 排序，返回空列表(比返回doing nothing的完整列表更有效率！)。
        vgroups = getattr(data, propname)
        helper_funcs = bpy.types.UI_UL_list

        # 默认返回值
        flt_flags = []
        flt_neworder = []

        # Pre-compute of vgroups data, CPU-intensive. :/
        vgroups_empty = self.filter_items_empty_vgroups(context, vgroups)

        # Filtering by name
        if self.filter_name:
            flt_flags = helper_funcs.filter_items_by_name(self.filter_name, self.bitflag_filter_item, vgroups, "name",
                                                          reverse=self.use_filter_name_reverse)
        if not flt_flags:
            flt_flags = [self.bitflag_filter_item] * len(vgroups)

        # Filter by emptiness.
        for idx, vg in enumerate(vgroups):
            if vgroups_empty[vg.index][0]:
                flt_flags[idx] |= self.VGROUP_EMPTY
                if self.use_filter_empty and self.use_filter_empty_reverse:
                    flt_flags[idx] &= ~self.bitflag_filter_item
            elif self.use_filter_empty and not self.use_filter_empty_reverse:
                flt_flags[idx] &= ~self.bitflag_filter_item

        # Reorder by name or average weight.
        if self.use_order_name:
            flt_neworder = helper_funcs.sort_items_by_name(vgroups, "name")
        elif self.use_order_importance:
            _sort = [(idx, vgroups_empty[vg.index][1]) for idx, vg in enumerate(vgroups)]
            flt_neworder = helper_funcs.sort_items_helper(_sort, lambda e: e[1], True)

        return flt_flags, flt_neworder


# 使用上述UIList的最少代码...
class UIListPanelExample2(bpy.types.Panel):
    """在对象属性窗口中创建面板"""
    bl_label = "UIList Example 2 Panel"
    bl_idname = "OBJECT_PT_ui_list_example_2"
    bl_space_type = 'PROPERTIES'
    bl_region_type = 'WINDOW'
    bl_context = "object"

    def draw(self, context):
        layout = self.layout
        obj = context.object

        # template_list 现在需要两个新参数。
        # # 第一个参数：注册UIList的要用的标识符（如果你只要默认列表，没有自定义的绘制代码，使用 "UI_UL_list"）。
        layout.template_list("MESH_UL_vgroups_slow", "", obj, "vertex_groups", obj.vertex_groups, "active_index")


def register():
    bpy.utils.register_class(MESH_UL_vgroups_slow)
    bpy.utils.register_class(UIListPanelExample2)


def unregister():
    bpy.utils.unregister_class(UIListPanelExample2)
    bpy.utils.unregister_class(MESH_UL_vgroups_slow)


if __name__ == "__main__":
    register()

```

### classbpy.types.UIList()

全名：classbpy.types.UIList(bpy_struct)

说明：包含一个集合的元素的 UI 列表

### bitflag_filter_item

说明：保留的 bitflag（位标志） 'FILTER_ITEM'的值（在 filter_flags 值中）。

类型：int in [0, inf], default 0, 只读

### bl_idname

说明：自定义 ID，不设置则使用类名

类型：字符串， default “”, 不会为 None

### filter_name

说明：只显示与此名称匹配的项目（使用'\*'作为通配符）。

类型：字符串，默认为"", 不会为 None

### layout_type

说明：布局方式

类型：

- DEFAULT：默认布局。使用默认的多行布局。
- COMPACT：紧凑型布局。使用紧凑型的单行布局。
- GRID：网格布局。使用基于网格的布局。

### list_id

说明：列表的标识符，如果给 "template_list() "传递 "list_id "参数的话。

类型：字符串， 默认"", 只读，可选

### use_filter_invert

说明：反向过滤（显示隐藏项目，反之亦然）。

类型：布尔型，默认为 False

### use_filter_show

说明：显示过滤选项

类型：布尔型，默认为 False

### use_filter_sort_alpha

说明：按项目名称排序

类型：布尔型，默认为 False

### use_filter_sort_lock

说明：锁定显示项目的顺序（用户不能改变）。

类型：布尔型，默认为 False

### use_filter_sort_reverse

说明：将显示的项目顺序颠倒过来

类型：布尔型，默认为 False

### draw_item()

全名：draw_item(context, layout, data, item, icon, active_data, active_property,
index=0, flt_flag=0)

说明：在一个列表中绘制元素

NOTE：绘制自己的 draw_item 函数时，请保证 item 类型 正确

参数：

- layout： UILayout, 不会为 None，用于绘制元素的布局。
- data：AnyType，包含在集合属性中的对象（RNA 对象）。
- item：AnyType，集合(collection)中当前绘制的项目(item) 。
- icon：整数 [0, inf]，集合中项目的图标。如材料或纹理有自定义的图标 ID。
- active_data：AnyType, 不会为 None，包含在集合内活动属性的 RNA 对象。
- active_property：字符串，可选，Identifier of property in active_data, for the active element。
- index： 整数 [0, inf]，可选，集合中当前项目的索引。
- flt_flag： 整数 [0, inf]，可选，The filter-flag result for this item 。

### draw_filter()

全名：draw_filter(context, layout)

说明：绘制过滤选项

参数：layout，UILayout, 不会为 None 。用于绘制元素的布局。

### filter_items()

全名：filter_items(context, data, property)

说明：过滤 和/或 重新排序集合中的项目（在 filter_flags 中输出过滤结果，在 filter_neworder 中输出重新排序结果）。

参数：

- data：AnyType 。从集合属性获取的数据
- property ：字符串， 不会为 None。数据中属性的标识符，用于集合。

Return ：(filter_flags, filter_neworder)

- filter_flags：过滤标志的数组，集合中的每个项目都有一个（注意：FILTER_ITEM 位是保留的，它决定项目是否被显示），[0, inf]。
- filter_neworder： 索引数组，集合中的每个项目都有一个索引，将 org 索引映射到新的索引， [0, inf]

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id (字符串) 。The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

## UI_UL_list(UIList)

基类：bpy_struct,UIList

### filter_items_by_name()

全名：static filter_items_by_name(pattern, bitflag, items, propname='name',flags=None, reverse=False)

说明：设置 FILTER_ITEM，用于与 filter_name 匹配的项目（不区分大小写）。

参数：

- pattern：过滤模式。
- propname：用于过滤的字符串属性名称。
- flags：必须是与项目相同长度的整数列表，或者是 None！

返回：一个标志列表（如果不是 None，则基于给定的标志），如果没有给定的标志，则返回一个空列表，不进行过滤。

### static sort_items_helper()

全名：static sort_items_helper ( sort_data , key , reverse = False )

说明：常见的排序工具。

参数：

- sort_data：一个（无序的）图元列表[(org_idx, ...), (org_idx, ...), ...].
- key：用于 sorted()内置函数的同种可调用文件。

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass ( id ,默认= None )

参数：id ( string ) -- RNA 类型标识符。

返回：未找到时的 RNA 类型或默认值。

返回类型：bpy.types.Struct 子类

### classmethod bl_rna_get_subclass_py()

全名：bl_rna_get_subclass_py ( id ,默认= None )

参数：id ( string ) -- RNA 类型标识符。

返回：类，未找到时为默认。

返回类型：type
