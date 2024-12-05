---
title: usd_setactive
order: 122
---
| Since | 17.5 |
| --- | --- |

`int  usd_setactive(int stagehandle, string primpath, int flag)`

This function sets the primitive’s active state.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`flag`

Non-zero to make the primitive active or 0 to make it inactive.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the sphere primitive as active.
usd_setactive(0, "/geo/sphere", true);

```
