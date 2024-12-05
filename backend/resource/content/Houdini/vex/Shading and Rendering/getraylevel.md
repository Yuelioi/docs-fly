---
title: getraylevel
order: 29
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`int  getraylevel()`

Returns the depth of the ray tree for the current shading. If the
returned value is 0, this represents a ray from the camera to the scene.
If the ray level is 1, the ray represents either a reflection/refraction
ray. If the level is 2, then this represents a reflection/refraction
which appears in a previous reflection/refraction etc.
