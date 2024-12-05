---
title: ndcdepth
order: 8
---
| Since | 18.0 |
| --- | --- |

`float  ndcdepth(float z)`

When Karma converts a position to NDC space, the z depth value is

transformed based on the render camera projection and clipping planes. This
function converts the NDC z-depth into the camera space. That is, the distance
down the z-axis in the space of the camera.

function returns the argument passed in.
Examples

## examples

```vex
vector ndc = ptransform("space:ndc", P);
float pz_camera = ndcspace(ndc.z);

// This value can also be computed using
float pz_camera = -ptransform("space:camera", P).z;

```
