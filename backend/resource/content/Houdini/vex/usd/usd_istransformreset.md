---
title: usd_istransformreset
order: 81
---
| Since | 17.5 |
| --- | --- |

`int  usd_istransformreset(<stage>stage, string primpath)`

This function checks if the primitive transform is reset, i.e., whether it uses the world co-ordinate system as the initial space, or whether it inherits the space transformation from the parent (default).

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

`1` if the primitive transform is reset (i.e., the primitive uses the world space), or `0` if it inherits the space from the parent (the default behavior).

Examples

## examples

```vex
// Check if the cube's transform is reset.
int is_xform_reset = usd_istransformreset(1, "/geo/cube");

```
