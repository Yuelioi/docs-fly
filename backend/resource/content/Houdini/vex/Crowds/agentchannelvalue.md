---
title: agentchannelvalue
order: 4
---
| Since | 18.0 |
| --- | --- |

`float  agentchannelvalue(<geometry>geometry, int prim, int channel)`

Returns zero if `prim` is out of range, `prim` is not an agent primitive, or `channel` is out of range.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`channel`

Index of a channel in the agent’s rig.
