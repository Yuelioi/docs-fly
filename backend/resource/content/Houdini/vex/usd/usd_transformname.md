---
title: usd_transformname
order: 145
---
| Since | 18.0 |
| --- | --- |

`string  usd_transformname(int transformtype, string suffix)`

This function returns the full name of a transform operation for the given type and suffix.

Show/hide arguments

`transformtype`

The numerical code for the transformation type. See the VEX “usd.h” header for defines, such as `USD_XFORM_TRANSLATE`, `USD_XFORM_TRANSFORM`, or `USD_XFORM_ROTATE_XYZ`.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

Returns

The full name of transform operation.

Examples

## examples

```vex
// Construct a full name for a translation operation with suffix "cone_pivot"
string pivot_xform_name = usd_transformname(USD_XFORM_TRANSLATE, "cone_pivot");

```
