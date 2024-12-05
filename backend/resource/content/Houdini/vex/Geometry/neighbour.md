---
title: neighbour
order: 14
---
This function lets you walk the points connected to a point (separated by a single edge). To get a list of all connected points at once, use [neighbours](neighbours.html "Returns an array of the point numbers of the neighbours of a point.").

`int  neighbour(<geometry>geometry, int point_num, int neighbour_num)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`point_num`

The number of the point whose neighbours you want to find.

`neighbour_num`

Which neighbour to find. The neighbours are in no particular order. Use [neighbourcount](neighbourcount.html "Returns the number of points that are connected to the specified point.") to get the total number of connected points.

Returns

The point index of the neighbour of the point. The order is undefined, but will be consistent for consistent geometry. Returns `-1` if the `neighbournum` is out of range for that point, or the point is out of range for that input, or the input doesn’t exist.
