---
title: agentlayershapes
order: 29
---
`string [] agentlayershapes(<geometry>geometry, int prim, string layername, string shapetype)`

`string [] agentlayershapes(<geometry>geometry, int prim, int layerindex, string shapetype)`

Returns the names of all shapes that are referenced by the layer and satisfy the `shapetype` filter.

`string [] agentlayershapes(<geometry>geometry, int prim, string layername, int transform)`

`string [] agentlayershapes(<geometry>geometry, int prim, int layerindex, int transform)`

Returns the names of all shapes that are referenced by the layer and are bound to the specified transform.

Returns an empty array if `layername` is not one of the agent’s [layers](agentlayers.html "Returns all of the layers that have been loaded for an agent primitive."), `shapetype` is invalid, `transform` is out of range, `prim` is out of range, or `prim` is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`layername`

The name of one of the agent’s layers.

`layerindex`

Index of a layer in the agent’s definition.
A layer’s index can be obtained via [agentfindlayer](agentfindlayer.html "Finds the index of a layer in an agent’s definition.").

`shapetype`

Whether to inspect `"static"`, `"deforming"`, or `"all"` shapes from the specified layer.

`transform`

Index of a transform in the agent’s rig.
