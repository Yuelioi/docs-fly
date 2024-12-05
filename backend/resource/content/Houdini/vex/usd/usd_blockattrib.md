---
title: usd_blockattrib
order: 23
---
| Since | 18.0 |
| --- | --- |

`int  usd_blockattrib(int stagehandle, string primpath, string name)`

This function blocks the attribute. I.e., removes all time samples and sets the *block* as default value.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Attribute name.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Block the attribute.
usd_blockattrib(0, "/geo/sphere", "attribute_name");

```
