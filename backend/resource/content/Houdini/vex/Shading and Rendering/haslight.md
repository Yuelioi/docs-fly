---
title: haslight
order: 36
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  haslight(material mat, vector P, int light, ...)`

Returns 1 if the given light is used to illuminate the material at the point
specified.

This function accepts PBR keyword arguments to specify sampling types. These
options may exclude lights which don’t match the desired sampling paths.

The PBR sampling keywords include:

Show/hide arguments

`label`

A string specifying a specific label. This keyword argument may be specified multiple times.

`direct`

Expects an 0 or 1 integer value which will limit lights based on indirect or direct contribution categories.

Examples

## examples

```vex
haslight(material(), P, light_num, "direct", 0);

```
