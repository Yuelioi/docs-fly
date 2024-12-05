---
title: chprim_keyvalues
order: 8
---
`float [] chprim_keyvalues(<geometry>geometry, int prim)`

Returns an array of the values of each key in the channel primitive.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`prim`

The primitive number of the channel primitive.
