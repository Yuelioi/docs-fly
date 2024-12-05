---
title: refract
order: 63
---
`vector  refract(vector direction, vector normal, float index)`

Returns the refraction ray given an incoming direction, the
normalized normal and an index of refraction.

The index is a relative index of refraction, the ratio between
the interior and exterior index of refraction, where the exterior
is defined by the direction of the normals (normals point away from
the interior).

In the case of total internal reflection, this function returns the
reflection vector.

For example:

```vex
refract(normalize(I), normalize(N), outside_to_inside_ior)

```
