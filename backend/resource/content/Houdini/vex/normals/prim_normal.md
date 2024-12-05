---
title: prim_normal
order: 2
---
`vector  prim_normal(<geometry>geometry, int prim_number, vector uvw)`

`vector  prim_normal(<geometry>geometry, int prim_number, float u, float v)`

`vector  prim_normal(<geometry>geometry, int prim_number, float u, float v, float w)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`uvw`, `u`, `v`, `w`

When w is not given, it is treated as zero.

Returns

The normal of the primitive (prim_number) at parametric location u, v, w.
