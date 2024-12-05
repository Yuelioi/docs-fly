---
title: xnoised
order: 41
---
`void  xnoised(float x, float &v, float &dvdx)`

`void  xnoised(float x, vector &v, vector &dvdx)`

`void  xnoised(float x, float y, float &v, float &dvdx, float &dvdy)`

`void  xnoised(float x, float y, vector &v, vector &dvdx, vector &dvdy)`

`void  xnoised(vector xyz, float &v, float &dvdx, float &dvdy, float &dvdz)`

`void  xnoised(vector xyz, vector &v, vector &dvdx, vector &dvdy, vector &dvdz)`

`void  xnoised(vector4 xyzw, float &v, float &dvdx, float &dvdy, float &dvdz, float &dvdw)`

`void  xnoised(vector4 xyzw, vector &v, vector &dvdx, vector &dvdy, vector &dvdz, vector &dvdw)`

This computes both the simplex noise value, and the derivatives of the
noise along each axis. This can be performed quite efficiently as
there are analytic derivatives available.

See [noise and randomness](../random.html) in the VEX language
guide for more information.
