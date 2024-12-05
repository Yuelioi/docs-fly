---
title: setpointgroup
order: 10
---
`int  setpointgroup(int geohandle, string name, int point_num, int value, string mode="set")`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`name`

The name of the group to modify.

`point_num`

The point number to add or remove from the group.

`value`

`1` to put the point in the group, `0` to remove the point from the group.
This is ignored if `mode` is `"toggle"`.

`mode`

Use `"set"` to set the point’s membership according to the `value`.
Use `"toggle"` to toggle the point’s membership, regardless of the `value`.
