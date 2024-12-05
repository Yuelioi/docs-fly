---
title: vnoise
order: 38
---
`void  vnoise(float position, float jitter, int &seed, float &f1, float &f2, float &pos1, float &pos2)`

Generates 1D noise.

`void  vnoise(float position, float jitter, int &seed, float &f1, float &f2, float &pos1, float &pos2, int period)`

Generates periodic 1D noise.

`void  vnoise(float posx, float posy, float jittx, float jitty, int &seed, float &f1, float &f2, float &pos1x, float &pos1y, float &pos2x, float &pos2y)`

Generates 2D noise. This is similar to the other forms but uses pairs of floats instead of a vector.

`void  vnoise(float posx, float posy, float jittx, float jitty, int &seed, float &f1, float &f2, float &pos1x, float &pos1y, float &pos2x, float &pos2, int periodx, int periody)`

Generates periodic 2D noise.

`void  vnoise(vector position, vector jitter, int &seed, float &f1, float &f2, vector &pos1, vector &pos2)`

Generates 3D noise.

`void  vnoise(vector position, vector jitter, int &seed, float &f1, float &f2, vector &pos1, vector &pos2, int periodx, int periody, int periodz)`

`void  vnoise(vector position, vector jitter, int &seed, float &f1, float &f2, vector &pos1, vector &pos2, vector period)`

Generates periodic 3D noise.

`void  vnoise(vector4 position, vector4 jitter, int &seed, float &f1, float &f2, vector4 &pos1, vector4 &pos2)`

Generates 4D noise.

`void  vnoise(vector4 position, vector4 jitter, int &seed, float &f1, float &f2, vector4 &pos1, vector4 &pos2, int periodx, int periody, int periodz, int periodw)`

`void  vnoise(vector4 position, vector4 jitter, int &seed, float &f1, float &f2, vector4 &pos1, vector4 &pos2, vector4 period)`

Generates periodic 4D noise.

Show/hide arguments

`position`

The position at which to sample the noise.

`jitter`

The amount of randomness to add to the noise in each axis.

`seed`

Outputs an integer value associated with the nearest seed point. This seed is pretty much guaranteed to be unique for every point (meaning that it’s unlikely that two points close by have the same seed associated with them).

`pos1`, `pos2`

These variables are overwritten with the positions of the two nearest seed points, in order of closeness.

`f1`, `f2`

These variables are overwritten with the distances to the nearest seed points, in order of closeness.

You can combine these distances to generate noise patterns. The noise generated tends to be very “cellular” in nature. In fact, one of the nice things is that you can determine “cell” boundaries by using the expression: `if (f2 - f1)` which will be true if the point in space is crossing the boundary between two cells.

`period`, `periodx`, `periody`, `periodz`, `periodw`

If you include the period argument(s), the function generates repeating (periodic) noise.

Voronoi noise gives almost identical results to the Worley noise function ([wnoise](wnoise.html "Generates Worley (cellular) noise.")). However, this function has controls over jittering (i.e. how randomly the points are scattered through space) and also return the actual locations of the two nearest seed points, whereas [wnoise](wnoise.html "Generates Worley (cellular) noise.") only returns the distances to the two nearest seed points.

Though this function is slightly more expensive than [wnoise](wnoise.html "Generates Worley (cellular) noise."), since it returns the actual point positions, you can overcome some of the artifacts of Worley noise. For example, to get even boundaries along the cell boundaries:

```vex
if (f2 - f1 < tolerance * (distance(p1, p2) / (f1 + f2)) ...

```

This will “normalize” the boundary width based on the distance between the two random points in space.

There are also periodic forms of vnoise().
Examples

## examples

```vex
// 1D noise
float    fp0, fp1, p1x, p1y, p2x, p2y;
vector    vp0, vp1;
vnoise(s*10, 0.8, seed, f1, f2, fp0, fp1);
vnoise(s*10, t*10, 0.8, 0.8, seed, f1, f2, p1x, p1y, p2x, p2y);
vnoise(P*10, {.8, .8, .8}, seed, f1, f2, vp0, vp1);

```
