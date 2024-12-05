---
title: usd_addprim
order: 6
---
| Since | 18.0 |
| --- | --- |

`int  usd_addprim(int stagehandle, string primpath, string typename)`

This function creates a new primitive of a given type and at a given path, if it one does not exist yet.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`typename`

The name or an alias of the type.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Adds a sphere primitive.
usd_addprim(0, "/geo/sphere", "Sphere");

```
