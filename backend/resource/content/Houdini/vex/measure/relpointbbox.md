---
title: relpointbbox
order: 16
---
`vector  relpointbbox(<geometry>geometry, vector position)`

Returns the relative position of the point given with respect to the
bounding box of the points in the geometry.

`vector  relpointbbox(<geometry>geometry, string pointgroup, vector position)`

Use the bounding box of the primitives in the named point group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
