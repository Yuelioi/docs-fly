---
title: priminteriorweights
order: 52
---
`int  priminteriorweights(<geometry>geometry, int prim_num, vector uvw, int &verts[], float &weights[])`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim_num`

The primitive number to read the attribute from.

`uvw`

The primitive UVW coordinates at which to read the attribute.

Finds the indices and weightings of the vertices that will compute an interior
point given the UVW coordinates. The indices are linear vertex indices.
