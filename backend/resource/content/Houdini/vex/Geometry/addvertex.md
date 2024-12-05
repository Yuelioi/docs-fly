---
title: addvertex
order: 3
---
`int  addvertex(int geohandle, int prim_num, int point_num)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`prim_num`

The primitive number to add the vertex to.

`point_num`

The point number to wire the new vertex to.

Returns

Returns a *linear* vertex index, or `-1` if the vertex could not be added. You can use this number with [setvertexattrib](setvertexattrib.html "Sets a vertex attribute in a geometry.") to set attributes on the new vertex, however this number may not be the final vertex index.
