---
title: sample_normal
order: 24
---
`float  sample_normal(float u)`

`float  sample_normal(float mean, float stddev, float u)`

`float  sample_normal(float origmean, float origstddev, float minvalue, float maxvalue, float u)`

`vector2  sample_normal(vector2 u)`

`vector  sample_normal(vector u)`

`vector4  sample_normal(vector4 u)`

Show/hide arguments

`u`

A number, or multiple numbers, in the range `[0,1)`.

`mean`

The mean of the distribution, or 0 if not specified.

`origmean`

The mean the distribution would have, were it not for `minvalue`
and `maxvalue`, limiting the range.

`stddev`

The standard deviation (scale) of the distribution, or 1 if not specified.

`origstddev`

The standard deviation (scale) the distribution would have, were it
not for `minvalue` and `maxvalue`, limiting the range.

`minvalue`,`maxvalue`

When given, instead of sampling the full normal distribution,
the distribution with its range limited to `[minvalue,maxvalue]` will be
sampled.

Samples the normal distribution with the specified `mean` and `stddev`, optionally
with a `minvalue` and `maxvalue`.
Given uniform random `u` values in `[0,1)`, this will return normally
distributed random numbers. The return value will be monotone increasing
with respect to `u`.

The `vector2`, `vector`, and `vector4` versions
return multiple samples with mean 0 and standard deviation 1. The
distribution of these vectors is naturally isotropic, i.e. rotating
the distribution won’t change it, which can be useful in simulations.
To add a maximum distance from the origin, while keeping the distribution
isotropic, use:

`sample_normal(0,1,0,maxdist,u.x) * sample_direction_uniform(set(u.y,u.z))`
