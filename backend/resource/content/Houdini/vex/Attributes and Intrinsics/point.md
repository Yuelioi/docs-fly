---
title: point
order: 34
---
`<type> point(<geometry>geometry, string attribute_name, int pointnumber)`

`<type>[] point(<geometry>geometry, string attribute_name, int pointnumber)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`pointnumber`

The point number to read the attribute on.

Returns

The value of the given attribute on the given point number, or `0` if the attribute or point do not exist.
