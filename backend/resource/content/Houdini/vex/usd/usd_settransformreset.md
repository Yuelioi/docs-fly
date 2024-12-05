---
title: usd_settransformreset
order: 140
---
| Since | 18.0 |
| --- | --- |

`int  usd_settransformreset(int stagehandle, string primpath, int flag)`

This function sets the primitive’s transform reset flag, i.e., whether the primitive uses the world co-ordinate system as the initial space, or whether it inherits the space transformation from the parent (default).

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`flag`

If `1`, the primitive will have its transform reset, i.e. will use the world co-ordinate system as the initial space. If `0` it will inherit the space transform from the parent.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Ignore parent's transform.
usd_settransformreset(0, "/geo/cone", 1);

```
