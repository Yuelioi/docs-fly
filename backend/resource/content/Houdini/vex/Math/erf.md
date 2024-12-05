---
title: erf
order: 24
---
`float  erf(float v)`

`vector2  erf(vector2 v)`

The [Gauss error function](http://en.wikipedia.org/wiki/Error_function). Houdini uses the Boost
library’s implementation internally.

The erf(vector2) version computes the complex error function and returns
the complex result. This function is much slower than the erf(float)
function for real values.
