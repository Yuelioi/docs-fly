---
title: sample_sphere_shell_uniform
order: 29
---
| Since | 17.0 |
| --- | --- |

`vector  sample_sphere_shell_uniform(vector u, float alpha)`

Show/hide arguments

`u`

Three numbers between 0 and 1.

`alpha`

The inner radius to be bounded by. A number between 0 and 1.

Returns a vector of length \< 1, based on `u`.
Given uniform random `u` vectors of three values in `[0,1)`, and a number in `[0,1]`, the returned vectors will be
uniform random and continuous with respect to `u` inside the unit sphere shell with the inner raidus alpha.
