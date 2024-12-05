---
title: getderiv
order: 12
---
| On this page | * [Derivatives options](#derivatives-options) * [Examples](#examples) |
| --- | --- |
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |

`void  getderiv(float attr, string attrName, int isVertexAttr, float s, float t, float &du, float &dv, ...)`

`void  getderiv(<vector>attr, string attrName, int isVertexAttr, float s, float t, <vector>&du, <vector>&dv, ...)`

Note
If derivatives are queried for a polygonal mesh it is interally sampled as a Subdivision Surface.

Show/hide arguments

`attr`

Attribute value.

`attrName`

Name of attribute to evaluate.

`isVertexAttr`

Set to `1` to indicate the attribute is a vertex type.

`s`

Parametric S shading value. This should be passed from the `s` global variable.

`t`

Parametric <type> shading value. This should be passed from the `t` global variable.

`du`

Derivative of attribute in U direction.

`dv`

Derivative of attribute in V direction.

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

```vex
// Get derivatives of point attribute 'N'
vector dNdu, dNdv;
getderiv(N, "N", 0, s, t, dNdu, dNdv);

```
