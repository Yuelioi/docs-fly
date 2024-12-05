---
title: setpointlocaltransforms
order: 68
---
| Since | 18.5 |
| --- | --- |

`int  setpointlocaltransforms(int geohandle, int pnts[], matrix transforms[])`

Sets an array of local transforms associated with the point indices. This function set the `4@localtransform` attribute.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`pnts`

The array of point indices to set.

`transforms`

The array of transforms to set.
