---
title: vertexindex
order: 40
---
`int  vertexindex(<geometry>geometry, int primnum, int vertex)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primnum`

The primitive number to get a vertex from.

`vertex`

The vertex number inside of a primitive. 0 is the first vertex.

Returns

The linear vertex index corresponding to the given primitive vertex.
Returns `-1` if the function could not find the equivalent linear vertex index.

Examples

## examples

```vex
int linearvtx;

// Get the linear vertex value of vertex 2 of primitive 3.
linearvtx = vertexindex("defgeo.bgeo", 3, 2);

```
