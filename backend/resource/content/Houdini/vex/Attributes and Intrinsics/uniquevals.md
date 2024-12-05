---
title: uniquevals
order: 75
---
| Since | 17.0 |
| --- | --- |

`int [] uniquevals(<geometry>geometry, string attribclass, string attribute_name)`

`string [] uniquevals(<geometry>geometry, string attribclass, string attribute_name)`

If any points/primitives/vertices in the geometry have the same value for the given attribute, the set of *unique* values will be smaller than the total number of points/primitives/vertices. This function lets you acquire the set of unique values.

This function only works with string and integer attributes.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

`attribute_name`

The name of the attribute (or intrinsic) to read.
