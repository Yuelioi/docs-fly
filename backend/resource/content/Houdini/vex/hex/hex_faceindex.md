---
title: hex_faceindex
order: 2
---
`int  hex_faceindex(int faceno, int vtxno)`

Returns `-1` if an invalid number is specified.

Returns `0` to `7` to refer to the eight vertices of a generic hexahedron.

Show/hide arguments

`faceno`

The face on the hexahedron. Ranges from `0` to `5`.

`vtxno`

Which vertex on the quadrilateral to return, `0` to `3`. Starts with
the lowest number and follows Houdini’s winding convention.
