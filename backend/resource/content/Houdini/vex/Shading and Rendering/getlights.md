---
title: getlights
order: 19
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int [] getlights(...)`

`int [] getlights(vector P, ...)`

`int [] getlights(material mat, vector P, ...)`

With this signature, the light mask is determined solely by the material (the
`lightmask` and `categories` keyword parameters are ignored).

This version
also accepts PBR keyword arguments to limit lights based on their Light
Contribution parameter.

Show/hide arguments

"label",
`string`

A specific label. This keyword argument may be specified multiple times.

"direct",
`int`

0 = limit lights to indirect contributions, 1 = limit lights to direct contributions.

Examples

## examples

```vex
getlights("lightmask", "light*,^light2");
getlights("categories", "shadow|occlusion");
getlights(material(), P, "direct", 0);

```
