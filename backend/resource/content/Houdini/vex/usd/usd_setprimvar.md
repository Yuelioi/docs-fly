---
title: usd_setprimvar
order: 132
---
| Since | 18.0 |
| --- | --- |

`int  usd_setprimvar(int stagehandle, string primpath, string name, <type>value)`

`int  usd_setprimvar(int stagehandle, string primpath, string name, <type>value[])`

This function sets the primvar’s value.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the value of some primvars.
usd_setprimvar(0, "/geo/sphere", "float_primvar", 0.25);
usd_setprimvar(0, "/geo/sphere", "string_primvar", "foo bar baz");
usd_setprimvar(0, "/geo/sphere", "vector_primvar", {1.25, 1.50, 1.75});

float  f_arr[] = {0, 0.25, 0.5, 0.75, 1};
usd_setprimvar(0, "/geo/sphere", "float_array_primvar", f_arr);

```
