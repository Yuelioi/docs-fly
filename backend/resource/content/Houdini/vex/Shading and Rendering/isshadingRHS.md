---
title: isshadingRHS
order: 45
---
| Context(s) | [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html)  [displace](../contexts/displace.html) |
| --- | --- |

`int  isshadingRHS()`

Shading can occur in either a left or right handed vector space. This function
can be used to check the native orientation in shading space. By default,
Mantra is left-handed space while Karma is in right-handed space.
