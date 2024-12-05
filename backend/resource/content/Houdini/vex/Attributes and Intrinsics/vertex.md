---
title: vertex
order: 77
---
`<type> vertex(<geometry>geometry, string attribute_name, int linear_vertex_index)`

`<type>[] vertex(<geometry>geometry, string attribute_name, int linear_vertex_index)`

Specifies the vertex using the linear index into the list of all vertices.

`<type> vertex(<geometry>geometry, string attribute_name, int prim_num, int vertex_num)`

`<type>[] vertex(<geometry>geometry, string attribute_name, int prim_num, int vertex_num)`

Specifies the vertex as a primitive number and then an offset into the list of vertices on that primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

Returns

The value of the given attribute on the given vertex, or `0` if the attribute or vertex do not exist.
