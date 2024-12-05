---
title: usd_addschemaattrib
order: 11
---
| Since | 20.0 |
| --- | --- |

`int  usd_addschemaattrib(int stagehandle, string primpath, string name, string typename)`

This function adds an attribute of a given type to the primitive. There are a few attributes which are considered part of a schema but which are not automatically added to a prim (such as the `extentsHint` attribute of the \`GeomModelAPI schema). In these rare cases, this function will add the attribute to a prim, marking it as a non-custom attribute. The data type of the attribute is not verified, so be sure to set the type expected by the schema. To create custom attributes, use [usd_addattrib](usd_addattrib.html "Creates an attribute of a given type on a primitive.").

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

Attribute name.

`typename`

The name or an alias of the type.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Adds a half-precision float attribute and sets its falue.
usd_applyapi(0, "/geo", "GeomModelAPI");
usd_addschemaattrib(0, "/geo", "extentsHint", "float[]");

```
