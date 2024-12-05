---
title: noise
order: 23
---
`float  noise(float pos)`

`vector  noise(float pos)`

Sample one or three numbers at the given position from 1D noise.

`float  noise(float posx, float posy)`

`vector  noise(float posx, float posy)`

Sample one or three numbers at the given position from 2D noise.

`float  noise(vector pos)`

`vector  noise(vector pos)`

Sample one or three numbers at the given position from 3D noise.

`float  noise(vector4 pos)`

`vector  noise(vector4 pos)`

Sample one or three numbers at the given position from 4D noise.

There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space.

Note
This function generates non-periodic noise. Use the
[pnoise](pnoise.html "There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space.") function to generate periodic Perlin noise.

The various functions return the noise value at a 4D (vector4 argument),
3D (vector argument), 2D (two float arguments) or 1D (float argument)
position. You can get a random float value or a vector of three random
values.

The noise is in the range 0-1 with a median of 0.5. The distribution of
the noise depends on the dimension, with higher dimensions approaching a
Gaussian distribution of noise values.
