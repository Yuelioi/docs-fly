---
title: setprimgroup
order: 11
---
`int  setprimgroup(int geohandle, string name, int prim_num, int value, string mode="set")`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`name`

The name of the group to modify.

`prim_num`

The primitive number to add or remove from the group.

`value`

`1` to put the primitive in the group, `0` to remove the primitive from the group.
This is ignored if `mode` is `"toggle"`.

`mode`

Use `"set"` to set the primitive’s membership according to the `value`.
Use `"toggle"` to toggle the primitive’s membership, regardless of the `value`.
