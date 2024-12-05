---
title: hedge_srcvertex
order: 17
---
`int  hedge_srcvertex(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The vertex number of the source vertex of the `hedge`.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int srcvtx;

// Get the source vertex of half-edge number 3.
srcvtx = hedge_srcvertex("defgeo.bgeo", 3);

```
