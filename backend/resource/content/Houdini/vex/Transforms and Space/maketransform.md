---
title: maketransform
order: 7
---
`matrix3  maketransform(vector zaxis, vector yaxis)`

`matrix3  maketransform(int xyz, vector angles)`

`matrix  maketransform(vector zaxis, vector yaxis, vector translate)`

`matrix  maketransform(int trs, int xyz, vector t, vector r)`

`matrix  maketransform(int trs, int xyz, vector t, vector r, vector s)`

`matrix  maketransform(int trs, int xyz, vector t, vector r, vector s, vector p)`

`matrix  maketransform(int trs, int xyz, vector t, vector r, vector s, vector p, vector pr)`

`matrix  maketransform(int trs, int xyz, vector t, vector r, vector s, vector p, vector pr, vector shears)`

Builds a 3×3 or 4×4 transform matrix.

`maketransform(int trs, ...)` builds a general 4×4 transform matrix
given an order of transformations (trs), an order for rotations
(xyz), a vector representing the translation (t), rotation
(r), scale (s) (and optionally a pivot (p), pivot rotation
(pr), and shears (shears)).

`maketransform(int xyz, vector angles)` builds a 3×3 rotation matrix
using the same rules as `maketransform(int trs, ...)` but only using
the rotation parameters.

`maketransform(vector zaxis, yaxis, ...)` builds either a 3×3 transform
matrix or a 4×4 transform matrix. The matrix will be constructed so that the
z-axis will be transformed to the z-axis specified with the given up vector
(yaxis). Thus, maketransform({0,0,1}, {0,1,0}) will result in an identity
matrix. The version which returns a 4×4 transform will apply the translation
to the 4×4 matrix. This function is very similar to the [lookat](lookat.html "Computes a rotation matrix or angles to orient the negative z-axis along the vector (to-from) under the transformation.")
function. The vectors passed in are *not* normalized meaning that scales
should be preserved in construction of the transform.

Note
Unlike most VEX functions, this function expects rotations
in *degrees*, not radians.

Show/hide arguments

`trs`

One of the transform order constants listed below, which can be imported from `$HFS/houdini/vex/include/math.h`.

| Constant name | Transform Order |
| --- | --- |
| XFORM_SRT | Scale, Rotate, Translate |
| XFORM_STR | Scale, Translate, Rotate |
| XFORM_RST | Rotate, Scale, Translate |
| XFORM_RTS | Rotate, Translate, Scale |
| XFORM_TSR | Translate, Scale, Rotate |
| XFORM_TRS | Translate, Rotate, Scale |

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
