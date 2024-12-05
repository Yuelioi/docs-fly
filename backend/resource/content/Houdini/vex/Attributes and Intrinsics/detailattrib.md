---
title: detailattrib
order: 15
---
`<type> detailattrib(<geometry>geometry, string attribute_name, int ignored, int &success)`

`<type>[] detailattrib(<geometry>geometry, string attribute_name, int ignored, int &success)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`ignored`

Pass `0` for this argument.

`success`

The function sets this variable to `1` if the attribute was successfully read, or `0` otherwise.

Returns

`0` if importing the attribute failed, the value of the attribute on success.
