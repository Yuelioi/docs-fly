---
title: usd_setrelationshiptargets
order: 138
---
| Since | 18.0 |
| --- | --- |

`int  usd_setrelationshiptargets(int stagehandle, string primpath, string name, string targets[])`

This function sets the targets in the primitive’s relationship.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

The relationship name.

`targets`

The the target paths to set the relationship to.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the cube's relationship.
usd_setrelationshiptargets(0, "/geo/cube", "new_relation", array("/geo/sphere6", "/geo/sphere7"));

```
