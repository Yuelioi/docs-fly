---
title: sample_circle_ring_uniform
order: 10
---
| Since | 17.0 |
| --- | --- |

`vector2  sample_circle_ring_uniform(vector2 u, float alpha)`

Show/hide arguments

`u`

Pair of numbers between 0 and 1.

`alpha`

The inner radius to be bounded by. A number between 0 and 1.

Returns a vector2 of length \< 1, based on `u`.
Given uniform random `u` pairs of values in `[0,1)`,
and value `alpha` in `[0,1]`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit circle ring with inner radius of `alpha`.
Specifically, it returns `scale*(cos(angle),sin(angle))`, where `angle` is `2*pi*u.x`
and `scale` is `sqrt((1-alpha^2)*u.y+alpha^2)`.
