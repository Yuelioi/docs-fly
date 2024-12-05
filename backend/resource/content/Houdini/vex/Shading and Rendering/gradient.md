---
title: gradient
order: 35
---
| On this page | * [Derivatives options](#derivatives-options) * [Examples](#examples) |
| --- | --- |

`vector  gradient(float val, ...)`

`vector  gradient(vector P, float val, ...)`

This method computes the derivative of a volume field using the partial
derivatives with respect to a given position (`Du`, `Dv`, and `Dw`). If no
position is provided, `P` is assumed in shading contexts. If only `Du` and
`Dv` are defined, the derivative will be tangent to the surface being
rendered.
Derivatives options

## derivatives-options

Functions which compute derivatives take additional arguments to
allow tuning of the derivative computation.

Show/hide arguments

"`extrapolate`",
`int`
`=0`

Whether derivatives are
“smooth” across patch boundaries. In most cases this is true and if
extrapolation is turned on, derivative computation should be exact
for C2 surfaces. However, when the VEX variables are changing with a
high frequency (for example, a high frequency displacement map
causing high frequency changes to the P variable), extrapolation of
derivative computation may cause exaggeration of discontinuities
between patch boundaries.

"`smooth`",
`int`
`=1`

Adjust the magnitude of the
differentials non-uniformly over patches. This will usually reduce
patch discontinuities in displacement/textured shaders. However, in
some odd cases you may want to turn this feature off.

```vex
N = computenormal(P, "extrapolate", 1, "smooth", 0);

```

Examples

## examples

Return the gradient of the density field:

```vex
surface test_grad(float density = 0)
{
Cf = gradient(density);
}

```
