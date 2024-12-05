---
title: setagentlocaltransforms
order: 61
---
`void  setagentlocaltransforms(int geohandle, int prim, matrix transforms[])`

When modifying a single transform, using [setagentlocaltransform](setagentlocaltransform.html "Overrides the local space transform of an agent primitive’s bone.") instead can be significantly faster.

Show/hide arguments

`geohandle`

Handle to the geometry to write to. `geoself()` can be used to get a handle to the current geometry.

`prim`

The primitive number.

`transforms`

The new transform (in local space) of each bone in the agent’s rig.
