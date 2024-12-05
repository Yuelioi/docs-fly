---
title: Mesh
order: 2
---

# Mesh

## Description

[Mesh(ID)](https://docs.blender.org/api/master/bpy.types.Mesh.html)

基类：bpy_struct, ID

说明：定义几何表面的网格数据块（data-block ）

### Mesh Data

网格数据（Mesh
Data）是以对象模式访问的，目的为了紧凑地存储，如果要从 python 中更灵活地编辑网格，请看[bmesh](https://docs.blender.org/api/master/bmesh.html#module-bmesh)。

Blender 存储了 4 个主要数组来定义网格的几何形状。

- [`Mesh.vertices`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.vertices "bpy.types.Mesh.vertices") ：顶点，空间中 3 点
- [`Mesh.edges`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.edges "bpy.types.Mesh.edges") ：边，引用两个顶点
- [`Mesh.loops`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.loops "bpy.types.Mesh.loops") ：循环，引用单个顶点和边
- [`Mesh.polygons`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.polygons "bpy.types.Mesh.polygons")：多边形，引用多个循环

每个多边形引用循环数组中一个切片（slice），这样一来，多边形不直接存储顶点或角（corner）数据，比如 UV，只引用多边形使用的循环。

[`Mesh.loops`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.loops "bpy.types.Mesh.loops"),
[`Mesh.uv_layers`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.uv_layers "bpy.types.Mesh.uv_layers")
[`Mesh.vertex_colors`](https://docs.blender.org/api/master/bpy.types.Mesh.html#bpy.types.Mesh.vertex_colors "bpy.types.Mesh.vertex_colors") 都是对齐的，所以相同的多边形循环索引可以寻找 UV 和顶点的颜色，就像顶点一样。

要比较网格 API 选项，请参见： [NGons and TessellationFaces](https://docs.blender.org/api/master/info_gotcha.html#info-gotcha-mesh-faces)

假设活动对象是个带 UV 的网格。本示例打印每个多边形的顶点和 UV

```python
import bpy

me = bpy.context.object.data
uv_layer = me.uv_layers.active.data


for poly in me.polygons:
    print("Polygon index: %d, length: %d" % (poly.index, poly.loop_total))

    # 在这里使用range为了展示多边形如何引用循环(loop),
    # 方便起见，也可以使用'poly.loop_indices'替代

    for loop_index in range(poly.loop_start, poly.loop_start + poly.loop_total):

        print("    Vertex: %d" % me.loops[loop_index].vertex_index)
        print("    UV: %r" % uv_layer[loop_index].uv)

```

## animation_data

说明：Animation data for this data-block

类型：AnimData, 只读

## attributes

说明：Geometry attributes

类型：AttributeGroup bpy_prop_collection of Attribute, 只读

## auto_smooth_angle

说明：Maximum angle between face normals that will be considered as smooth (unused if custom split normals data are available)

类型：float in [0, 3.14159], default 0.523599

## auto_texspace

说明：Adjust active object’s texture space automatically when transforming object

类型：boolean, default True

## cycles

说明：Cycles mesh settings

类型：CyclesMeshSettings, 只读

## edges

说明：Edges of the mesh

类型：MeshEdges bpy_prop_collection of MeshEdge, 只读

## face_maps

说明：类型：MeshFaceMapLayers bpy_prop_collection of MeshFaceMapLayer, 只读

## has_custom_normals

说明：True if there are custom split normals data in this mesh

类型：boolean, default False, 只读

## is_editmode

说明：True when used in editmode

类型：boolean, default False, 只读

## loop_triangles

说明：Tessellation of mesh polygons into triangles

类型：MeshLoopTriangles bpy_prop_collection of MeshLoopTriangle, 只读

## loops

说明：Loops of the mesh (polygon corners)

类型：MeshLoops bpy_prop_collection of MeshLoop, 只读

## materials

说明：类型：IDMaterials bpy_prop_collection of Material, 只读

## polygon_layers_float

说明：类型：PolygonFloatProperties bpy_prop_collection of MeshPolygonFloatPropertyLayer, 只读

## polygon_layers_int

说明：类型：PolygonIntProperties bpy_prop_collection of MeshPolygonIntPropertyLayer, 只读

## polygon_layers_string

说明：类型：PolygonStringProperties bpy_prop_collection of MeshPolygonStringPropertyLayer, 只读

## polygons

说明：Polygons of the mesh

类型：MeshPolygons bpy_prop_collection of MeshPolygon, 只读

## remesh_mode

说明：VOXEL Voxel 。 Use the voxel remesher.

QUAD Quad 。 Use the quad remesher.

类型：enum in [‘VOXEL’, ‘QUAD’], default ‘VOXEL’

## remesh_voxel_adaptivity

说明：Reduces the final face count by simplifying geometry where detail is not needed, generating triangles. A value greater than 0 disables Fix Poles

类型：float in [0, 1], default 0.0

## remesh_voxel_size

说明：Size of the voxel in object space used for volume evaluation. Lower values preserve finer details

类型：float in [0.0001, inf], default 0.1

## sculpt_vertex_colors

说明：All vertex colors

类型：VertColors bpy_prop_collection of MeshVertColorLayer, 只读

## shape_keys

说明：类型：Key, 只读

## skin_vertices

说明：All skin vertices

类型：bpy_prop_collection of MeshSkinVertexLayer, 只读

## texco_mesh

说明：Derive texture coordinates from another mesh

类型：Mesh

## texspace_location

说明：Texture space location

类型：float array of 3 items in [-inf, inf], default (0.0, 0.0, 0.0)

## texspace_size

说明：Texture space size

类型：float array of 3 items in [-inf, inf], default (1.0, 1.0, 1.0)

## texture_mesh

说明：Use another mesh for texture indices (vertex indices must be aligned)

类型：Mesh

## total_edge_sel

说明：Selected edge count in editmode

类型：int in [0, inf], default 0, 只读

## total_face_sel

说明：Selected face count in editmode

类型：int in [0, inf], default 0, 只读

## total_vert_sel

说明：Selected vertex count in editmode

类型：int in [0, inf], default 0, 只读

## use_auto_smooth

说明：Auto smooth (based on smooth/sharp faces/edges and angle between faces), or use custom split normals data if available

类型：boolean, default False

## use_auto_texspace

说明：Adjust active object’s texture space automatically when transforming object

类型：boolean, default True

## use_customdata_edge_bevel

说明：类型：boolean, default False

## use_customdata_edge_crease

说明：类型：boolean, default False

## use_customdata_vertex_bevel

说明：类型：boolean, default False

## use_mirror_topology

说明：Use topology based mirroring (for when both sides of mesh have matching, unique topology)

类型：boolean, default False

## use_mirror_vertex_groups

说明：Mirror the left/right vertex groups when painting. The symmetry axis is determined by the symmetry settings

类型：boolean, default True

## use_mirror_x

说明：Enable symmetry in the X axis

类型：boolean, default False

## use_mirror_y

说明：Enable symmetry in the Y axis

类型：boolean, default False

## use_mirror_z

说明：Enable symmetry in the Z axis

类型：boolean, default False

## use_paint_mask

说明：Face selection masking for painting

类型：boolean, default False

## use_paint_mask_vertex

说明：Vertex selection masking for painting

类型：boolean, default False

## use_remesh_fix_poles

说明：Produces less poles and a better topology flow

类型：boolean, default True

## use_remesh_preserve_paint_mask

说明：Keep the current mask on the new mesh

类型：boolean, default False

## use_remesh_preserve_sculpt_face_sets

说明：Keep the current Face Sets on the new mesh

类型：boolean, default False

## use_remesh_preserve_vertex_colors

说明：Keep the current vertex colors on the new mesh

类型：boolean, default False

## use_remesh_preserve_volume

说明：Projects the mesh to preserve the volume and details of the original mesh

类型：boolean, default True

## uv_layer_clone

说明：UV loop layer to be used as cloning source

类型：MeshUVLoopLayer

## uv_layer_clone_index

说明：Clone UV loop layer index

类型：int in [0, inf], default 0

## uv_layer_stencil

说明：UV loop layer to mask the painted area

类型：MeshUVLoopLayer

## uv_layer_stencil_index

说明：Mask UV loop layer index

类型：int in [0, inf], default 0

## uv_layers

说明：All UV loop layers

类型：UVLoopLayers bpy_prop_collection of MeshUVLoopLayer, 只读

## vertex_colors

说明：All vertex colors

类型：LoopColors bpy_prop_collection of MeshLoopColorLayer, 只读

## vertex_layers_float

说明：类型：VertexFloatProperties bpy_prop_collection of MeshVertexFloatPropertyLayer, 只读

## vertex_layers_int

说明：类型：VertexIntProperties bpy_prop_collection of MeshVertexIntPropertyLayer, 只读

## vertex_layers_string

说明：类型：VertexStringProperties bpy_prop_collection of MeshVertexStringPropertyLayer, 只读

## vertex_paint_masks

说明：Vertex paint mask

类型：bpy_prop_collection of MeshPaintMaskLayer, 只读

## vertices

说明：Vertices of the mesh

类型：MeshVertices bpy_prop_collection of MeshVertex, 只读

## edge_keys

说明：只读

## transform()

全名：transform(matrix, shape_keys=False)

说明：Transform mesh vertices by a matrix (Warning: inverts normals if matrix is negative)

参数：matrix (float multi-dimensional array of 4 \* 4 items in [-inf, inf]) 。 Matrix

shape_keys (boolean，可选) 。 Transform Shape Keys

## flip_normals()

全名：flip_normals()

说明：Invert winding of all polygons (clears tessellation, does not handle custom normals)

## calc_normals()

全名：calc_normals()

Calculate vertex normals

## create_normals_split()

全名：create_normals_split()

说明：Empty split vertex normals

## calc_normals_split()

全名：calc_normals_split()

说明：Calculate split vertex normals, which preserve sharp edges

## free_normals_split()

全名：free_normals_split()

说明：Free split vertex normals

## split_faces()

全名：split_faces(free_loop_normals=True)

说明：Split faces based on the edge angle

参数：free_loop_normals (boolean，可选) 。 Free Loop Normals, Free loop normals custom data layer

## calc_tangents()

全名：calc_tangents(uvmap='')

说明：Compute tangents and bitangent signs, to be used together with the split normals to get a complete tangent space for normal mapping (split normals are also computed if not yet present)

参数：uvmap (string, (可选，可选)) 。 Name of the UV map to use for tangent space computation

## free_tangents()

全名：free_tangents()

说明：Free tangents

## calc_loop_triangles()

全名：calc_loop_triangles()

说明：Calculate loop triangle tessellation (supports editmode too)

## calc_smooth_groups()

全名：calc_smooth_groups(use_bitflags=False)

说明：Calculate smooth groups from sharp edges

参数：use_bitflags，boolean，可选。Produce bitflags groups instead of simple numeric values

返回：(poly_groups, groups)

- poly_groups：Smooth Groups, int array of 1 items in [-inf, inf]

- groups：Total number of groups, int in [0, inf]

## normals_split_custom_set()

全名：normals_split_custom_set(normals)

Define custom split normals of this mesh (use zero-vectors to keep auto ones)

参数：normals (float multi-dimensional array of 1 \* 3 items in [-1, 1]) 。 Normals

## normals_split_custom_set_from_vertices()

全名：normals_split_custom_set_from_vertices(normals)

说明：Define custom split normals of this mesh, from vertices’ normals (use zero-vectors to keep auto ones)

参数：normals (float multi-dimensional array of 1 \* 3 items in [-1, 1]) 。 Normals

## update()

全名：update(calc_edges=False, calc_edges_loose=False)

说明：update

参数：

- calc_edges：boolean，可选。 Calculate Edges, Force recalculation of edges

- calc_edges_loose：boolean，可选。 Calculate Loose Edges, Calculate the loose state of each edge

## update_gpu_tag()

全名：update_gpu_tag()

update_gpu_tag

## unit_test_compare()

全名：unit_test_compare(mesh=None, threshold=7.1526e-06)

unit_test_compare

参数：mesh (Mesh，可选) 。 Mesh to compare to

threshold (float in [0, inf]，可选) 。 Threshold, Comparison tolerance threshold

返回：Return value, String description of result of comparison

返回类型：string, 不会为 None

## clear_geometry()

全名：clear_geometry()

Remove all geometry from the mesh. Note that this does not free shape keys or materials

## validate()

全名：validate(verbose=False, clean_customdata=True)

Validate geometry, return True when the mesh has had invalid geometry corrected/removed

参数：verbose (boolean，可选) 。 Verbose, Output information about the errors found

clean_customdata (boolean，可选) 。 Clean Custom Data, Remove temp/cached custom-data layers, like e.g. normals…

返回：Result

返回类型：boolean

## validate_material_indices()

全名：validate_material_indices()

Validate material indices of polygons, return True when the mesh has had invalid indices corrected (to default 0)

返回：Result

返回类型：boolean

## count_selected_items()

全名：count_selected_items()

Return the number of selected items (vert, edge, face)

返回：Result

返回类型：int array of 3 items in [0, inf]

## from_pydata()

全名：from_pydata(vertices, edges, faces)

Make a mesh from a list of vertices/edges/faces Until we have a nicer way to make geometry, use this.

参数：vertices (iterable object) 。 float triplets each representing (X, Y, Z) eg: [(0.0, 1.0, 0.5), …].

edges (iterable object) 。

int pairs, each pair contains two indices to the vertices argument. eg: [(1, 2), …]

When an empty iterable is passed in, the edges are inferred from the polygons.

faces (iterable object) 。 iterator of faces, each faces contains three or more indices to the vertices argument. eg: [(5, 6, 8, 9), (1, 2, 3), …]

Warning

Invalid mesh data (out of range indices, edges with matching indices, 2 sided faces… etc) are not prevented. If the data used for mesh creation isn’t known to be valid, run Mesh.validate after this function.

## classmethod bl_rna_get_subclass()

全名：classmethod bl_rna_get_subclass(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：RNA 类型，未找到则为默认。

返回类型：bpy.types.Struct subclass

## classmethod bl_rna_get_subclass_py()

全名：classmethod bl_rna_get_subclass_py(id, default=None)

参数：id(字符串)。 The RNA type identifier.

返回：类，未找到则为默认。

返回类型：type
