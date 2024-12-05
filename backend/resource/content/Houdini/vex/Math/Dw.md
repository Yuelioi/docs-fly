---
title: Dw
order: 3
---
`float  Dw(float p, ...)`

`vector  Dw(vector p, ...)`

`vector4  Dw(vector4 p, ...)`

Show/hide arguments

Returns

Returns the derivative of `p` with respect to W.

When rendering surfaces, this function returns 0.

In shading contexts, this is the change in the variable over the volume being shaded.
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
