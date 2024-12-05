---
title: pnoise
order: 27
---
`float|vector pnoise(float x, int px)`

`float|vector pnoise(vector x, vector p)`

`float|vector pnoise(vector4 xyzt, vector4 p)`

`float|vector pnoise(float x, float y, int px, int py)`

`float|vector pnoise(vector xyz, int px, int py, int pz)`

`float|vector pnoise(vector4 xyzt, int px, int py, int pz, int pt)`

There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space.

This function generates periodic noise. Use the [noise](noise.html "There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space.")
function to generate non-periodic Perlin noise.

The various functions return the noise value at a 4D (vector4 argument),
3D (vector argument), 2D (two float arguments) or 1D (float argument)
position. You can get a random float value or a vector of three random
values.

The “p” int or vector arguments specify the range of periodicity. For
example, if you're making a 2D image and you want it to tile:

```vex
clr = pnoise(X * 4, Y * 5, _4, 5_)

```

In this example, X is in the range 0-4, Y is in the range 0-5, and the
noise is periodic within that segment.

If a period argument is 0, VEX treats that as *no* periodicity. Use
this to make the noise periodic in one dimension but not another.

The distribution of the noise depends on the dimension, with higher
dimensions approaching a Gaussian distribution of noise values.
