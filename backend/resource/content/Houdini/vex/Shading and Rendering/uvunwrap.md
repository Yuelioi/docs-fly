---
title: uvunwrap
order: 79
---
`int  uvunwrap(string object_path, float u, float v, float time, vector &P, vector &I)`

`int  uvunwrap(string object_path, float u, float v, float time, vector &P, vector &I, vector &mikkelsenUtan, vector &mikkelsenVtan)`

This function **only makes sense in a Mantra context**, for use in **texture baking** or in a **lens shader**. The function unfortunately must be “context-less” so it’s available to the CVEX lens shader, but in any other context it will fail and return `0`.

For any other kind of texture sampling, use the superior [uvsample](uvsample.html "Interpolates the value of an attribute at certain UV coordinates using a UV attribute.") or [uvintersect](uvintersect.html "This function computes the intersection of the specified ray with the geometry in uv space.") functions instead of this.

Show/hide arguments

`object_path`

The object being unwrapped.

`u`, `v`

The UV coordinates specifying where on the surface to get the position and normal.

`time`

The time along the timeline at which to measure the geometry, in seconds.

`&P`

If it succeeds, the function overwrites this variable with the world space position of the given point.

`&I`

If it succeeds, the function overwrites this variable with the normal at the given point.

`&mikkelsenUtan`, `&mikkelsenVtan`

The function overwrites these variables with the Mikkelsen tangent vectors.

Returns

`1` if the UV coordinates specified a valid point on the surface, or `0` otherwise.
