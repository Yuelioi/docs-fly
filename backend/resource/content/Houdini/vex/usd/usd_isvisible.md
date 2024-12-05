---
title: usd_isvisible
order: 83
---
| Since | 17.5 |
| --- | --- |

`int  usd_isvisible(<stage>stage, string primpath)`

`int  usd_isvisible(<stage>stage, string primpath, float timecode)`

This function checks whether the given primitive is visible.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

1 if the primitive is visible, and 0 otherwise.

Examples

## examples

```vex
// Check if the sphere primitive is visible.
int is_visible = usd_isvisible(0, "/geometry/sphere");

```
