---
title: prim
order: 44
---
`<type> prim(<geometry>geometry, string attribute_name, int primnumber)`

`<type>[] prim(<geometry>geometry, string attribute_name, int primnumber)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`primnumber`

The primitive number to read the attribute on.

Returns

The value of the given attribute on the given point number, or `0` if the attribute or point do not exist.
