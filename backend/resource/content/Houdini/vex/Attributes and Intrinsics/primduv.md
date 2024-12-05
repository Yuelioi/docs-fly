---
title: primduv
order: 51
---
`vector  primduv(<geometry>geometry, int prim_number, vector2 uv, int du, int dv)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim_number`

The number of the primitive on which to measure the derivative.

`uv`

The parametric coordinates on the primitive at which to measure the derivative.

`du`, `dv`

Represent the derivative order to query.

On a curve, the curve direction is given by `du==1` and the curvature is given by `du==2`.

`dv` is used for parametric surfaces.
