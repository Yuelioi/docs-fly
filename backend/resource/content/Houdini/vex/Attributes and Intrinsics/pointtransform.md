---
title: pointtransform
order: 40
---
| Since | 18.5 |
| --- | --- |

`matrix  pointtransform(<geometry>geometry, int pnt)`

Computes a transform for the point, using the [standard instancing point attributes](../../copy/instanceattrs.html).

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`pnt`

The point index to query.
