---
title: limport
order: 51
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

Note
This function is only valid inside an [illuminance](illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") loop.

`int  limport(string name, <type>&value)`

Show/hide arguments

`name`

The name of the shader variable to read.

`&value`

If the named variable is defined and exported, the function overwrites this variable with the variable’s value.

Returns

Returns `1` if the shader variable is defined and exported, or `0` otherwise.
