---
title: curlnoise2d
order: 5
---
`vector  curlnoise2d(float x, float y)`

`vector  curlnoise2d(vector xyt)`

Computes a divergence free vector field based on the cross product of the derivatives of two Perlin noise functions.

The resulting vectors all lie along the X-Y plane.

Note
This is not the same as projecting a `curlnoise` to a plane.

See [noise and randomness](../random.html) in the VEX language
guide for more information.
