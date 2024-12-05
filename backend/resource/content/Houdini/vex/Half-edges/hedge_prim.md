---
title: hedge_prim
order: 14
---
`int  hedge_prim(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The primitive number of primitive that contains (the source and destination vertices of) `hedge`.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int prim;

// Get the primitive number of half-edge number 3.
prim = hedge_prim("defgeo.bgeo", 3);

```
