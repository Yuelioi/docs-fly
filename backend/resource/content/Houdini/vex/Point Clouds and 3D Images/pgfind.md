---
title: pgfind
order: 34
---
`int [] pgfind(<geometry>geometry, vector P, float radius, int maxpoints, float divsize)`

`int [] pgfind(<geometry>geometry, string ptgroup, vector P, float radius, int maxpoints, float divsize)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

These functions are very similar to the `pcfind` functions. The difference is that they use a grid-based acceleration structure. This can provide faster initialization and lookups, provided the right grid size tuning parameter is used.

If you are searching a point cloud using an near-constant search radius, that radius can be used as the division size.

Note
The division size must be not vary per point.

Note
The division size is clamped above 3.0×105.

The `ptgroup` is a point group that limits the points to search. This is a SOP-style group pattern, so can be something like `0-10` or `@Cd.x>0.5`. A blank string is treated as matching all points.
