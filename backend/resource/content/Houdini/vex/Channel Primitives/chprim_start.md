---
title: chprim_start
order: 13
---
`float  chprim_start(<geometry>geometry, int prim)`

Returns the start time in seconds of a channel primitive.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`prim`

The primitive number of the channel primitive.
