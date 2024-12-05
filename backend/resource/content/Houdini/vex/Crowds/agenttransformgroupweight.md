---
title: agenttransformgroupweight
order: 44
---
See [Transform Groups](../../crowds/agents.html#xformgroups) for more information.

`float  agenttransformgroupweight(<geometry>geometry, int prim, int transformgroup, int transform)`

Returns the transform’s weight if it is a member of the specified transform group, and zero otherwise.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transformgroup`

Index of a transform group in the agent’s definition.
A transform group’s index can be obtained via [agentfindtransformgroup](agentfindtransformgroup.html "Finds the index of a transform group in an agent’s definition.").

`transform`

Index of a transform in the agent’s rig.
