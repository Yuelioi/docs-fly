---
title: 数据访问(bpy.data)
order: 7
---

# 数据访问(bpy.data)

## Description

该模块用于所有 Blender/Python 访问

详见：[bpy.types.BlendData](https://docs.blender.org/api/master/bpy.types.BlendData.html)

actions，armatures，as_pointer()，batch_remove()，bl_rna，bl_rna_get_subclass()，bl_rna_get_subclass_py()，brushes，cache_files，cameras，collections，curves，driver_add()，driver_remove()，filepath，fonts，get()，grease_pencils，id_data，id_properties_clear()，id_properties_ensure()，id_properties_ui()，images，is_dirty，is_property_hidden()，is_property_overridable_library()，is_property_readonly()，is_property_set()，is_saved，items()，keyframe_delete()，keyframe_insert()，keys()，lattices，libraries，lightprobes，lights，linestyles，masks，materials，meshes，metaballs，movieclips，node_groups，objects，orphans_purge()，paint_curves，palettes，particles，path_from_id()，path_resolve()，pop()，property_overridable_library_set()，property_unset()，rna_type，scenes，screens，shape_keys，sounds，speakers，temp_data()，texts，textures，type_recast()，use_autopack，user_map()，values()，version，volumes，window_managers，workspaces，worlds

### bpy.data.data

说明：访问 Blender 内部数据

类型：`bpy.types.BlendData`

示例：

```python
import bpy


# 打印所有对象
for obj in bpy.data.objects:
    print(obj.name)


# 打印场景名称列表
print(bpy.data.scenes.keys())


# 移除网格数组中的“Cube”
if "Cube" in bpy.data.meshes:
    mesh = bpy.data.meshes["Cube"]
    print("removing mesh", mesh)
    bpy.data.meshes.remove(mesh)


# 将图像写入bl文件旁
import os
with open(os.path.splitext(bpy.data.filepath)[0] + ".txt", 'w') as fs:
    for image in bpy.data.images:
        fs.write("%s %d x %d\n" % (image.filepath, image.size[0], image.size[1]))

```
