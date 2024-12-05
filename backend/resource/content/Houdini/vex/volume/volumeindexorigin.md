---
title: volumeindexorigin
order: 8
---
`vector  volumeindexorigin(<geometry>geometry, int primnum)`

`vector  volumeindexorigin(<geometry>geometry, string volumename)`

Show/hide arguments

Returns

The index of the bottom left of a volume primitive.
For Volume primitives, this is always zero. However, for VDB primitives,
this represents the bottom left of their active bounding box of voxels.

Returns 0 if `primnum` is out of range, the geometry is invalid, or the given primitive is not a volume primitive.
