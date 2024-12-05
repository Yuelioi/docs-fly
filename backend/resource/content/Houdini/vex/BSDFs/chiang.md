---
title: chiang
order: 4
---
| Since | 19.0 |
| --- | --- |

`bsdf  chiang(vector nn, vector tanV, float R_v, float R_s, float TT_v, float TT_s, float TRT_v, float TRT_s, float shift, vector absorption_coeff, float ior, ...)`

Creates a BSDF for computation of the physically based hair
model described in paper “A Practical and Controllable Hair
and Fur Model for Production Path Tracing” by Chiang et al.

Suitable only for curve geometry.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

Show/hide arguments

`nn`

bumped/shading normal

`tanV`

tangent vector along V

`R_v`

Longitudinal roughness value “v” for lobe R (section 4.1 of the paper)

`R_s`

Azimuthal roughness value “s” for lobe R (section 4.1 of the paper)

`TT_v`

Longitudinal roughness value “v” for lobe TT (section 4.1 of the paper)

`TT_s`

Azimuthal roughness value “s” for lobe TT (section 4.1 of the paper)

`TRT_v`

Longitudinal roughness value “v” for lobe TRT (section 4.1 of the paper)

`TRT_s`

Azimuthal roughness value “s” for lobe TRT (section 4.1 of the paper)

`shift`

Represents the cuticle angle, which affects the position of the specular highlight. Input range of -1 to 1 is internally mapped to -90 to 90 (eg meaning 3-degrees would be 3/90 = 0.03333)

`absorption_coeff`

The absorption coefficient (section 4.2 of the paper)

`ior`

Index of refraction (eg 1.55)
