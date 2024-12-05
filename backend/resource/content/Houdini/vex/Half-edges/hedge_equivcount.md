---
title: hedge_equivcount
order: 3
---
`int  hedge_equivcount(<geometry>geometry, int hedge)`

Note
Equivalent half-edges may be oppositely oriented, i.e. the source of one can be the destination of the other and vice versa.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The number of half-edges that have the same endpoint as `hedge` (including `hedge`), or `-1` if the half-edge is not valid.

Examples

## examples

```vex
int is_boundary = 0;
int is_interior = 0;
int is_nonmanifold = 0;

// Determine the type of edge represented by half-edge number 3:
int numeq;
numeq = hedge_equivcount("defgeo.bgeo", 3);
if (numeq == 1)
is_boundary = 1;
else if (numeq >= 3)
is_nonmanifold = 1;
else
is_interior = 1;

```
