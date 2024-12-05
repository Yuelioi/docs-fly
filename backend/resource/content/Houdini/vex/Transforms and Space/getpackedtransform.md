---
title: getpackedtransform
order: 3
---
| Since | 17.0 |
| --- | --- |

`matrix  getpackedtransform(int input, int primnum)`

Gets the transform of a packed primitive. This is constructed from the `P`
attribute of the primitive’s point and its intrinsic `transform`.

Warning
This function builds a transform from *only* the `P` (position) attribute and the `transform` intrinsic attribute. It ignores various details that the `packedfulltransform` intrinsic attribute includes:

- The packed primitive’s `pivot` intrinsic attribute.
- Instancing attributes such as `orient` (when the `pointinstancetransform` intrinsic attribute is on, as with crowd agents).
- The `packedlocaltransform` intrinsic attribute (Alembic primitives).

So this function will not return the expected transform in several different cases.

The [setpackedtransform](setpackedtransform.html "Sets the transform of a packed primitive.") function has the same issues since it *only* overwrites `P` and `transform`. So, for example, the `packedfulltransform` won’t contain the matrix you would expect if there is a non-zero packed pivot, or in the other cases listed above.

Examples

## examples

```vex
// matrix to transform by
matrix transform = ident();
rotate(transform, radians(45), {0,1,0});
translate(transform, {0,1,0});

matrix tf = getpackedtransform(0, @primnum);
setpackedtransform(0, @primnum, transform * tf);

```
