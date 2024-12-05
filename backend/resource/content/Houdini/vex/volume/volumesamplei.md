---
title: volumesamplei
order: 16
---
`int  volumesamplei(<geometry>geometry, int primnum, vector pos)`

`int  volumesamplei(<geometry>geometry, string volumename, vector pos)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The volume primitive’s sampled value at the given position.

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a volume or vdb primitive.
