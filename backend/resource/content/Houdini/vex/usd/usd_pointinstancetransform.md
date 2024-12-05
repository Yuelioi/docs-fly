---
title: usd_pointinstancetransform
order: 104
---
| Since | 18.0 |
| --- | --- |

`matrix  usd_pointinstancetransform(<stage>stage, string primpath, int index)`

`matrix  usd_pointinstancetransform(<stage>stage, string primpath, int index, float timecode)`

This function returns point instance transform.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`index`

The index of the instance within the point instancer.

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

The transform of the point instance.

Examples

## examples

```vex
// Get the transform of the third instance.
matrix xform = usd_pointinstancetransform(0, "/src/instanced_cubes", 2);

```
