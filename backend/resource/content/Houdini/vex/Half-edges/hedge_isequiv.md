---
title: hedge_isequiv
order: 4
---
`int  hedge_isequiv(<geometry>geometry, int hedge1, int hedge2)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge1`

The integer representing the first half-edge.

`hedge2`

The integer representing the second half-edge.

Returns

`1` if `hedge1` and `hedge2` are equivalent, i.e. represent the same
edge in the geometry. This is the case when either source and destination points
of `hedge1` and `hedge2` are respectively the same, or source of each of `hedge1`
and `hedge2` is the same as the destination of the other.

Examples

## examples

```vex
int opposite = 0;

// test if hedges 2 and 3 are oppositely oriented equivalent half-edges
if (hedge_isequiv("defgeo.bgeo", 2, 3))
{
if (hedge_srcpoint("defgeo.bgeo", 2) == hedge_dstpoint("defgeo.bgeo", 3))
opposite = 1;
}

```
