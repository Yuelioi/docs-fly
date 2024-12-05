---
title: usd_setattribelement
order: 124
---
| Since | 18.0 |
| --- | --- |

`int  usd_setattribelement(int stagehandle, string primpath, string name, int index, <type>value)`

This function sets the element value in an array attribute.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Attribute name.

`index`

An index of an element in the array attribute.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the value of element at index 2 in the array attribute.
usd_setattribelement(0, "/geo/sphere", "float_array_attrib", 2, 0.25);

```
