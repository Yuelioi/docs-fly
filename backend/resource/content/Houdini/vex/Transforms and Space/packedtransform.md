---
title: packedtransform
order: 14
---
| Since | 17.0 |
| --- | --- |

`void  packedtransform(int input, int primnum, matrix transform)`

Transforms a packed primitive by the specified transform. This modifies the `P`
attribute of the primitive’s point as well as its intrinsic `transform`.

This is equivalent to the following code.

```vex
// matrix to transform by
matrix transform = ident();
rotate(transform, radians(45), {0,1,0});
translate(transform, {0,1,0});

// get current packed transform
matrix3 primtf = primintrinsic(0, "transform", primnum);
setprimintrinsic(0, "transform", primnum, primtf * (matrix3)transform);
int primpoint = primpoint(0, primnum, 0);
vector pos = point(0, "P", primpoint);
setpointattrib(0, "P", primpoint, pos * transform);

```
