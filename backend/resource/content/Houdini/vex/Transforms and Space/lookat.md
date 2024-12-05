---
title: lookat
order: 6
---
`matrix3  lookat(vector from, vector to)`

`matrix3  lookat(vector from, vector to, float roll)`

`matrix3  lookat(vector from, vector to, vector up)`

`vector  lookat(vector from, vector to, float roll, int xyz)`

`vector  lookat(vector from, vector to, vector up, int xyz)`

Computes a rotation matrix or angles to orient the negative z-axis along the
vector (to-from) under the transformation. If an up vector is specified, this
will determine the roll.

Show/hide arguments

`xyz`

One of the rotation order constants listed below, which can be imported from `$HFS/houdini/vex/include/math.h`.

| Constant name | Rotation Order |
| --- | --- |
| XFORM_XYZ | Rotate order X, Y, Z |
| XFORM_XZY | Rotate order X, Z, Y |
| XFORM_YXZ | Rotate order Y, X, Z |
| XFORM_YZX | Rotate order Y, Z, X |
| XFORM_ZXY | Rotate order Z, X, Y |
| XFORM_ZYX | Rotate order Z, Y, X |
