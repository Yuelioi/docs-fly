---
title: agentlayerbindings
order: 27
---
`int [] agentlayerbindings(<geometry>geometry, int prim, string layername, string shapetype)`

`int [] agentlayerbindings(<geometry>geometry, int prim, int layerindex, string shapetype)`

Returns an empty array if `layername` is not one of the agent’s [layers](agentlayers.html "Returns all of the layers that have been loaded for an agent primitive."), `shapetype` is invalid, `prim` is out of range, or `prim` is not an agent primitive.

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

Examples

## examples

Find the current world transform of each static shape in the collision layer.

```vex
string layer = agentcollisionlayer(0, @primnum);
int[] bindings = agentlayerbindings(0, @primnum, layer, "static");
matrix xforms[] = agentworldtransforms(0, @primnum);

foreach (int idx; bindings) {
matrix xform = xforms[idx];
}

```
