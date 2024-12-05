---
title: inprimgroup
order: 6
---
`int  inprimgroup(<geometry>geometry, string groupname, int primnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

`1` if the groups exists and the primitive is in the group, or `0` otherwise.

This can use ad-hoc groups, like `0-3` or `@Cd.x>0.5`. It matches the SOP
group naming convention, in particular that an empty string means all
primitives.
