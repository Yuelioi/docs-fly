---
title: setvertexgroup
order: 12
---
`int  setvertexgroup(int geohandle, string name, int prim_num, int vertex_num, int value, string mode="set")`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

**To use a linear vertex index**, set the `prim_num` to the **linear vertex number** and set `vertex_num` to `-1`. Note that **this is different** from how most other vertex functions work.

Show/hide arguments

`name`

The name of the group to modify.

`prim_num`

The number of the primitive containing the vertex you want to add/remove.

`vertex_num`

The vertex offset on the primitive of the vertex you want to add/remove.

`value`

`1` to put the vertex in the group, `0` to remove the vertex from the group.
This is ignored if `mode` is `"toggle"`.

`mode`

Use `"set"` to set the vertex’s membership according to the `value`.
Use `"toggle"` to toggle the vertex’s membership, regardless of the `value`.
