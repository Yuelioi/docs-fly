---
title: agentclipsampleworld
order: 15
---
`matrix [] agentclipsampleworld(<geometry>geometry, int prim, string clipname, float time)`

`matrix [] agentclipsampleworld(<geometry>geometry, int prim, int clipindex, float time)`

`matrix  agentclipsampleworld(<geometry>geometry, int prim, string clipname, float time, int transform)`

`matrix  agentclipsampleworld(<geometry>geometry, int prim, int clipindex, float time, int transform)`

Evaluates the clip at the given time and returns the world transforms of the agent’s rig.
Returns an empty array if `clipname` is not one of the agent’s [animation clips](agentclipcatalog.html "Returns all of the animation clips that have been loaded for an agent primitive."), `prim` is out of range, `prim` is not an agent primitive, or `transform` is [out of range](agenttransformcount.html "Returns the number of transforms in an agent primitive’s rig.").
The `matrix` signature is more efficient than the `matrix[]` signature for sampling a single transform.

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

`time`

The time (in seconds) to evaluate the clip at. If this time is greater than the [clip’s length](agentcliplength.html "Returns the length (in seconds) of an agent’s animation clip."), it will be wrapped around.

`transform`

Index of a transform in the agent’s rig.

Examples

## examples

Sample the world transforms of the walk clip after 1.2 seconds.

```vex
matrix xforms[] = agentclipsampleworld(0, @primnum, "agent1_clip.walk", 1.2);

```
