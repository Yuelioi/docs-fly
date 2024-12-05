---
title: makebasis
order: 41
---
`void  makebasis(vector &xaxis, vector &yaxis, vector zaxis)`

`void  makebasis(vector &xaxis, vector &yaxis, vector zaxis, vector u)`

Completes an orthonormal basis for the given `zaxis` vector consisting of
the `xaxis` and `yaxis` basis vectors. When only the `zaxis` vector is
given, the basis will have an arbitrary orientation. When a second vector
`u` is provided, the `yaxis` vector will be constrained so that it is aligned
with that vector.
