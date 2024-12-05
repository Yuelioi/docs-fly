---
title: surfacedist
order: 17
---
| Since | 17.0 |
| --- | --- |

`float  surfacedist(<geometry>geometry, string ptgroup, string P_attribute, int search_pt, int &closest_pt, string distance_metric)`

`float  surfacedist(<geometry>geometry, string ptgroup, string P_attribute, int search_pt, float max_radius, int &closest_pt, string distance_metric)`

Returns the distance from the search point to the closest point in the point
group.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`ptgroup`

The name of a point group or a pattern to generate a point
group. Uses the same semantics as a SOP group, so empty strings
will match all points. Attribute groups like `@Cd.x>0` can
also be used, but note that the `@` may need to be escaped with
a backslash in a [![](../../icons/COMMON/wrangle.svg)Snippet VOP](../../nodes/vop/snippet.html "Runs a VEX snippet to modify the incoming values.").

`P_attribute`

The name of the vector attribute to use to measure distance between
connected points. Using “P” will give the world distance along the surface,
but a custom attribute can be used to measure along a different metric.

`search_pt`

The point to measure the distance for.

`max_radius`

The maximum distance to measure the surface distance. This can speed things
up by allowing the search to quit early if the point is not within the
radius. Points outside the radius will return a value of `-1` for both the
distance and the lead point.

`&closest_pt`

Index of the closest point in the source group.

`-1` if no closest point was found.

`distance_metric`

The method to use to measure distance. Accepted values are `edge` and
`surface`. Edge distance is measured along the edges of the model, while
surface distance is measured along edges and across single polygons. Surface
distance is a better approximation of the true geodesic distance, but is
also more expensive to compute.

Returns

The distance from the search point to the closest point in the point group.

Returns `-1` if no closest point was found.
