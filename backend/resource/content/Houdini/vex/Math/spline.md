---
title: spline
order: 79
---
`float  spline(string basis, float sample_pos, float value1, ...)`

`vector  spline(string basis, float sample_pos, vector value1, ...)`

`vector4  spline(string basis, float sample_pos, vector4 value1, ...)`

This version takes a single basis to use for all keys, and takes the (linearly spaced) key values as variadic arguments.

`float  spline(string basis, float sample_pos, float values[], ...)`

`vector  spline(string basis, float sample_pos, vector values[], ...)`

`vector4  spline(string basis, float sample_pos, vector4 values[], ...)`

This version takes a single basis to use for all keys, and takes the (linearly spaced) key values as an array.

`float  spline(string bases[], float sample_pos, float values[], ...)`

`vector  spline(string bases[], float sample_pos, vector values[], ...)`

`vector4  spline(string bases[], float sample_pos, vector4 values[], ...)`

This version takes an array specifying the bases to use between each pair of keys, and the (linearly spaced) key values as an array.

`float  spline(string bases[], float sample_pos, float values[], float positions[], ...)`

`vector  spline(string bases[], float sample_pos, vector values[], float positions[], ...)`

`vector4  spline(string bases[], float sample_pos, vector4 values[], float positions[], ...)`

This version takes an array specifying the bases to use between each pair of keys, an array of key values, and an array of key positions.

These forms take an array of strings specifying the interpolation
bases between the keys, an array of key values, and an array of key positions.
They ensure that the interpolation curve is smooth (tangent-continuous) across
the control points (keys) if the adjoining segments have the same basis, even if
the key positions are not evenly spaced (i.e., are non-uniform and the distances
between them are not equal).

Show/hide arguments

`basis`, `bases`

These are the same interpolations supported by ramp parameters.

`"constant"`

Maintains each key value until the next key, creating a “stair step” curve.

`"linear"`

Connects the key points with a polyline.

For example, if you specified four values:

```vex
spline("linear", t, v0, v1, v2, v3)

```

![](../../images/vex/spline_linear.svg)
…the function returns the height of the orange dot at position sample_pos.

`"cubic"` (or `"catmullrom"`, `"cspline"`)

Connect the point values with a Catmull-Rom spline.

Note that the first and last values are outside the sample area to
provide the slope of the curve at the second point (at the start of the
sample range) and the second-to-last point (at the end of the sample
range).

For example, if you specified six values:

```vex
spline("catrom", t, v0, v1, v2, v3, v4, v5)

```

![](../../images/vex/spline_catrom.svg)
…the function returns the height of the orange dot at position t.

(This image is for illustration only, it does not show the correct
curve for the shown points.)

`"linearsolve"` (or `"solvelinear"`)

Maps between a set of non-uniform positions and a set of values.
The [kspline](kspline.html "Returns an interpolated value along a curve defined by a basis and key/position pairs.") function does this mapping implicitly.

```vex
tk = spline("linearsolve", t, k0, k1, k2, k3, ...);
v = spline(basis, tk, v1, v2, v3, ...);

```

(Technically, `linearsolve` interprets the values as key values, solves the
intersection of the spline, and returns the intercept point.)

`"monotonecubic"`

A cubic spline with no overshoot of the control values.

`"bezier"`

A bezier spline.

`"bspline"`

A b-spline basis.

`"hermite"`

A hermite spline.

`sample_pos`

The position along the curve at which to sample the value.

Returns

The value at sample_pos along a polyline or cubic spline.

Note
For b-spline basis, this function implicitly assumes the multiplicity of 3
for b-spline curve end point, even though the given control points and
knots are not explicitly repeated. This ensures the curve passes through the
end control points, making it easier to create continuous ramp curves with
mixed interpolation bases (e.g., b-spline basis segments surrounded by
linear interpolation segments).
