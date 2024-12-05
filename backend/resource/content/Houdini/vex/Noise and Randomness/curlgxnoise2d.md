---
title: curlgxnoise2d
order: 3
---
`vector  curlgxnoise2d(vector2 xy)`

`vector  curlgxnoise2d(float x, float y)`

`vector  curlgxnoise2d(vector xyz)`

`vector  curlgxnoise2d(vector4 xyzt)`

Evaluates a divergence free vector field by computing curl of one-dimensional
simplex noise (see [gxnoise](gxnoise.html "Evaluates a simplex noise field.")). The first two components of the
output vector contain the divergence free noise, and its final component is 0.
