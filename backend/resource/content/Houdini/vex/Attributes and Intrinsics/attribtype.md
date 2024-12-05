---
title: attribtype
order: 11
---
If you know the attribute class ahead of time, using [detailattribtype](detailattribtype.html "Returns the type of a geometry detail attribute."), [primattribtype](primattribtype.html "Returns the type of a geometry prim attribute."), [pointattribtype](pointattribtype.html "Returns the type of a geometry point attribute."), or [vertexattribtype](vertexattribtype.html "Returns the type of a geometry vertex attribute.") may be faster.

`int  attribtype(<geometry>geometry, string attribclass, string attribute_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

Returns

A numeric code indicating the attribute type:

| `-1` | Attribute not found, or unknown type. |
| --- | --- |
| `0` | Integer |
| `1` | Float or vector |
| `2` | String |
| `3` | Array of integers (or integer tuples) |
| `4` | Array of floats (or float tuples) |
| `5` | Array of strings. |
| `6` | Dictionary |
| `7` | Array of Dictionaries |

Examples

## examples

```vex
// Get the type of the position attribute of "defgeo.bgeo"
int type = attribtype("defgeo.bgeo", "point", "P");

```
