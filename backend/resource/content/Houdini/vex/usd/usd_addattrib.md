---
title: usd_addattrib
order: 1
---
| Since | 18.0 |
| --- | --- |

`int  usd_addattrib(int stagehandle, string primpath, string name, string typename)`

This function adds an attribute of a given type to the primitive, if such attribute is not part of a schema. It is useful for controlling the exact type of a custom attribute. For attributes defined by primitive’s schema, this call has no effect, because the schema already determines their type. To create an attribute but mark it as not being a custm attribute, use [usd_addschemaattrib](usd_addschemaattrib.html "Creates an attribute of a given type on a primitive, and sets the custom metadata flag to False.").

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
usd_addattrib(0, "/geo/sphere", "half_attrib", "half3");
usd_setattrib(0, "/geo/sphere", "half_attrib", {1.25, 1.50, 1.75});

```
