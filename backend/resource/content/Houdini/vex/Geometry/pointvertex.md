---
title: pointvertex
order: 22
---
`int  pointvertex(<geometry>geometry, int point_num)`

Use this to find linear vertex number of the first vertex to share this point.
Then you can use [vertexnext](vertexnext.html "Returns the linear vertex number of the next vertex sharing a point with a given vertex.") to iterate over the other vertices in the point.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

Returns the linear vertex number of the first vertex to share this point.
Returns `-1`if no vertices share this point.

Examples

## examples

```vex
int        vtx;

// Get the linear vertex  of point 3
vtx = pointvertex("defgeo.bgeo", 3);

```
