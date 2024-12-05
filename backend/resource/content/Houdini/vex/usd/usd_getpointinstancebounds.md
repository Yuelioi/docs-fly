---
title: usd_getpointinstancebounds
order: 48
---
| Since | 18.0 |
| --- | --- |

`int  usd_getpointinstancebounds(<stage>stage, string primpath, int instance_index, string purpose, vector &min, vector &max)`

`int  usd_getpointinstancebounds(<stage>stage, string primpath, int instance_index, string purpose[], vector &min, vector &max)`

`int  usd_getpointinstancebounds(<stage>stage, string primpath, int instance_index, string purpose, float timecode, vector &min, vector &max)`

`int  usd_getpointinstancebounds(<stage>stage, string primpath, int instance_index, string purpose[], float timecode, vector &min, vector &max)`

This function returns primitive’s axis-aligned bounding box of a particular instance in a point instancer primitive. The point corresponding to the minimum corner of the bounding box will be returned in min, while the maximum will be in max. Always returns 1.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`instance_index`

The index of the instance whose bounding box to return.

`purpose`

The primitive’s purpose for which to return the bounding box (e.g., “default”, “render”).

`timecode`

The USD time code at which to evaluate the attribute. A USD time code roughly corresponds to a frame in Houdini. If not given, the time code corresponding to the current frame is used.

Returns

Always 1.

Examples

## examples

```vex
// Get the second sphere's bounding box.
vector min, max;
usd_getpointinstancebounds(0, "/src/instanced_spheres", 1, "render", min, max);

```
