---
title: getbbox_max
order: 5
---
`vector  getbbox_max(<geometry>geometry)`

Computes the maximum of the bounding box for the geometry.

`vector  getbbox_max(<geometry>geometry, string primgroup)`

Computes the maximum of the bounding box of the primitives in the given group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
