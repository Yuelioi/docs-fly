---
title: hedge_postdstpoint
order: 9
---
`int  hedge_postdstpoint(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The point number of the point into which the vertex that comes after the destination vertex of the
half-edge `hedge` in the primitive that contains `hedge` is wired.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int postdstpt;

// Get the destination vertex half-edge number 3.
postdstpt = hedge_postdstpoint("defgeo.bgeo", 3);

```
