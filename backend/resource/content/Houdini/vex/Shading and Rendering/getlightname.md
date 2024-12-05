---
title: getlightname
order: 18
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`string  getlightname()`

Returns the name of the current light when called from within an
[illuminance](illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") loop or when a current light has been set using
[setcurrentlight](setcurrentlight.html "Sets the current light").

`string  getlightname(int lightid)`

Returns the name of the light referred to by the given integer
light ID. Integer light IDs are used by some low-level VEX functions
instead of strings for efficiency.
