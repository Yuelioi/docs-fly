---
title: shadow
order: 70
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [surface](../contexts/surface.html) |
| --- | --- |
This function can only be called from within an [illuminance](illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") statement.

`void  shadow(vector &Cl)`

`vector  shadow(vector Cl)`

Uses the `P` and `L` global variables.

`void  shadow(vector &Cl, vector P, vector L)`

`vector  shadow(vector Cl, vector P, vector L)`
