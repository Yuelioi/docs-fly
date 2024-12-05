---
title: mx_worley
order: 22
---
`float|vector|vector2 mx_worley(vector2 pos, float jitter, int metric)`

`float|vector|vector2 mx_worley(vector pos, float jitter, int metric)`

`float|vector|vector2 mx_worley(vector2 pos, float jitter, int periodx, int periody, int periodz)`

`float|vector|vector2 mx_worley(vector pos, float jitter, int periodx, int periody, int periodz, int periodw)`

Returns a Worley Noise value that matches the values as in the standard MaterialX library.

Jitter should normally be clamped between 0 and 1.

The metric is an integer representing how the distance is measured for Worley noise

- 0 - Euclidean Distance
- 1 - Distance Squared
- 2 - Manhattan Distance
- 3 - Chebyshev Distance
