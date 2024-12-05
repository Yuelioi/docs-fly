---
title: volume
order: 1
---
| On this page | * [Derivatives options](#derivatives-options) * [Examples](#examples) |
| --- | --- |

`float  volume(vector pos, ...)`

Note
This function relies on the fact that VEX “knows” that `pos` has
derivatives (`dPdu`, `dPdv`, and `dPdz`).
Passing a literal vector instead of a special variables such as `P`
will return `0` since VEX will not be able to access the derivatives.

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

Return the volume of the current micro-voxel in camera space:

```vex
volume(P)

```

Returns `0`, since the argument is not a variable VEX knows the derivatives for:

```vex
volume({0.1, 2.3, 4.5})

```
