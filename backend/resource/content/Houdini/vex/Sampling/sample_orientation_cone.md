---
title: sample_orientation_cone
order: 25
---
`vector4  sample_orientation_cone(vector4 center, float maxangle, vector u)`

Show/hide arguments

`center`

Orientation in the center of the cone. This does not need to be normalized.

`maxangle`

Maximum angle, in radians, away from `center` that any sample of the cone
will be, so long as all `u` values are between 0 and 1.

Note
This angle is the maximum quaternion rotation angle between the orientation quaternion that the output represents and `center`, which is twice the Euclidean cone angle of the unit 4D hypersphere being sampled. A `maxangle` of π would sample all orientations, but only half of all 4D unit vectors; a `maxangle` of 2π would sample all 4D unit vectors.

`u`

Three numbers between 0 and 1.

Returns a unit vector4, i.e. a vector4 of length 1, based on `u`.
Given uniform random `u` vectors of three values in `[0,1)`, the returned quaternion
orientations will be uniform random and continuous with respect to `u`,
in the region of orientations within `maxangle` of `center`.
