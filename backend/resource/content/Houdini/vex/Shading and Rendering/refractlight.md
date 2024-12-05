---
title: refractlight
order: 64
---
| On this page | * [Light inclusion/exclusion options](#light-inclusion-exclusion-options) * [Area sampling options](#area-sampling-options) * [Ray options](#ray-options) * [Examples](#examples) |
| --- | --- |
| Context(s) | [shading](../contexts/shading.html) |

`void  refractlight(vector &cf, vector &of, float &af, vector P, vector D, float bias, float max_contrib, ...)`

`void  refractlight(vector &cf, vector &of, float &af, vector P, vector N, vector I, float eta, float bias, float max_contrib, ...)`

Computes the illumination of surfaces refracted by the current surface.
Computes and outputs the output color (cf), opacity (of) and
alpha (af). See [opacity vs. alpha](../contexts/shading_contexts.html#opacity) .

bias is typically a small number (for example 0.005) used to help
eliminate self-reflection.

max_contrib tells the renderer how much the reflected light will
contribute to the final color of the pixel. This has no effect on the
resultant color.

The first form of the refractlight() function takes a position and
direction, typically computed by the [refract](refract.html "Returns the refraction ray given an incoming direction, the
normalized normal and an index of refraction.") or
[fresnel](fresnel.html "Computes the fresnel reflection/refraction contributions given an
incoming vector, surface normal (both normalized), and an index of
refraction (eta).") functions.

To prevent the renderer from computing standard transparency (i.e.
non-refracted transparency), the Of variable must be set to {1,1,1} to
make the surface “opaque”. The Af variable can be set to any arbitrary
value.
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

Area sampling options

## area-sampling-options

For area sampling, you must specify both the angle and sample variadic parameters. For example:

```vex
surface
blurry_mirror(float angle = 3; int samples = 16; float bias=0.05)
{
    Cf = reflectlight(bias, 1, "angle", angle, "samples", samples);
}

```

Ray options

## ray-options

Tip
When you specify a texture, such as with the `"environment"` keyword,
you can also use the image filtering keyword arguments. See [environment](environment.html "Returns the color of the environment texture.")
for a listing of the image filter keyword arguments.

Show/hide arguments

"`scope`",
`string`

A list of objects which can be hit by the rays. When specified, `scope` overrides the default scope that would have been selected for the given `raystyle`. The `"scope:default"` value will cause the `scope` argument to use the default scope for the current context - as if the argument were not specified.

Allows an override of the [scope](../contexts/shading_contexts.html#scope) for ray-intersections.
A special scope argument, `scope:self`, will match the currently
shading object.

"`currentobject`",
`material`

Used to specify what the current shading object is. For example, when used with the scope argument, `scope:self` will match the object passed in by this argument.

"`maxdist`",
`float`
`=-1`

The maximum distance to search for objects. This can be used to limit the search of objects to nearby objects only. If the `maxdist` given is negative, then it will act as if there is no maximum distance.

Allows an override of the maximum distance the ray can
travel when testing for intersections. Some functions (such as
[fastshadow](fastshadow.html "Sends a ray from the position P along the direction specified by the
direction D.")) have the maximum distance implicitly defined (by
the length of the ray) and should probably avoid using this
option. However, this option can be used effectively when
computing reflections, global illumination, refraction etc.

"`variancevar`",
`string`

The name of a VEX export variable to use for variance anti-aliasing. The renderer compares the value with adjacent micropolygons in micropolygon rendering to decide what shading points need additional samples (using `vm_variance` [property](../../props/index.html "Properties let you set up flexible and powerful hierarchies of rendering, shading, lighting, and camera parameters.") as a threshold). If more samples are required, the algorithm takes samples up to the specified maximum ray samples.

This variable must be imported from the hit surface, so it must be in the list of imported names (see “importing information back from the ray” below). If the named variable is not imported, this option will be ignored.

Variance antialiasing puts more samples in areas of the image with high variance, for example a sharp shadow edge. It is only used when `vm_dorayvariance` is enabled. Otherwise, only the min ray samples (or an explicitly supplied `"samples"` value) are used for antialiasing of the gather loop.

Overrides the global variance control (mantra’s -v option)
which is used to determine anti-aliasing quality of ray tracing.
For more information please refer to the documentation on
mantra.

"`angle`",
`float`
`=0`

The distribution angle (specified in radians). For gather(), rays will be distributed over this angle. For trace(), the angle is used to indicate the rate at which the filter width should increase with increasing intersection distance. Larger angles will cause farther hit surfaces to use larger derivatives, leading to improved texturing and displacement performance.

To be effective, the samples parameter should also be specified.

"`samples`",
`int|float`
`=1`

How many samples should be sent out to filter rays. For the
irradiance and occlusion functions, specifying a samples
parameter will override the default irradiance sampling.

"`environment`",
`string`

If the ray sent out to the scene misses everything, then
it’s possible to specify an environment map to evaluate.

Using the ray’s direction, the environment map specified
will be evaluated and the resulting color will be returned.
Most likely, it will be necessary to specify a transform
space for the environment map evaluations.

In the case of refractlight and trace the Of and Af
variables will be set to 0 regardless of the background
color specified. the resulting color.

When an environment map is specified, the filtering options
from texture() are also supported.

See [how to create an environment/reflection map](../../render/envmaps.html).

"`envobject`",
`string`

If an environment map is used, the orientation of the
environment map can be specified by transforming the ray
into the space of another object, light or fog object in the
scene. In Houdini, null objects can be used to specify the
orientation. For example:

```vex
Cf = R*reflectlight(bias, max(R), "environment", "map.rat", "envobject", "null_object_name");

```

"`envlight`",
`string`

If an environment map is used, the orientation of the
environment map can be specified by transforming the ray
into the space of a light in the scene.

"`envtint`",
`vector`

If an environment map is used, tint it with this color.

"`background`",
`vector`

If a ray misses all objects, use this as the
background color of the scene. In the case of refractlight and
trace the Of and Af variables will be set to 0 regardless of the
background color specified.

"`distribution`",
`string`

**Functions**: [irradiance](irradiance.html "Computes irradiance (global illumination) at the point P with the normal N."), [occlusion](occlusion.html "Computes ambient occlusion.")

Distribution for computing irradiance. The default is to use
a cosine distribution (diffuse illumination). The possible
values for the style are `"nonweighted"` for uniform sampling
or `"cosine"` for cosine weighted sampling.

Examples

## examples

```vex
surface
glass(float eta=1.3, bias = 0.005)
{
    float    Kr, Kt;
    vector    R, T;
    vector    cf, of;
    float    af;
    frensel(normalize(I), normalize(N), eta, Kr, Kt, R, T);
    Cf  = Kr * reflectlight(P, R, bias, Kr);
    refractlight(cf, of, af, P, T, bias, Kt);
    Cf += Kt * cf;
    Af = clamp(Kr + af*Kt, 0, 1);
    Of = 1;
}

```
