---
title: getglobalraylevel
order: 14
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  getglobalraylevel()`

Returns the depth of the ray tree for computing global illumination. If
this function returns a non-zero value, the shader is being called for
the purpose of evaluating global illumination.
