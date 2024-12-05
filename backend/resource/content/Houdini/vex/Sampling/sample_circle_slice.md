---
title: sample_circle_slice
order: 11
---
`vector2  sample_circle_slice(vector2 center, float maxangle, vector2 u)`

Show/hide arguments

`center`

Direction in the center of the slice. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the slice
will be, so long as all `u` values are between 0 and 1.

`u`

Pair of numbers between 0 and 1.

Returns a vector2 of length \< 1, based on `u`.
Given uniform random `u` pairs of values in `[0,1)`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit circle,
in the slice within `maxangle` of the direction indicated by `center`.
