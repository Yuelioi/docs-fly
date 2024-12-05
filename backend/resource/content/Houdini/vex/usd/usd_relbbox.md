---
title: usd_relbbox
order: 120
---
| Since | 18.0 |
| --- | --- |

`vector  usd_relbbox(<stage>stage, string primpath, string purpose, vector position)`

Returns the relative position of the given point with respect to the bounding box of the primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`purpose`

The primitive’s purpose for which to use the bounding box (e.g., “render”).

Returns

The relative position of the given point with respect to the bounding box of the primitive.

Examples

## examples

```vex
// Get the points relative position.
vector pt = {1, 0, 0};
vector rel_pt = usd_relbbox(0, "/src/sphere", "render", pt);

```
