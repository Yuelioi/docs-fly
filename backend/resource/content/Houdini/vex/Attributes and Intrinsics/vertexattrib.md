---
title: vertexattrib
order: 78
---
`<type> vertexattrib(<geometry>geometry, string attribute_name, int linear_vertex_index, int &success)`

`<type>[] vertexattrib(<geometry>geometry, string attribute_name, int linear_vertex_index, int &success)`

Unlike [vertex](vertex.html "Reads a vertex attribute value from a geometry."), this function does not have a version that takes a primitive number and primitive vertex number. If you have a primitive number and primitive vertex number, you can convert them into a linear index using [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`linear_vertex_index`

A linear index into the list of all vertices. If you have a primitive number and primitive vertex number, you can convert them into a linear index using [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.").

`success`

The function overwrites this variable with `1` if the attribute exists and was read successfully, or `0` otherwise.

Returns

The value of the given attribute on the given point number.
