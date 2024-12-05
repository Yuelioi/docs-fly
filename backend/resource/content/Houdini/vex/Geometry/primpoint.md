---
title: primpoint
order: 26
---
`int  primpoint(<geometry>geometry, int primnum, int vertex)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primnum`

The primitive number to get a vertex from.

`vertex`

The vertex number inside of a primitive. 0 is the first vertex.

Returns

The point number the vertex is wired to.
Returns `-1` if failed to find the corresponding point.
