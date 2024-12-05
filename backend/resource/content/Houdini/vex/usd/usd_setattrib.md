---
title: usd_setattrib
order: 123
---
| Since | 17.5 |
| --- | --- |

`int  usd_setattrib(int stagehandle, string primpath, string name, <type>value)`

`int  usd_setattrib(int stagehandle, string primpath, string name, <type>value[])`

This function sets the attribute value.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Attribute name.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the value of some attributes.
usd_setattrib(0, "/geo/sphere", "float_attrib", 0.25);
usd_setattrib(0, "/geo/sphere", "string_attrib", "foo bar baz");
usd_setattrib(0, "/geo/sphere", "vector_attrib", {1.25, 1.50, 1.75});

float  f_arr[] = {0, 0.25, 0.5, 0.75, 1};
usd_setattrib(0, "/geo/sphere", "float_array_attrib", f_arr);

```
