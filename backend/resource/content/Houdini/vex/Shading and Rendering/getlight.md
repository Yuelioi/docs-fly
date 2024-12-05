---
title: getlight
order: 16
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`light  getlight(int lid)`

Given a integer light identifier (lid), you can get a reference to a light
object representing the light (see details on the [mantra specific type](../lang.html#mantratypes))

```vex
int[] lights = getlights();
int nlights = len(lights);
for (int i = 0; i < nlights; i++)
{
light lp = getlight(i);
lp->illuminate(...);
}

```
