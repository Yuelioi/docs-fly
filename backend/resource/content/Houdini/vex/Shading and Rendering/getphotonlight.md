---
title: getphotonlight
order: 26
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  getphotonlight()`

Returns `-1` if the shader is not generating photons from a light.

The return value is an integer identifying a light. You can get convert this
integer ID to the light name with [getlightname](getlightname.html "Returns the name of the current light when called from within an illuminance loop, or converts an integer light ID into the light’s name.")
