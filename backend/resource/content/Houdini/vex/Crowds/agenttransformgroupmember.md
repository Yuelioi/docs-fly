---
title: agenttransformgroupmember
order: 41
---
See [Transform Groups](../../crowds/agents.html#xformgroups) for more information.

`int  agenttransformgroupmember(<geometry>geometry, int prim, string transformgroup, int transform)`

`int  agenttransformgroupmember(<geometry>geometry, int prim, int transformgroupidx, int transform)`

Returns a non-zero value if the transform is a member of the specified transform group, and zero otherwise.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transformgroup`

Name of a transform group in the agent’s definition.

`transformgroupidx`

Index of a transform group in the agent’s definition.
A transform group’s index can be obtained via [agentfindtransformgroup](agentfindtransformgroup.html "Finds the index of a transform group in an agent’s definition.").

`transform`

Index of a transform in the agent’s rig.
