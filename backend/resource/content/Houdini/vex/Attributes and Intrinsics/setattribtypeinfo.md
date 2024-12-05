---
title: setattribtypeinfo
order: 64
---
`int  setattribtypeinfo(int geohandle, string attribclass, string name, string typeinfo)`

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

`name`

The name of the attribute for which to change the transformation info.

`typeinfo`

The meaning of the attribute, which is used by transform nodes to determine how to modify the attribute. It is one of:

| `"none"` | Don’t transform. |
| --- | --- |
| `"point"` | Apply scales, rotations, and transformations. |
| `"hpoint"` | Apply scales, rotations, and transformations to this vector4. |
| `"vector"` | Apply scales and rotations, but not transformations. |
| `"normal"` | Apply rotations, apply scales with inverse-transpose. |
| `"color"` | Don’t transform. |
| `"matrix"` | Apply scales, rotations, and transformations to this matrix. |
| `"quaternion"` | Apply rotations. |
| `"indexpair"` | Don’t transform. |
| `"integer"` | Do not blend this value when points are averaged. |
| `"integer-blend"` | Integer values that blend when points are averaged. |
| `"texturecoord"` | Don’t transform, and try to preserve seams when interpolating.  Attributes with this type will show up in the UV viewport menu. |
