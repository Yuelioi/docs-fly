---
title: random
order: 30
---
`float  random(float|intposition)`

`vector  random(float|intposition)`

`vector2  random(int position)`

`vector4  random(float|intposition)`

Generate 1D, 2D, 3D, or 4D noise based on a 1D position.

`float  random(float|intxpos, float|intypos)`

`vector  random(float|intxpos, float|intypos)`

`vector4  random(float|intxpos, float|intypos)`

Specify a 2D position in the noise field using two numbers.

`float  random(vector position)`

`vector  random(vector position)`

`vector4  random(vector position)`

Specify a 3D position in the noise field using a vector.

`float  random(vector4 position)`

`vector  random(vector4 position)`

`vector4  random(vector4 position)`

Specify a 4D position in the noise field using a vector4.

Generate a random number based on the integer position in `N` dimensional space
(where N is 1 to 4 dimensions). Unlike the noise functions, the random
functions do not smoothly interpolate the random values between integer
lattice points. The `random()` functions are very efficient ways of doing
something like `noise(floor(position))`.

While `random()` takes floats, it only varies the random effect for
integer changes. To have a random result that varies with even the
smallest float changes, use `rand()`.

The result of this is in the half-open interval `[0, 1)`.
