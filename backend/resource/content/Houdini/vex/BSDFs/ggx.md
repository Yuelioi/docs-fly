---
title: ggx
order: 11
---
| Since | 18.0 |
| --- | --- |

`bsdf  ggx(vector ng, vector nn, vector xg, vector yg, vector F0, vector F90, float alphax, float alphay, int masking, int fresblend, float eta, float reflect, float refract, int reflectmask, int refractmask, float dispersion, ...)`

Creates a BSDF for computation of the GGX microfacet model, used
for rough specular reflection and refraction.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

Show/hide arguments

`ng`

Normalized geometry normal

`nn`

Normalized bumped/shading normal

`xg`

Normalized x tangent vector

`yg`

Normalized y tangent vector

`F0`

Color tint at oblique angles

`F90`

Color tint at grazing angles

`alphax`

Roughness along the x tangent vector

`alphay`

Roughness along the y tangent vector (use the same value as alphax for isotropic)

`masking`

Enable/Disable microfacet masking

`fresblend`

0 = disable fresnel, 1 = enable fresnel, 2 = use critical angle to select between reflection and refraction.

`eta`

Index of refraction

`reflect`

Explicit scalar on reflection (0->1). Or -1 to let the function decide itself on the appriate value based on geometric information.

`refract`

Explicit scalar on refraction (0->1). Or -1 to let the function decide itself on the appriate value based on geometric information.

`reflectmask`

Bitmask representing the desired reflection behaviour. Simply passing in `bouncemask(reflectlabel)` will suffice

`refractmask`

Bitmask representing the desired refraction behaviour. Simply passing in `bouncemask(refractlabel)` will suffice

`dispersion`

Amount of dispersion

Light inclusion/exclusion options

## light-inclusion-exclusion-options

Show/hide arguments

"`categories`",
`string`
`="*"`

Specifies lights to include/exclude by their “category” tags.
This is the preferred include/exclude lights rather than pattern matching
light names with the `"lightmask"` keyword argument.

For example:

```vex
diff = diffuse(nml, "lightmask", "hero | fill");

```

See [light categories](../../render/lights.html#categories) for more information.

"`lightmask`",
`string`
`="*"`

When evaluating light and shadow shaders, objects have pre-defined light
masks. This mask is usually specified in the geometry object and
specifies a list of lights which are used to illuminate a surface or fog
shader. It is possible to override the default light mask by specifying
a “lightmask” argument.

For example:

```vex
diff = diffuse(nml, "lightmask", "light*,^light2");

```

…will cause all lights whose names begin with “light” except for a
light named “light2” to be considered for diffuse illumination.

All Houdini scoping patterns, excepting group expansion, are supported:

- `*` - wild-card match
- `?` - single character match
- `^` - exclusion operator
- `[list]` - character list match
