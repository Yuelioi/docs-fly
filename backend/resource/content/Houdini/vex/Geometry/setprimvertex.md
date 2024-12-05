---
title: setprimvertex
order: 36
---
`int  setprimvertex(int geohandle, int prim, int vtxofprim, int pt)`

Rewires a specified vertex to a point number.

If the point number is -1, no rewiring is done.

If prim is -1, `vtxofprim` is treated as a linear index, and vice versa. Otherwise, the pair (`prim`, `vtxofprim`) is used to identify a vertex in a primitive’s vertex list.

Since this sets a vertex’s point, not a primitive’s vertex, it’s recommended to use the equivalent function [setvertexpoint](setvertexpoint.html "Rewires a vertex in the geometry to a different point.") for clarity, instead.
