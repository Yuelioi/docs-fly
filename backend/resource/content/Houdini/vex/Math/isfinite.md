---
title: isfinite
order: 33
---
`int  isfinite(float x)`

Returns 1 if the given value is a normal, finite, number.

Returns 0 if it is a NAN or infinite. Note that VEX generally
will not produce these: 3/0 is defined as 0, for example.
