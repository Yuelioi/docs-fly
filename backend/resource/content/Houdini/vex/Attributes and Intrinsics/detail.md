---
title: detail
order: 14
---
`<type> detail(<geometry>geometry, string attribute_name, int ignored=0)`

`<type>[] detail(<geometry>geometry, string attribute_name, int ignored=0)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`ignored`

The last argument is always ignored.
It is just there so you can change a prim/point/vertex call (which each have an element number argument) to a detail call by changing the name without having to change the arguments as well.

Returns

`0` if importing the attribute failed, or the value of the attribute on success.
