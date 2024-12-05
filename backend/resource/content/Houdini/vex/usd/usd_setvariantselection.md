---
title: usd_setvariantselection
order: 141
---
| Since | 17.5 |
| --- | --- |

`int  usd_setvariantselection(<stage>stage, string primpath, string variantset, string variant)`

This function sets the selected variant in a given variant set.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`variantset`

The variant set name.

`variant`

The name of a variant to select.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Set the variant "cone" in a variant set "shapes" on the "shape_shifter" primitive.
usd_setvariantselection(0, "/geo/shape_shifter", "shapes", "cone");

```
