---
title: xnoise
order: 40
---
`float  xnoise(float x)`

`vector  xnoise(float x)`

`float  xnoise(float x, float y)`

`vector  xnoise(float x, float y)`

`float  xnoise(vector xyz)`

`vector  xnoise(vector xyz)`

`float  xnoise(vector4 xyzt)`

`vector  xnoise(vector4 xyzt)`

Simplex noise is very close to Perlin noise, except with the samples on a
simplex mesh rather than a grid. This results in less grid artifacts. It also
uses a higher order `bspline` to provide better derivatives.

The various functions return the noise value at a 4D (vector4 argument),
3D (vector argument), 2D (two float arguments) or 1D (float argument)
position. You can get a random float value or a vector of three random
values.

The noise is in the range 0-1 with a median of 0.5. The distribution of
the noise depends on the dimension, with higher dimensions approaching a
Gaussian distribution of noise values.
