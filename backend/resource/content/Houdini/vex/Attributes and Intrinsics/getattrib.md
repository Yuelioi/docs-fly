---
title: getattrib
order: 22
---
`<type> getattrib(<geometry>geometry, string attribclass, string attribute_name, int elemnum, int &success)`

`<type>[] getattrib(<geometry>geometry, string attribclass, string attribute_name, int elemnum, int &success)`

This general form lets you specify the attribute “class” at run-time. This can be useful for writing general code that can work on different classes.
If you know the class of attribute you want to read ahead of time, using [detailattrib](detailattrib.html "Reads a detail attribute value from a geometry."), [primattrib](primattrib.html "Reads a primitive attribute value from a geometry, outputting a success flag."), [pointattrib](pointattrib.html "Reads a point attribute value from a geometry and outputs a success/fail flag."), or [vertexattrib](vertexattrib.html "Reads a vertex attribute value from a geometry.") may be faster.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

`attribute_name`

The name of the attribute (or intrinsic) to read.

`elemnum`

The point/primitive/vertex number to read the attribute value from. For detail attributes, use `0` here (the argument is ignored for detail attributes).

To get the linear vertex number given a primitive number and the vertex number on the primitive, use the [primvertex](primvertex.html "Converts a primitive/vertex pair into a linear vertex.") function.

`success`

If the given attribute exists and can be read, the function sets this variable to `1`. Otherwise, it sets this variable to `0`.

Returns

The value of the attribute.
