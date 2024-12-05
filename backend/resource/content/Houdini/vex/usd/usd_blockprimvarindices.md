---
title: usd_blockprimvarindices
order: 25
---
| Since | 18.0 |
| --- | --- |

`int  usd_blockprimvarindices(int stagehandle, string primpath, string name)`

This function blocks the primvar indices. I.e., removes all time samples and sets the *block* as default value. It changes the indexed primvar into a non-indexed primvar.

Note, you may also want to block the primvar itself with [usd_blockprimvar](usd_blockprimvar.html "Blocks the primvar.").

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Block the primvar indices.
usd_blockprimvarindices(0, "/geo/sphere", "primvar_name");

```
