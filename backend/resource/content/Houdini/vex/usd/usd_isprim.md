---
title: usd_isprim
order: 77
---
| Since | 17.5 |
| --- | --- |

`int  usd_isprim(<stage>stage, string primpath)`

This function checks whether the path points to a valid USD primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

Returns

1 if the primitive at the given path is valid, and 0 otherwise.

Examples

## examples

```vex
// Check if the stage coming on the first input has a sphere primitive 
// at scene graph location "/geometry/sphere".
int is_valid_primitive = usd_isprim(0, "/geometry/sphere");

```
