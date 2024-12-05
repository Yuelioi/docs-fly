---
title: usd_blockprimvar
order: 24
---
| Since | 18.0 |
| --- | --- |

`int  usd_blockprimvar(int stagehandle, string primpath, string name)`

This function blocks the primvar. I.e., removes all time samples and sets the *block* as default value.

Note, if primvar is indexed, you may also want to block the indices with
[usd_blockprimvarindices](usd_blockprimvarindices.html "Blocks the primvar.").

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
// Block the primvar.
usd_blockprimvar(0, "/geo/sphere", "primvar_name");

```
