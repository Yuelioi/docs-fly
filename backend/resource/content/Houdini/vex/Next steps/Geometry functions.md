---
title: Geometry functions
order: 2
---
| On this page | * [Overview](#overview) * [Geometry traversal](#geometry-traversal) |
| --- | --- |

Overview

## overview

In Houdini, geometry is made up of *primitives*, with the most common primitive being a polygonal faces. Each corner of a polygonal face is a *vertex*. To keep individual faces together as part of a surface, vertices of separate faces can be fused together as a *point*. When a point moves, Houdini automatically moves each of the point’s vertices as well.

Geometry traversal

## geometry-traversal

- You can use [vertexpoint](functions/vertexpoint.html "Returns the point number of linear vertex in a geometry.") to get from a vertex to its point.
- You can use [pointvertex](functions/pointvertex.html "Returns a linear vertex number of a point in a geometry.") to get from a point to its first vertex and [vertexnext](functions/vertexnext.html "Returns the linear vertex number of the next vertex sharing a point with a given vertex.") and [vertexprev](functions/vertexprev.html "Returns the linear vertex number of the previous vertex sharing a point with a given vertex.") to iterate through the list of vertices sharing that point.
- To iterate through the vertices of a polygonal face, use [primvertexcount](functions/primvertexcount.html "Returns number of vertices in a primitive in a geometry.") to get the number of vertices, then you can use the primitive number and a counter from `0` to `primitive count - 1` to refer to the vertices.
- A *vertex number* is the 0-based index of a certain vertex of a primitive. A *linear vertex number* is the 0-based index of a certain primitive *across all the geometry*. It is just a slightly more convenient way to refer to a vertex with one number instead of two (primitive number and vertex number). You can use [vertexindex](functions/vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") to get a linear vertex number from a primitive number and a vertex number. Use both [vertexprim](functions/vertexprim.html "Returns the number of the primitive containing a given vertex.") and [vertexprimindex](functions/vertexprimindex.html "Converts a linear vertex index into a primitive vertex number.") to convert the other way, from a linear vertex number to a primitive number and a vertex number.

| Returns | Function | Argument(s) |
| --- | --- | --- |
| Point number | [vertexpoint](functions/vertexpoint.html "Returns the point number of linear vertex in a geometry.") | Linear vertex number |
| Linear vertex number | [pointvertex](functions/pointvertex.html "Returns a linear vertex number of a point in a geometry.") | Point number |
| Linear vertex number | [vertexnext](functions/vertexnext.html "Returns the linear vertex number of the next vertex sharing a point with a given vertex.") | Linear vertex number |
| Linear vertex number | [vertexprev](functions/vertexprev.html "Returns the linear vertex number of the previous vertex sharing a point with a given vertex.") | Linear vertex number |
| Linear vertex number | [vertexindex](functions/vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") | Primitive number,  Vertex number |
| Number of vertices | [primvertexcount](functions/primvertexcount.html "Returns number of vertices in a primitive in a geometry.") | Primitive number |
| Primitive number | [vertexprim](functions/vertexprim.html "Returns the number of the primitive containing a given vertex.") | Linear vertex number |
| Vertex number | [vertexprimindex](functions/vertexprimindex.html "Converts a linear vertex index into a primitive vertex number.") | Linear vertex number |
