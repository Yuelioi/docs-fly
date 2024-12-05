---
title: getattribute
order: 23
---
`int  getattribute(string geometry, <type>&value, string attribclass, string attribute_name, int element_number, int vertex_number)`

`int  getattribute(string geometry, <type>&value[], string attribclass, string attribute_name, int element_number, int vertex_number)`

Warning

This function does not allow reading from inputs in a node context, and is harder to use than the other attribute functions. You probably want to use one of the other attribute functions instead, such as [getattrib](getattrib.html "Reads an attribute value from geometry, with validity check.").

Returns `0` if the attribute does not exist or cannot be read, `1` on success. If the function returns `0` (failed), the given variable may remain uninitialized.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`value`

The variable to overwrite with the attribute value.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

`attribute_name`

The name of the attribute (or intrinsic) to read.

`element_number`

The point or primitive number. If you are reading a detail attribute, use `0` here.

`vertex_number`

- When reading a vertex attribute, you can specify the primitive number in the `element_number` argument and the primitive’s vertex number here.
- To use a linear vertex index, use `-1` as the `element_number` and the vertex index here.
- If you are not reading a vertex attribute, this argument is ignored.

Examples

## examples

```vex
vector pos, uv, clr;
// Get the position of point 3 in "defgeo.bgeo"
getattribute("defgeo.bgeo", pos, "point", "P", 3, 0);

// Get the value of the "uv" attribute for vertex 2 of primitive
// number 3 in the file defgeo.bgeo
getattribute("defgeo.bgeo", uv, "vertex", "uv", 3, 2);

// Get the value of the "Cd" attribute for primitive 7
// in the SOP specified by the path "/obj/geo1/color1" (Houdini
// only)
getattribute("op:/obj/geo1/color1", clr, "primitive", "Cd", 7);

```
