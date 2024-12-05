---
title: Working with geometry groups in VEX
order: 9
---
| On this page | * [Dedicated functions](#dedicated-functions) * [Getting the existing groups](#getting-the-existing-groups) * [Reading groups as attributes](#reading-groups-as-attributes) |
| --- | --- |

Dedicated functions

## dedicated-functions

VEX has a number of functions for working with geometry groups:

| Get the number of elements in a group | * [nprimitivesgroup](functions/nprimitivesgroup.html "Returns the number of primitives in the group.") * [npointsgroup](functions/npointsgroup.html "Returns the number of points in the group.") * [nverticesgroup](functions/nverticesgroup.html "Returns the number of vertices in the group.") |
| --- | --- |
| Test if a given element number is in a group | * [inprimgroup](functions/inprimgroup.html "Returns 1 if the primitive specified by the primitive number is in the group specified by the string.") * [inpointgroup](functions/inpointgroup.html "Returns 1 if the point specified by the point number is in the group specified by the string.") * [invertexgroup](functions/invertexgroup.html "Returns 1 if the vertex specified by the vertex number is in the group specified by the string.") |
| Convert a group name or group syntax into a list of point/primitive numbers | * [expandprimgroup](functions/expandprimgroup.html "Returns an array of prim numbers corresponding to a group string.") * [expandpointgroup](functions/expandpointgroup.html "Returns an array of point numbers corresponding to a group string.") |
| Modify the contents of a named group | * [setprimgroup](functions/setprimgroup.html "Adds or removes a primitive to/from a group in a geometry.") * [setpointgroup](functions/setpointgroup.html "Adds or removes a point to/from a group in a geometry.") * [setvertexgroup](functions/setvertexgroup.html "Adds or removes a vertex to/from a group in a geometry.") |

Adding an element to a non-existent group creates the group.

Getting the existing groups

## getting-the-existing-groups

You can get an array of existing group names by reading the “intrinsic” detail attributes `edgegroups`, `pointgroups`, `primitivegroups`, or `vertexgroups`.

```vex
string groups[] = detailintrinsic(0, "pointgroups")
```

Reading groups as attributes

## reading-groups-as-attributes

The generic attribute reading function [attrib](functions/attrib.html "Reads the value of an attribute from geometry.") takes an `attribclass` argument that lets you specify what geometry level (detail, primitive, point, or vertex) to read the attribute from.

You can instead specify `"primgroup"`, `"pointgroup"`, or `"vertexgroup"` as the class to read group contents as if they were attributes.
