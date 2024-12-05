---
title: vertexprim
order: 44
---
`int  vertexprim(<geometry>geometry, int linearvertex)`

Note
To convert the linear index into a primitive number and primitive vertex number,
use [vertexprim](vertexprim.html "Returns the number of the primitive containing a given vertex.") and [vertexprimindex](vertexprimindex.html "Converts a linear vertex index into a primitive vertex number.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`linearvertex`

The linear index of a vertex.
If you have a point number and point vertex number, you can use [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") to convert them to a linear index.

Returns

The primitive number of the primitive containing the vertex,
or `-1` if the vertex has no primitive.

Examples

## examples

```vex
int        pt;

// Get the primitive of vertex 3
pt = vertexprim("defgeo.bgeo", 3);

```
