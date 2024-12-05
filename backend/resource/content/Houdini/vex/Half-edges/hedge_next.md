---
title: hedge_next
order: 7
---
`int  hedge_next(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The number of the half-edges that follow (its source is the destination of) `hedge` in the polygon that contains it. Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int nexthedge;

// Get the next half-edge of half-edge number 3.
nexthedge = hedge_next("defgeo.bgeo", 3);

```
