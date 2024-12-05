---
title: usd_transformorder
order: 146
---
| Since | 17.5 |
| --- | --- |

`string [] usd_transformorder(<stage>stage, string primpath)`

This function returns primitive’s local transform. Transform order is a sequence of transform operations, whose full names are stored in `xformOpOrder` attribute as a string array. Thus, this function returns the value of that attribute.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

The primitive’s transform order.

Examples

## examples

```vex
// Get the cube's transform order.
string cube_xform_ops[] = usd_transformorder(0, "/geo/cube");

```
