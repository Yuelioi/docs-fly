---
title: pointhedgenext
order: 20
---
`int  pointhedgenext(<geometry>geometry, int point)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`point`

The point number in the geometry. `0` is the first point.

Returns

The next half-edge that has the same source as `hedge`.

Successive calls to this function iterate over all outgoing half-edges out of the same point.
The iteration order does not necessarily agree with the order of the edges around a point in a manifold setting.

Returns `-1` if `hedge` is not valid, or the there are no more shared vertices with the source vertex of this hedge (same as `op:vertexnext`).

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
}

```
