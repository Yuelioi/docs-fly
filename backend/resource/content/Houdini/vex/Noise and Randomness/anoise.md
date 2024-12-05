---
title: anoise
order: 1
---
`float  anoise(vector pos)`

`vector  anoise(vector pos)`

`float  anoise(vector pos, int turbulence, float rough, float atten)`

`vector  anoise(vector pos, int turbulence, float rough, float atten)`

`float  anoise(vector pos, int periodX, int periodY, int periodZ)`

`vector  anoise(vector pos, int periodX, int periodY, int periodZ)`

`float  anoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

`vector  anoise(vector pos, int periodX, int periodY, int periodZ, int turbulence, float rough, float atten)`

These functions generate “alligator” noise, a type of cellular noise similar to
Worley noise ([wnoise](wnoise.html "Generates Worley (cellular) noise.")). It is currently not possible to simulate alligator noise
using the Worley functions, but it’s possible to get a very similar “look”.

The bounds on the noise are roughly (0, 1). This function only supports 3D noise.
