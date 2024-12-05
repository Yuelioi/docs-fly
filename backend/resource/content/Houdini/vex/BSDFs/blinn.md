---
title: blinn
order: 3
---
`bsdf  blinn(float exponent, ...)`

`bsdf  blinn(vector nml, float exponent, ...)`

Returns a Blinn BSDF.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

`vector  blinn(vector nml, vector V, float roughness, ...)`

Computes Blinn shading.

Show/hide arguments

`nml`

The normal of the surface to use for evaluate.

`V`

The incidence vector.

`exponent`

The exponent value. The higher the value, the tighter the specular lobe.

Each of the functions can also take an optional light mask.

![](../../images/rendering/blinn.png)
