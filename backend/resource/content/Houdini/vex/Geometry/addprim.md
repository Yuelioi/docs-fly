---
title: addprim
order: 2
---
`int  addprim(int geohandle, string type)`

Create a polygon or polyline without any points. You can then add vertices to the primitive with [addvertex](addvertex.html "Adds a vertex to a primitive in a geometry.").

Make sure to add at least one vertex to the created primitive. While we attempt to make sure Houdini’s code can deal with empty polygons, it’s possible they could cause strange results or failures.

`int  addprim(int geohandle, string type, int pt0)`

`int  addprim(int geohandle, string type, int pt0, int pt1)`

`int  addprim(int geohandle, string type, int pt0, int pt1, int pt2)`

`int  addprim(int geohandle, string type, int pt0, int pt1, int pt2, int pt3)`

`int  addprim(int geohandle, string type, int pt0, int pt1, int pt2, int pt3, int pt4, int pt5, int pt6, int pt7)`

Create a primitive using points specified by point numbers.

`int  addprim(int geohandle, string type, int points[])`

Create a primitive using points specified in an array of point numbers.

`void  addprim(int &prim_num, int geohandle, string type, int pt0, int &vertices[])`

`void  addprim(int &prim_num, int geohandle, string type, int pt0, int pt1, int &vertices[])`

`void  addprim(int &prim_num, int geohandle, string type, int pt0, int pt1, int pt2, int &vertices[])`

`void  addprim(int &prim_num, int geohandle, string type, int pt0, int pt1, int pt2, int pt3, int &vertices[])`

`void  addprim(int &prim_num, int geohandle, string type, int pt0, int pt1, int pt2, int pt3, int pt4, int pt5, int pt6, int pt7, int &vertices[])`

`void  addprim(int &prim_num, int geohandle, string type, int points[], int &vertices[])`

These signatures fill the given `vertices` array with the new primitive’s vertex numbers corresponding to the given points. You can use these numbers with [setvertexattrib](setvertexattrib.html "Sets a vertex attribute in a geometry.") to set attributes on the vertices, however they may not be the final numbers of the vertices.

If the primitive was created, but any points were invalid, the corresponding vertex numbers in the array will be `-1`.

These signatures overwrite the `primnum` variable with the new primitive number instead of returning it.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`type`

One of the following strings:

| `"poly"` | Closed polygon. Can use 0 or more points. |
| --- | --- |
| `"polyline"` | Open polygon. Can use 0 or more points. |
| `"tet"` | Tetrahedron primitive. A triangular pyramid. Requires exactly 4 points. You cannot add vertices to this primitive. |
| “hex” | Hexahedron primitive. A distorted cube. Requires eactly 8 points. You cannot add vertices to this primitive. |
| `"sphere"`, `"circle"`, `"tube"`, `"metaball"`, `"metasquad"` | Sphere, circle, tube, metaball, or metasuperquadric primitive. The radius and shape are controled by transform primitive intrinsics. Require exactly 1 point. You cannot add vertices to these primitives. |
| `"channel"` | Channel Primitive. Stores channel data in a primitive. Can use 0 or more points. |
| `"AlembicRef"`, `"PackedDisk"` | Packed Alembic or packed disk primitive. Require exactly 1 point. You cannot add vertices to these primitives. |

Returns

A primitive number for the created primitive, or `-1` if the point could not be created.

You can use the return value with [setprimattrib](setprimattrib.html "Sets a primitive attribute in a geometry.") to set attributes on the new point, however it may not be the final number of the point.

You can set a primitive’s transforms using [setprimintrinsic](setprimintrinsic.html "Sets the value of a writeable primitive intrinsic attribute."), for example:

```vex
matrix3 transform_value = {{0, 0, 0}, {0, 0, 0}, {1, 1, 1}};
setprimintrinsic(geoself(), "transform", prim, transform_value);

```

You can also set Alembic and packed primitive intrinsics, for example:

```vex
setprimintrinsic(geoself(), "unexpandedfilename", prim, "test.bgeo");`

```
