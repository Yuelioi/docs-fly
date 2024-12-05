---
title: agentfindtransformgroup
order: 26
---
`int  agentfindtransformgroup(<geometry>geometry, int prim, string transformgroup)`

Returns the index of a transform group in the agent’s definition.
Returns -1 if `prim` is out of range, `prim` is not an agent primitive, or the transform group does not exist.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transformgroup`

Name of a transform group in the agent’s definition.
