---
title: osd_facecount
order: 1
---
`int  osd_facecount(<geometry>geometry)`

Returns the number of coarse faces in the subdivision hull. This is different that the number of patches in the subdivision surface.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.
