---
title: pointtransformsrigid
order: 43
---
| Since | 18.5 |
| --- | --- |

`matrix [] pointtransformsrigid(<geometry>geometry, int pnts[])`

`matrix [] pointtransformsrigid(<geometry>geometry)`

Returns an array of rigid transforms associated with the point indices.
This function uses the [standard instancing point attributes](../../copy/instanceattrs.html) to build the matrix, and [polar decomposition](polardecomp.html "Computes the polar decomposition of a matrix.") is performed to make it rigid.
Returns all the point transforms if the point indices argument is omitted.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`pnts`

The array of point indices to query.
