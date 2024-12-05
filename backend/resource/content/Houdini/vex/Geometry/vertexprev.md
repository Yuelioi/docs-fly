---
title: vertexprev
order: 43
---
`int  vertexprev(<geometry>geometry, int linearvertex)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`linearvertex`

The linear index of a vertex.
If you have a point number and point vertex number, you can use [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") to convert them to a linear index.

Returns

The linear index of the previous vertex sharing the same point with the given vertex,
or `-1` if the vertex has no earlier shared vertices.
(To go in the other direction, use [vertexnext](vertexnext.html "Returns the linear vertex number of the next vertex sharing a point with a given vertex.").)

Examples

## examples

```vex
int        vtx;

// Get the previous vertex of vertex 3
vtx = vertexprev("defgeo.bgeo", 3);

```
