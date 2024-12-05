---
title: getpointbbox_min
order: 12
---
`vector  getpointbbox_min(<geometry>geometry)`

`vector  getpointbbox_min(<geometry>geometry, string pointgroup)`

This is the same as [getbbox_min](getbbox_min.html "Returns the minimum of the bounding box for the geometry.") except it only computes the bounding box of the *points*. So if a primitive has extents that don’t have points (for example, the boundary of a primitive sphere), they will not be included in the box.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
