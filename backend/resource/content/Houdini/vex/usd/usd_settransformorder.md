---
title: usd_settransformorder
order: 139
---
| Since | 18.0 |
| --- | --- |

`int  usd_settransformorder(int stagehandle, string primpath, string transformorder[])`

This function sets the primitive’s transform order. Transform order is a sequence of transform operations, whose full names are stored in `xformOpOrder` attribute as a string array. Thus, this function sets that attribute.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`transformorder`

The transform order to set on the primitive. It is a list of transform operation full names.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
string  ops[] = {"xformOp:translate:xform_cube_t", "xformOp:rotateZ:xform_cube_r", "xformOp:rotateXYZ:xform_cube_r", "xformOp:scale:xform_cube_s"};
usd_settransformorder(0, "/geo/cube", ops);

```
