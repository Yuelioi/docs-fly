---
title: primvertices
order: 30
---
`int [] primvertices(<geometry>geometry, int primnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primnum`

The primitive number to get a vertex from.

Returns

An array of linear vertex indices, in the same order as stored on the primitive itself.
If the primitive number is not valid, the array will be empty.
