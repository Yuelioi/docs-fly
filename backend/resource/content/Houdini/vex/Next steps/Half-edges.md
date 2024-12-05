---
title: Half-edges
order: 3
---
| On this page | * [Overview](#overview) * [Equivalence](#equivalence) * [Functions](#functions) |
| --- | --- |

Overview

## overview

|  | In Houdini, edges are normally treated as *undirected* and *shared* between faces. However, for some tasks (such as simplifying or cleaning geometry) it is useful to view faces as each having their own edges. |
| --- | --- |
|  | You can think of this as splitting each shared undirected edge along its length into two *half edges*. (Boundary edges of course will only have one “half-edge”.) |
|  | Each half-edge is *directed* (it has a start vertex and an end vertex). |
|  | The half-edge functions let you quickly find a half-edge’s source and destination vertex, the next half-edge, get the other half-edges from the same edge, find all half-edges sharing a given point, and other manipulations. |

Equivalence

## equivalence

Two half-edges are *equivalent* if they are “split” from the same shared edge. In the following figure, three primitives share an edge. The half-edges *e*, *f*, and *g* are *equivalent* because they are from the same edge. Directionality does not matter – *e* is considered equivalent to *f* and *g* even though they have opposite directions.
![](../images/model/half_edge_equiv.png)
The source (and destination) vertex of a half-edge uniquely identify it, meaning that there can be at most one half-edge with a given vertex as the source. This is because a vertex can belong to only one primitive. However, because several vertices can be wired to the same point, half-edges can have shared source and destination *points*. Another way of thinking about equivalence is that two half-edges are equivalent if their vertices are shared by the same two points.

One of a set of equivalent half-edges is considered the *primary* half-edge in that class. The VEX functions let you check if a given half-edge is primary, or given a non-primary half-edge, find the equivalent primary half-edge. In the case of an unshared edge, it only has a single primary half-edge.

Functions

## functions

| [hedge_isvalid](functions/hedge_isvalid.html "Determines whether a half-edge number corresponds to a valid half-edge.") | Check whether a half-edge is valid |
| --- | --- |
| [hedge_isprimary](functions/hedge_isprimary.html "Determines whether a half-edge number corresponds to a primary half-edge.") | Check whether a half-edge is primary |
| [hedge_primary](functions/hedge_primary.html "Returns the primary half-edge equivalent to a given half-edge.") | Find the primary equivalent half-edge. If it belongs to an unshared edge, then the same half-edge is returned. |
| [hedge_isequiv](functions/hedge_isequiv.html "Determines whether a two half-edges are equivalent (represent the same edge).") | Test whether two half-edges are equivalent |
| [hedge_nextequiv](functions/hedge_nextequiv.html "Returns the next half-edges equivalent to a given half-edge.") | Iterate over the half-edges in an equivalence class |
| [hedge_equivcount](functions/hedge_equivcount.html "Returns the number of half-edges equivalent to a given half-edge.") | Get the number of half-edges in an equivalence class |
| [hedge_srcvertex](functions/hedge_srcvertex.html "Returns the source vertex of a half-edge.") | Get the source vertex of a half-edge |
| [hedge_dstvertex](functions/hedge_dstvertex.html "Returns the destination vertex of a half-edge.") | Get the destination vertex of a half-edge |
| [hedge_presrcvertex](functions/hedge_presrcvertex.html "Returns the vertex that precedes the source vertex of a half-edge in its primitive.") | Get the vertex preceding the source vertex in half-edge primitive |
| [hedge_postdstvertex](functions/hedge_postdstvertex.html "Returns the vertex following the destination vertex of a half-edge in its primitive.") | Get the vertex following the destination vertex in half-edge primitive |
| [hedge_prim](functions/hedge_prim.html "Returns the primitive that contains a half-edge.") | Get the primitive of a half-edge |
| [hedge_srcpoint](functions/hedge_srcpoint.html "Returns the source point of a half-edge.") | Get the source point of a half-edge |
| [hedge_dstpoint](functions/hedge_dstpoint.html "Returns the destination point of a half-edge.") | Get the destination point of a half-edge |
| [hedge_presrcpoint](functions/hedge_presrcpoint.html "Returns the point into which the vertex that precedes the source vertex of a half-edge in its primitive is wired.") | Get the pre-source point of a half-edge |
| [hedge_postdstpoint](functions/hedge_postdstpoint.html "Returns the point into which the vertex following the destination vertex of a half-edge in its primitive is wired.") | Get the post-destination point of a half-edge |
| [hedge_next](functions/hedge_next.html "Returns the half-edge that follows a given half-edge in its polygon.") | Get the next half-edge in primitive |
| [hedge_prev](functions/hedge_prev.html "Returns the half-edge that precedes a given half-edge in its polygon.") | Get the previous half-edge in primitive |
| [pointedge](functions/pointedge.html "Finds and returns a half-edge with the given endpoints.") | Find half-edge between two points |
| [pointhedge](functions/pointhedge.html "Finds and returns a half-edge with a given source point or with given source and destination points.") | Find half-edge with given source and destination |
| [pointhedge](functions/pointhedge.html "Finds and returns a half-edge with a given source point or with given source and destination points.") | Find a half-edge with a given source point |
| [pointhedgenext](functions/pointhedgenext.html "Returns the next half-edge with the same source as a given half-edge.") | Iterate over half-edges with a given source point |
| [vertexhedge](functions/vertexhedge.html "Returns the half-edge which has a vertex as source.") | Get the half-edge with a source vertex |
| [primhedge](functions/primhedge.html "Returns one of the half-edges contained in a primitive.") | Get a half-edge on a given primitive |
