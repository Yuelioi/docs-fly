---
title: getbbox_size
order: 7
---
`vector  getbbox_size(<geometry>geometry)`

Computes the size of the bounding box for the geometry.

`vector  getbbox_size(<geometry>geometry, string primgroup)`

Computes the size of the bounding box of the primitives in the given group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
