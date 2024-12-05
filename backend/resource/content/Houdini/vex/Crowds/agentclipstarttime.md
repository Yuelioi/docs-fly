---
title: agentclipstarttime
order: 16
---
| Since | 18.0 |
| --- | --- |

`float  agentclipstarttime(<geometry>geometry, int prim, string clipname)`

`float  agentclipstarttime(<geometry>geometry, int prim, int clipindex)`

Returns `0` if `prim` is out of range or is not an agent primitive, or if `clipname` is not one of the agent’s [animation clips](agentclipcatalog.html "Returns all of the animation clips that have been loaded for an agent primitive.").

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
