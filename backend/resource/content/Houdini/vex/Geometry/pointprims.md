---
title: pointprims
order: 21
---
`int [] pointprims(<geometry>geometry, int ptnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`ptnum`

The point number to get a primitive from.

Returns

An array of primitive numbers. These will be in ascending order
and not contain duplicates.

If no primitives own the given point the array will be empty.
