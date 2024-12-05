---
title: intersect
order: 8
---
To get a list of *all* intersections along a ray, use [intersect_all](intersect_all.html "Computes all intersections of the specified ray with geometry.").

`int  intersect(<geometry>geometry, vector orig, vector dir, vector &p, float &u, float &v)`

`int  intersect(<geometry>geometry, vector orig, vector dir, vector &p, float &u, float &v, ...)`

Computes the first intersection of the specified ray with the geometry.
To get *all* intersections along a vector, use [intersect_all](intersect_all.html "Computes all intersections of the specified ray with geometry.") instead.
The variadic argument `"farthest"` can be used to indicate whether to return the last intersection instead of the first.

`int  intersect(<geometry>geometry, vector orig, vector dir, vector &p, vector &uvw)`

Computes the first intersection of the specified ray with the geometry.
To get *all* intersections along a vector, use [intersect_all](intersect_all.html "Computes all intersections of the specified ray with geometry.") instead.

`int  intersect(<geometry>geometry, string group, vector orig, vector dir, vector &p, vector &uvw)`

Computes the intersection of the specified ray with primitives in the given group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`group`

If given, only intersect primitives in this group.

`orig`

The ray origin point.

`dir`

The ray direction *and maximum distance*.
This function does not expect a normalized direction vector.
Instead, it uses the length of the vector as the maximum distance to search.

`&p`

If the ray intersects a primitive, this variable is overwritten with the intersection position in world space.

`&u`, `&v`, `&uvw`

If the ray intersects a primitive, this/these variable(s) is/are overwritten with the parametric intersection position on the primitive.

Returns

The intersected primitive number, or `-1` if there was an error or the ray didn’t intersect anything.

Note
When intersections are performed against metaball geometry, it is
impossible to determine the primitive number of the metaball which
was hit. In this case, the function returns the number of primitives
in the intersection geometry.

Examples

## examples

```vex
// Intersect against the second input's geometry, using a ray at the current
// point's position and in the direction of its velocity vector.
vector origin = @P;
float max_dist = 1000;
vector dir = max_dist * normalize(@v);

vector isect_pos;
float isect_u, isect_v;
int isect_prim = intersect(@OpInput2, origin, dir, isect_pos, isect_u, isect_v);

// Return the farthest intersection instead.
isect_prim = intersect(@OpInput2, origin, dir, isect_pos, isect_u, isect_v, "farthest", 1);

```
