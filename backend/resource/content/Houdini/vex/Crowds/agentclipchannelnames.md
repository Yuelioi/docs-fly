---
title: agentclipchannelnames
order: 8
---
`string [] agentclipchannelnames(<geometry>geometry, int prim, string clipname)`

`string [] agentclipchannelnames(<geometry>geometry, int prim, int clipindex)`

Returns an empty array if `clipname` is not one of the agent’s [animation clips](agentclipcatalog.html "Returns all of the animation clips that have been loaded for an agent primitive."), `prim` is out of range, or `prim` is not an agent primitive.

For a list of the agent’s transforms, use [agenttransformnames](agenttransformnames.html "Returns the name of each transform in an agent primitive’s rig.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`clipname`

The name of the animation clip.

`clipindex`

Index of a clip in the agent’s definition.
A clip’s index can be obtained via [agentfindclip](agentfindclip.html "Finds the index of a clip in an agent’s definition.").
