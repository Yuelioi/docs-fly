---
title: usd_addrelationshiptarget
order: 8
---
| Since | 18.0 |
| --- | --- |

`int  usd_addrelationshiptarget(int stagehandle, string primpath, string name, string target)`

This function adds a target to the primitive’s relationship.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

The relationship name.

`target`

The the target path to add.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Add the sphere to cube's relationship.
usd_addrelationshiptarget(0, "/geo/cube", "relationship_name", "/geo/sphere");

```
