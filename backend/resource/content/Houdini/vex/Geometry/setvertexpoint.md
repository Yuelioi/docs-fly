---
title: setvertexpoint
order: 37
---
`int  setvertexpoint(int geohandle, int prim, int vtxofprim, int pt)`

Rewires a specified vertex to a point number.

If the point number is -1, no rewiring is done.

If prim is -1, `vtxofprim` is treated as a linear index, and vice versa. Otherwise, the pair (`prim`, `vtxofprim`) is used to identify a vertex in a primitive’s vertex list.

This function is a new name for the equivalent function [setprimvertex](setprimvertex.html "Rewires a vertex in the geometry to a different point."), added for clarity.
