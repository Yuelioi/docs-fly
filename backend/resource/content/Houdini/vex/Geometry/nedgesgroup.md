---
title: nedgesgroup
order: 13
---
`int  nedgesgroup(<geometry>geometry, string groupname)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`groupname`

A group name or ad-hoc group, like `0` or `p0-1`. This matches the
SOP group naming convention, in particular that an empty string means all
edges.
