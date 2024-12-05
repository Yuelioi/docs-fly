---
title: hedge_nextequiv
order: 8
---
`int  hedge_nextequiv(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The next half-edge equivalent to `hedge`, or `hedge` if there are no other half-edges equivalent to it.
Successive calls to `hedge_nextequiv()` cycle through all the equivalent half-edges.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
// Determine the number of half-edges equivalent to half-edge number 3 (including itself)
int num_equiv = 0;
int h = 3;
do
{
h = hedge_nextequiv("defgeo.bgeo", h);
num_equiv++;
} while (h != 3);

```
