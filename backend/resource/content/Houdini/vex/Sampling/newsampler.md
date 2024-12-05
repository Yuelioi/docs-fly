---
title: newsampler
order: 4
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  newsampler(...)`

`int  newsampler(int seed, ...)`

Returns an initialized sampler sequence for use as the first argument
to the [nextsample](nextsample.html) function.

Show/hide arguments

`seed`

You can specify a seed value for the sequence.
Using the same seed will generate the same sequence.
This can be useful when random sampling into a point cloud, for consistent results.
