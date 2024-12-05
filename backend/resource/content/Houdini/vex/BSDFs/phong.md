---
title: phong
order: 17
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

![](../../images/rendering/phong.png)

`bsdf  phong(float exponent, ...)`

`bsdf  phong(vector nml, float exponent, ...)`

See [writing a PBR shader](../pbr.html) for information on BSDFs.

`vector  phong(vector nml, vector V, float shinyness, ...)`

V represents the normalized vector from the surface to the eye
(-normalize(I)). shinyness is the Phong exponent (typically around
20 or higher). roughness represents the surface roughness (typically
0 to 1).

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
