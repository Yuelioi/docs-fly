---
title: ID
order: 1
---

# ID

## Description

[ID(bpy_struct)](https://docs.blender.org/api/master/bpy.types.ID.html)

基类： bpy_struct

子类：Action, Armature, Brush, CacheFile, Camera, Collection, Curve,FreestyleLineStyle, GreasePencil, Hair, Image, Key, Lattice, Library, Light,LightProbe, Mask, Material, Mesh, MetaBall, MovieClip, NodeTree, Object,PaintCurve, Palette, ParticleSettings, PointCloud, Scene, Screen, Simulation,Sound, Speaker, Text, Texture, VectorFont, Volume, WindowManager, WorkSpace,
World

说明：数据块（data-blocks）的基础类型，定义唯一的名称，从其他库中链接和垃圾收集。

### asset_data

说明：Additional data for an asset data-block

类型：AssetMetaData, 只读

### is_embedded_data

说明：This data-block is not an independent one, but is actually a sub-data ofanother ID (typical example: root node trees or master collections)

类型：boolean, default False, 只读

### is_evaluated

说明：Whether this ID is runtime-only, evaluated data-block, or actual data from
.blend file

类型：boolean, default False, 只读

### is_library_indirect

说明：Is this ID block linked indirectly

类型：boolean, default False, 只读

### library

说明：Library file the data-block is linked from

类型：Library, 只读

### library_weak_reference

说明：Weak reference to a data-block in another library .blend file (used to re-use already appended data instead of appending new copies)

类型：LibraryWeakReference, 只读

### name

说明：Unique data-block ID name

类型：string, default “”, 不会为 None

### name_full

说明：Unique data-block ID name, including library one is any

类型：string, default “”, (只读，可选)

### original

说明：Actual data-block from .blend file (Main database) that generated that
evaluated one

类型：ID, 只读

### override_library

说明：Library override data

类型：IDOverrideLibrary, 只读

### preview

说明：Preview image and icon of this data-block (always None if not supported for
this type of data)

类型：ImagePreview, 只读

### tag

说明：Tools can use this to tag data for their own purposes (initial state is
undefined)

类型：boolean, default False

### use_fake_user

说明：Save this data-block even if it has no users

类型：boolean, default False

### users

说明：Number of times this data-block is referenced

类型：int in [0, inf], default 0, 只读

### evaluated_get()

全名：evaluated_get(depsgraph)

Get corresponding evaluated ID from the given dependency graph

参数：depsgraph (Depsgraph, 不会为 None) 。 Dependency graph to perform lookup in

返回：New copy of the ID

返回类型：ID

### copy()

全名：copy()

Create a copy of this data-block (not supported for all data-blocks)

返回：New copy of the ID

返回类型：ID

### asset_mark()

全名：asset_mark()

Enable easier reuse of the data-block through the Asset Browser, with the helpof customizable metadata (like previews, descriptions and tags)

### asset_clear()

全名：asset_clear()

Delete all asset metadata and turn the asset data-block back into a normal
data-block

### asset_generate_preview()

全名：asset_generate_preview()

Generate preview image (might be scheduled in a background thread)

### override_create()

全名：override_create(remap_local_usages=False)

Create an overridden local copy of this linked data-block (not supported for
all data-blocks)

参数：remap_local_usages (boolean，可选) 。 Whether local usages of the linked IDshould be remapped to the new library override of it

返回：New overridden local copy of the ID

返回类型：ID

### override_hierarchy_create()

全名：override_hierarchy_create(scene, view_layer, reference=None)

Create an overridden local copy of this linked data-block, and most of itsdependencies when it is a Collection or and Object

参数：scene (Scene, 不会为 None) 。 In which scene the new overrides should be
instantiated

view_layer (ViewLayer, 不会为 None) 。 In which view layer the new overrides should
be instantiated

reference (ID，可选) 。 Another ID (usually an Object or Collection) used todecide where to instantiate the new overrides

返回：New overridden local copy of the root ID

返回类型：ID

### override_template_create()

全名：override_template_create()

Create an override template for this ID

### user_clear()

全名：user_clear()

Clear the user count of a data-block so its not saved, on reload the data will
be removed

This function is for advanced use only, misuse can crash blender since theuser count is used to prevent data being removed when it is used.

### user_remap()

全名：user_remap(new_id)

Replace all usage in the .blend file of this ID by new given one

参数：new_id (ID, 不会为 None) 。 New ID to use

### make_local()

全名：make_local(clear_proxy=True)

Make this datablock local, return local one (may be a copy of the original, incase it is also indirectly used)

参数：clear_proxy (boolean，可选) 。 Whether to clear proxies (the default behavior,note that if object has to be duplicated to be made local, proxies are always
cleared)

返回：This ID, or the new ID if it was copied

返回类型：ID

### user_of_id()

全名：user_of_id(id)

Count the number of times that ID uses/references given one

参数：id (ID, 不会为 None) 。 ID to count usages

返回：Number of usages/references of given id by current data-block

返回类型：int in [0, inf]

### animation_data_create()

全名：animation_data_create()

Create animation data to this ID, note that not all ID types support this

返回：New animation data or NULL

返回类型：AnimData

### animation_data_clear()

全名：animation_data_clear()

Clear animation on this this ID

### update_tag()

全名：update_tag(refresh={})

Tag the ID to update its display data, e.g. when calling
bpy.types.Scene.update

参数：refresh (enum set in {'OBJECT', 'DATA', 'TIME'}，可选) 。 Type of updates to
perform

### preview_ensure()

全名：preview_ensure()

Ensure that this ID has preview data (if ID type supports it)

返回：The existing or created preview

返回类型：ImagePreview

### classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

### classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id（字符串） 。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
