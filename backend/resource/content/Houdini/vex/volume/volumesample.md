---
title: volumesample
order: 15
---
`float  volumesample(<geometry>geometry, int primnum, vector pos)`

`float  volumesample(<geometry>geometry, string volumename, vector pos)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The volume primitive’s sampled value at the given position. Values between voxels will be trilinearly interpolated.

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a volume or vdb primitive.

![](../../images/vex/volumesample.png)
Example of interpolation of one and two dimensional data using `volumesample`. The visualized normal is computed using the `volumegradient` function.
