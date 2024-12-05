---
title: prim_attribute
order: 45
---
This function specifies the position using *intrinsic primitive UVs*. To use UVs stored in UV attribute, use [uvsample](uvsample.html "Interpolates the value of an attribute at certain UV coordinates using a UV attribute.") instead.

`int  prim_attribute(<geometry>geometry, <type>&value, string attribute_name, int prim_number, float u, float v)`

`int  prim_attribute(<geometry>geometry, <type>&value[], string attribute_name, int prim_number, float u, float v)`

Samples the attribute value at the given UV coordinates on the primitive.

`int  prim_attribute(<geometry>geometry, <type>&value, string attribute_name, int prim_number, vector uvw)`

`int  prim_attribute(<geometry>geometry, <type>&value[], string attribute_name, int prim_number, vector uvw)`

Specify the UVW coordinates using a vector instead of two floats.

If you don’t need to test for errors, you can use [primuv](primuv.html "Interpolates the value of an attribute at a certain parametric (uvw) position.") instead.
This function does not work with certain primitive types such as tetrahedra and polysoups.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`value`

The function overwrites this variable with the interpolated value from the primitive.

`attribute_name`

The name of the attribute to read. **For point and vertex attributes, the value will at the given UV coordinates will be interpolated** from the surrounding points/vertices.

`prim_number`

The primitive number to read the attribute from.

`u`, `v`

The primitive UV coordinates at which to read the attribute.

Returns

Returns `1` on success or `0` on an error (for example, the attribute doesn’t exist).

Returns `0` if the type of `value` is larger than the primitive type. For example, you can’t read a vector attribute into a matrix variable.
