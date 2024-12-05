---
title: noised
order: 24
---
`void  noised(float x, float &v, float &dvdx)`

`void  noised(float x, vector &v, vector &dvdx)`

`void  noised(float x, float y, float &v, float &dvdx, float &dvdy)`

`void  noised(float x, float y, vector &v, vector &dvdx, vector &dvdy)`

`void  noised(vector xyz, float &v, float &dvdx, float &dvdy, float &dvdz)`

`void  noised(vector xyz, vector &v, vector &dvdx, vector &dvdy, vector &dvdz)`

`void  noised(vector4 xyzw, float &v, float &dvdx, float &dvdy, float &dvdz, float &dvdw)`

`void  noised(vector4 xyzw, vector &v, vector &dvdx, vector &dvdy, vector &dvdz, vector &dvdw)`

This computes both the perlin noise value, and the derivatives of the
noise along each axis. This can be performed quite efficiently as
there are analytic derivatives available.

See [noise and randomness](../random.html) in the VEX language
guide for more information.
