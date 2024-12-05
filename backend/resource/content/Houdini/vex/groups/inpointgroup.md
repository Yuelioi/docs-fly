---
title: inpointgroup
order: 5
---
`int  inpointgroup(<geometry>geometry, string groupname, int pointnum)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

`1` if the group exists and the point is in the group, or `0` otherwise.

This can use ad-hoc groups, like `0-3` or `@Cd.x>0.5`. It matches the SOP
group naming convention, in particular that an empty string means all points.
