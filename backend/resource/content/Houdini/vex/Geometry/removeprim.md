---
title: removeprim
order: 33
---
`int  removeprim(int geohandle, int prim_number, int andpoints)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`prim_number`

The index of the primitive to remove. If this is `-1`, the function does nothing.

`andpoints`

If this is `1`, the function will also delete any points associated with the primitive that are not associated with any other primitives.

Note
If some primitives are being removed with `andpoints` set to `0` and some are being removed with `andpoints` set to `1`, all of the primitives with `andpoints` set to `0` will be deleted before all of the primitives with `andpoints` set to `1`.
