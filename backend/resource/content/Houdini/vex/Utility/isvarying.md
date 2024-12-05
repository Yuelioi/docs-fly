---
title: isvarying
order: 7
---
`int  isvarying(<type>x)`

`int  isvarying(<type>x[])`

Returns 1 when the given variable is varying, or 0 when it is uniform. A
variable is varying when it may have a different value for each processor
in the VEX SIMD array. If a value is varying, shader execution will
generally be slower - so this function can be useful when debugging shader
performance. Any variable type can be passed to the `isvarying()` function.
