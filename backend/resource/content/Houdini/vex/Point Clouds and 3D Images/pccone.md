---
title: pccone
order: 5
---
| Since | 18.0 |
| --- | --- |

`int [] pccone(<geometry>geometry, string PChannel, vector P, vector dir, float angle, float max_distance, int maxpoints)`

`int [] pccone(<geometry>geometry, string ptgroup, string PChannel, vector P, vector dir, float angle, float max_distance, int maxpoints)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

These functions open a geometry file and return a list of points within the cone whose apex is P, opens in the vector direction dir, and with angle angle to dir. Additionally, it only returns the closest maxpoints points within distance max_distance of P.

The `ptgroup` is a point group that limits the points to search. This is a SOP-style group pattern, so can be something like `0-10` or `@Cd.x>0.5`. A blank string is treated as matching all points.
