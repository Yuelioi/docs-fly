---
title: pcfind
order: 11
---
`int [] pcfind(<geometry>geometry, string Pchannel, vector P, float radius, int maxpoints)`

`int [] pcfind(<geometry>geometry, string ptgroup, string Pchannel, vector P, float radius, int maxpoints)`

`int [] pcfind(<geometry>geometry, string Pchannel, vector P, float radius, int maxpoints, float &distances[])`

`int [] pcfind(<geometry>geometry, string ptgroup, string Pchannel, vector P, float radius, int maxpoints, float &distances[])`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

These functions open a geometry file and return a list of points with the location P within radius, based on point positions found in Pchannel. Only the maxpoints closest points within the given radius will be returned. The file name may use the `op:` syntax to reference SOP geometry in the OP contexts. The Pchannel parameter indicates the attribute which contains the positions to be searched.

The `ptgroup` is a point group that limits the points to search. This is a SOP-style group pattern, so can be something like `0-10` or `@Cd.x>0.5`. A blank string is treated as matching all points.

The function also optionally takes a float array `distances`, which it modifies with the distances to each point.

The closest point is in entry 0 of the returned array, and the other points are sorted by increasing distance.
Examples

## examples

Performing a proximity query:

```vex
int closept[] = pcfind(filename, "P", P, maxdistance, maxpoints);
P = 0;
foreach (int ptnum; closept)
{
    vector closepos = point(filename, "P", ptnum);
    P += closepos;
}
P /= len(closept);

```
