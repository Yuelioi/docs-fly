---
title: uvsample
order: 76
---
This function specifies the position using UVs from a UV attribute. To use *intrinsic primitive UVs*, use [primuv](primuv.html "Interpolates the value of an attribute at a certain parametric (uvw) position.") instead.

`<type> uvsample(<geometry>geometry, string attr_name, string uv_attr_name, vector uvw)`

`<type>[] uvsample(<geometry>geometry, string attr_name, string uv_attr_name, vector uvw)`

`<type> uvsample(<geometry>geometry, string primgroup, string attr_name, string uv_attr_name, vector uvw)`

`<type>[] uvsample(<geometry>geometry, string primgroup, string attr_name, string uv_attr_name, vector uvw)`

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

`attr_name`

The name of a point or vertex attribute to sample. For primitive attributes, the value is taken from the primitive under the given UVs. **For point and vertex attributes, the value at the given UV coordinates will be interpolated from the surrounding points/vertices**. The values are taken from the “lowest” level at which an attribute with this name exists.

This must be a 3-float attribute.

`uv_attr_name`

The name of a point or vertex attribute containing UVs. The default UVs created by Houdini are in an attribute named `uv`. The named attribute can be 2D (UV) or 3D (UVW) in any vector type.

`uvw`

The position in UV(W) space at which to sample the attribute.
