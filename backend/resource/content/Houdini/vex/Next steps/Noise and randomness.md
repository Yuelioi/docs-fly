---
title: Noise and randomness
order: 4
---
| On this page | * [Noise generators](#noise-generators) * [Random number generators](#random-number-generators) |
| --- | --- |

Noise generators

## noise-generators

You can generate noise with the [noise](functions/noise.html "There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space."),
[wnoise](functions/wnoise.html "Generates Worley (cellular) noise."), [vnoise](functions/vnoise.html "Generates Voronoi (cellular) noise."), [onoise](functions/onoise.html "These functions are similar to wnoise and vnoise."),
[snoise](functions/snoise.html "These functions are similar to wnoise."), and [anoise](functions/anoise.html "Generates alligator noise.") functions.

Each represents a different algorithm for generating noise. This
allows you to make speed tradeoffs between different methods, as
well as ensure comparability with noise values generated through
other or older methods.

The relative costs for computing noise of different types is
roughly:

| Noise type | Relative cost |
| --- | --- |
| Perlin noise ([noise](functions/noise.html "There are two forms of Perlin-style noise: a non-periodic noise which changes randomly throughout N-dimensional space, and a periodic form which repeats over a given range of space.")) | 1.0 |
| Original perlin noise ([onoise](functions/onoise.html "These functions are similar to wnoise and vnoise.")) | 1.1 |
| Worley noise ([wnoise](functions/wnoise.html "Generates Worley (cellular) noise.")) | 1.8 |
| Sparse Convolution noise ([snoise](functions/snoise.html "These functions are similar to wnoise.")) | 2.1 |
| Alligator noise ([anoise](functions/anoise.html "Generates alligator noise.")) | 2.3 |

Random number generators

## random-number-generators

[random](functions/random.html "Generate a random number based on the integer position in 1-4D space.") generates a random number based on the position
in N dimensional space (where N is 1 to 4 dimensions). Unlike the
noise functions which smoothly interpolate the random values between
integer lattice points, the random functions do not. The random()
functions are very efficient ways of doing something like:
`noise(floor(position))`

[hscript_rand](functions/hscript_rand.html "Produces the exact same results as the Houdini expression function of
the same name.") produces the exact same results
as the Houdini rand() expression. This function will generate
different random values for every floating point seed. This is
different from [random](functions/random.html "Generate a random number based on the integer position in 1-4D space.") which converts the floating point
argument to an integer seed. [hscript_rand](functions/hscript_rand.html "Produces the exact same results as the Houdini expression function of
the same name.") may
produce different results on different hardware or operating
systems.

[nrandom](functions/nrandom.html "Non-deterministic random number generation function.") is a non-deterministic random number
generator. Numbers generated will be between 0 and 1. These
functions will generate the same sequence of random numbers if
called in precisely the same order. However, there is no seed
involved so it is not possible to reproduce the same random number
or sequence multiple times.
