---
title: invlerp
order: 9
---
| Since | 18.5 |
| --- | --- |

`float  invlerp(float a, float min, float max)`

`<vector> invlerp(<vector>a, <vector>min, <vector>max)`

Returns the amount to mix `min` and `max` to generate the
input value `a`. This is the inverse of the `lerp` function.

The vector version operates component-wise, so the resulting
vector will be the independent mixing amount for each dimension.

If `a` is outside the range `min` to `max`, values greater than
`1` or less than `0` will be produced.

If `min` and `max` are equal, the mixing value is `0.5`.
