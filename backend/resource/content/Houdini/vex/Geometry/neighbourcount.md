---
title: neighbourcount
order: 15
---
`int  neighbourcount(<geometry>geometry, int point_num)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`point_num`

The number of the point whose neighbours you want to count.

Returns

The number of points that are connected to the specified point.
A point is connected if it is adjacent in some polygon, is one of the
four surrounding points in a grid or NURBs surface, or in some other
manner directly shares an edge with `point_num`. Returns 0 if there
is no input, or if the point number is out of range.
