---
title: agentchannelvalues
order: 5
---
| Since | 18.0 |
| --- | --- |

`float [] agentchannelvalues(<geometry>geometry, int prim)`

If only a single channel value is needed, using [agentchannelvalue](agentchannelvalue.html "Returns the current value of an agent primitive’s channel.") instead can be significantly faster.

Returns an empty array if `prim` is out of range or is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.
