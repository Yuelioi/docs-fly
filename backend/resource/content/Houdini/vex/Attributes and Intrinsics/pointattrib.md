---
title: pointattrib
order: 35
---
`<type> pointattrib(<geometry>geometry, string attribute_name, int pointnumber, int &success)`

`<type>[] pointattrib(<geometry>geometry, string attribute_name, int pointnumber, int &success)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`&success`

The function overwrites this variable with `1` if the attribute exists and was read successfully, or `0` otherwise.

Returns

The value of the given attribute on the given point number, or `0` if the attribute or point do not exist.
