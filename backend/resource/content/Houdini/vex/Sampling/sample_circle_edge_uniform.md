---
title: sample_circle_edge_uniform
order: 9
---
`vector2  sample_circle_edge_uniform(float u)`

Show/hide arguments

`u`

Number between 0 and 1.

Returns a unit vector2, i.e. a vector2 of length 1, based on `u`.
Given uniform random `u` values in `[0,1)`, the returned unit vectors will be
uniform random and continuous with respect to `u` on the edge of the unit circle.
Specifically, it returns `(cos(angle),sin(angle))`, where `angle` is `2*pi*u`.
