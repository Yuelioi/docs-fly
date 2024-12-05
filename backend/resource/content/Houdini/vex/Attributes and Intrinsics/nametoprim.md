---
title: nametoprim
order: 32
---
`int  nametoprim(<geometry>geometry, string name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The number of the point with the given value in the `name` attribute. Returns `-1` if no point has the given ID, or if the geometry has no `name` attribute.

To look up a point by its `id` attribute value, use [idtopoint](idtopoint.html "Finds a point by its id attribute."). To look up a point by an arbitrary string or int attribute value, use [findattribval](findattribval.html "Finds a primitive/point/vertex that has a certain attribute value.").
