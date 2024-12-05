---
title: nearpoint
order: 11
---
`int  nearpoint(<geometry>geometry, vector pt)`

`int  nearpoint(<geometry>geometry, vector pt, float maxdist)`

`int  nearpoint(<geometry>geometry, string ptgroup, vector pt)`

`int  nearpoint(<geometry>geometry, string ptgroup, vector pt, float maxdist)`

Returns the number of the closest point on the geometry.
This will only search against points, not the surface locations
of the geometry. Use [xyzdist](xyzdist.html "Finds the distance from a point to the closest location on surface geometry.") to find the closest point on surfaces or curves.

-1 is returned if no point is found in the search distance.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`ptgroup`

A point group pattern to limit the search to. Can be a SOP-style group
pattern such as `0-10` or `@Cd.x>0.5`. An empty string will match all points.

`pt`

The position in space to find the closest point on the geometry to.

`maxdist`

The maximum distance to search. The operation can be sped up if it
is allowed to quit early.
