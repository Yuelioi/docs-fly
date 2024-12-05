---
title: usd_addprimvar
order: 7
---
| Since | 18.0 |
| --- | --- |

`int  usd_addprimvar(int stagehandle, string primpath, string name, string typename)`

`int  usd_addprimvar(int stagehandle, string primpath, string name, string typename, string interpolation)`

This function adds a primvar of a given type to the primitive, if such primvar is not part of a schema. It is useful for controlling the exact type of a custom primvar. For primvars defined by primitive’s schema, this call has no effect, because the schema already determines their type.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`typename`

The name or an alias of the type.

`interpolation`

The name of the interpolation to be used for this primvar (e.g., “constant”, “vertex”, “faceVarying”, etc).

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Adds a half-precision float primvar and sets its falue.
usd_addprimvar(0, "/geo/sphere", "half_primvar", "half3");
usd_setprimvar(0, "/geo/sphere", "half_primvar", {1.25, 1.50, 1.75});

// Adds a color primitive with 'vertex' interpolation.
usd_addprimvar(0, pp, "color_primvar", "color3d[]", "vertex");
usd_setprimvar(0, pp, "color_primvar",  vector[](array({1,0,0}, {0,1,0}, {0,0,1})));

```
