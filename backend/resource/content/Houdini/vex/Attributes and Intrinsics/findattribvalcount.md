---
title: findattribvalcount
order: 21
---
`int  findattribvalcount(<geometry>geometry, string attribclass, string attribute_name, int|stringvalue)`

Returns the number of elements that has that integer or string value set on the given attribute name.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

`attribute_name`

The name of the attribute to read.

`value`

The value to match. Must be of the same type as the attribute.
