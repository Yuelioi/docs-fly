---
title: sample_direction_cone
order: 13
---
`vector  sample_direction_cone(vector center, float maxangle, vector2 u)`

Show/hide arguments

`center`

Direction in the center of the cone. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the cone
will be, so long as all `u` values are between 0 and 1.

`u`

Pair of numbers between 0 and 1.

Returns a unit vector, i.e. a vector of length 1, based on `u`.
Given uniform random `u` pairs of values in `[0,1)`, the returned unit vectors will be
uniform random and continuous with respect to `u` on the surface of the unit sphere,
in the area within `maxangle` of the direction indicated by `center`.
