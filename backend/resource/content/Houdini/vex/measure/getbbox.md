---
title: getbbox
order: 3
---
`void  getbbox(<geometry>geometry, vector &min, vector &max)`

Sets the vectors to the minimum and maximum corners of the bounding box
for the geometry. This outputs the primitive bounding box, which includes
the extents of spheres and volumes.

`void  getbbox(<geometry>geometry, string primgroup, vector &min, vector &max)`

Outputs the bounding box of the primitive in the given group.
An empty primgroup string will include all primitives.
The string supports Ad-hoc patterns like `0-10` and `@Cd.x>0`.

`void  getbbox(vector &min, vector &max)`

Warning
This form of `getbbox` is deprecated and may be removed in the future.
Use the other forms as needed.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
