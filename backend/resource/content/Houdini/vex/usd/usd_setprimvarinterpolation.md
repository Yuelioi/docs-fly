---
title: usd_setprimvarinterpolation
order: 136
---
| Since | 18.0 |
| --- | --- |

`int  usd_setprimvarinterpolation(int stagehandle, string primpath, string name, string interpolation)`

This function sets the interpolation style of the given primvar.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`interpolation`

The new interpolation style for the primvar.

The standard interpolation styles are

- “constant” - same value over the entire surface (i.e., detail)
- “uniform” - one value for each uv patch or a face (i.e., primitive)
- “vertex” - values interpolated between each vertex using surface’s basis function (i.e., point)
- “varying” - four values interpolated over uv patch or a face (i.e., vertex)
- “faceVarying” - for polygons and subdivision surfaces, four values are interpolated over each face of the mesh (i.e., vertex)

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the primvar's interpolation style.
usd_setprimvarinterpolation(0, "/geo/mesh", "primvar_name", "faceVarying");

```
