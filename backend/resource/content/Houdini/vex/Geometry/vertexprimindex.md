---
title: vertexprimindex
order: 45
---
`int  vertexprimindex(<geometry>geometry, int linearindex)`

Note
To convert the linear index into a primitive number and primitive vertex number,
use [vertexprim](vertexprim.html "Returns the number of the primitive containing a given vertex.") and [vertexprimindex](vertexprimindex.html "Converts a linear vertex index into a primitive vertex number.").

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`linearindex`

The linear index of a vertex

Returns

The vertex’s number within the primitive that contains it, or
`-1` if the vertex has no primitive.

To get the primitive number of the containing primitive, use [vertexprim](vertexprim.html "Returns the number of the primitive containing a given vertex.").

Note
Due to the nature of the geometry structure, the first time this is run on
a geometry it has to run over all primitives to find the look up table.
This will be amortized out if most vertices are invoking this function.

Examples

## examples

```vex
int prim, vtx;

// Find the primitive and vertex offset of the linear vertex 6.
prim = vertexprim("defgeo.bgeo", 6);
vtx = vertexprimindex("defgeo.bgeo", 6);

```
