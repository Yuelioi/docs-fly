---
title: uvintersect
order: 38
---
`int  uvintersect(<geometry>geometry, string uvname, vector orig, vector dir, vector &pos, vector &primuv)`

`int  uvintersect(<geometry>geometry, string primgroup, string uvname, vector orig, vector dir, vector &pos, vector &primuv)`

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

This function computes the intersection of the specified ray with the geometry in uv space. The primitive number is returned, or -1 if there is an error or no intersection found.

The position in uvspace of the intersection location is stored in p. The corresponding parametric location of the intersection is stored in primuv. In the case of multiple intersections, the intersection closest to the ray origin is used.

This function does not expect a normalized direction vector. Instead, it uses the length of the vector as the maximum distance. The integer result is the primitive hit.

Note
It can be hard to visualize a 3D intersection of a ray in a 3D UV space. One trick that can be used is to unwrap the geometry in SOP to get better visualization of the space. This can be done by using a [Split Vertex SOP](../../nodes/sop/splitvertex.html) followed by a [![](../../icons/SOP/attribcopy.svg)Attribute Copy SOP](../../nodes/sop/attribcopy.html "Copies attributes between groups of vertices,
points, or primitives."). This will disconnect the faces at uv boundaries and stamp the uvw values on top of the `P` attribute.

Note
When intersections are performed against metaball geometry, it is impossible to determine the primitive number of the metaball which was hit. In this case, the function returns the number of primitives in the intersection geometry.
