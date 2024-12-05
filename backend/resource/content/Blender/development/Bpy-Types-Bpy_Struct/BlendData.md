---
title: BlendData
order: 4
---

# BlendData

## Description

[BlendData(bpy_struct)](https://docs.blender.org/api/master/bpy.types.BlendData.html)

基类：bpy_struct

说明：代表一个.blend 文件及其所有数据块的主要数据结构

数据详情，请在[子类中](https://www.yuelili.com/docs/blender-master/bpy-types/)查找

![](https://cdn.yuelili.com/20220115144246.png)

### actions

说明：Action data-blocks

类型：BlendDataActions bpy_prop_collection of Action, 只读

### armatures

说明：Armature data-blocks

类型：BlendDataArmatures bpy_prop_collection of Armature, 只读

### brushes 笔刷

说明：Brush data-blocks

类型：BlendDataBrushes bpy_prop_collection of Brush, 只读

### cache_files 缓存文件

说明：Cache Files data-blocks

类型：BlendDataCacheFiles bpy_prop_collection of CacheFile, 只读

### cameras 摄像机

说明：Camera data-blocks

类型：BlendDataCameras bpy_prop_collection of Camera, 只读

### collections 集合

说明：Collection data-blocks

类型：BlendDataCollections bpy_prop_collection of Collection, 只读

### curves 曲线

说明：Curve data-blocks

类型：BlendDataCurves bpy_prop_collection of Curve, 只读

### filepath 文件路径

说明：Path to the .blend file

类型：string, default “”, (readonly, ，可选)

### fonts 字体

说明：Vector font data-blocks

类型：BlendDataFonts bpy_prop_collection of VectorFont, 只读

### grease_pencils

说明：Grease Pencil data-blocks

类型：BlendDataGreasePencils bpy_prop_collection of GreasePencil, 只读

### hairs 毛发

说明：Hair data-blocks

类型：BlendDataHairs bpy_prop_collection of Hair, 只读

### images 图片

说明：Image data-blocks

类型：BlendDataImages bpy_prop_collection of Image, 只读

### is_dirty 是否修改并保存

说明：Have recent edits been saved to disk

类型：boolean, default False, 只读

### is_saved 是否保存

说明：Has the current session been saved to disk as a .blend file

类型：boolean, default False, 只读

### lattices 晶格

说明：Lattice data-blocks

类型：BlendDataLattices bpy_prop_collection of Lattice, 只读

### libraries

说明：Library data-blocks

类型：BlendDataLibraries bpy_prop_collection of Library, 只读

### lightprobes

说明：Light Probe data-blocks

类型：BlendDataProbes bpy_prop_collection of LightProbe, 只读

### lights 灯光

说明：Light data-blocks

类型：BlendDataLights bpy_prop_collection of Light, 只读

### linestyles

说明：Line Style data-blocks

类型：BlendDataLineStyles bpy_prop_collection of FreestyleLineStyle, 只读

### masks 蒙版

说明：Masks data-blocks

类型：BlendDataMasks bpy_prop_collection of Mask, 只读

### materials 材质

说明：Material data-blocks

类型：BlendDataMaterials bpy_prop_collection of Material, 只读

### meshes 网格

说明：Mesh data-blocks

类型：BlendDataMeshes bpy_prop_collection of Mesh, 只读

### metaballs 融球

说明：Metaball data-blocks

类型：BlendDataMetaBalls bpy_prop_collection of MetaBall, 只读

### movieclips

说明：Movie Clip data-blocks

类型：BlendDataMovieClips bpy_prop_collection of MovieClip, 只读

### node_groups

说明：Node group data-blocks

类型：BlendDataNodeTrees bpy_prop_collection of NodeTree, 只读

### objects 对象

说明：Object data-blocks

类型：BlendDataObjects bpy_prop_collection of Object, 只读

### paint_curves

说明：Paint Curves data-blocks

类型：BlendDataPaintCurves bpy_prop_collection of PaintCurve, 只读

### palettes

说明：Palette data-blocks

类型：BlendDataPalettes bpy_prop_collection of Palette, 只读

### particles 粒子

说明：Particle data-blocks

类型：BlendDataParticles bpy_prop_collection of ParticleSettings, 只读

### pointclouds

说明：Point cloud data-blocks

类型：BlendDataPointClouds bpy_prop_collection of PointCloud, 只读

### scenes 场景

说明：Scene data-blocks

类型：BlendDataScenes bpy_prop_collection of Scene, 只读

### screens 屏幕

说明：Screen data-blocks

类型：BlendDataScreens bpy_prop_collection of Screen, 只读

### shape_keys

说明：Shape Key data-blocks

类型：bpy_prop_collection of Key, 只读

### simulations 模拟

说明：Simulation data-blocks

类型：BlendDataSimulations bpy_prop_collection of Simulation, 只读

### sounds

说明：Sound data-blocks

类型：BlendDataSounds bpy_prop_collection of Sound, 只读

### speakers

说明：Speaker data-blocks

类型：BlendDataSpeakers bpy_prop_collection of Speaker, 只读

### texts

说明：Text data-blocks

类型：BlendDataTexts bpy_prop_collection of Text, 只读

### textures 纹理

说明：Texture data-blocks

类型：BlendDataTextures bpy_prop_collection of Texture, 只读

### use_autopack

说明：Automatically pack all external data into .blend file

类型：boolean, default False

### version 版本

说明：File format version the .blend file was saved with

类型：int array of 3 items in [0, inf], default (0, 0, 0), 只读

### volumes 体积

说明：Volume data-blocks

类型：BlendDataVolumes bpy_prop_collection of Volume, 只读

### window_managers

说明：Window manager data-blocks

类型：BlendDataWindowManagers bpy_prop_collection of WindowManager, 只读

### workspaces 工作空间

说明：Workspace data-blocks

类型：BlendDataWorkSpaces bpy_prop_collection of WorkSpace, 只读

### worlds

说明：World data-blocks

类型：BlendDataWorlds bpy_prop_collection of World, 只读

### batch_remove()

全名：batch_remove(ids)

Remove (delete) several IDs at once.

WARNING: Considered experimental feature currently.

Note that this function is quicker than individual calls to remove() (frombpy.types.BlendData ID collections), but less safe/versatile (it can breakBlender, e.g. by removing all scenes…).

参数：ids 。 Iterables of IDs (types can be mixed).

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type

### orphans_purge()

全名：orphans_purge()

Remove (delete) all IDs with no user.

参数：

do_local_ids (bool, 字符串) 。 Include unused local IDs in the deletion, defaults
to True

do_linked_ids (bool, 字符串) 。 Include unused linked IDs in the deletion,
defaults to True

do_recursive (bool, 字符串) 。 Recursively check for unused IDs, ensuring noorphaned one remain after a single run of that function, defaults to False

返回：The number of deleted IDs.

### temp_data()

全名：temp_data(filepath=None)

A context manager that temporarily creates blender file data.

参数：filepath (str or NoneType) 。 The file path for the newly temporary data.When None, the path of the currently open file is used.

返回：Blend file data which is freed once the context exists.

返回类型：bpy.types.BlendData

### user_map()

全名：user_map(subset, key_types, value_types)

Returns a mapping of all ID data-blocks in current bpy.data to a set of all
datablocks using them.

For list of valid set members for key_types & value_types, see:bpy.types.KeyingSetPath.id_type.

参数：

subset (sequence) 。 When passed, only these data-blocks and their users willbe included as keys/values in the map.

key_types (set of strings) 。 Filter the keys mapped by ID types.

value_types (set of strings) 。 Filter the values in the set by ID types.

返回：dictionary of bpy.types.ID instances, with sets of ID’s as their values.

返回类型：dict
