---
title: usd_blockrelationship
order: 26
---
| Since | 18.0 |
| --- | --- |

`int  usd_blockrelationship(int stagehandle, string primpath, string name)`

This function blocks the primitive’s relationship, i.e., clears the targets from the relationship.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

The relationship name.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Clear the the cube's relationship.
usd_blockrelationship(0, "/geo/cube", "relationship_name");

```
