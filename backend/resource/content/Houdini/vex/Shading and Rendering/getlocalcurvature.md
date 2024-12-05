---
title: getlocalcurvature
order: 21
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`vector  getlocalcurvature(float s, float t)`

Returns 0 vector if the object does not have subdivision enabled, or has no displacement shader assigned.
Otherwise, the measured convexity and concavity will be returned in `x` and `y` components respectively.

Show/hide arguments

`s`

Parametric S shading value. This should be passed from the `s` global variable.

`t`

Parametric <type> shading value. This should be passed from the `t` global variable.
