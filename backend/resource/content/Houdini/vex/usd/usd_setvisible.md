---
title: usd_setvisible
order: 143
---
| Since | 17.5 |
| --- | --- |

`int  usd_setvisible(int stagehandle, string primpath, int flag)`

This function makes the primitive visible or invisible, depending on the given flag parameter.

NOTE: This function is similar to `usd_setvisibility()` which, in addition to making a primitive visible or invisible, can also configure it to inherit visibility from the parent.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`flag`

Non-zero to make the primitive visible or 0 to make it invisible.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the sphere primitive as visible.
usd_setvisible(0, "/geo/sphere", true);

```
