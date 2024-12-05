---
title: usd_pointinstance_getbbox_max
order: 100
---
| Since | 18.0 |
| --- | --- |

`vector  usd_pointinstance_getbbox_max(<stage>stage, string primpath, int instance_index, string purpose)`

Computes the largest position of the bounding box for the instance geometry.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`instance_index`

The index of the instance whose bounding box to return.

`purpose`

The primitive’s purpose for which to return the bounding box maximum (e.g., “render”).

Returns

The maximum position of the instance’s bounding box.

Examples

## examples

```vex
// Get the max of the first instance's boundsng box.
vector max = usd_pointinstance_getbbox_max(0, "/src/instanced_spheres", 0, "render");

```
