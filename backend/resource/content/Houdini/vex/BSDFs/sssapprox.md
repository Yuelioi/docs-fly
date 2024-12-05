---
title: sssapprox
order: 23
---
`bsdf  sssapprox(vector albedo, float meanFreePath, float roughness, float scale, ...)`

Show/hide arguments

`albedo`

Average surface reflectance.

`meanFreePath`

Average distance between scattering events.

`roughness`

A value from '0.0' to '1.0' blending to an ideal diffuse transmission reflectance profile.

`scale`

Physical scale of the material. Smaller scale values will make the material more transmissive.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

Models SSS illumination based on an approximate reflectance profile.
