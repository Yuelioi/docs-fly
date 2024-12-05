---
title: usd_setdrawmode
order: 128
---
| Since | 17.5 |
| --- | --- |

`int  usd_setdrawmode(int stagehandle, string primpath, string mode)`

This function sets the primitive’s draw mode.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`mode`

The draw mode to set the primitive to.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the sphere to draw as bounds and the cube to draw as default.
usd_setdrawmode(0, "/geo/sphere", "bounds");
usd_setdrawmode(0, "/geo/cube", "default");

```
