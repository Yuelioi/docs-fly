---
title: gather
order: 9
---
| On this page | * [Light inclusion/exclusion options](#light-inclusion-exclusion-options) * [Ray options](#ray-options) * [Ray sending options](#ray-sending-options) * [Sending information to the surface’s shader](#sending-information-to-the-surface-s-shader) * [Importing information from the ray](#importing-information-from-the-ray) * [Sample filtering options](#sample-filtering-options) * [Pipeline options](#pipeline-options) * [Examples](#examples) |
| --- | --- |
You can use VEX’s `gather` loop statement to send rays into the scene and gather the resulting information.

```vex
gather(vector origin, vector direction, ...)
{
    // Statements for rays that hit other surfaces
} else {
    // Statements for rays which didn't hit any surface
}

```

After the origin and direction arguments, any additional arguments are interpreted as keyword/value pairs.

```vex
gather(origin, dir, "samples", 100) {
    ...
}

```

Note
If you specify 0 samples, it will still take at least 1 sample.

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

Ray sending options

## ray-sending-options

Show/hide arguments

"`width`",
`float`
`=-1`

Specifies the filter width at the source of the ray. If `angle` is also specified, the filter width will become larger with increasing distance along the ray. By default, the filter width will be initialized from the current shading context, so it’s normally not necessary to specify `width` directly. Negative values are ignored and will also cause the filter width to be initialized from the current shading context.

"`distribution`",
`string`
`="cosine"`

Determines the sampling distribution.

For [gather](gather.html "Sends rays into the scene and returns information from the shaders of
surfaces hit by the rays."):

- `cosine` – Rays are distributed by the cosine (diffuse) function over the hemisphere.
- `uniform` – Rays are distributed uniformly over the hemisphere

For [sample_geometry](sample_geometry.html "Samples geometry in the scene and returns information from the shaders of surfaces that were sampled."):

- `area` – Samples are distributed by primitive area
- `parametric` – Samples are distributed by primitive ID, subdivision ID, and parametric surface coordinates (s, t).
- `solidangle` – Samples are distributed either by primitive area or by primitive area and solid angle subtended by the primitive.

"`biasdir`",
`vector`
`=Ng`

Overrides the bias direction when **Bias Along Normal** is enabled. When no `biasdir` is specified, the geometric normal `Ng` is used. When bias along normal is disabled, this option has no effect.

"`SID`",
`int`
`=0`

Sample identifier to be passed to the called shader. If the calling shader has used SID to generate samples, it can be useful to pass the modified sample identifier to the called shader so that it can begin sampling at the specified offset. This value will be used to initialize the SID global in the hit surface.

"`rayweight`",
`float`
`=1`

A hint to mantra to indicate the relative contribution of this ray to the final shading. This value is used by the ray clip threshold to limit sending of rays (similar to ray bounce).

"`raystyle`",
`string`
`="refract"`

The type of rays you are sending. Mantra will use `raystyle` to determine both the default raytracing mask and bounce limit used for ray termination.

- `reflect` – Sending reflection rays. Mantra will use the reflection mask and reflection limit to terminate raytracing.
- `refract` – (default) Sending refraction rays. Mantra will use the refraction mask and refraction limit to terminate raytracing.
- `diffuse` – Sending diffuse rays. Mantra will use the diffuse limit for diffuse rays.
- `shadow` – Sending shadow rays. Mantra will not modify the raytracing level and will trace against `shadowmask` if inside a shadow or light shader.
- `primary` – Sending primary rays. This style can be used when a shader needs to change the direction of a primary ray without affecting the behavior of render settings that apply only to directly visible objects (such as matte and phantom). Mantra will still increment the raytracing level when sending `primary` rays.
- `nolimit` – Sending reflection rays with no limit on the number of raytracing bounces. Mantra will still increment the raytracing level when sending `nolimit` rays.

"`categories`",
`string`

A category expression used to select the objects which can be hit by rays. When specified, this overrides the existing `reflectcategories` and `refractcategories` parameters.

For example, `^hidden` will hit all objects which do not have the hidden category, and `shiny|happy` will hit all objects with either the shiny or happy category.

The intersection of the scope and categories parameters are used to choose the objects which can be hit by rays.

"`samplebase`",
`float`
`=0`

Typically, rays are distributed over the surface of the micro-polygon being shaded. This argument can be used to control the area. A value of 0 will force all rays to be sent from the same point. A value of 1 will cover the entire micro-polygon. (Gather only)

"`transparentsamples`",
`int`
`=1`

The number of transparent samples to take for stochastic transparency with array outputs. Normally this value should be set to 1 unless you have requested exports in array variables - in which case the ray tracer will insert an entry in the array for each sample along the ray.

Note
`transparentsamples` must be 1 when importing `F` or `ray:material` using `screendoor` `samplefilter`.

Sending information to the surface’s shader

## sending-information-to-the-surface-s-shader

Using a keyword in the form `"send:name", value`, you can pass data from the originating surface to surfaces which are intersected by the ray. These arguments pass any values you want.

```vex
gather(P, dir, "send:N", normalize(N)) { ... }

```

You can extract this passed data on the receiving end (that is, in the surface being hit by the ray) with the [rayimport](rayimport.html "Imports a value sent by a shader in a gather loop.") function. The first argument is the name (without the `send:` prefix) and the second argument is a variable in which to store the imported value.

`int rayimport(string name, <type> &value)`

`rayimport` returns `1` if the value was imported successfully.

Importing information from the ray

## importing-information-from-the-ray

You can specify names of global or exported variables to import from the hit shader in the form `"varname", &var`, typically including `Cf` (color vector of surface hit) and `Of` (opacity vector of surface hit).

```vex
vector  hitcf;
gather(P, dir, "bias", 0.01, "Cf", hitcf) {...}

```

In addition, you can import the following special keywords to get information about the ray itself:

Show/hide arguments

"`ray:origin`",
`&vector`

The origin of the ray (defined in `else` clause also).

"`ray:direction`",
`&vector`

The direction of the ray (defined in `else` clause also).

"`ray:length`",
`&float`

The distance to the first surface which was hit by the ray.

"`ray:area`",
`&float`

The total surface area of all the geometry in the raytracing scope.

"`ray:solidangle`",
`&float`

The estimated solid angle subtended by all geometry in the raytracing scope. For large objects close to or enclosing the ray origin, this may be a very poor estimate while for individual primitives the estimate can be very good.

You can retrieve information about more than one hit along
the ray by requesting data in an array variable. When an imported
value is of an array type, the [trace](trace.html "Sends a ray from P along the normalized vector D.") function will automatically append
an entry in the array for each individual hit point that was composited
during ray tracing. For the `opacity` sample filter (see below), an entry will be
created in the array for each semi-transparent sample encountered until
full opacity is reached. When using array outputs, it may also be useful
to use the `all` sample filter, which will cause all hits along the ray to
be inserted regardless of whether the opacity limit was exceeded.

```vex
// Find the position and normal for all hit points along the ray,
// regardless of visibility.
vector a_pos[];
vector a_nml[];
trace(P, dir, Time,
        "samplefilter", "all",
            "P", a_pos,
            "N", a_nml);

```

Sample filtering options

## sample-filtering-options

By default, Houdini composites the global variables using opacity blending. In some cases, it’s more useful to get the value from the closest surface (regardless of whether it’s transparent). You can use the special `samplefilter` keyword with a string value of either `closest` or `opacity` to control whether the value of a global is from the closest surface or opacity blended.

Show/hide arguments

"`samplefilter`",
`string`

When the `samplefilter` keyword is encountered in the argument list, *all following* import variables will use the specified filtering mode. You can specify multiple `samplefilter` arguments in a single gather statement to filter different variables in different ways.

The current types of allowed for `samplefilter` are

`minimum`

Take the minimum value of all the samples. Note that with tuple values, the minimum value of each component will be used.

`maximum`

Take the maximum value of all the samples. Note that with tuple values, the maximum value of each component will be used.

`opacity`

Composite samples using the over operation.

`closest`

This is the default behavior, returning only the closest surface.

`screendoor`

Use stochastic compositing of the samples.

`sum`

Return the sum of the values for all samples.

`sum_square`

Return the sum of the squares of the values of all samples.

`sum_reciprocal`

Return the sum of the reciprocals of each sample.

Note
When using [sample_geometry](sample_geometry.html "Samples geometry in the scene and returns information from the shaders of surfaces that were sampled."), the default `samplefilter` is set to `closest` by default, since opacity blending only works when compositing data along a ray.

```vex
gather(P, dir,
        "samplefilter", "opacity",
            "Cf", hitCf,
            "Of", hitOf,
        "samplefilter", "closest",
            "P", hitP,
            "N", hitN)
{
    trace(pos, dir, time,
            // Composite the bsdf of the hit surfaces using stochastic transparency
            "samplefilter", "screendoor",
            "F", hitF,
            // But find the closest sample's position
            "samplefilter", "closest",
            "P", hitP);
}

```

Pipeline options

## pipeline-options

Show/hide arguments

"`pipeline`",
`string`

As you specify variables, you can intersperse `pipeline` keyword options to control where in the pipeline to fill out read/write variables. The value can be one of `surface`, `atmosphere`, or `displacement`. You can specify the `pipeline` option multiple times. Each use of the option affects any variables specified after it (up to the next usage of `pipeline` if any).

```vex
gather(p, d, "pipeline", "surface", "Cf", surfCf,
             "pipeline", "atmosphere" "Cf", fogCf, "P", hitP)

```

Examples

## examples

Here are two shaders which communicate through `gather`.

This shader sends rays into the scene. It sends `1.234` in a message labeled `exportvalue`. Since the surface hit by the rays may not have an exported parameter called `amp`, it’s important to initialize this variable before the gather loop.

```vex
surface
sendray(vector background=0; int nsamples=16)
{
    float   amp;

    amp = 0;
    Cf = 0;

    gather(P, N,
        "bias", 0.01,
        "angle", radians(15),
        "samples", nsamples,
        "samplebase", 1,
        "distribution", "cosine",
        "send:exportvalue", 0.8,
        "amp", amp)
    {
        Cf += amp;
        amp = 0;
    }
    else
    {
        Cf += background;
    }

    Cf *= 1.0 / nsamples;
    Of = 1;
}

```

The following shader imports values which were sent to it by the gather loop. The exported parameter can be imported by the gather loop.

```vex
surface
hitsurf(export float amp=1)
{
    float   sendvalue = 0;

    if (!rayimport("exportvalue", sendvalue))
        printf("Error importing the send:exportvalue\n");

    amp = sendvalue;

    Cf = 1;
    Of = 1;
}

```
