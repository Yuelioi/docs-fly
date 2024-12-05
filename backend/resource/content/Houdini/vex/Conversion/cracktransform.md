---
title: cracktransform
order: 3
---
`vector  cracktransform(int trs, int xyz, int c, vector pivot, vector pivot_rotate, matrix xform)`

`vector  cracktransform(int trs, int xyz, int c, vector pivot, matrix xform)`

Depending on the value of c, returns the translate (`c=0`), rotate
(`c=1` or `c=4`), scale (`c=2`) or shears (`c=3`) component of the transform (xform). The
function uses the given transform and rotation orders (trs and
xyz) , the given pivot point (pivot) and optional pivot rotation (pr) to calculate the return
value.

Note
Rotation angles (when `c=1`) are returned in degrees, whereas many other VEX functions use radians.
You can use the [radians](radians.html "Converts the argument from degrees into radians.") VEX function to convert the vector of angles in degrees to a vector of angles in radians.
For example: `vector angles = radians(cracktransform(XFORM_TRS, XFORM_XYZ, 1, {0,0,0}, M));`

Note
Rotation angles (when `c=4`) are returned in radians.

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

`void  cracktransform(int trs, int xyz, vector pivot, vector pivot_rotate, matrix xform, vector &t, vector &r, vector &s, vector &shears)`

Returns the translate, rotate, scale, and shear components of xform in t, r, s, and shears, respectively.
If more than one component is needed, using this overload is more efficient than making multiple calls to the other function signature.

`void  cracktransform(int trs, int xyz, vector pivot, matrix xform, vector &t, vector &r, vector &s)`

Returns the translate, rotate, and scale of xform in t, r, s respectively.
This overload doesn’t support pivot_rotate or shears.
If more than one component is needed, using this overload is more efficient than making multiple calls to the other function signature.
