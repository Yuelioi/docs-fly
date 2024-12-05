---
title: dimport
order: 1
---
| Context(s) | [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  dimport(string name, <type>&out)`

Reads a variable from the displacement shader for the surface.

Mantra runs the shaders for a surface in a fixed order:

1. Displacement
1. Surface (possibly calling light shaders in `illuminance` loops)
1. Fog (possibly calling light shaders in `illuminance` loops)

Once the displacement shader has run, you can use `dimport`
to retrieve exported variables from it. Once the surface shader
has run, you can use [simport](simport.html "Imports a variable sent by a surface shader in an illuminance loop.") to retrieve exported variables
from it.

If the shader variable named by the first argument is defined and
exported, the function returns 1 and puts the value in the second
argument. Otherwise, it returns 0.
