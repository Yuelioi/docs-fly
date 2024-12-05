---
title: hedge_srcpoint
order: 16
---
`int  hedge_srcpoint(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The point number of the source point of the `hedge`.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int srcpt;

// Get the source point of half-edge number 3.
srcpt = hedge_srcpoint("defgeo.bgeo", 3);

```
