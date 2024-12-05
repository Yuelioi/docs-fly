---
title: usd_getbbox_size
order: 46
---
| Since | 18.0 |
| --- | --- |

`vector  usd_getbbox_size(<stage>stage, string primpath, string purpose)`

Computes the size of the bounding box for the geometry.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`purpose`

The primitive’s purpose for which to return the bounding box size (e.g., “render”).

Returns

The size of the primitive’s bounding box.

Examples

## examples

```vex
// Get the sphere's bounding box.
vector size = usd_getbbox_size(0, "/src/sphere", "render");

```
