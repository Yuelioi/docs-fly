---
title: volumepostoindex
order: 13
---
`vector  volumepostoindex(<geometry>geometry, int primnum, vector position)`

`vector  volumepostoindex(<geometry>geometry, string volumename, vector position)`

Show/hide arguments

Returns

The index of a voxel at the given position.

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a vector volume primitive.
