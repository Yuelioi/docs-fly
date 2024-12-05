---
title: hex_adjacent
order: 1
---
`int  hex_adjacent(<geometry>geometry, int primindex, int faceno)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primindex`

The primitive number.

`faceno`

The face on the hexahedron. Ranges from `0` to `5`.

Returns

The primitive number of the hexahedron attached to the given face.
Returns `-1` f the primitive is not a hex or doesn’t have an adjacent hexahedron.

Use [hex_faceindex](hex_faceindex.html "Returns vertex indices of each face of a hexahedron.") to get the vertex indices of each face of a hexahedron.
