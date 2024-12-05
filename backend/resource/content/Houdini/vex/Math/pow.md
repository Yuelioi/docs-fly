---
title: pow
order: 53
---
`float  pow(float n, float exponent)`

`<vector> pow(<vector>v, float exponent)`

Raises `n` to the power of `exponent`. For vectors, this is done per-component.

When `n` is less than zero, the exponent will be rounded to the closest integer.
