---
title: gxnoise
order: 11
---
`float|vector|vector2 gxnoise(vector2 xy)`

`float|vector|vector2 gxnoise(float x, float y)`

`float|vector|vector2 gxnoise(vector xyz)`

`float|vector|vector2 gxnoise(vector4 xyzt)`

Simplex noise is similar to Perlin noise, except with samples on a simplex mesh
rather than a grid. This family of simplex noise functions uses a different
lattice structure and a cheaper accumulation method compared to
[xnoise](xnoise.html "Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid.  This results in less grid artifacts.  It also uses a higher order bspline to provide better derivatives.").

The various functions return noise value at a 4D (`vector4` argument), 3D
(`vector` argument), or 2D (a single `vector2` argument or two `float` inputs)
position. You can also get a random floating point value or a vector of two or
three entries.

Noise values will be in the 0-1 range. Nature of the noise field depends on the
number of input dimensions. Higher-dimensional noise uses tighter noise
elements, and the resultant noise field appears more structured and less smooth.
Consider using the slower [xnoise](xnoise.html "Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid.  This results in less grid artifacts.  It also uses a higher order bspline to provide better derivatives.") function in higher dimensions if
this function gives you undesirable results.
