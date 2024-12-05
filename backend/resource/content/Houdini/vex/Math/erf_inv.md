---
title: erf_inv
order: 25
---
`float  erf_inv(float v)`

The inverse of the [Gauss error function](http://en.wikipedia.org/wiki/Error_function).

`erf_inv(erf(v)) = v = erf(erf_inv(v))`

To generate a normally-distributed random number, `n`, with mean `mu` and standard deviation `sigma`,
from a uniformly-distributed random number, `u`, between 0 and 1,

`n = mu + sqrt(2)*sigma*erf_inv(2*u - 1)`
