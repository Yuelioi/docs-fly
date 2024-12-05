---
title: sample_lognormal
order: 22
---
`float  sample_lognormal(float mu, float sigma, float u)`

`float  sample_lognormal(float mu, float sigma, float minvalue, float maxvalue, float u)`

Show/hide arguments

`mu`

The mean of the underlying normal distribution.

`sigma`

The standard deviation of the underlying normal distribution.

`u`

A number in the range `[0,1)`.

`minvalue`,`maxvalue`

When given, instead of sampling the full log-normal distribution,
the distribution with its range limited to `[minvalue,maxvalue]` will be
sampled.

Samples the log-normal distribution with the specified `mu` and `sigma`, optionally
with a `minvalue` and `maxvalue`. To use parameters that are more understandable,
`median` and `stddev`, please use `sample_lognormal_by_median`.
Given uniform random `u` values in `[0,1)`, this will return log-normally
distributed random numbers. The return value will be monotone increasing
with respect to `u`.

The log-normal distribution is sampled by sampling a normal distribution
and exponentiating the result, giving a value that is always positive, so
this distribution is often used for generating random point scales.
