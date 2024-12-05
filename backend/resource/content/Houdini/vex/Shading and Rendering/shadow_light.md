---
title: shadow_light
order: 71
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [surface](../contexts/surface.html) |
| --- | --- |

`vector  shadow_light(int lightid, vector pos, vector dir, float time, ...)`

This operation is similar to the shadow() function but it allows execution
of the shadow shader outside an illuminance loop. The position and
direction toward the light source are provided directly, and the shadow
shader is executed - returning the shadow multiplier. To produce the final
shadowed color, multiply the shaded color by the value returned by
shadow_light.

Keyword variadic arguments can be passed to the shadow shader, for import
in the shadow shader with with simport().

Show/hide arguments

`lightid`

A light identifier, as returned by [getlights](getlights.html "Returns an array of light identifiers for the currently shaded surface.").

`pos`

The origin of the ray (such as the global variable `P`).

`dir`

Direction vector from the origin. The length of this vector should be
the distance from *pos* to the light source.

`time`

Time to send the ray at.
