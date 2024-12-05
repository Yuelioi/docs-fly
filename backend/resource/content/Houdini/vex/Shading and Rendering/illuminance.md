---
title: illuminance
order: 37
---
| On this page | * [Overview](#overview) * [Light inclusion/exclusion options](#light-inclusion-exclusion-options) * [Sending information to the light’s shader](#sending-information-to-the-light-s-shader) * [Message passing](#message-passing) * [lightexport keyword argument](#lightexport-keyword-argument) |
| --- | --- |

Overview

## overview

// Need to indent this so the include below doesn’t go under this heading

```vex
illuminance(position, [axis], [angle], [light_typemask], [lightmask])
{
    // Here, Cl and L will be set to the value/direction for the
    // current light source.
    // To force the shadow shader to be called, use:
    // shadow(Cl);
}

```

The shadow shader is not called unless you explicitly call it. However, once the shadow shader has been called, the value of `Cl` will be changed for the duration of the surface shader. The shadow shader is automatically called when using any of the built-in lighting calls (e.g. [diffuse](diffuse.html "Returns a diffuse BSDF or computes diffuse shading."), [specular](specular.html "Returns a specular BSDF or computes specular shading."), [ambient](ambient.html "Returns the color of ambient light in the scene.")).

The default value for the axis is the surface normal. The default value for the angle is PI/2. The default value for the light mask is LIGHT_DIFFUSE|LIGHT_SPECULAR (please see shading.h for the light definitions).

The `illuminance` statement loops through all light sources for which `dot(L, axis) > cos(angle)`.

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

Sending information to the light’s shader

## sending-information-to-the-light-s-shader

You can give additional pairs of string/value arguments to `illuminance` to pass named values to each light’s shader. For example, to pass the value of the `N` variable as `orgN`:

```vex
illuminance (P, nf, M_PI/2, "orgN", N) {
...
}

```

In the light’s shader, you can receive the value from the illuminance loop with the [simport](simport.html "Imports a variable sent by a surface shader in an illuminance loop.") function.

```vex
vector orgN;
simport("orgN", orgN);

```

The `simport` function returns 1 if the import succeeds and 0 otherwise, so you can use it as the condition in an `if` statement.

Here’s a full example:

```vex
surface
exporter()
{
vector nf = frontface(normalize(N), I);
Cf = 0;
illuminance(P, nf, M_PI/2, "orgN", N)
{
Cf += Cl;
}

light
importer()
{
vector orgN;
if (!simport("orgN", orgN))
orgN = N;
// Use original N
Cl = orgN;
}

```

Message passing

## message-passing

Within the illuminance loop, you can retrieve values from the light shader
with the [limport](limport.html "Imports a variable from the light shader for the surface.") function.

The light shader can retrieve any “keyword” arguments passed to the illuminance
statement with the [simport](simport.html "Imports a variable sent by a surface shader in an illuminance loop.") function.

For example, to send down the vector variable `uv` to the light shader…

```vex
vector      uv = set(s, t, 0);
illuminance(P, dir, "uv", uv) { ... }

```

The light shader would be able to read this using…

```vex
vector      uv;
if (simport("uv", uv))
printf("Imported: %g from surface\n", uv);

```

lightexport keyword argument

## lightexport-keyword-argument

You can supply the extra string argument `"lightexport"` followed by
a string argument containing the name(s) of the exported variables to
assign within the loop.

In some shaders, multiple illuminance loops are used to define different
light contributions. The `lightexport` argument is useful in these cases
to specify which variables should be exported from the different
loops.

The `lightexport` value can be a space-separated list of wildcard
patterns. For example, `illuminance(pos, dir, "lightexport", "Front*")`
exports variables whose names start with `Front`.

```vex
surface
light_export_test(export vector diff=0;
export vector spec=0)
{
vector     nn = normalize(frontface(N, I));
vector     vv = -normalize(I);
vector     clr;

Cf = 0;
// This illuminance loop only exports to the "diff" variable
illuminance(P, nn, "lightexport", "diff")
{
clr = Cl * diffuseBRDF(normalize(L), nn);
Cf += clr;
diff = clr;
}
// This illuminance loop only exports to the "spec" variable
illuminance(P, nn, "lightexport", "spec")
{
clr = Cl * specularBRDF(normalize(L), nn, vv, 0.1);
Cf += clr;
spec = clr;
}

```
