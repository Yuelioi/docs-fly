---
title: nearpoints
order: 12
---
`int [] nearpoints(<geometry>geometry, vector pt, float maxdist)`

`int [] nearpoints(<geometry>geometry, vector pt, float maxdist, int maxpts)`

`int [] nearpoints(<geometry>geometry, string ptgroup, vector pt, float maxdist)`

`int [] nearpoints(<geometry>geometry, string ptgroup, vector pt, float maxdist, int maxpts)`

Show/hide arguments

`opinput`

The number of the input to the current node, starting with `0` being the first input.

`geometry`

The name of the geometry file to reference. Inside Houdini,
this may be “op:full_path_to_sop” to reference a SOP.

`ptgroup`

A point group pattern to limit the search to. Can be a SOP-style group
pattern such as `0-10` or `@Cd.x>0.5`. An empty string will match
all points.

`pt`

The position in space to find the closest point on the geometry to.

`maxdist`

The maximum distance to search.

`maxpts`

The maximum number of points to find.

Returns

An array of point numbers
This will only search against points, not the surface locations of the geometry.
Use [xyzdist](xyzdist.html "Finds the distance from a point to the closest location on surface geometry.") to find the closest point on surfaces or curves.
