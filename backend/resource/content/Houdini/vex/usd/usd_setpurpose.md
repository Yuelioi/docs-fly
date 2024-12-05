---
title: usd_setpurpose
order: 137
---
| Since | 17.5 |
| --- | --- |

`int  usd_setpurpose(int stagehandle, string primpath, string purpose)`

This function sets the primitive’s purpose, e.g., “default”, “render”, “proxy”, “guide”, etc.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`purpose`

The primitive’s purpose to set.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the sphere primitive to be traversable only for rendering.
usd_setpurpose(0, "/geo/sphere", "render");

```
