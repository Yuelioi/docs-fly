---
title: Dv
order: 2
---
`float  Dv(float n, ...)`

`vector  Dv(vector n, ...)`

`vector4  Dv(vector4 n, ...)`

Show/hide arguments

Returns

The derivative of `n` with respect to V.

In shading contexts, this is the change in the variable over the area
being shaded.
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
