---
title: intersect_lights
order: 40
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [surface](../contexts/surface.html) |
| --- | --- |

Note
This function only works with area lights.

`int  intersect_lights(int lightids[], vector pos, vector dir, float time, int &idx, float &dist, vector &clr, float &scale, ...)`

Show/hide arguments

`lightids`

An array of light IDs, as returned by [getlights](getlights.html "Returns an array of light identifiers for the currently shaded surface.").

`pos`

The origin of the ray (such as the global variable `P`).

`dir`

Direction vector from the origin. The length of this vector does not affect
the distance the ray will travel.

`time`

Time to send the ray at.

The function modifies the values of the following arguments:

Show/hide arguments

`idx`

The light index for the light that was hit by the ray, or -1 if no intersection was found.

`dist`

The distance to the nearest intersected light.

`clr`

The light color set by the light shader.

`scale`

The light average hemispherical intensity (for area lights).

Returns

A [component bitmask](bouncemask.html) indicating what types of component bounces the light affects,
or `0` if the ray did not hit a light.
