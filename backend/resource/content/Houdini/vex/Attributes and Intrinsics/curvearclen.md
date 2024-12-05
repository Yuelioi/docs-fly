---
title: curvearclen
order: 13
---
| Since | 18.5 |
| --- | --- |

`float  curvearclen(vector positions[], float uv1, float uv2, int closedflag, int fmt, int order)`

`float  curvearclen(vector positions[], float uv1, float uv2, int closedflag, int fmt, int order, int divs)`

`float  curvearclen(vector positions[], float uv1, float uv2, int closedflag, int fmt, int order, int divs, int primuvmode)`

`float  curvearclen(vector positions[], float uv1, float uv2, int closedflag, int fmt, int order, int divs, int primuvmode, float primuvtol)`

Returns the arc length between two parametric UV coordinates on a given primitive. This lets you measure the distance across a polygon face or along a curve.

Show/hide arguments

`geometry`

A string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`points`

The array of control points to define the curve.

`uv1`

The start coordinate in the primitive’s parametric space to measure between.

`uv2`

The end coordinate in the primitive’s parametric space to measure between.

`closedflag`

The curve can be closed or open. In a closed curve, the last control point will be connected to the first one.

`fmt`

The curve type to create. You can used constants defined in math.h, or 0 to create a polygon curve, 1 to create a bezier curve or 2 to create a NURBS curve.

`order`

The curve order for NURBS or Bezier curves. This is ignored for polygon curves.

`divs`

The number of divisions per segment to use or 10 if not supplied.

`primuvmode`

Define the uv1 and uv2 coordinates units. See [primuvconvert](primuvconvert.html "Convert parametric UV locations on curve primitives between different spaces.") for the list of modes.

`primuvtol`

A tolerance used when computing the curve length to do uv coordinates conversions.
