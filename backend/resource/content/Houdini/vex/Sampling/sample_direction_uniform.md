---
title: sample_direction_uniform
order: 14
---
`vector  sample_direction_uniform(vector2 u)`

Show/hide arguments

`u`

Pair of numbers between 0 and 1.

Returns a unit vector, i.e. a vector of length 1, based on `u`.
Given uniform random `u` pairs of values in `[0,1)`, the returned unit vectors will be
uniform random and continuous with respect to `u` on the surface of the unit sphere.
