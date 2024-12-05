---
title: sample_circle_uniform
order: 12
---
`vector2  sample_circle_uniform(vector2 u)`

Show/hide arguments

`u`

Pair of numbers between 0 and 1.

Returns a vector2 of length \< 1, based on `u`.
Given uniform random `u` pairs of values in `[0,1)`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit circle.
Specifically, it returns `scale*(cos(angle),sin(angle))`, where `angle` is `2*pi*u.x`
and `scale` is `sqrt(u.y)`.
