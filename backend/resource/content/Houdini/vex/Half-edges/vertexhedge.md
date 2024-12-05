---
title: vertexhedge
order: 22
---
`int  vertexhedge(<geometry>geometry, int vertex)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`vertex`

The linear vertex number in the geometry. `0` is the first vertex.

Returns

The number for the half-edge that has `vertex` as source and the vertex following `vertex` in the primitive of `vertex` as destination.
Returns `-1` if failed to find the corresponding vertex.

Examples

## examples

```vex
int vtxhedge;

// Get the hedge out of vertex vertex number 2.
vtxhedge = vertexhedge("defgeo.bgeo", 2);

```
