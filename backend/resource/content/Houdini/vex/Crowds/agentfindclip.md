---
title: agentfindclip
order: 24
---
`int  agentfindclip(<geometry>geometry, int prim, string clipname)`

Returns the index of a clip in the agent’s definition.
Returns -1 if `prim` is out of range, `prim` is not an agent primitive, or the clip does not exist.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`clipname`

Name of a clip in the agent’s definition.
