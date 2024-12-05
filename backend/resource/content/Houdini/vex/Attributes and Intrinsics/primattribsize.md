---
title: primattribsize
order: 48
---
`int  primattribsize(<geometry>geometry, string attribute_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

See [attribsize](attribsize.html "Returns the size of a geometry attribute.") for more information.

Returns `0` if the attribute cannot be found.
