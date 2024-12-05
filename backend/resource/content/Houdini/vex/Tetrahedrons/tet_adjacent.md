---
title: tet_adjacent
order: 1
---
`int  tet_adjacent(<geometry>geometry, int primindex, int faceno)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primindex`

The primitive number.

`faceno`

The face on the tetrahedron. Face 0 is the triangle that doesn’t
have vertex 0.

Returns

The primitive number of the tetrahedron opposite the given vertex.
Returns `-1` f the primitive is not a tet or doesn’t have an adjacent tetrahedron.

Use [tet_faceindex](tet_faceindex.html "Returns vertex indices of each face of a tetrahedron.") to get the vertex indices of each face of a tetrahedron.
