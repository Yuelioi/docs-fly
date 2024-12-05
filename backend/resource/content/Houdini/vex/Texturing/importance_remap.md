---
title: importance_remap
order: 6
---
| Since | 18.5 |
| --- | --- |

`vector2  importance_remap(string map, vector2 uv, ...)`

This function remaps texture coordinates to new texture coordinates based on importance sampling of the texture.

Show/hide arguments

`map`

The filename of the texture map to use to guide resampling.

`uv`

Components should be in the range `0` to `1`. The function remaps these coordinates so more “important” areas of the input text (that is, brighter areas) get more samples.

"`maxres`",
`int`
`=0`

While building lookup tables for importance sampling, the function resamples the texture for faster evaluation. This argument clamps the resolution of the resampled map. Depending on the use, importance sample tables can often be significantly smaller resolution than the source image with no perceptible loss.

A value of `0` (the default) just uses the original texture size.

Note
It’s a good idea to limit the lookup table size, since you typically don’t need much resolution for importance sampling, and a large texture can generate a huge LUT.

Returns

The remapped texture coordinates.
