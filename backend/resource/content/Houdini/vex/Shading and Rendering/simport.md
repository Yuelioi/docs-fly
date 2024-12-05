---
title: simport
order: 73
---
| Context(s) | [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html) |
| --- | --- |

`int  simport(string name, <type>&value)`

Imports a variable from the surface shader.

Mantra runs the shaders for a surface in a fixed order:

1. Displacement
1. Surface (possibly calling light shaders in `illuminance` loops)
1. Fog (possibly calling light shaders in `illuminance` loops)

Once the displacement shader has run, you can use [dimport](dimport.html "Reads a variable from the displacement shader for the surface.")
to retrieve exported variables from it. Once the surface shader
has run, you can use `simport` to retrieve exported variables
from it.

If the shader variable named by the first argument is defined and
exported, the function returns 1 and puts the value in the second
argument. Otherwise, it returns 0.
