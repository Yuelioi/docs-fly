---
title: intersect_all
order: 9
---
To get only the *first* intersection, use [intersect](intersect.html "This function computes the first intersection of a ray with geometry.").

`int  intersect_all(<geometry>geometry, string group, vector orig, vector dir, vector &pos[], int &prim[], vector &uvw[], float tol=0.01, float ttol=0.01 )`

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

`&pos`

The function overwrites this array with the world space positions of each hit.

`&prim`

The function overwrites this array with the primitive numbers of the primitives hit by the ray.

`&uvw`

The function overwrites this array with the parametric UVW coordinates of where each intersection occurred on the primitive.

`tol`, `ttol`

`tol` is the 3D tolerance. `ttol` is the ray tolerance.
Collision points within the parametric ray tolerance, `ttol` will be merged
together, often useful to avoid getting extra intersects when hitting the edges
of geometry.

To get *all* intersections without merging, set `ttol` to `-1`.

Returns

The number of intersections, or `0` if the ray didn’t hit anything.

Note
When intersections are performed against metaball geometry, it is
impossible to determine the primitive number of the metaball which
was hit. In this case, the function returns the number of primitives
in the intersection geometry.
