---
title: setpackedtransform
order: 24
---
| Since | 17.0 |
| --- | --- |

`void  setpackedtransform(int input, int primnum, matrix transform)`

Sets the transform of a packed primitive. This modifies the `P` attribute of
the primitive’s point as well as its intrinsic `transform`.

Warning
This function replaces *only* the `P` (position) attribute and the `transform` intrinsic attribute. It ignores various details that the `packedfulltransform` intrinsic attribute includes:

- The packed primitive’s `pivot` intrinsic attribute.
- Instancing attributes such as `orient` (when the `pointinstancetransform` intrinsic attribute is on, as with crowd agents).
- The `packedlocaltransform` intrinsic attribute (Alembic primitives).

So this function will not apply the expected transform in several different cases.

The [getpackedtransform](getpackedtransform.html "Gets the transform of a packed primitive.") function has the same issues since it returns a transform *only* based on `P` and `transform`.

Examples

## examples

```vex
// matrix to transform by
matrix tf = ident();
rotate(tf, radians(45), {0,1,0});
translate(tf, {0,1,0});

matrix transform = getpackedtransform(0, @primnum);
setpackedtransform(0, @primnum, transform * tf);

```
