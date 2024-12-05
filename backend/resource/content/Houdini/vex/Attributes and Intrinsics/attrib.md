---
title: attrib
order: 7
---
`<type> attrib(<geometry>geometry, string attribclass, string name, int elemnum)`

`<type>[] attrib(<geometry>geometry, string attribclass, string name, int elemnum)`

This general form lets you specify the attribute “class” at run-time. This can be useful for writing general code that can work on different classes.
If you know the class of attribute you want to read ahead of time, using [detail](detail.html "Reads the value of a detail attribute value from a geometry."), [prim](prim.html "Reads a primitive attribute value from a geometry."), [point](point.html "Reads a point attribute value from a geometry."), or [vertex](vertex.html "Reads a vertex attribute value from a geometry.") may be faster.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

`name`

The name of the attribute, group, or intrinsic to read from.

`elemnum`

Which element (e.g. point number, primitive number, vertex number) to read from. Ignored for detail attributes. You can use [vertexindex](vertexindex.html "Converts a primitive/vertex pair into a linear vertex.") to convert a primitive/point pair into a vertex number.

Returns

Zero/empty value if the attribute does not exist. Use [getattrib](getattrib.html "Reads an attribute value from geometry, with validity check.") if you want to check whether the attribute existed.
