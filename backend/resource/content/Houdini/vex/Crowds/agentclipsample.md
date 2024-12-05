---
title: agentclipsample
order: 12
---
`float  agentclipsample(<geometry>geometry, int prim, string clipname, float time, int channel_index)`

`float  agentclipsample(<geometry>geometry, int prim, string clipname, float time, string channel)`

`float  agentclipsample(<geometry>geometry, int prim, int clipindex, float time, int channel_index)`

`float  agentclipsample(<geometry>geometry, int prim, int clipindex, float time, string channel)`

Evaluates the clip at the given time and returns the value of the specified channel.
Returns zero if `clipname` is not one of the agent’s [animation clips](agentclipcatalog.html "Returns all of the animation clips that have been loaded for an agent primitive."), `prim` is out of range, `prim` is not an agent primitive, `channel_index` is out of range, or `channel` does not exist.

For sampling the clip’s transform channels, use [agentclipsamplelocal](agentclipsamplelocal.html "Samples an agent’s animation clip at a specific time.") or [agentclipsampleworld](agentclipsampleworld.html "Samples an agent’s animation clip at a specific time.") instead.

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

`channel_index`

Index of a channel in the animation clip, as returned by [agentclipchannel](agentclipchannel.html "Finds the index of a channel in an agent’s animation clip.").

`channel`

Name of a channel in the animation clip.

Examples

## examples

Sample a channel of the walk clip after 1.2 seconds.

```vex
float value = agentclipsample(0, @primnum, "walk", 1.2, "latch_leftfoot");

```
