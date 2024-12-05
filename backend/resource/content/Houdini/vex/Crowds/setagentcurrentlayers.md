---
title: setagentcurrentlayers
order: 59
---
| Since | 19.0 |
| --- | --- |

`int  setagentcurrentlayers(int geohandle, int prim, string layernames[])`

`int  setagentcurrentlayers(int geohandle, int prim, int layer_ids[])`

Show/hide arguments

`geohandle`

Handle to the geometry to write to. `geoself()` can be used to get a handle to the current geometry.

`prim`

The primitive number.

`layernames`

A list of agent layer names.

`layer_ids`

A list of agent layer indices, as returned by [agentfindlayer](agentfindlayer.html "Finds the index of a layer in an agent’s definition.").
