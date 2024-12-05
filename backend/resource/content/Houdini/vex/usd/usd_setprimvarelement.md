---
title: usd_setprimvarelement
order: 133
---
| Since | 18.0 |
| --- | --- |

`int  usd_setprimvarelement(int stagehandle, string primpath, string name, int index, <type>value)`

This function sets the element value in a array primvar.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`index`

An index of an element in the array primvar.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the value of an element at index 2 in the array primvar.
usd_setprimvarelement(0, "/geo/sphere", "float_arr_primvar", 2, 0.25);

```
