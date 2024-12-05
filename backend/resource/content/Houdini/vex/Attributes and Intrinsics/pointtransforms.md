---
title: pointtransforms
order: 42
---
| Since | 18.5 |
| --- | --- |

`matrix [] pointtransforms(<geometry>geometry, int pnts[])`

`matrix [] pointtransforms(<geometry>geometry)`

Returns an array of transforms associated with the point indices, using the [standard instancing point attributes](../../copy/instanceattrs.html).
Returns all the point transforms if the point indices argument is omitted.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`pnts`

The array of point indices to query.
