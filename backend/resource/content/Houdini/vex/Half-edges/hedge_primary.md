---
title: hedge_primary
order: 15
---
`int  hedge_primary(<geometry>geometry, int hedge)`

Each class of equivalent half-edges has precisely one primary half-edge. In particular, a half-edge which is equivalent to no other half-edges is always primary. Primary half-edges are useful for accounting for each edge exactly once as each edge may be realized by any number of equivalent half-edges.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

Input half-edge.

Returns

The primary half-edge `hedge` that shares the source and destination of `hedge` (possibly in reverse order).
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int primhedge;

// Get the primary half-edge equivalent  to half-edge number 3.
primhedge = hedge_primary("defgeo.bgeo", 3);

```
