---
title: hasvertexattrib
order: 28
---
`int  hasvertexattrib(<geometry>geometry, string attribute_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns `1` if the attribute exists, or `0` otherwise.

Examples

## examples

```vex
// Determine if the P attribute exists.
int exists = hasvertexattrib("defgeo.bgeo", "P");

```
