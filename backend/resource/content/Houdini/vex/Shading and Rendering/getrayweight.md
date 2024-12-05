---
title: getrayweight
order: 30
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`float  getrayweight()`

Returns an approximation to the contribution of the ray to the final
pixel color. Often, this is a better metric for judging contribution to
the final pixel color than [getraylevel](getraylevel.html "Returns the depth of the ray tree for the current shading."). However, this
relies on prior shaders giving good estimates on the contribution to
their shading (see [reflectlight](reflectlight.html "Computes the amount of reflected light which hits the surface.")).
