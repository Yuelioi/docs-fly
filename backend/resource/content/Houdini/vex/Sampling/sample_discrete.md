---
title: sample_discrete
order: 15
---
`int  sample_discrete(int nvalues, float u)`

`int  sample_discrete(float weights[], float u)`

Show/hide arguments

`nvalues`

The returned integer will be uniform in the range `[0,nvalues-1]`,
returning 0 if `u==0`, and returning `nvalues-1` if `u==1`. The
output will be clamped to that range in case `u` is out of the range
`[0,1)`, to reduce the risk of roundoff on `u` causing problems.

`weights`

Relative weights, (the sum does not need to be 1), of each integer
value in the range `[0,len(weights)-1]`.

`u`

A number between 0 and 1.

Returns an integer, based on `u`, either uniformly weighted from 0 to
`nvalues-1`, or weighted based on the `weights` array from 0 to
`len(weights)-1`.
Given uniform random `u` values in `[0,1)`, the version taking `nvalues`
will return uniform random integers in `[0,nvalues-1]`, and the version
taking `weights` will return random integers in `[0,len(weights)-1]`, where
the probability of `i` is `weights[i]/sum_of_weights`.
