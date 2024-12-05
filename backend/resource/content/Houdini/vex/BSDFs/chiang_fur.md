---
title: chiang_fur
order: 5
---
| Since | 20.0 |
| --- | --- |

`bsdf  chiang_fur(vector nn, vector tanV, float mask, float cuticle, float R_v, float R_s, float TT_v, float TT_s, float TRT_v, float TRT_s, float shift, vector absorption_coeff, float ior, float R2_v, float R2_s, vector R2_color, ...)`

Creates a BSDF for computation of the physically based fur and hair
model described in papers “A Practical and Controllable Hair
and Fur Model for Production Path Tracing” by Chiang, “Physically-AccurAate
Fur Reflectance: Modeling, Measurement and Rendering” by Ling-Qi Yan et al.
Chiang Fur is an extension of the [Chiang model - shader](chiang.html "Returns a chiang BSDF.").
This model considers a structural feature of fur and thick hair:
the so-called medulla. Hair and fur have three main components:

Suitable only for curve geometry.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

Show/hide arguments

`nn`

bumped/shading normal

`tanV`

tangent vector along V

`mask`

Masking main lobes in favour of medulla

`cuticle`

Modulates fresnel factor of the hair’s outermost layer

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

`R2_v`

Longitudinal roughness value “v” of extra lobe R2

`R2_s`

Azimuthal roughness value “s” of extra lobe R2

`R2_color`

Extra R2 lobes has a color input to colorize the reflection. This could be used to promote iridescence colors for example.
