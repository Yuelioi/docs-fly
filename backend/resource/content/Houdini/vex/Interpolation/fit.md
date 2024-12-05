---
title: fit
order: 5
---
`float  fit(float value, float omin, float omax, float nmin, float nmax)`

`<vector> fit(<vector>value, <vector>omin, <vector>omax, <vector>nmin, <vector>nmax)`

Takes the value in the range (omin, omax) and shifts it to the corresponding value in the new range (nmin, nmax).

The function clamps the given value the range (omin, omax) before fitting, so the resulting value will be guaranteed to be in the range (nmin, nmax). To avoid clamping use [efit](efit.html "Takes the value in one range and shifts it to the corresponding value in a new range.") instead.

Examples

## examples

```vex
fit(.3, 0, 1, 10, 20) == 13

```
