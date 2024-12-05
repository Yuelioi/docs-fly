---
title: getbbox_min
order: 6
---
`vector  getbbox_min(<geometry>geometry)`

Computes the minimum of the bounding box for the geometry.

`vector  getbbox_min(<geometry>geometry, string primgroup)`

Computes the minimum of the bounding box of the primitives in the given group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
