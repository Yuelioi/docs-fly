---
title: sample_circle_arc
order: 8
---
`vector2  sample_circle_arc(vector2 center, float maxangle, float u)`

Show/hide arguments

`center`

Direction in the center of the arc. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the arc
will be, so long as `u` is between 0 and 1.

`u`

Number between 0 and 1.

Returns a unit vector2, i.e. a vector2 of length 1, based on `u`.
Given uniform random `u` values in `[0,1)`, the returned unit vectors will be
uniform random and continuous with respect to `u` on the edge of the unit circle,
in the arc within `maxangle` of the direction indicated by `center`.
