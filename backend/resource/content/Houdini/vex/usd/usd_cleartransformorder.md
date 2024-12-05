---
title: usd_cleartransformorder
order: 30
---
| Since | 17.5 |
| --- | --- |

`int  usd_cleartransformorder(int stagehandle, string primpath)`

This function clears the primitive’s transform order. Transform order is a sequence of transform operations, whose full names are stored in `xformOpOrder` attribute as a string array. Thus, this function clears that attribute.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
usd_cleartransformorder(0, "/geo/cone");

```
