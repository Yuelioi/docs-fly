---
title: isuvrendering
order: 47
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  isuvrendering()`

Returns 1 if the shader is being called while evaluating uv rendering (texture unwrapping), and 0 for normal rendering.

Use this function to evaluate shaders differently when baking illumination.
