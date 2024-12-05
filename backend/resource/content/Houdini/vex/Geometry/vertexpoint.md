---
title: vertexpoint
order: 42
---
`int  vertexpoint(<geometry>geometry, int linearvertex)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`linearvertex`

The linear vertex number. The `vertexindex` function can be used
to compute a linear vertex from a primitive number and vertex
number pair.

Returns

The point number associated with the vertex, or `-1` if the vertex has no point.

Examples

## examples

```vex
int        pt;

// Get the point of vertex 3
pt = vertexpoint("defgeo.bgeo", 3);

```
