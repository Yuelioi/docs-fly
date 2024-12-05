---
title: minpos
order: 10
---
`vector  minpos(<geometry>geometry, vector point)`

Returns the position of the closest point in geometry to point.

`vector  minpos(<geometry>geometry, vector point, float maxdist)`

Returns the position of the closest point in geometry to point,
within the a radius of maxdist.
If it doesn’t find a point within maxdist, it returns point.

`vector  minpos(<geometry>geometry, string primgroup, vector point)`

Returns the position of the closest point in geometry to point,
limiting the search to primitives in the named primgroup.
If it doesn’t find a point on a primitive in the group, it returns point.

`vector  minpos(<geometry>geometry, string primgroup, vector point, float maxdist)`

Returns the position of the closest point in geometry to point,
limiting the search to primitives in the named primgroup,
within the a radius of maxdist.
If it doesn’t find a point on a primitive in the group or within maxdist,
it returns point.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`primgroup`

If specified, only report points on
This can be the name of a group on the geometry, or
[group specification syntax](../../model/groups.html#manual) like `@Cd.x>0`.
An empty string is ignored (matches all primitives).

Note
If you use group syntax with `@` in a Wrangle node snippet, you may need to
put a backslash (`\`) before the `@` to prevent the Wrangle node from trying
to interpret it.

`point`

The point (in world space) to find the closest point to on the geometry.

`maxdist`

The maximum distance to search.
This can speed up the function since it may allow quitting the search early.

Returns

The position of the nearest point on the geometry, or point if no nearest point was found.
