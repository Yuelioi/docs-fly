---
title: hedge_prev
order: 13
---
`int  hedge_prev(<geometry>geometry, int hedge)`

Returns `-1` if `hedge` is invalid. Otherwise, returns the number of the half-edge
that precedes (its destination is the source of) `hedge` in the polygon that contains it.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The number of the half-edge that precedes (its destination is the source of) `hedge` in the polygon that contains it.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int prev;

// Get the previous half-edge of half-edge number 3.
prevhedge = hedge_prev("defgeo.bgeo", 3);

```
