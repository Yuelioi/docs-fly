---
title: blinnBRDF
order: 2
---
`float  blinnBRDF(vector L, vector N, vector V, float rough)`

[specularBRDF](specularBRDF.html "Returns the computed BRDFs for the different lighting models used in VEX shading."), [phongBRDF](phongBRDF.html), `blinnBRDF`,
and [diffuseBRDF](diffuseBRDF.html) return the computed BRDF for the
different lighting models used in VEX shading. You can use them in
custom [illuminance](illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") loops to replicate the lighting models of the
corresponding VEX lighting functions.

See [specularBRDF](specularBRDF.html "Returns the computed BRDFs for the different lighting models used in VEX shading.") for some example code.
