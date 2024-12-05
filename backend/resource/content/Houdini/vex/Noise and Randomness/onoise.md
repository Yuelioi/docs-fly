---
title: onoise
order: 26
---
`float  onoise(vector pos)`

`vector  onoise(vector pos)`

`float  onoise(vector pos, int turbulence, float rough, float atten)`

`vector  onoise(vector pos, int turbulence, float rough, float atten)`

`float  onoise(vector pos, int periodX, int periodY, int periodZ)`

`vector  onoise(vector pos, int periodX, int periodY, int periodZ)`

`float  onoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

`vector  onoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

These functions are similar to [wnoise](wnoise.html "Generates Worley (cellular) noise.") and
[vnoise](vnoise.html "Generates Voronoi (cellular) noise."). However, they are marginally less efficient in
computation and don’t have the same characteristics. The bounds on the
noise are roughly (-1, 1). Only 3D noise is supported. However, this
noise has the ability to compute turbulence with roughness and
attenuation on the noise.
