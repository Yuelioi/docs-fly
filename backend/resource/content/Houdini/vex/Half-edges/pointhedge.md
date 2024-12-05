---
title: pointhedge
order: 19
---
`int  pointhedge(<geometry>geometry, int point)`

`int  pointhedge(<geometry>geometry, int srcpoint, int dstpoint)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`point`

The point number in the geometry for the source point of the returned half-edge. `0` is the first point.

`srcpoint`, `dstpoint`

The point numbers in the geometry for source and destination of returned half-edge. `0` is the first point.

Returns

The number of a half-edge that has `point` as source or has `srcpoint` as source and `dstpoint` as destination.
In the former case, using `op:pointhedgenext` one can enumerate over all the half-edges that have `point` as source.
Returns `-1` if the half-edge is not valid.

Examples

## examples

```vex
int edge_count = 0;

// Count number of *edges* (not half-edges) incident to point number 23.
int hout = pointhedge("defgeo.bgeo", 23);
while ( hout != -1 )
{
    if (hedge_isprimary("defgeo.bgeo", hout))
        edge_count++;
    int hin = hedge_prev("defgeo.bgeo", hout);
    if (hedge_isprimary("defgeo.bgeo", hin))
        edge_count++;
    hout = pointhedgenext("defgeo", hout);
};

```
