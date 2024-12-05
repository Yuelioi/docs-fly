---
title: sample_hypersphere_uniform
order: 20
---
`vector4  sample_hypersphere_uniform(vector4 u)`

Show/hide arguments

`u`

Four numbers between 0 and 1.

Returns a vector4 of length \< 1, based on `u`.
Given uniform random `u` vectors of four values in `[0,1)`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit hypersphere.
