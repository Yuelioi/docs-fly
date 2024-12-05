---
title: usd_setprimvarindices
order: 135
---
| Since | 18.0 |
| --- | --- |

`int  usd_setprimvarindices(int stagehandle, string primpath, string name, int indices[])`

This function sets the indices for a given primvar, thus making it an indexed primvar if it was not already.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Primvar name (without namespace).

`indices`

The index array to set.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set the primvar's value and indices.
float values[]  = array(0, 100, 200, 300, 400, 500);
int   indices[] = array(5,5,4,4,3,3,2,2,1,1,0,0);
usd_setprimvar(0, "/geo/mesh", "primvar_name", values); 
usd_setprimvarindices(0, "/geo/mesh", "primvar_name", indices);

```
