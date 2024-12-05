---
title: sample_sphere_cone
order: 28
---
`vector  sample_sphere_cone(vector center, float maxangle, vector u)`

Show/hide arguments

`center`

Direction in the center of the cone. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the cone
will be, so long as all `u` values are between 0 and 1.

`u`

Three numbers between 0 and 1.

Returns a vector of length \< 1, based on `u`.
Given uniform random `u` vectors of three values in `[0,1)`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit sphere,
in the volume within `maxangle` of the direction indicated by `center`.
