---
title: pxnoised
order: 28
---
| Since | 17.0 |
| --- | --- |

`void  pxnoised(float x, int px, float &v, float &dvdx)`

`void  pxnoised(float x, int px, vector &v, vector &dvdx)`

`void  pxnoised(float x, float y, int px, int py, float &v, float &dvdx, float &dvdy)`

`void  pxnoised(float x, float y, int px, int py, vector &v, vector &dvdx, vector &dvdy)`

`void  pxnoised(vector xyz, int px, int py, int pz, float &v, float &dvdx, float &dvdy, float &dvdz)`

`void  pxnoised(vector xyz, int px, int py, int pz, vector &v, vector &dvdx, vector &dvdy, vector &dvdz)`

`void  pxnoised(vector4 xyzw, int px, int py, int pz, int pw, float &v, float &dvdx, float &dvdy, float &dvdz, float &dvdw)`

`void  pxnoised(vector4 xyzw, int px, int py, int pz, int pw, vector &v, vector &dvdx, vector &dvdy, vector &dvdz, vector &dvdw)`

This computes both the simplex noise value, and the derivatives of the
noise along each axis. This can be performed quite efficiently as
there are analytic derivatives available.

See [xnoise](xnoise.html "Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid.  This results in less grid artifacts.  It also uses a higher order bspline to provide better derivatives.") and [pxnoise](pxnoise.html "Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid.  This results in less grid artifacts.  It also uses a higher order bspline to provide better derivatives. This is the periodic simplex noise") for other versions of xnoise.

See [noise and randomness](../random.html) in the VEX language
guide for more information.
