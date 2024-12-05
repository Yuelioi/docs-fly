---
title: agentrigfindchannel
order: 37
---
| Since | 18.0 |
| --- | --- |

`int  agentrigfindchannel(<geometry>geometry, int prim, string channelname)`

Returns `-1` if `channelname` was not found in the rig, `prim` is out of range, or `prim` is not an agent primitive.
This index can be used to access an agent’s current channel value using [agentchannelvalue](agentchannelvalue.html "Returns the current value of an agent primitive’s channel.") and [setagentchannelvalue](setagentchannelvalue.html "Overrides the value of an agent primitive’s channel."), or to sample the channel value from any clip using [agentclipsample](agentclipsample.html "Samples a channel of an agent’s clip at a specific time.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`channelname`

The name of a channel in the agent’s rig.
