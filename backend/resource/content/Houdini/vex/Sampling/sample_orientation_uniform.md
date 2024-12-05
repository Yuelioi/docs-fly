---
title: sample_orientation_uniform
order: 26
---
`vector4  sample_orientation_uniform(vector u)`

Show/hide arguments

`u`

Three numbers between 0 and 1.

Returns a unit vector4, i.e. a vector4 of length 1, based on `u`.
Given uniform random `u` vectors of three values in `[0,1)`, the returned unit vectors will be
uniform random and continuous with respect to `u` on the surface of the unit hypersphere.
In other words, they will be uniform random orientation quaternions.
