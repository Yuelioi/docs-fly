---
title: instance
order: 5
---
`matrix  instance(vector P, vector N)`

`matrix  instance(vector P, vector N, vector scale)`

`matrix  instance(vector P, vector N, vector scale, vector pivot)`

`matrix  instance(vector P, vector N, vector scale, vector4 rotate, vector up)`

`matrix  instance(vector P, vector N, vector scale, vector4 rotate, vector up, vector pivot)`

`matrix  instance(vector P, vector N, vector scale, vector4 rotate, vector4 orient)`

`matrix  instance(vector P, vector N, vector scale, vector4 rotate, vector4 orient, vector pivot)`

Creates a transform matrix from the given arguments, using the same
method that the [Copy SOP](../../nodes/sop/copy.html) uses to transform its
output instances. The instance is placed at point `P`, oriented
along the normal direction `N`, and, optionally, scaled by `scale`. An
optional `pivot` parameter can be supplied as the local transformation point
for the instance.

The function supports two methods for setting rotation. The first method
requires an explicit `up` vector, which should be tangent to `N`. This
`up` vector, along with `N`, is used to construct an orthonormal frame
in which the rotation takes place.
The second method uses an explicit orientation relative to the XYZ axis
to construct the frame.
