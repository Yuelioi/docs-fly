---
title: usd_setkind
order: 129
---
| Since | 17.5 |
| --- | --- |

`int  usd_setkind(int stagehandle, string primpath, string kind)`

This function sets the primitive’s kind.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`kind`

The kind to set the primitive to.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the sphere primitive to be an assembly.
usd_setkind(0, "/geo/sphere", "assembly");

```
