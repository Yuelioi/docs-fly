---
title: sample_cauchy
order: 6
---
`float  sample_cauchy(float u)`

`float  sample_cauchy(float scale, float u)`

`vector2  sample_cauchy(float scale, vector2 u)`

`float  sample_cauchy(float origscale, float minvalue, float maxvalue, float u)`

`<vector> sample_cauchy(<vector>u)`

Sample multivariate Cauchy distributions with median 0 and scale 1. The
distribution of these vectors is forced to be isotropic, i.e. rotating
the distribution won’t change it, which can be useful in simulations.
This wouldn’t be the case if one generated components of the vectors as
independent samples of the univariate Cauchy distribution.

Show/hide arguments

`u`

A number, or multiple numbers, in the range `[0,1)`.

`scale`

The scale of the distribution, or 1 if not specified.
This is the difference between the 50th percentile and the 75th percentile.

`origscale`

The scale the distribution would have, were it not for `minvalue`
and `maxvalue`, limiting the range.

`minvalue`,`maxvalue`

When given, instead of sampling the full Cauchy distribution,
the distribution with its range limited to `[minvalue,maxvalue]` will be
sampled.

Returns

Monotonically increasing value with respect to `u`.

Samples the Cauchy distribution with median zero and the specified `scale`,
optionally with a `minvalue` and `maxvalue`.
Given uniform random `u` values in `[0,1)`, this will return Cauchy
distributed random numbers.

Note that without limits, the Cauchy distribution has
no defined mean or variance, which can cause statistical problems if not
dealt with carefully.

To add a maximum distance from the origin, while keeping the distribution
isotropic, use:

```vex
!vex
sample_cauchy(1,0,maxdist,u.x) * sample_direction_uniform(set(u.y,u.z))
```

The 2D Cauchy distribution is the distribution of photons hitting a plane,
coming from a point light that is distance `scale` from the plane.
