---
title: pointedge
order: 18
---
`int  pointedge(<geometry>geometry, int point1, int point2)`

Returns `-1` if no such half-edge exists. Otherwise returns the number of a half-edge that either has `point1` as source or has `point2` as desination, or the other way around.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`point1`, `point2`

The point numbers in the geometry for the two endpoints of the returned half-edge. `0` is the first point.

Examples

## examples

```vex
int edge_count = 0;

// Determine if there is an edge between points 23 and 25:
int h0 = pointedge("defgeo.bgeo", 23, 25);
if (h0 != -1)
{
// Edge exists!
}

```
