---
title: setpointtransform
order: 69
---
| Since | 18.5 |
| --- | --- |

`int  setpointtransform(int geohandle, int pt, matrix transform)`

`int  setpointtransform(int geohandle, int pt, matrix transform, int constrain)`

Sets the `v@P` and the `3@transform` attributes on the given point from the given 4×4 matrix.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`pt`

The point index to modify.

`transform`

The 4×4 transform.

`constrain`

When True, update the children point transforms when modifying a point world
transform. When False, the children points stay in place like when using
Child Compensation on a transform handle. Defaults to True;
