---
title: eval_bsdf
order: 9
---
| On this page | * [Variadic arguments](#variadic-arguments) * [Examples](#examples) |
| --- | --- |

`vector  eval_bsdf(bsdf b, vector viewer, vector light, ...)`

`vector  eval_bsdf(bsdf b, vector viewer, vector light, int mask, ...)`

`vector  eval_bsdf(bsdf b, vector viewer, vector light, float &pdf, int mask, ...)`

`vector  eval_bsdf(bsdf b, vector viewer, vector light, vector normal, ...)`

`vector  eval_bsdf(bsdf b, vector viewer, vector light, vector normal, int mask, ...)`

`vector  eval_bsdf(bsdf b, vector viewer, vector light, vector normal, float &pdf, int mask, ...)`

Show/hide arguments

`b`

BSDF to evaluate.

`viewer`

Vector toward viewer.

`light`

Vector toward light.

`normal`

Surface normal.

`mask`

A bitmask indicating which types of shading component bounces to evaluate.

See [bouncemask](bouncemask.html) for information on component label bitmasks.

`&pdf`

The function overwrites this variable with the computed PDF for the given directions, scaled by the albedo.

Variadic arguments

## variadic-arguments

The `eval_bsdf` function passes any extra `"name", value` argument pairs to the BSDF being
evaluated. For custom BSDFs these keyword arguments are bound to shader
arguments (e.g. indicating whether the BSDF is being evaluated for direct or
indirect illumination). It’s also possible for a BSDF to pass information back
to `eval_bsdf`. To indicate that a keyword argument value should be imported
from the BSDF prefix the keyword with “import:”

Examples

## examples

```vex
v = eval_bsdf(F, inI, dir,
    "direct", 0,                        // Specify indirect illumination
    "import:sssmfp", sssmfp,        // Read the exported sssmfp parameter
    ...
);

```
