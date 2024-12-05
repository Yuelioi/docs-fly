---
title: primarclen
order: 46
---
`float  primarclen(<geometry>geometry, vector2 uv1, vector2 uv2, int prim_num)`

`float  primarclen(<geometry>geometry, vector2 uv1, vector2 uv2, int prim_num, int divs)`

`float  primarclen(<geometry>geometry, vector2 uv1, vector2 uv2, int prim_num, int divs, int primuvmode)`

`float  primarclen(<geometry>geometry, vector2 uv1, vector2 uv2, int prim_num, int divs, int primuvmode, float primuvtol)`

Returns the arc length between two parametric UV coordinates on a given primitive. This lets you measure the distance across a polygon face or along a curve.

Show/hide arguments

`geometry`

A string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`uv1`

The start coordinate in the primitive’s parametric space to measure between.

`uv2`

The end coordinate in the primitive’s parametric space to measure between.

`prim_num`

The number of the primitive across which to measure the distance.

`divs`

The number of divisions per segment to use or 10 if not supplied.

`primuvmode`

Define the uv1 and uv2 coordinates units. See [primuvconvert](primuvconvert.html "Convert parametric UV locations on curve primitives between different spaces.") for the list of modes.

`primuvtol`

A tolerance used when computing the curve length to do uv coordinates conversions.

Tip

You can also read the `arclength` primitive intrinsic attribute to get a curve’s total arc length.
