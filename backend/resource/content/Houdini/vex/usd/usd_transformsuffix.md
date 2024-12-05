---
title: usd_transformsuffix
order: 147
---
| Since | 18.0 |
| --- | --- |

`string  usd_transformsuffix(string name)`

This function returns the suffix contained in the full name of a transform operation.

Show/hide arguments

`name`

The full name of the transform operation, which includes the standard namespace, encodes transformation type, and optionally contains the suffix.

Returns

The suffix contained in the transform operation name.

Examples

## examples

```vex
// Get the suffix of the first transform operation on the cube
string cube_xform_ops[] = usd_transformorder(0, "/geo/cube");
string suffix = usd_transformsuffix(cube_xform_ops[0]);

```
