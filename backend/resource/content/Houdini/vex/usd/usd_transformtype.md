---
title: usd_transformtype
order: 148
---
| Since | 18.0 |
| --- | --- |

`int  usd_transformtype(string name)`

This function returns the transform operation type implied by the full name.

Show/hide arguments

`name`

The full name of the transform operation, which includes the standard namespace, encodes transformation type, and optionally contains the suffix.

Returns

The the numerical code for the inferred transform operation type implied by the transform operation name. See the VEX “usd.h” header for defines, such as `USD_XFORM_TRANSLATE`, `USD_XFORM_TRANSFORM`, or `USD_XFORM_ROTATE_XYZ`.

Examples

## examples

```vex
// Get the type of the first transform operation on the cube
string cube_xform_ops[] = usd_transformorder(0, "/geo/cube");
int type = usd_transformtype(cube_xform_ops[0]);

```
