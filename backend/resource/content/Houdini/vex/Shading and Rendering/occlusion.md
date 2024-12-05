---
title: occlusion
order: 56
---
| On this page | * [Sample adaptation options](#sample-adaptation-options) * [Ray options](#ray-options) |
| --- | --- |
| Context(s) | [shading](../contexts/shading.html) |

`vector  occlusion(vector P, vector N, ...)`

Computes ambient occlusion at the point P with the normal N. Just as
in the [irradiance](irradiance.html "Computes irradiance (global illumination) at the point P with the normal N.") function, the hemisphere is
sampled. However, unlike irradiance, surfaces intersected during the
hemisphere sampling are not shaded. For this function to work
properly, either a constant background color or an environment map
must be specified in the
[optional scope parameters](../contexts/shading_contexts.html#scope).

`void  occlusion(float &coverage, vector &missed_direction, vector P, vector N, ...)`

Instead of computing color information from ambient occlusion, this
form computes the coverage (the percentage of occlusion) and the
average direction of empty space. The average direction can be used to
look up the color in a pre-blurred environment map.

Sample adaptation options

## sample-adaptation-options

Show/hide arguments

"`adaptive`",

`1` or `0`. Turns on an automatic optimization that will reduce the
number of samples when there is little variation in occlusion
above the sample point. This can improve performance at the
expense of some possible flickering or additional noise.

Show/hide arguments

"`bias`",
`float`
`=0.001`

Ray-tracing bias. Gives control over self-intersection.

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
