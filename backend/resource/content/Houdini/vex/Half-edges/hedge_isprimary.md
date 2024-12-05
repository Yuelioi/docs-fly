---
title: hedge_isprimary
order: 5
---
`int  hedge_isprimary(string geometry, int hedge)`

`int  hedge_isprimary(int opinput, int hendge)`

Show/hide arguments

`geometry`

The name of the geometry file to reference. Inside Houdini, this may be `op:full_path_to_sop` to reference a SOP.

`hedge`

The integer representing a half-edge.

Returns

`1` if `hedge` represents a primary half-edge in the referenced geometry, or `0` otherwise.

Examples

## examples

```vex
int numedges;

// Count the number of edges

if (hedge_isprimary("defgeo.bgeo", 3))
numedges++;

```
