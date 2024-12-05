---
title: pointlocaltransforms
order: 39
---
| Since | 18.5 |
| --- | --- |

`matrix [] pointlocaltransforms(<geometry>geometry, int pnts[])`

Returns an array of local transforms associated with the point indices. This function queries the `4@localtransform` attribute.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`pnts`

The array of point indices to query.
