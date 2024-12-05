---
title: hedge_dstpoint
order: 1
---
`int  hedge_dstpoint(<geometry>geometry, int hedge)`

Returns `-1` if the half-edge hedge is invalid. Otherwise, returns the point number of the destination point of the half-edge `hedge`.

Show/hide arguments

`geometry`

The name of the geometry file to reference. Inside Houdini, this may be `op:full_path_to_sop` to reference a SOP.

`hedge`

Input half-edge.

Examples

## examples

```vex
int dstpt;

// Get vertex number of half-edge number 3.
dstpt = hedge_dstpoint("defgeo.bgeo", 3);

```
