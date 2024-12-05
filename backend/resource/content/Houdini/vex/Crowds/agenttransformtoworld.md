---
title: agenttransformtoworld
order: 47
---
`int  agenttransformtoworld(<geometry>geometry, int prim, matrix &transforms[])`

Returns `-1` if `len(transforms)` does not match the number of transforms in the agent’s rig, `prim` is out of range, or `prim` is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transforms`

The transforms to convert from local space to world space.
