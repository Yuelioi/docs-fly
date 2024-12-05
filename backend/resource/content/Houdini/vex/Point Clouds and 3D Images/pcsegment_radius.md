---
title: pcsegment_radius
order: 30
---
| Since | 18.0 |
| --- | --- |

`int [] pcsegment_radius(<geometry>geometry, string PChannel, string RadChannel, float radscale, vector P0, vector P1, float max_distance, int maxpoints)`

`int [] pcsegment_radius(<geometry>geometry, string ptgroup, string PChannel, string RadChannel, float radscale, vector P0, vector P1, float max_distance, int maxpoints)`

`int [] pcsegment_radius(<geometry>geometry, string PChannel, string RadChannel, float radscale, vector P0, vector P1, float max_distance, int maxpoints, float &distances[])`

`int [] pcsegment_radius(<geometry>geometry, string ptgroup, string PChannel, string RadChannel, float radscale, vector P0, vector P1, float max_distance, int maxpoints, float &distances[])`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

These functions open a geometry file and return a list of points (treated as spheres) within max_distance of the line segment passing from P0 to P1. Each point is treated as a sphere with radius equal to its RadChannel attribute,

The `ptgroup` is a point group that limits the points to search. This is a SOP-style group pattern, so can be something like `0-10` or `@Cd.x>0.5`. A blank string is treated as matching all points.

The function also optionally takes a float array `distances`, which it modifies with the distances to each point.
