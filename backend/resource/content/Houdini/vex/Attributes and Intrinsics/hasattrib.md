---
title: hasattrib
order: 24
---
If you know the attribute class ahead of time, using [hasdetailattrib](hasdetailattrib.html "Returns if a geometry detail attribute exists."), [hasprimattrib](hasprimattrib.html "Returns if a geometry prim attribute exists."), [haspointattrib](haspointattrib.html "Returns if a geometry point attribute exists."), or [hasvertexattrib](hasvertexattrib.html "Returns if a geometry vertex attribute exists.") may be faster.

`int  hasattrib(<geometry>geometry, string attribclass, string attribute_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

Returns `1` if the attribute exists, or `0` otherwise.

Examples

## examples

```vex
// Check whether the point group "pointstouse" exists.
if (hasattrib("defgeo.bgeo", "pointgroup", "pointstouse")) {
    // Do something with the point group
}

```
