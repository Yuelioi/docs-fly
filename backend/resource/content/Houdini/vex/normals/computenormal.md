---
title: computenormal
order: 1
---
`vector  computenormal(vector P, ...)`

In shading contexts, computes the normal for position P using the cross product of the derivatives of P.

`vector  computenormal(vector P, vector N, vector Ng, ...)`

In shading contexts, computes the normal for position P using the cross product of the derivatives of P.
N is the original surface normal and Ng is the geometric normal.
This version “adjusts” the computed normal so interpolated normals will be relatively correct.

`void  computenormal(int i)`

(Obsolete) In SOP context, sets the hint for whether normals should be recomputed when `P` or `N` change (0=never, 1=automatic, 2=always).

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
