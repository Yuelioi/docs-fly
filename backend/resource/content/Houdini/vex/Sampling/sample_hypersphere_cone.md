---
title: sample_hypersphere_cone
order: 19
---
`vector4  sample_hypersphere_cone(vector4 center, float maxangle, vector4 u)`

Show/hide arguments

`center`

Direction in the center of the cone. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the cone
will be, so long as all `u` values are between 0 and 1.

`u`

Four numbers between 0 and 1.

Returns a vector4 of length \< 1, based on `u`.
Given uniform random `u` vectors of four values in `[0,1)`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit hypersphere,
in the hypervolume within `maxangle` of the direction indicated by `center`.
