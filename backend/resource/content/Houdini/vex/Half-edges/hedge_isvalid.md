---
title: hedge_isvalid
order: 6
---
`int  hedge_isvalid(<geometry>geometry, int hedge)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`hedge`

The integer representing a half-edge.

Returns

`1` if `hedge` represents a valid half-edge in the referenced geometry, or `0` otherwise.

Examples

## examples

```vex
int srcpt;

// find the source point of a half-edge number 3 if it is valid
if (hedge_isvalid("defgeo.bgeo", 3))
srcpt = hedge_srcpoint("defgeo.bgeo", 3);

```
