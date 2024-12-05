---
title: snoise
order: 37
---
`float  snoise(vector pos)`

`vector  snoise(vector pos)`

`float  snoise(vector pos, int turbulence, float rough, float atten)`

`vector  snoise(vector pos, int turbulence, float rough, float atten)`

`float  snoise(vector pos, int periodX, int periodY, int periodZ)`

`vector  snoise(vector pos, int periodX, int periodY, int periodZ)`

`float  snoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

`vector  snoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

These functions are similar to [wnoise](wnoise.html "Generates Worley (cellular) noise."). The noise returned
is based on the weights of all of the closest points, with each point’s
contribution based on a meta-ball like rolloff curve. That is, if the
sample point is close to the sphere, its contribution will be greater.

The bounds on the noise are roughly (-1.7, 1.7). Only 3D noise is
supported. However, this noise has the ability to compute turbulence
with roughness and attenuation on the noise.
