---
title: smoothrotation
order: 25
---
`vector  smoothrotation(int order, vector r, vector r_reference)`

Returns the Euler rotations that have the closest values to r_reference while still describing the same orientation as r.
Typically, r_reference will be the rotations from the previous sample or frame.

The angles are in radians. Use the `radians()` function to convert degrees into radians.

Show/hide arguments

`order`

One of the rotation order constants listed below, which can be imported from `$HFS/houdini/vex/include/math.h`.

| Constant name | Rotation Order |
| --- | --- |
| XFORM_XYZ | Rotate order X, Y, Z |
| XFORM_XZY | Rotate order X, Z, Y |
| XFORM_YXZ | Rotate order Y, X, Z |
| XFORM_YZX | Rotate order Y, Z, X |
| XFORM_ZXY | Rotate order Z, X, Y |
| XFORM_ZYX | Rotate order Z, Y, X |
