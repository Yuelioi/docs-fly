---
title: quaterniontoeuler
order: 8
---
| Since | 17.0 |
| --- | --- |

`vector  quaterniontoeuler(vector4 orient, int order)`

Creates a vector representing euler angles from a vector4 representing a quaternion.

The angles are in radians. Use the `degrees()` function to convert radians into degrees.

For more information, see [Data types](../lang.html#data-types) and [Dot operator](../lang.html#dot-operator).

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
