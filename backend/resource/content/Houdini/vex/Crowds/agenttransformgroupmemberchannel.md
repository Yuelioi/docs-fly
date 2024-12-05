---
title: agenttransformgroupmemberchannel
order: 42
---
| Since | 18.0 |
| --- | --- |
See [Transform Groups](../../crowds/agents.html#xformgroups) for more information.

`int  agenttransformgroupmemberchannel(<geometry>geometry, int prim, int transformgroupidx, int channel)`

Returns a non-zero value if the channel is a member of the specified transform group, and zero otherwise.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transformgroupidx`

Index of a transform group in the agent’s definition.
A transform group’s index can be obtained via [agentfindtransformgroup](agentfindtransformgroup.html "Finds the index of a transform group in an agent’s definition.").

`channel`

Index of a channel in the agent’s rig.
