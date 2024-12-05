---
title: removepoint
order: 32
---
`int  removepoint(int geohandle, int point_number)`

`int  removepoint(int geohandle, int point_number, int and_prims)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`point_number`

If this is `-1`, the function has no effect.

`and_prims`

If this is `1`, the function deletes any *degenerate* primitives that referred to the removed point (for example, closed polygons with fewer than 3 vertices or open polygons with fewer than 2 vertices).

If this is `0`, the function only deletes primitives that become invalid because of the removed point (for example, tetrahedra with fewer than 4 vertices, or volumes with zero vertices).
