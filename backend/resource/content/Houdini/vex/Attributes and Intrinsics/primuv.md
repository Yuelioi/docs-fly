---
title: primuv
order: 54
---
This function specifies the position using *intrinsic primitive UVs*. To use UVs stored in UV attribute, use [uvsample](uvsample.html "Interpolates the value of an attribute at certain UV coordinates using a UV attribute.") instead.

`<type> primuv(<geometry>geometry, string attribute_name, int prim_num, vector uvw)`

`<type>[] primuv(<geometry>geometry, string attribute_name, int prim_num, vector uvw)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute to read. **For point and vertex attributes, the value at the given UV coordinates will be interpolated** from the surrounding points/vertices.

`prim_num`

The primitive number to read the attribute from.

`uvw`

The primitive UVW coordinates at which to read the attribute.

- Returns the (possibly interpolated) value of the attribute at the given coordinates. If the attribute or primitive number doesn’t exist, returns `0`.
- If you need to test for errors, you can use [prim_attribute](prim_attribute.html "Interpolates the value of an attribute at a certain parametric (u, v) position and copies it into a variable.") instead.
