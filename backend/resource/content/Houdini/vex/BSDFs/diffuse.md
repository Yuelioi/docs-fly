---
title: diffuse
order: 8
---
![](../../images/rendering/diffuse.png)

`bsdf  diffuse(...)`

`bsdf  diffuse(float roughness, ...)`

`bsdf  diffuse(vector nml, ...)`

`bsdf  diffuse(vector nml, float roughness, ...)`

`bsdf  diffuse(vector nml, vector geo_normal, ...)`

`bsdf  diffuse(vector nml, vector geo_normal, float roughness, ...)`

Diffuse reflections. This BSDF has an albedo of 0.5. If your shader or
geometry has smooth normals (N and Ng differ) you should avoid the
signature `diffuse(vector nml)` since it assumes that the
shading normal matches the geometric normal.

See [writing a PBR shader](../pbr.html) for information on BSDFs.

`vector  diffuse(vector nml, ...)`

`vector  diffuse(vector nml, vector V, float roughness, ...)`

This form uses the
Oren-Nayar lighting model to compute the diffuse illumination for the
surface. The Oren-Nayar lighting model is a more sophisticated lighting
model than Lambertian lighting. The V vector represents a vector from
the surface to the eye (i.e. -normalize(I)). With a roughness of 0, the
Oren-Nayar lighting model is equivalent to the Lambertian model. As
roughness increases toward 1, the illumination changes to mimic rougher
materials (like clay). The Oren-Nayar form of diffuse() is more
expensive than Lambertian diffuse lighting.

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
