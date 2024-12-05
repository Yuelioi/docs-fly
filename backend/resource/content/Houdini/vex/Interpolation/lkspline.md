---
title: lkspline
order: 11
---
`float  lkspline(float sample_pos, float value1, float key_pos1, ...)`

Samples a polyline defined by a series of value/position pairs.
This is useful for specifying a 1D data ramp.

`vector  lkspline(float sample_pos, vector value1, float key_pos1, ...)`

`vector4  lkspline(float sample_pos, vector4 value1, float key_pos1, ...)`

Samples a polyline defined by a series of vector value/position pairs.
This is useful for specifying a color ramp.

If you just want linearly spaced keys, use [lspline](lspline.html "Samples a polyline defined by linearly spaced values.") instead.

Show/hide arguments

`sample_pos`

The position along the curve at which to sample.

`valuen`, `key_posn`

To define the shape of the curve, you pass a number of value/position pairs specifying the key points through which the curve passes.

You must specify key positions in ascending order or the results will be unpredictable.

Returns

The value of the curve at the sampled position.

Tip
The [spline](spline.html "Samples a value along a polyline or spline curve.") function is a more flexible superset of this function.
