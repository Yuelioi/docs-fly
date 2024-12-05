---
title: pcline
order: 23
---
| Since | 18.0 |
| --- | --- |

`int [] pcline(<geometry>geometry, string PChannel, vector P, vector dir, float max_distance, int maxpoints)`

`int [] pcline(<geometry>geometry, string ptgroup, string PChannel, vector P, vector dir, float max_distance, int maxpoints)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

These functions open a geometry file and return a list of points within max_distance of the line passing through P with direction dir.

The `ptgroup` is a point group that limits the points to search. This is a SOP-style group pattern, so can be something like `0-10` or `@Cd.x>0.5`. A blank string is treated as matching all points.
