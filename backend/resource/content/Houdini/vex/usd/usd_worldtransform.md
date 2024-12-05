---
title: usd_worldtransform
order: 154
---
| Since | 17.5 |
| --- | --- |

`matrix  usd_worldtransform(<stage>stage, string primpath)`

`matrix  usd_worldtransform(<stage>stage, string primpath, float timecode)`

This function returns primitive’s world transform.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The primitive’s world transform.

Examples

## examples

```vex
// Get the cube's world transform.
matrix cube_world_xform = usd_worldtransform(0, "/src/cube");

```
