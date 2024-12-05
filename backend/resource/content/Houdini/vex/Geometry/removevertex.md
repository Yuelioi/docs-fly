---
title: removevertex
order: 34
---
| Since | 18.0 |
| --- | --- |

`int  removevertex(int geohandle, int linear_vertex_index)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`linear_vertex_index`

If this is `-1`, the function has no effect. This is a linear
vertex index, so `vertexindex` may be needed to convert from
a primitive and vertex number.

This removes the given vertex from the geometry. Note: This is done as
a post process, not immediately when invoked.

Only polygons currently support the removal of vertices.

This can result in degenerate (0 vertex) polygons, as the primitive
is not deleted.

Removing many vertices from polygons with high vertex counts can be slow.
