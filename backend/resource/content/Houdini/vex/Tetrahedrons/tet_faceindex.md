---
title: tet_faceindex
order: 2
---
`int  tet_faceindex(int faceno, int vtxno)`

Returns `-1` if an invalid number is specified.

Returns `0` to `3` to refer to the four vertices of a generic tetrahedron.

Show/hide arguments

`faceno`

The face on the tetrahedron. Face 0 is the triangle that doesn’t
have vertex 0.

`vtxno`

Which vertex on the triangle to return, `0` to `2`. Starts with
the lowest number and follows Houdini’s winding convention, ie,
face 0 is vertices 1, 2, and 3.
