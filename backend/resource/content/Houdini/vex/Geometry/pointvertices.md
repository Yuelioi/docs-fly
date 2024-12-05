---
title: pointvertices
order: 23
---
`int [] pointvertices(<geometry>geometry, int ptnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`ptnum`

The point number to get a vertex from.

Returns

An array of vertices that are wired to the given point. You should not rely on the numbers being in a particular order.

If the given point contains no vertices, the array will be empty.
