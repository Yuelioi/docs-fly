---
title: isfogray
order: 42
---
| Context(s) | [light](../contexts/light.html)  [shadow](../contexts/shadow.html) |
| --- | --- |

`int  isfogray()`

Returns 1 if the shader is being called to evaluate illumination for fog
objects, or 0 if the light or shadow shader is being called to evaluate
surface illumination.

Use this function to simplify light shaders when evaluating for fog.
