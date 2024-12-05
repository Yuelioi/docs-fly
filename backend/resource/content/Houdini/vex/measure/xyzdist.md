---
title: xyzdist
order: 21
---
`float  xyzdist(<geometry>geometry, vector origin)`

Finds the distance from origin to the closest location on the given geometry.

`float  xyzdist(<geometry>geometry, vector origin, int &prim, vector &uv)`

`float  xyzdist(<geometry>geometry, vector origin, int &prim, vector &uv, float maxdist)`

Finds the distance from origin to the closest location on the geometry, and
writes the primitive number and UV coordinates of the closest location into the output arguments.

`float  xyzdist(<geometry>geometry, string primgroup, vector origin)`

`float  xyzdist(<geometry>geometry, string primgroup, vector origin, int &prim, vector &uv)`

`float  xyzdist(<geometry>geometry, string primgroup, vector origin, int &prim, vector &uv, float maxdist)`

Finds the distance from origin to the closest location in the given primitive group on the given geometry,
and writes the primitive number and UV coordinates of the closest location into the output arguments.

Note: Distances to packed primitives and sphere/tube/circle primitives with
non-uniform scales may not represent the actual closest point as the closest
point is found in the untransformed space.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primgroup`

The name of a primitive group or a pattern to generate a primitive
group. Uses the same semantics as a SOP group, so empty strings
will match all primitives. Attribute groups like `@Cd.x>0` can
also be used, but note that the `@` may need to be escaped with
a backslash in a [![](../../icons/COMMON/wrangle.svg)Snippet VOP](../../nodes/vop/snippet.html "Runs a VEX snippet to modify the incoming values.").

`origin`

The position in space to find the closest position on the geometry to.

`&prim`

The function overwrites this variable with the number of the closest primitive, or `-1` if no primitive was found.

`&uv`

The function overwrites this variable with the UV coordinates of the closest point on the closest primitive.
You can use [primuv](primuv.html "Interpolates the value of an attribute at a certain parametric (uvw) position.") to sample an attribute value at this location.

`maxdist`

The maximum distance to search. Specifying this can speed up the function by allowing it to quit early.

Returns

The distance from the origin point to the closest location on the geometry.
