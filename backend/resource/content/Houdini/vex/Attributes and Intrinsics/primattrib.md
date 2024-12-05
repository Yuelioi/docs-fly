---
title: primattrib
order: 47
---
`<type> primattrib(<geometry>geometry, string attribute_name, int prim, int &success)`

`<type>[] primattrib(<geometry>geometry, string attribute_name, int prim, int &success)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribute_name`

The name of the attribute (or intrinsic) to read.

`prim`

The primitive number.

`&success`

Set to `1` if the import was successful,
`0` on error (for example, the attribute or primitive number don’t exist).
