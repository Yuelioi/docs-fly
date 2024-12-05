---
title: primhedge
order: 21
---
`int  primhedge(<geometry>geometry, int prim)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number in the geometry. `0` is the first primitive.

Returns

The number of an arbitrary half-edge contained in `prim`.
Returns `-1` if the primitive number is not valid.
