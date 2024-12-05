---
title: volumevoxeldiameter
order: 23
---
`float  volumevoxeldiameter(<geometry>geometry, int primnum)`

`float  volumevoxeldiameter(<geometry>geometry, string primname)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The diameter of a voxel in the given primitive.
To find the length of a side of a voxel, divide by `sqrt(3)`.

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a vector volume primitive.
