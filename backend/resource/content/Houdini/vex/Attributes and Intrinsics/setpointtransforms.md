---
title: setpointtransforms
order: 70
---
| Since | 18.5 |
| --- | --- |

`int  setpointtransforms(int geohandle, int pnts[], matrix transforms[])`

`int  setpointtransforms(int geohandle, int pnts[], matrix transforms[], int constrain)`

Sets an array of transforms associated with the point indices. This function set the `v@P` and the `3@transform` attributes.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`pnts`

The array of point indices to set.

`transforms`

The array of transforms to set.

`constrain`

When True, update the children point transforms when modifying a point world
transform. When False, the children points stay in place like when using
Child Compensation on a transform handle. Defaults to True;
