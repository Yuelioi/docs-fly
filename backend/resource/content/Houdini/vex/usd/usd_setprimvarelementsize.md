---
title: usd_setprimvarelementsize
order: 134
---
| Since | 18.0 |
| --- | --- |

`int  usd_setprimvarelementsize(int stagehandle, string primpath, string name, int size)`

This function sets the element size of a given primvar.

The primvar element size applies to array primvars, but it does not encode the length of the array. It specifies how many consecutive array elements should be taken as an atomic element to be interpolated over a gprim. So, on a mesh, array length relates to element size like this `array_length = element_size * face_count`.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`size`

The new element size for the primvar.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the primvar's element size to 2.
usd_setprimvarelementsize(0, "/geo/mesh", "primvar_name", 2);

```
