---
title: rotate_x_to
order: 22
---
`vector2  rotate_x_to(vector2 direction, vector2 v)`

`vector  rotate_x_to(vector direction, vector v)`

`vector4  rotate_x_to(vector4 direction, vector4 v)`

Show/hide arguments

`direction`

Direction to which a vector along the positive x-axis, e.g. (1,0,0),
would be rotated. This does not need to be normalized.

`v`

Vector to which to apply the rotation.

Applies to `v`, the rotation that would move (1,0), (1,0,0), or (1,0,0,0) to
`direction` most directly.

In the case of `direction` being (-1,0,0), there are multiple distinct
rotations that would move (1,0,0) to (-1,0,0) with a
half turn rotation, so one is chosen arbitrarily, negating `v.x` and `v.z`.
In 2D, there is only one distinct rotation that moves (1,0) to (-1,0),
equivalent to negating `v`. In 4D, the rotation where `v` is negated
is also chosen.

This is used by functions like `sample_direction_cone` and `sample_sphere_cone`
to rotate the cone centre from (1,0,0) to a given direction vector.
