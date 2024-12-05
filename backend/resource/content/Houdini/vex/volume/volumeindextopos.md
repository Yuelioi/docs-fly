---
title: volumeindextopos
order: 10
---
`vector  volumeindextopos(<geometry>geometry, int primnum, vector voxel)`

`vector  volumeindextopos(<geometry>geometry, string volumename, vector voxel)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The position of the center of a voxel.

Returns 0 if `primnum` is out of range, the geometry is invalid, or the given primitive is not a volume primitive.
