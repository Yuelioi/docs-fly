---
title: addpoint
order: 1
---
`int  addpoint(int geohandle, int point_number)`

Creates a new point with all the attributes and group memberships of the point with the given point number.

`int  addpoint(int geohandle, vector pos)`

Creates a new point with the given position.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

Returns

A point number for the created point, or `-1` if the point could not be created.

You can use the return value with [setpointattrib](setpointattrib.html "Sets a point attribute in a geometry.") to set attributes on the new point, however it may not be the final number of the point.
