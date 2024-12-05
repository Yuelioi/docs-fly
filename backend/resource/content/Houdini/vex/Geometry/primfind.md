---
title: primfind
order: 25
---
`int [] primfind(<geometry>geometry, vector min, vector max)`

Find all the primitives whose bounding boxes overlap the given box.

`int [] primfind(<geometry>geometry, string group, vector min, vector max)`

Find all primitives in a group whose bounding boxes overlap the given box.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`min`, `max`

These vectors define the minimum and maximum corners of the bounding box to search.

`group`

If given, only match primitives in this group.
An empty group string will include all primitives.
The string supports Ad-hoc patterns like `0-10` and `@Cd.x>0`.

Returns

An array of primitive numbers.

Note
These functions are intended to be used as an optimization to finding primitives
in a particular area for processing. For instance, to find all the curves
from one input intersecting polygons on another input, we may naively iterate
over all polygons for each curve to determine their intersection. To speed this
process, we may find which primitives may intersect a particular curve using
these functions, and iterate solely over the potentially intersecting
primitives. This significantly improves performance since `primfind` uses an
underlying tree structure to speed up search.

Examples

## examples

Remove primitives that may be intersecting the unit box centered at the origin:

```vex
int[] prims = primfind(geometry, {-0.5, -0.5, -0.5}, {0.5, 0.5, 0.5});
foreach ( int prim; prims )
{
    removeprim("primitives.bgeo", prim, 1);
}

```

Alternatively, we can use a query bounding box from an auxiliary source:

```vex
vector min, max;
getbbox("bbox.bgeo", min, max);
int[] prims = primfind(geometry, min, max);
foreach ( int prim; prims )
{
    removeprim("primitives.bgeo", prim, 1);
}

```

To see the performance benefit of `primfind`, compare it to the following equivalent
implementation of the function above:

```vex
float tol = 1e-5;
vector min, max;
getbbox("bbox.bgeo",min,max);
int n = nprimitives(0);
for ( int prim = 0; prim < n; ++prim )
{
    int[] verts = primvertices("primitives.bgeo", prim);

    // compute primitive bounding box and store it in prim_min and prim_max
    vector vert_pos = point("primitives.bgeo", "P", vertexpoint("primitives.bgeo", verts[0]));
    vector prim_min = vert_pos, prim_max = vert_pos;
    for ( int v = 1; v < len(verts); ++v )
    {
        vert_pos = point("primitives.bgeo", "P", vertexpoint("primitives.bgeo", verts[v]));
        prim_min = min(prim_min, vert_pos);
        prim_max = max(prim_max, vert_pos);
    }

    // bounding box intersection test
    if ( prim_max.x - min.x < -tol ) continue;
    if ( prim_max.y - min.y < -tol ) continue;
    if ( prim_max.z - min.z < -tol ) continue;
    if ( prim_min.x - max.x > tol ) continue;
    if ( prim_min.y - max.y > tol ) continue;
    if ( prim_min.z - max.z > tol ) continue;
    removeprim("primitives.bgeo", prim, 1);
}

```
