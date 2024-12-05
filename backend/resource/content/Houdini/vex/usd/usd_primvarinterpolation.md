---
title: usd_primvarinterpolation
order: 110
---
| Since | 18.0 |
| --- | --- |

`string  usd_primvarinterpolation(<stage>stage, string primpath, string name)`

This function returns the interpolation style of the primvar found directly on the given primitive. E.g. “constant”, “varying”, etc.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The primvar’s interpolation. The standard interpolation styles are:

- “constant” - same value over the entire surface (i.e., detail)
- “uniform” - one value for each uv patch or a face (i.e., primitive)
- “vertex” - values interpolated between each vertex using surface’s basis function (i.e., point)
- “varying” - four values interpolated over uv patch or a face (i.e., vertex)
- “faceVarying” - for polygons and subdivision surfaces, four values are interpolated over each face of the mesh (i.e., vertext)

Examples

## examples

```vex
// Get the interpolation style of the primvar on the cube.
string interpolation = usd_primvarinterpolation(0, "/geo/cube", "primvar_name");

```
