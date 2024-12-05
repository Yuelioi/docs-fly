---
title: shimport
order: 72
---
| Context(s) | [fog](../contexts/fog.html)  [surface](../contexts/surface.html) |
| --- | --- |
This function is only valid inside an [illuminance](illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") loop.

`int  shimport(string variable_name, <type>&value)`

Show/hide arguments

`variable_name`

The variable to import from the shadow shader.

`value`

If the variable is successfully read, the value is copied into this variable.

Returns

`1` if the variable is defined and exported, `0` otherwise.
