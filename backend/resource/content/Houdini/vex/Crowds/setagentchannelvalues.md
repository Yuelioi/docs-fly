---
title: setagentchannelvalues
order: 51
---
| Since | 18.0 |
| --- | --- |

`void  setagentchannelvalues(int geohandle, int prim, float values[])`

When modifying a single channel, using [setagentchannelvalue](setagentchannelvalue.html "Overrides the value of an agent primitive’s channel.") instead can be significantly faster.

Show/hide arguments

`geohandle`

Handle to the geometry to write to. `geoself()` can be used to get a handle to the current geometry.

`prim`

The primitive number.

`values`

The new value of each channel in the agent’s rig.
