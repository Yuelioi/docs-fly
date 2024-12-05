---
title: volumeres
order: 14
---
`vector  volumeres(<geometry>geometry, int primnum)`

`vector  volumeres(<geometry>geometry, string volumename)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The resolution of the volume primitive.

For VDBs, the valid range of indices is not `0..res-1`, but
`volumeindexorigin..volumeindexorigin+res-1`

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a vector volume primitive.
