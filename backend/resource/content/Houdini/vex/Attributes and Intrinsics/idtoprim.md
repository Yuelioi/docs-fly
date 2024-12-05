---
title: idtoprim
order: 30
---
`int  idtoprim(<geometry>geometry, int id)`

Returns the number of the primitive with the given value in the `id` attribute. Returns `-1` if no primitive has the given ID.

If the geometry doesn’t have an `id` attribute, primitive numbers are used as ids. In this case, the function will return the given `id` value, unless it is greater than the number of points in the source geometry, in which case the function will return `-1`.

To look up a primitive by its `name` attribute value, use [nametoprim](nametoprim.html "Finds a primitive by its name attribute."). To look up a primitive by an arbitrary string or int attribute value, use [findattribval](findattribval.html "Finds a primitive/point/vertex that has a certain attribute value.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
